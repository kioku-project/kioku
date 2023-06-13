package handler

import (
	"context"
	"errors"
	"github.com/kioku-project/kioku/pkg/converter"
	"github.com/kioku-project/kioku/pkg/helper"
	pbCardDeck "github.com/kioku-project/kioku/services/carddeck/proto"
	pbSrs "github.com/kioku-project/kioku/services/srs/proto"
	"go-micro.dev/v4/logger"

	"github.com/kioku-project/kioku/pkg/model"
	pb "github.com/kioku-project/kioku/services/collaboration/proto"
	pbUser "github.com/kioku-project/kioku/services/user/proto"
	"github.com/kioku-project/kioku/store"
)

type Collaboration struct {
	store           store.CollaborationStore
	userService     pbUser.UserService
	srsService      pbSrs.SrsService
	cardDeckService pbCardDeck.CardDeckService
}

func New(s store.CollaborationStore, uS pbUser.UserService, srsS pbSrs.SrsService, cdS pbCardDeck.CardDeckService) *Collaboration {
	return &Collaboration{store: s, userService: uS, srsService: srsS, cardDeckService: cdS}
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

func (e *Collaboration) checkUserRoleAccessWithGroupAndRoleReturn(_ context.Context, userID string, groupID string, requiredRole pb.GroupRole) (group *model.Group, protoRole pb.GroupRole, err error) {
	logger.Infof("Find group with id %s", groupID)
	group, err = helper.FindStoreEntity(e.store.FindGroupByID, groupID, helper.CollaborationServiceID)
	if err != nil {
		return
	}
	logger.Infof("Requesting group role for user (%s)", userID)
	role, err := e.store.GetGroupUserRole(userID, groupID)
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

func (e *Collaboration) generateGroupMemberAdmissionResponse(ctx context.Context, groupAdmissions []model.GroupAdmission) (memberAdmissions []*pb.MemberAdmission, err error) {
	userIDs := converter.ConvertToTypeArray(groupAdmissions, converter.StoreGroupAdmissionToProtoUserIDConverter)
	logger.Infof("Requesting information of users in group from user service")
	users, err := e.userService.GetUserInformation(ctx, &pbUser.UserInformationRequest{UserIDs: userIDs})
	if err != nil {
		return
	}
	memberAdmissions = make([]*pb.MemberAdmission, len(users.Users))
	for i, user := range users.Users {
		memberAdmissions[i] = &pb.MemberAdmission{
			AdmissionID: groupAdmissions[i].ID,
			User: &pb.User{
				UserID: user.UserID,
				Name:   user.UserName,
			},
		}
	}
	logger.Infof("Successfully received user information from %d users and added it to request information", len(users.Users))
	return
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

func (e *Collaboration) ManageGroupInvitation(ctx context.Context, req *pb.ManageGroupInvitationRequest, rsp *pb.SuccessResponse) error {
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
		if err = e.store.ChangeInvitedUserToFullGroupMember(groupAdmission.UserID, groupAdmission.GroupID); err != nil {
			return err
		}
		logPrefix = "accepted"
	} else {
		if err = e.store.RemoveUserFromGroup(groupAdmission.UserID, groupAdmission.GroupID); err != nil {
			return err
		}
	}
	logger.Infof("User %s %s request to join group with id %s", groupAdmission.UserID, logPrefix, groupAdmission.GroupID)
	if err = e.store.DeleteGroupAdmission(groupAdmission); err != nil {
		return err
	}
	logger.Infof("Deleted group admission with id %s", groupAdmission.ID)

	// add cardbindings for user
	if err := e.createUserCardBindingsForWholeGroup(ctx, groupAdmission.UserID, groupAdmission.GroupID); err != nil {
		return err
	}
	rsp.Success = true
	logger.Infof("Successfully handled invitation for user %s to join group %s", groupAdmission.UserID, groupAdmission.GroupID)
	return nil
}

func (e *Collaboration) createUserCardBindingsForWholeGroup(ctx context.Context, userID string, groupID string) error {
	decks, err := e.cardDeckService.GetGroupDecks(ctx, &pbCardDeck.GroupDecksRequest{
		UserID:  userID,
		GroupID: groupID,
	})
	if err != nil {
		return err
	}
	for _, deck := range decks.Decks {
		cards, err := e.cardDeckService.GetDeckCards(ctx, &pbCardDeck.IDRequest{
			UserID:   userID,
			EntityID: deck.DeckID,
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

func (e *Collaboration) GetUserGroups(_ context.Context, req *pb.UserIDRequest, rsp *pb.UserGroupsResponse) error {
	logger.Infof("Received Collaboration.GetUserGroups request: %v", req)
	groups, err := helper.FindStoreEntity(e.store.FindGroupsByUserID, req.UserID, helper.CollaborationServiceID)
	if err != nil {
		return err
	}
	protoGroups := converter.ConvertToTypeArray(groups, converter.StoreGroupToProtoGroupConverter)
	protoRoles := make([]pb.GroupRole, len(protoGroups))
	for index, group := range protoGroups {
		role, err := e.store.GetGroupUserRole(req.UserID, group.GroupID)
		if err != nil {
			return err
		}
		protoRoles[index] = converter.MigrateModelRoleToProtoRole(role)
	}
	protoGroupsWithUserRole := make([]*pb.GroupWithUserRole, len(protoGroups))
	for index := range protoGroups {
		protoGroupsWithUserRole[index] = &pb.GroupWithUserRole{
			Group: protoGroups[index],
			Role:  protoRoles[index],
		}
	}
	rsp.Groups = protoGroupsWithUserRole
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
		Name:        req.GroupName,
		Description: req.GroupDescription,
		IsDefault:   req.IsDefault,
		GroupType:   model.Private,
	}
	err = e.store.CreateNewGroupWithAdmin(req.UserID, &newGroup)
	if err != nil {
		return err
	}
	rsp.ID = newGroup.ID
	logger.Infof("Successfully created new group (%s) with user (%s) as admin", newGroup.ID, req.UserID)
	return nil
}

func (e *Collaboration) GetGroup(ctx context.Context, req *pb.GroupRequest, rsp *pb.GroupWithUserRole) error {
	logger.Infof("Received Collaboration.GetGroup request: %v", req)
	group, protoRole, err := e.checkUserRoleAccessWithGroupAndRoleReturn(ctx, req.UserID, req.GroupID, pb.GroupRole_INVITED)
	if err != nil {
		return err
	}
	protoGroup := converter.StoreGroupToProtoGroupConverter(*group)
	*rsp = pb.GroupWithUserRole{
		Group: protoGroup,
		Role:  protoRole,
	}
	logger.Infof("Successfully got information for group %s", req.GroupID)
	return nil
}

func (e *Collaboration) ModifyGroup(ctx context.Context, req *pb.ModifyGroupRequest, rsp *pb.SuccessResponse) error {
	logger.Infof("Received Collaboration.ModifyGroup request: %v", req)
	group, _, err := e.checkUserRoleAccessWithGroupAndRoleReturn(ctx, req.UserID, req.GroupID, pb.GroupRole_ADMIN)
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
	if req.GroupDescription != nil {
		group.Description = *req.GroupDescription
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
	group, _, err := e.checkUserRoleAccessWithGroupAndRoleReturn(ctx, req.UserID, req.GroupID, pb.GroupRole_ADMIN)
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
	if _, _, err := e.checkUserRoleAccessWithGroupAndRoleReturn(ctx, req.UserID, req.GroupID, pb.GroupRole_INVITED); err != nil {
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
				Name:   user.UserName,
			},
			GroupRole: converter.MigrateModelRoleToProtoRole(groupMembers[i].RoleType),
		}
	}
	logger.Infof("Found %d users in group with id %s", len(rsp.Users), req.GroupID)
	return nil
}

func (e *Collaboration) GetGroupMemberRequests(ctx context.Context, req *pb.GroupRequest, rsp *pb.GroupMemberAdmissionResponse) error {
	logger.Infof("Received Collaboration.GetGroupMemberRequests request: %v", req)
	if _, _, err := e.checkUserRoleAccessWithGroupAndRoleReturn(ctx, req.UserID, req.GroupID, pb.GroupRole_ADMIN); err != nil {
		return err
	}
	groupRequests, err := e.store.FindGroupRequestsByGroupID(req.GroupID)
	if err != nil && !errors.Is(err, helper.ErrStoreNoEntryWithID) {
		return err
	}
	logger.Infof("Found %d requests for group with id %s", len(groupRequests), req.GroupID)
	rsp.MemberAdmissions, err = e.generateGroupMemberAdmissionResponse(ctx, groupRequests)
	if err != nil {
		return err
	}
	return nil
}

func (e *Collaboration) ManageGroupMemberRequest(ctx context.Context, req *pb.ManageGroupMemberRequestRequest, rsp *pb.SuccessResponse) error {
	logger.Infof("Received Collaboration.ManageGroupMemberRequest request: %v", req)
	if _, _, err := e.checkUserRoleAccessWithGroupAndRoleReturn(ctx, req.UserID, req.GroupID, pb.GroupRole_ADMIN); err != nil {
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

	// add cardbindings for user
	if err := e.createUserCardBindingsForWholeGroup(ctx, groupAdmission.UserID, groupAdmission.GroupID); err != nil {
		return err
	}

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

func (e *Collaboration) GetInvitationsForGroup(ctx context.Context, req *pb.GroupRequest, rsp *pb.GroupMemberAdmissionResponse) error {
	logger.Infof("Received Collaboration.GetInvitationsForGroup request: %v", req)
	if _, _, err := e.checkUserRoleAccessWithGroupAndRoleReturn(ctx, req.UserID, req.GroupID, pb.GroupRole_ADMIN); err != nil {
		return err
	}
	memberInvitations, err := e.store.FindGroupInvitationsByGroupID(req.GroupID)
	if err != nil {
		return err
	}
	rsp.MemberAdmissions, err = e.generateGroupMemberAdmissionResponse(ctx, memberInvitations)
	if err != nil {
		return err
	}
	return nil
}

func (e *Collaboration) InviteUserToGroup(ctx context.Context, req *pb.GroupInvitationRequest, rsp *pb.SuccessResponse) error {
	logger.Infof("Received Collaboration.InviteUserToGroup request: %v", req)
	if _, _, err := e.checkUserRoleAccessWithGroupAndRoleReturn(ctx, req.UserID, req.GroupID, pb.GroupRole_ADMIN); err != nil {
		return err
	}
	logger.Infof("Requesting user id from invited user from user service by email %s", req.InvitedUserEmail)
	userRsp, err := e.userService.GetUserIDFromEmail(ctx, &pbUser.UserIDRequest{UserEmail: req.InvitedUserEmail})
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
	if err = e.store.AddInvitedUserToGroup(newAdmission.UserID, newAdmission.GroupID); err != nil {
		return err
	}
	rsp.Success = true
	logger.Infof("Successfully invited user %s to group %s", userRsp.UserID, req.GroupID)
	return nil
}

func (e *Collaboration) GetGroupUserRole(_ context.Context, req *pb.GroupRequest, rsp *pb.GroupRoleResponse) error {
	logger.Infof("Received Collaboration.GetUserGroupRole request: %v", req)
	_, err := helper.FindStoreEntity(e.store.FindGroupByID, req.GroupID, helper.CollaborationServiceID)
	if err != nil {
		return err
	}
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

func (e *Collaboration) FindGroupByID(_ context.Context, req *pb.GroupRequest, rsp *pb.Group) error {
	logger.Infof("Received Collaboration.FindGroupByID request: %v", req)
	group, err := helper.FindStoreEntity(e.store.FindGroupByID, req.GroupID, helper.CollaborationServiceID)
	if err != nil {
		return err
	}
	*rsp = *converter.StoreGroupToProtoGroupConverter(*group)
	logger.Infof("Successfully found group with id %s", req.GroupID)
	return nil
}
