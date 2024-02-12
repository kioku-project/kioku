package handler

import (
	"context"
	"math"
	"sort"
	"time"

	"github.com/kioku-project/kioku/pkg/helper"
	"github.com/kioku-project/kioku/pkg/model"
	pbCommon "github.com/kioku-project/kioku/pkg/proto"
	pbCardDeck "github.com/kioku-project/kioku/services/carddeck/proto"
	"go-micro.dev/v4/logger"

	"github.com/kioku-project/kioku/store"
)

type StaticSrs struct {
	store           store.SrsStore
	cardDeckService pbCardDeck.CardDeckService
}

func New(s store.SrsStore, cds pbCardDeck.CardDeckService) *StaticSrs {
	return &StaticSrs{store: s, cardDeckService: cds}
}

func (e *StaticSrs) Push(ctx context.Context, req *pbCommon.SrsPushRequest, rsp *pbCommon.Success) error {
	logger.Infof("Received StaticSrs.Push request: %v", req)
	cardBinding, err := e.store.FindCardBinding(ctx, req.UserID, req.CardID)
	if err != nil {
		return err
	}
	if cardBinding.DeckID != req.DeckID {
		return helper.NewMicroWrongDeckIDErr(helper.SrsServiceID)
	}

	// calculate new due date
	switch req.Rating {
	case 0: // Forgotten
		newInterval := math.Max(float64(cardBinding.LastInterval-1), 0)
		cardBinding.Due = time.Now().Add(time.Hour * 24 * time.Duration(newInterval)).Unix()
		cardBinding.LastInterval = uint32(newInterval)
	case 1: // Hard
		newIvl := cardBinding.LastInterval
		cardBinding.Due = time.Now().Add(time.Hour * 24 * time.Duration(newIvl)).Unix()
		cardBinding.LastInterval = newIvl
	case 2: // Easy
		newIvl := cardBinding.LastInterval + 1
		cardBinding.Due = time.Now().Add(time.Hour * 24 * time.Duration(newIvl)).Unix()
		cardBinding.LastInterval = uint32(newIvl)
	default:
		return helper.NewMicroWrongRatingErr(helper.SrsServiceID)
	}
	if err = e.store.ModifyUserCard(ctx, cardBinding); err != nil {
		return err
	}

	// Add revlog entry
	if err = e.store.CreateRevlog(ctx,
		&model.Revlog{
			CardID: req.CardID,
			UserID: req.UserID,
			Date:   time.Now().Unix(),
			Rating: req.Rating,
		}); err != nil {
		return err
	}
	rsp.Success = true
	return nil
}

func (e *StaticSrs) Pull(ctx context.Context, req *pbCommon.DeckRequest, rsp *pbCommon.Card) error {
	logger.Infof("Received StaticSrs.Pull request: %v", req)
	if _, err := e.cardDeckService.AddUserActiveDeck(ctx, &pbCommon.DeckRequest{UserID: req.UserID, Deck: &pbCommon.Deck{DeckID: req.Deck.DeckID}}); err != nil {
		return err
	}
	cards, err := e.store.FindUserDeckCards(
		ctx, req.UserID,
		req.Deck.DeckID,
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
		*rsp = pbCommon.Card{
			CardID: "",
			Sides:  nil,
		}
		return nil
	}

	returnedCard := &pbCommon.Card{CardID: dueCards[0].CardID}

	// get content of card
	cardWithContent, err := e.cardDeckService.GetCard(ctx, &pbCommon.CardRequest{
		UserID: req.UserID,
		Card: &pbCommon.Card{
			CardID: returnedCard.CardID,
		},
	})
	if err != nil {
		return err
	}
	*rsp = pbCommon.Card{
		CardID: cardWithContent.CardID,
		DeckID: cardWithContent.DeckID,
		Sides:  cardWithContent.Sides,
	}
	return nil
}

func (e *StaticSrs) AddUserCardBinding(ctx context.Context, req *pbCommon.BindingRequest, rsp *pbCommon.Success) error {
	logger.Infof("Received StaticSrs.AddUserCardBinding request: %v", req)
	err := e.store.CreateUserCard(ctx,
		&model.UserCardBinding{
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

func (e *StaticSrs) GetDeckCardsDue(ctx context.Context, req *pbCommon.DeckRequest, rsp *pbCommon.UserDueResponse) error {
	logger.Infof("Received StaticSrs.GetDeckCardsDue request: %v", req)
	cards, err := e.store.FindUserDeckCards(
		ctx,
		req.UserID,
		req.Deck.DeckID,
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
	rsp.DueCards = int64(len(dueCards))
	return nil
}
func (e *StaticSrs) GetUserCardsDue(ctx context.Context, req *pbCommon.User, rsp *pbCommon.UserDueResponse) error {
	logger.Infof("Received StaticSrs.GetUserCardsDue request: %v", req)
	cards, err := e.store.FindUserCards(ctx, req.UserID)
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
