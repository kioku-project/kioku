package util

import (
	"encoding/json"
	"os"

	"github.com/SherClockHolmes/webpush-go"
	"github.com/kioku-project/kioku/pkg/model"
	"go-micro.dev/v4/logger"
)

type PushHandler struct {
	privateVapidKey string
	publicVapidKey  string
}

func New() *PushHandler {
	privateVapidKey, success := os.LookupEnv("VAPID_PRIVATE_KEY")
	if !success {
		logger.Fatal("VAPID_PRIVATE_KEY not set")
	}
	publicVapidKey, success := os.LookupEnv("VAPID_PUBLIC_KEY")
	if !success {
		logger.Fatal("VAPID_PUBLIC_KEY not set")
	}
	return &PushHandler{
		privateVapidKey: privateVapidKey,
		publicVapidKey:  publicVapidKey,
	}
}

func (ph *PushHandler) SendNotification(subscription *model.PushSubscription, notification *model.PushNotification) {
	s := &webpush.Subscription{
		Endpoint: subscription.Endpoint,
		Keys: webpush.Keys{
			P256dh: subscription.P256DH,
			Auth:   subscription.Auth,
		},
	}
	jsonNotification, err := json.Marshal(notification)
	if err != nil {
		logger.Errorf("Error while marshalling subscriptions: %s", err)
		logger.Info(notification)
		return
	}

	resp, err := webpush.SendNotification(jsonNotification, s, &webpush.Options{
		Subscriber:      "web-push@kioku.dev",
		VAPIDPublicKey:  ph.publicVapidKey,
		VAPIDPrivateKey: ph.privateVapidKey,
		TTL:             30,
	})
	if err != nil {
		logger.Errorf("Error while sending push message: %s", err)
	}
	defer resp.Body.Close()
}
