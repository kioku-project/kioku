package util

import (
	"encoding/json"
	"os"

	"github.com/SherClockHolmes/webpush-go"
	"github.com/kioku-project/kioku/pkg/model"
	"go-micro.dev/v4/logger"
)

func SendNotification(subscription *model.PushSubscription, notification *model.PushNotification) {
	privateKey, success := os.LookupEnv("VAPID_PRIVATE_KEY")
	if !success {
		logger.Fatal("VAPID_PRIVATE_KEY not set")
	}
	publicKey, success := os.LookupEnv("VAPID_PUBLIC_KEY")
	if !success {
		logger.Fatal("VAPID_PUBLIC_KEY not set")
	}

	s := &webpush.Subscription{
		Endpoint: subscription.Endpoint,
		Keys: webpush.Keys{
			P256dh: subscription.P256DH,
			Auth:   subscription.Auth,
		},
	}
	jsonNotification, err := json.Marshal(notification)
	if err != nil {
		logger.Errorf("Cronjob: Error while marshalling subscriptions: %s", err)
		logger.Info(notification)
	}

	resp, err := webpush.SendNotification(jsonNotification, s, &webpush.Options{
		Subscriber:      "web-push@kioku.dev",
		VAPIDPublicKey:  publicKey,
		VAPIDPrivateKey: privateKey,
		TTL:             30,
	})
	if err != nil {
		logger.Errorf("Cronjob: Error while sending push message: %s", err)
	}
	defer resp.Body.Close()
}
