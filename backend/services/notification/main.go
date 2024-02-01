package main

import (
	"context"
	"fmt"
	"os"

	"github.com/kioku-project/kioku/pkg/model"
	"github.com/kioku-project/kioku/pkg/util"

	"github.com/kioku-project/kioku/pkg/helper"
	pbCommon "github.com/kioku-project/kioku/pkg/proto"
	pbSrs "github.com/kioku-project/kioku/services/srs/proto"
	"go-micro.dev/v4/server"

	"github.com/kioku-project/kioku/services/notification/handler"
	pb "github.com/kioku-project/kioku/services/notification/proto"
	"github.com/kioku-project/kioku/store"

	"go-micro.dev/v4"
	"go-micro.dev/v4/logger"

	_ "github.com/go-micro/plugins/v4/registry/kubernetes"
	"github.com/go-micro/plugins/v4/wrapper/trace/opentelemetry"

	grpcc "github.com/go-micro/plugins/v4/client/grpc"
	grpcs "github.com/go-micro/plugins/v4/server/grpc"

	"github.com/robfig/cron/v3"
)

var (
	service        = "notification"
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
	privateVapidKey, success := os.LookupEnv("VAPID_PRIVATE_KEY")
	if !success {
		logger.Fatal("VAPID_PRIVATE_KEY not set")
	}
	publicVapidKey, success := os.LookupEnv("VAPID_PUBLIC_KEY")
	if !success {
		logger.Fatal("VAPID_PUBLIC_KEY not set")
	}
	pushHandler := util.NewNotification(publicVapidKey, privateVapidKey)

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
	svc := handler.NewNotification(dbStore, pushHandler, srsService)

	// Register handler
	if err := pb.RegisterNotificationHandler(srv.Server(), svc); err != nil {
		logger.Fatal(err)
	}

	c := cron.New()
	c.AddFunc("0 11,14,17 * * *", func() {
		logger.Info("Cronjob: Sending daily notification reminder")
		subscriptions, err := dbStore.FindAllPushSubscriptions(ctx)
		if err != nil {
			logger.Errorf("Cronjob: Error while gathering subscriptions: %s", err)
			return
		}
		for _, subscription := range subscriptions {
			userDueRsp, err := srsService.GetUserCardsDue(ctx, &pbCommon.User{
				UserID: subscription.UserID,
			})
			if err != nil {
				logger.Error(err)
				continue
			}
			if userDueRsp.DueCards == 0 {
				continue
			}

			notification := &model.PushNotification{
				Title: "Don't forget to review your cards!",
				Options: model.PushNotificationOptions{
					Body: fmt.Sprintf("You have %d %s in %d %s to learn",
						userDueRsp.DueCards,
						util.PluralSingularSelector(userDueRsp.DueCards, "card", "cards"),
						userDueRsp.DueDecks,
						util.PluralSingularSelector(userDueRsp.DueDecks, "deck", "decks")),
					Vibrate: []int{200, 100, 200},
					Actions: []map[string]string{},
					Tag:     "Kioku",
				},
			}
			if err := pushHandler.SendNotification(ctx, subscription, notification); err != nil {
				logger.Errorf("Cronjob: Error while sending push message: %s", err)
			}
		}
	})
	c.Start()

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
