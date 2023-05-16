package handler

import (
	"context"

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
	err := e.store.CreateNewGroupWithAdmin(uint(req.UserID), &newGroup)
	if err != nil {
		return err
	}
	rsp.Success = true
	return nil
}

func (e *Collaboration) GetGroupUserRole(ctx context.Context, req *pb.GroupRequest, rsp *pb.GroupRoleResponse) error {
	logger.Infof("Received Collaboration.GetUserGroupRole request: %v", req)
	group, err := e.store.FindGroupByPublicID(req.GroupPublicID)
	if err != nil {
		return err
	}
	logger.Infof("Obtain role entry from database: UserID(%v) GroupID(%v)", req.UserID, group.ID)
	role, err := e.store.GetGroupUserRole(uint(req.UserID), group.ID)
	if err != nil {
		return err
	}
	rsp.GroupID = uint64(group.ID)
	if *role == model.RoleRead {
		rsp.GroupRole = pb.GroupRole_WRITE
	} else if *role == model.RoleWrite {
		rsp.GroupRole = pb.GroupRole_WRITE
	} else if *role == model.RoleAdmin {
		rsp.GroupRole = pb.GroupRole_ADMIN
	}
	return nil
}

func (e *Collaboration) GetUserGroups(ctx context.Context, req *pb.UserGroupsRequest, rsp *pb.UserGroupsResponse) error {
	logger.Infof("Received Carddeck.GetUserGroups request: %v", req)
	group, err := e.store.FindGroupsByUserID(uint(req.UserID))
	if err != nil {
		return err
	}
	rsp.Groups = make([]*pb.Group, len(group))
	for i, group := range group {
		rsp.Groups[i] = &pb.Group{
			GroupPublicID: group.PublicID,
			GroupName:     group.Name,
		}
	}
	return nil
}

func (e *Collaboration) FindGroupByPublicID(ctx context.Context, req *pb.GroupRequest, rsp *pb.GroupResponse) error {
	logger.Infof("Received Carddeck.FindGroupByPublicID request: %v", req)
	group, err := e.store.FindGroupByPublicID(req.GroupPublicID)
	if err != nil {
		return err
	}
	rsp.GroupID = uint64(group.ID)
	rsp.GroupName = group.Name
	return nil
}
