package handler

import (
	"context"
	"errors"
	"time"

	"go-micro.dev/v4/logger"

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

func (e *CardDeck) securityRoleHandler(ctx context.Context, userID string, groupID string, requiredRole pbCollaboration.GroupRole) error {
	logger.Infof("Requesting group role for user (%s)", userID)
	roleRsp, err := e.collaborationService.GetGroupUserRole(ctx, &pbCollaboration.GroupRequest{UserID: userID, GroupID: groupID})
	if err != nil {
		return err
	}
	logger.Infof("Obtained group role (%s) for user (%s)", roleRsp.GroupRole.String(), userID)
	if !helper.IsAuthorized(roleRsp.GroupRole, requiredRole) {
		return helper.ErrMicroNotAuthorized(helper.CardDeckServiceID)
	}
	logger.Infof("Authenticated group role (%s) for user (%s)", roleRsp.GroupRole.String(), userID)
	return nil
}

func (e *CardDeck) CreateCard(ctx context.Context, req *pb.CreateCardRequest, rsp *pb.IDResponse) error {
	logger.Infof("Received CardDeck.CreateCard request: %v", req)
	deck, err := e.store.FindDeckByID(req.DeckID)
	if err != nil {
		if errors.Is(err, helper.ErrStoreNoEntryWithID) {
			return helper.ErrMicroNoEntryWithID(helper.CardDeckServiceID)
		}
		return err
	}
	logger.Infof("Found deck with id %s", req.DeckID)
	if err := e.securityRoleHandler(ctx, req.UserID, deck.GroupID, pbCollaboration.GroupRole_WRITE); err != nil {
		return err
	}
	newCard := model.Card{
		DeckID:    deck.ID,
		Frontside: req.Frontside,
		Backside:  req.Backside,
	}
	err = e.store.CreateCard(&newCard)
	if err != nil {
		return err
	}
	rsp.ID = newCard.ID
	logger.Infof("Successfully created new card (%s) in deck (%s)", newCard.ID, req.DeckID)
	return nil
}

func (e *CardDeck) ModifyCard(ctx context.Context, req *pb.ModifyCardRequest, rsp *pb.SuccessResponse) error {
	logger.Infof("Received CardDeck.ModifyCard request: %v", req)
	card, err := e.store.FindCardByID(req.CardID)
	if err != nil {
		if errors.Is(err, helper.ErrStoreNoEntryWithID) {
			return helper.ErrMicroNoEntryWithID(helper.CardDeckServiceID)
		}
		return err
	}
	logger.Infof("Found card with id %s", req.CardID)
	if err := e.securityRoleHandler(ctx, req.UserID, card.Deck.GroupID, pbCollaboration.GroupRole_WRITE); err != nil {
		return err
	}
	if req.Frontside != nil {
		card.Frontside = *req.Frontside
	}
	if req.Backside != nil {
		card.Backside = *req.Backside
	}
	err = e.store.ModifyCard(card)
	if err != nil {
		return err
	}
	rsp.Success = true
	logger.Infof("Successfully modified card (%s) in deck (%s)", req.CardID, card.DeckID)
	return nil
}

func (e *CardDeck) DeleteCard(ctx context.Context, req *pb.DeleteCardRequest, rsp *pb.SuccessResponse) error {
	logger.Infof("Received CardDeck.DeleteCard request: %v", req)
	card, err := e.store.FindCardByID(req.CardID)
	if err != nil {
		if errors.Is(err, helper.ErrStoreNoEntryWithID) {
			return helper.ErrMicroNoEntryWithID(helper.CardDeckServiceID)
		}
		return err
	}
	logger.Infof("Found card with id %s", req.CardID)
	if err := e.securityRoleHandler(ctx, req.UserID, card.Deck.GroupID, pbCollaboration.GroupRole_WRITE); err != nil {
		return err
	}
	err = e.store.DeleteCard(card)
	if err != nil {
		return err
	}
	rsp.Success = true
	logger.Infof("Successfully deleted card (%s) in deck (%s)", req.CardID, card.DeckID)
	return nil
}

func (e *CardDeck) CreateDeck(ctx context.Context, req *pb.CreateDeckRequest, rsp *pb.IDResponse) error {
	logger.Infof("Received CardDeck.CreateDeck request: %v", req)
	if err := e.securityRoleHandler(ctx, req.UserID, req.GroupID, pbCollaboration.GroupRole_WRITE); err != nil {
		return err
	}
	newDeck := model.Deck{
		Name:      req.DeckName,
		CreatedAt: time.Now(),
		GroupID:   req.GroupID,
	}
	err := e.store.CreateDeck(&newDeck)
	if err != nil {
		return err
	}
	rsp.ID = newDeck.ID
	logger.Infof("Successfully created new deck (%s) in group (%s)", req.GroupID, newDeck.ID)
	return nil
}

func (e *CardDeck) ModifyDeck(ctx context.Context, req *pb.ModifyDeckRequest, rsp *pb.SuccessResponse) error {
	logger.Infof("Received CardDeck.ModifyCard request: %v", req)
	deck, err := e.store.FindDeckByID(req.DeckID)
	if err != nil {
		if errors.Is(err, helper.ErrStoreNoEntryWithID) {
			return helper.ErrMicroNoEntryWithID(helper.CardDeckServiceID)
		}
		return err
	}
	logger.Infof("Found deck with id %s", req.DeckID)
	if err := e.securityRoleHandler(ctx, req.UserID, deck.GroupID, pbCollaboration.GroupRole_WRITE); err != nil {
		return err
	}
	if req.DeckName != nil {
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

func (e *CardDeck) DeleteDeck(ctx context.Context, req *pb.DeckRequest, rsp *pb.SuccessResponse) error {
	logger.Infof("Received CardDeck.DeleteDeck request: %v", req)
	deck, err := e.store.FindDeckByID(req.DeckID)
	if err != nil {
		if errors.Is(err, helper.ErrStoreNoEntryWithID) {
			return helper.ErrMicroNoEntryWithID(helper.CardDeckServiceID)
		}
		return err
	}
	logger.Infof("Found deck with id %s", req.DeckID)
	if err := e.securityRoleHandler(ctx, req.UserID, deck.GroupID, pbCollaboration.GroupRole_WRITE); err != nil {
		return err
	}
	err = e.store.DeleteDeck(deck)
	if err != nil {
		return err
	}
	rsp.Success = true
	logger.Infof("Successfully deleted deck (%s) in group (%s)", req.DeckID, deck.GroupID)
	return nil
}

func (e *CardDeck) GetDeckCards(ctx context.Context, req *pb.DeckRequest, rsp *pb.DeckCardsResponse) error {
	logger.Infof("Received CardDeck.GetDeckCards request: %v", req)
	deck, err := e.store.FindDeckByID(req.DeckID)
	logger.Infof("Found deck with id %s", req.DeckID)
	if err != nil {
		if errors.Is(err, helper.ErrStoreNoEntryWithID) {
			return helper.ErrMicroNoEntryWithID(helper.CardDeckServiceID)
		}
		return err
	}
	if err := e.securityRoleHandler(ctx, req.UserID, deck.GroupID, pbCollaboration.GroupRole_READ); err != nil {
		return err
	}
	rsp.Cards = make([]*pb.Card, len(deck.Cards))
	for i, card := range deck.Cards {
		rsp.Cards[i] = &pb.Card{
			CardID:    card.ID,
			Frontside: card.Frontside,
			Backside:  card.Backside,
		}
	}
	logger.Infof("Found %d cards in deck with id %s", len(deck.Cards), req.DeckID)
	return nil
}

func (e *CardDeck) GetGroupDecks(ctx context.Context, req *pb.GroupDecksRequest, rsp *pb.GroupDecksResponse) error {
	logger.Infof("Received CardDeck.GetGroupDecks request: %v", req)
	if err := e.securityRoleHandler(ctx, req.UserID, req.GroupID, pbCollaboration.GroupRole_READ); err != nil {
		return err
	}
	decks, err := e.store.FindDecksByGroupID(req.GroupID)
	if err != nil {
		if errors.Is(err, helper.ErrStoreNoEntryWithID) {
			return helper.ErrMicroNoEntryWithID(helper.CardDeckServiceID)
		}
		return err
	}
	rsp.Decks = make([]*pb.Deck, len(decks))
	for i, deck := range decks {
		rsp.Decks[i] = &pb.Deck{
			DeckID:   deck.ID,
			DeckName: deck.Name,
		}
	}
	logger.Infof("Found %d decks in group with id %s", len(decks), req.GroupID)
	return nil
}
