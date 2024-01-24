package handler

import (
	"context"

	"github.com/SherClockHolmes/webpush-go"
	"github.com/kioku-project/kioku/pkg/helper"
	"github.com/kioku-project/kioku/pkg/model"
	pbCommon "github.com/kioku-project/kioku/pkg/proto"
	"github.com/kioku-project/kioku/pkg/util"
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

func (e *Notifications) Subscribe(ctx context.Context, req *pb.PushSubscriptionRequest, rsp *pb.PushSubscription) error {
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
	notification := &model.PushNotification{
		Title: "Welcome to Kioku!",
		Options: model.PushNotificationOptions{
			Body:    "You will now receive reminders so you don't forget your cards'!",
			Actions: []map[string]string{},
			Vibrate: []int{200, 100, 200},
			Tag:     "Kioku",
		},
	}
	util.SendNotification(subscription, notification)
	rsp.SubscriptionID = subscription.ID
	return nil
}

func (e *Notifications) Unsubscribe(ctx context.Context, req *pb.PushSubscriptionRequest, rsp *pbCommon.Success) error {
	logger.Infof("Received Notifications.Unenroll request: %v", req)
	subscription, err := e.store.FindPushSubscriptionByID(ctx, req.Subscription.SubscriptionID)
	if err != nil {
		return err
	}
	if subscription.UserID != req.UserID {
		return helper.NewMicroNotAuthorizedErr(helper.NotificationsServiceID)
	}

	if err := e.store.DeletePushSubscription(ctx, subscription); err != nil {
		return err
	}
	rsp.Success = true
	return nil
}
