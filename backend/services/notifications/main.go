package main

import (
	"context"
	"fmt"
	"os"

	"github.com/kioku-project/kioku/pkg/helper"
	pbCardDeck "github.com/kioku-project/kioku/services/carddeck/proto"
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
	svc := handler.New(dbStore, pbCardDeck.NewCardDeckService("cardDeck", srv.Client()))

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
			// TODO: Handle error
		}
		privateKey, publicKey := "", ""
		if err != nil {
			// TODO: Handle error
		}
		for _, subscription := range subscriptions {
						s := &webpush.Subscription{
				Endpoint: subscription.Endpoint,
				Keys: webpush.Keys{
					P256dh: subscription.P256DH,
					Auth:   subscription.Auth,
				},
			}
			resp, err := webpush.SendNotification([]byte("Test"), s, &webpush.Options{
				VAPIDPublicKey:  publicKey,
				VAPIDPrivateKey: privateKey,
				TTL:             30,
			})
			if err != nil {
				// TODO: Handle error
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
