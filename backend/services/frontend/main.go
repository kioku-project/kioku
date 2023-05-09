package main

import (
	"crypto/x509"
	pem2 "encoding/pem"
	"fmt"
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
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

	pem, _ := pem2.Decode([]byte(os.Getenv("JWT_PRIVATE_KEY")))
	priv, err := x509.ParseECPrivateKey(pem.Bytes)
	if err != nil {
		logger.Info(err)
	}

	app.Get("/api/reauth", func(c *fiber.Ctx) error {
		tokenString := c.Cookies("refresh_token", "NOT_GIVEN")
		if tokenString == "NOT_GIVEN" {
			return fiber.NewError(fiber.StatusUnauthorized, "Please re-authenticate")
		}
		refreshToken, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Validate the alg is what you expect:
			if _, ok := token.Method.(*jwt.SigningMethodECDSA); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}
			return priv.Public(), nil
		})
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}
		claims, ok := refreshToken.Claims.(jwt.MapClaims)
		if !ok || !refreshToken.Valid {
			return fiber.NewError(fiber.StatusUnauthorized, "Please re-authenticate")
		}

		pem, _ := pem2.Decode([]byte(os.Getenv("JWT_PRIVATE_KEY")))
		priv, err := x509.ParseECPrivateKey(pem.Bytes)
		if err != nil {
			logger.Info(err)
		}

		aTExp := time.Now().Add(time.Minute * 30)
		rTExp := time.Now().Add(time.Hour * 24 * 7)

		accessClaims := jwt.MapClaims{
			"sub":   claims["sub"],
			"email": claims["email"],
			"name":  claims["name"],
			"exp":   aTExp.Unix(),
		}
		refreshClaims := jwt.MapClaims{
			"sub":   claims["sub"],
			"email": claims["email"],
			"name":  claims["name"],
			"exp":   rTExp.Unix(),
		}

		// Create token
		accessToken := jwt.NewWithClaims(jwt.SigningMethodES512, accessClaims)
		newRefreshToken := jwt.NewWithClaims(jwt.SigningMethodES512, refreshClaims)
		// Generate encoded token and send it as response.
		aTString, err := accessToken.SignedString(priv)

		if err != nil {
			logger.Infof("%v", err)
			return c.SendStatus(fiber.StatusInternalServerError)
		}
		rTString, err := newRefreshToken.SignedString(priv)
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
		SigningKey:    priv.Public(),
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
