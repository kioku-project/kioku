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
	case 0: // Hard
		newInterval := 0.0
		cardBinding.Due = time.Now().Add(time.Hour * 24 * time.Duration(newInterval)).Unix()
		cardBinding.LastInterval = newInterval
	case 1: // Medium
		newIvl := 3.0
		cardBinding.Due = time.Now().Add(time.Hour * 24 * time.Duration(newIvl)).Unix()
		cardBinding.LastInterval = newIvl
	case 2: // Easy
		newIvl := 5.0
		cardBinding.Due = time.Now().Add(time.Hour * 24 * time.Duration(newIvl)).Unix()
		cardBinding.LastInterval = newIvl
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
	dueCards, err := e.store.FindUserDeckDueCards(ctx, req.UserID, req.Deck.DeckID)
	if err != nil {
		return err
	}
	// TODO: Implement carddeck service call here
	targetNewCards := int64(5)

	// get new cards learned today
	currentNewCards, err := e.store.FindUserDeckNewCardsLearnedToday(ctx, req.UserID, req.Deck.DeckID)
	if err != nil {
		return err
	}
	newCards, err := e.store.FindUserDeckNewCards(ctx, req.UserID, req.Deck.DeckID)
	if err != nil {
		return err
	}
	var returnedCard *pbCommon.Card
	if len(dueCards) >= 5 || ((currentNewCards >= targetNewCards || len(newCards) == 0) && len(dueCards) > 0) {
		// determine smartest card to return
		// sort by oldest first
		now := time.Now().Unix()
		sort.Slice(dueCards, func(i, j int) bool {
			return (math.Abs(float64(dueCards[i].Due - now))) > math.Abs(float64((dueCards[j].Due - now)))
		})
		returnedCard = &pbCommon.Card{CardID: dueCards[0].CardID}
	} else if currentNewCards < targetNewCards && len(newCards) > 0 {
		sort.Slice(newCards, func(i, j int) bool {
			return (newCards[i].CardID) > (newCards[j].CardID)
		})
		returnedCard = &pbCommon.Card{CardID: newCards[0].CardID}
	} else {
		*rsp = pbCommon.Card{
			CardID: "",
			Sides:  nil,
		}
		return nil
	}

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
			Factor:       0,
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
