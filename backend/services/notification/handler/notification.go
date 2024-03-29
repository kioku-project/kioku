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

type Notification struct {
	store       store.NotificationStore
	pushHandler *util.PushHandler
	srsService  pbSrs.SrsService
}

func NewNotification(s store.NotificationStore, ph *util.PushHandler, cds pbSrs.SrsService) *Notification {
	return &Notification{store: s, pushHandler: ph, srsService: cds}
}

func (e *Notification) Subscribe(ctx context.Context, req *pbCommon.PushSubscriptionRequest, rsp *pbCommon.PushSubscription) error {
	logger.Infof("Received Notification.Subscribe request: %v", req)
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
			Actions: []map[string]string{},
			Tag:     "Kioku",
		},
	}
	if err := e.pushHandler.SendNotification(ctx, subscription, notification); err != nil {
		return err
	}
	rsp.SubscriptionID = subscription.ID
	return nil
}

func (e *Notification) GetUserNotificationSubscriptions(ctx context.Context, req *pbCommon.User, rsp *pbCommon.PushSubscriptions) error {
	logger.Infof("Received Notification.GetUserNotificationSubscriptions request: %v", req)
	subscriptions, err := e.store.FindPushSubscriptionsByUserID(ctx, req.UserID)
	if err != nil {
		return err
	}
	protoSubscriptions := converter.ConvertToTypeArray(subscriptions, converter.StoreNotificationSubscriptionToProtoNotificationSubscriptionConverter)
	*rsp = pbCommon.PushSubscriptions{Subscriptions: protoSubscriptions}
	return nil

}

func (e *Notification) Unsubscribe(ctx context.Context, req *pbCommon.PushSubscriptionRequest, rsp *pbCommon.Success) error {
	logger.Infof("Received Notification.Unsubscribe request: %v", req)
	subscription, err := e.store.FindPushSubscriptionByID(ctx, req.Subscription.SubscriptionID)
	if err != nil {
		return err
	}
	if subscription.UserID != req.UserID {
		return helper.NewMicroNotAuthorizedErr(helper.NotificationServiceID)
	}

	if err := e.store.DeletePushSubscription(ctx, subscription); err != nil {
		return err
	}
	rsp.Success = true
	return nil
}
