package handler

import (
	"context"
	"github.com/kioku-project/kioku/pkg/converter"
	"go-micro.dev/v4/logger"
	"time"

	"github.com/kioku-project/kioku/pkg/helper"
	"github.com/kioku-project/kioku/pkg/model"
	pb "github.com/kioku-project/kioku/services/carddeck/proto"
	pbCollaboration "github.com/kioku-project/kioku/services/collaboration/proto"
	"github.com/kioku-project/kioku/store"
)

type CardDeck struct {
	store                store.CardDeckStore
	collaborationService pbCollaboration.CollaborationService
}

func New(s store.CardDeckStore, cS pbCollaboration.CollaborationService) *CardDeck {
	return &CardDeck{store: s, collaborationService: cS}
}

func (e *CardDeck) checkUserRoleAccess(ctx context.Context, userID string, groupID string, requiredRole pbCollaboration.GroupRole) error {
	logger.Infof("Requesting group role for user (%s)", userID)
	roleRsp, err := e.collaborationService.GetGroupUserRole(ctx, &pbCollaboration.GroupRequest{UserID: userID, GroupID: groupID})
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

func (e *CardDeck) getCardSideAndCheckForValidAccess(ctx context.Context, userID string, cardSideID string) (*model.CardSide, error) {
	cardSide, err := helper.FindStoreEntity(e.store.FindCardSideByID, cardSideID, helper.CardDeckServiceID)
	if err != nil {
		return nil, err
	}
	deck, err := helper.FindStoreEntity(e.store.FindDeckByID, cardSide.Card.DeckID, helper.CardDeckServiceID)
	if err != nil {
		return nil, err
	}
	if err := e.checkUserRoleAccess(ctx, userID, deck.GroupID, pbCollaboration.GroupRole_WRITE); err != nil {
		return nil, err
	}
	return cardSide, nil
}

func (e *CardDeck) GetGroupDecks(ctx context.Context, req *pb.GroupDecksRequest, rsp *pb.GroupDecksResponse) error {
	logger.Infof("Received CardDeck.GetGroupDecks request: %v", req)
	if err := e.checkUserRoleAccess(ctx, req.UserID, req.GroupID, pbCollaboration.GroupRole_READ); err != nil {
		return err
	}
	decks, err := helper.FindStoreEntity(e.store.FindDecksByGroupID, req.GroupID, helper.CardDeckServiceID)
	if err != nil {
		return err
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
	newDeck := model.Deck{
		Name:      req.DeckName,
		CreatedAt: time.Now(),
		GroupID:   req.GroupID,
	}
	if err := e.store.CreateDeck(&newDeck); err != nil {
		return err
	}
	rsp.ID = newDeck.ID
	logger.Infof("Successfully created new deck (%s) in group (%s)", req.GroupID, newDeck.ID)
	return nil
}

func (e *CardDeck) ModifyDeck(ctx context.Context, req *pb.ModifyDeckRequest, rsp *pb.SuccessResponse) error {
	logger.Infof("Received CardDeck.ModifyCard request: %v", req)
	deck, err := helper.FindStoreEntity(e.store.FindDeckByID, req.DeckID, helper.CardDeckServiceID)
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
	err = e.store.ModifyDeck(deck)
	if err != nil {
		return err
	}
	rsp.Success = true
	logger.Infof("Successfully modified deck (%s) in group (%s)", req.DeckID, deck.GroupID)
	return nil
}

func (e *CardDeck) DeleteDeck(ctx context.Context, req *pb.DeleteWithIDRequest, rsp *pb.SuccessResponse) error {
	logger.Infof("Received CardDeck.DeleteDeck request: %v", req)
	deck, err := helper.FindStoreEntity(e.store.FindDeckByID, req.EntityID, helper.CardDeckServiceID)
	if err != nil {
		return err
	}
	if err := e.checkUserRoleAccess(ctx, req.UserID, deck.GroupID, pbCollaboration.GroupRole_ADMIN); err != nil {
		return err
	}
	err = e.store.DeleteDeck(deck)
	if err != nil {
		return err
	}
	rsp.Success = true
	logger.Infof("Successfully deleted deck (%s) in group (%s)", req.EntityID, deck.GroupID)
	return nil
}

func (e *CardDeck) GetDeckCards(ctx context.Context, req *pb.DeckRequest, rsp *pb.DeckCardsResponse) error {
	logger.Infof("Received CardDeck.GetDeckCards request: %v", req)
	deck, err := helper.FindStoreEntity(e.store.FindDeckByID, req.DeckID, helper.CardDeckServiceID)
	if err != nil {
		return err
	}
	if err := e.checkUserRoleAccess(ctx, req.UserID, deck.GroupID, pbCollaboration.GroupRole_READ); err != nil {
		return err
	}
	rsp.Cards = make([]*pb.Card, len(deck.Cards))
	for i, card := range deck.Cards {
		cardSides, err := e.store.FindCardSidesByCardID(card.ID)
		if err != nil {
			return err
		}
		rsp.Cards[i] = &pb.Card{
			CardID: card.ID,
			Sides:  converter.ConvertToTypeArray(cardSides, converter.StoreCardSideToProtoCardSideConverter),
		}
	}
	logger.Infof("Found %d cards in deck with id %s", len(deck.Cards), req.DeckID)
	return nil
}

func (e *CardDeck) CreateCard(ctx context.Context, req *pb.CreateCardRequest, rsp *pb.IDResponse) error {
	logger.Infof("Received CardDeck.CreateCard request: %v", req)
	deck, err := helper.FindStoreEntity(e.store.FindDeckByID, req.DeckID, helper.CardDeckServiceID)
	if err != nil {
		return err
	}
	if err := e.checkUserRoleAccess(ctx, req.UserID, deck.GroupID, pbCollaboration.GroupRole_WRITE); err != nil {
		return err
	}
	newCard := model.Card{
		DeckID: deck.ID,
	}
	if err = e.store.CreateCard(&newCard); err != nil {
		return err
	}
	var previousCardSide *model.CardSide
	var firstCardSideID string
	for _, side := range req.Sides {
		newCardSide := &model.CardSide{
			CardID:  newCard.ID,
			Content: side,
		}
		if err := e.store.CreateCardSide(newCardSide); err != nil {
			return err
		}
		if previousCardSide != nil {
			previousCardSide.NextCardSideID = newCardSide.ID
			if err := e.store.ModifyCardSide(previousCardSide); err != nil {
				return err
			}
			newCardSide.PreviousCardSideID = previousCardSide.ID
			if err := e.store.ModifyCardSide(newCardSide); err != nil {
				return err
			}
		} else {
			firstCardSideID = newCardSide.ID
		}
		previousCardSide = newCardSide
	}
	newCard.FirstCardSideID = firstCardSideID
	if err := e.store.ModifyCard(&newCard); err != nil {
		return err
	}
	rsp.ID = newCard.ID
	logger.Infof("Successfully created new card (%s) in deck (%s)", newCard.ID, req.DeckID)
	return nil
}

func (e *CardDeck) DeleteCard(ctx context.Context, req *pb.DeleteWithIDRequest, rsp *pb.SuccessResponse) error {
	logger.Infof("Received CardDeck.DeleteCard request: %v", req)
	card, err := helper.FindStoreEntity(e.store.FindCardByID, req.EntityID, helper.CardDeckServiceID)
	if err != nil {
		return err
	}
	if err := e.checkUserRoleAccess(ctx, req.UserID, card.Deck.GroupID, pbCollaboration.GroupRole_WRITE); err != nil {
		return err
	}
	err = e.store.DeleteCard(card)
	if err != nil {
		return err
	}
	rsp.Success = true
	logger.Infof("Successfully deleted card (%s) in deck (%s)", req.EntityID, card.DeckID)
	return nil
}

func (e *CardDeck) ModifyCardSide(ctx context.Context, req *pb.ModifyCardSideRequest, rsp *pb.SuccessResponse) error {
	logger.Infof("Received CardDeck.ModifyCardSide request: %v", req)
	cardSide, err := e.getCardSideAndCheckForValidAccess(ctx, req.UserID, req.CardSideID)
	if err != nil {
		return err
	}
	if req.Content != nil {
		cardSide.Content = *req.Content
	}
	err = e.store.ModifyCardSide(cardSide)
	if err != nil {
		return err
	}
	rsp.Success = true
	logger.Infof("Successfully modified card side %s of card %s", req.CardSideID, cardSide.CardID)
	return nil
}

func (e *CardDeck) DeleteCardSide(ctx context.Context, req *pb.DeleteWithIDRequest, rsp *pb.SuccessResponse) error {
	logger.Infof("Received CardDeck.DeleteCardSide request: %v", req)
	cardSideToDelete, err := e.getCardSideAndCheckForValidAccess(ctx, req.UserID, req.EntityID)
	if err != nil {
		return err
	}
	cardSidesToUpdate := [2]string{cardSideToDelete.PreviousCardSideID, cardSideToDelete.NextCardSideID}
	for index, cardSideToUpdate := range cardSidesToUpdate {
		if cardSideToUpdate != "" {
			side, err := helper.FindStoreEntity(e.store.FindCardSideByID, cardSideToUpdate, helper.CollaborationServiceID)
			if err != nil {
				return err
			}
			if index%2 == 0 {
				side.NextCardSideID = cardSideToDelete.NextCardSideID
			} else {
				side.PreviousCardSideID = cardSideToDelete.PreviousCardSideID
			}
			err = e.store.ModifyCardSide(side)
			if err != nil {
				return err
			}
		} else if index%2 == 0 {
			card, err := helper.FindStoreEntity(e.store.FindCardByID, cardSideToDelete.CardID, helper.CollaborationServiceID)
			if err != nil {
				return err
			}
			card.FirstCardSideID = cardSideToDelete.NextCardSideID
		}
	}
	rsp.Success = true
	logger.Infof("Successfully deleted card side %s of card %s", req.EntityID, cardSideToDelete.CardID)
	return nil
}
