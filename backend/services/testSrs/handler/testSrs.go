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

type TestSrs struct {
	store           store.SrsStore
	cardDeckService pbCardDeck.CardDeckService
}

func New(s store.SrsStore, cds pbCardDeck.CardDeckService) *TestSrs {
	return &TestSrs{store: s, cardDeckService: cds}
}

func (e *TestSrs) Push(ctx context.Context, req *pbCommon.SrsPushRequest, rsp *pbCommon.Success) error {
	logger.Infof("Received TestSrs.Push request: %v", req)
	cardBinding, err := e.store.FindCardBinding(ctx, req.UserID, req.CardID)
	if err != nil {
		return err
	}
	if cardBinding.DeckID != req.DeckID {
		return helper.NewMicroWrongDeckIDErr(helper.SrsServiceID)
	}
	now := time.Now()

	// Add revlog entry
	if err = e.store.CreateRevlog(ctx,
		&model.Revlog{
			CardID: req.CardID,
			UserID: req.UserID,
			Date:   now.Unix(),
			Rating: req.Rating,
			Due:    cardBinding.Due,
		}); err != nil {
		return err
	}

	// calculate new due date
	cardBinding.Due = 1
	if err = e.store.ModifyUserCard(ctx, cardBinding); err != nil {
		return err
	}

	rsp.Success = true
	return nil
}

func (e *TestSrs) Pull(ctx context.Context, req *pbCommon.DeckRequest, rsp *pbCommon.Card) error {
	logger.Infof("Received TestSrs.Pull request: %v", req)
	newCards, err := e.store.FindUserDeckNewCards(ctx, req.UserID, req.Deck.DeckID)
	if err != nil {
		return err
	}
	var returnedCard *pbCommon.Card
	if len(newCards) > 0 {
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

func (e *TestSrs) AddUserCardBinding(ctx context.Context, req *pbCommon.BindingRequest, rsp *pbCommon.Success) error {
	logger.Infof("Received TestSrs.AddUserCardBinding request: %v", req)
	err := e.store.CreateUserCard(ctx,
		&model.UserCardBinding{
			UserID:       req.UserID,
			CardID:       req.CardID,
			DeckID:       req.DeckID,
			Type:         0,
			Due:          0,
			LastInterval: 0,
			Factor:       0,
		})
	if err != nil {
		return err
	}
	rsp.Success = true
	return nil
}

func (e *TestSrs) GetDeckCardsDue(ctx context.Context, req *pbCommon.DeckRequest, rsp *pbCommon.UserDueResponse) error {
	logger.Infof("Received TestSrs.GetDeckCardsDue request: %v", req)
	newCards, err := e.store.FindUserDeckNewCards(
		ctx,
		req.UserID,
		req.Deck.DeckID,
	)
	if err != nil {
		return err
	}
	rsp.DueCards = int64(0)
	rsp.NewCards = int64(len(newCards))
	return nil
}
func (e *TestSrs) GetUserCardsDue(ctx context.Context, req *pbCommon.User, rsp *pbCommon.UserDueResponse) error {
	logger.Infof("Received TestSrs.GetUserCardsDue request: %v", req)
	activeDecks, err := e.cardDeckService.GetUserActiveDecks(ctx, &pbCommon.User{UserID: req.UserID})
	if err != nil {
		return err
	}
	for _, deck := range activeDecks.Decks {
		dueCards, err := e.store.FindUserDeckDueCards(ctx, req.UserID, deck.DeckID)
		if err != nil {
			return err
		}
		// TODO: Implement carddeck service call here
		userNewCardsPerDay := int64(5)

		// get new cards learned today
		newCardsAmount, err := e.store.FindUserDeckNewCardsLearnedToday(ctx, req.UserID, deck.DeckID)
		unrestrictedNewCards, err := e.store.FindUserDeckNewCards(
			ctx,
			req.UserID,
			deck.DeckID,
		)
		if err != nil {
			return err
		}
		newCards := int64(math.Min(math.Max(float64(userNewCardsPerDay)-float64(newCardsAmount), 0), float64(len(unrestrictedNewCards))))
		if len(dueCards) > 0 || newCards > 0 {
			rsp.DueDecks++
		}
		rsp.DueCards += int64(len(dueCards))
		rsp.NewCards += newCards
	}
	return nil
}
