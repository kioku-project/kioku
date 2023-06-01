// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.12
// source: proto/carddeck.proto

package carddeck

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GroupDecksRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID  string `protobuf:"bytes,1,opt,name=userID,proto3" json:"userID,omitempty"`
	GroupID string `protobuf:"bytes,2,opt,name=groupID,proto3" json:"groupID,omitempty"`
}

func (x *GroupDecksRequest) Reset() {
	*x = GroupDecksRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_carddeck_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GroupDecksRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GroupDecksRequest) ProtoMessage() {}

func (x *GroupDecksRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_carddeck_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GroupDecksRequest.ProtoReflect.Descriptor instead.
func (*GroupDecksRequest) Descriptor() ([]byte, []int) {
	return file_proto_carddeck_proto_rawDescGZIP(), []int{0}
}

func (x *GroupDecksRequest) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

func (x *GroupDecksRequest) GetGroupID() string {
	if x != nil {
		return x.GroupID
	}
	return ""
}

type GroupDecksResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Decks []*Deck `protobuf:"bytes,1,rep,name=decks,proto3" json:"decks,omitempty"`
}

func (x *GroupDecksResponse) Reset() {
	*x = GroupDecksResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_carddeck_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GroupDecksResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GroupDecksResponse) ProtoMessage() {}

func (x *GroupDecksResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_carddeck_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GroupDecksResponse.ProtoReflect.Descriptor instead.
func (*GroupDecksResponse) Descriptor() ([]byte, []int) {
	return file_proto_carddeck_proto_rawDescGZIP(), []int{1}
}

func (x *GroupDecksResponse) GetDecks() []*Deck {
	if x != nil {
		return x.Decks
	}
	return nil
}

type Deck struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DeckID   string `protobuf:"bytes,1,opt,name=deckID,proto3" json:"deckID,omitempty"`
	DeckName string `protobuf:"bytes,2,opt,name=deckName,proto3" json:"deckName,omitempty"`
}

func (x *Deck) Reset() {
	*x = Deck{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_carddeck_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Deck) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Deck) ProtoMessage() {}

func (x *Deck) ProtoReflect() protoreflect.Message {
	mi := &file_proto_carddeck_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Deck.ProtoReflect.Descriptor instead.
func (*Deck) Descriptor() ([]byte, []int) {
	return file_proto_carddeck_proto_rawDescGZIP(), []int{2}
}

func (x *Deck) GetDeckID() string {
	if x != nil {
		return x.DeckID
	}
	return ""
}

func (x *Deck) GetDeckName() string {
	if x != nil {
		return x.DeckName
	}
	return ""
}

type CreateDeckRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID   string `protobuf:"bytes,1,opt,name=userID,proto3" json:"userID,omitempty"`
	GroupID  string `protobuf:"bytes,2,opt,name=groupID,proto3" json:"groupID,omitempty"`
	DeckName string `protobuf:"bytes,3,opt,name=deckName,proto3" json:"deckName,omitempty"`
}

func (x *CreateDeckRequest) Reset() {
	*x = CreateDeckRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_carddeck_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateDeckRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateDeckRequest) ProtoMessage() {}

func (x *CreateDeckRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_carddeck_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateDeckRequest.ProtoReflect.Descriptor instead.
func (*CreateDeckRequest) Descriptor() ([]byte, []int) {
	return file_proto_carddeck_proto_rawDescGZIP(), []int{3}
}

func (x *CreateDeckRequest) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

func (x *CreateDeckRequest) GetGroupID() string {
	if x != nil {
		return x.GroupID
	}
	return ""
}

func (x *CreateDeckRequest) GetDeckName() string {
	if x != nil {
		return x.DeckName
	}
	return ""
}

type IDResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID string `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
}

func (x *IDResponse) Reset() {
	*x = IDResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_carddeck_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IDResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IDResponse) ProtoMessage() {}

func (x *IDResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_carddeck_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IDResponse.ProtoReflect.Descriptor instead.
func (*IDResponse) Descriptor() ([]byte, []int) {
	return file_proto_carddeck_proto_rawDescGZIP(), []int{4}
}

func (x *IDResponse) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

type ModifyDeckRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID   string  `protobuf:"bytes,1,opt,name=userID,proto3" json:"userID,omitempty"`
	DeckID   string  `protobuf:"bytes,2,opt,name=deckID,proto3" json:"deckID,omitempty"`
	DeckName *string `protobuf:"bytes,3,opt,name=deckName,proto3,oneof" json:"deckName,omitempty"`
}

func (x *ModifyDeckRequest) Reset() {
	*x = ModifyDeckRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_carddeck_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ModifyDeckRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ModifyDeckRequest) ProtoMessage() {}

func (x *ModifyDeckRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_carddeck_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ModifyDeckRequest.ProtoReflect.Descriptor instead.
func (*ModifyDeckRequest) Descriptor() ([]byte, []int) {
	return file_proto_carddeck_proto_rawDescGZIP(), []int{5}
}

func (x *ModifyDeckRequest) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

func (x *ModifyDeckRequest) GetDeckID() string {
	if x != nil {
		return x.DeckID
	}
	return ""
}

func (x *ModifyDeckRequest) GetDeckName() string {
	if x != nil && x.DeckName != nil {
		return *x.DeckName
	}
	return ""
}

type SuccessResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *SuccessResponse) Reset() {
	*x = SuccessResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_carddeck_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SuccessResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SuccessResponse) ProtoMessage() {}

func (x *SuccessResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_carddeck_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SuccessResponse.ProtoReflect.Descriptor instead.
func (*SuccessResponse) Descriptor() ([]byte, []int) {
	return file_proto_carddeck_proto_rawDescGZIP(), []int{6}
}

func (x *SuccessResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type DeleteWithIDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID   string `protobuf:"bytes,1,opt,name=userID,proto3" json:"userID,omitempty"`
	EntityID string `protobuf:"bytes,2,opt,name=entityID,proto3" json:"entityID,omitempty"`
}

func (x *DeleteWithIDRequest) Reset() {
	*x = DeleteWithIDRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_carddeck_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteWithIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteWithIDRequest) ProtoMessage() {}

func (x *DeleteWithIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_carddeck_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteWithIDRequest.ProtoReflect.Descriptor instead.
func (*DeleteWithIDRequest) Descriptor() ([]byte, []int) {
	return file_proto_carddeck_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteWithIDRequest) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

func (x *DeleteWithIDRequest) GetEntityID() string {
	if x != nil {
		return x.EntityID
	}
	return ""
}

type DeckRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID string `protobuf:"bytes,1,opt,name=userID,proto3" json:"userID,omitempty"`
	DeckID string `protobuf:"bytes,2,opt,name=deckID,proto3" json:"deckID,omitempty"`
}

func (x *DeckRequest) Reset() {
	*x = DeckRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_carddeck_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeckRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeckRequest) ProtoMessage() {}

func (x *DeckRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_carddeck_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeckRequest.ProtoReflect.Descriptor instead.
func (*DeckRequest) Descriptor() ([]byte, []int) {
	return file_proto_carddeck_proto_rawDescGZIP(), []int{8}
}

func (x *DeckRequest) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

func (x *DeckRequest) GetDeckID() string {
	if x != nil {
		return x.DeckID
	}
	return ""
}

type DeckCardsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cards []*Card `protobuf:"bytes,1,rep,name=cards,proto3" json:"cards,omitempty"`
}

func (x *DeckCardsResponse) Reset() {
	*x = DeckCardsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_carddeck_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeckCardsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeckCardsResponse) ProtoMessage() {}

func (x *DeckCardsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_carddeck_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeckCardsResponse.ProtoReflect.Descriptor instead.
func (*DeckCardsResponse) Descriptor() ([]byte, []int) {
	return file_proto_carddeck_proto_rawDescGZIP(), []int{9}
}

func (x *DeckCardsResponse) GetCards() []*Card {
	if x != nil {
		return x.Cards
	}
	return nil
}

type Card struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CardID string      `protobuf:"bytes,1,opt,name=cardID,proto3" json:"cardID,omitempty"`
	Sides  []*CardSide `protobuf:"bytes,2,rep,name=sides,proto3" json:"sides,omitempty"`
}

func (x *Card) Reset() {
	*x = Card{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_carddeck_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Card) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Card) ProtoMessage() {}

func (x *Card) ProtoReflect() protoreflect.Message {
	mi := &file_proto_carddeck_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Card.ProtoReflect.Descriptor instead.
func (*Card) Descriptor() ([]byte, []int) {
	return file_proto_carddeck_proto_rawDescGZIP(), []int{10}
}

func (x *Card) GetCardID() string {
	if x != nil {
		return x.CardID
	}
	return ""
}

func (x *Card) GetSides() []*CardSide {
	if x != nil {
		return x.Sides
	}
	return nil
}

type CardSide struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CardSideID string `protobuf:"bytes,1,opt,name=cardSideID,proto3" json:"cardSideID,omitempty"`
	Content    string `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *CardSide) Reset() {
	*x = CardSide{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_carddeck_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CardSide) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CardSide) ProtoMessage() {}

func (x *CardSide) ProtoReflect() protoreflect.Message {
	mi := &file_proto_carddeck_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CardSide.ProtoReflect.Descriptor instead.
func (*CardSide) Descriptor() ([]byte, []int) {
	return file_proto_carddeck_proto_rawDescGZIP(), []int{11}
}

func (x *CardSide) GetCardSideID() string {
	if x != nil {
		return x.CardSideID
	}
	return ""
}

func (x *CardSide) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

type CreateCardRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID string   `protobuf:"bytes,1,opt,name=userID,proto3" json:"userID,omitempty"`
	DeckID string   `protobuf:"bytes,2,opt,name=deckID,proto3" json:"deckID,omitempty"`
	Sides  []string `protobuf:"bytes,3,rep,name=sides,proto3" json:"sides,omitempty"`
}

func (x *CreateCardRequest) Reset() {
	*x = CreateCardRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_carddeck_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateCardRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCardRequest) ProtoMessage() {}

func (x *CreateCardRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_carddeck_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCardRequest.ProtoReflect.Descriptor instead.
func (*CreateCardRequest) Descriptor() ([]byte, []int) {
	return file_proto_carddeck_proto_rawDescGZIP(), []int{12}
}

func (x *CreateCardRequest) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

func (x *CreateCardRequest) GetDeckID() string {
	if x != nil {
		return x.DeckID
	}
	return ""
}

func (x *CreateCardRequest) GetSides() []string {
	if x != nil {
		return x.Sides
	}
	return nil
}

type ModifyCardSideRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID     string  `protobuf:"bytes,1,opt,name=userID,proto3" json:"userID,omitempty"`
	CardSideID string  `protobuf:"bytes,2,opt,name=cardSideID,proto3" json:"cardSideID,omitempty"`
	Content    *string `protobuf:"bytes,3,opt,name=content,proto3,oneof" json:"content,omitempty"`
}

func (x *ModifyCardSideRequest) Reset() {
	*x = ModifyCardSideRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_carddeck_proto_msgTypes[13]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ModifyCardSideRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ModifyCardSideRequest) ProtoMessage() {}

func (x *ModifyCardSideRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_carddeck_proto_msgTypes[13]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ModifyCardSideRequest.ProtoReflect.Descriptor instead.
func (*ModifyCardSideRequest) Descriptor() ([]byte, []int) {
	return file_proto_carddeck_proto_rawDescGZIP(), []int{13}
}

func (x *ModifyCardSideRequest) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

func (x *ModifyCardSideRequest) GetCardSideID() string {
	if x != nil {
		return x.CardSideID
	}
	return ""
}

func (x *ModifyCardSideRequest) GetContent() string {
	if x != nil && x.Content != nil {
		return *x.Content
	}
	return ""
}

var File_proto_carddeck_proto protoreflect.FileDescriptor

var file_proto_carddeck_proto_rawDesc = []byte{
	0x0a, 0x14, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x61, 0x72, 0x64, 0x64, 0x65, 0x63, 0x6b,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x63, 0x61, 0x72, 0x64, 0x64, 0x65, 0x63, 0x6b,
	0x22, 0x45, 0x0a, 0x11, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x44, 0x65, 0x63, 0x6b, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x18, 0x0a,
	0x07, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x67, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x44, 0x22, 0x3a, 0x0a, 0x12, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x44, 0x65, 0x63, 0x6b, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x24, 0x0a,
	0x05, 0x64, 0x65, 0x63, 0x6b, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x63,
	0x61, 0x72, 0x64, 0x64, 0x65, 0x63, 0x6b, 0x2e, 0x44, 0x65, 0x63, 0x6b, 0x52, 0x05, 0x64, 0x65,
	0x63, 0x6b, 0x73, 0x22, 0x3a, 0x0a, 0x04, 0x44, 0x65, 0x63, 0x6b, 0x12, 0x16, 0x0a, 0x06, 0x64,
	0x65, 0x63, 0x6b, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x65, 0x63,
	0x6b, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x65, 0x63, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x65, 0x63, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x22,
	0x61, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x18, 0x0a, 0x07,
	0x67, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x67,
	0x72, 0x6f, 0x75, 0x70, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x65, 0x63, 0x6b, 0x4e, 0x61,
	0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x65, 0x63, 0x6b, 0x4e, 0x61,
	0x6d, 0x65, 0x22, 0x1c, 0x0a, 0x0a, 0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44,
	0x22, 0x71, 0x0a, 0x11, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x44, 0x65, 0x63, 0x6b, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x16, 0x0a,
	0x06, 0x64, 0x65, 0x63, 0x6b, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64,
	0x65, 0x63, 0x6b, 0x49, 0x44, 0x12, 0x1f, 0x0a, 0x08, 0x64, 0x65, 0x63, 0x6b, 0x4e, 0x61, 0x6d,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x08, 0x64, 0x65, 0x63, 0x6b, 0x4e,
	0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x64, 0x65, 0x63, 0x6b, 0x4e,
	0x61, 0x6d, 0x65, 0x22, 0x2b, 0x0a, 0x0f, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x22, 0x49, 0x0a, 0x13, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x57, 0x69, 0x74, 0x68, 0x49, 0x44,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12,
	0x1a, 0x0a, 0x08, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x49, 0x44, 0x22, 0x3d, 0x0a, 0x0b, 0x44,
	0x65, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x65, 0x63, 0x6b, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x64, 0x65, 0x63, 0x6b, 0x49, 0x44, 0x22, 0x39, 0x0a, 0x11, 0x44, 0x65,
	0x63, 0x6b, 0x43, 0x61, 0x72, 0x64, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x24, 0x0a, 0x05, 0x63, 0x61, 0x72, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e,
	0x2e, 0x63, 0x61, 0x72, 0x64, 0x64, 0x65, 0x63, 0x6b, 0x2e, 0x43, 0x61, 0x72, 0x64, 0x52, 0x05,
	0x63, 0x61, 0x72, 0x64, 0x73, 0x22, 0x48, 0x0a, 0x04, 0x43, 0x61, 0x72, 0x64, 0x12, 0x16, 0x0a,
	0x06, 0x63, 0x61, 0x72, 0x64, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63,
	0x61, 0x72, 0x64, 0x49, 0x44, 0x12, 0x28, 0x0a, 0x05, 0x73, 0x69, 0x64, 0x65, 0x73, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x63, 0x61, 0x72, 0x64, 0x64, 0x65, 0x63, 0x6b, 0x2e,
	0x43, 0x61, 0x72, 0x64, 0x53, 0x69, 0x64, 0x65, 0x52, 0x05, 0x73, 0x69, 0x64, 0x65, 0x73, 0x22,
	0x44, 0x0a, 0x08, 0x43, 0x61, 0x72, 0x64, 0x53, 0x69, 0x64, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x63,
	0x61, 0x72, 0x64, 0x53, 0x69, 0x64, 0x65, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x63, 0x61, 0x72, 0x64, 0x53, 0x69, 0x64, 0x65, 0x49, 0x44, 0x12, 0x18, 0x0a, 0x07, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x59, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43,
	0x61, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x65, 0x63, 0x6b, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x64, 0x65, 0x63, 0x6b, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x69,
	0x64, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x73, 0x69, 0x64, 0x65, 0x73,
	0x22, 0x7a, 0x0a, 0x15, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x43, 0x61, 0x72, 0x64, 0x53, 0x69,
	0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x44, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x61, 0x72, 0x64, 0x53, 0x69, 0x64, 0x65, 0x49, 0x44, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x61, 0x72, 0x64, 0x53, 0x69, 0x64, 0x65, 0x49,
	0x44, 0x12, 0x1d, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x48, 0x00, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x88, 0x01, 0x01,
	0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x32, 0x9e, 0x05, 0x0a,
	0x08, 0x43, 0x61, 0x72, 0x64, 0x44, 0x65, 0x63, 0x6b, 0x12, 0x4c, 0x0a, 0x0d, 0x47, 0x65, 0x74,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x44, 0x65, 0x63, 0x6b, 0x73, 0x12, 0x1b, 0x2e, 0x63, 0x61, 0x72,
	0x64, 0x64, 0x65, 0x63, 0x6b, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x44, 0x65, 0x63, 0x6b, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x63, 0x61, 0x72, 0x64, 0x64, 0x65,
	0x63, 0x6b, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x44, 0x65, 0x63, 0x6b, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x41, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x44, 0x65, 0x63, 0x6b, 0x12, 0x1b, 0x2e, 0x63, 0x61, 0x72, 0x64, 0x64, 0x65, 0x63, 0x6b,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x14, 0x2e, 0x63, 0x61, 0x72, 0x64, 0x64, 0x65, 0x63, 0x6b, 0x2e, 0x49, 0x44,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x46, 0x0a, 0x0a, 0x4d, 0x6f,
	0x64, 0x69, 0x66, 0x79, 0x44, 0x65, 0x63, 0x6b, 0x12, 0x1b, 0x2e, 0x63, 0x61, 0x72, 0x64, 0x64,
	0x65, 0x63, 0x6b, 0x2e, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x44, 0x65, 0x63, 0x6b, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x63, 0x61, 0x72, 0x64, 0x64, 0x65, 0x63, 0x6b,
	0x2e, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x48, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x44, 0x65, 0x63, 0x6b,
	0x12, 0x1d, 0x2e, 0x63, 0x61, 0x72, 0x64, 0x64, 0x65, 0x63, 0x6b, 0x2e, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x57, 0x69, 0x74, 0x68, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x19, 0x2e, 0x63, 0x61, 0x72, 0x64, 0x64, 0x65, 0x63, 0x6b, 0x2e, 0x53, 0x75, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x44, 0x0a, 0x0c,
	0x47, 0x65, 0x74, 0x44, 0x65, 0x63, 0x6b, 0x43, 0x61, 0x72, 0x64, 0x73, 0x12, 0x15, 0x2e, 0x63,
	0x61, 0x72, 0x64, 0x64, 0x65, 0x63, 0x6b, 0x2e, 0x44, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x63, 0x61, 0x72, 0x64, 0x64, 0x65, 0x63, 0x6b, 0x2e, 0x44,
	0x65, 0x63, 0x6b, 0x43, 0x61, 0x72, 0x64, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x41, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x61, 0x72, 0x64,
	0x12, 0x1b, 0x2e, 0x63, 0x61, 0x72, 0x64, 0x64, 0x65, 0x63, 0x6b, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x43, 0x61, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e,
	0x63, 0x61, 0x72, 0x64, 0x64, 0x65, 0x63, 0x6b, 0x2e, 0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x48, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43,
	0x61, 0x72, 0x64, 0x12, 0x1d, 0x2e, 0x63, 0x61, 0x72, 0x64, 0x64, 0x65, 0x63, 0x6b, 0x2e, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x57, 0x69, 0x74, 0x68, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x19, 0x2e, 0x63, 0x61, 0x72, 0x64, 0x64, 0x65, 0x63, 0x6b, 0x2e, 0x53, 0x75,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x4e, 0x0a, 0x0e, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x43, 0x61, 0x72, 0x64, 0x53, 0x69, 0x64,
	0x65, 0x12, 0x1f, 0x2e, 0x63, 0x61, 0x72, 0x64, 0x64, 0x65, 0x63, 0x6b, 0x2e, 0x4d, 0x6f, 0x64,
	0x69, 0x66, 0x79, 0x43, 0x61, 0x72, 0x64, 0x53, 0x69, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x19, 0x2e, 0x63, 0x61, 0x72, 0x64, 0x64, 0x65, 0x63, 0x6b, 0x2e, 0x53, 0x75,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x4c, 0x0a, 0x0e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x61, 0x72, 0x64, 0x53, 0x69, 0x64,
	0x65, 0x12, 0x1d, 0x2e, 0x63, 0x61, 0x72, 0x64, 0x64, 0x65, 0x63, 0x6b, 0x2e, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x57, 0x69, 0x74, 0x68, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x19, 0x2e, 0x63, 0x61, 0x72, 0x64, 0x64, 0x65, 0x63, 0x6b, 0x2e, 0x53, 0x75, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x12, 0x5a,
	0x10, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x63, 0x61, 0x72, 0x64, 0x64, 0x65, 0x63,
	0x6b, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_carddeck_proto_rawDescOnce sync.Once
	file_proto_carddeck_proto_rawDescData = file_proto_carddeck_proto_rawDesc
)

func file_proto_carddeck_proto_rawDescGZIP() []byte {
	file_proto_carddeck_proto_rawDescOnce.Do(func() {
		file_proto_carddeck_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_carddeck_proto_rawDescData)
	})
	return file_proto_carddeck_proto_rawDescData
}

var file_proto_carddeck_proto_msgTypes = make([]protoimpl.MessageInfo, 14)
var file_proto_carddeck_proto_goTypes = []interface{}{
	(*GroupDecksRequest)(nil),     // 0: carddeck.GroupDecksRequest
	(*GroupDecksResponse)(nil),    // 1: carddeck.GroupDecksResponse
	(*Deck)(nil),                  // 2: carddeck.Deck
	(*CreateDeckRequest)(nil),     // 3: carddeck.CreateDeckRequest
	(*IDResponse)(nil),            // 4: carddeck.IDResponse
	(*ModifyDeckRequest)(nil),     // 5: carddeck.ModifyDeckRequest
	(*SuccessResponse)(nil),       // 6: carddeck.SuccessResponse
	(*DeleteWithIDRequest)(nil),   // 7: carddeck.DeleteWithIDRequest
	(*DeckRequest)(nil),           // 8: carddeck.DeckRequest
	(*DeckCardsResponse)(nil),     // 9: carddeck.DeckCardsResponse
	(*Card)(nil),                  // 10: carddeck.Card
	(*CardSide)(nil),              // 11: carddeck.CardSide
	(*CreateCardRequest)(nil),     // 12: carddeck.CreateCardRequest
	(*ModifyCardSideRequest)(nil), // 13: carddeck.ModifyCardSideRequest
}
var file_proto_carddeck_proto_depIdxs = []int32{
	2,  // 0: carddeck.GroupDecksResponse.decks:type_name -> carddeck.Deck
	10, // 1: carddeck.DeckCardsResponse.cards:type_name -> carddeck.Card
	11, // 2: carddeck.Card.sides:type_name -> carddeck.CardSide
	0,  // 3: carddeck.CardDeck.GetGroupDecks:input_type -> carddeck.GroupDecksRequest
	3,  // 4: carddeck.CardDeck.CreateDeck:input_type -> carddeck.CreateDeckRequest
	5,  // 5: carddeck.CardDeck.ModifyDeck:input_type -> carddeck.ModifyDeckRequest
	7,  // 6: carddeck.CardDeck.DeleteDeck:input_type -> carddeck.DeleteWithIDRequest
	8,  // 7: carddeck.CardDeck.GetDeckCards:input_type -> carddeck.DeckRequest
	12, // 8: carddeck.CardDeck.CreateCard:input_type -> carddeck.CreateCardRequest
	7,  // 9: carddeck.CardDeck.DeleteCard:input_type -> carddeck.DeleteWithIDRequest
	13, // 10: carddeck.CardDeck.ModifyCardSide:input_type -> carddeck.ModifyCardSideRequest
	7,  // 11: carddeck.CardDeck.DeleteCardSide:input_type -> carddeck.DeleteWithIDRequest
	1,  // 12: carddeck.CardDeck.GetGroupDecks:output_type -> carddeck.GroupDecksResponse
	4,  // 13: carddeck.CardDeck.CreateDeck:output_type -> carddeck.IDResponse
	6,  // 14: carddeck.CardDeck.ModifyDeck:output_type -> carddeck.SuccessResponse
	6,  // 15: carddeck.CardDeck.DeleteDeck:output_type -> carddeck.SuccessResponse
	9,  // 16: carddeck.CardDeck.GetDeckCards:output_type -> carddeck.DeckCardsResponse
	4,  // 17: carddeck.CardDeck.CreateCard:output_type -> carddeck.IDResponse
	6,  // 18: carddeck.CardDeck.DeleteCard:output_type -> carddeck.SuccessResponse
	6,  // 19: carddeck.CardDeck.ModifyCardSide:output_type -> carddeck.SuccessResponse
	6,  // 20: carddeck.CardDeck.DeleteCardSide:output_type -> carddeck.SuccessResponse
	12, // [12:21] is the sub-list for method output_type
	3,  // [3:12] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_proto_carddeck_proto_init() }
func file_proto_carddeck_proto_init() {
	if File_proto_carddeck_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_carddeck_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GroupDecksRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_carddeck_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GroupDecksResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_carddeck_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Deck); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_carddeck_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateDeckRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_carddeck_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IDResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_carddeck_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ModifyDeckRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_carddeck_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SuccessResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_carddeck_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteWithIDRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_carddeck_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeckRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_carddeck_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeckCardsResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_carddeck_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Card); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_carddeck_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CardSide); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_carddeck_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateCardRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_carddeck_proto_msgTypes[13].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ModifyCardSideRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_proto_carddeck_proto_msgTypes[5].OneofWrappers = []interface{}{}
	file_proto_carddeck_proto_msgTypes[13].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_carddeck_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   14,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_carddeck_proto_goTypes,
		DependencyIndexes: file_proto_carddeck_proto_depIdxs,
		MessageInfos:      file_proto_carddeck_proto_msgTypes,
	}.Build()
	File_proto_carddeck_proto = out.File
	file_proto_carddeck_proto_rawDesc = nil
	file_proto_carddeck_proto_goTypes = nil
	file_proto_carddeck_proto_depIdxs = nil
}
