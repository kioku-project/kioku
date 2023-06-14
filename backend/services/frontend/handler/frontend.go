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
	pbSrs "github.com/kioku-project/kioku/services/srs/proto"
	pbUser "github.com/kioku-project/kioku/services/user/proto"
	"go-micro.dev/v4/logger"
)

type Frontend struct {
	userService          pbUser.UserService
	cardDeckService      pbCardDeck.CardDeckService
	collaborationService pbCollaboration.CollaborationService
	srsService           pbSrs.SrsService
}

func New(
	userService pbUser.UserService,
	cardDeckService pbCardDeck.CardDeckService,
	collaborationService pbCollaboration.CollaborationService,
	srsService pbSrs.SrsService,
) *Frontend {
	return &Frontend{
		userService:          userService,
		cardDeckService:      cardDeckService,
		collaborationService: collaborationService,
		srsService:           srsService,
	}
}

func (e *Frontend) RegisterHandler(c *fiber.Ctx) error {
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return err
	}
	if data["userEmail"] == "" {
		return helper.NewFiberBadRequestErr("no e-mail given")
	}
	if data["userName"] == "" {
		return helper.NewFiberBadRequestErr("no name given")
	}
	if data["userPassword"] == "" {
		return helper.NewFiberBadRequestErr("no password given")
	}
	rspRegister, err := e.userService.Register(c.Context(), &pbUser.RegisterRequest{
		UserEmail:    data["userEmail"],
		UserName:     data["userName"],
		UserPassword: data["userPassword"]})
	if err != nil {
		return err
	}
	return c.SendString(rspRegister.UserName)
}

func (e *Frontend) LoginHandler(c *fiber.Ctx) error {
	var reqUser model.User
	if err := c.BodyParser(&reqUser); err != nil {
		return err
	}
	if reqUser.Email == "" {
		return helper.NewFiberBadRequestErr("no e-mail given")
	}
	if reqUser.Password == "" {
		return helper.NewFiberBadRequestErr("no password given")
	}
	rspLogin, err := e.userService.Login(c.Context(), &pbUser.LoginRequest{
		UserEmail:    reqUser.Email,
		UserPassword: reqUser.Password})
	if err != nil {
		return err
	}

	// Generate encoded tokens and send them as response.
	aTExp := time.Now().Add(time.Minute * 30)
	aTString, err := helper.CreateJWTTokenString(aTExp, rspLogin.UserID, reqUser.Email, rspLogin.UserName)
	if err != nil {
		logger.Infof("%v", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	rTExp := time.Now().Add(time.Hour * 24 * 7)
	rTString, err := helper.CreateJWTTokenString(rTExp, rspLogin.UserID, reqUser.Email, rspLogin.UserName)
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

func (e *Frontend) LogoutHandler(c *fiber.Ctx) error {
	c.Cookie(&fiber.Cookie{
		Name:    "access_token",
		Path:    "/",
		Expires: time.Now().Add(-time.Minute * 30),
	})
	c.Cookie(&fiber.Cookie{
		Name:     "refresh_token",
		Path:     "/",
		Expires:  time.Now().Add(-time.Minute * 30),
		HTTPOnly: true,
	})
	return c.SendStatus(200)
}

func (e *Frontend) GetUserHandler(c *fiber.Ctx) error {
	userID := helper.GetUserIDFromContext(c)
	rspGetUser, err := e.userService.GetUserProfileInformation(c.Context(), &pbUser.UserID{
		UserID: userID,
	})
	if err != nil {
		return err
	}
	return c.JSON(rspGetUser)
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

func (e *Frontend) GetUserGroupsHandler(c *fiber.Ctx) error {
	userID := helper.GetUserIDFromContext(c)
	rspUserGroups, err := e.collaborationService.GetUserGroups(c.Context(), &pbCollaboration.UserIDRequest{
		UserID: userID,
	})
	if err != nil {
		return err
	}
	return c.JSON(converter.FiberGetUserGroupsResponseBody{
		Groups: converter.ConvertToTypeArray(rspUserGroups.Groups, converter.ProtoGroupWithRoleToFiberGroupConverter),
	})
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
			UserID:           userID,
			GroupName:        groupName,
			GroupDescription: data["groupDescription"],
			IsDefault:        false,
		},
	)
	if err != nil {
		return err
	}
	return c.SendString(rspCreateGroup.ID)
}

func (e *Frontend) GetGroupHandler(c *fiber.Ctx) error {
	userID := helper.GetUserIDFromContext(c)
	rspGetGroup, err := e.collaborationService.GetGroup(c.Context(), &pbCollaboration.GroupRequest{
		UserID:  userID,
		GroupID: c.Params("groupID"),
	})
	if err != nil {
		return err
	}
	return c.JSON(converter.ProtoGroupWithRoleToFiberGroupConverter(rspGetGroup))
}

func (e *Frontend) ModifyGroupHandler(c *fiber.Ctx) error {
	var data map[string]*string
	if err := c.BodyParser(&data); err != nil {
		return err
	}
	userID := helper.GetUserIDFromContext(c)
	var groupType pbCollaboration.GroupType
	if data["groupType"] != nil {
		if gt := strings.TrimSpace(*data["groupType"]); gt != "" {
			groupType = converter.MigrateStringGroupTypeToProtoGroupType(gt)
		}
	}
	rspModifyGroup, err := e.collaborationService.ModifyGroup(c.Context(), &pbCollaboration.ModifyGroupRequest{
		UserID:           userID,
		GroupID:          c.Params("groupID"),
		GroupName:        data["groupName"],
		GroupDescription: data["groupDescription"],
		GroupType:        &groupType,
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
	return c.JSON(converter.FiberGetGroupMembersResponseBody{
		Members: converter.ConvertToTypeArray(rspGroupMembers.Users, converter.ProtoUserWithRoleToFiberGroupMember),
	})
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
	return c.JSON(converter.FiberGetGroupMemberRequestsResponseBody{
		MemberRequests: converter.ConvertToTypeArray(rspMemberRequests.MemberAdmissions, converter.ProtoGroupMemberRequestToFiberGroupMemberRequestConverter),
	})
}

func (e *Frontend) AddUserGroupRequestHandler(c *fiber.Ctx) error {
	userID := helper.GetUserIDFromContext(c)
	rspJoinGroupRequest, err := e.collaborationService.AddGroupUserRequest(c.Context(), &pbCollaboration.GroupUserRequest{
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

func (e *Frontend) RemoveUserGroupRequestHandler(c *fiber.Ctx) error {
	userID := helper.GetUserIDFromContext(c)
	rspJoinGroupRequest, err := e.collaborationService.RemoveGroupUserRequest(c.Context(), &pbCollaboration.GroupUserRequest{
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

func (e *Frontend) GetInvitationsForGroupHandler(c *fiber.Ctx) error {
	userID := helper.GetUserIDFromContext(c)
	rspInvitationsForGroup, err := e.collaborationService.GetInvitationsForGroup(c.Context(), &pbCollaboration.GroupRequest{
		UserID:  userID,
		GroupID: c.Params("groupID"),
	})
	if err != nil {
		return err
	}
	return c.JSON(converter.FiberGetInvitationsForGroupResponseBody{
		MemberRequests: converter.ConvertToTypeArray(rspInvitationsForGroup.MemberAdmissions, converter.ProtoGroupMemberRequestToFiberGroupMemberRequestConverter),
	})
}

func (e *Frontend) AddUserGroupInviteHandler(c *fiber.Ctx) error {
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return err
	}
	var invitedUserEmail string
	if invitedUserEmail = strings.TrimSpace(data["invitedUserEmail"]); invitedUserEmail == "" {
		return helper.NewFiberBadRequestErr("no email for user to invite given")
	}
	userID := helper.GetUserIDFromContext(c)
	rspInviteUser, err := e.collaborationService.AddGroupUserInvite(c.Context(), &pbCollaboration.GroupUserInvite{
		UserID:          userID,
		GroupID:         c.Params("groupID"),
		InviteUserEmail: data["invitedUserEmail"],
	})
	if err != nil {
		return err
	} else if !rspInviteUser.Success {
		return helper.NewMicroNotSuccessfulResponseErr(helper.FrontendServiceID)
	}
	return c.SendStatus(200)
}

func (e *Frontend) RemoveUserGroupInviteHandler(c *fiber.Ctx) error {
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return err
	}
	var invitedUserEmail string
	if invitedUserEmail = strings.TrimSpace(data["invitedUserEmail"]); invitedUserEmail == "" {
		return helper.NewFiberBadRequestErr("no email for user to invite given")
	}
	userID := helper.GetUserIDFromContext(c)
	rspInviteUser, err := e.collaborationService.RemoveGroupUserInvite(c.Context(), &pbCollaboration.GroupUserInvite{
		UserID:          userID,
		GroupID:         c.Params("groupID"),
		InviteUserEmail: data["invitedUserEmail"],
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
	return c.SendString(rspCreateDeck.ID)
}

func (e *Frontend) GetDeckHandler(c *fiber.Ctx) error {
	userID := helper.GetUserIDFromContext(c)
	rspGetDeck, err := e.cardDeckService.GetDeck(c.Context(), &pbCardDeck.IDRequest{
		UserID:   userID,
		EntityID: c.Params("deckID"),
	})
	if err != nil {
		return err
	}
	return c.JSON(rspGetDeck)
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
	rspDeleteDeck, err := e.cardDeckService.DeleteDeck(c.Context(), &pbCardDeck.IDRequest{
		UserID:   userID,
		EntityID: c.Params("deckID"),
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
	rspDeckCards, err := e.cardDeckService.GetDeckCards(c.Context(), &pbCardDeck.IDRequest{
		UserID:   userID,
		EntityID: c.Params("deckID"),
	})
	if err != nil {
		return err
	}
	return c.JSON(rspDeckCards)
}

func (e *Frontend) CreateCardHandler(c *fiber.Ctx) error {
	var data converter.FiberCreateCardRequestBody
	if err := c.BodyParser(&data); err != nil {
		return err
	}
	if len(data.Sides) == 0 {
		return helper.NewFiberBadRequestErr("new card should at least have one side")
	}
	userID := helper.GetUserIDFromContext(c)
	rspCreateCard, err := e.cardDeckService.CreateCard(c.Context(), &pbCardDeck.CreateCardRequest{
		UserID: userID,
		DeckID: c.Params("deckID"),
		Sides:  converter.ConvertToTypeArray(data.Sides, converter.FiberCardSideContentToProtoCardSideContent),
	})
	if err != nil {
		return err
	}
	return c.SendString(rspCreateCard.ID)
}

func (e *Frontend) GetCardHandler(c *fiber.Ctx) error {
	userID := helper.GetUserIDFromContext(c)
	rspGetCard, err := e.cardDeckService.GetCard(c.Context(), &pbCardDeck.IDRequest{
		UserID:   userID,
		EntityID: c.Params("cardID"),
	})
	if err != nil {
		return err
	}
	return c.JSON(rspGetCard)
}

func (e *Frontend) ModifyCardHandler(c *fiber.Ctx) error {
	var data converter.FiberModifyCardRequestBody
	if err := c.BodyParser(&data); err != nil {
		return err
	}
	if len(data.Sides) == 0 {
		return helper.NewFiberBadRequestErr("modified card should at least have one side")
	}
	userID := helper.GetUserIDFromContext(c)
	rspModifyCard, err := e.cardDeckService.ModifyCard(c.Context(), &pbCardDeck.ModifyCardRequest{
		UserID: userID,
		CardID: c.Params("cardID"),
		Sides:  converter.ConvertToTypeArray(data.Sides, converter.FiberCardSideContentToProtoCardSideContent),
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
	rspDeleteCard, err := e.cardDeckService.DeleteCard(c.Context(), &pbCardDeck.IDRequest{
		UserID:   userID,
		EntityID: c.Params("cardID"),
	})
	if err != nil {
		return err
	} else if !rspDeleteCard.Success {
		return helper.NewMicroNotSuccessfulResponseErr(helper.FrontendServiceID)
	}
	return c.SendStatus(200)
}

func (e *Frontend) CreateCardSideHandler(c *fiber.Ctx) error {
	var data converter.FiberCreateCardSideRequestBody
	if err := c.BodyParser(&data); err != nil {
		return err
	}
	userID := helper.GetUserIDFromContext(c)
	rspCreateCardSide, err := e.cardDeckService.CreateCardSide(c.Context(), &pbCardDeck.CreateCardSideRequest{
		UserID:                userID,
		CardID:                c.Params("cardID"),
		PlaceBeforeCardSideID: data.PlaceBeforeCardSideID,
		Content:               converter.FiberCardSideContentToProtoCardSideContent(data.FiberCardSideContent),
	})
	if err != nil {
		return err
	}
	return c.SendString(rspCreateCardSide.ID)
}

func (e *Frontend) ModifyCardSideHandler(c *fiber.Ctx) error {
	var data converter.FiberModifyCardSideRequestBody
	if err := c.BodyParser(&data); err != nil {
		return err
	}
	userID := helper.GetUserIDFromContext(c)
	rspModifyCardSide, err := e.cardDeckService.ModifyCardSide(c.Context(), &pbCardDeck.ModifyCardSideRequest{
		UserID:     userID,
		CardSideID: c.Params("cardSideID"),
		Content:    converter.FiberCardSideContentToProtoCardSideContent(data.FiberCardSideContent),
	})
	if err != nil {
		return err
	} else if !rspModifyCardSide.Success {
		return helper.NewMicroNotSuccessfulResponseErr(helper.FrontendServiceID)
	}
	return c.SendStatus(200)
}

func (e *Frontend) DeleteCardSideHandler(c *fiber.Ctx) error {
	userID := helper.GetUserIDFromContext(c)
	rspDeleteCardSide, err := e.cardDeckService.DeleteCardSide(c.Context(), &pbCardDeck.IDRequest{
		UserID:   userID,
		EntityID: c.Params("cardSideID"),
	})
	if err != nil {
		return err
	} else if !rspDeleteCardSide.Success {
		return helper.NewMicroNotSuccessfulResponseErr(helper.FrontendServiceID)
	}
	return c.SendStatus(200)
}

func (e *Frontend) SrsPullHandler(c *fiber.Ctx) error {
	userID := helper.GetUserIDFromContext(c)
	rspSrsPull, err := e.srsService.Pull(c.Context(), &pbSrs.DeckPullRequest{UserID: userID, DeckID: c.Params("deckID")})
	if err != nil {
		return err
	}
	return c.JSON(rspSrsPull.Card)
}
func (e *Frontend) SrsPushHandler(c *fiber.Ctx) error {
	var data converter.FiberPushCardRequestBody
	if err := c.BodyParser(&data); err != nil {
		return err
	}
	if data.CardID == "" {
		return helper.NewFiberBadRequestErr("no cardID given")
	}
	userID := helper.GetUserIDFromContext(c)
	rspSrsPush, err := e.srsService.Push(c.Context(), &pbSrs.SrsPushRequest{UserID: userID, CardID: data.CardID, DeckID: c.Params("deckID"), Rating: data.Rating})
	if err != nil {
		return err
	}
	if err != nil || !rspSrsPush.Success {
		return err
	}
	return c.SendStatus(200)
}

func (e *Frontend) SrsDeckDueHandler(c *fiber.Ctx) error {

	userID := helper.GetUserIDFromContext(c)
	rspSrsDue, err := e.srsService.GetDeckCardsDue(c.Context(), &pbSrs.DeckPullRequest{
		UserID: userID,
		DeckID: c.Params("deckID"),
	})
	if err != nil {
		return err
	}
	if rspSrsDue.Due == 0 {
		return c.JSON(0)
	}
	return c.JSON(rspSrsDue.Due)
}

func (e *Frontend) SrsUserDueHandler(c *fiber.Ctx) error {

	userID := helper.GetUserIDFromContext(c)
	dueCards, err := e.srsService.GetUserCardsDue(c.Context(), &pbSrs.UserDueRequest{
		UserID: userID,
	})
	if err != nil {
		return err
	}
	return c.JSON(converter.FiberGetUserDueCardsResponseBody{
		DueCards: dueCards.DueCards,
		DueDecks: dueCards.DueDecks,
	})
}
