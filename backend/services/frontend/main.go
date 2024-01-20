package main

import (
	"context"
	"fmt"
	"os"

	pbNotifications "github.com/kioku-project/kioku/services/notifications/proto"
	pbSrs "github.com/kioku-project/kioku/services/srs/proto"
	microErrors "go-micro.dev/v4/errors"

	"github.com/gofiber/contrib/otelfiber"
	"github.com/gofiber/fiber/v2"
	jwtWare "github.com/gofiber/jwt/v3"
	"github.com/joho/godotenv"
	"github.com/kioku-project/kioku/pkg/helper"
	pbCardDeck "github.com/kioku-project/kioku/services/carddeck/proto"
	pbCollaboration "github.com/kioku-project/kioku/services/collaboration/proto"
	"github.com/kioku-project/kioku/services/frontend/handler"
	pbUser "github.com/kioku-project/kioku/services/user/proto"

	"go-micro.dev/v4"
	"go-micro.dev/v4/logger"
	"go-micro.dev/v4/server"

	_ "github.com/go-micro/plugins/v4/registry/kubernetes"
	"github.com/go-micro/plugins/v4/wrapper/trace/opentelemetry"

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

	tp, err := helper.SetupTracing(context.TODO(), service)
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
	svc := handler.New(
		pbUser.NewUserService("user", srv.Client()),
		pbCardDeck.NewCardDeckService("cardDeck", srv.Client()),
		pbCollaboration.NewCollaborationService("collaboration", srv.Client()),
		pbSrs.NewSrsService("srs", srv.Client()),
		pbNotifications.NewNotificationsService("notifications", srv.Client()),
	)

	fiberConfig := fiber.Config{
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			parsedError := microErrors.Parse(err.Error())
			if parsedError.Code == 0 {
				parsedError.Code = fiber.StatusInternalServerError
			}
			logger.Infof("Error from %s containing code (%d) and error detail (%s)",
				parsedError.Id, parsedError.Code, parsedError.Detail)
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
	app.Use(otelfiber.Middleware(
		otelfiber.WithTracerProvider(tp),
	))
	app.Use(jwtWare.New(jwtWare.Config{
		SigningMethod: "ES512",
		SigningKey:    pub,
	}))
	////
	// - add endpoints where authentication is needed below this block.
	////
	app.Post("/api/logout", svc.LogoutHandler)

	app.Get("/api/user", svc.GetUserHandler)
	app.Put("/api/user", svc.ModifyUserHandler)
	app.Delete("/api/user", svc.DeleteUserHandler)
	app.Get("/api/user/dueCards", svc.SrsUserDueHandler)
	app.Get("/api/user/invitations", svc.GetGroupInvitationsHandler)

	app.Get("/api/groups", svc.GetUserGroupsHandler)
	app.Post("/api/groups", svc.CreateGroupHandler)
	app.Get("/api/groups/:groupID", svc.GetGroupHandler)
	app.Put("/api/groups/:groupID", svc.ModifyGroupHandler)
	app.Delete("/api/groups/:groupID", svc.DeleteGroupHandler)

	app.Get("/api/groups/:groupID/members", svc.GetGroupMembersHandler)
	app.Delete("/api/groups/:groupID/members", svc.LeaveGroupHandler)
	app.Put("/api/groups/:groupID/members/:userID", svc.ModifyGroupMemberHandler)
	app.Delete("/api/groups/:groupID/members/:userID", svc.KickGroupMemberHandler)
	app.Get("/api/groups/:groupID/members/requests", svc.GetGroupMemberRequestsHandler)
	app.Get("/api/groups/:groupID/members/invitations", svc.GetInvitationsForGroupHandler)

	app.Post("/api/groups/:groupID/members/invitation", svc.AddUserGroupInviteHandler)
	app.Delete("/api/groups/:groupID/members/invitation", svc.RemoveUserGroupInviteHandler)
	app.Post("/api/groups/:groupID/members/request", svc.AddUserGroupRequestHandler)
	app.Delete("/api/groups/:groupID/members/request", svc.RemoveUserGroupRequestHandler)

	app.Get("/api/decks/favorites", svc.GetFavoriteDecksHandler)
	app.Post("/api/decks/favorites", svc.AddFavoriteDeckHandler)
	app.Delete("/api/decks/favorites", svc.DeleteFavoriteDeckHandler)
	app.Get("/api/decks/active", svc.GetActiveDecksHandler)
	app.Delete("/api/decks/active", svc.DeleteActiveDeckHandler)

	app.Get("/api/groups/:groupID/decks", svc.GetGroupDecksHandler)
	app.Post("/api/groups/:groupID/decks", svc.CreateDeckHandler)
	app.Get("/api/decks/:deckID", svc.GetDeckHandler)
	app.Put("/api/decks/:deckID", svc.ModifyDeckHandler)
	app.Delete("/api/decks/:deckID", svc.DeleteDeckHandler)
	app.Post("/api/decks/:deckID/copy", svc.CopyDeckHandler)

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

	app.Post("/api/user/notifications", svc.SubscribeNotificationsHandler)
	app.Delete("/api/user/notifications/:subscriptionID", svc.UnsubscribeNotificationsHandler)

	// Register the handler with the micro framework
	// if err := micro.RegisterHandler(srv.Server(), grpcHandler); err != nil {
	// 	logger.Fatal(err)
	// }

	// Register handler

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
