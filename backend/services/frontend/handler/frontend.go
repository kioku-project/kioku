package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/kioku-project/kioku/pkg/helper"
	"github.com/kioku-project/kioku/pkg/model"
	pblogin "github.com/kioku-project/kioku/services/login/proto"
	pbregister "github.com/kioku-project/kioku/services/register/proto"
	"go-micro.dev/v4/logger"
	"time"
)

type Frontend struct {
	loginService    pblogin.LoginService
	registerService pbregister.RegisterService
}

func New(loginService pblogin.LoginService, registerService pbregister.RegisterService) *Frontend {
	return &Frontend{loginService: loginService, registerService: registerService}
}

func (e *Frontend) ReauthHandler(c *fiber.Ctx) error {
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

	// Generate encoded tokens and send them as response.
	aTExp := time.Now().Add(time.Minute * 30)
	aTString, err := helper.CreateJWTTokenString(aTExp, rspLogin.Id, reqUser.Email, rspLogin.Name)
	if err != nil {
		logger.Infof("%v", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	rTExp := time.Now().Add(time.Hour * 24 * 7)
	rTString, err := helper.CreateJWTTokenString(rTExp, rspLogin.Id, reqUser.Email, rspLogin.Name)
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
