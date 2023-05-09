package handler

import (
	"context"
	"time"

	"go-micro.dev/v4/logger"

	"github.com/kioku-project/kioku/pkg/model"
	pb "github.com/kioku-project/kioku/services/carddeck/proto"
	pbcollab "github.com/kioku-project/kioku/services/collaboration/proto"
	"github.com/kioku-project/kioku/store"
)

type Carddeck struct {
	store                store.CardDeckStore
	collaborationService pbcollab.CollaborationService
}

func New(s store.CardDeckStore, cS pbcollab.CollaborationService) *Carddeck {
	return &Carddeck{store: s, collaborationService: cS}
}

func (e *Carddeck) CreateCard(ctx context.Context, req *pb.CardRequest, rsp *pb.PublicIDResponse) error {
	logger.Infof("Received Carddeck.CreateCard request: %v", req)
	deck, err := e.store.FindDeckByPublicID(req.DeckPublicID)
	if err != nil {
		return err
	}
	newCard := model.Card{
		DeckID:    deck.GroupID,
		Frontside: req.Frontside,
		Backside:  req.Backside,
	}
	err = e.store.CreateCard(&newCard)
	if err != nil {
		return err
	}
	rsp.PublicID = newCard.PublicID
	return nil
}

func (e *Carddeck) CreateDeck(ctx context.Context, req *pb.DeckRequest, rsp *pb.PublicIDResponse) error {
	logger.Infof("Received Carddeck.CreateDeck request: %v", req)
	roleRsp, err := e.collaborationService.GetGroupUserRole(context.TODO(), &pbcollab.GroupRequest{UserID: req.UserID, GroupPublicID: req.GroupPublicID})
	if err != nil {
		return err
	}
	newDeck := model.Deck{
		Name:      req.DeckName,
		CreatedAt: time.Now(),
		GroupID:   uint(roleRsp.GroupID),
	}
	err = e.store.CreateDeck(&newDeck)
	if err != nil {
		return err
	}
	rsp.PublicID = newDeck.PublicID
	return nil
}
