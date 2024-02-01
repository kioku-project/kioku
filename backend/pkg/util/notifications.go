package util

import (
	"context"
	"encoding/json"

	"github.com/SherClockHolmes/webpush-go"
	"github.com/kioku-project/kioku/pkg/model"
	"go-micro.dev/v4/logger"
)

type PushHandler struct {
	privateVapidKey string
	publicVapidKey  string
}

func NewNotification(publicVapidKey string, privateVapidKey string) *PushHandler {
	return &PushHandler{
		privateVapidKey: privateVapidKey,
		publicVapidKey:  publicVapidKey,
	}
}

func (ph *PushHandler) SendNotification(ctx context.Context, subscription *model.PushSubscription, notification *model.PushNotification) error {
	jsonNotification, err := json.Marshal(notification)
	if err != nil {
		return err
	}

	resp, err := webpush.SendNotification(jsonNotification,
		&webpush.Subscription{
			Endpoint: subscription.Endpoint,
			Keys: webpush.Keys{
				P256dh: subscription.P256DH,
				Auth:   subscription.Auth,
			},
		},
		&webpush.Options{
			Subscriber:      "web-push@kioku.dev",
			VAPIDPublicKey:  ph.publicVapidKey,
			VAPIDPrivateKey: ph.privateVapidKey,
			TTL:             30,
		})
	if err != nil {
		logger.Errorf("Error while sending push message: %s", err)
		return err
	}
	defer resp.Body.Close()
	return nil
}
