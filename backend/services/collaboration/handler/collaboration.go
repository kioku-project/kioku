package handler

import (
	"context"
	"errors"
	"github.com/kioku-project/kioku/pkg/converter"
	"github.com/kioku-project/kioku/pkg/helper"

	"go-micro.dev/v4/logger"

	"github.com/kioku-project/kioku/pkg/model"
	pb "github.com/kioku-project/kioku/services/collaboration/proto"
	pbUser "github.com/kioku-project/kioku/services/user/proto"
	"github.com/kioku-project/kioku/store"
)

type Collaboration struct {
	store       store.CollaborationStore
	userService pbUser.UserService
}

func New(s store.CollaborationStore, uS pbUser.UserService) *Collaboration {
	return &Collaboration{store: s, userService: uS}
}

func (e *Collaboration) checkForGroupAndAdmission(userID string, groupID string) error {
	if _, err := e.store.GetGroupUserRole(userID, groupID); err == nil {
		return helper.NewMicroUserAlreadyInGroupErr(helper.CollaborationServiceID)
	}
	logger.Infof("User with id %s is not part of group with id %s", userID, groupID)
	if _, err := e.store.FindGroupAdmissionByUserAndGroupID(userID, groupID); err == nil {
		return helper.NewMicroUserAdmissionInProgressErr(helper.CollaborationServiceID)
	}
	logger.Infof("User with id %s has no running admission process for group with id %s", userID, groupID)
	return nil
}

func (e *Collaboration) checkUserRoleAccess(_ context.Context, userID string, groupID string, requiredRole pb.GroupRole) error {
	logger.Infof("Requesting group role for user (%s)", userID)
	role, err := e.store.GetGroupUserRole(userID, groupID)
	if err != nil {
		if errors.Is(err, helper.ErrStoreNoEntryWithID) {
			return helper.NewMicroNoEntryWithIDErr(helper.CollaborationServiceID)
		}
		return err
	}
	protoRole := converter.MigrateModelRoleToProtoRole(role)
	logger.Infof("Obtained group role (%s) for user (%s)", protoRole.String(), userID)
	if !helper.IsAuthorized(protoRole, requiredRole) {
		return helper.NewMicroNotAuthorizedErr(helper.CardDeckServiceID)
	}
	logger.Infof("Authenticated group role (%s) for user (%s)", protoRole.String(), userID)
	return nil
}

func (e *Collaboration) GetGroupInvitations(_ context.Context, req *pb.UserIDRequest, rsp *pb.GroupInvitationsResponse) error {
	logger.Infof("Received Collaboration.GetGroupInvitations request: %v", req)
	groupInvitations, err := e.store.FindGroupInvitationsByUserID(req.UserID)
	if err != nil && !errors.Is(err, helper.ErrStoreNoEntryWithID) {
		return err
	}
	rsp.GroupInvitation = converter.ConvertToTypeArray(groupInvitations, converter.StoreGroupAdmissionToProtoGroupInvitationConverter)
	logger.Infof("Found %d invitations for user with id %s", len(groupInvitations), req.UserID)
	return nil
}

func (e *Collaboration) ManageGroupInvitation(_ context.Context, req *pb.ManageGroupInvitationRequest, rsp *pb.SuccessResponse) error {
	logger.Infof("Received Collaboration.ManageGroupInvitation request: %v", req)
	groupAdmission, err := e.store.FindGroupAdmissionByID(req.AdmissionID)
	if err != nil {
		return err
	}
	logger.Infof("Found group admission with id %s", groupAdmission.ID)
	if groupAdmission.UserID != req.UserID || groupAdmission.AdmissionStatus != model.Invited {
		return helper.NewMicroInvalidParameterDataErr(helper.CollaborationServiceID)
	}
	logPrefix := "rejected"
	if req.RequestResponse {
		if err = e.store.AddNewMemberToGroup(groupAdmission.UserID, groupAdmission.GroupID); err != nil {
			return err
		}
		logPrefix = "accepted"
	}
	logger.Infof("User %s %s request to join group with id %s", groupAdmission.UserID, logPrefix, groupAdmission.GroupID)
	if err = e.store.DeleteGroupAdmission(groupAdmission); err != nil {
		return err
	}
	logger.Infof("Deleted group admission with id %s", groupAdmission.ID)
	rsp.Success = true
	logger.Infof("Successfully handled invitation for user %s to join group %s", groupAdmission.UserID, groupAdmission.GroupID)
	return nil
}

func (e *Collaboration) GetUserGroups(_ context.Context, req *pb.UserIDRequest, rsp *pb.UserGroupsResponse) error {
	logger.Infof("Received Collaboration.GetUserGroups request: %v", req)
	groups, err := helper.FindStoreEntity(e.store.FindGroupsByUserID, req.UserID, helper.CollaborationServiceID)
	if err != nil {
		return err
	}
	rsp.Groups = converter.ConvertToTypeArray(groups, converter.StoreGroupToProtoGroupConverter)
	logger.Infof("Found %d groups for user with id %s", len(groups), req.UserID)
	return nil
}

func (e *Collaboration) CreateNewGroupWithAdmin(_ context.Context, req *pb.CreateGroupRequest, rsp *pb.IDResponse) error {
	logger.Infof("Received Collaboration.CreateNewGroupWithAdmin request: %v", req)
	err := helper.CheckForValidName(req.GroupName, helper.GroupAndDeckNameRegex, helper.UserServiceID)
	if err != nil {
		return err
	}
	newGroup := model.Group{
		Name:      req.GroupName,
		IsDefault: req.IsDefault,
		GroupType: model.Private,
	}
	err = e.store.CreateNewGroupWithAdmin(req.UserID, &newGroup)
	if err != nil {
		return err
	}
	rsp.ID = newGroup.ID
	logger.Infof("Successfully created new group (%s) with user (%s) as admin", newGroup.ID, req.UserID)
	return nil
}

func (e *Collaboration) ModifyGroup(ctx context.Context, req *pb.ModifyGroupRequest, rsp *pb.SuccessResponse) error {
	logger.Infof("Received Collaboration.ModifyGroup request: %v", req)
	if err := e.checkUserRoleAccess(ctx, req.UserID, req.GroupID, pb.GroupRole_ADMIN); err != nil {
		return err
	}
	group, err := helper.FindStoreEntity(e.store.FindGroupByID, req.GroupID, helper.CollaborationServiceID)
	if err != nil {
		return err
	}
	if group.IsDefault {
		logger.Infof("Cannot modify group %s as it is default group for user %s", req.GroupID, req.UserID)
		return helper.NewMicroNotAuthorizedErr(helper.CollaborationServiceID)
	}
	if req.GroupName != nil {
		err = helper.CheckForValidName(*req.GroupName, helper.GroupAndDeckNameRegex, helper.UserServiceID)
		if err != nil {
			return err
		}
		group.Name = *req.GroupName
	}
	if req.GroupType != nil && *req.GroupType != pb.GroupType_INVALID {
		group.GroupType = converter.MigrateProtoGroupTypeToModelGroupType(*req.GroupType)
	}
	err = e.store.ModifyGroup(group)
	if err != nil {
		return err
	}
	rsp.Success = true
	logger.Infof("Successfully modified group (%s)", req.GroupID)
	return nil
}

func (e *Collaboration) DeleteGroup(ctx context.Context, req *pb.GroupRequest, rsp *pb.SuccessResponse) error {
	logger.Infof("Received Collaboration.DeleteGroup request: %v", req)
	if err := e.checkUserRoleAccess(ctx, req.UserID, req.GroupID, pb.GroupRole_ADMIN); err != nil {
		return err
	}
	group, err := helper.FindStoreEntity(e.store.FindGroupByID, req.GroupID, helper.CollaborationServiceID)
	if err != nil {
		return err
	}
	if group.IsDefault {
		logger.Infof("Cannot delete group %s as it is default group for user %s", req.GroupID, req.UserID)
		return helper.NewMicroNotAuthorizedErr(helper.CollaborationServiceID)
	}
	err = e.store.DeleteGroup(group)
	if err != nil {
		return err
	}
	rsp.Success = true
	logger.Infof("Successfully deleted group (%s)", req.GroupID)
	return nil
}

func (e *Collaboration) GetGroupMembers(ctx context.Context, req *pb.GroupRequest, rsp *pb.GroupMembersResponse) error {
	logger.Infof("Received Collaboration.GetGroupMembers request: %v", req)
	if err := e.checkUserRoleAccess(ctx, req.UserID, req.GroupID, pb.GroupRole_READ); err != nil {
		return err
	}
	groupMembers, err := helper.FindStoreEntity(e.store.GetGroupMemberRoles, req.GroupID, helper.CollaborationServiceID)
	if err != nil {
		return err
	}
	logger.Infof("Found %d member roles in group with id %s", len(groupMembers), req.GroupID)
	userIDs := converter.ConvertToTypeArray(groupMembers, converter.StoreGroupUserRoleToProtoUserIDConverter)
	logger.Infof("Requesting information of users in group from user service")
	users, err := e.userService.GetUserInformation(ctx, &pbUser.UserInformationRequest{UserIDs: userIDs})
	if err != nil {
		return err
	}
	rsp.Users = make([]*pb.UserWithRole, len(users.Users))
	for i, user := range users.Users {
		rsp.Users[i] = &pb.UserWithRole{
			User: &pb.User{
				UserID: user.UserID,
				Name:   user.Name,
			},
			GroupRole: converter.MigrateModelRoleToProtoRole(groupMembers[i].RoleType),
		}
	}
	logger.Infof("Found %d users in group with id %s", len(rsp.Users), req.GroupID)
	return nil
}

func (e *Collaboration) GetGroupMemberRequests(ctx context.Context, req *pb.GroupRequest, rsp *pb.GroupMemberRequestsResponse) error {
	logger.Infof("Received Collaboration.GetGroupMemberRequests request: %v", req)
	if err := e.checkUserRoleAccess(ctx, req.UserID, req.GroupID, pb.GroupRole_ADMIN); err != nil {
		return err
	}
	groupRequests, err := e.store.FindGroupRequestsByGroupID(req.GroupID)
	if err != nil && !errors.Is(err, helper.ErrStoreNoEntryWithID) {
		return err
	}
	logger.Infof("Found %d requests for group with id %s", len(groupRequests), req.GroupID)
	userIDs := converter.ConvertToTypeArray(groupRequests, converter.StoreGroupAdmissionToProtoUserIDConverter)
	logger.Infof("Requesting information of users in group from user service")
	users, err := e.userService.GetUserInformation(ctx, &pbUser.UserInformationRequest{UserIDs: userIDs})
	if err != nil {
		return err
	}
	rsp.MemberRequests = make([]*pb.MemberRequest, len(users.Users))
	for i, user := range users.Users {
		rsp.MemberRequests[i] = &pb.MemberRequest{
			AdmissionID: groupRequests[i].ID,
			User: &pb.User{
				UserID: user.UserID,
				Name:   user.Name,
			},
		}
	}
	logger.Infof("Successfully received user information from %d users and added it to request information", len(users.Users))
	return nil
}

func (e *Collaboration) ManageGroupMemberRequest(ctx context.Context, req *pb.ManageGroupMemberRequestRequest, rsp *pb.SuccessResponse) error {
	logger.Infof("Received Collaboration.ManageGroupMemberRequest request: %v", req)
	if err := e.checkUserRoleAccess(ctx, req.UserID, req.GroupID, pb.GroupRole_ADMIN); err != nil {
		return err
	}
	groupAdmission, err := e.store.FindGroupAdmissionByID(req.AdmissionID)
	if err != nil {
		return err
	}
	logger.Infof("Found group admission with id %s", groupAdmission.ID)
	if groupAdmission.GroupID != req.GroupID || groupAdmission.AdmissionStatus != model.Requested {
		return helper.NewMicroInvalidParameterDataErr(helper.CollaborationServiceID)
	}
	logPrefix := "Rejected"
	if req.RequestResponse {
		if err = e.store.AddNewMemberToGroup(groupAdmission.UserID, groupAdmission.GroupID); err != nil {
			return err
		}
		logPrefix = "Accepted"
	}
	logger.Infof("%s request of user %s for group %s", logPrefix, groupAdmission.UserID, groupAdmission.GroupID)
	if err = e.store.DeleteGroupAdmission(groupAdmission); err != nil {
		return err
	}
	logger.Infof("Deleted group admission with id %s", groupAdmission.ID)
	rsp.Success = true
	logger.Infof("Successfully handled member request of user %s for group %s", groupAdmission.UserID, groupAdmission.GroupID)
	return nil
}

func (e *Collaboration) RequestToJoinGroup(_ context.Context, req *pb.GroupRequest, rsp *pb.SuccessResponse) error {
	logger.Infof("Received Collaboration.RequestToJoinGroup request: %v", req)
	group, err := helper.FindStoreEntity(e.store.FindGroupByID, req.GroupID, helper.CollaborationServiceID)
	if err != nil {
		return err
	}
	logger.Infof("Found group with id %s", group.ID)
	if group.GroupType == model.Private {
		logger.Infof("Group with id %s is private", group.ID)
		return helper.NewMicroNotAuthorizedErr(helper.CollaborationServiceID)
	}
	logger.Infof("Group with id %s is public", group.ID)
	err = e.checkForGroupAndAdmission(req.UserID, req.GroupID)
	if err != nil {
		return err
	}
	newAdmission := model.GroupAdmission{
		UserID:          req.UserID,
		GroupID:         req.GroupID,
		AdmissionStatus: model.Requested,
	}
	err = e.store.CreateNewGroupAdmission(&newAdmission)
	if err != nil {
		return err
	}
	rsp.Success = true
	logger.Infof("Successfully requested to join group %s as user with id %s", req.GroupID, req.UserID)
	return nil
}

func (e *Collaboration) InviteUserToGroup(ctx context.Context, req *pb.GroupInvitationRequest, rsp *pb.SuccessResponse) error {
	logger.Infof("Received Collaboration.InviteUserToGroup request: %v", req)
	if err := e.checkUserRoleAccess(ctx, req.UserID, req.GroupID, pb.GroupRole_ADMIN); err != nil {
		return err
	}
	logger.Infof("Requesting user id from invited user from user service by email %s", req.InvitedUserEmail)
	userRsp, err := e.userService.GetUserIDFromEmail(ctx, &pbUser.UserIDRequest{Email: req.InvitedUserEmail})
	if err != nil {
		return err
	}
	logger.Infof("Got user id %s for email %s", userRsp.UserID, req.InvitedUserEmail)
	if err = e.checkForGroupAndAdmission(userRsp.UserID, req.GroupID); err != nil {
		return err
	}
	newAdmission := model.GroupAdmission{
		UserID:          userRsp.UserID,
		GroupID:         req.GroupID,
		AdmissionStatus: model.Invited,
	}
	if err = e.store.CreateNewGroupAdmission(&newAdmission); err != nil {
		return err
	}
	rsp.Success = true
	logger.Infof("Successfully invited user %s to group %s", userRsp.UserID, req.GroupID)
	return nil
}

func (e *Collaboration) GetGroupUserRole(_ context.Context, req *pb.GroupRequest, rsp *pb.GroupRoleResponse) error {
	logger.Infof("Received Collaboration.GetUserGroupRole request: %v", req)
	role, err := e.store.GetGroupUserRole(req.UserID, req.GroupID)
	if err != nil {
		if errors.Is(err, helper.ErrStoreNoEntryWithID) {
			return helper.NewMicroNoEntryWithIDErr(helper.CollaborationServiceID)
		}
		return err
	}
	logger.Infof("Found group with id %s by obtaining role", req.GroupID)
	rsp.GroupID = req.GroupID
	protoRole := converter.MigrateModelRoleToProtoRole(role)
	rsp.GroupRole = protoRole
	logger.Infof("Obtained role (%s) for group (%s) for user (%s)", rsp.GroupRole.String(), req.GroupID, req.UserID)
	return nil
}

func (e *Collaboration) FindGroupByID(_ context.Context, req *pb.GroupRequest, rsp *pb.GroupResponse) error {
	logger.Infof("Received Collaboration.FindGroupByID request: %v", req)
	group, err := helper.FindStoreEntity(e.store.FindGroupByID, req.GroupID, helper.CollaborationServiceID)
	if err != nil {
		return err
	}
	rsp.GroupID = group.ID
	rsp.GroupName = group.Name
	return nil
}
