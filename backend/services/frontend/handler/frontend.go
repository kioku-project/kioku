package handler

import (
	"fmt"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/kioku-project/kioku/pkg/converter"

	"github.com/gofiber/fiber/v2"
	"github.com/kioku-project/kioku/pkg/helper"
	"github.com/kioku-project/kioku/pkg/model"
	pbCardDeck "github.com/kioku-project/kioku/services/carddeck/proto"
	pbCollaboration "github.com/kioku-project/kioku/services/collaboration/proto"
	pbSrs "github.com/kioku-project/kioku/services/srs/proto"
	pbUser "github.com/kioku-project/kioku/services/user/proto"
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
		return helper.NewFiberMissingEmailErr()
	}
	if data["userName"] == "" {
		return helper.NewFiberMissingNameErr()
	}
	if data["userPassword"] == "" {
		return helper.NewFiberMissingPasswordErr()
	}
	rspRegister, err := e.userService.Register(c.UserContext(), &pbUser.RegisterRequest{
		UserEmail:    data["userEmail"],
		UserName:     data["userName"],
		UserPassword: data["userPassword"],
	})
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
		return helper.NewFiberMissingEmailErr()
	}
	if reqUser.Password == "" {
		return helper.NewFiberMissingPasswordErr()
	}
	rspLogin, err := e.userService.Login(c.UserContext(), &pbUser.LoginRequest{
		UserEmail:    reqUser.Email,
		UserPassword: reqUser.Password,
	})
	if err != nil {
		return err
	}

	// Generate encoded tokens and send them as response.
	if err := helper.GenerateAccessToken(c, rspLogin.UserID, reqUser.Email, rspLogin.UserName); err != nil {
		return err
	}
	if err := helper.GenerateRefreshToken(c, rspLogin.UserID, reqUser.Email, rspLogin.UserName); err != nil {
		return err
	}

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

	rsp, err := e.userService.VerifyUserExists(c.UserContext(), &pbUser.VerificationRequest{
		UserID:    fmt.Sprint(claims["sub"]),
		UserEmail: fmt.Sprint(claims["email"]),
	})
	if err != nil {
		return err
	}
	if !rsp.Success {
		return helper.NewFiberUnauthorizedErr("User does not exist")
	}

	// Generate encoded tokens and send them as response.
	helper.GenerateAccessToken(c, claims["sub"].(string), claims["email"].(string), claims["name"].(string))
	helper.GenerateRefreshToken(c, claims["sub"].(string), claims["email"].(string), claims["name"].(string))
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
	rspGetUser, err := e.userService.GetUserProfileInformation(c.UserContext(), &pbUser.UserID{
		UserID: userID,
	})
	if err != nil {
		return err
	}
	return c.JSON(rspGetUser)
}

type modifyUserPayload struct {
	UserEmail    *string `json:"userEmail"`
	UserName     *string `json:"userName"`
	UserPassword *string `json:"userPassword"`
}

func (e *Frontend) ModifyUserHandler(c *fiber.Ctx) error {
	userID := helper.GetUserIDFromContext(c)
	var data modifyUserPayload
	if err := c.BodyParser(&data); err != nil {
		return err
	}
	rspModUser, err := e.userService.ModifyUserProfileInformation(c.Context(), &pbUser.ModifyRequest{
		UserID:       userID,
		UserName:     data.UserName,
		UserEmail:    data.UserEmail,
		UserPassword: data.UserPassword,
	})
	if err != nil {
		return err
	}
	if !rspModUser.Success {
		return helper.NewMicroNotSuccessfulResponseErr(helper.FrontendServiceID)
	}
	rspGetUser, err := e.userService.GetUserProfileInformation(c.Context(), &pbUser.UserID{
		UserID: userID,
	})
	if err != nil {
		return err
	}
	helper.GenerateAccessToken(c, userID, rspGetUser.UserEmail, rspGetUser.UserName)
	helper.GenerateRefreshToken(c, userID, rspGetUser.UserEmail, rspGetUser.UserName)
	return c.SendStatus(200)
}

func (e *Frontend) DeleteUserHandler(c *fiber.Ctx) error {
	userID := helper.GetUserIDFromContext(c)
	rspDeleteUser, err := e.userService.DeleteUser(c.UserContext(), &pbUser.UserID{
		UserID: userID,
	})
	if err != nil {
		return err
	}
	if !rspDeleteUser.Success {
		return helper.NewMicroNotSuccessfulResponseErr(helper.FrontendServiceID)
	}
	return c.SendStatus(200)
}

func (e *Frontend) GetGroupInvitationsHandler(c *fiber.Ctx) error {
	userID := helper.GetUserIDFromContext(c)
	rspGroupInvitations, err := e.collaborationService.GetGroupInvitations(c.UserContext(), &pbCollaboration.UserIDRequest{
		UserID: userID,
	})
	if err != nil {
		return err
	}
	return c.JSON(rspGroupInvitations)
}

func (e *Frontend) GetUserGroupsHandler(c *fiber.Ctx) error {
	userID := helper.GetUserIDFromContext(c)
	rspUserGroups, err := e.collaborationService.GetUserGroups(c.UserContext(), &pbCollaboration.UserIDRequest{
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
		c.UserContext(),
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

func (e *Frontend) LeaveGroupHandler(c *fiber.Ctx) error {
	groupID := c.Params("groupID")
	userID := helper.GetUserIDFromContext(c)
	rspLeaveGroup, err := e.collaborationService.LeaveGroupSafe(
		c.UserContext(),
		&pbCollaboration.GroupRequest{
			UserID:  userID,
			GroupID: groupID,
		},
	)
	if err != nil {
		return err
	}
	if !rspLeaveGroup.Success {
		return helper.NewMicroNotSuccessfulResponseErr(helper.FrontendServiceID)
	}
	return c.SendStatus(200)
}

func (e *Frontend) GetGroupHandler(c *fiber.Ctx) error {
	userID := helper.GetUserIDFromContext(c)
	rspGetGroup, err := e.collaborationService.GetGroup(c.UserContext(), &pbCollaboration.GroupRequest{
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
	rspModifyGroup, err := e.collaborationService.ModifyGroup(c.UserContext(), &pbCollaboration.ModifyGroupRequest{
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
	rspDeleteGroup, err := e.collaborationService.DeleteGroup(c.UserContext(), &pbCollaboration.GroupRequest{
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
	rspGroupMembers, err := e.collaborationService.GetGroupMembers(c.UserContext(), &pbCollaboration.GroupRequest{
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

func (e *Frontend) ModifyGroupMemberHandler(c *fiber.Ctx) error {
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return err
	}
	userID := helper.GetUserIDFromContext(c)
	var groupRole pbCollaboration.GroupRole
	if dataGroupRole, ok := data["groupRole"]; ok {
		if gr := strings.TrimSpace(dataGroupRole); gr != "" {
			groupRole = converter.MigrateStringRoleToProtoRole(gr)
		}
	}
	rspModifyGroupUser, err := e.collaborationService.ModifyGroupUserRequest(
		c.UserContext(),
		&pbCollaboration.GroupModUserRequest{
			UserID:    userID,
			GroupID:   c.Params("groupID"),
			ModUserID: c.Params("userID"),
			NewRole:   groupRole,
		},
	)
	if err != nil {
		return err
	}
	if !rspModifyGroupUser.Success {
		return helper.NewMicroNotSuccessfulResponseErr(helper.FrontendServiceID)
	}
	return c.SendStatus(200)
}

func (e *Frontend) KickGroupMemberHandler(c *fiber.Ctx) error {
	userID := helper.GetUserIDFromContext(c)
	rspKickGroupUser, err := e.collaborationService.KickGroupUser(c.UserContext(), &pbCollaboration.GroupKickUserRequest{
		UserID:    userID,
		GroupID:   c.Params("groupID"),
		DelUserID: c.Params("userID"),
	})
	if err != nil {
		return err
	}
	if !rspKickGroupUser.Success {
		return helper.NewMicroNotSuccessfulResponseErr(helper.FrontendServiceID)
	}
	return c.SendStatus(200)
}

func (e *Frontend) GetGroupMemberRequestsHandler(c *fiber.Ctx) error {
	userID := helper.GetUserIDFromContext(c)
	rspMemberRequests, err := e.collaborationService.GetGroupMemberRequests(c.UserContext(), &pbCollaboration.GroupRequest{
		UserID:  userID,
		GroupID: c.Params("groupID"),
	})
	if err != nil {
		return err
	}
	return c.JSON(converter.FiberGetGroupMemberRequestsResponseBody{
		MemberRequests: converter.ConvertToTypeArray(
			rspMemberRequests.MemberAdmissions,
			converter.ProtoGroupMemberRequestToFiberGroupMemberRequestConverter,
		),
	})
}

func (e *Frontend) AddUserGroupRequestHandler(c *fiber.Ctx) error {
	userID := helper.GetUserIDFromContext(c)
	rspJoinGroupRequest, err := e.collaborationService.AddGroupUserRequest(
		c.UserContext(),
		&pbCollaboration.GroupUserRequest{
			UserID:  userID,
			GroupID: c.Params("groupID"),
		},
	)
	if err != nil {
		return err
	} else if !rspJoinGroupRequest.Success {
		return helper.NewMicroNotSuccessfulResponseErr(helper.FrontendServiceID)
	}
	return c.SendStatus(200)
}

func (e *Frontend) RemoveUserGroupRequestHandler(c *fiber.Ctx) error {
	userID := helper.GetUserIDFromContext(c)
	rspJoinGroupRequest, err := e.collaborationService.RemoveGroupUserRequest(
		c.UserContext(),
		&pbCollaboration.GroupUserRequest{
			UserID:  userID,
			GroupID: c.Params("groupID"),
		},
	)
	if err != nil {
		return err
	} else if !rspJoinGroupRequest.Success {
		return helper.NewMicroNotSuccessfulResponseErr(helper.FrontendServiceID)
	}
	return c.SendStatus(200)
}

func (e *Frontend) GetInvitationsForGroupHandler(c *fiber.Ctx) error {
	userID := helper.GetUserIDFromContext(c)
	rspInvitationsForGroup, err := e.collaborationService.GetInvitationsForGroup(
		c.UserContext(),
		&pbCollaboration.GroupRequest{
			UserID:  userID,
			GroupID: c.Params("groupID"),
		},
	)
	if err != nil {
		return err
	}
	return c.JSON(converter.FiberGetInvitationsForGroupResponseBody{
		MemberRequests: converter.ConvertToTypeArray(
			rspInvitationsForGroup.MemberAdmissions,
			converter.ProtoGroupMemberRequestToFiberGroupMemberRequestConverter,
		),
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
	rspInviteUser, err := e.collaborationService.AddGroupUserInvite(c.UserContext(), &pbCollaboration.GroupUserInvite{
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
	rspInviteUser, err := e.collaborationService.RemoveGroupUserInvite(c.UserContext(), &pbCollaboration.GroupUserInvite{
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
	rspGroupDecks, err := e.cardDeckService.GetGroupDecks(c.UserContext(), &pbCardDeck.GroupDecksRequest{
		UserID:  userID,
		GroupID: c.Params("groupID"),
	})
	if err != nil {
		return err
	}
	decks := converter.ConvertToTypeArray(rspGroupDecks.Decks, converter.ProtoDeckToFiberDeckConverter)
	for _, deck := range decks {
		deck.GroupID = c.Params("groupID")
	}
	return c.JSON(converter.FiberGetGroupDecksResponseBody{
		Decks: decks,
	})
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
	deckType := pbCardDeck.DeckType_PRIVATE
	if dt := strings.TrimSpace(data["deckType"]); dt != "" {
		deckType = converter.MigrateStringDeckTypeToProtoDeckType(dt)
	}
	rspCreateDeck, err := e.cardDeckService.CreateDeck(c.UserContext(), &pbCardDeck.CreateDeckRequest{
		UserID:   userID,
		GroupID:  c.Params("groupID"),
		DeckName: data["deckName"],
		DeckType: deckType,
	})
	if err != nil {
		return err
	}
	return c.SendString(rspCreateDeck.ID)
}

func (e *Frontend) GetDeckHandler(c *fiber.Ctx) error {
	userID := helper.GetUserIDFromContext(c)
	rspGetDeck, err := e.cardDeckService.GetDeck(c.UserContext(), &pbCardDeck.IDRequest{
		UserID:   userID,
		EntityID: c.Params("deckID"),
	})
	if err != nil {
		return err
	}
	return c.JSON(converter.ProtoDeckRespToFiberDeckConverter(rspGetDeck))
}

func (e *Frontend) ModifyDeckHandler(c *fiber.Ctx) error {
	var data map[string]*string
	if err := c.BodyParser(&data); err != nil {
		return err
	}
	var deckType pbCardDeck.DeckType
	if data["deckType"] != nil {
		if dt := strings.TrimSpace(*data["deckType"]); dt != "" {
			deckType = converter.MigrateStringDeckTypeToProtoDeckType(dt)
		}
	}
	userID := helper.GetUserIDFromContext(c)
	rspModifyDeck, err := e.cardDeckService.ModifyDeck(c.UserContext(), &pbCardDeck.ModifyDeckRequest{
		UserID:   userID,
		DeckID:   c.Params("deckID"),
		DeckName: data["deckName"],
		DeckType: &deckType,
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
	rspDeleteDeck, err := e.cardDeckService.DeleteDeck(c.UserContext(), &pbCardDeck.IDRequest{
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
	rspDeckCards, err := e.cardDeckService.GetDeckCards(c.UserContext(), &pbCardDeck.IDRequest{
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
	rspCreateCard, err := e.cardDeckService.CreateCard(c.UserContext(), &pbCardDeck.CreateCardRequest{
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
	rspGetCard, err := e.cardDeckService.GetCard(c.UserContext(), &pbCardDeck.IDRequest{
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
	rspModifyCard, err := e.cardDeckService.ModifyCard(c.UserContext(), &pbCardDeck.ModifyCardRequest{
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
	rspDeleteCard, err := e.cardDeckService.DeleteCard(c.UserContext(), &pbCardDeck.IDRequest{
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
	rspCreateCardSide, err := e.cardDeckService.CreateCardSide(c.UserContext(), &pbCardDeck.CreateCardSideRequest{
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
	rspModifyCardSide, err := e.cardDeckService.ModifyCardSide(c.UserContext(), &pbCardDeck.ModifyCardSideRequest{
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
	rspDeleteCardSide, err := e.cardDeckService.DeleteCardSide(c.UserContext(), &pbCardDeck.IDRequest{
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
	rspSrsPull, err := e.srsService.Pull(c.UserContext(), &pbSrs.DeckPullRequest{
		UserID: userID,
		DeckID: c.Params("deckID"),
	})
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
	rspSrsPush, err := e.srsService.Push(c.UserContext(), &pbSrs.SrsPushRequest{
		UserID: userID,
		CardID: data.CardID,
		DeckID: c.Params("deckID"),
		Rating: data.Rating,
	})
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
	rspSrsDue, err := e.srsService.GetDeckCardsDue(c.UserContext(), &pbSrs.DeckPullRequest{
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
	dueCards, err := e.srsService.GetUserCardsDue(c.UserContext(), &pbSrs.UserDueRequest{
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
