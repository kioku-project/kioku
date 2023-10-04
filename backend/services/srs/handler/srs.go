package handler

import (
	"context"
	"github.com/kioku-project/kioku/pkg/converter"
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
	cardBinding, err := e.store.FindCardBinding(req.UserID, req.CardID)
	if err != nil {
		return err
	}
	if cardBinding.DeckID != req.DeckID {
		return helper.NewMicroWrongDeckIDErr(helper.SrsServiceID)
	}

	// calculate new due date
	switch req.Rating {
	case 0: // Forgotten
		newInterval := 0
		cardBinding.Due = time.Now().Add(time.Hour * 24 * time.Duration(newInterval)).Unix()
		cardBinding.LastInterval = uint32(newInterval)
	case 1: // Hard
		newIvl := math.Max(float64(cardBinding.LastInterval), 1)
		cardBinding.Due = time.Now().Add(time.Hour * 24 * time.Duration(newIvl)).Unix()
		cardBinding.LastInterval = uint32(newIvl)
	case 2: // Easy
		newIvl := math.Max(float64(cardBinding.LastInterval*2), 1)
		cardBinding.Due = time.Now().Add(time.Hour * 24 * time.Duration(newIvl)).Unix()
		cardBinding.LastInterval = uint32(newIvl)
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

func (e *Srs) Pull(ctx context.Context, req *pb.DeckPullRequest, rsp *pb.SrsPullResponse) error {
	logger.Infof("Received Srs.Pull request: %v", req)
	cards, err := e.store.FindDeckCards(
		req.UserID,
		req.DeckID,
	)
	if err != nil {
		return err
	}
	var dueCards []*model.UserCardBinding
	for _, card := range cards {
		if card.Due <= time.Now().Unix() {
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
	returnedCard = converter.CardDeckProtoCardToSrsProtoCardConverter(cardWithContent)
	rsp.Card = returnedCard
	return nil
}

func (e *Srs) AddUserCardBinding(ctx context.Context, req *pb.BindingRequest, rsp *pb.SuccessResponse) error {
	logger.Infof("Received Srs.AddUserCardBinding request: %v", req)
	err := e.store.CreateUserCard(&model.UserCardBinding{
		UserID:       req.UserID,
		CardID:       req.CardID,
		DeckID:       req.DeckID,
		Type:         0,
		Due:          time.Now().Unix(),
		LastInterval: 0,
		Factor:       1,
	})
	if err != nil {
		return err
	}
	rsp.Success = true
	return nil
}
func (e *Srs) GetDeckCardsDue(ctx context.Context, req *pb.DeckPullRequest, rsp *pb.DueResponse) error {
	logger.Infof("Received Srs.GetDeckCardsDue request: %v", req)
	cards, err := e.store.FindDeckCards(
		req.UserID,
		req.DeckID,
	)
	if err != nil {
		return err
	}
	var dueCards []*model.UserCardBinding
	for _, card := range cards {
		if card.Due <= time.Now().Unix() {
			dueCards = append(dueCards, card)
		}
	}
	rsp.Due = int64(len(dueCards))
	return nil
}
func (e *Srs) GetUserCardsDue(ctx context.Context, req *pb.UserDueRequest, rsp *pb.UserDueResponse) error {
	logger.Infof("Received Srs.GetUserCardsDue request: %v", req)
	cards, err := e.store.FindUserCards(req.UserID)
	if err != nil {
		return err
	}
	decksDueMap := map[string]int64{}
	decksDueCount := 0
	var dueCards []*model.UserCardBinding
	for _, card := range cards {
		if card.Due <= time.Now().Unix() {
			dueCards = append(dueCards, card)
			if _, ok := decksDueMap[card.DeckID]; !ok {
				decksDueMap[card.DeckID] = 0
				decksDueCount++
			}
			decksDueMap[card.DeckID]++
		}
	}
	rsp.DueCards = int64(len(dueCards))
	rsp.DueDecks = int64(decksDueCount)
	return nil
}
