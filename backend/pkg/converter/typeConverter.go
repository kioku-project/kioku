package converter

import (
	"github.com/kioku-project/kioku/pkg/helper"
	"github.com/kioku-project/kioku/pkg/model"
	pbCommon "github.com/kioku-project/kioku/pkg/proto"
)

func MigrateModelRoleToProtoRole(modelRole model.RoleType) (protoRole pbCommon.GroupRole) {
	if modelRole == model.RoleRequested {
		protoRole = pbCommon.GroupRole_REQUESTED
	} else if modelRole == model.RoleInvited {
		protoRole = pbCommon.GroupRole_INVITED
	} else if modelRole == model.RoleRead {
		protoRole = pbCommon.GroupRole_READ
	} else if modelRole == model.RoleWrite {
		protoRole = pbCommon.GroupRole_WRITE
	} else if modelRole == model.RoleAdmin {
		protoRole = pbCommon.GroupRole_ADMIN
	}
	return
}

func MigrateProtoRoleToModelRole(protoRole pbCommon.GroupRole) (modelRole model.RoleType) {
	switch protoRole {
	case pbCommon.GroupRole_REQUESTED:
		modelRole = model.RoleRequested
	case pbCommon.GroupRole_INVITED:
		modelRole = model.RoleInvited
	case pbCommon.GroupRole_READ:
		modelRole = model.RoleRead
	case pbCommon.GroupRole_WRITE:
		modelRole = model.RoleWrite
	case pbCommon.GroupRole_ADMIN:
		modelRole = model.RoleAdmin
	}
	return
}

func MigrateModelAlgorithmToProtoAlgorithm(modelAlgo model.AlgorithmType) (protoAlgo pbCommon.AlgoType) {
	if modelAlgo == model.AlgoDynamicSRS {
		protoAlgo = pbCommon.AlgoType_DYNAMIC_SRS
	} else if modelAlgo == model.AlgoLinearSRS {
		protoAlgo = pbCommon.AlgoType_LINEAR_SRS
	} else if modelAlgo == model.AlgoStaticSRS {
		protoAlgo = pbCommon.AlgoType_STATIC_SRS
	} else if modelAlgo == model.AlgoTestSRS {
		protoAlgo = pbCommon.AlgoType_TEST_SRS
	}
	return
}

func MigrateProtoAlgorithmToModelAlgorithm(protoAlgo pbCommon.AlgoType) (modelAlgo model.AlgorithmType) {
	switch protoAlgo {
	case pbCommon.AlgoType_DYNAMIC_SRS:
		modelAlgo = model.AlgoDynamicSRS
	case pbCommon.AlgoType_LINEAR_SRS:
		modelAlgo = model.AlgoLinearSRS
	case pbCommon.AlgoType_STATIC_SRS:
		modelAlgo = model.AlgoStaticSRS
	case pbCommon.AlgoType_TEST_SRS:
		modelAlgo = model.AlgoTestSRS
	}
	return
}

func MigrateStringRoleToProtoRole(stringRole string) (protoRole pbCommon.GroupRole) {
	switch stringRole {
	case pbCommon.GroupRole_REQUESTED.String():
		protoRole = pbCommon.GroupRole_REQUESTED
	case pbCommon.GroupRole_INVITED.String():
		protoRole = pbCommon.GroupRole_INVITED
	case pbCommon.GroupRole_READ.String():
		protoRole = pbCommon.GroupRole_READ
	case pbCommon.GroupRole_WRITE.String():
		protoRole = pbCommon.GroupRole_WRITE
	case pbCommon.GroupRole_ADMIN.String():
		protoRole = pbCommon.GroupRole_ADMIN
	}
	return
}

func MigrateModelGroupTypeToProtoGroupType(modelType model.GroupType) (protoType pbCommon.GroupType) {
	if modelType == model.OpenGroupType {
		protoType = pbCommon.GroupType_OPEN
	} else if modelType == model.RequestGroupType {
		protoType = pbCommon.GroupType_REQUEST
	} else if modelType == model.ClosedGroupType {
		protoType = pbCommon.GroupType_CLOSED
	}
	return
}

func MigrateProtoGroupTypeToModelGroupType(protoType pbCommon.GroupType) (modelType model.GroupType) {
	if protoType == pbCommon.GroupType_OPEN {
		modelType = model.OpenGroupType
	} else if protoType == pbCommon.GroupType_REQUEST {
		modelType = model.RequestGroupType
	} else if protoType == pbCommon.GroupType_CLOSED {
		modelType = model.ClosedGroupType
	}
	return
}

func MigrateStringGroupTypeToProtoGroupType(stringType string) pbCommon.GroupType {
	if stringType == pbCommon.GroupType_OPEN.String() {
		return pbCommon.GroupType_OPEN
	}
	if stringType == pbCommon.GroupType_REQUEST.String() {
		return pbCommon.GroupType_REQUEST
	}
	if stringType == pbCommon.GroupType_CLOSED.String() {
		return pbCommon.GroupType_CLOSED
	}
	return pbCommon.GroupType_GT_INVALID
}

func MigrateModelDeckTypeToProtoDeckType(modelType model.DeckType) (protoType pbCommon.DeckType, err error) {
	if modelType == model.PublicDeckType {
		protoType = pbCommon.DeckType_PUBLIC
	} else if modelType == model.PrivateDeckType {
		protoType = pbCommon.DeckType_PRIVATE
	} else {
		err = helper.NewMicroDeckTypeNotValidErr(helper.CardDeckServiceID)
	}
	return
}

func MigrateProtoDeckTypeToModelDeckType(protoType pbCommon.DeckType) (modelType model.DeckType, err error) {
	if protoType == pbCommon.DeckType_PUBLIC {
		modelType = model.PublicDeckType
	} else if protoType == pbCommon.DeckType_PRIVATE {
		modelType = model.PrivateDeckType
	} else {
		err = helper.NewMicroDeckTypeNotValidErr(helper.CardDeckServiceID)
	}
	return
}

func MigrateStringDeckTypeToProtoDeckType(stringType string) pbCommon.DeckType {
	if stringType == pbCommon.DeckType_PUBLIC.String() {
		return pbCommon.DeckType_PUBLIC
	}
	if stringType == pbCommon.DeckType_PRIVATE.String() {
		return pbCommon.DeckType_PRIVATE
	}
	return pbCommon.DeckType_DT_INVALID
}

func StoreUserToProtoUserProfileInformationResponseConverter(user model.User) *pbCommon.User {
	return &pbCommon.User{
		UserID:    user.ID,
		UserEmail: user.Email,
		UserName:  user.Name,
	}
}

func StoreGroupUserRoleToProtoUserIDConverter(role model.GroupUserRole) *pbCommon.User {
	return &pbCommon.User{UserID: role.UserID}
}

func StoreGroupAdmissionToProtoUserIDConverter(groupRole model.GroupUserRole) *pbCommon.User {
	return &pbCommon.User{UserID: groupRole.UserID}
}

func StoreGroupAdmissionToProtoGroupInvitationConverter(groupRole model.GroupUserRole) *pbCommon.Group {
	return &pbCommon.Group{
		GroupID:   groupRole.GroupID,
		GroupName: groupRole.Group.Name,
	}
}

func ProtoGroupMemberRequestToFiberGroupMemberRequestConverter(groupMemberRequest *pbCommon.User) FiberGroupMemberAdmission {
	return FiberGroupMemberAdmission{
		UserID: groupMemberRequest.UserID,
		Name:   groupMemberRequest.UserName,
		Email:  groupMemberRequest.UserEmail,
	}
}

func StoreGroupToProtoGroupConverter(group model.Group) *pbCommon.Group {
	return &pbCommon.Group{
		GroupID:          group.ID,
		GroupName:        group.Name,
		GroupDescription: group.Description,
		IsDefault:        group.IsDefault,
		GroupType:        MigrateModelGroupTypeToProtoGroupType(group.GroupType),
	}
}

func ProtoGroupWithRoleToFiberGroupConverter(group *pbCommon.Group) FiberGroup {
	return FiberGroup{
		GroupID:          group.GroupID,
		GroupName:        group.GroupName,
		GroupDescription: group.GroupDescription,
		IsDefault:        group.IsDefault,
		GroupType:        group.GroupType.String(),
		GroupRole:        group.Role.String(),
	}
}

func ProtoDeckToFiberDeckConverter(deck *pbCommon.Deck) FiberDeck {
	return FiberDeck{
		DeckID:          deck.DeckID,
		DeckName:        deck.DeckName,
		DeckDescription: deck.DeckDescription,
		DeckType:        deck.DeckType.String(),
		GroupID:         deck.GroupID,
		IsActive:        deck.IsActive,
		IsFavorite:      deck.IsFavorite,
		DeckRole:        deck.DeckRole.String(),
	}
}

func ProtoDeckRespToFiberDeckConverter(deck *pbCommon.Deck) FiberDeck {
	return FiberDeck{
		DeckID:          deck.DeckID,
		DeckName:        deck.DeckName,
		DeckDescription: deck.DeckDescription,
		DeckType:        deck.DeckType.String(),
		GroupID:         deck.GroupID,
		IsActive:        deck.IsActive,
		IsFavorite:      deck.IsFavorite,
		DeckRole:        deck.DeckRole.String(),
	}
}

func StoreDeckToProtoDeckConverter(deck model.Deck) *pbCommon.Deck {
	dt, _ := MigrateModelDeckTypeToProtoDeckType(deck.DeckType)
	return &pbCommon.Deck{
		DeckID:          deck.ID,
		DeckName:        deck.Name,
		DeckDescription: deck.Description,
		DeckType:        dt,
		GroupID:         deck.GroupID,
		IsFavorite:      deck.IsFavorite,
		IsActive:        deck.IsActive,
		Algorithm:       MigrateModelAlgorithmToProtoAlgorithm(deck.Algorithm),
	}
}

func StoreCardToProtoCardConverter(card model.Card) *pbCommon.Card {
	return &pbCommon.Card{
		CardID: card.ID,
		Sides:  ConvertToTypeArray(card.CardSides, StoreCardSideToProtoCardSideConverter),
	}
}

func StoreCardSideToProtoCardSideConverter(cardSide model.CardSide) *pbCommon.CardSide {
	return &pbCommon.CardSide{
		CardSideID:  cardSide.ID,
		Header:      cardSide.Header,
		Description: cardSide.Description,
	}
}

func FiberCardSideContentToProtoCardSideContent(cardSide FiberCardSideContent) *pbCommon.CardSide {
	return &pbCommon.CardSide{
		Header:      cardSide.Header,
		Description: cardSide.Description,
	}
}

func ProtoUserWithRoleToFiberGroupMember(groupMembers *pbCommon.User) FiberGroupMember {
	return FiberGroupMember{
		UserID:    groupMembers.UserID,
		Name:      groupMembers.UserName,
		GroupRole: groupMembers.GroupRole.String(),
	}
}

func StoreNotificationSubscriptionToProtoNotificationSubscriptionConverter(subscription *model.PushSubscription) *pbCommon.PushSubscription {
	return &pbCommon.PushSubscription{
		SubscriptionID: subscription.ID,
		Endpoint:       subscription.Endpoint,
		P256Dh:         subscription.P256DH,
		Auth:           subscription.Auth,
	}
}

func ProtoNotificationSubscriptionToIDStringConverter(subscription *pbCommon.PushSubscription) string {
	return subscription.SubscriptionID
}
