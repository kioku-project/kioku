package handler

import (
	"context"
	"time"

	"go-micro.dev/v4/logger"

	"github.com/kioku-project/kioku/pkg/helper"
	"github.com/kioku-project/kioku/pkg/model"
	pb "github.com/kioku-project/kioku/services/carddeck/proto"
	pbcollab "github.com/kioku-project/kioku/services/collaboration/proto"
	"github.com/kioku-project/kioku/store"
)

type CardDeck struct {
	store                store.CardDeckStore
	collaborationService pbcollab.CollaborationService
}

func New(s store.CardDeckStore, cS pbcollab.CollaborationService) *CardDeck {
	return &CardDeck{store: s, collaborationService: cS}
}

func (e *CardDeck) CreateCard(ctx context.Context, req *pb.CreateCardRequest, rsp *pb.IDResponse) error {
	logger.Infof("Received Carddeck.CreateCard request: %v", req)
	deck, err := e.store.FindDeckByID(req.DeckID)
	if err != nil {
		return err
	}
	roleRsp, err := e.collaborationService.GetGroupUserRole(context.TODO(), &pbcollab.GroupRequest{UserID: req.UserID, GroupID: deck.GroupID})
	if err != nil {
		return err
	}
	if roleRsp.GroupRole != pbcollab.GroupRole_ADMIN && roleRsp.GroupRole != pbcollab.GroupRole_WRITE {
		return helper.ErrNotAuthorized
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
	return nil
}

func (e *CardDeck) CreateDeck(ctx context.Context, req *pb.CreateDeckRequest, rsp *pb.IDResponse) error {
	logger.Infof("Received Carddeck.CreateDeck request: %v", req)
	roleRsp, err := e.collaborationService.GetGroupUserRole(context.TODO(), &pbcollab.GroupRequest{UserID: req.UserID, GroupID: req.GroupID})
	if err != nil {
		return err
	}
	if roleRsp.GroupRole != pbcollab.GroupRole_ADMIN && roleRsp.GroupRole != pbcollab.GroupRole_WRITE {
		return helper.ErrNotAuthorized
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
	return nil
}

func (e *CardDeck) GetDeckCards(ctx context.Context, req *pb.DeckCardsRequest, rsp *pb.DeckCardsResponse) error {
	logger.Infof("Received Carddeck.GetDeckCards request: %v", req)
	deck, err := e.store.FindDeckByID(req.DeckID)
	if err != nil {
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
	return nil
}

func (e *CardDeck) GetGroupDecks(ctx context.Context, req *pb.GroupDecksRequest, rsp *pb.GroupDecksResponse) error {
	logger.Infof("Received Carddeck.GetGroupDecks request: %v", req)
	groupRsp, err := e.collaborationService.FindGroupByID(context.TODO(), &pbcollab.GroupRequest{UserID: req.UserID, GroupID: req.GroupID})
	if err != nil {
		return err
	}
	decks, err := e.store.FindDecksByGroupID(groupRsp.GroupID)
	if err != nil {
		return err
	}
	rsp.Decks = make([]*pb.Deck, len(decks))
	for i, deck := range decks {
		rsp.Decks[i] = &pb.Deck{
			DeckID:   deck.ID,
			DeckName: deck.Name,
		}
	}
	return nil
}
