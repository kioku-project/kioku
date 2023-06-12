package converter_test

import (
	"testing"
	"time"

	"github.com/kioku-project/kioku/pkg/converter"
	"github.com/kioku-project/kioku/pkg/model"

	pbCardDeck "github.com/kioku-project/kioku/services/carddeck/proto"
	pbCollaboration "github.com/kioku-project/kioku/services/collaboration/proto"
	"github.com/stretchr/testify/assert"
)

const (
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

	pbCard = pbCardDeck.Card{
		CardID: id,
		Sides:  []*pbCardDeck.CardSide{},
	}

	cardSide = model.CardSide{
		ID:          id,
		Header:      header,
		Description: desc,
	}

	pbCardSide = pbCardDeck.CardSide{
		Header:      header,
		Description: desc,
	}

	cardSideFiber = converter.FiberCardSideContent{
		Header:      header,
		Description: desc,
	}

	groupMembers = pbCollaboration.UserWithRole{
		User: &pbCollaboration.User{
			UserID: id,
			Name:   name,
		},
		GroupRole: pbCollaboration.GroupRole_READ,
	}
)

func TestMigrateModelRoleToProtoRole(t *testing.T) {
	modelRoles := []model.RoleType{
		model.RoleRead,
		model.RoleWrite,
		model.RoleAdmin,
	}

	protoRoles := []pbCollaboration.GroupRole{
		pbCollaboration.GroupRole_READ,
		pbCollaboration.GroupRole_WRITE,
		pbCollaboration.GroupRole_ADMIN,
	}

	for idx, modelRole := range modelRoles {
		converted := converter.MigrateModelRoleToProtoRole(modelRole)
		assert.Equal(t, protoRoles[idx], converted)
	}
}

func TestMigrateModelGroupTypeToProtoGroupType(t *testing.T) {
	modelTypes := []model.GroupType{
		model.Public,
		model.Private,
	}

	protoTypes := []pbCollaboration.GroupType{
		pbCollaboration.GroupType_PUBLIC,
		pbCollaboration.GroupType_PRIVATE,
	}

	for idx, modelType := range modelTypes {
		converted := converter.MigrateModelGroupTypeToProtoGroupType(modelType)
		assert.Equal(t, protoTypes[idx], converted)
	}
}

func TestMigrateStringGroupTypeToProtoGroupType(t *testing.T) {
	stringTypes := []string{
		pbCollaboration.GroupType_PUBLIC.String(),
		pbCollaboration.GroupType_PRIVATE.String(),
		// InvalidStringGroupType
		"",
	}

	protoTypes := []pbCollaboration.GroupType{
		pbCollaboration.GroupType_PUBLIC,
		pbCollaboration.GroupType_PRIVATE,
		pbCollaboration.GroupType_INVALID,
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
	assert.Equal(t, user.Email, conv.Email)
	assert.Equal(t, user.Name, conv.Name)
}

func TestStoreGroupUserRoleToProtoUserIDConverter(t *testing.T) {
	role := model.GroupUserRole{
		UserID: id,
	}

	conv := converter.StoreGroupUserRoleToProtoUserIDConverter(role)
	assert.Equal(t, id, conv.UserID)
}

func TestStoreGroupAdmissionToProtoUserIDConverter(t *testing.T) {
	role := model.GroupAdmission{
		UserID: id,
	}

	conv := converter.StoreGroupAdmissionToProtoUserIDConverter(role)
	assert.Equal(t, id, conv.UserID)
}

func TestStoreGroupAdmissionToProtoGroupInvitationConverter(t *testing.T) {
	admission := model.GroupAdmission{
		ID:      id,
		GroupID: groupID,
		Group: model.Group{
			Name: groupName,
		},
	}

	conv := converter.StoreGroupAdmissionToProtoGroupInvitationConverter(admission)
	assert.Equal(t, id, conv.AdmissionID)
	assert.Equal(t, groupID, conv.GroupID)
	assert.Equal(t, groupName, conv.GroupName)
}

func TestProtoGroupMemberRequestToFiberGroupMemberRequestConverter(t *testing.T) {
	groupMemberRequest := pbCollaboration.MemberRequest{
		AdmissionID: id,
		User: &pbCollaboration.User{
			UserID: id,
			Name:   name,
		},
	}

	conv := converter.ProtoGroupMemberRequestToFiberGroupMemberRequestConverter(&groupMemberRequest)

	assert.Equal(t, id, conv.AdmissionID)
	assert.Equal(t, id, conv.UserID)
	assert.Equal(t, name, conv.Name)
}

func TestStoreGroupToProtoGroupConverter(t *testing.T) {
	group := model.Group{
		ID:          id,
		Name:        name,
		Description: desc,
		IsDefault:   isDefault,
		GroupType:   model.Public,
	}

	conv := converter.StoreGroupToProtoGroupConverter(group)

	assert.Equal(t, id, conv.GroupID)
	assert.Equal(t, name, conv.GroupName)
	assert.Equal(t, desc, conv.GroupDescription)
	assert.Equal(t, isDefault, conv.IsDefault)
	assert.Equal(t, converter.MigrateModelGroupTypeToProtoGroupType(model.Public), conv.GroupType)
}

func TestProtoGroupToFiberGroupConverter(t *testing.T) {
	group := pbCollaboration.Group{
		GroupID:          id,
		GroupName:        name,
		GroupDescription: desc,
		IsDefault:        isDefault,
	}

	conv := converter.ProtoGroupToFiberGroupConverter(&group)

	assert.Equal(t, id, conv.GroupID)
	assert.Equal(t, name, conv.GroupName)
	assert.Equal(t, desc, conv.GroupDescription)
	assert.Equal(t, isDefault, conv.IsDefault)
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
	assert.Equal(t, timeConstant.Unix(), conv.CreatedAT)
	assert.Equal(t, id, conv.GroupID)
}

func TestStoreCardToProtoCardConverter(t *testing.T) {

	conv := converter.StoreCardToProtoCardConverter(card)

	assert.Equal(t, id, conv.CardID)
}

func TestCardDeckProtoCardToSrsProtoCardConverter(t *testing.T) {
	conv := converter.CardDeckProtoCardToSrsProtoCardConverter(&pbCard)

	assert.Equal(t, id, conv.CardID)
}

func TestStoreCardSideToProtoCardSideConverter(t *testing.T) {
	conv := converter.StoreCardSideToProtoCardSideConverter(cardSide)

	assert.Equal(t, id, conv.CardSideID)
	assert.Equal(t, header, conv.Header)
	assert.Equal(t, desc, conv.Description)
}

func TestCardDeckProtoCardSideToSrsProtoCardSideConverter(t *testing.T) {
	conv := converter.CardDeckProtoCardSideToSrsProtoCardSideConverter(&pbCardSide)

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
	assert.Equal(t, pbCollaboration.GroupRole_READ.String(), conv.GroupRole)
}
