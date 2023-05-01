package handler

import (
	"github.com/gofiber/fiber/v2"
	pblogin "github.com/kioku-project/kioku/services/login/proto"
	pbregister "github.com/kioku-project/kioku/services/register/proto"
)

type Frontend struct {
	loginService    pblogin.LoginService
	registerService pbregister.RegisterService
}

func New(loginService pblogin.LoginService, registerService pbregister.RegisterService) *Frontend {
	return &Frontend{loginService: loginService, registerService: registerService}
}

func (e *Frontend) LoginHandler(c *fiber.Ctx) error {
	data := map[string]string{}
	if err := c.BodyParser(&data); err != nil {
		return err
	}
	if data["email"] == "" {
		return fiber.NewError(fiber.StatusBadRequest, "No E-Mail given")
	}
	if data["password"] == "" {
		return fiber.NewError(fiber.StatusBadRequest, "No Password given")
	}
	rspLogin, err := e.loginService.Login(c.Context(), &pblogin.LoginRequest{Email: data["email"], Password: data["password"]})
	if err != nil {
		return err
	}
	return c.SendString(rspLogin.Name)
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
