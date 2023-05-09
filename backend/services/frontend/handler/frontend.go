package handler

import (
	"crypto/x509"
	pem2 "encoding/pem"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/kioku-project/kioku/pkg/model"
	pblogin "github.com/kioku-project/kioku/services/login/proto"
	pbregister "github.com/kioku-project/kioku/services/register/proto"
	"go-micro.dev/v4/logger"
	"os"
	"time"
)

type Frontend struct {
	loginService    pblogin.LoginService
	registerService pbregister.RegisterService
}

func New(loginService pblogin.LoginService, registerService pbregister.RegisterService) *Frontend {
	return &Frontend{loginService: loginService, registerService: registerService}
}
func (e *Frontend) LoginHandler(c *fiber.Ctx) error {
	var reqUser model.User
	if err := c.BodyParser(&reqUser); err != nil {
		return err
	}
	if reqUser.Email == "" {
		return fiber.NewError(fiber.StatusBadRequest, "No E-Mail given")
	}
	if reqUser.Password == "" {
		return fiber.NewError(fiber.StatusBadRequest, "No Password given")
	}
	rspLogin, err := e.loginService.Login(c.Context(), &pblogin.LoginRequest{Email: reqUser.Email, Password: reqUser.Password})
	if err != nil {
		return err
	}

	pem, _ := pem2.Decode([]byte(os.Getenv("JWT_PRIVATE_KEY")))
	priv, err := x509.ParseECPrivateKey(pem.Bytes)
	if err != nil {
		logger.Info(err)
		return fiber.NewError(fiber.StatusInternalServerError, "Could not load private key for jwt signing")
	}

	aTExp := time.Now().Add(time.Minute * 30)
	rTExp := time.Now().Add(time.Hour * 24 * 7)
	accessClaims := jwt.MapClaims{
		"sub":   rspLogin.Id,
		"email": reqUser.Email,
		"name":  rspLogin.Name,
		"exp":   aTExp.Unix(),
	}
	refreshClaims := jwt.MapClaims{
		"sub":   rspLogin.Id,
		"email": reqUser.Email,
		"name":  rspLogin.Name,
		"exp":   rTExp.Unix(),
	}

	// Create token
	accessToken := jwt.NewWithClaims(jwt.SigningMethodES512, accessClaims)
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodES512, refreshClaims)

	// Generate encoded tokens and send them as response.
	aTString, err := accessToken.SignedString(priv)
	if err != nil {
		logger.Infof("%v", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	rTString, err := refreshToken.SignedString(priv)
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
}

func (e *Frontend) RegisterHandler(c *fiber.Ctx) error {
	data := map[string]string{}
	if err := c.BodyParser(&data); err != nil {
		return err
	}
	if data["email"] == "" {
		return fiber.NewError(fiber.StatusBadRequest, "No E-Mail given")
	}
	if data["name"] == "" {
		return fiber.NewError(fiber.StatusBadRequest, "No Name given")
	}
	if data["password"] == "" {
		return fiber.NewError(fiber.StatusBadRequest, "No Password given")
	}
	rspRegister, err := e.registerService.Register(c.Context(), &pbregister.RegisterRequest{Email: data["email"], Name: data["name"], Password: data["password"]})
	if err != nil {
		return err
	}
	return c.SendString(rspRegister.Name)
}
