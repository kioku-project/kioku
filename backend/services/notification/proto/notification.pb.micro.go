// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: services/notification/proto/notification.proto

package notification

import (
	fmt "fmt"
	proto1 "github.com/kioku-project/kioku/pkg/proto"
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

// Api Endpoints for Notification service

func NewNotificationEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Notification service

type NotificationService interface {
	Subscribe(ctx context.Context, in *proto1.PushSubscriptionRequest, opts ...client.CallOption) (*proto1.PushSubscription, error)
	Unsubscribe(ctx context.Context, in *proto1.PushSubscriptionRequest, opts ...client.CallOption) (*proto1.Success, error)
	GetUserNotificationSubscriptions(ctx context.Context, in *proto1.User, opts ...client.CallOption) (*proto1.PushSubscriptions, error)
}

type notificationService struct {
	c    client.Client
	name string
}

func NewNotificationService(name string, c client.Client) NotificationService {
	return &notificationService{
		c:    c,
		name: name,
	}
}

func (c *notificationService) Subscribe(ctx context.Context, in *proto1.PushSubscriptionRequest, opts ...client.CallOption) (*proto1.PushSubscription, error) {
	req := c.c.NewRequest(c.name, "Notification.Subscribe", in)
	out := new(proto1.PushSubscription)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notificationService) Unsubscribe(ctx context.Context, in *proto1.PushSubscriptionRequest, opts ...client.CallOption) (*proto1.Success, error) {
	req := c.c.NewRequest(c.name, "Notification.Unsubscribe", in)
	out := new(proto1.Success)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notificationService) GetUserNotificationSubscriptions(ctx context.Context, in *proto1.User, opts ...client.CallOption) (*proto1.PushSubscriptions, error) {
	req := c.c.NewRequest(c.name, "Notification.GetUserNotificationSubscriptions", in)
	out := new(proto1.PushSubscriptions)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Notification service

type NotificationHandler interface {
	Subscribe(context.Context, *proto1.PushSubscriptionRequest, *proto1.PushSubscription) error
	Unsubscribe(context.Context, *proto1.PushSubscriptionRequest, *proto1.Success) error
	GetUserNotificationSubscriptions(context.Context, *proto1.User, *proto1.PushSubscriptions) error
}

func RegisterNotificationHandler(s server.Server, hdlr NotificationHandler, opts ...server.HandlerOption) error {
	type notification interface {
		Subscribe(ctx context.Context, in *proto1.PushSubscriptionRequest, out *proto1.PushSubscription) error
		Unsubscribe(ctx context.Context, in *proto1.PushSubscriptionRequest, out *proto1.Success) error
		GetUserNotificationSubscriptions(ctx context.Context, in *proto1.User, out *proto1.PushSubscriptions) error
	}
	type Notification struct {
		notification
	}
	h := &notificationHandler{hdlr}
	return s.Handle(s.NewHandler(&Notification{h}, opts...))
}

type notificationHandler struct {
	NotificationHandler
}

func (h *notificationHandler) Subscribe(ctx context.Context, in *proto1.PushSubscriptionRequest, out *proto1.PushSubscription) error {
	return h.NotificationHandler.Subscribe(ctx, in, out)
}

func (h *notificationHandler) Unsubscribe(ctx context.Context, in *proto1.PushSubscriptionRequest, out *proto1.Success) error {
	return h.NotificationHandler.Unsubscribe(ctx, in, out)
}

func (h *notificationHandler) GetUserNotificationSubscriptions(ctx context.Context, in *proto1.User, out *proto1.PushSubscriptions) error {
	return h.NotificationHandler.GetUserNotificationSubscriptions(ctx, in, out)
}