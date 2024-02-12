package handler

import (
	"context"
	"errors"

	pbCommon "github.com/kioku-project/kioku/pkg/proto"
	pbCardDeck "github.com/kioku-project/kioku/services/carddeck/proto"
	"gorm.io/gorm"

	"github.com/kioku-project/kioku/pkg/converter"
	"github.com/kioku-project/kioku/pkg/helper"
	"github.com/kioku-project/kioku/pkg/model"
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
	requiredRole pbCommon.GroupRole,
) error {
	logger.Infof("Requesting group role for user (%s)", userID)
	roleRsp, err := e.collaborationService.GetGroupUserRole(ctx, &pbCommon.GroupRequest{
		UserID: userID,
		Group: &pbCommon.Group{
			GroupID: groupID,
		},
	})
	if err != nil {
		return err
	}
	logger.Infof("Obtained group role (%s) for user (%s)", roleRsp.Role.String(), userID)
	if !helper.IsAuthorized(roleRsp.Role, requiredRole) {
		return helper.NewMicroNotAuthorizedErr(helper.CardDeckServiceID)
	}
	logger.Infof("Authenticated group role (%s) for user (%s)", roleRsp.Role.String(), userID)
	return nil
}

func (e *CardDeck) checkUserDeckAccess(
	ctx context.Context,
	userID string,
	deckID string,
) error {
	deck, err := e.store.FindDeckByID(ctx, deckID, userID)
	if err != nil {
		return err
	}
	if deck.DeckType == model.PrivateDeckType {
		logger.Infof("Requesting group role for user (%s)", userID)
		if err = e.checkUserRoleAccess(ctx, userID, deck.GroupID, pbCommon.GroupRole_READ); err != nil {
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
	deck, err := e.store.FindDeckByID(ctx, cardSide.Card.DeckID, userID)
	if err != nil {
		return nil, err
	}
	if err = e.checkUserRoleAccess(ctx, userID, deck.GroupID, pbCommon.GroupRole_WRITE); err != nil {
		return nil, err
	}
	return cardSide, nil
}

func (e *CardDeck) generateCardSidesForCard(ctx context.Context, card model.Card, sides []*pbCommon.CardSide) error {
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

func (e *CardDeck) copyCards(ctx context.Context, cards []*model.Card, deckID string) error {
	for _, card := range cards {
		newCard := &model.Card{
			DeckID: deckID,
		}
		if err := e.store.CreateCard(ctx, newCard); err != nil {
			return err
		}
		cardSides, err := e.store.FindCardSidesByCardID(ctx, card.ID)
		if err != nil {
			return err
		}
		pbCardSides := make([]*pbCommon.CardSide, 0, len(cardSides))
		for _, cardSide := range cardSides {
			pbCardSides = append(pbCardSides, &pbCommon.CardSide{
				Header:      cardSide.Header,
				Description: cardSide.Description,
			})
		}
		if err := e.generateCardSidesForCard(ctx, *newCard, pbCardSides); err != nil {
			return err
		}
	}
	return nil
}

func (e *CardDeck) GetGroupDecks(ctx context.Context, req *pbCommon.GroupRequest, rsp *pbCommon.Decks) error {
	logger.Infof("Received CardDeck.GetGroupDecks request: %v", req)

	var decks []model.Deck
	err := e.checkUserRoleAccess(ctx, req.UserID, req.Group.GroupID, pbCommon.GroupRole_INVITED)
	if err != nil {
		decks, err = helper.FindStoreEntity(ctx, e.store.FindPublicDecksByGroupID, req.Group.GroupID, helper.CardDeckServiceID)
		if err != nil {
			return err
		}
	} else {
		decks, err = e.store.FindDecksByGroupID(ctx, req.Group.GroupID, req.UserID)
		if err != nil {
			return err
		}
	}

	rsp.Decks = converter.ConvertToTypeArray(decks, converter.StoreDeckToProtoDeckConverter)
	for _, deck := range rsp.Decks {
		deckRole, err := e.collaborationService.GetGroupUserRole(ctx, &pbCommon.GroupRequest{
			UserID: req.UserID,
			Group: &pbCommon.Group{
				GroupID: deck.GroupID,
			},
		})
		if err != nil {
			return err
		}
		deck.DeckRole = deckRole.Role
	}
	logger.Infof("Found %d decks in group with id %s", len(decks), req.Group.GroupID)
	return nil
}

func (e *CardDeck) CreateDeck(ctx context.Context, req *pbCommon.DeckRequest, rsp *pbCommon.Deck) error {
	logger.Infof("Received CardDeck.CreateDeck request: %v", req)
	if err := e.checkUserRoleAccess(ctx, req.UserID, req.Deck.GroupID, pbCommon.GroupRole_WRITE); err != nil {
		return err
	}
	if err := helper.CheckForValidName(req.Deck.DeckName, helper.GroupAndDeckNameRegex, helper.UserServiceID); err != nil {
		return err
	}
	dt, err := converter.MigrateProtoDeckTypeToModelDeckType(req.Deck.DeckType)
	if err != nil {
		return err
	}
	newDeck := model.Deck{
		Name:        req.Deck.DeckName,
		GroupID:     req.Deck.GroupID,
		DeckType:    dt,
		Description: req.Deck.DeckDescription,
	}
	if err := e.store.CreateDeck(ctx, &newDeck); err != nil {
		return err
	}
	rsp.DeckID = newDeck.ID
	logger.Infof("Successfully created new deck (%s) in group (%s)", req.Deck.GroupID, newDeck.ID)
	return nil
}

func (e *CardDeck) CopyDeck(ctx context.Context, req *pbCardDeck.CopyDeckRequest, rsp *pbCommon.Deck) error {
	logger.Infof("Received CardDeck.CopyDeck request: %v", req)
	if err := e.checkUserDeckAccess(ctx, req.UserID, req.Deck.DeckID); err != nil {
		return err
	}
	if err := e.checkUserRoleAccess(ctx, req.UserID, req.TargetGroupID, pbCommon.GroupRole_WRITE); err != nil {
		return err
	}
	var (
		deckType model.DeckType
		err      error
	)
	if req.NewDeck.DeckType != pbCommon.DeckType_DT_INVALID {
		deckType, err = converter.MigrateProtoDeckTypeToModelDeckType(req.Deck.DeckType)
		if err != nil {
			return err
		}
	} else {
		deckType = model.PrivateDeckType
	}

	newDeck := &model.Deck{
		GroupID:  req.TargetGroupID,
		Name:     req.NewDeck.DeckName,
		DeckType: deckType,
	}
	if err := e.store.CreateDeck(ctx, newDeck); err != nil {
		return err
	}
	cards, err := e.store.FindDeckCards(ctx, req.Deck.DeckID)
	if err != nil {
		return err
	}
	if err := e.copyCards(ctx, cards, newDeck.ID); err != nil {
		return err
	}
	rsp.DeckID = newDeck.ID
	return nil
}

func (e *CardDeck) GetDeck(ctx context.Context, req *pbCommon.DeckRequest, rsp *pbCommon.Deck) error {
	logger.Infof("Received CardDeck.GetDeck request: %v", req)
	deck, err := e.store.FindDeckByID(ctx, req.Deck.DeckID, req.UserID)
	if err != nil {
		return err
	}
	if err = e.checkUserDeckAccess(ctx, req.UserID, deck.ID); err != nil {
		return err
	}
	*rsp = *converter.StoreDeckToProtoDeckConverter(*deck)
	role, err := e.collaborationService.GetGroupUserRole(ctx, &pbCommon.GroupRequest{
		UserID: req.UserID,
		Group: &pbCommon.Group{
			GroupID: deck.GroupID,
		},
	})
	if err != nil {
		return err
	}
	rsp.DeckRole = role.Role
	logger.Infof("Successfully got information for deck %s", req.Deck.DeckID)
	return nil
}

func (e *CardDeck) ModifyDeck(ctx context.Context, req *pbCommon.DeckRequest, rsp *pbCommon.Success) error {
	logger.Infof("Received CardDeck.ModifyDeck request: %v", req)
	deck, err := e.store.FindDeckByID(ctx, req.Deck.DeckID, req.UserID)
	if err != nil {
		return err
	}
	if err := e.checkUserRoleAccess(ctx, req.UserID, deck.GroupID, pbCommon.GroupRole_WRITE); err != nil {
		return err
	}
	if req.Deck.DeckName != "" {
		err := helper.CheckForValidName(req.Deck.DeckName, helper.GroupAndDeckNameRegex, helper.UserServiceID)
		if err != nil {
			return err
		}
		deck.Name = req.Deck.DeckName
	}
	if req.Deck.DeckDescription != "" {
		err := helper.CheckForValidName(req.Deck.DeckDescription, helper.GroupAndDeckNameRegex, helper.UserServiceID)
		if err != nil {
			return err
		}
		deck.Description = req.Deck.DeckDescription
	}
	if req.Deck.DeckType != pbCommon.DeckType_DT_INVALID {
		dt, err := converter.MigrateProtoDeckTypeToModelDeckType(req.Deck.DeckType)
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
	logger.Infof("Successfully modified deck (%s) in group (%s)", req.Deck.DeckID, deck.GroupID)
	return nil
}

func (e *CardDeck) DeleteDeck(ctx context.Context, req *pbCommon.DeckRequest, rsp *pbCommon.Success) error {
	logger.Infof("Received CardDeck.DeleteDeck request: %v", req)
	deck, err := e.store.FindDeckByID(ctx, req.Deck.DeckID, req.UserID)
	if err != nil {
		return err
	}
	if err := e.checkUserRoleAccess(ctx, req.UserID, deck.GroupID, pbCommon.GroupRole_ADMIN); err != nil {
		return err
	}
	err = e.store.DeleteDeck(ctx, deck)
	if err != nil {
		return err
	}
	rsp.Success = true
	logger.Infof("Successfully deleted deck (%s) in group (%s)", req.Deck.DeckID, deck.GroupID)
	return nil
}

func (e *CardDeck) GetDeckCards(ctx context.Context, req *pbCommon.DeckRequest, rsp *pbCommon.Cards) error {
	logger.Infof("Received CardDeck.GetDeckCards request: %v", req)
	deck, err := e.store.FindDeckByID(ctx, req.Deck.DeckID, req.UserID)
	if err != nil {
		return err
	}
	if err := e.checkUserDeckAccess(ctx, req.UserID, deck.ID); err != nil {
		return err
	}
	slices.SortFunc(deck.Cards, cardModelDateComparator)
	rsp.Cards = make([]*pbCommon.Card, len(deck.Cards))
	for i, card := range deck.Cards {
		cardSides, err := e.store.FindCardSidesByCardID(ctx, card.ID)
		if err != nil {
			return err
		}
		rsp.Cards[i] = &pbCommon.Card{
			CardID: card.ID,
			Sides:  converter.ConvertToTypeArray(cardSides, converter.StoreCardSideToProtoCardSideConverter),
		}
	}
	logger.Infof("Found %d cards in deck with id %s", len(deck.Cards), req.Deck.DeckID)
	return nil
}

func (e *CardDeck) CreateCard(ctx context.Context, req *pbCommon.CardRequest, rsp *pbCommon.Card) error {
	logger.Infof("Received CardDeck.CreateCard request: %v", req)
	deck, err := e.store.FindDeckByID(ctx, req.Card.DeckID, req.UserID)
	if err != nil {
		return err
	}
	if err := e.checkUserRoleAccess(ctx, req.UserID, deck.GroupID, pbCommon.GroupRole_WRITE); err != nil {
		return err
	}
	newCard := model.Card{
		DeckID: deck.ID,
	}
	if err = e.store.CreateCard(ctx, &newCard); err != nil {
		return err
	}
	if err = e.generateCardSidesForCard(ctx, newCard, req.Card.Sides); err != nil {
		return err
	}

	///// add usercardbindings
	membersResp, err := e.collaborationService.GetGroupMembers(ctx, &pbCommon.GroupRequest{
		UserID: req.UserID,
		Group: &pbCommon.Group{
			GroupID: deck.GroupID,
		},
	})
	if err != nil {
		return err
	}
	for _, user := range membersResp.Users {
		if _, err = e.srsService.AddUserCardBinding(ctx, &pbCommon.BindingRequest{
			UserID: user.UserID,
			CardID: newCard.ID,
			DeckID: newCard.DeckID,
		}); err != nil {
			return err
		}
	}
	rsp.CardID = newCard.ID
	logger.Infof("Successfully created new card (%s) in deck (%s)", newCard.ID, req.Card.DeckID)
	return nil
}

func (e *CardDeck) GetCard(ctx context.Context, req *pbCommon.CardRequest, rsp *pbCommon.Card) error {
	logger.Infof("Received CardDeck.GetCard request: %v", req)
	card, err := helper.FindStoreEntity(ctx, e.store.FindCardByID, req.Card.CardID, helper.CardDeckServiceID)
	if err != nil {
		return err
	}
	cardSides, err := helper.FindStoreEntity(ctx, e.store.FindCardSidesByCardID, req.Card.CardID, helper.CardDeckServiceID)
	if err != nil {
		return err
	}
	card.CardSides = cardSides
	if err := e.checkUserRoleAccess(ctx, req.UserID, card.Deck.GroupID, pbCommon.GroupRole_INVITED); err != nil {
		return err
	}
	*rsp = *converter.StoreCardToProtoCardConverter(*card)
	logger.Infof("Successfully got information for card %s", req.Card.CardID)
	return nil
}

func (e *CardDeck) ModifyCard(ctx context.Context, req *pbCommon.CardRequest, rsp *pbCommon.Success) error {
	logger.Infof("Received CardDeck.ModifyCard request: %v", req)
	card, err := helper.FindStoreEntity(ctx, e.store.FindCardByID, req.Card.CardID, helper.CardDeckServiceID)
	if err != nil {
		return err
	}
	if err := e.checkUserRoleAccess(ctx, req.UserID, card.Deck.GroupID, pbCommon.GroupRole_WRITE); err != nil {
		return err
	}
	if err = e.store.DeleteCardSidesOfCardByID(ctx, card.ID); err != nil {
		return err
	}
	if err = e.generateCardSidesForCard(ctx, *card, req.Card.Sides); err != nil {
		return err
	}
	rsp.Success = true
	logger.Infof("Successfully modified card (%s) in deck (%s)", card.ID, card.DeckID)
	return nil
}

func (e *CardDeck) DeleteCard(ctx context.Context, req *pbCommon.CardRequest, rsp *pbCommon.Success) error {
	logger.Infof("Received CardDeck.DeleteCard request: %v", req)
	card, err := helper.FindStoreEntity(ctx, e.store.FindCardByID, req.Card.CardID, helper.CardDeckServiceID)
	if err != nil {
		return err
	}
	if err := e.checkUserRoleAccess(ctx, req.UserID, card.Deck.GroupID, pbCommon.GroupRole_WRITE); err != nil {
		return err
	}
	err = e.store.DeleteCard(ctx, card)
	if err != nil {
		return err
	}
	rsp.Success = true
	logger.Infof("Successfully deleted card (%s) in deck (%s)", req.Card.CardID, card.DeckID)
	return nil
}

func (e *CardDeck) CreateCardSide(ctx context.Context, req *pbCommon.CardSideRequest, rsp *pbCommon.CardSide) error {
	logger.Infof("Received CardDeck.CreateCardSide request: %v", req)
	card, err := helper.FindStoreEntity(ctx, e.store.FindCardByID, req.CardID, helper.CardDeckServiceID)
	if err != nil {
		return err
	}
	if err := e.checkUserRoleAccess(ctx, req.UserID, card.Deck.GroupID, pbCommon.GroupRole_WRITE); err != nil {
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
		Header:             req.CardSide.Header,
		Description:        req.CardSide.Description,
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
	rsp.CardSideID = newCardSide.ID
	logger.Infof("Successfully created card side (%s) in card (%s)", newCardSide.ID, card.ID)
	return nil
}

func (e *CardDeck) ModifyCardSide(ctx context.Context, req *pbCommon.CardSideRequest, rsp *pbCommon.Success) error {
	logger.Infof("Received CardDeck.ModifyCardSide request: %v", req)
	cardSide, err := e.getCardSideAndCheckForValidAccess(ctx, req.UserID, req.CardSide.CardSideID)
	if err != nil {
		return err
	}
	cardSide.Header = req.CardSide.Header
	cardSide.Description = req.CardSide.Description
	err = e.store.ModifyCardSide(ctx, cardSide)
	if err != nil {
		return err
	}
	rsp.Success = true
	logger.Infof("Successfully modified card side %s of card %s", req.CardSide.CardSideID, cardSide.CardID)
	return nil
}

func (e *CardDeck) DeleteCardSide(ctx context.Context, req *pbCommon.CardSideRequest, rsp *pbCommon.Success) error {
	logger.Infof("Received CardDeck.DeleteCardSide request: %v", req)
	cardSideToDelete, err := e.getCardSideAndCheckForValidAccess(ctx, req.UserID, req.CardSide.CardSideID)
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
	logger.Infof("Successfully deleted card side %s of card %s", req.CardSide.CardSideID, cardSideToDelete.CardID)
	return nil
}

func (e *CardDeck) GetUserFavoriteDecks(ctx context.Context, req *pbCommon.User, rsp *pbCommon.Decks) error {
	logger.Infof("Received CardDeck.GetUserFavoriteDecks request: %v", req)
	favoriteDecks, err := helper.FindStoreEntity(ctx, e.store.FindFavoriteDecks, req.UserID, helper.CardDeckServiceID)
	if err != nil {
		return err
	}
	rsp.Decks = converter.ConvertToTypeArray(favoriteDecks, converter.StoreDeckToProtoDeckConverter)
	logger.Infof("Successfully retrieved user %s's favorite decks.", req.UserID)
	return nil
}

func (e *CardDeck) AddUserFavoriteDeck(ctx context.Context, req *pbCommon.DeckRequest, rsp *pbCommon.Success) error {
	logger.Infof("Received CardDeck.AddUserFavoriteDeck request: %v", req)
	if err := e.store.AddFavoriteDeck(ctx, req.UserID, req.Deck.DeckID); err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return helper.NewMicroDeckAlreadyFavoriteErr(helper.CardDeckServiceID)
		}
		return err
	}
	rsp.Success = true
	logger.Infof("Successfully added %s to user %s's favorite decks.", req.Deck.DeckID, req.UserID)
	return nil
}

func (e *CardDeck) DeleteUserFavoriteDeck(ctx context.Context, req *pbCommon.DeckRequest, rsp *pbCommon.Success) error {
	logger.Infof("Received CardDeck.DeleteUserFavoriteDeck request: %v", req)
	if err := e.store.DeleteFavoriteDeck(ctx, req.UserID, req.Deck.DeckID); err != nil {
		return err
	}
	rsp.Success = true
	logger.Infof("Successfully deleted %s from user %s's favorite decks.", req.Deck.DeckID, req.UserID)
	return nil
}

func (e *CardDeck) GetUserActiveDecks(ctx context.Context, req *pbCommon.User, rsp *pbCommon.Decks) error {
	logger.Infof("Received CardDeck.GetUserActiveDecks request: %v", req)
	activeDecks, err := helper.FindStoreEntity(ctx, e.store.FindActiveDecks, req.UserID, helper.CardDeckServiceID)
	if err != nil {
		return err
	}
	rsp.Decks = converter.ConvertToTypeArray(activeDecks, converter.StoreDeckToProtoDeckConverter)
	logger.Infof("Successfully retrieved user %s's active decks.", req.UserID)
	return nil
}

func (e *CardDeck) AddUserActiveDeck(ctx context.Context, req *pbCommon.DeckRequest, rsp *pbCommon.Success) error {
	logger.Infof("Received CardDeck.AddUserActiveDeck request: %v", req)
	if err := e.store.AddActiveDeck(ctx, req.UserID, req.Deck.DeckID); err != nil {
		if !errors.Is(err, gorm.ErrDuplicatedKey) {
			return err
		}
	}
	rsp.Success = true
	logger.Infof("Successfully added %s to user %s's active decks.", req.Deck.DeckID, req.UserID)
	return nil
}

func (e *CardDeck) DeleteUserActiveDeck(ctx context.Context, req *pbCommon.DeckRequest, rsp *pbCommon.Success) error {
	logger.Infof("Received CardDeck.DeleteUserActiveDeck request: %v", req)
	if err := e.store.DeleteActiveDeck(ctx, req.UserID, req.Deck.DeckID); err != nil {
		return err
	}
	rsp.Success = true
	logger.Infof("Successfully deleted %s from user %s's active decks.", req.Deck.DeckID, req.UserID)
	return nil
}
