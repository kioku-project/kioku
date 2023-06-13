package main

import (
	"fmt"
	pbSrs "github.com/kioku-project/kioku/services/srs/proto"
	microErrors "go-micro.dev/v4/errors"
	"os"

	"github.com/gofiber/fiber/v2"
	jwtWare "github.com/gofiber/jwt/v3"
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
		pbCardDeck.NewCardDeckService("cardDeck", srv.Client()),
		pbCollaboration.NewCollaborationService("collaboration", srv.Client()),
		pbSrs.NewSrsService("srs", srv.Client()),
	)

	fiberConfig := fiber.Config{
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			parsedError := microErrors.Parse(err.Error())
			if parsedError.Code == 0 {
				parsedError.Code = fiber.StatusInternalServerError
			}
			logger.Infof("Error from %s containing code (%d) and error detail (%s)", parsedError.Id, parsedError.Code, parsedError.Detail)
			return ctx.Status(int(parsedError.Code)).SendString(parsedError.Detail)
		},
	}

	app := fiber.New(fiberConfig)
	app.Post("/api/register", svc.RegisterHandler)
	app.Post("/api/login", svc.LoginHandler)
	app.Get("/api/reauth", svc.ReauthHandler)
	// JWT Middleware
	pub, err := helper.GetJWTPublicKey()
	if err != nil {
		panic("Could not parse JWT public / private keypair")
	}
	app.Use(jwtWare.New(jwtWare.Config{
		SigningMethod: "ES512",
		SigningKey:    pub,
	}))
	////
	// - add endpoints where authentication is needed below this block.
	////
	app.Post("/api/logout", svc.LogoutHandler)

	app.Get("/api/user", svc.GetUserHandler)
	app.Get("/api/user/dueCards", svc.SrsUserDueHandler)
	app.Get("/api/user/invitations", svc.GetGroupInvitationsHandler)

	app.Get("/api/groups", svc.GetUserGroupsHandler)
	app.Post("/api/groups", svc.CreateGroupHandler)
	app.Get("/api/groups/:groupID", svc.GetGroupHandler)
	app.Put("/api/groups/:groupID", svc.ModifyGroupHandler)
	app.Delete("/api/groups/:groupID", svc.DeleteGroupHandler)

	app.Get("/api/groups/:groupID/members", svc.GetGroupMembersHandler)
	app.Get("/api/groups/:groupID/members/requests", svc.GetGroupMemberRequestsHandler)
	app.Get("/api/groups/:groupID/members/invitations", svc.GetInvitationsForGroupHandler)

	app.Post("/api/groups/:groupID/members/invitation", svc.AddUserGroupInviteHandler)
	app.Delete("/api/groups/:groupID/members/invitation", svc.RemoveUserGroupInviteHandler)
	app.Post("/api/groups/:groupID/members/request", svc.AddUserGroupRequestHandler)
	app.Delete("/api/groups/:groupID/members/request", svc.RemoveUserGroupRequestHandler)

	app.Get("/api/groups/:groupID/decks", svc.GetGroupDecksHandler)
	app.Post("/api/groups/:groupID/decks", svc.CreateDeckHandler)
	app.Get("/api/decks/:deckID", svc.GetDeckHandler)
	app.Put("/api/decks/:deckID", svc.ModifyDeckHandler)
	app.Delete("/api/decks/:deckID", svc.DeleteDeckHandler)

	app.Get("/api/decks/:deckID/cards", svc.GetDeckCardsHandler)
	app.Post("/api/decks/:deckID/cards", svc.CreateCardHandler)
	app.Get("/api/cards/:cardID", svc.GetCardHandler)
	app.Put("/api/cards/:cardID", svc.ModifyCardHandler)
	app.Delete("/api/cards/:cardID", svc.DeleteCardHandler)

	app.Post("/api/cards/:cardID/cardSides", svc.CreateCardSideHandler)
	app.Put("/api/cardSides/:cardSideID", svc.ModifyCardSideHandler)
	app.Delete("/api/cardSides/:cardSideID", svc.DeleteCardSideHandler)

	app.Get("/api/decks/:deckID/pull", svc.SrsPullHandler)
	app.Post("/api/decks/:deckID/push", svc.SrsPushHandler)
	app.Get("/api/decks/:deckID/dueCards", svc.SrsDeckDueHandler)

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
