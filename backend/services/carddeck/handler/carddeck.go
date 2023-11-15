package handler

import (
	"context"

	"github.com/kioku-project/kioku/pkg/converter"
	"github.com/kioku-project/kioku/pkg/helper"
	"github.com/kioku-project/kioku/pkg/model"
	pb "github.com/kioku-project/kioku/services/carddeck/proto"
	pbCollaboration "github.com/kioku-project/kioku/services/collaboration/proto"
	pbSrs "github.com/kioku-project/kioku/services/srs/proto"
	"github.com/kioku-project/kioku/store"
	"go-micro.dev/v4/logger"
	"golang.org/x/exp/slices"
)

type CardDeck struct {
	store                store.CardDeckStore
	collaborationService pbCollaboration.CollaborationService
	srsService           pbSrs.SrsService
}

func New(s store.CardDeckStore, cS pbCollaboration.CollaborationService, srsS pbSrs.SrsService) *CardDeck {
	return &CardDeck{store: s, collaborationService: cS, srsService: srsS}
}

func (e *CardDeck) checkUserRoleAccess(
	ctx context.Context,
	userID string,
	groupID string,
	requiredRole pbCollaboration.GroupRole,
) error {
	logger.Infof("Requesting group role for user (%s)", userID)
	roleRsp, err := e.collaborationService.GetGroupUserRole(ctx, &pbCollaboration.GroupRequest{
		UserID:  userID,
		GroupID: groupID,
	})
	if err != nil {
		return err
	}
	logger.Infof("Obtained group role (%s) for user (%s)", roleRsp.GroupRole.String(), userID)
	if !helper.IsAuthorized(roleRsp.GroupRole, requiredRole) {
		return helper.NewMicroNotAuthorizedErr(helper.CardDeckServiceID)
	}
	logger.Infof("Authenticated group role (%s) for user (%s)", roleRsp.GroupRole.String(), userID)
	return nil
}

func (e *CardDeck) checkUserDeckAccess(
	ctx context.Context,
	userID string,
	deckID string,
) error {
	e.store.FindDeckByID(deckID)
	deck, err := helper.FindStoreEntity(e.store.FindDeckByID, deckID, helper.CardDeckServiceID)
	if err != nil {
		return err
	}
	if deck.DeckType == model.PrivateDeckType {
		logger.Infof("Requesting group role for user (%s)", userID)
		if err = e.checkUserRoleAccess(ctx, userID, deck.GroupID, pbCollaboration.GroupRole_READ); err != nil {
			return err
		}
	}
	logger.Infof("Authenticated user (%s) for deck (%s)", userID, deckID)
	return nil
}

func (e *CardDeck) getCardSideAndCheckForValidAccess(
	ctx context.Context,
	userID string,
	cardSideID string,
) (*model.CardSide, error) {
	cardSide, err := helper.FindStoreEntity(ctx, e.store.FindCardSideByID, cardSideID, helper.CardDeckServiceID)
	if err != nil {
		return nil, err
	}
	deck, err := helper.FindStoreEntity(ctx, e.store.FindDeckByID, cardSide.Card.DeckID, helper.CardDeckServiceID)
	if err != nil {
		return nil, err
	}
	if err = e.checkUserRoleAccess(ctx, userID, deck.GroupID, pbCollaboration.GroupRole_WRITE); err != nil {
		return nil, err
	}
	return cardSide, nil
}

func (e *CardDeck) generateCardSidesForCard(ctx context.Context, card model.Card, sides []*pb.CardSideContent) error {
	var previousCardSide *model.CardSide
	var firstCardSideID string
	for _, side := range sides {
		newCardSide := &model.CardSide{
			CardID:      card.ID,
			Header:      side.Header,
			Description: side.Description,
		}
		if err := e.store.CreateCardSide(ctx, newCardSide); err != nil {
			return err
		}
		if previousCardSide != nil {
			previousCardSide.NextCardSideID = newCardSide.ID
			if err := e.store.ModifyCardSide(ctx, previousCardSide); err != nil {
				return err
			}
			newCardSide.PreviousCardSideID = previousCardSide.ID
			if err := e.store.ModifyCardSide(ctx, newCardSide); err != nil {
				return err
			}
		} else {
			firstCardSideID = newCardSide.ID
		}
		logger.Infof("Created card side %s and updated references", newCardSide.ID)
		previousCardSide = newCardSide
	}
	card.FirstCardSideID = firstCardSideID
	if err := e.store.ModifyCard(ctx, &card); err != nil {
		return err
	}
	logger.Infof("Modified first card reference in card %s", card.ID)
	return nil
}

func (e *CardDeck) updateCardSideReferences(ctx context.Context, cardSidesToUpdate [2]string, index int) error {
	side, err := helper.FindStoreEntity(
		ctx,
		e.store.FindCardSideByID,
		cardSidesToUpdate[index],
		helper.CollaborationServiceID,
	)
	if err != nil {
		return err
	}
	if index%2 == 0 {
		side.NextCardSideID = cardSidesToUpdate[1]
	} else {
		side.PreviousCardSideID = cardSidesToUpdate[0]
	}
	if err = e.store.ModifyCardSide(ctx, side); err != nil {
		return err
	}
	return nil
}

func (e *CardDeck) updateCardReferences(ctx context.Context, cardSideToDelete *model.CardSide) (bool, error) {
	isLastCardSide := false
	card, err := helper.FindStoreEntity(
		ctx,
		e.store.FindCardByID,
		cardSideToDelete.CardID,
		helper.CollaborationServiceID,
	)
	if err != nil {
		return false, err
	}
	if cardSideToDelete.NextCardSideID == "" {
		isLastCardSide = true
		logger.Infof("Card side %s is the last from card %s - will delete card",
			card.ID, cardSideToDelete.CardID)
		if err = e.store.DeleteCard(ctx, card); err != nil {
			return false, err
		}
	} else {
		card.FirstCardSideID = cardSideToDelete.NextCardSideID
		if err = e.store.ModifyCard(ctx, card); err != nil {
			return false, err
		}
	}
	return isLastCardSide, nil
}

func cardModelDateComparator(a, b model.Card) int {
	return a.CreatedAt.Compare(b.CreatedAt)
}

func (e *CardDeck) GetGroupDecks(ctx context.Context, req *pb.GroupDecksRequest, rsp *pb.GroupDecksResponse) error {
	logger.Infof("Received CardDeck.GetGroupDecks request: %v", req)

	var decks []model.Deck
	err := e.checkUserRoleAccess(ctx, req.UserID, req.GroupID, pbCollaboration.GroupRole_INVITED)
	if err != nil {
		decks, err = helper.FindStoreEntity(ctx, e.store.FindPublicDecksByGroupID, req.GroupID, helper.CardDeckServiceID)
		if err != nil {
			return err
		}
	} else {
		decks, err = helper.FindStoreEntity(ctx, e.store.FindDecksByGroupID, req.GroupID, helper.CardDeckServiceID)
		if err != nil {
			return err
		}
	}

	rsp.Decks = converter.ConvertToTypeArray(decks, converter.StoreDeckToProtoDeckConverter)
	logger.Infof("Found %d decks in group with id %s", len(decks), req.GroupID)
	return nil
}

func (e *CardDeck) CreateDeck(ctx context.Context, req *pb.CreateDeckRequest, rsp *pb.IDResponse) error {
	logger.Infof("Received CardDeck.CreateDeck request: %v", req)
	if err := e.checkUserRoleAccess(ctx, req.UserID, req.GroupID, pbCollaboration.GroupRole_WRITE); err != nil {
		return err
	}
	if err := helper.CheckForValidName(req.DeckName, helper.GroupAndDeckNameRegex, helper.UserServiceID); err != nil {
		return err
	}
	err, dt := converter.MigrateProtoDeckTypeToModelDeckType(req.DeckType)
	if err != nil {
		return err
	}
	newDeck := model.Deck{
		Name:     req.DeckName,
		GroupID:  req.GroupID,
		DeckType: dt,
	}
	if err := e.store.CreateDeck(ctx, &newDeck); err != nil {
		return err
	}
	rsp.ID = newDeck.ID
	logger.Infof("Successfully created new deck (%s) in group (%s)", req.GroupID, newDeck.ID)
	return nil
}

func (e *CardDeck) GetDeck(ctx context.Context, req *pb.IDRequest, rsp *pb.DeckResponse) error {
	logger.Infof("Received CardDeck.GetDeck request: %v", req)
	deck, err := helper.FindStoreEntity(ctx, e.store.FindDeckByID, req.EntityID, helper.CardDeckServiceID)
	if err != nil {
		return err
	}
	if err := e.checkUserDeckAccess(ctx, req.UserID, deck.ID); err != nil {
		return err
	}
	*rsp = *converter.StoreDeckToProtoDeckResponseConverter(*deck)
	logger.Infof("Successfully got information for deck %s", req.EntityID)
	return nil
}

func (e *CardDeck) ModifyDeck(ctx context.Context, req *pb.ModifyDeckRequest, rsp *pb.SuccessResponse) error {
	logger.Infof("Received CardDeck.ModifyCard request: %v", req)
	deck, err := helper.FindStoreEntity(ctx, e.store.FindDeckByID, req.DeckID, helper.CardDeckServiceID)
	if err != nil {
		return err
	}
	if err := e.checkUserRoleAccess(ctx, req.UserID, deck.GroupID, pbCollaboration.GroupRole_WRITE); err != nil {
		return err
	}
	if req.DeckName != nil {
		err := helper.CheckForValidName(*req.DeckName, helper.GroupAndDeckNameRegex, helper.UserServiceID)
		if err != nil {
			return err
		}
		deck.Name = *req.DeckName
	}
	if req.DeckType != nil {
		err, dt := converter.MigrateProtoDeckTypeToModelDeckType(*req.DeckType)
		if err != nil {
			return err
		}
		deck.DeckType = dt
	}
	err = e.store.ModifyDeck(ctx, deck)
	if err != nil {
		return err
	}
	rsp.Success = true
	logger.Infof("Successfully modified deck (%s) in group (%s)", req.DeckID, deck.GroupID)
	return nil
}

func (e *CardDeck) DeleteDeck(ctx context.Context, req *pb.IDRequest, rsp *pb.SuccessResponse) error {
	logger.Infof("Received CardDeck.DeleteDeck request: %v", req)
	deck, err := helper.FindStoreEntity(ctx, e.store.FindDeckByID, req.EntityID, helper.CardDeckServiceID)
	if err != nil {
		return err
	}
	if err := e.checkUserRoleAccess(ctx, req.UserID, deck.GroupID, pbCollaboration.GroupRole_ADMIN); err != nil {
		return err
	}
	err = e.store.DeleteDeck(ctx, deck)
	if err != nil {
		return err
	}
	rsp.Success = true
	logger.Infof("Successfully deleted deck (%s) in group (%s)", req.EntityID, deck.GroupID)
	return nil
}

func (e *CardDeck) GetDeckCards(ctx context.Context, req *pb.IDRequest, rsp *pb.DeckCardsResponse) error {
	logger.Infof("Received CardDeck.GetDeckCards request: %v", req)
	deck, err := helper.FindStoreEntity(ctx, e.store.FindDeckByID, req.EntityID, helper.CardDeckServiceID)
	if err != nil {
		return err
	}
	if err := e.checkUserDeckAccess(ctx, req.UserID, deck.ID); err != nil {
		return err
	}
	slices.SortFunc(deck.Cards, cardModelDateComparator)
	rsp.Cards = make([]*pb.Card, len(deck.Cards))
	for i, card := range deck.Cards {
		cardSides, err := e.store.FindCardSidesByCardID(ctx, card.ID)
		if err != nil {
			return err
		}
		rsp.Cards[i] = &pb.Card{
			CardID: card.ID,
			Sides:  converter.ConvertToTypeArray(cardSides, converter.StoreCardSideToProtoCardSideConverter),
		}
	}
	logger.Infof("Found %d cards in deck with id %s", len(deck.Cards), req.EntityID)
	return nil
}

func (e *CardDeck) CreateCard(ctx context.Context, req *pb.CreateCardRequest, rsp *pb.IDResponse) error {
	logger.Infof("Received CardDeck.CreateCard request: %v", req)
	deck, err := helper.FindStoreEntity(ctx, e.store.FindDeckByID, req.DeckID, helper.CardDeckServiceID)
	if err != nil {
		return err
	}
	if err := e.checkUserRoleAccess(ctx, req.UserID, deck.GroupID, pbCollaboration.GroupRole_WRITE); err != nil {
		return err
	}
	newCard := model.Card{
		DeckID: deck.ID,
	}
	if err = e.store.CreateCard(ctx, &newCard); err != nil {
		return err
	}
	if err = e.generateCardSidesForCard(ctx, newCard, req.Sides); err != nil {
		return err
	}

	///// add usercardbindings
	membersResp, err := e.collaborationService.GetGroupMembers(ctx, &pbCollaboration.GroupRequest{
		UserID:  req.UserID,
		GroupID: deck.GroupID,
	})
	if err != nil {
		return err
	}
	for _, user := range membersResp.Users {
		if _, err = e.srsService.AddUserCardBinding(ctx, &pbSrs.BindingRequest{
			UserID: user.User.UserID,
			CardID: newCard.ID,
			DeckID: newCard.DeckID,
		}); err != nil {
			return err
		}
	}
	rsp.ID = newCard.ID
	logger.Infof("Successfully created new card (%s) in deck (%s)", newCard.ID, req.DeckID)
	return nil
}

func (e *CardDeck) GetCard(ctx context.Context, req *pb.IDRequest, rsp *pb.Card) error {
	logger.Infof("Received CardDeck.GetCard request: %v", req)
	card, err := helper.FindStoreEntity(ctx, e.store.FindCardByID, req.EntityID, helper.CardDeckServiceID)
	if err != nil {
		return err
	}
	cardSides, err := helper.FindStoreEntity(ctx, e.store.FindCardSidesByCardID, req.EntityID, helper.CardDeckServiceID)
	if err != nil {
		return err
	}
	card.CardSides = cardSides
	if err := e.checkUserRoleAccess(ctx, req.UserID, card.Deck.GroupID, pbCollaboration.GroupRole_INVITED); err != nil {
		return err
	}
	*rsp = *converter.StoreCardToProtoCardConverter(*card)
	logger.Infof("Successfully got information for card %s", req.EntityID)
	return nil
}

func (e *CardDeck) ModifyCard(ctx context.Context, req *pb.ModifyCardRequest, rsp *pb.SuccessResponse) error {
	logger.Infof("Received CardDeck.ModifyCard request: %v", req)
	card, err := helper.FindStoreEntity(ctx, e.store.FindCardByID, req.CardID, helper.CardDeckServiceID)
	if err != nil {
		return err
	}
	if err := e.checkUserRoleAccess(ctx, req.UserID, card.Deck.GroupID, pbCollaboration.GroupRole_WRITE); err != nil {
		return err
	}
	if err = e.store.DeleteCardSidesOfCardByID(ctx, card.ID); err != nil {
		return err
	}
	if err = e.generateCardSidesForCard(ctx, *card, req.Sides); err != nil {
		return err
	}
	rsp.Success = true
	logger.Infof("Successfully modified card (%s) in deck (%s)", card.ID, card.DeckID)
	return nil
}

func (e *CardDeck) DeleteCard(ctx context.Context, req *pb.IDRequest, rsp *pb.SuccessResponse) error {
	logger.Infof("Received CardDeck.DeleteCard request: %v", req)
	card, err := helper.FindStoreEntity(ctx, e.store.FindCardByID, req.EntityID, helper.CardDeckServiceID)
	if err != nil {
		return err
	}
	if err := e.checkUserRoleAccess(ctx, req.UserID, card.Deck.GroupID, pbCollaboration.GroupRole_WRITE); err != nil {
		return err
	}
	err = e.store.DeleteCard(ctx, card)
	if err != nil {
		return err
	}
	rsp.Success = true
	logger.Infof("Successfully deleted card (%s) in deck (%s)", req.EntityID, card.DeckID)
	return nil
}

func (e *CardDeck) CreateCardSide(ctx context.Context, req *pb.CreateCardSideRequest, rsp *pb.IDResponse) error {
	logger.Infof("Received CardDeck.CreateCardSide request: %v", req)
	card, err := helper.FindStoreEntity(ctx, e.store.FindCardByID, req.CardID, helper.CardDeckServiceID)
	if err != nil {
		return err
	}
	if err := e.checkUserRoleAccess(ctx, req.UserID, card.Deck.GroupID, pbCollaboration.GroupRole_WRITE); err != nil {
		return err
	}
	var previousCardSideIDForNewCardSide string
	var newPreviousCard, newNextCard *model.CardSide
	if req.PlaceBeforeCardSideID != "" {
		newNextCard, err = helper.FindStoreEntity(
			ctx,
			e.store.FindCardSideByID,
			req.PlaceBeforeCardSideID,
			helper.CardDeckServiceID,
		)
		if err != nil {
			return err
		}
		if newNextCard.CardID != card.ID {
			return helper.NewMicroCardSideNotInGivenCardErr(helper.CardDeckServiceID)
		}
		if newNextCard.PreviousCardSideID != "" {
			if newPreviousCard, err = helper.FindStoreEntity(
				ctx,
				e.store.FindCardSideByID,
				newNextCard.PreviousCardSideID,
				helper.CardDeckServiceID,
			); err != nil {
				return err
			}
		}
		previousCardSideIDForNewCardSide = newNextCard.PreviousCardSideID
	} else {
		if card.FirstCardSideID != "" {
			newPreviousCard, err = helper.FindStoreEntity(
				ctx,
				e.store.FindLastCardSideOfCardByID,
				req.CardID,
				helper.CardDeckServiceID,
			)
			if err != nil {
				return err
			}
			previousCardSideIDForNewCardSide = newPreviousCard.ID
		}
	}
	newCardSide := model.CardSide{
		CardID:             card.ID,
		Header:             req.Content.Header,
		Description:        req.Content.Description,
		PreviousCardSideID: previousCardSideIDForNewCardSide,
		NextCardSideID:     req.PlaceBeforeCardSideID,
	}
	err = e.store.CreateCardSide(ctx, &newCardSide)
	if err != nil {
		return err
	}
	if newPreviousCard != nil {
		newPreviousCard.NextCardSideID = newCardSide.ID
		err = e.store.ModifyCardSide(ctx, newPreviousCard)
		if err != nil {
			return err
		}
	} else {
		card.FirstCardSideID = newCardSide.ID
		err = e.store.ModifyCard(ctx, card)
		if err != nil {
			return err
		}
	}
	if newNextCard != nil {
		newNextCard.PreviousCardSideID = newCardSide.ID
		err = e.store.ModifyCardSide(ctx, newNextCard)
		if err != nil {
			return err
		}
	}
	rsp.ID = newCardSide.ID
	logger.Infof("Successfully created card side (%s) in card (%s)", newCardSide.ID, card.ID)
	return nil
}

func (e *CardDeck) ModifyCardSide(ctx context.Context, req *pb.ModifyCardSideRequest, rsp *pb.SuccessResponse) error {
	logger.Infof("Received CardDeck.ModifyCardSide request: %v", req)
	cardSide, err := e.getCardSideAndCheckForValidAccess(ctx, req.UserID, req.CardSideID)
	if err != nil {
		return err
	}
	if req.Content != nil {
		cardSide.Header = req.Content.Header
		cardSide.Description = req.Content.Description
	}
	err = e.store.ModifyCardSide(ctx, cardSide)
	if err != nil {
		return err
	}
	rsp.Success = true
	logger.Infof("Successfully modified card side %s of card %s", req.CardSideID, cardSide.CardID)
	return nil
}

func (e *CardDeck) DeleteCardSide(ctx context.Context, req *pb.IDRequest, rsp *pb.SuccessResponse) error {
	logger.Infof("Received CardDeck.DeleteCardSide request: %v", req)
	cardSideToDelete, err := e.getCardSideAndCheckForValidAccess(ctx, req.UserID, req.EntityID)
	if err != nil {
		return err
	}
	isLastCardSide := false
	cardSidesToUpdate := [2]string{cardSideToDelete.PreviousCardSideID, cardSideToDelete.NextCardSideID}
	for index, cardSideToUpdate := range cardSidesToUpdate {
		if cardSideToUpdate != "" {
			err = e.updateCardSideReferences(ctx, cardSidesToUpdate, index)
			if err != nil {
				return err
			}
		} else if index%2 == 0 {
			isLastCardSide, err = e.updateCardReferences(ctx, cardSideToDelete)
			if err != nil {
				return err
			}
		}
	}
	if !isLastCardSide {
		if err = e.store.DeleteCardSide(ctx, cardSideToDelete); err != nil {
			return err
		}
	}
	rsp.Success = true
	logger.Infof("Successfully deleted card side %s of card %s", req.EntityID, cardSideToDelete.CardID)
	return nil
}
