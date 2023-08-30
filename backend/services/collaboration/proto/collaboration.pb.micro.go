// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/collaboration.proto

package collaboration

import (
	fmt "fmt"
	proto "google.golang.org/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "go-micro.dev/v4/api"
	client "go-micro.dev/v4/client"
	server "go-micro.dev/v4/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for Collaboration service

func NewCollaborationEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Collaboration service

type CollaborationService interface {
	GetGroupInvitations(ctx context.Context, in *UserIDRequest, opts ...client.CallOption) (*GroupInvitationsResponse, error)
	GetUserGroups(ctx context.Context, in *UserIDRequest, opts ...client.CallOption) (*UserGroupsResponse, error)
	CreateNewGroupWithAdmin(ctx context.Context, in *CreateGroupRequest, opts ...client.CallOption) (*IDResponse, error)
	GetGroup(ctx context.Context, in *GroupRequest, opts ...client.CallOption) (*GroupWithUserRole, error)
	ModifyGroup(ctx context.Context, in *ModifyGroupRequest, opts ...client.CallOption) (*SuccessResponse, error)
	DeleteGroup(ctx context.Context, in *GroupRequest, opts ...client.CallOption) (*SuccessResponse, error)
	GetGroupMembers(ctx context.Context, in *GroupRequest, opts ...client.CallOption) (*GroupMembersResponse, error)
	GetGroupMemberRequests(ctx context.Context, in *GroupRequest, opts ...client.CallOption) (*GroupMemberAdmissionResponse, error)
	GetInvitationsForGroup(ctx context.Context, in *GroupRequest, opts ...client.CallOption) (*GroupMemberAdmissionResponse, error)
	GetGroupUserRole(ctx context.Context, in *GroupRequest, opts ...client.CallOption) (*GroupRoleResponse, error)
	FindGroupByID(ctx context.Context, in *GroupRequest, opts ...client.CallOption) (*Group, error)
	AddGroupUserRequest(ctx context.Context, in *GroupUserRequest, opts ...client.CallOption) (*SuccessResponse, error)
	RemoveGroupUserRequest(ctx context.Context, in *GroupUserRequest, opts ...client.CallOption) (*SuccessResponse, error)
	AddGroupUserInvite(ctx context.Context, in *GroupUserInvite, opts ...client.CallOption) (*SuccessResponse, error)
	RemoveGroupUserInvite(ctx context.Context, in *GroupUserInvite, opts ...client.CallOption) (*SuccessResponse, error)
	LeaveGroup(ctx context.Context, in *GroupRequest, opts ...client.CallOption) (*SuccessResponse, error)
	LeaveGroupSafe(ctx context.Context, in *GroupRequest, opts ...client.CallOption) (*SuccessResponse, error)
}

type collaborationService struct {
	c    client.Client
	name string
}

func NewCollaborationService(name string, c client.Client) CollaborationService {
	return &collaborationService{
		c:    c,
		name: name,
	}
}

func (c *collaborationService) GetGroupInvitations(ctx context.Context, in *UserIDRequest, opts ...client.CallOption) (*GroupInvitationsResponse, error) {
	req := c.c.NewRequest(c.name, "Collaboration.GetGroupInvitations", in)
	out := new(GroupInvitationsResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *collaborationService) GetUserGroups(ctx context.Context, in *UserIDRequest, opts ...client.CallOption) (*UserGroupsResponse, error) {
	req := c.c.NewRequest(c.name, "Collaboration.GetUserGroups", in)
	out := new(UserGroupsResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *collaborationService) CreateNewGroupWithAdmin(ctx context.Context, in *CreateGroupRequest, opts ...client.CallOption) (*IDResponse, error) {
	req := c.c.NewRequest(c.name, "Collaboration.CreateNewGroupWithAdmin", in)
	out := new(IDResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *collaborationService) GetGroup(ctx context.Context, in *GroupRequest, opts ...client.CallOption) (*GroupWithUserRole, error) {
	req := c.c.NewRequest(c.name, "Collaboration.GetGroup", in)
	out := new(GroupWithUserRole)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *collaborationService) ModifyGroup(ctx context.Context, in *ModifyGroupRequest, opts ...client.CallOption) (*SuccessResponse, error) {
	req := c.c.NewRequest(c.name, "Collaboration.ModifyGroup", in)
	out := new(SuccessResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *collaborationService) DeleteGroup(ctx context.Context, in *GroupRequest, opts ...client.CallOption) (*SuccessResponse, error) {
	req := c.c.NewRequest(c.name, "Collaboration.DeleteGroup", in)
	out := new(SuccessResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *collaborationService) GetGroupMembers(ctx context.Context, in *GroupRequest, opts ...client.CallOption) (*GroupMembersResponse, error) {
	req := c.c.NewRequest(c.name, "Collaboration.GetGroupMembers", in)
	out := new(GroupMembersResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *collaborationService) GetGroupMemberRequests(ctx context.Context, in *GroupRequest, opts ...client.CallOption) (*GroupMemberAdmissionResponse, error) {
	req := c.c.NewRequest(c.name, "Collaboration.GetGroupMemberRequests", in)
	out := new(GroupMemberAdmissionResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *collaborationService) GetInvitationsForGroup(ctx context.Context, in *GroupRequest, opts ...client.CallOption) (*GroupMemberAdmissionResponse, error) {
	req := c.c.NewRequest(c.name, "Collaboration.GetInvitationsForGroup", in)
	out := new(GroupMemberAdmissionResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *collaborationService) GetGroupUserRole(ctx context.Context, in *GroupRequest, opts ...client.CallOption) (*GroupRoleResponse, error) {
	req := c.c.NewRequest(c.name, "Collaboration.GetGroupUserRole", in)
	out := new(GroupRoleResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *collaborationService) FindGroupByID(ctx context.Context, in *GroupRequest, opts ...client.CallOption) (*Group, error) {
	req := c.c.NewRequest(c.name, "Collaboration.FindGroupByID", in)
	out := new(Group)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *collaborationService) AddGroupUserRequest(ctx context.Context, in *GroupUserRequest, opts ...client.CallOption) (*SuccessResponse, error) {
	req := c.c.NewRequest(c.name, "Collaboration.AddGroupUserRequest", in)
	out := new(SuccessResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *collaborationService) RemoveGroupUserRequest(ctx context.Context, in *GroupUserRequest, opts ...client.CallOption) (*SuccessResponse, error) {
	req := c.c.NewRequest(c.name, "Collaboration.RemoveGroupUserRequest", in)
	out := new(SuccessResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *collaborationService) AddGroupUserInvite(ctx context.Context, in *GroupUserInvite, opts ...client.CallOption) (*SuccessResponse, error) {
	req := c.c.NewRequest(c.name, "Collaboration.AddGroupUserInvite", in)
	out := new(SuccessResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *collaborationService) RemoveGroupUserInvite(ctx context.Context, in *GroupUserInvite, opts ...client.CallOption) (*SuccessResponse, error) {
	req := c.c.NewRequest(c.name, "Collaboration.RemoveGroupUserInvite", in)
	out := new(SuccessResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *collaborationService) LeaveGroup(ctx context.Context, in *GroupRequest, opts ...client.CallOption) (*SuccessResponse, error) {
	req := c.c.NewRequest(c.name, "Collaboration.LeaveGroup", in)
	out := new(SuccessResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *collaborationService) LeaveGroupSafe(ctx context.Context, in *GroupRequest, opts ...client.CallOption) (*SuccessResponse, error) {
	req := c.c.NewRequest(c.name, "Collaboration.LeaveGroupSafe", in)
	out := new(SuccessResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Collaboration service

type CollaborationHandler interface {
	GetGroupInvitations(context.Context, *UserIDRequest, *GroupInvitationsResponse) error
	GetUserGroups(context.Context, *UserIDRequest, *UserGroupsResponse) error
	CreateNewGroupWithAdmin(context.Context, *CreateGroupRequest, *IDResponse) error
	GetGroup(context.Context, *GroupRequest, *GroupWithUserRole) error
	ModifyGroup(context.Context, *ModifyGroupRequest, *SuccessResponse) error
	DeleteGroup(context.Context, *GroupRequest, *SuccessResponse) error
	GetGroupMembers(context.Context, *GroupRequest, *GroupMembersResponse) error
	GetGroupMemberRequests(context.Context, *GroupRequest, *GroupMemberAdmissionResponse) error
	GetInvitationsForGroup(context.Context, *GroupRequest, *GroupMemberAdmissionResponse) error
	GetGroupUserRole(context.Context, *GroupRequest, *GroupRoleResponse) error
	FindGroupByID(context.Context, *GroupRequest, *Group) error
	AddGroupUserRequest(context.Context, *GroupUserRequest, *SuccessResponse) error
	RemoveGroupUserRequest(context.Context, *GroupUserRequest, *SuccessResponse) error
	AddGroupUserInvite(context.Context, *GroupUserInvite, *SuccessResponse) error
	RemoveGroupUserInvite(context.Context, *GroupUserInvite, *SuccessResponse) error
	LeaveGroup(context.Context, *GroupRequest, *SuccessResponse) error
	LeaveGroupSafe(context.Context, *GroupRequest, *SuccessResponse) error
}

func RegisterCollaborationHandler(s server.Server, hdlr CollaborationHandler, opts ...server.HandlerOption) error {
	type collaboration interface {
		GetGroupInvitations(ctx context.Context, in *UserIDRequest, out *GroupInvitationsResponse) error
		GetUserGroups(ctx context.Context, in *UserIDRequest, out *UserGroupsResponse) error
		CreateNewGroupWithAdmin(ctx context.Context, in *CreateGroupRequest, out *IDResponse) error
		GetGroup(ctx context.Context, in *GroupRequest, out *GroupWithUserRole) error
		ModifyGroup(ctx context.Context, in *ModifyGroupRequest, out *SuccessResponse) error
		DeleteGroup(ctx context.Context, in *GroupRequest, out *SuccessResponse) error
		GetGroupMembers(ctx context.Context, in *GroupRequest, out *GroupMembersResponse) error
		GetGroupMemberRequests(ctx context.Context, in *GroupRequest, out *GroupMemberAdmissionResponse) error
		GetInvitationsForGroup(ctx context.Context, in *GroupRequest, out *GroupMemberAdmissionResponse) error
		GetGroupUserRole(ctx context.Context, in *GroupRequest, out *GroupRoleResponse) error
		FindGroupByID(ctx context.Context, in *GroupRequest, out *Group) error
		AddGroupUserRequest(ctx context.Context, in *GroupUserRequest, out *SuccessResponse) error
		RemoveGroupUserRequest(ctx context.Context, in *GroupUserRequest, out *SuccessResponse) error
		AddGroupUserInvite(ctx context.Context, in *GroupUserInvite, out *SuccessResponse) error
		RemoveGroupUserInvite(ctx context.Context, in *GroupUserInvite, out *SuccessResponse) error
		LeaveGroup(ctx context.Context, in *GroupRequest, out *SuccessResponse) error
		LeaveGroupSafe(ctx context.Context, in *GroupRequest, out *SuccessResponse) error
	}
	type Collaboration struct {
		collaboration
	}
	h := &collaborationHandler{hdlr}
	return s.Handle(s.NewHandler(&Collaboration{h}, opts...))
}

type collaborationHandler struct {
	CollaborationHandler
}

func (h *collaborationHandler) GetGroupInvitations(ctx context.Context, in *UserIDRequest, out *GroupInvitationsResponse) error {
	return h.CollaborationHandler.GetGroupInvitations(ctx, in, out)
}

func (h *collaborationHandler) GetUserGroups(ctx context.Context, in *UserIDRequest, out *UserGroupsResponse) error {
	return h.CollaborationHandler.GetUserGroups(ctx, in, out)
}

func (h *collaborationHandler) CreateNewGroupWithAdmin(ctx context.Context, in *CreateGroupRequest, out *IDResponse) error {
	return h.CollaborationHandler.CreateNewGroupWithAdmin(ctx, in, out)
}

func (h *collaborationHandler) GetGroup(ctx context.Context, in *GroupRequest, out *GroupWithUserRole) error {
	return h.CollaborationHandler.GetGroup(ctx, in, out)
}

func (h *collaborationHandler) ModifyGroup(ctx context.Context, in *ModifyGroupRequest, out *SuccessResponse) error {
	return h.CollaborationHandler.ModifyGroup(ctx, in, out)
}

func (h *collaborationHandler) DeleteGroup(ctx context.Context, in *GroupRequest, out *SuccessResponse) error {
	return h.CollaborationHandler.DeleteGroup(ctx, in, out)
}

func (h *collaborationHandler) GetGroupMembers(ctx context.Context, in *GroupRequest, out *GroupMembersResponse) error {
	return h.CollaborationHandler.GetGroupMembers(ctx, in, out)
}

func (h *collaborationHandler) GetGroupMemberRequests(ctx context.Context, in *GroupRequest, out *GroupMemberAdmissionResponse) error {
	return h.CollaborationHandler.GetGroupMemberRequests(ctx, in, out)
}

func (h *collaborationHandler) GetInvitationsForGroup(ctx context.Context, in *GroupRequest, out *GroupMemberAdmissionResponse) error {
	return h.CollaborationHandler.GetInvitationsForGroup(ctx, in, out)
}

func (h *collaborationHandler) GetGroupUserRole(ctx context.Context, in *GroupRequest, out *GroupRoleResponse) error {
	return h.CollaborationHandler.GetGroupUserRole(ctx, in, out)
}

func (h *collaborationHandler) FindGroupByID(ctx context.Context, in *GroupRequest, out *Group) error {
	return h.CollaborationHandler.FindGroupByID(ctx, in, out)
}

func (h *collaborationHandler) AddGroupUserRequest(ctx context.Context, in *GroupUserRequest, out *SuccessResponse) error {
	return h.CollaborationHandler.AddGroupUserRequest(ctx, in, out)
}

func (h *collaborationHandler) RemoveGroupUserRequest(ctx context.Context, in *GroupUserRequest, out *SuccessResponse) error {
	return h.CollaborationHandler.RemoveGroupUserRequest(ctx, in, out)
}

func (h *collaborationHandler) AddGroupUserInvite(ctx context.Context, in *GroupUserInvite, out *SuccessResponse) error {
	return h.CollaborationHandler.AddGroupUserInvite(ctx, in, out)
}

func (h *collaborationHandler) RemoveGroupUserInvite(ctx context.Context, in *GroupUserInvite, out *SuccessResponse) error {
	return h.CollaborationHandler.RemoveGroupUserInvite(ctx, in, out)
}

func (h *collaborationHandler) LeaveGroup(ctx context.Context, in *GroupRequest, out *SuccessResponse) error {
	return h.CollaborationHandler.LeaveGroup(ctx, in, out)
}

func (h *collaborationHandler) LeaveGroupSafe(ctx context.Context, in *GroupRequest, out *SuccessResponse) error {
	return h.CollaborationHandler.LeaveGroupSafe(ctx, in, out)
}
