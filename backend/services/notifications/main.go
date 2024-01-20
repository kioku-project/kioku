package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/kioku-project/kioku/pkg/model"
	"os"

	"github.com/kioku-project/kioku/pkg/helper"
	pbCommon "github.com/kioku-project/kioku/pkg/proto"
	pbSrs "github.com/kioku-project/kioku/services/srs/proto"
	"go-micro.dev/v4/server"

	"github.com/kioku-project/kioku/services/notifications/handler"
	pb "github.com/kioku-project/kioku/services/notifications/proto"
	"github.com/kioku-project/kioku/store"

	"go-micro.dev/v4"
	"go-micro.dev/v4/logger"

	_ "github.com/go-micro/plugins/v4/registry/kubernetes"
	"github.com/go-micro/plugins/v4/wrapper/trace/opentelemetry"

	grpcc "github.com/go-micro/plugins/v4/client/grpc"
	grpcs "github.com/go-micro/plugins/v4/server/grpc"

	webpush "github.com/SherClockHolmes/webpush-go"
	"github.com/robfig/cron/v3"
)

var (
	service        = "notifications"
	version        = "latest"
	serviceAddress = fmt.Sprintf("%s%s", os.Getenv("HOSTNAME"), ":8080")
)

func main() {
	ctx := context.TODO()
	// Initialize the database connection
	dbStore, err := store.NewNotificationStore(ctx)
	if err != nil {
		logger.Fatal("Failed to initialize database:", err)
	}

	logger.Info("Trying to listen on: ", serviceAddress)

	tp, err := helper.SetupTracing(ctx, service)
	if err != nil {
		logger.Fatal("Error setting up tracer: %v", err)
	}
	defer func() {
		if err := tp.Shutdown(context.Background()); err != nil {
			logger.Error("Error shutting down tracer provider: %v", err)
		}
	}()

	// Create service
	srv := micro.NewService(
		micro.Server(grpcs.NewServer(server.Address(serviceAddress))),
		micro.Client(grpcc.NewClient()),
		micro.WrapClient(opentelemetry.NewClientWrapper(opentelemetry.WithTraceProvider(tp))),
		micro.WrapHandler(opentelemetry.NewHandlerWrapper(opentelemetry.WithTraceProvider(tp))),
		micro.WrapSubscriber(opentelemetry.NewSubscriberWrapper(opentelemetry.WithTraceProvider(tp))),
	)
	srv.Init(
		micro.Name(service),
		micro.Version(version),
		micro.Address(serviceAddress),
	)

	// Create a new instance of the service handler with the initialized database connection
	srsService := pbSrs.NewSrsService("srs", srv.Client())
	svc := handler.New(dbStore, srsService)

	// Register handler
	if err := pb.RegisterNotificationsHandler(srv.Server(), svc); err != nil {
		logger.Fatal(err)
	}

	c := cron.New()
	// TODO: change to daily at 6pm
	c.AddFunc("* * * * *", func() {
		logger.Info("Cronjob: Sending daily notification reminder")
		subscriptions, err := dbStore.FindAllPushSubscriptions(ctx)
		if err != nil {
			logger.Errorf("Cronjob: Error while gathering subscriptions: %s", err)
		}
		privateKey, success := os.LookupEnv("VAPID_PRIVATE_KEY")
		if !success {
			logger.Fatal("VAPID_PRIVATE_KEY not set")
		}
		publicKey, success := os.LookupEnv("VAPID_PUBLIC_KEY")
		if !success {
			logger.Fatal("VAPID_PUBLIC_KEY not set")
		}
		for _, subscription := range subscriptions {

			userDueRsp, err := srsService.GetUserCardsDue(ctx, &pbCommon.User{
				UserID: subscription.UserID,
			})

			s := &webpush.Subscription{
				Endpoint: subscription.Endpoint,
				Keys: webpush.Keys{
					P256dh: subscription.P256DH,
					Auth:   subscription.Auth,
				},
			}

			notification := &model.PushNotification{
				Title: "Don't forget to review your cards!",
				Options: model.PushNotificationOptions{
					Body:    fmt.Sprintf("You have %d cards in %d decks to learn", userDueRsp.DueCards, userDueRsp.DueDecks),
					Actions: []map[string]string{},
					Vibrate: []int{200, 100, 200},
					Tag:     "Kioku",
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
	})
	c.Start()

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
