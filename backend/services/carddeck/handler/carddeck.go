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
	roleRsp, err := e.collaborationService.GetGroupUserRole(context.TODO(), &pbCollaboration.GroupRequest{UserID: req.UserID, GroupID: deck.GroupID})
	if err != nil {
		return err
	}
	logger.Infof("Obtained group role (%s) for user (%s)", roleRsp.GroupRole.String(), req.UserID)
	if roleRsp.GroupRole != pbCollaboration.GroupRole_ADMIN && roleRsp.GroupRole != pbCollaboration.GroupRole_WRITE {
		return helper.ErrMicroNotAuthorized(helper.CardDeckServiceID)
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
	logger.Infof("Successfully created new card in deck (%s): %s", req.DeckID, newCard.ID)
	return nil
}

func (e *CardDeck) CreateDeck(ctx context.Context, req *pb.CreateDeckRequest, rsp *pb.IDResponse) error {
	logger.Infof("Received CardDeck.CreateDeck request: %v", req)
	roleRsp, err := e.collaborationService.GetGroupUserRole(context.TODO(), &pbCollaboration.GroupRequest{UserID: req.UserID, GroupID: req.GroupID})
	if err != nil {
		return err
	}
	logger.Infof("Obtained role (%s) for group (%s) for user (%s)", roleRsp.GroupRole.String(), req.GroupID, req.UserID)
	if roleRsp.GroupRole != pbCollaboration.GroupRole_ADMIN && roleRsp.GroupRole != pbCollaboration.GroupRole_WRITE {
		return helper.ErrMicroNotAuthorized(helper.CardDeckServiceID)
	}
	newDeck := model.Deck{
		Name:      req.DeckName,
		CreatedAt: time.Now(),
		GroupID:   roleRsp.GroupID,
	}
	err = e.store.CreateDeck(&newDeck)
	if err != nil {
		return err
	}
	rsp.ID = newDeck.ID
	logger.Infof("Successfully created new deck (%s) in group (%s)", req.GroupID, newDeck.ID)
	return nil
}

func (e *CardDeck) GetDeckCards(ctx context.Context, req *pb.DeckCardsRequest, rsp *pb.DeckCardsResponse) error {
	logger.Infof("Received CardDeck.GetDeckCards request: %v", req)
	deck, err := e.store.FindDeckByID(req.DeckID)
	logger.Infof("Found deck with id %s", req.DeckID)
	if err != nil {
		if errors.Is(err, helper.ErrStoreNoEntryWithID) {
			return helper.ErrMicroNoEntryWithID(helper.CardDeckServiceID)
		}
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
	groupRsp, err := e.collaborationService.FindGroupByID(context.TODO(), &pbCollaboration.GroupRequest{UserID: req.UserID, GroupID: req.GroupID})
	if err != nil {
		return err
	}
	logger.Infof("Found group with id %s", req.GroupID)
	decks, err := e.store.FindDecksByGroupID(groupRsp.GroupID)
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
