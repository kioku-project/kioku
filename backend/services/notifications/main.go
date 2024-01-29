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

	"github.com/kioku-project/kioku/services/notifications/handler"
	pb "github.com/kioku-project/kioku/services/notifications/proto"
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
	pushHandler := notifications.New()

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
	svc := handler.New(dbStore, pushHandler, srsService)

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
		for _, subscription := range subscriptions {

			userDueRsp, err := srsService.GetUserCardsDue(ctx, &pbCommon.User{
				UserID: subscription.UserID,
			})
			if err != nil {
				logger.Fatal(err)
			}
			if userDueRsp.DueCards == 0 {
				continue
			}

			cardString := "card"
			if userDueRsp.DueCards > 1 {
				cardString += "s"
			}
			deckString := "deck"
			if userDueRsp.DueDecks > 1 {
				deckString += "s"
			}

			notification := &model.PushNotification{
				Title: "Don't forget to review your cards!",
				Options: model.PushNotificationOptions{
					Body:    fmt.Sprintf("You have %d %s in %d %s to learn", userDueRsp.DueCards, cardString, userDueRsp.DueDecks, deckString),
					Actions: []map[string]string{},
					Vibrate: []int{200, 100, 200},
					Tag:     "Kioku",
				},
			}
			pushHandler.SendNotification(subscription, notification)
		}
	})
	c.Start()

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
