package converter_test

import (
	pbCommon "github.com/kioku-project/kioku/pkg/proto"
	"testing"
	"time"

	"github.com/kioku-project/kioku/pkg/converter"
	"github.com/kioku-project/kioku/pkg/model"

	"github.com/stretchr/testify/assert"
)

var (
	id        = "id"
	name      = "name"
	email     = "email"
	password  = "password"
	groupID   = "groupID"
	groupName = "groupName"
	desc      = "desc"
	isDefault = true
	header    = "header"
)

var (
	timeConstant = time.Now()

	deck = model.Deck{
		ID:        id,
		Name:      name,
		CreatedAt: timeConstant,
		GroupID:   id,
	}

	card = model.Card{
		ID:        id,
		CardSides: []model.CardSide{},
	}

	pbCard = pbCommon.Card{
		CardID: id,
		Sides:  []*pbCommon.CardSide{},
	}

	cardSide = model.CardSide{
		ID:          id,
		Header:      header,
		Description: desc,
	}

	pbCardSide = pbCommon.CardSide{
		Header:      header,
		Description: desc,
	}

	cardSideFiber = converter.FiberCardSideContent{
		Header:      header,
		Description: desc,
	}

	groupMembers = pbCommon.User{
		UserID:    id,
		UserName:  name,
		GroupRole: pbCommon.GroupRole_READ,
	}
)

func TestMigrateModelRoleToProtoRole(t *testing.T) {
	modelRoles := []model.RoleType{
		model.RoleRead,
		model.RoleWrite,
		model.RoleAdmin,
	}

	protoRoles := []pbCommon.GroupRole{
		pbCommon.GroupRole_READ,
		pbCommon.GroupRole_WRITE,
		pbCommon.GroupRole_ADMIN,
	}

	for idx, modelRole := range modelRoles {
		converted := converter.MigrateModelRoleToProtoRole(modelRole)
		assert.Equal(t, protoRoles[idx], converted)
	}
}

func TestMigrateModelGroupTypeToProtoGroupType(t *testing.T) {
	modelTypes := []model.GroupType{
		model.OpenGroupType,
		model.RequestGroupType,
		model.ClosedGroupType,
	}

	protoTypes := []pbCommon.GroupType{
		pbCommon.GroupType_OPEN,
		pbCommon.GroupType_REQUEST,
		pbCommon.GroupType_CLOSED,
	}

	for idx, modelType := range modelTypes {
		converted := converter.MigrateModelGroupTypeToProtoGroupType(modelType)
		assert.Equal(t, protoTypes[idx], converted)
	}
}

func TestMigrateStringGroupTypeToProtoGroupType(t *testing.T) {
	stringTypes := []string{
		pbCommon.GroupType_OPEN.String(),
		pbCommon.GroupType_REQUEST.String(),
		pbCommon.GroupType_CLOSED.String(),
		// InvalidStringGroupType
		"",
	}

	protoTypes := []pbCommon.GroupType{
		pbCommon.GroupType_OPEN,
		pbCommon.GroupType_REQUEST,
		pbCommon.GroupType_CLOSED,
		pbCommon.GroupType_GT_INVALID,
	}

	for idx, modelType := range stringTypes {
		converted := converter.MigrateStringGroupTypeToProtoGroupType(modelType)
		assert.Equal(t, protoTypes[idx], converted)
	}
}

func TestStoreUserToProtoUserProfileInformationResponseConverter(t *testing.T) {
	user := model.User{
		ID:       id,
		Name:     name,
		Email:    email,
		Password: password,
	}

	conv := converter.StoreUserToProtoUserProfileInformationResponseConverter(user)
	assert.Equal(t, user.ID, conv.UserID)
	assert.Equal(t, user.Email, conv.UserEmail)
	assert.Equal(t, user.Name, conv.UserName)
}

func TestStoreGroupUserRoleToProtoUserIDConverter(t *testing.T) {
	role := model.GroupUserRole{
		UserID: id,
	}

	conv := converter.StoreGroupUserRoleToProtoUserIDConverter(role)
	assert.Equal(t, id, conv.UserID)
}

func TestStoreGroupAdmissionToProtoUserIDConverter(t *testing.T) {
	role := model.GroupUserRole{
		UserID: id,
	}

	conv := converter.StoreGroupAdmissionToProtoUserIDConverter(role)
	assert.Equal(t, id, conv.UserID)
}

func TestStoreGroupAdmissionToProtoGroupInvitationConverter(t *testing.T) {
	admission := model.GroupUserRole{
		GroupID: groupID,
		Group: model.Group{
			Name: groupName,
		},
	}

	conv := converter.StoreGroupAdmissionToProtoGroupInvitationConverter(admission)
	assert.Equal(t, groupID, conv.GroupID)
	assert.Equal(t, groupName, conv.GroupName)
}

func TestProtoGroupMemberRequestToFiberGroupMemberRequestConverter(t *testing.T) {
	groupMemberRequest := pbCommon.User{
		UserID:    id,
		UserName:  name,
		UserEmail: email,
	}

	conv := converter.ProtoGroupMemberRequestToFiberGroupMemberRequestConverter(&groupMemberRequest)

	assert.Equal(t, id, conv.UserID)
	assert.Equal(t, name, conv.Name)
}

func TestStoreGroupToProtoGroupConverter(t *testing.T) {
	group := model.Group{
		ID:          id,
		Name:        name,
		Description: desc,
		IsDefault:   isDefault,
		GroupType:   model.OpenGroupType,
	}

	conv := converter.StoreGroupToProtoGroupConverter(group)

	assert.Equal(t, id, conv.GroupID)
	assert.Equal(t, name, conv.GroupName)
	assert.Equal(t, desc, conv.GroupDescription)
	assert.Equal(t, isDefault, conv.IsDefault)
	assert.Equal(t, converter.MigrateModelGroupTypeToProtoGroupType(model.OpenGroupType), conv.GroupType)
}

func TestProtoGroupToFiberGroupConverter(t *testing.T) {
	group := pbCommon.Group{
		GroupID:          id,
		GroupName:        name,
		GroupDescription: desc,
		IsDefault:        isDefault,
		GroupType:        pbCommon.GroupType_OPEN,
		Role:             pbCommon.GroupRole_READ,
	}

	conv := converter.ProtoGroupWithRoleToFiberGroupConverter(&group)

	assert.Equal(t, id, conv.GroupID)
	assert.Equal(t, name, conv.GroupName)
	assert.Equal(t, desc, conv.GroupDescription)
	assert.Equal(t, isDefault, conv.IsDefault)
	assert.Equal(t, pbCommon.GroupType_OPEN.String(), conv.GroupType)
	assert.Equal(t, pbCommon.GroupRole_READ.String(), conv.GroupRole)
}

func TestStoreDeckToProtoDeckConverter(t *testing.T) {
	conv := converter.StoreDeckToProtoDeckConverter(deck)

	assert.Equal(t, id, conv.DeckID)
	assert.Equal(t, name, conv.DeckName)
}

func TestStoreDeckToProtoDeckResponseConverter(t *testing.T) {
	conv := converter.StoreDeckToProtoDeckResponseConverter(deck)

	assert.Equal(t, id, conv.DeckID)
	assert.Equal(t, name, conv.DeckName)
	assert.Equal(t, timeConstant.Unix(), conv.CreatedAt)
	assert.Equal(t, id, conv.GroupID)
}

func TestStoreCardToProtoCardConverter(t *testing.T) {
	conv := converter.StoreCardToProtoCardConverter(card)

	assert.Equal(t, id, conv.CardID)
}

func TestStoreCardSideToProtoCardSideConverter(t *testing.T) {
	conv := converter.StoreCardSideToProtoCardSideConverter(cardSide)

	assert.Equal(t, id, conv.CardSideID)
	assert.Equal(t, header, conv.Header)
	assert.Equal(t, desc, conv.Description)
}

func TestFiberCardSideContentToProtoCardSideContent(t *testing.T) {
	conv := converter.FiberCardSideContentToProtoCardSideContent(cardSideFiber)

	assert.Equal(t, header, conv.Header)
	assert.Equal(t, desc, conv.Description)
}

func TestProtoUserWithRoleToFiberGroupMember(t *testing.T) {
	conv := converter.ProtoUserWithRoleToFiberGroupMember(&groupMembers)

	assert.Equal(t, id, conv.UserID)
	assert.Equal(t, name, conv.Name)
	assert.Equal(t, pbCommon.GroupRole_READ.String(), conv.GroupRole)
}
