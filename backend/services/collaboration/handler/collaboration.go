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

func (e *Collaboration) CreateNewGroupWithAdmin(ctx context.Context, req *pb.CreateGroupRequest, rsp *pb.SuccessResponse) error {
	logger.Infof("Received Collaboration.CreateNewGroupWithAdmin request: %v", req)
	newGroup := model.Group{
		Name: req.GroupName,
	}
	err := e.store.CreateNewGroupWithAdmin(req.UserID, &newGroup)
	if err != nil {
		return err
	}
	rsp.Success = true
	logger.Infof("Successfully created new group (%s) with user (%s) as admin", newGroup.ID, req.UserID)
	return nil
}

func (e *Collaboration) GetGroupUserRole(ctx context.Context, req *pb.GroupRequest, rsp *pb.GroupRoleResponse) error {
	logger.Infof("Received Collaboration.GetUserGroupRole request: %v", req)
	group, err := e.store.FindGroupByID(req.GroupID)
	if err != nil {
		if errors.Is(err, helper.ErrStoreNoEntryWithID) {
			return helper.ErrMicroNoEntryWithID(helper.CollaborationServiceID)
		}
		return err
	}
	logger.Infof("Found group with id %s", req.GroupID)
	role, err := e.store.GetGroupUserRole(req.UserID, group.ID)
	if err != nil {
		if errors.Is(err, helper.ErrStoreNoEntryWithID) {
			return helper.ErrMicroNoEntryWithID(helper.CollaborationServiceID)
		}
		return err
	}
	rsp.GroupID = group.ID
	if role == model.RoleRead {
		rsp.GroupRole = pb.GroupRole_READ
	} else if role == model.RoleWrite {
		rsp.GroupRole = pb.GroupRole_WRITE
	} else if role == model.RoleAdmin {
		rsp.GroupRole = pb.GroupRole_ADMIN
	}
	logger.Infof("Obtained role (%s) for group (%s) for user (%s)", rsp.GroupRole.String(), req.GroupID, req.UserID)
	return nil
}

func (e *Collaboration) GetUserGroups(ctx context.Context, req *pb.UserGroupsRequest, rsp *pb.UserGroupsResponse) error {
	logger.Infof("Received Collaboration.GetUserGroups request: %v", req)
	groups, err := e.store.FindGroupsByUserID(req.UserID)
	if err != nil {
		if errors.Is(err, helper.ErrStoreNoEntryWithID) {
			return helper.ErrMicroNoEntryWithID(helper.CollaborationServiceID)
		}
		return err
	}
	rsp.Groups = make([]*pb.Group, len(groups))
	for i, group := range groups {
		rsp.Groups[i] = &pb.Group{
			GroupID:   group.ID,
			GroupName: group.Name,
		}
	}
	logger.Infof("Found %d groups for user with id %s", len(groups), req.UserID)
	return nil
}

func (e *Collaboration) FindGroupByID(ctx context.Context, req *pb.GroupRequest, rsp *pb.GroupResponse) error {
	logger.Infof("Received Collaboration.FindGroupByID request: %v", req)
	group, err := e.store.FindGroupByID(req.GroupID)
	if err != nil {
		if errors.Is(err, helper.ErrStoreNoEntryWithID) {
			return helper.ErrMicroNoEntryWithID(helper.CollaborationServiceID)
		}
		return err
	}
	rsp.GroupID = group.ID
	rsp.GroupName = group.Name
	logger.Infof("Found group with id %s", req.GroupID)
	return nil
}
