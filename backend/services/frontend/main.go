package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
	helper "github.com/kioku-project/kioku/pkg/helper"
	"os"
	"time"

	"github.com/kioku-project/kioku/services/frontend/handler"
	pb "github.com/kioku-project/kioku/services/frontend/proto"
	pblogin "github.com/kioku-project/kioku/services/login/proto"
	pbregister "github.com/kioku-project/kioku/services/register/proto"

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
)

func main() {

	logger.Info("Trying to listen on: ", serviceAddress)
	_ = godotenv.Load("../.env", "../.env.example")

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
	svc := handler.New(
		pblogin.NewLoginService("login", srv.Client()),
		pbregister.NewRegisterService("register", srv.Client()),
	)

	app := fiber.New()
	app.Post("/api/login", svc.LoginHandler)
	app.Post("/api/register", svc.RegisterHandler)

	app.Get("/api/reauth", func(c *fiber.Ctx) error {
		tokenString := c.Cookies("refresh_token", "NOT_GIVEN")
		refreshToken, err := helper.ParseJWTToken(tokenString)
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}
		claims, ok := refreshToken.Claims.(jwt.MapClaims)
		if !ok || !refreshToken.Valid {
			return fiber.NewError(fiber.StatusUnauthorized, "Please re-authenticate")
		}

		// Generate encoded token and send it as response.
		aTExp := time.Now().Add(time.Minute * 30)
		aTString, err := helper.CreateJWTTokenString(aTExp, claims["sub"], claims["email"], claims["name"])
		if err != nil {
			logger.Infof("%v", err)
			return c.SendStatus(fiber.StatusInternalServerError)
		}
		rTExp := time.Now().Add(time.Hour * 24 * 7)
		rTString, err := helper.CreateJWTTokenString(rTExp, claims["sub"], claims["email"], claims["name"])
		if err != nil {
			logger.Infof("%v", err)
			return c.SendStatus(fiber.StatusInternalServerError)
		}

		c.Cookie(&fiber.Cookie{
			Name:    "access_token",
			Value:   aTString,
			Path:    "/",
			Expires: aTExp,
		})
		c.Cookie(&fiber.Cookie{
			Name:     "refresh_token",
			Value:    rTString,
			Path:     "/",
			Expires:  rTExp,
			HTTPOnly: true,
		})
		return c.SendStatus(200)
	})
	// JWT Middleware
	app.Use(jwtware.New(jwtware.Config{
		SigningMethod: "ES512",
		SigningKey:    helper.GetJWTPublicKey(),
	}))
	////
	// - add endpoints where authentication is needed below this block.
	////

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
