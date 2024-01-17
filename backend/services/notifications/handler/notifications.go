package handler

import (
	"context"

	"github.com/kioku-project/kioku/pkg/model"
	pbCommon "github.com/kioku-project/kioku/pkg/proto"
	pbCardDeck "github.com/kioku-project/kioku/services/carddeck/proto"
	"go-micro.dev/v4/logger"

	pb "github.com/kioku-project/kioku/services/notifications/proto"
	"github.com/kioku-project/kioku/store"
)

type Notifications struct {
	store           store.NotificationsStore
	cardDeckService pbCardDeck.CardDeckService
}

func New(s store.NotificationsStore, cds pbCardDeck.CardDeckService) *Notifications {
	return &Notifications{store: s, cardDeckService: cds}
}

func (e *Notifications) Enroll(ctx context.Context, req *pb.PushSubscriptionRequest, rsp *pbCommon.Success) error {
	logger.Infof("Received Notifications.Enroll request: %v", req)
	subscription := &model.PushSubscription{
		UserID:   req.UserID,
		Endpoint: req.Subscription.Endpoint,
		P256DH:   req.Subscription.P256Dh,
		Auth:     req.Subscription.Auth,
	}
	if err := e.store.CreatePushSubscription(ctx, subscription); err != nil {
		return err
	}
	rsp.Success = true
	return nil
}
