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

// Api Endpoints for CardDeck service

func NewCardDeckEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for CardDeck service

type CardDeckService interface {
	GetGroupDecks(ctx context.Context, in *GroupDecksRequest, opts ...client.CallOption) (*GroupDecksResponse, error)
	CreateDeck(ctx context.Context, in *CreateDeckRequest, opts ...client.CallOption) (*IDResponse, error)
	ModifyDeck(ctx context.Context, in *ModifyDeckRequest, opts ...client.CallOption) (*SuccessResponse, error)
	DeleteDeck(ctx context.Context, in *DeleteWithIDRequest, opts ...client.CallOption) (*SuccessResponse, error)
	GetDeckCards(ctx context.Context, in *DeckRequest, opts ...client.CallOption) (*DeckCardsResponse, error)
	CreateCard(ctx context.Context, in *CreateCardRequest, opts ...client.CallOption) (*IDResponse, error)
	DeleteCard(ctx context.Context, in *DeleteWithIDRequest, opts ...client.CallOption) (*SuccessResponse, error)
	CreateCardSide(ctx context.Context, in *CreateCardSideRequest, opts ...client.CallOption) (*IDResponse, error)
	ModifyCardSide(ctx context.Context, in *ModifyCardSideRequest, opts ...client.CallOption) (*SuccessResponse, error)
	DeleteCardSide(ctx context.Context, in *DeleteWithIDRequest, opts ...client.CallOption) (*SuccessResponse, error)
}

type cardDeckService struct {
	c    client.Client
	name string
}

func NewCardDeckService(name string, c client.Client) CardDeckService {
	return &cardDeckService{
		c:    c,
		name: name,
	}
}

func (c *cardDeckService) GetGroupDecks(ctx context.Context, in *GroupDecksRequest, opts ...client.CallOption) (*GroupDecksResponse, error) {
	req := c.c.NewRequest(c.name, "CardDeck.GetGroupDecks", in)
	out := new(GroupDecksResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cardDeckService) CreateDeck(ctx context.Context, in *CreateDeckRequest, opts ...client.CallOption) (*IDResponse, error) {
	req := c.c.NewRequest(c.name, "CardDeck.CreateDeck", in)
	out := new(IDResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cardDeckService) ModifyDeck(ctx context.Context, in *ModifyDeckRequest, opts ...client.CallOption) (*SuccessResponse, error) {
	req := c.c.NewRequest(c.name, "CardDeck.ModifyDeck", in)
	out := new(SuccessResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cardDeckService) DeleteDeck(ctx context.Context, in *DeleteWithIDRequest, opts ...client.CallOption) (*SuccessResponse, error) {
	req := c.c.NewRequest(c.name, "CardDeck.DeleteDeck", in)
	out := new(SuccessResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cardDeckService) GetDeckCards(ctx context.Context, in *DeckRequest, opts ...client.CallOption) (*DeckCardsResponse, error) {
	req := c.c.NewRequest(c.name, "CardDeck.GetDeckCards", in)
	out := new(DeckCardsResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cardDeckService) CreateCard(ctx context.Context, in *CreateCardRequest, opts ...client.CallOption) (*IDResponse, error) {
	req := c.c.NewRequest(c.name, "CardDeck.CreateCard", in)
	out := new(IDResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cardDeckService) DeleteCard(ctx context.Context, in *DeleteWithIDRequest, opts ...client.CallOption) (*SuccessResponse, error) {
	req := c.c.NewRequest(c.name, "CardDeck.DeleteCard", in)
	out := new(SuccessResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cardDeckService) CreateCardSide(ctx context.Context, in *CreateCardSideRequest, opts ...client.CallOption) (*IDResponse, error) {
	req := c.c.NewRequest(c.name, "CardDeck.CreateCardSide", in)
	out := new(IDResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cardDeckService) ModifyCardSide(ctx context.Context, in *ModifyCardSideRequest, opts ...client.CallOption) (*SuccessResponse, error) {
	req := c.c.NewRequest(c.name, "CardDeck.ModifyCardSide", in)
	out := new(SuccessResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cardDeckService) DeleteCardSide(ctx context.Context, in *DeleteWithIDRequest, opts ...client.CallOption) (*SuccessResponse, error) {
	req := c.c.NewRequest(c.name, "CardDeck.DeleteCardSide", in)
	out := new(SuccessResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for CardDeck service

type CardDeckHandler interface {
	GetGroupDecks(context.Context, *GroupDecksRequest, *GroupDecksResponse) error
	CreateDeck(context.Context, *CreateDeckRequest, *IDResponse) error
	ModifyDeck(context.Context, *ModifyDeckRequest, *SuccessResponse) error
	DeleteDeck(context.Context, *DeleteWithIDRequest, *SuccessResponse) error
	GetDeckCards(context.Context, *DeckRequest, *DeckCardsResponse) error
	CreateCard(context.Context, *CreateCardRequest, *IDResponse) error
	DeleteCard(context.Context, *DeleteWithIDRequest, *SuccessResponse) error
	CreateCardSide(context.Context, *CreateCardSideRequest, *IDResponse) error
	ModifyCardSide(context.Context, *ModifyCardSideRequest, *SuccessResponse) error
	DeleteCardSide(context.Context, *DeleteWithIDRequest, *SuccessResponse) error
}

func RegisterCardDeckHandler(s server.Server, hdlr CardDeckHandler, opts ...server.HandlerOption) error {
	type cardDeck interface {
		GetGroupDecks(ctx context.Context, in *GroupDecksRequest, out *GroupDecksResponse) error
		CreateDeck(ctx context.Context, in *CreateDeckRequest, out *IDResponse) error
		ModifyDeck(ctx context.Context, in *ModifyDeckRequest, out *SuccessResponse) error
		DeleteDeck(ctx context.Context, in *DeleteWithIDRequest, out *SuccessResponse) error
		GetDeckCards(ctx context.Context, in *DeckRequest, out *DeckCardsResponse) error
		CreateCard(ctx context.Context, in *CreateCardRequest, out *IDResponse) error
		DeleteCard(ctx context.Context, in *DeleteWithIDRequest, out *SuccessResponse) error
		CreateCardSide(ctx context.Context, in *CreateCardSideRequest, out *IDResponse) error
		ModifyCardSide(ctx context.Context, in *ModifyCardSideRequest, out *SuccessResponse) error
		DeleteCardSide(ctx context.Context, in *DeleteWithIDRequest, out *SuccessResponse) error
	}
	type CardDeck struct {
		cardDeck
	}
	h := &cardDeckHandler{hdlr}
	return s.Handle(s.NewHandler(&CardDeck{h}, opts...))
}

type cardDeckHandler struct {
	CardDeckHandler
}

func (h *cardDeckHandler) GetGroupDecks(ctx context.Context, in *GroupDecksRequest, out *GroupDecksResponse) error {
	return h.CardDeckHandler.GetGroupDecks(ctx, in, out)
}

func (h *cardDeckHandler) CreateDeck(ctx context.Context, in *CreateDeckRequest, out *IDResponse) error {
	return h.CardDeckHandler.CreateDeck(ctx, in, out)
}

func (h *cardDeckHandler) ModifyDeck(ctx context.Context, in *ModifyDeckRequest, out *SuccessResponse) error {
	return h.CardDeckHandler.ModifyDeck(ctx, in, out)
}

func (h *cardDeckHandler) DeleteDeck(ctx context.Context, in *DeleteWithIDRequest, out *SuccessResponse) error {
	return h.CardDeckHandler.DeleteDeck(ctx, in, out)
}

func (h *cardDeckHandler) GetDeckCards(ctx context.Context, in *DeckRequest, out *DeckCardsResponse) error {
	return h.CardDeckHandler.GetDeckCards(ctx, in, out)
}

func (h *cardDeckHandler) CreateCard(ctx context.Context, in *CreateCardRequest, out *IDResponse) error {
	return h.CardDeckHandler.CreateCard(ctx, in, out)
}

func (h *cardDeckHandler) DeleteCard(ctx context.Context, in *DeleteWithIDRequest, out *SuccessResponse) error {
	return h.CardDeckHandler.DeleteCard(ctx, in, out)
}

func (h *cardDeckHandler) CreateCardSide(ctx context.Context, in *CreateCardSideRequest, out *IDResponse) error {
	return h.CardDeckHandler.CreateCardSide(ctx, in, out)
}

func (h *cardDeckHandler) ModifyCardSide(ctx context.Context, in *ModifyCardSideRequest, out *SuccessResponse) error {
	return h.CardDeckHandler.ModifyCardSide(ctx, in, out)
}

func (h *cardDeckHandler) DeleteCardSide(ctx context.Context, in *DeleteWithIDRequest, out *SuccessResponse) error {
	return h.CardDeckHandler.DeleteCardSide(ctx, in, out)
}
