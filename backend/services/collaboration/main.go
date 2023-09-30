package main

import (
	"fmt"
	"os"

	pbCardDeck "github.com/kioku-project/kioku/services/carddeck/proto"
	pbSrs "github.com/kioku-project/kioku/services/srs/proto"

	"github.com/kioku-project/kioku/services/collaboration/handler"
	pb "github.com/kioku-project/kioku/services/collaboration/proto"
	pbUser "github.com/kioku-project/kioku/services/user/proto"
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
	service        = "collaboration"
	version        = "latest"
	serviceAddress = fmt.Sprintf("%s%s", os.Getenv("HOSTNAME"), ":8080")
)

func main() {

	// Initialize the database connection
	dbStore, err := store.NewCollaborationStore()
	if err != nil {
		logger.Fatal("Failed to initialize database:", err)
	}

	logger.Info("Trying to listen on: ", serviceAddress)

	// Create service
	srv := micro.NewService(
		micro.Server(grpcServer.NewServer(server.Address(serviceAddress), server.Wait(nil))),
		micro.Client(grpcClient.NewClient()),
		micro.WrapClient(opentelemetry.NewClientWrapper()),
		micro.WrapHandler(opentelemetry.NewHandlerWrapper()),
		micro.WrapSubscriber(opentelemetry.NewSubscriberWrapper()),
	)
	srv.Init(
		micro.Name(service),
		micro.Version(version),
		micro.Address(serviceAddress),
	)

	// Create a new instance of the service handler with the initialized database connection
	svc := handler.New(
		dbStore,
		pbUser.NewUserService("user", srv.Client()),
		pbSrs.NewSrsService("srs", srv.Client()),
		pbCardDeck.NewCardDeckService("cardDeck", srv.Client()),
	)

	// Register handler
	if err := pb.RegisterCollaborationHandler(srv.Server(), svc); err != nil {
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
