package main

import (
	"fmt"
	pbCardDeck "github.com/kioku-project/kioku/services/carddeck/proto"
	"go-micro.dev/v4/server"
	"os"

	"github.com/kioku-project/kioku/services/srs/handler"
	pb "github.com/kioku-project/kioku/services/srs/proto"
	"github.com/kioku-project/kioku/store"

	"go-micro.dev/v4"
	"go-micro.dev/v4/logger"

	_ "github.com/go-micro/plugins/v4/registry/kubernetes"

	grpcc "github.com/go-micro/plugins/v4/client/grpc"
	grpcs "github.com/go-micro/plugins/v4/server/grpc"
)

var (
	service        = "srs"
	version        = "latest"
	serviceAddress = fmt.Sprintf("%s%s", os.Getenv("HOSTNAME"), ":8080")
)

func main() {

	// Initialize the database connection
	dbStore, err := store.NewSrsStore()
	if err != nil {
		logger.Fatal("Failed to initialize database:", err)
	}

	logger.Info("Trying to listen on: ", serviceAddress)

	// Create service
	srv := micro.NewService(
		micro.Server(grpcs.NewServer(server.Address(serviceAddress))),
		micro.Client(grpcc.NewClient()),
	)
	srv.Init(
		micro.Name(service),
		micro.Version(version),
		micro.Address(serviceAddress),
	)

	// Create a new instance of the service handler with the initialized database connection
	svc := handler.New(dbStore, pbCardDeck.NewCardDeckService("cardDeck", srv.Client()))

	// Register handler
	if err := pb.RegisterSrsHandler(srv.Server(), svc); err != nil {
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
