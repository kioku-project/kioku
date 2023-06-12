package handler

import (
	"context"
	"github.com/kioku-project/kioku/pkg/helper"
	"github.com/kioku-project/kioku/pkg/model"
	pbCardDeck "github.com/kioku-project/kioku/services/carddeck/proto"
	"go-micro.dev/v4/logger"
	"math"
	"sort"
	"time"

	pb "github.com/kioku-project/kioku/services/srs/proto"
	"github.com/kioku-project/kioku/store"
)

type Srs struct {
	store           store.SrsStore
	cardDeckService pbCardDeck.CardDeckService
}

func New(s store.SrsStore, cds pbCardDeck.CardDeckService) *Srs {
	return &Srs{store: s, cardDeckService: cds}
}

func (e *Srs) Push(ctx context.Context, req *pb.SrsPushRequest, rsp *pb.SuccessResponse) error {
	logger.Infof("Received Srs.Push request: %v", req)
	cardBinding, err := e.store.GetCardBinding(req.UserID, req.CardID)
	if err != nil {
		return err
	}

	// calculate new due date
	switch req.Rating {
	case 0: // Forgotten
		newIvl := 0
		cardBinding.Due = time.Now().Add(time.Hour * 24 * time.Duration(newIvl)).Unix()
		cardBinding.LastIvl = uint32(newIvl)
	case 1: // Hard
		newIvl := math.Max(float64(cardBinding.LastIvl), 1)
		cardBinding.Due = time.Now().Add(time.Hour * 24 * time.Duration(newIvl)).Unix()
		cardBinding.LastIvl = uint32(newIvl)
	case 2: // Easy
		newIvl := math.Max(float64(cardBinding.LastIvl*2), 1)
		cardBinding.Due = time.Now().Add(time.Hour * 24 * time.Duration(newIvl)).Unix()
		cardBinding.LastIvl = uint32(newIvl)
	default:
		return helper.NewMicroWrongRatingErr(helper.SrsServiceID)
	}
	e.store.ModifyUserCard(cardBinding)

	// Add revlog entry
	e.store.CreateRevlog(&model.Revlog{
		CardID: req.CardID,
		UserID: req.UserID,
		Date:   time.Now().Unix(),
		Rating: req.Rating,
	})
	rsp.Success = true
	return nil
}

func (e *Srs) Pull(ctx context.Context, req *pb.SrsPullRequest, rsp *pb.SrsPullResponse) error {
	logger.Infof("Received Srs.Pull request: %v", req)
	cards, err := e.store.GetDeckCards(
		req.UserID,
		req.DeckID,
	)
	if err != nil {
		return err
	}
	var dueCards []*model.UserCardBinding
	for _, card := range cards {
		if card.Due < time.Now().Unix() {
			dueCards = append(dueCards, card)
		}
	}
	// determine smartest card to return
	// sort by oldest first
	sort.Slice(dueCards, func(i, j int) bool {
		return dueCards[i].Due < dueCards[j].Due
	})
	// if no more cards are due, return empty card
	if len(dueCards) == 0 {
		rsp.Card = &pb.Card{
			CardID: "",
			Sides:  nil,
		}
		return nil
	}

	returnedCard := &pb.Card{CardID: dueCards[0].CardID}

	// get content of card
	cardWithContent, err := e.cardDeckService.GetCard(ctx, &pbCardDeck.IDRequest{
		UserID:   req.UserID,
		EntityID: returnedCard.CardID,
	})
	if err != nil {
		return err
	}
	returnedCard.Sides = make([]*pb.Side, len(cardWithContent.Sides))
	for i, card := range cardWithContent.Sides {
		returnedCard.Sides[i] = &pb.Side{
			Header:      card.Header,
			Description: card.Description,
		}
	}
	rsp.Card = returnedCard
	return nil
}

func (e *Srs) AddUserCardBinding(ctx context.Context, req *pb.BindingRequest, rsp *pb.SuccessResponse) error {
	logger.Infof("Received Srs.AddUserCardBinding request: %v", req)
	err := e.store.CreateUserCard(&model.UserCardBinding{
		UserID:  req.UserID,
		CardID:  req.CardID,
		DeckID:  req.DeckID,
		Type:    0,
		Due:     time.Now().Unix(),
		LastIvl: 0,
		Factor:  1,
	})
	if err != nil {
		return err
	}
	rsp.Success = true
	return nil
}
