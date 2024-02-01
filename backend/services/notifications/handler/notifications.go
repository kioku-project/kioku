package handler

import (
	"context"

	"github.com/kioku-project/kioku/pkg/converter"
	"github.com/kioku-project/kioku/pkg/helper"
	"github.com/kioku-project/kioku/pkg/model"
	pbCommon "github.com/kioku-project/kioku/pkg/proto"
	"github.com/kioku-project/kioku/pkg/util"
	pbSrs "github.com/kioku-project/kioku/services/srs/proto"
	"go-micro.dev/v4/logger"

	"github.com/kioku-project/kioku/store"
)

type Notifications struct {
	store       store.NotificationsStore
	pushHandler *util.PushHandler
	srsService  pbSrs.SrsService
}

func NewNotifications(s store.NotificationsStore, ph *util.PushHandler, cds pbSrs.SrsService) *Notifications {
	return &Notifications{store: s, pushHandler: ph, srsService: cds}
}

func (e *Notifications) Subscribe(ctx context.Context, req *pbCommon.PushSubscriptionRequest, rsp *pbCommon.PushSubscription) error {
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
		Title: "Welcome!",
		Options: model.PushNotificationOptions{
			Body:    "Reminders will be sent to help you keep track of due cards.",
			Vibrate: []int{200, 100, 200},
			Tag:     "Kioku",
		},
	}
	if err := e.pushHandler.SendNotification(ctx, subscription, notification); err != nil {
		return err
	}
	rsp.SubscriptionID = subscription.ID
	return nil
}

func (e *Notifications) GetUserNotificationSubscriptions(ctx context.Context, req *pbCommon.User, rsp *pbCommon.PushSubscriptions) error {
	logger.Infof("Received Notifications.GetUserNotificationSubscriptions request: %v", req)
	subscriptions, err := e.store.FindPushSubscriptionsByUserID(ctx, req.UserID)
	if err != nil {
		return err
	}
	protoSubscriptions := converter.ConvertToTypeArray(subscriptions, converter.StoreNotificationSubscriptionToProtoNotificationSubscriptionConverter)
	*rsp = pbCommon.PushSubscriptions{Subscriptions: protoSubscriptions}
	return nil

}

func (e *Notifications) Unsubscribe(ctx context.Context, req *pbCommon.PushSubscriptionRequest, rsp *pbCommon.Success) error {
	logger.Infof("Received Notifications.Unsubscribe request: %v", req)
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
