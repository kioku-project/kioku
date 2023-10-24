package converter

import (
	"github.com/kioku-project/kioku/pkg/helper"
	"github.com/kioku-project/kioku/pkg/model"
	pbCardDeck "github.com/kioku-project/kioku/services/carddeck/proto"
	pbCollaboration "github.com/kioku-project/kioku/services/collaboration/proto"
	pbSrs "github.com/kioku-project/kioku/services/srs/proto"
	pbUser "github.com/kioku-project/kioku/services/user/proto"
)

func MigrateModelRoleToProtoRole(modelRole model.RoleType) (protoRole pbCollaboration.GroupRole) {
	if modelRole == model.RoleRequested {
		protoRole = pbCollaboration.GroupRole_REQUESTED
	} else if modelRole == model.RoleInvited {
		protoRole = pbCollaboration.GroupRole_INVITED
	} else if modelRole == model.RoleRead {
		protoRole = pbCollaboration.GroupRole_READ
	} else if modelRole == model.RoleWrite {
		protoRole = pbCollaboration.GroupRole_WRITE
	} else if modelRole == model.RoleAdmin {
		protoRole = pbCollaboration.GroupRole_ADMIN
	}
	return
}

func MigrateProtoRoleToModelRole(protoRole pbCollaboration.GroupRole) (modelRole model.RoleType) {
	switch protoRole {
	case pbCollaboration.GroupRole_REQUESTED:
		modelRole = model.RoleRequested
	case pbCollaboration.GroupRole_INVITED:
		modelRole = model.RoleInvited
	case pbCollaboration.GroupRole_READ:
		modelRole = model.RoleRead
	case pbCollaboration.GroupRole_WRITE:
		modelRole = model.RoleWrite
	case pbCollaboration.GroupRole_ADMIN:
		modelRole = model.RoleAdmin
	}
	return
}

func MigrateStringRoleToProtoRole(stringRole string) (protoRole pbCollaboration.GroupRole) {
	switch stringRole {
	case "requested":
		protoRole = pbCollaboration.GroupRole_REQUESTED
	case "invited":
		protoRole = pbCollaboration.GroupRole_INVITED
	case "read":
		protoRole = pbCollaboration.GroupRole_READ
	case "write":
		protoRole = pbCollaboration.GroupRole_WRITE
	case "admin":
		protoRole = pbCollaboration.GroupRole_ADMIN
	}
	return
}

func MigrateModelGroupTypeToProtoGroupType(modelType model.GroupType) (protoType pbCollaboration.GroupType) {
	if modelType == model.OpenGroupType {
		protoType = pbCollaboration.GroupType_OPEN
	} else if modelType == model.RequestGroupType {
		protoType = pbCollaboration.GroupType_REQUEST
	} else if modelType == model.ClosedGroupType {
		protoType = pbCollaboration.GroupType_CLOSED
	}
	return
}

func MigrateProtoGroupTypeToModelGroupType(protoType pbCollaboration.GroupType) (modelType model.GroupType) {
	if protoType == pbCollaboration.GroupType_OPEN {
		modelType = model.OpenGroupType
	} else if protoType == pbCollaboration.GroupType_REQUEST {
		modelType = model.RequestGroupType
	} else if protoType == pbCollaboration.GroupType_CLOSED {
		modelType = model.ClosedGroupType
	}
	return
}

func MigrateStringGroupTypeToProtoGroupType(stringType string) pbCollaboration.GroupType {
	if stringType == pbCollaboration.GroupType_OPEN.String() {
		return pbCollaboration.GroupType_OPEN
	}
	if stringType == pbCollaboration.GroupType_REQUEST.String() {
		return pbCollaboration.GroupType_REQUEST
	}
	if stringType == pbCollaboration.GroupType_CLOSED.String() {
		return pbCollaboration.GroupType_CLOSED
	}
	return pbCollaboration.GroupType_INVALID
}

func MigrateModelDeckTypeToProtoDeckType(modelType model.DeckType) (protoType pbCardDeck.DeckType, err error) {
	if modelType == model.PublicDeckType {
		protoType = pbCardDeck.DeckType_PUBLIC
	} else if modelType == model.PrivateDeckType {
		protoType = pbCardDeck.DeckType_PRIVATE
	} else {
		err = helper.NewMicroDeckTypeNotValidErr(helper.CardDeckServiceID)
	}
	return
}

func MigrateProtoDeckTypeToModelDeckType(protoType pbCardDeck.DeckType) (err error, modelType model.DeckType) {
	if protoType == pbCardDeck.DeckType_PUBLIC {
		modelType = model.PublicDeckType
	} else if protoType == pbCardDeck.DeckType_PRIVATE {
		modelType = model.PrivateDeckType
	} else {
		err = helper.NewMicroDeckTypeNotValidErr(helper.CardDeckServiceID)
	}
	return
}

func MigrateStringDeckTypeToProtoDeckType(stringType string) pbCardDeck.DeckType {
	if stringType == pbCardDeck.DeckType_PUBLIC.String() {
		return pbCardDeck.DeckType_PUBLIC
	}
	if stringType == pbCardDeck.DeckType_PRIVATE.String() {
		return pbCardDeck.DeckType_PRIVATE
	}
	return pbCardDeck.DeckType_INVALID
}

func StoreUserToProtoUserProfileInformationResponseConverter(user model.User) *pbUser.UserProfileInformationResponse {
	return &pbUser.UserProfileInformationResponse{
		UserID:    user.ID,
		UserEmail: user.Email,
		UserName:  user.Name,
	}
}

func StoreGroupUserRoleToProtoUserIDConverter(role model.GroupUserRole) *pbUser.UserID {
	return &pbUser.UserID{UserID: role.UserID}
}

func StoreGroupAdmissionToProtoUserIDConverter(groupRole model.GroupUserRole) *pbUser.UserID {
	return &pbUser.UserID{UserID: groupRole.UserID}
}

func StoreGroupAdmissionToProtoGroupInvitationConverter(groupRole model.GroupUserRole) *pbCollaboration.GroupInvitation {
	return &pbCollaboration.GroupInvitation{
		GroupID:   groupRole.GroupID,
		GroupName: groupRole.Group.Name,
	}
}

func ProtoGroupMemberRequestToFiberGroupMemberRequestConverter(groupMemberRequest *pbCollaboration.MemberAdmission) FiberGroupMemberAdmission {
	return FiberGroupMemberAdmission{
		UserID: groupMemberRequest.User.UserID,
		Name:   groupMemberRequest.User.Name,
		Email:  *groupMemberRequest.User.Email,
	}
}

func StoreGroupToProtoGroupConverter(group model.Group) *pbCollaboration.Group {
	return &pbCollaboration.Group{
		GroupID:          group.ID,
		GroupName:        group.Name,
		GroupDescription: group.Description,
		IsDefault:        group.IsDefault,
		GroupType:        MigrateModelGroupTypeToProtoGroupType(group.GroupType),
	}
}

func ProtoGroupWithRoleToFiberGroupConverter(group *pbCollaboration.GroupWithUserRole) FiberGroup {
	return FiberGroup{
		GroupID:          group.Group.GroupID,
		GroupName:        group.Group.GroupName,
		GroupDescription: group.Group.GroupDescription,
		IsDefault:        group.Group.IsDefault,
		GroupType:        group.Group.GroupType.String(),
		GroupRole:        group.Role.String(),
	}
}

func ProtoDeckToFiberDeckConverter(deck *pbCardDeck.Deck) FiberDeck {
	return FiberDeck{
		DeckID:   deck.DeckID,
		DeckName: deck.DeckName,
		DeckType: deck.DeckType.String(),
	}
}

func ProtoDeckRespToFiberDeckConverter(deck *pbCardDeck.DeckResponse) FiberDeck {
	return FiberDeck{
		DeckID:   deck.DeckID,
		DeckName: deck.DeckName,
		DeckType: deck.DeckType.String(),
		GroupID:  deck.GroupID,
	}
}

func StoreDeckToProtoDeckConverter(deck model.Deck) *pbCardDeck.Deck {
	dt, _ := MigrateModelDeckTypeToProtoDeckType(deck.DeckType)
	return &pbCardDeck.Deck{
		DeckID:   deck.ID,
		DeckName: deck.Name,
		DeckType: dt,
	}
}

func StoreDeckToProtoDeckResponseConverter(deck model.Deck) *pbCardDeck.DeckResponse {
	dt, _ := MigrateModelDeckTypeToProtoDeckType(deck.DeckType)
	return &pbCardDeck.DeckResponse{
		DeckID:    deck.ID,
		DeckName:  deck.Name,
		CreatedAt: deck.CreatedAt.Unix(),
		DeckType:  dt,
		GroupID:   deck.GroupID,
	}
}

func StoreCardToProtoCardConverter(card model.Card) *pbCardDeck.Card {
	return &pbCardDeck.Card{
		CardID: card.ID,
		Sides:  ConvertToTypeArray(card.CardSides, StoreCardSideToProtoCardSideConverter),
	}
}

func CardDeckProtoCardToSrsProtoCardConverter(card *pbCardDeck.Card) *pbSrs.Card {
	return &pbSrs.Card{
		CardID: card.CardID,
		Sides:  ConvertToTypeArray(card.Sides, CardDeckProtoCardSideToSrsProtoCardSideConverter),
	}
}

func StoreCardSideToProtoCardSideConverter(cardSide model.CardSide) *pbCardDeck.CardSide {
	return &pbCardDeck.CardSide{
		CardSideID:  cardSide.ID,
		Header:      cardSide.Header,
		Description: cardSide.Description,
	}
}

func CardDeckProtoCardSideToSrsProtoCardSideConverter(cardSide *pbCardDeck.CardSide) *pbSrs.Side {
	return &pbSrs.Side{
		Header:      cardSide.Header,
		Description: cardSide.Description,
	}
}

func FiberCardSideContentToProtoCardSideContent(cardSide FiberCardSideContent) *pbCardDeck.CardSideContent {
	return &pbCardDeck.CardSideContent{
		Header:      cardSide.Header,
		Description: cardSide.Description,
	}
}

func ProtoUserWithRoleToFiberGroupMember(groupMembers *pbCollaboration.UserWithRole) FiberGroupMember {
	return FiberGroupMember{
		UserID:    groupMembers.User.UserID,
		Name:      groupMembers.User.Name,
		GroupRole: groupMembers.GroupRole.String(),
	}
}
