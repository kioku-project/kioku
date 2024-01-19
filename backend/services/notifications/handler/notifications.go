package handler

import (
	"context"

	"github.com/kioku-project/kioku/pkg/model"
	pbCommon "github.com/kioku-project/kioku/pkg/proto"
	pbSrs "github.com/kioku-project/kioku/services/srs/proto"
	"go-micro.dev/v4/logger"

	pb "github.com/kioku-project/kioku/services/notifications/proto"
	"github.com/kioku-project/kioku/store"
)

type Notifications struct {
	store      store.NotificationsStore
	srsService pbSrs.SrsService
}

func New(s store.NotificationsStore, cds pbSrs.SrsService) *Notifications {
	return &Notifications{store: s, srsService: cds}
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
