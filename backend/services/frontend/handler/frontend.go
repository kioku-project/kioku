package handler

import (
	"strconv"

	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/kioku-project/kioku/pkg/helper"
	"github.com/kioku-project/kioku/pkg/model"
	pbcarddeck "github.com/kioku-project/kioku/services/carddeck/proto"
	pbcollab "github.com/kioku-project/kioku/services/collaboration/proto"
	pbuser "github.com/kioku-project/kioku/services/user/proto"
	"go-micro.dev/v4/logger"
)

type Frontend struct {
	userService          pbuser.UserService
	carddeckService      pbcarddeck.CarddeckService
	collaborationService pbcollab.CollaborationService
}

func New(userService pbuser.UserService, carddeckService pbcarddeck.CarddeckService, collaborationService pbcollab.CollaborationService) *Frontend {
	return &Frontend{userService: userService, carddeckService: carddeckService, collaborationService: collaborationService}
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
	rspLogin, err := e.userService.Login(c.Context(), &pbuser.LoginRequest{Email: reqUser.Email, Password: reqUser.Password})
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
	rspRegister, err := e.userService.Register(c.Context(), &pbuser.RegisterRequest{Email: data["email"], Name: data["name"], Password: data["password"]})
	if err != nil {
		return err
	}
	return c.SendString(rspRegister.Name)
}

func (e *Frontend) CreateDeckHandler(c *fiber.Ctx) error {
	data := map[string]string{}
	if err := c.BodyParser(&data); err != nil {
		return err
	}
	if data["userID"] == "" {
		return fiber.NewError(fiber.StatusBadRequest, "No user id given")
	}
	if data["deckName"] == "" {
		return fiber.NewError(fiber.StatusBadRequest, "No deck name given")
	}
	userID, err := strconv.ParseUint(data["userID"], 10, 64)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid user id")
	}
	rspCardDeck, err := e.carddeckService.CreateDeck(c.Context(), &pbcarddeck.DeckRequest{UserID: userID, GroupPublicID: data["groupPublicID"], DeckName: data["deckName"]})
	if err != nil {
		return err
	}
	strSuccess := rspCardDeck.PublicID
	return c.SendString(strSuccess)
}

func (e *Frontend) CreateCardHandler(c *fiber.Ctx) error {
	data := map[string]string{}
	if err := c.BodyParser(&data); err != nil {
		return err
	}
	if data["userID"] == "" {
		return fiber.NewError(fiber.StatusBadRequest, "No user id given")
	}
	if data["deckPublicID"] == "" {
		return fiber.NewError(fiber.StatusBadRequest, "No deck public id given")
	}
	if data["frontside"] == "" {
		return fiber.NewError(fiber.StatusBadRequest, "No frontside given")
	}
	if data["backside"] == "" {
		return fiber.NewError(fiber.StatusBadRequest, "No backside given")
	}
	userID, err := strconv.ParseUint(data["userID"], 10, 64)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid user id")
	}
	rspCardDeck, err := e.carddeckService.CreateCard(c.Context(), &pbcarddeck.CardRequest{UserID: userID, DeckPublicID: data["deckPublicID"], Frontside: data["frontside"], Backside: data["backside"]})
	if err != nil {
		return err
	}
	strSuccess := rspCardDeck.PublicID
	return c.SendString(strSuccess)
}
