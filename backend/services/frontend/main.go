package main

import (
	"context"
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/joho/godotenv"
	helper "github.com/kioku-project/kioku/pkg/helper"
	pbcarddeck "github.com/kioku-project/kioku/services/carddeck/proto"
	pbcollab "github.com/kioku-project/kioku/services/collaboration/proto"
	"github.com/kioku-project/kioku/services/frontend/handler"
	pb "github.com/kioku-project/kioku/services/frontend/proto"
	pbuser "github.com/kioku-project/kioku/services/user/proto"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"github.com/gofiber/contrib/otelfiber"

	"go-micro.dev/v4"
	"go-micro.dev/v4/logger"
	"go-micro.dev/v4/server"

	_ "github.com/go-micro/plugins/v4/registry/kubernetes"

	grpcc "github.com/go-micro/plugins/v4/client/grpc"
	grpcs "github.com/go-micro/plugins/v4/server/grpc"
)

var (
	service        = "frontend"
	version        = "latest"
	serviceAddress = fmt.Sprintf("%s%s", os.Getenv("HOSTNAME"), ":8080")
	jaegerUrl = os.Getenv("JAEGER_ADDRESS")
	tracer = otel.Tracer(service)
)

func main() {

	logger.Info("Trying to listen on: ", serviceAddress)
	_ = godotenv.Load("../.env", "../.env.example")

	tp, err := helper.NewTracerProvider(service, service, version, jaegerUrl)

	if err != nil {
		logger.Error("Error starting tracer provider: %v", err)
	}

	defer func() {
		if err := tp.Shutdown(context.Background()); err != nil {
			logger.Error("Error shutting down tracer provider: %v", err)
		}
	}()

	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))

	// Create service
	srv := micro.NewService(
		micro.Server(grpcs.NewServer(server.Address(serviceAddress), server.Wait(nil))),
		micro.Client(grpcc.NewClient()),
	)
	srv.Init(
		micro.Name(service),
		micro.Version(version),
		micro.Address(serviceAddress),
	)

	// Create a new instance of the service handler with the initialized database connection
	svc := handler.New(
		pbuser.NewUserService("user", srv.Client()),
		pbcarddeck.NewCarddeckService("carddeck", srv.Client()),
		pbcollab.NewCollaborationService("collaboration", srv.Client()),
		tracer,
	)

	app := fiber.New()
	app.Use(otelfiber.Middleware())
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
	app.Get("/api/group/:groupID/deck", svc.GetGroupDecksHandler)
	app.Post("/api/group/:groupID/deck", svc.CreateDeckHandler)
	app.Get("/api/deck/:deckID/card", svc.GetDeckCardsHandler)
	app.Post("/api/deck/:deckID/card", svc.CreateCardHandler)

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
