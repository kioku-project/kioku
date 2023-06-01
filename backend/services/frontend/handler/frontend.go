package handler

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/kioku-project/kioku/pkg/converter"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/kioku-project/kioku/pkg/helper"
	"github.com/kioku-project/kioku/pkg/model"
	pbCardDeck "github.com/kioku-project/kioku/services/carddeck/proto"
	pbCollaboration "github.com/kioku-project/kioku/services/collaboration/proto"
	pbUser "github.com/kioku-project/kioku/services/user/proto"
	"go-micro.dev/v4/logger"
)

type Frontend struct {
	userService          pbUser.UserService
	cardDeckService      pbCardDeck.CardDeckService
	collaborationService pbCollaboration.CollaborationService
}

func New(
	userService pbUser.UserService,
	cardDeckService pbCardDeck.CardDeckService,
	collaborationService pbCollaboration.CollaborationService,
) *Frontend {
	return &Frontend{
		userService:          userService,
		cardDeckService:      cardDeckService,
		collaborationService: collaborationService,
	}
}

func (e *Frontend) RegisterHandler(c *fiber.Ctx) error {
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return err
	}
	if strings.TrimSpace(data["email"]) == "" {
		return helper.NewFiberBadRequestErr("no e-mail given")
	}
	if strings.TrimSpace(data["name"]) == "" {
		return helper.NewFiberBadRequestErr("no name given")
	}
	if strings.TrimSpace(data["password"]) == "" {
		return helper.NewFiberBadRequestErr("no password given")
	}
	rspRegister, err := e.userService.Register(c.Context(), &pbUser.RegisterRequest{
		Email:    data["email"],
		Name:     data["name"],
		Password: data["password"]})
	if err != nil {
		return err
	}
	return c.SendString(rspRegister.Name)
}

func (e *Frontend) LoginHandler(c *fiber.Ctx) error {
	var reqUser model.User
	if err := c.BodyParser(&reqUser); err != nil {
		return err
	}
	if strings.TrimSpace(reqUser.Email) == "" {
		return helper.NewFiberBadRequestErr("no e-mail given")
	}
	if strings.TrimSpace(reqUser.Password) == "" {
		return helper.NewFiberBadRequestErr("no password given")
	}
	rspLogin, err := e.userService.Login(c.Context(), &pbUser.LoginRequest{
		Email:    reqUser.Email,
		Password: reqUser.Password})
	if err != nil {
		return err
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

func (e *Frontend) ReauthHandler(c *fiber.Ctx) error {
	tokenString := c.Cookies("refresh_token")
	refreshToken, err := helper.ParseJWTToken(tokenString)
	if err != nil {
		return helper.NewFiberUnauthorizedErr(err.Error())
	}
	claims, ok := refreshToken.Claims.(jwt.MapClaims)
	if !ok || !refreshToken.Valid {
		return helper.NewFiberUnauthorizedErr("Please re-authenticate")
	}

	// Generate encoded tokens and send them as response.
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

func (e *Frontend) GetGroupInvitationsHandler(c *fiber.Ctx) error {
	userID := helper.GetUserIDFromContext(c)
	rspGroupInvitations, err := e.collaborationService.GetGroupInvitations(c.Context(), &pbCollaboration.UserIDRequest{
		UserID: userID,
	})
	if err != nil {
		return err
	}
	return c.JSON(rspGroupInvitations)
}

func (e *Frontend) ManageGroupInvitationHandler(c *fiber.Ctx) error {
	var data map[string]bool
	if err := c.BodyParser(&data); err != nil {
		return err
	}
	userID := helper.GetUserIDFromContext(c)
	rspManageInvitation, err := e.collaborationService.ManageGroupInvitation(
		c.Context(),
		&pbCollaboration.ManageGroupInvitationRequest{
			UserID:          userID,
			AdmissionID:     c.Params("requestID"),
			RequestResponse: data["isAccepted"],
		},
	)
	if err != nil {
		return err
	} else if !rspManageInvitation.Success {
		return helper.NewMicroNotSuccessfulResponseErr(helper.FrontendServiceID)
	}
	return c.SendStatus(200)
}

func (e *Frontend) GetUserGroupsHandler(c *fiber.Ctx) error {
	userID := helper.GetUserIDFromContext(c)
	rspUserGroups, err := e.collaborationService.GetUserGroups(c.Context(), &pbCollaboration.UserIDRequest{
		UserID: userID,
	})
	if err != nil {
		return err
	}
	return c.JSON(rspUserGroups)
}

func (e *Frontend) CreateGroupHandler(c *fiber.Ctx) error {
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return err
	}
	var groupName string
	if groupName = strings.TrimSpace(data["groupName"]); groupName == "" {
		return helper.NewFiberBadRequestErr("no group name given")
	}
	userID := helper.GetUserIDFromContext(c)
	rspCreateGroup, err := e.collaborationService.CreateNewGroupWithAdmin(
		c.Context(),
		&pbCollaboration.CreateGroupRequest{
			UserID:    userID,
			GroupName: groupName,
			IsDefault: false,
		},
	)
	if err != nil {
		return err
	}
	strSuccess := rspCreateGroup.ID
	return c.SendString(strSuccess)
}

func (e *Frontend) ModifyGroupHandler(c *fiber.Ctx) error {
	var data map[string]*string
	if err := c.BodyParser(&data); err != nil {
		return err
	}
	userID := helper.GetUserIDFromContext(c)
	var groupType pbCollaboration.GroupType
	if gt := strings.TrimSpace(*data["groupType"]); gt != "" {
		groupType = converter.MigrateStringGroupTypeToProtoGroupType(gt)
	}
	rspModifyGroup, err := e.collaborationService.ModifyGroup(c.Context(), &pbCollaboration.ModifyGroupRequest{
		UserID:    userID,
		GroupID:   c.Params("groupID"),
		GroupName: data["groupName"],
		GroupType: &groupType,
	})
	if err != nil {
		return err
	} else if !rspModifyGroup.Success {
		return helper.NewMicroNotSuccessfulResponseErr(helper.FrontendServiceID)
	}
	return c.SendStatus(200)
}

func (e *Frontend) DeleteGroupHandler(c *fiber.Ctx) error {
	userID := helper.GetUserIDFromContext(c)
	rspDeleteGroup, err := e.collaborationService.DeleteGroup(c.Context(), &pbCollaboration.GroupRequest{
		UserID:  userID,
		GroupID: c.Params("groupID"),
	})
	if err != nil {
		return err
	} else if !rspDeleteGroup.Success {
		return helper.NewMicroNotSuccessfulResponseErr(helper.FrontendServiceID)
	}
	return c.SendStatus(200)
}

func (e *Frontend) GetGroupMembersHandler(c *fiber.Ctx) error {
	userID := helper.GetUserIDFromContext(c)
	rspGroupMembers, err := e.collaborationService.GetGroupMembers(c.Context(), &pbCollaboration.GroupRequest{
		UserID:  userID,
		GroupID: c.Params("groupID"),
	})
	if err != nil {
		return err
	}
	return c.JSON(rspGroupMembers)
}

func (e *Frontend) GetGroupMemberRequestsHandler(c *fiber.Ctx) error {
	userID := helper.GetUserIDFromContext(c)
	rspMemberRequests, err := e.collaborationService.GetGroupMemberRequests(c.Context(), &pbCollaboration.GroupRequest{
		UserID:  userID,
		GroupID: c.Params("groupID"),
	})
	if err != nil {
		return err
	}
	return c.JSON(rspMemberRequests)
}

func (e *Frontend) ManageGroupMemberRequestHandler(c *fiber.Ctx) error {
	var data map[string]bool
	if err := c.BodyParser(&data); err != nil {
		return err
	}
	userID := helper.GetUserIDFromContext(c)
	rspManageRequest, err := e.collaborationService.ManageGroupMemberRequest(
		c.Context(),
		&pbCollaboration.ManageGroupMemberRequestRequest{
			UserID:          userID,
			GroupID:         c.Params("groupID"),
			AdmissionID:     c.Params("requestID"),
			RequestResponse: data["isAccepted"],
		},
	)
	if err != nil {
		return err
	} else if !rspManageRequest.Success {
		return helper.NewMicroNotSuccessfulResponseErr(helper.FrontendServiceID)
	}
	return c.SendStatus(200)
}

func (e *Frontend) RequestToJoinGroupHandler(c *fiber.Ctx) error {
	userID := helper.GetUserIDFromContext(c)
	rspJoinGroupRequest, err := e.collaborationService.RequestToJoinGroup(c.Context(), &pbCollaboration.GroupRequest{
		UserID:  userID,
		GroupID: c.Params("groupID"),
	})
	if err != nil {
		return err
	} else if !rspJoinGroupRequest.Success {
		return helper.NewMicroNotSuccessfulResponseErr(helper.FrontendServiceID)
	}
	return c.SendStatus(200)
}

func (e *Frontend) InviteUserToGroupHandler(c *fiber.Ctx) error {
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return err
	}
	var invitedUserEmail string
	if invitedUserEmail = strings.TrimSpace(data["invitedUserEmail"]); invitedUserEmail == "" {
		return helper.NewFiberBadRequestErr("no email for user to invite given")
	}
	userID := helper.GetUserIDFromContext(c)
	rspInviteUser, err := e.collaborationService.InviteUserToGroup(c.Context(), &pbCollaboration.GroupInvitationRequest{
		UserID:           userID,
		GroupID:          c.Params("groupID"),
		InvitedUserEmail: data["invitedUserEmail"],
	})
	if err != nil {
		return err
	} else if !rspInviteUser.Success {
		return helper.NewMicroNotSuccessfulResponseErr(helper.FrontendServiceID)
	}
	return c.SendStatus(200)
}

func (e *Frontend) GetGroupDecksHandler(c *fiber.Ctx) error {
	userID := helper.GetUserIDFromContext(c)
	rspGroupDecks, err := e.cardDeckService.GetGroupDecks(c.Context(), &pbCardDeck.GroupDecksRequest{
		UserID:  userID,
		GroupID: c.Params("groupID"),
	})
	if err != nil {
		return err
	}
	return c.JSON(rspGroupDecks)
}

func (e *Frontend) CreateDeckHandler(c *fiber.Ctx) error {
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return err
	}
	var deckName string
	if deckName = data["deckName"]; deckName == "" {
		return helper.NewFiberBadRequestErr("no deck name given")
	}
	userID := helper.GetUserIDFromContext(c)
	rspCreateDeck, err := e.cardDeckService.CreateDeck(c.Context(), &pbCardDeck.CreateDeckRequest{
		UserID:   userID,
		GroupID:  c.Params("groupID"),
		DeckName: data["deckName"],
	})
	if err != nil {
		return err
	}
	strSuccess := rspCreateDeck.ID
	return c.SendString(strSuccess)
}

func (e *Frontend) ModifyDeckHandler(c *fiber.Ctx) error {
	var data map[string]*string
	if err := c.BodyParser(&data); err != nil {
		return err
	}
	userID := helper.GetUserIDFromContext(c)
	rspModifyDeck, err := e.cardDeckService.ModifyDeck(c.Context(), &pbCardDeck.ModifyDeckRequest{
		UserID:   userID,
		DeckID:   c.Params("deckID"),
		DeckName: data["deckName"],
	})
	if err != nil {
		return err
	} else if !rspModifyDeck.Success {
		return helper.NewMicroNotSuccessfulResponseErr(helper.FrontendServiceID)
	}
	return c.SendStatus(200)
}

func (e *Frontend) DeleteDeckHandler(c *fiber.Ctx) error {
	userID := helper.GetUserIDFromContext(c)
	rspDeleteDeck, err := e.cardDeckService.DeleteDeck(c.Context(), &pbCardDeck.DeckRequest{
		UserID: userID,
		DeckID: c.Params("deckID"),
	})
	if err != nil {
		return err
	} else if !rspDeleteDeck.Success {
		return helper.NewMicroNotSuccessfulResponseErr(helper.FrontendServiceID)
	}
	return c.SendStatus(200)
}

func (e *Frontend) GetDeckCardsHandler(c *fiber.Ctx) error {
	userID := helper.GetUserIDFromContext(c)
	rspDeckCards, err := e.cardDeckService.GetDeckCards(c.Context(), &pbCardDeck.DeckRequest{
		UserID: userID,
		DeckID: c.Params("deckID"),
	})
	if err != nil {
		return err
	}
	return c.JSON(rspDeckCards)
}

func (e *Frontend) CreateCardHandler(c *fiber.Ctx) error {
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return err
	}
	if data["frontside"] == "" {
		return helper.NewFiberBadRequestErr("no card frontside given")
	}
	if data["backside"] == "" {
		return helper.NewFiberBadRequestErr("no card backside given")
	}
	userID := helper.GetUserIDFromContext(c)
	rspCreateCard, err := e.cardDeckService.CreateCard(c.Context(), &pbCardDeck.CreateCardRequest{
		UserID:    userID,
		DeckID:    c.Params("deckID"),
		Frontside: data["frontside"],
		Backside:  data["backside"],
	})
	if err != nil {
		return err
	}
	strSuccess := rspCreateCard.ID
	return c.SendString(strSuccess)
}

func (e *Frontend) ModifyCardHandler(c *fiber.Ctx) error {
	var data map[string]*string
	if err := c.BodyParser(&data); err != nil {
		return err
	}
	userID := helper.GetUserIDFromContext(c)
	rspModifyCard, err := e.cardDeckService.ModifyCard(c.Context(), &pbCardDeck.ModifyCardRequest{
		UserID:    userID,
		CardID:    c.Params("cardID"),
		Frontside: data["frontside"],
		Backside:  data["backside"],
	})
	if err != nil {
		return err
	} else if !rspModifyCard.Success {
		return helper.NewMicroNotSuccessfulResponseErr(helper.FrontendServiceID)
	}
	return c.SendStatus(200)
}

func (e *Frontend) DeleteCardHandler(c *fiber.Ctx) error {
	userID := helper.GetUserIDFromContext(c)
	rspDeleteCard, err := e.cardDeckService.DeleteCard(c.Context(), &pbCardDeck.DeleteCardRequest{
		UserID: userID,
		CardID: c.Params("cardID"),
	})
	if err != nil {
		return err
	} else if !rspDeleteCard.Success {
		return helper.NewMicroNotSuccessfulResponseErr(helper.FrontendServiceID)
	}
	return c.SendStatus(200)
}
