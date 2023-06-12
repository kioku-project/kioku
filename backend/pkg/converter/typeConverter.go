package converter

import (
	"github.com/kioku-project/kioku/pkg/model"
	pbCardDeck "github.com/kioku-project/kioku/services/carddeck/proto"
	pbCollaboration "github.com/kioku-project/kioku/services/collaboration/proto"
	pbSrs "github.com/kioku-project/kioku/services/srs/proto"
	pbUser "github.com/kioku-project/kioku/services/user/proto"
)

func MigrateModelRoleToProtoRole(modelRole model.RoleType) (protoRole pbCollaboration.GroupRole) {
	if modelRole == model.RoleRead {
		protoRole = pbCollaboration.GroupRole_READ
	} else if modelRole == model.RoleWrite {
		protoRole = pbCollaboration.GroupRole_WRITE
	} else if modelRole == model.RoleAdmin {
		protoRole = pbCollaboration.GroupRole_ADMIN
	}
	return
}

func MigrateModelGroupTypeToProtoGroupType(modelType model.GroupType) (protoType pbCollaboration.GroupType) {
	if modelType == model.Public {
		protoType = pbCollaboration.GroupType_PUBLIC
	} else if modelType == model.Private {
		protoType = pbCollaboration.GroupType_PRIVATE
	}
	return
}

func MigrateProtoGroupTypeToModelGroupType(protoType pbCollaboration.GroupType) (modelType model.GroupType) {
	if protoType == pbCollaboration.GroupType_PUBLIC {
		modelType = model.Public
	} else if protoType == pbCollaboration.GroupType_PRIVATE {
		modelType = model.Private
	}
	return
}

func MigrateStringGroupTypeToProtoGroupType(stringType string) pbCollaboration.GroupType {
	if stringType == pbCollaboration.GroupType_PUBLIC.String() {
		return pbCollaboration.GroupType_PUBLIC
	} else if stringType == pbCollaboration.GroupType_PRIVATE.String() {
		return pbCollaboration.GroupType_PRIVATE
	}
	return pbCollaboration.GroupType_INVALID
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

func StoreGroupAdmissionToProtoUserIDConverter(groupAdmission model.GroupAdmission) *pbUser.UserID {
	return &pbUser.UserID{UserID: groupAdmission.UserID}
}

func StoreGroupAdmissionToProtoGroupInvitationConverter(groupAdmission model.GroupAdmission) *pbCollaboration.GroupInvitation {
	return &pbCollaboration.GroupInvitation{
		AdmissionID: groupAdmission.ID,
		GroupID:     groupAdmission.GroupID,
		GroupName:   groupAdmission.Group.Name,
	}
}

func ProtoGroupMemberRequestToFiberGroupMemberRequestConverter(groupMemberRequest *pbCollaboration.MemberAdmission) FiberGroupMemberRequest {
	return FiberGroupMemberRequest{
		AdmissionID: groupMemberRequest.AdmissionID,
		UserID:      groupMemberRequest.User.UserID,
		Name:        groupMemberRequest.User.Name,
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

func StoreDeckToProtoDeckConverter(deck model.Deck) *pbCardDeck.Deck {
	return &pbCardDeck.Deck{
		DeckID:   deck.ID,
		DeckName: deck.Name,
	}
}

func StoreDeckToProtoDeckResponseConverter(deck model.Deck) *pbCardDeck.DeckResponse {
	return &pbCardDeck.DeckResponse{
		DeckID:    deck.ID,
		DeckName:  deck.Name,
		CreatedAt: deck.CreatedAt.Unix(),
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
