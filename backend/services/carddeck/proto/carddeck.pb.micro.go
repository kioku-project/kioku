// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/carddeck.proto

package carddeck

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

// Api Endpoints for Carddeck service

func NewCarddeckEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Carddeck service

type CarddeckService interface {
	CreateCard(ctx context.Context, in *CardRequest, opts ...client.CallOption) (*SuccessResponse, error)
	CreateDeck(ctx context.Context, in *DeckRequest, opts ...client.CallOption) (*SuccessResponse, error)
}

type carddeckService struct {
	c    client.Client
	name string
}

func NewCarddeckService(name string, c client.Client) CarddeckService {
	return &carddeckService{
		c:    c,
		name: name,
	}
}

func (c *carddeckService) CreateCard(ctx context.Context, in *CardRequest, opts ...client.CallOption) (*SuccessResponse, error) {
	req := c.c.NewRequest(c.name, "Carddeck.CreateCard", in)
	out := new(SuccessResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *carddeckService) CreateDeck(ctx context.Context, in *DeckRequest, opts ...client.CallOption) (*SuccessResponse, error) {
	req := c.c.NewRequest(c.name, "Carddeck.CreateDeck", in)
	out := new(SuccessResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Carddeck service

type CarddeckHandler interface {
	CreateCard(context.Context, *CardRequest, *SuccessResponse) error
	CreateDeck(context.Context, *DeckRequest, *SuccessResponse) error
}

func RegisterCarddeckHandler(s server.Server, hdlr CarddeckHandler, opts ...server.HandlerOption) error {
	type carddeck interface {
		CreateCard(ctx context.Context, in *CardRequest, out *SuccessResponse) error
		CreateDeck(ctx context.Context, in *DeckRequest, out *SuccessResponse) error
	}
	type Carddeck struct {
		carddeck
	}
	h := &carddeckHandler{hdlr}
	return s.Handle(s.NewHandler(&Carddeck{h}, opts...))
}

type carddeckHandler struct {
	CarddeckHandler
}

func (h *carddeckHandler) CreateCard(ctx context.Context, in *CardRequest, out *SuccessResponse) error {
	return h.CarddeckHandler.CreateCard(ctx, in, out)
}

func (h *carddeckHandler) CreateDeck(ctx context.Context, in *DeckRequest, out *SuccessResponse) error {
	return h.CarddeckHandler.CreateDeck(ctx, in, out)
}
