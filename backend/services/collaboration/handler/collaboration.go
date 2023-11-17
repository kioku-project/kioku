package handler

import (
	"context"
	"errors"

	"github.com/kioku-project/kioku/pkg/converter"
	"github.com/kioku-project/kioku/pkg/helper"
	"github.com/kioku-project/kioku/pkg/model"
	pbCommon "github.com/kioku-project/kioku/pkg/proto"
	pbCardDeck "github.com/kioku-project/kioku/services/carddeck/proto"
	pbSrs "github.com/kioku-project/kioku/services/srs/proto"
	pbUser "github.com/kioku-project/kioku/services/user/proto"
	"github.com/kioku-project/kioku/store"
	"go-micro.dev/v4/logger"
)

type Collaboration struct {
	store           store.CollaborationStore
	userService     pbUser.UserService
	srsService      pbSrs.SrsService
	cardDeckService pbCardDeck.CardDeckService
}

func New(
	s store.CollaborationStore,
	uS pbUser.UserService,
	srsS pbSrs.SrsService,
	cdS pbCardDeck.CardDeckService,
) *Collaboration {
	return &Collaboration{store: s, userService: uS, srsService: srsS, cardDeckService: cdS}
}

func (e *Collaboration) checkUserRoleAccessWithGroupAndRoleReturn(
	_ context.Context,
	userID string,
	groupID string,
	requiredRole pbCommon.GroupRole,
) (group *model.Group, protoRole pbCommon.GroupRole, err error) {
	logger.Infof("Find group with id %s", groupID)
	group, err = helper.FindStoreEntity(e.store.FindGroupByID, groupID, helper.CollaborationServiceID)
	if err != nil {
		return
	}
	logger.Infof("Requesting group role for user (%s)", userID)
	role, err := e.store.FindGroupUserRole(userID, groupID)
	if err != nil {
		if errors.Is(err, helper.ErrStoreNoEntryWithID) {
			err = helper.NewMicroNotAuthorizedErr(helper.CollaborationServiceID)
		}
		return
	}
	protoRole = converter.MigrateModelRoleToProtoRole(role)
	logger.Infof("Obtained group role (%s) for user (%s)", protoRole.String(), userID)
	if !helper.IsAuthorized(protoRole, requiredRole) {
		err = helper.NewMicroNotAuthorizedErr(helper.CollaborationServiceID)
		return
	}
	logger.Infof("Authenticated group role (%s) for user (%s)", protoRole.String(), userID)
	return
}

func (e *Collaboration) generateGroupMemberAdmissionResponse(
	ctx context.Context,
	groupAdmissions []model.GroupUserRole,
) (memberAdmissions []*pbCommon.User, err error) {
	userIDs := converter.ConvertToTypeArray(groupAdmissions, converter.StoreGroupAdmissionToProtoUserIDConverter)
	logger.Infof("Requesting information of users in group from user service")
	users, err := e.userService.GetUserInformation(ctx, &pbCommon.Users{Users: userIDs})
	if err != nil {
		return
	}
	memberAdmissions = make([]*pbCommon.User, len(users.Users))
	for i, user := range users.Users {
		memberAdmissions[i] = &pbCommon.User{
			UserID:    user.UserID,
			UserName:  user.UserName,
			UserEmail: user.UserEmail,
		}
	}
	logger.Infof("Successfully received user information from %d users and added it to request information",
		len(users.Users))
	return
}

func (e *Collaboration) GetGroupInvitations(
	_ context.Context,
	req *pbCommon.User,
	rsp *pbCommon.Groups,
) error {
	logger.Infof("Received Collaboration.GetGroupInvitations request: %v", req)
	groupInvitations, err := e.store.FindGroupInvitationsByUserID(req.UserID)
	if err != nil && !errors.Is(err, helper.ErrStoreNoEntryWithID) {
		return err
	}
	rsp.Groups = converter.ConvertToTypeArray(
		groupInvitations,
		converter.StoreGroupAdmissionToProtoGroupInvitationConverter,
	)
	logger.Infof("Found %d invitations for user with id %s", len(groupInvitations), req.UserID)
	return nil
}

func (e *Collaboration) createUserCardBindingsForWholeGroup(ctx context.Context, userID string, groupID string) error {
	decks, err := e.cardDeckService.GetGroupDecks(ctx, &pbCommon.GroupRequest{
		UserID: userID,
		Group: &pbCommon.Group{
			GroupID: groupID,
		},
	})
	if err != nil {
		return err
	}
	for _, deck := range decks.Decks {
		cards, err := e.cardDeckService.GetDeckCards(ctx, &pbCommon.DeckRequest{
			UserID: userID,
			Deck: &pbCommon.Deck{
				DeckID: deck.DeckID,
			},
		})
		if err != nil {
			return err
		}
		for _, card := range cards.Cards {
			_, err := e.srsService.AddUserCardBinding(ctx, &pbSrs.BindingRequest{
				UserID: userID,
				CardID: card.CardID,
				DeckID: deck.DeckID,
			})
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (e *Collaboration) GetUserGroups(_ context.Context, req *pbCommon.User, rsp *pbCommon.Groups) error {
	logger.Infof("Received Collaboration.GetUserGroups request: %v", req)
	groups, err := helper.FindStoreEntity(e.store.FindGroupsByUserID, req.UserID, helper.CollaborationServiceID)
	if err != nil {
		return err
	}
	protoGroups := converter.ConvertToTypeArray(groups, converter.StoreGroupToProtoGroupConverter)
	protoRoles := make([]pbCommon.GroupRole, len(protoGroups))
	for index, group := range protoGroups {
		role, err := e.store.FindGroupUserRole(req.UserID, group.GroupID)
		if err != nil {
			return err
		}
		if role != model.RoleRequested && role != model.RoleInvited {
			protoRoles[index] = converter.MigrateModelRoleToProtoRole(role)
		}
	}
	protoGroupsWithUserRole := make([]*pbCommon.Group, len(protoGroups))
	for index := range protoGroups {
		protoGroupsWithUserRole[index] = &pbCommon.Group{
			GroupID:          protoGroups[index].GroupID,
			GroupName:        protoGroups[index].GroupName,
			GroupDescription: protoGroups[index].GroupDescription,
			IsDefault:        protoGroups[index].IsDefault,
			GroupType:        protoGroups[index].GroupType,
			Role:             protoRoles[index],
		}
	}
	rsp.Groups = protoGroupsWithUserRole
	logger.Infof("Found %d groups for user with id %s", len(groups), req.UserID)
	return nil
}

func (e *Collaboration) CreateNewGroupWithAdmin(
	_ context.Context,
	req *pbCommon.GroupRequest,
	rsp *pbCommon.Group,
) error {
	logger.Infof("Received Collaboration.CreateNewGroupWithAdmin request: %v", req)
	err := helper.CheckForValidName(req.Group.GroupName, helper.GroupAndDeckNameRegex, helper.UserServiceID)
	if err != nil {
		return err
	}
	newGroup := model.Group{
		Name:        req.Group.GroupName,
		Description: req.Group.GroupDescription,
		IsDefault:   req.Group.IsDefault,
		GroupType:   model.RequestGroupType,
	}
	err = e.store.CreateNewGroupWithAdmin(req.UserID, &newGroup)
	if err != nil {
		return err
	}
	rsp.GroupID = newGroup.ID
	logger.Infof("Successfully created new group (%s) with user (%s) as admin", newGroup.ID, req.UserID)
	return nil
}

func (e *Collaboration) GetGroup(ctx context.Context, req *pbCommon.GroupRequest, rsp *pbCommon.Group) error {
	logger.Infof("Received Collaboration.GetGroup request: %v", req)
	group, err := e.store.FindGroupByID(req.Group.GroupID)
	if err != nil {
		return err
	}
	protoGroup := converter.StoreGroupToProtoGroupConverter(*group)

	_, err = e.store.FindGroupUserRole(req.UserID, req.Group.GroupID)
	if err != nil {
		if errors.Is(err, helper.ErrStoreNoEntryWithID) {
			logger.Infof("User does not have a group role")
			*rsp = pbCommon.Group{GroupID: protoGroup.GroupID, GroupName: protoGroup.GroupName, GroupDescription: protoGroup.GroupDescription, GroupType: protoGroup.GroupType, Role: pbCommon.GroupRole_GR_EXTERNAL}
			return nil
		}
		return err
	}
	group, protoRole, err := e.checkUserRoleAccessWithGroupAndRoleReturn(
		ctx,
		req.UserID,
		req.Group.GroupID,
		pbCommon.GroupRole_GR_INVITED,
	)
	if err != nil {
		return err
	}
	*rsp = pbCommon.Group{
		GroupID:          protoGroup.GroupID,
		GroupName:        protoGroup.GroupName,
		GroupDescription: protoGroup.GroupDescription,
		IsDefault:        protoGroup.IsDefault,
		GroupType:        protoGroup.GroupType,
		Role:             protoRole,
	}
	logger.Infof("Successfully got information for group %s", req.Group.GroupID)
	return nil
}

func (e *Collaboration) ModifyGroup(ctx context.Context, req *pbCommon.GroupRequest, rsp *pbCommon.Success) error {
	logger.Infof("Received Collaboration.ModifyGroup request: %v", req)
	group, _, err := e.checkUserRoleAccessWithGroupAndRoleReturn(ctx, req.UserID, req.Group.GroupID, pbCommon.GroupRole_GR_ADMIN)
	if err != nil {
		return err
	}
	if group.IsDefault {
		logger.Infof("Cannot modify group %s as it is default group for user %s", req.Group.GroupID, req.UserID)
		return helper.NewMicroNotAuthorizedErr(helper.CollaborationServiceID)
	}
	if req.Group.GroupName != "" {
		err = helper.CheckForValidName(req.Group.GroupName, helper.GroupAndDeckNameRegex, helper.UserServiceID)
		if err != nil {
			return err
		}
		group.Name = req.Group.GroupName
	}
	if req.Group.GroupDescription != "" {
		group.Description = req.Group.GroupDescription
	}
	if req.Group.GroupType != pbCommon.GroupType_GT_INVALID {
		group.GroupType = converter.MigrateProtoGroupTypeToModelGroupType(req.Group.GroupType)
	}
	err = e.store.ModifyGroup(group)
	if err != nil {
		return err
	}
	rsp.Success = true
	logger.Infof("Successfully modified group (%s)", req.Group.GroupID)
	return nil
}

func (e *Collaboration) DeleteGroup(ctx context.Context, req *pbCommon.GroupRequest, rsp *pbCommon.Success) error {
	logger.Infof("Received Collaboration.DeleteGroup request: %v", req)
	group, _, err := e.checkUserRoleAccessWithGroupAndRoleReturn(ctx, req.UserID, req.Group.GroupID, pbCommon.GroupRole_GR_ADMIN)
	if err != nil {
		return err
	}
	if group.IsDefault {
		logger.Infof("Cannot delete group %s as it is default group for user %s", req.Group.GroupID, req.UserID)
		return helper.NewMicroNotAuthorizedErr(helper.CollaborationServiceID)
	}
	err = e.store.DeleteGroup(group)
	if err != nil {
		return err
	}
	rsp.Success = true
	logger.Infof("Successfully deleted group (%s)", req.Group.GroupID)
	return nil
}

func (e *Collaboration) GetGroupMembers(ctx context.Context, req *pbCommon.GroupRequest, rsp *pbCommon.GroupMembers) error {
	logger.Infof("Received Collaboration.GetGroupMembers request: %v", req)
	if _, _, err := e.checkUserRoleAccessWithGroupAndRoleReturn(
		ctx,
		req.UserID,
		req.Group.GroupID,
		pbCommon.GroupRole_GR_INVITED,
	); err != nil {
		return err
	}
	groupMembers, err := helper.FindStoreEntity(e.store.FindGroupMemberRoles, req.Group.GroupID, helper.CollaborationServiceID)
	if err != nil {
		return err
	}
	logger.Infof("Found %d member roles in group with id %s", len(groupMembers), req.Group.GroupID)
	userIDs := converter.ConvertToTypeArray(groupMembers, converter.StoreGroupUserRoleToProtoUserIDConverter)
	logger.Infof("Requesting information of users in group from user service")
	users, err := e.userService.GetUserInformation(ctx, &pbCommon.Users{Users: userIDs})
	if err != nil {
		return err
	}
	rsp.Users = make([]*pbCommon.User, len(users.Users))
	for i, user := range users.Users {
		rsp.Users[i] = &pbCommon.User{
			UserID:    user.UserID,
			UserName:  user.UserName,
			GroupRole: converter.MigrateModelRoleToProtoRole(groupMembers[i].RoleType),
		}
	}
	logger.Infof("Found %d users in group with id %s", len(rsp.Users), req.Group.GroupID)
	return nil
}

func (e *Collaboration) GetGroupMemberRequests(
	ctx context.Context,
	req *pbCommon.GroupRequest,
	rsp *pbCommon.GroupMembers,
) error {
	logger.Infof("Received Collaboration.GetGroupMemberRequests request: %v", req)
	if _, _, err := e.checkUserRoleAccessWithGroupAndRoleReturn(
		ctx,
		req.UserID,
		req.Group.GroupID,
		pbCommon.GroupRole_GR_ADMIN,
	); err != nil {
		return err
	}
	groupRequests, err := e.store.FindGroupRequestsByGroupID(req.Group.GroupID)
	if err != nil && !errors.Is(err, helper.ErrStoreNoEntryWithID) {
		return err
	}
	logger.Infof("Found %d requests for group with id %s", len(groupRequests), req.Group.GroupID)
	rsp.Users, err = e.generateGroupMemberAdmissionResponse(ctx, groupRequests)
	if err != nil {
		return err
	}
	return nil
}

func (e *Collaboration) AddGroupUserRequest(
	_ context.Context,
	req *pbCommon.GroupRequest,
	rsp *pbCommon.Success,
) error {
	logger.Infof("Received Collaboration.AddGroupUserRequest request: %v", req)

	groupRole, err := e.store.FindGroupUserRole(req.UserID, req.Group.GroupID)
	if err != nil {
		if !errors.Is(err, helper.ErrStoreNoEntryWithID) {
			return err
		}
		group, err := helper.FindStoreEntity(e.store.FindGroupByID, req.Group.GroupID, helper.CollaborationServiceID)
		if err != nil {
			return err
		}
		if group.GroupType == model.ClosedGroupType {
			return helper.NewMicroNotAuthorizedErr(helper.CollaborationServiceID)
		}
		if group.GroupType == model.OpenGroupType {
			err := e.store.AddNewMemberToGroup(req.UserID, req.Group.GroupID)
			rsp.Success = true
			return err
		}
		if err = e.store.AddRequestingUserToGroup(req.UserID, req.Group.GroupID); err != nil {
			return err
		}
		rsp.Success = true
		return nil
	}
	if groupRole == model.RoleRequested {
		return helper.NewMicroAlreadyRequestedErr(helper.CollaborationServiceID)
	}
	if groupRole == model.RoleInvited {
		if err = e.store.PromoteUserToFullGroupMember(req.UserID, req.Group.GroupID); err != nil {
			return err
		}
		rsp.Success = true
		return nil
	}
	return helper.NewMicroUserAlreadyInGroupErr(helper.CollaborationServiceID)
}

func (e *Collaboration) ModifyGroupUserRequest(
	_ context.Context,
	req *pbCommon.GroupModUserRequest,
	rsp *pbCommon.Success,
) error {
	logger.Infof("Received Collaboration.ModifyGroupUserRequest request: %v", req)

	role, err := e.store.FindGroupUserRole(req.UserID, req.Group.GroupID)
	if err != nil {
		return err
	}
	if role != model.RoleAdmin {
		return helper.NewMicroNotAuthorizedErr(helper.CollaborationServiceID)
	}
	modUserCurrRole, err := e.store.FindGroupUserRole(req.ModUserID, req.Group.GroupID)
	if err != nil {
		return err
	}
	if modUserCurrRole == model.RoleRequested || modUserCurrRole == model.RoleInvited {
		return helper.NewMicroUserAdmissionInProgressErr(helper.CollaborationServiceID)
	}
	if modUserCurrRole == model.RoleAdmin {
		return helper.NewMicroNotAuthorizedErr(helper.CollaborationServiceID)
	}
	newModelRole := converter.MigrateProtoRoleToModelRole(req.Group.Role)
	switch newModelRole {
	case model.RoleAdmin, model.RoleWrite, model.RoleRead:
		if err := e.store.ModifyUserRole(req.ModUserID, req.Group.GroupID, newModelRole); err != nil {
			return err
		}
	default:
		return helper.NewMicroNotAuthorizedErr(helper.CollaborationServiceID)
	}
	rsp.Success = true
	return nil
}

func (e *Collaboration) KickGroupUser(_ context.Context, req *pbCommon.GroupModUserRequest, rsp *pbCommon.Success) error {
	logger.Infof("Received Collaboration.KickGroupUserRequest request: %v", req)

	if req.UserID == req.ModUserID {
		return helper.NewMicroNotAuthorizedErr(helper.CollaborationServiceID)
	}
	role, err := e.store.FindGroupUserRole(req.UserID, req.Group.GroupID)
	if err != nil {
		return err
	}
	if role != model.RoleAdmin {
		return helper.NewMicroNotAuthorizedErr(helper.CollaborationServiceID)
	}
	delUserCurrRole, err := e.store.FindGroupUserRole(req.ModUserID, req.Group.GroupID)
	if err != nil {
		return err
	}
	if delUserCurrRole == model.RoleRequested || delUserCurrRole == model.RoleInvited {
		return helper.NewMicroUserAdmissionInProgressErr(helper.CollaborationServiceID)
	}
	if delUserCurrRole == model.RoleAdmin {
		return helper.NewMicroNotAuthorizedErr(helper.CollaborationServiceID)
	}
	if err := e.store.RemoveUserFromGroup(req.ModUserID, req.Group.GroupID); err != nil {
		return err
	}

	rsp.Success = true
	return nil
}

func (e *Collaboration) RemoveGroupUserRequest(
	_ context.Context,
	req *pbCommon.GroupRequest,
	rsp *pbCommon.Success,
) error {
	logger.Infof("Received Collaboration.RemoveGroupUserRequest request: %v", req)

	groupRole, err := e.store.FindGroupUserRole(req.UserID, req.Group.GroupID)
	if err != nil {
		if errors.Is(err, helper.ErrStoreNoEntryWithID) {
			return helper.NewMicroNoEntryWithIDErr(helper.CollaborationServiceID)
		}
		return err
	}
	if groupRole == model.RoleRequested || groupRole == model.RoleInvited {
		if err := e.store.RemoveUserFromGroup(req.UserID, req.Group.GroupID); err != nil {
			return err
		}
		rsp.Success = true
		return nil
	} else {
		return helper.NewMicroUserAlreadyInGroupErr(helper.CollaborationServiceID)
	}
}

func (e *Collaboration) AddGroupUserInvite(
	ctx context.Context,
	req *pbCommon.GroupInviteRequest,
	rsp *pbCommon.Success,
) error {
	logger.Infof("Received Collaboration.AddGroupUserInvite request: %v", req)

	requestingUserRole, err := e.store.FindGroupUserRole(req.UserID, req.Group.GroupID)
	if err != nil {
		return helper.NewMicroNotAuthorizedErr(helper.CollaborationServiceID)
	}
	if requestingUserRole != model.RoleAdmin {
		return helper.NewMicroNotAuthorizedErr(helper.CollaborationServiceID)
	}

	userRsp, err := e.userService.GetUserIDFromEmail(ctx, &pbCommon.User{UserEmail: req.InviteUserEmail})
	if err != nil {
		return err
	}
	invitedUserRole, err := e.store.FindGroupUserRole(userRsp.UserID, req.Group.GroupID)
	if err != nil {
		if errors.Is(err, helper.ErrStoreNoEntryWithID) {
			if err := e.store.AddInvitedUserToGroup(userRsp.UserID, req.Group.GroupID); err != nil {
				return err
			}
			rsp.Success = true
			return nil
		}
		return err
	}
	if invitedUserRole == model.RoleInvited {
		return helper.NewMicroAlreadyInvitedErr(helper.CollaborationServiceID)
	} else if invitedUserRole == model.RoleRequested {
		if err := e.store.PromoteUserToFullGroupMember(userRsp.UserID, req.Group.GroupID); err != nil {
			return err
		}
		rsp.Success = true
		return nil
	} else {
		return helper.NewMicroUserAlreadyInGroupErr(helper.CollaborationServiceID)
	}
}

func (e *Collaboration) RemoveGroupUserInvite(
	ctx context.Context,
	req *pbCommon.GroupInviteRequest,
	rsp *pbCommon.Success,
) error {
	logger.Infof("Received Collaboration.RemoveGroupUserInvite request: %v", req)

	requestingUserRole, err := e.store.FindGroupUserRole(req.UserID, req.Group.GroupID)
	if err != nil {
		return helper.NewMicroNotAuthorizedErr(helper.CollaborationServiceID)
	}
	if requestingUserRole != model.RoleAdmin {
		return helper.NewMicroNotAuthorizedErr(helper.CollaborationServiceID)
	}
	userRsp, err := e.userService.GetUserIDFromEmail(ctx, &pbCommon.User{UserEmail: req.InviteUserEmail})
	if err != nil {
		return err
	}
	invitedUserRole, err := e.store.FindGroupUserRole(userRsp.UserID, req.Group.GroupID)
	if err != nil {
		if errors.Is(err, helper.ErrStoreNoEntryWithID) {
			return helper.NewMicroNoEntryWithIDErr(helper.CollaborationServiceID)
		}
		return err
	}
	if invitedUserRole == model.RoleInvited || invitedUserRole == model.RoleRequested {
		if err := e.store.RemoveUserFromGroup(userRsp.UserID, req.Group.GroupID); err != nil {
			return err
		}
		rsp.Success = true
		return nil
	} else {
		return helper.NewMicroUserAlreadyInGroupErr(helper.CollaborationServiceID)
	}
}

func (e *Collaboration) GetInvitationsForGroup(
	ctx context.Context,
	req *pbCommon.GroupRequest,
	rsp *pbCommon.GroupMembers,
) error {
	logger.Infof("Received Collaboration.GetInvitationsForGroup request: %v", req)
	if _, _, err := e.checkUserRoleAccessWithGroupAndRoleReturn(
		ctx,
		req.UserID,
		req.Group.GroupID,
		pbCommon.GroupRole_GR_ADMIN,
	); err != nil {
		return err
	}
	memberInvitations, err := e.store.FindGroupInvitationsByGroupID(req.Group.GroupID)
	if err != nil {
		return err
	}
	rsp.Users, err = e.generateGroupMemberAdmissionResponse(ctx, memberInvitations)
	if err != nil {
		return err
	}
	return nil
}

func (e *Collaboration) GetGroupUserRole(_ context.Context, req *pbCommon.GroupRequest, rsp *pbCommon.Group) error {
	logger.Infof("Received Collaboration.GetUserGroupRole request: %v", req)
	_, err := helper.FindStoreEntity(e.store.FindGroupByID, req.Group.GroupID, helper.CollaborationServiceID)
	if err != nil {
		return err
	}
	role, err := e.store.FindGroupUserRole(req.UserID, req.Group.GroupID)
	if err != nil {
		if errors.Is(err, helper.ErrStoreNoEntryWithID) {
			return helper.NewMicroNoEntryWithIDErr(helper.CollaborationServiceID)
		}
		return err
	}
	logger.Infof("Found group with id %s by obtaining role", req.Group.GroupID)
	rsp.GroupID = req.Group.GroupID
	protoRole := converter.MigrateModelRoleToProtoRole(role)
	rsp.Role = protoRole
	logger.Infof("Obtained role (%s) for group (%s) for user (%s)",
		rsp.Role.String(), req.Group.GroupID, req.UserID)
	return nil
}

func (e *Collaboration) FindGroupByID(_ context.Context, req *pbCommon.GroupRequest, rsp *pbCommon.Group) error {
	logger.Infof("Received Collaboration.FindGroupByID request: %v", req)
	group, err := helper.FindStoreEntity(e.store.FindGroupByID, req.Group.GroupID, helper.CollaborationServiceID)
	if err != nil {
		return err
	}
	*rsp = *converter.StoreGroupToProtoGroupConverter(*group)
	logger.Infof("Successfully found group with id %s", req.Group.GroupID)
	return nil
}

func (e *Collaboration) LeaveGroupSafe(ctx context.Context, req *pbCommon.GroupRequest, rsp *pbCommon.Success) error {
	logger.Infof("Received Collaboration.LeaveGroupSafe request: %+v", req)
	remainingAdmins, err := e.store.FindGroupAdmins(req.Group.GroupID)
	if err != nil {
		return err
	}
	if len(remainingAdmins) == 1 && remainingAdmins[0].UserID == req.UserID {
		return helper.NewMicroCantLeaveAsLastAdminErr(helper.CollaborationServiceID)
	}
	return e.LeaveGroup(ctx, req, rsp)
}

func (e *Collaboration) LeaveGroup(_ context.Context, req *pbCommon.GroupRequest, rsp *pbCommon.Success) error {
	logger.Infof("Received Collaboration.LeaveGroup request: %+v", req)
	group, err := e.store.FindGroupByID(req.Group.GroupID)
	if err != nil {
		return err
	}
	groupUsers, err := helper.FindStoreEntity(e.store.FindGroupMemberRoles, req.Group.GroupID, helper.CollaborationServiceID)
	if err != nil {
		return err
	}
	if len(groupUsers) > 1 || group.GroupType == model.OpenGroupType {
		if err = e.store.RemoveUserFromGroup(req.UserID, req.Group.GroupID); err != nil {
			return err
		}
	} else {
		if err = e.store.DeleteGroup(group); err != nil {
			return err
		}
	}
	rsp.Success = true
	logger.Infof("User %s left group %s.", req.UserID, req.Group.GroupID)
	return nil
}
