package handler

import (
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/kioku-project/kioku/pkg/helper"
	"github.com/kioku-project/kioku/pkg/model"
	pbCardDeck "github.com/kioku-project/kioku/services/carddeck/proto"
	pbCollaboration "github.com/kioku-project/kioku/services/collaboration/proto"
	pbUser "github.com/kioku-project/kioku/services/user/proto"
	microError "go-micro.dev/v4/errors"
	"go-micro.dev/v4/logger"
)

type Frontend struct {
	userService          pbUser.UserService
	cardDeckService      pbCardDeck.CarddeckService
	collaborationService pbCollaboration.CollaborationService
}

func New(userService pbUser.UserService, cardDeckService pbCardDeck.CarddeckService, collaborationService pbCollaboration.CollaborationService) *Frontend {
	return &Frontend{userService: userService, cardDeckService: cardDeckService, collaborationService: collaborationService}
}

func handleMicroError(err error) error {
	parsedError := microError.Parse(err.Error())
	logger.Infof("Error from %s containing code (%d) and error detail (%s)", parsedError.Id, parsedError.Code, parsedError.Detail)
	if parsedError.Code == http.StatusBadRequest {
		return helper.ErrFiberBadRequest(parsedError.Detail)
	} else if parsedError.Code == http.StatusUnauthorized {
		return helper.ErrFiberUnauthorized(parsedError.Detail)
	} else {
		return helper.ErrFiberInternalServerError(parsedError.Detail)
	}
}

func (e *Frontend) ReauthHandler(c *fiber.Ctx) error {
	tokenString := c.Cookies("refresh_token", "NOT_GIVEN")
	refreshToken, err := helper.ParseJWTToken(tokenString)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	claims, ok := refreshToken.Claims.(jwt.MapClaims)
	if !ok || !refreshToken.Valid {
		return helper.ErrFiberUnauthorized("Please re-authenticate")
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
		return helper.ErrFiberBadRequest("no e-mail given")
	}
	if reqUser.Password == "" {
		return helper.ErrFiberBadRequest("no password given")
	}
	rspLogin, err := e.userService.Login(c.Context(), &pbUser.LoginRequest{Email: reqUser.Email, Password: reqUser.Password})
	if err != nil {
		return handleMicroError(err)
	}

	// Generate encoded tokens and send them as response.
	aTExp := time.Now().Add(time.Minute * 30)
	aTString, err := helper.CreateJWTTokenString(aTExp, rspLogin.ID, reqUser.Email, rspLogin.Name)
	if err != nil {
		logger.Infof("%v", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	rTExp := time.Now().Add(time.Hour * 24 * 7)
	rTString, err := helper.CreateJWTTokenString(rTExp, rspLogin.ID, reqUser.Email, rspLogin.Name)
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
		return helper.ErrFiberBadRequest("no e-mail given")
	}
	if data["name"] == "" {
		return helper.ErrFiberBadRequest("no name given")
	}
	if data["password"] == "" {
		return helper.ErrFiberBadRequest("no password given")
	}
	rspRegister, err := e.userService.Register(c.Context(), &pbUser.RegisterRequest{Email: data["email"], Name: data["name"], Password: data["password"]})
	if err != nil {
		return handleMicroError(err)
	}
	return c.SendString(rspRegister.Name)
}

func (e *Frontend) CreateDeckHandler(c *fiber.Ctx) error {
	data := map[string]string{}
	if err := c.BodyParser(&data); err != nil {
		return err
	}
	if data["deckName"] == "" {
		return helper.ErrFiberBadRequest("no deck name given")
	}
	userID := helper.GetUserIDFromContext(c)
	rspCardDeck, err := e.cardDeckService.CreateDeck(c.Context(), &pbCardDeck.CreateDeckRequest{UserID: userID, GroupID: c.Params("groupID"), DeckName: data["deckName"]})
	if err != nil {
		return handleMicroError(err)
	}
	strSuccess := rspCardDeck.ID
	return c.SendString(strSuccess)
}

func (e *Frontend) CreateCardHandler(c *fiber.Ctx) error {
	data := map[string]string{}
	if err := c.BodyParser(&data); err != nil {
		return err
	}
	if data["frontside"] == "" {
		return helper.ErrFiberBadRequest("no card frontside given")
	}
	if data["backside"] == "" {
		return helper.ErrFiberBadRequest("no card backside given")
	}
	userID := helper.GetUserIDFromContext(c)
	rspCardDeck, err := e.cardDeckService.CreateCard(c.Context(), &pbCardDeck.CreateCardRequest{UserID: userID, DeckID: c.Params("deckID"), Frontside: data["frontside"], Backside: data["backside"]})
	if err != nil {
		return handleMicroError(err)
	}
	strSuccess := rspCardDeck.ID
	return c.SendString(strSuccess)
}

func (e *Frontend) GetUserGroupsHandler(c *fiber.Ctx) error {
	userID := helper.GetUserIDFromContext(c)
	rspUserGroups, err := e.collaborationService.GetUserGroups(c.Context(), &pbCollaboration.UserGroupsRequest{UserID: userID})
	if err != nil {
		return handleMicroError(err)
	}
	return c.JSON(rspUserGroups)
}

func (e *Frontend) GetGroupDecksHandler(c *fiber.Ctx) error {
	userID := helper.GetUserIDFromContext(c)
	rspGroupDecks, err := e.cardDeckService.GetGroupDecks(c.Context(), &pbCardDeck.GroupDecksRequest{UserID: userID, GroupID: c.Params("groupID")})
	if err != nil {
		return handleMicroError(err)
	}
	return c.JSON(rspGroupDecks)
}

func (e *Frontend) GetDeckCardsHandler(c *fiber.Ctx) error {
	userID := helper.GetUserIDFromContext(c)
	rspDeckCards, err := e.cardDeckService.GetDeckCards(c.Context(), &pbCardDeck.DeckCardsRequest{UserID: userID, DeckID: c.Params("deckID")})
	if err != nil {
		return handleMicroError(err)
	}
	return c.JSON(rspDeckCards)
}
