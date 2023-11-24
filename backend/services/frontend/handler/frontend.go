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
	pbCommon "github.com/kioku-project/kioku/pkg/proto"
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
	var data model.User
	if err := c.BodyParser(&data); err != nil {
		return err
	}
	if data.Email == "" {
		return helper.NewFiberMissingEmailErr()
	}
	if data.Name == "" {
		return helper.NewFiberMissingNameErr()
	}
	if data.Password == "" {
		return helper.NewFiberMissingPasswordErr()
	}
	_, err := e.userService.Register(c.Context(), &pbCommon.User{
		UserEmail:    data.Email,
		UserName:     data.Name,
		UserPassword: data.Password,
	})
	// empty password in order to omit field on return
	data.Password = ""
	if err != nil {
		return err
	}
	return c.JSON(data)
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
	rspLogin, err := e.userService.Login(c.Context(), &pbCommon.User{
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

	rsp, err := e.userService.VerifyUserExists(c.Context(), &pbCommon.User{
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
	rspGetUser, err := e.userService.GetUserProfileInformation(c.Context(), &pbCommon.User{
		UserID: userID,
	})
	if err != nil {
		return err
	}
	return c.JSON(rspGetUser)
}

func (e *Frontend) ModifyUserHandler(c *fiber.Ctx) error {
	userID := helper.GetUserIDFromContext(c)
	var data model.User
	if err := c.BodyParser(&data); err != nil {
		return err
	}
	rspModUser, err := e.userService.ModifyUserProfileInformation(c.Context(), &pbCommon.User{
		UserID:       userID,
		UserName:     data.Name,
		UserEmail:    data.Email,
		UserPassword: data.Password,
	})
	if err != nil {
		return err
	}
	if !rspModUser.Success {
		return helper.NewMicroNotSuccessfulResponseErr(helper.FrontendServiceID)
	}
	rspGetUser, err := e.userService.GetUserProfileInformation(c.Context(), &pbCommon.User{
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
	rspDeleteUser, err := e.userService.DeleteUser(c.Context(), &pbCommon.User{
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
	rspGroupInvitations, err := e.collaborationService.GetGroupInvitations(c.Context(), &pbCommon.User{
		UserID: userID,
	})
	if err != nil {
		return err
	}
	return c.JSON(rspGroupInvitations)
}

func (e *Frontend) GetUserGroupsHandler(c *fiber.Ctx) error {
	userID := helper.GetUserIDFromContext(c)
	rspUserGroups, err := e.collaborationService.GetUserGroups(c.Context(), &pbCommon.User{
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
	var groupType pbCommon.GroupType
	if data["groupType"] != "" {
		groupType = converter.MigrateStringGroupTypeToProtoGroupType(data["groupType"])
	}
	userID := helper.GetUserIDFromContext(c)
	rspCreateGroup, err := e.collaborationService.CreateNewGroupWithAdmin(
		c.Context(),
		&pbCommon.GroupRequest{
			UserID: userID,
			Group: &pbCommon.Group{
				GroupName:        groupName,
				GroupDescription: data["groupDescription"],
				GroupType:        groupType,
				IsDefault:        false,
			},
		},
	)
	if err != nil {
		return err
	}
	return c.SendString(rspCreateGroup.GroupID)
}

func (e *Frontend) LeaveGroupHandler(c *fiber.Ctx) error {
	groupID := c.Params("groupID")
	userID := helper.GetUserIDFromContext(c)
	rspLeaveGroup, err := e.collaborationService.LeaveGroupSafe(
		c.Context(),
		&pbCommon.GroupRequest{
			UserID: userID,
			Group: &pbCommon.Group{
				GroupID: groupID,
			},
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
	rspGetGroup, err := e.collaborationService.GetGroup(c.Context(), &pbCommon.GroupRequest{
		UserID: userID,
		Group: &pbCommon.Group{
			GroupID: c.Params("groupID"),
		},
	})
	if err != nil {
		return err
	}
	return c.JSON(converter.ProtoGroupWithRoleToFiberGroupConverter(rspGetGroup))
}

func (e *Frontend) ModifyGroupHandler(c *fiber.Ctx) error {
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return err
	}
	userID := helper.GetUserIDFromContext(c)
	var groupType pbCommon.GroupType
	if data["groupType"] != "" {
		groupType = converter.MigrateStringGroupTypeToProtoGroupType(data["groupType"])
	}
	rspModifyGroup, err := e.collaborationService.ModifyGroup(c.Context(), &pbCommon.GroupRequest{
		UserID: userID,
		Group: &pbCommon.Group{
			GroupID:          c.Params("groupID"),
			GroupName:        data["groupName"],
			GroupDescription: data["groupDescription"],
			GroupType:        groupType,
		},
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
	rspDeleteGroup, err := e.collaborationService.DeleteGroup(c.Context(), &pbCommon.GroupRequest{
		UserID: userID,
		Group: &pbCommon.Group{
			GroupID: c.Params("groupID"),
		},
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
	rspGroupMembers, err := e.collaborationService.GetGroupMembers(c.Context(), &pbCommon.GroupRequest{
		UserID: userID,
		Group: &pbCommon.Group{
			GroupID: c.Params("groupID"),
		},
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
	var groupRole pbCommon.GroupRole
	if dataGroupRole, ok := data["groupRole"]; ok {
		if gr := strings.TrimSpace(dataGroupRole); gr != "" {
			groupRole = converter.MigrateStringRoleToProtoRole(gr)
		}
	}
	rspModifyGroupUser, err := e.collaborationService.ModifyGroupUserRequest(
		c.Context(),
		&pbCommon.GroupModUserRequest{
			UserID: userID,
			Group: &pbCommon.Group{
				GroupID: c.Params("groupID"),
				Role:    groupRole,
			},
			ModUserID: c.Params("userID"),
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
	rspKickGroupUser, err := e.collaborationService.KickGroupUser(c.Context(), &pbCommon.GroupModUserRequest{
		UserID: userID,
		Group: &pbCommon.Group{
			GroupID: c.Params("groupID"),
		},
		ModUserID: c.Params("userID"),
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
	rspMemberRequests, err := e.collaborationService.GetGroupMemberRequests(c.Context(), &pbCommon.GroupRequest{
		UserID: userID,
		Group: &pbCommon.Group{
			GroupID: c.Params("groupID"),
		},
	})
	if err != nil {
		return err
	}
	return c.JSON(converter.FiberGetGroupMemberRequestsResponseBody{
		MemberRequests: converter.ConvertToTypeArray(
			rspMemberRequests.Users,
			converter.ProtoGroupMemberRequestToFiberGroupMemberRequestConverter,
		),
	})
}

func (e *Frontend) AddUserGroupRequestHandler(c *fiber.Ctx) error {
	userID := helper.GetUserIDFromContext(c)
	rspJoinGroupRequest, err := e.collaborationService.AddGroupUserRequest(
		c.Context(),
		&pbCommon.GroupRequest{
			UserID: userID,
			Group: &pbCommon.Group{
				GroupID: c.Params("groupID"),
			},
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
		c.Context(),
		&pbCommon.GroupRequest{
			UserID: userID,
			Group: &pbCommon.Group{
				GroupID: c.Params("groupID"),
			},
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
		c.Context(),
		&pbCommon.GroupRequest{
			UserID: userID,
			Group: &pbCommon.Group{
				GroupID: c.Params("groupID"),
			},
		},
	)
	if err != nil {
		return err
	}
	return c.JSON(converter.FiberGetInvitationsForGroupResponseBody{
		MemberRequests: converter.ConvertToTypeArray(
			rspInvitationsForGroup.Users,
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
	rspInviteUser, err := e.collaborationService.AddGroupUserInvite(c.Context(), &pbCommon.GroupInviteRequest{
		UserID: userID,
		Group: &pbCommon.Group{
			GroupID: c.Params("groupID"),
		},
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
	rspInviteUser, err := e.collaborationService.RemoveGroupUserInvite(c.Context(), &pbCommon.GroupInviteRequest{
		UserID: userID,
		Group: &pbCommon.Group{
			GroupID: c.Params("groupID"),
		},
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
	rspGroupDecks, err := e.cardDeckService.GetGroupDecks(c.Context(), &pbCommon.GroupRequest{
		UserID: userID,
		Group: &pbCommon.Group{
			GroupID: c.Params("groupID"),
		},
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
	deckType := pbCommon.DeckType_PRIVATE
	if dt := strings.TrimSpace(data["deckType"]); dt != "" {
		deckType = converter.MigrateStringDeckTypeToProtoDeckType(dt)
	}
	rspCreateDeck, err := e.cardDeckService.CreateDeck(c.Context(), &pbCommon.DeckRequest{
		UserID: userID,
		Deck: &pbCommon.Deck{
			GroupID:  c.Params("groupID"),
			DeckName: data["deckName"],
			DeckType: deckType,
		},
	})
	if err != nil {
		return err
	}
	return c.SendString(rspCreateDeck.DeckID)
}

func (e *Frontend) GetDeckHandler(c *fiber.Ctx) error {
	userID := helper.GetUserIDFromContext(c)
	rspGetDeck, err := e.cardDeckService.GetDeck(c.Context(), &pbCommon.DeckRequest{
		UserID: userID,
		Deck: &pbCommon.Deck{
			DeckID: c.Params("deckID"),
		},
	})
	if err != nil {
		return err
	}
	return c.JSON(converter.ProtoDeckRespToFiberDeckConverter(rspGetDeck))
}

func (e *Frontend) ModifyDeckHandler(c *fiber.Ctx) error {
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return err
	}
	var deckType pbCommon.DeckType
	if data["deckType"] != "" {
		deckType = converter.MigrateStringDeckTypeToProtoDeckType(data["deckType"])
	}
	userID := helper.GetUserIDFromContext(c)
	rspModifyDeck, err := e.cardDeckService.ModifyDeck(c.Context(), &pbCommon.DeckRequest{
		UserID: userID,
		Deck: &pbCommon.Deck{
			DeckID:   c.Params("deckID"),
			DeckName: data["deckName"],
			DeckType: deckType,
		},
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
	rspDeleteDeck, err := e.cardDeckService.DeleteDeck(c.Context(), &pbCommon.DeckRequest{
		UserID: userID,
		Deck: &pbCommon.Deck{
			DeckID: c.Params("deckID"),
		},
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
	rspDeckCards, err := e.cardDeckService.GetDeckCards(c.Context(), &pbCommon.DeckRequest{
		UserID: userID,
		Deck: &pbCommon.Deck{
			DeckID: c.Params("deckID"),
		},
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
	rspCreateCard, err := e.cardDeckService.CreateCard(c.Context(), &pbCommon.CardRequest{
		UserID: userID,
		Card: &pbCommon.Card{
			DeckID: c.Params("deckID"),
			Sides:  converter.ConvertToTypeArray(data.Sides, converter.FiberCardSideContentToProtoCardSideContent),
		},
	})
	if err != nil {
		return err
	}
	return c.SendString(rspCreateCard.CardID)
}

func (e *Frontend) GetCardHandler(c *fiber.Ctx) error {
	userID := helper.GetUserIDFromContext(c)
	rspGetCard, err := e.cardDeckService.GetCard(c.Context(), &pbCommon.CardRequest{
		UserID: userID,
		Card: &pbCommon.Card{
			CardID: c.Params("cardID"),
		},
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
	rspModifyCard, err := e.cardDeckService.ModifyCard(c.Context(), &pbCommon.CardRequest{
		UserID: userID,
		Card: &pbCommon.Card{
			CardID: c.Params("cardID"),
			Sides:  converter.ConvertToTypeArray(data.Sides, converter.FiberCardSideContentToProtoCardSideContent),
		},
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
	rspDeleteCard, err := e.cardDeckService.DeleteCard(c.Context(), &pbCommon.CardRequest{
		UserID: userID,
		Card: &pbCommon.Card{
			CardID: c.Params("cardID"),
		},
	})
	if err != nil {
		return err
	} else if !rspDeleteCard.Success {
		return helper.NewMicroNotSuccessfulResponseErr(helper.FrontendServiceID)
	}
	return c.SendStatus(200)
}

func (e *Frontend) GetFavoriteDecksHandler(c *fiber.Ctx) error {
	userID := helper.GetUserIDFromContext(c)
	rspGetFavoriteDecks, err := e.cardDeckService.GetUserFavoriteDecks(c.Context(), &pbCommon.User{UserID: userID})
	if err != nil {
		return err
	}
	decks := converter.ConvertToTypeArray(rspGetFavoriteDecks.Decks, converter.ProtoDeckToFiberDeckConverter)

	return c.JSON(converter.FiberGetGroupDecksResponseBody{
		Decks: decks,
	})
}

func (e *Frontend) AddFavoriteDeckHandler(c *fiber.Ctx) error {
	var deck = &pbCommon.Deck{}
	err := c.BodyParser(deck)
	if err != nil {
		return err
	}
	if deck.DeckID == "" {
		return helper.NewFiberBadRequestErr("No DeckID provided")
	}
	userID := helper.GetUserIDFromContext(c)
	rspAddFavoriteDeck, err := e.cardDeckService.AddUserFavoriteDeck(c.Context(), &pbCommon.DeckRequest{
		UserID: userID,
		Deck:   deck,
	})
	if err != nil {
		return err
	}
	if !rspAddFavoriteDeck.Success {
		return helper.NewMicroNotSuccessfulResponseErr(helper.FrontendServiceID)
	}
	return c.SendStatus(200)
}

func (e *Frontend) DelFavoriteDeckHandler(c *fiber.Ctx) error {
	var deck = &pbCommon.Deck{}
	err := c.BodyParser(deck)
	if err != nil {
		return err
	}
	if deck.DeckID == "" {
		return helper.NewFiberBadRequestErr("No DeckID provided")
	}
	userID := helper.GetUserIDFromContext(c)
	rspDelFavoriteDeck, err := e.cardDeckService.DelUserFavoriteDeck(c.Context(), &pbCommon.DeckRequest{
		UserID: userID,
		Deck:   deck,
	})
	if err != nil {
		return err
	}
	if !rspDelFavoriteDeck.Success {
		return helper.NewMicroNotSuccessfulResponseErr(helper.FrontendServiceID)
	}
	return c.SendStatus(200)
}

func (e *Frontend) GetActiveDecksHandler(c *fiber.Ctx) error {
	userID := helper.GetUserIDFromContext(c)
	rspGetActiveDecks, err := e.cardDeckService.GetUserActiveDecks(c.Context(), &pbCommon.User{UserID: userID})
	if err != nil {
		return err
	}
	decks := converter.ConvertToTypeArray(rspGetActiveDecks.Decks, converter.ProtoDeckToFiberDeckConverter)

	return c.JSON(converter.FiberGetGroupDecksResponseBody{
		Decks: decks,
	})
}

func (e *Frontend) DelActiveDeckHandler(c *fiber.Ctx) error {
	var deck = &pbCommon.Deck{}
	err := c.BodyParser(deck)
	if err != nil {
		return err
	}
	if deck.DeckID == "" {
		return helper.NewFiberBadRequestErr("No DeckID provided")
	}
	userID := helper.GetUserIDFromContext(c)
	rspDelActiveDeck, err := e.cardDeckService.DelUserActiveDeck(c.Context(), &pbCommon.DeckRequest{
		UserID: userID,
		Deck:   deck,
	})
	if err != nil {
		return err
	}
	if !rspDelActiveDeck.Success {
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
	rspCreateCardSide, err := e.cardDeckService.CreateCardSide(c.Context(), &pbCommon.CardSideRequest{
		UserID: userID,
		CardID: c.Params("cardID"),
		CardSide: &pbCommon.CardSide{
			Header:      data.FiberCardSideContent.Header,
			Description: data.FiberCardSideContent.Description,
		},
		PlaceBeforeCardSideID: data.PlaceBeforeCardSideID,
	})
	if err != nil {
		return err
	}
	return c.SendString(rspCreateCardSide.CardSideID)
}

func (e *Frontend) ModifyCardSideHandler(c *fiber.Ctx) error {
	var data converter.FiberModifyCardSideRequestBody
	if err := c.BodyParser(&data); err != nil {
		return err
	}
	userID := helper.GetUserIDFromContext(c)
	rspModifyCardSide, err := e.cardDeckService.ModifyCardSide(c.Context(), &pbCommon.CardSideRequest{
		UserID: userID,
		CardSide: &pbCommon.CardSide{
			CardSideID:  c.Params("cardSideID"),
			Header:      data.FiberCardSideContent.Header,
			Description: data.FiberCardSideContent.Description,
		},
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
	rspDeleteCardSide, err := e.cardDeckService.DeleteCardSide(c.Context(), &pbCommon.CardSideRequest{
		UserID: userID,
		CardSide: &pbCommon.CardSide{
			CardSideID: c.Params("cardSideID"),
		},
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
	rspSrsPull, err := e.srsService.Pull(c.Context(), &pbCommon.DeckRequest{
		UserID: userID,
		Deck: &pbCommon.Deck{
			DeckID: c.Params("deckID"),
		},
	})
	if err != nil {
		return err
	}
	return c.JSON(rspSrsPull)
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
	rspSrsPush, err := e.srsService.Push(c.Context(), &pbSrs.SrsPushRequest{
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
	rspSrsDue, err := e.srsService.GetDeckCardsDue(c.Context(), &pbCommon.DeckRequest{
		UserID: userID,
		Deck: &pbCommon.Deck{
			DeckID: c.Params("deckID"),
		},
	})
	if err != nil {
		return err
	}
	return c.JSON(converter.FiberGetDueResponseBody{
		DueCards: rspSrsDue.DueCards,
	})
}

func (e *Frontend) SrsUserDueHandler(c *fiber.Ctx) error {
	userID := helper.GetUserIDFromContext(c)
	dueCards, err := e.srsService.GetUserCardsDue(c.Context(), &pbCommon.User{
		UserID: userID,
	})
	if err != nil {
		return err
	}
	return c.JSON(converter.FiberGetDueResponseBody{
		DueCards: dueCards.DueCards,
		DueDecks: dueCards.DueDecks,
	})
}
