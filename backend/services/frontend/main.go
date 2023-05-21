package main

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/joho/godotenv"
	"github.com/kioku-project/kioku/pkg/helper"
	pbCardDeck "github.com/kioku-project/kioku/services/carddeck/proto"
	pbCollaboration "github.com/kioku-project/kioku/services/collaboration/proto"
	"github.com/kioku-project/kioku/services/frontend/handler"
	pb "github.com/kioku-project/kioku/services/frontend/proto"
	pbUser "github.com/kioku-project/kioku/services/user/proto"

	"go-micro.dev/v4"
	"go-micro.dev/v4/logger"
	"go-micro.dev/v4/server"

	_ "github.com/go-micro/plugins/v4/registry/kubernetes"

	grpcClient "github.com/go-micro/plugins/v4/client/grpc"
	grpcServer "github.com/go-micro/plugins/v4/server/grpc"
)

var (
	service        = "frontend"
	version        = "latest"
	serviceAddress = fmt.Sprintf("%s%s", os.Getenv("HOSTNAME"), ":8080")
)

func main() {

	logger.Info("Trying to listen on: ", serviceAddress)
	_ = godotenv.Load("../.env", "../.env.example")

	// Create service
	srv := micro.NewService(
		micro.Server(grpcServer.NewServer(server.Address(serviceAddress), server.Wait(nil))),
		micro.Client(grpcClient.NewClient()),
	)
	srv.Init(
		micro.Name(service),
		micro.Version(version),
		micro.Address(serviceAddress),
	)

	// Create a new instance of the service handler with the initialized database connection
	svc := handler.New(
		pbUser.NewUserService("user", srv.Client()),
		pbCardDeck.NewCarddeckService("carddeck", srv.Client()),
		pbCollaboration.NewCollaborationService("collaboration", srv.Client()),
	)

	app := fiber.New()
	app.Post("/api/login", svc.LoginHandler)
	app.Post("/api/register", svc.RegisterHandler)
	app.Get("/api/reauth", svc.ReauthHandler)
	// JWT Middleware
	pub, err := helper.GetJWTPublicKey()
	if err != nil {
		panic("Could not parse JWT public / private keypair")
	}
	app.Use(jwtware.New(jwtware.Config{
		SigningMethod: "ES512",
		SigningKey:    pub,
	}))
	////
	// - add endpoints where authentication is needed below this block.
	////
	app.Get("/api/group", svc.GetUserGroupsHandler)
	app.Post("/api/group", svc.CreateGroupHandler)
	app.Put("/api/group/:groupID", svc.ModifyGroupHandler)
	app.Delete("/api/group/:groupID", svc.DeleteGroupHandler)
	app.Get("/api/group/:groupID/deck", svc.GetGroupDecksHandler)
	app.Post("/api/group/:groupID/deck", svc.CreateDeckHandler)
	app.Put("/api/deck/:deckID", svc.ModifyDeckHandler)
	app.Delete("/api/deck/:deckID", svc.DeleteDeckHandler)
	app.Get("/api/deck/:deckID/card", svc.GetDeckCardsHandler)
	app.Post("/api/deck/:deckID/card", svc.CreateCardHandler)
	app.Put("/api/card/:cardID", svc.ModifyCardHandler)
	app.Delete("/api/card/:cardID", svc.DeleteCardHandler)
	// Register the handler with the micro framework
	// if err := micro.RegisterHandler(srv.Server(), grpcHandler); err != nil {
	// 	logger.Fatal(err)
	// }

	// Register handler
	if err := pb.RegisterHealthHandler(srv.Server(), new(handler.Health)); err != nil {
		logger.Fatal(err)
	}

	go func() {
		if err := app.Listen(":80"); err != nil {
			logger.Fatal(err)
		}
	}()

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
