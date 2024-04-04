package main

import (
	"context"
	"fmt"
	"os"

	"github.com/kioku-project/kioku/pkg/helper"
	pbCardDeck "github.com/kioku-project/kioku/services/carddeck/proto"
	"go-micro.dev/v4/server"

	"github.com/kioku-project/kioku/services/testSrs/handler"
	pb "github.com/kioku-project/kioku/services/testSrs/proto"
	"github.com/kioku-project/kioku/store"

	"go-micro.dev/v4"
	"go-micro.dev/v4/logger"

	_ "github.com/go-micro/plugins/v4/registry/kubernetes"
	"github.com/go-micro/plugins/v4/wrapper/trace/opentelemetry"

	grpcc "github.com/go-micro/plugins/v4/client/grpc"
	grpcs "github.com/go-micro/plugins/v4/server/grpc"
)

var (
	service        = "testSrs"
	version        = "latest"
	serviceAddress = fmt.Sprintf("%s%s", os.Getenv("HOSTNAME"), ":8080")
)

func main() {
	ctx := context.TODO()
	// Initialize the database connection
	dbStore, err := store.NewSrsStore(ctx)
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
	if err := pb.RegisterTestSrsHandler(srv.Server(), svc); err != nil {
		logger.Fatal(err)
	}

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
