package main

import (
	"go-micro.dev/v4/server"

	"github.com/kioku-project/kioku/services/register/handler"
	"github.com/kioku-project/kioku/store"
	pb "github.com/kioku-project/kioku/services/register/proto"
	pblogin "github.com/kioku-project/kioku/services/login/proto"

	"go-micro.dev/v4"
	"go-micro.dev/v4/logger"

	grpcc "github.com/go-micro/plugins/v4/client/grpc"
	grpcs "github.com/go-micro/plugins/v4/server/grpc"
)

var (
	service = "register"
	version = "latest"
)

const (
    servicePort = ":8080" // You can change this to your desired port
)

func main() {


	// Initialize the database connection
	dbStore, err := store.NewPostgresStore()
	if err != nil {
		logger.Fatal("Failed to initialize database:", err)
	}

	// Create service
	srv := micro.NewService(
		micro.Server(grpcs.NewServer(server.Address(servicePort))),
		micro.Client(grpcc.NewClient()),
	)
	srv.Init(
		micro.Name(service),
		micro.Version(version),
	)

	// Create a new instance of the service handler with the initialized database connection
	svc := handler.New(dbStore, pblogin.NewLoginService("login", srv.Client()))

	// Register handler
	if err := pb.RegisterRegisterHandler(srv.Server(), svc); err != nil {
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
