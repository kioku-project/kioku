package main

import (
	"context"
	"fmt"
	"os"

	"github.com/kioku-project/kioku/pkg/helper"
	pbCollaboration "github.com/kioku-project/kioku/services/collaboration/proto"
	"github.com/kioku-project/kioku/services/user/handler"
	pb "github.com/kioku-project/kioku/services/user/proto"
	"github.com/kioku-project/kioku/store"

	"go-micro.dev/v4"
	"go-micro.dev/v4/logger"
	"go-micro.dev/v4/server"

	_ "github.com/go-micro/plugins/v4/registry/kubernetes"
	"github.com/go-micro/plugins/v4/wrapper/trace/opentelemetry"

	grpcClient "github.com/go-micro/plugins/v4/client/grpc"
	grpcServer "github.com/go-micro/plugins/v4/server/grpc"
)

var (
	service        = "user"
	version        = "latest"
	serviceAddress = fmt.Sprintf("%s%s", os.Getenv("HOSTNAME"), ":8080")
)

func main() {

	// Initialize the database connection
	dbStore, err := store.NewUserStore()
	if err != nil {
		logger.Fatal("Failed to initialize database:", err)
	}

	logger.Info("Trying to listen on: ", serviceAddress)

	tp, err := helper.SetupTracing(context.TODO(), service)
	if err != nil {
		logger.Error("Error setting up tracer: %v", err)
	}
	defer func() {
		if err := tp.Shutdown(context.Background()); err != nil {
			logger.Error("Error shutting down tracer provider: %v", err)
		}
	}()

	// Create service
	srv := micro.NewService(
		micro.Server(grpcServer.NewServer(server.Address(serviceAddress), server.Wait(nil))),
		micro.Client(grpcClient.NewClient()),
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
	svc := handler.New(dbStore, pbCollaboration.NewCollaborationService("collaboration", srv.Client()))

	// Register handler
	if err := pb.RegisterUserHandler(srv.Server(), svc); err != nil {
		logger.Fatal(err)
	}
	if err := pb.RegisterHealthHandler(srv.Server(), new(handler.Health)); err != nil {
		logger.Fatal(err)
	}

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
