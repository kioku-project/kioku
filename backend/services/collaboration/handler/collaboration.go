package handler

import (
	"context"
	"errors"
	"github.com/kioku-project/kioku/pkg/helper"

	"go-micro.dev/v4/logger"

	"github.com/kioku-project/kioku/pkg/model"
	pb "github.com/kioku-project/kioku/services/collaboration/proto"
	"github.com/kioku-project/kioku/store"
)

type Collaboration struct{ store store.CollaborationStore }

func New(s store.CollaborationStore) *Collaboration { return &Collaboration{store: s} }

func (e *Collaboration) migrateModelRoleToProtoRole(modelRole model.RoleType) (protoRole pb.GroupRole) {
	if modelRole == model.RoleRead {
		protoRole = pb.GroupRole_READ
	} else if modelRole == model.RoleWrite {
		protoRole = pb.GroupRole_WRITE
	} else if modelRole == model.RoleAdmin {
		protoRole = pb.GroupRole_ADMIN
	}
	return
}

func (e *Collaboration) securityRoleHandler(_ context.Context, userID string, groupID string, requiredRole pb.GroupRole) error {
	logger.Infof("Requesting group role for user (%s)", userID)
	role, err := e.store.GetGroupUserRole(userID, groupID)
	if err != nil {
		if errors.Is(err, helper.ErrStoreNoEntryWithID) {
			return helper.ErrMicroNoEntryWithID(helper.CollaborationServiceID)
		}
		return err
	}
	protoRole := e.migrateModelRoleToProtoRole(role)
	logger.Infof("Obtained group role (%s) for user (%s)", protoRole.String(), userID)
	if !helper.IsAuthorized(protoRole, requiredRole) {
		return helper.ErrMicroNotAuthorized(helper.CardDeckServiceID)
	}
	logger.Infof("Authenticated group role (%s) for user (%s)", protoRole.String(), userID)
	return nil
}

func (e *Collaboration) GetUserGroups(_ context.Context, req *pb.UserGroupsRequest, rsp *pb.UserGroupsResponse) error {
	logger.Infof("Received Collaboration.GetUserGroups request: %v", req)
	groups, err := helper.FindEntityWrapper(e.store.FindGroupsByUserID, req.UserID, helper.CollaborationServiceID)
	if err != nil {
		return err
	}
	rsp.Groups = make([]*pb.Group, len(groups))
	for i, group := range groups {
		rsp.Groups[i] = &pb.Group{
			GroupID:   group.ID,
			GroupName: group.Name,
			IsDefault: group.IsDefault,
		}
	}
	logger.Infof("Found %d groups for user with id %s", len(groups), req.UserID)
	return nil
}

func (e *Collaboration) CreateNewGroupWithAdmin(_ context.Context, req *pb.CreateGroupRequest, rsp *pb.IDResponse) error {
	logger.Infof("Received Collaboration.CreateNewGroupWithAdmin request: %v", req)
	newGroup := model.Group{
		Name:      req.GroupName,
		IsDefault: req.IsDefault,
	}
	err := e.store.CreateNewGroupWithAdmin(req.UserID, &newGroup)
	if err != nil {
		return err
	}
	rsp.ID = newGroup.ID
	logger.Infof("Successfully created new group (%s) with user (%s) as admin", newGroup.ID, req.UserID)
	return nil
}

func (e *Collaboration) ModifyGroup(ctx context.Context, req *pb.ModifyGroupRequest, rsp *pb.SuccessResponse) error {
	logger.Infof("Received Collaboration.ModifyGroup request: %v", req)
	if err := e.securityRoleHandler(ctx, req.UserID, req.GroupID, pb.GroupRole_WRITE); err != nil {
		return err
	}
	group, err := helper.FindEntityWrapper(e.store.FindGroupByID, req.GroupID, helper.CollaborationServiceID)
	if err != nil {
		return err
	}
	if group.IsDefault {
		logger.Infof("Cannot modify group %s as it is default group for user %s", req.GroupID, req.UserID)
		return helper.ErrMicroNotAuthorized(helper.CollaborationServiceID)
	}
	if req.GroupName != nil {
		group.Name = *req.GroupName
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
	if err := e.securityRoleHandler(ctx, req.UserID, req.GroupID, pb.GroupRole_ADMIN); err != nil {
		return err
	}
	group, err := helper.FindEntityWrapper(e.store.FindGroupByID, req.GroupID, helper.CollaborationServiceID)
	if err != nil {
		return err
	}
	if group.IsDefault {
		logger.Infof("Cannot delete group %s as it is default group for user %s", req.GroupID, req.UserID)
		return helper.ErrMicroNotAuthorized(helper.CollaborationServiceID)
	}
	err = e.store.DeleteGroup(group)
	if err != nil {
		return err
	}
	rsp.Success = true
	logger.Infof("Successfully deleted group (%s)", req.GroupID)
	return nil
}

func (e *Collaboration) GetGroupUserRole(_ context.Context, req *pb.GroupRequest, rsp *pb.GroupRoleResponse) error {
	logger.Infof("Received Collaboration.GetUserGroupRole request: %v", req)
	role, err := e.store.GetGroupUserRole(req.UserID, req.GroupID)
	if err != nil {
		if errors.Is(err, helper.ErrStoreNoEntryWithID) {
			return helper.ErrMicroNoEntryWithID(helper.CollaborationServiceID)
		}
		return err
	}
	logger.Infof("Found group with id %s by obtaining role", req.GroupID)
	rsp.GroupID = req.GroupID
	protoRole := e.migrateModelRoleToProtoRole(role)
	rsp.GroupRole = protoRole
	logger.Infof("Obtained role (%s) for group (%s) for user (%s)", rsp.GroupRole.String(), req.GroupID, req.UserID)
	return nil
}

func (e *Collaboration) FindGroupByID(_ context.Context, req *pb.GroupRequest, rsp *pb.GroupResponse) error {
	logger.Infof("Received Collaboration.FindGroupByID request: %v", req)
	group, err := helper.FindEntityWrapper(e.store.FindGroupByID, req.GroupID, helper.CollaborationServiceID)
	if err != nil {
		return err
	}
	rsp.GroupID = group.ID
	rsp.GroupName = group.Name
	return nil
}
