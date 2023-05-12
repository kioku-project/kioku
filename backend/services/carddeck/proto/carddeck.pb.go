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

type CreateCardRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID       uint64 `protobuf:"varint,1,opt,name=userID,proto3" json:"userID,omitempty"`
	DeckPublicID string `protobuf:"bytes,2,opt,name=deckPublicID,proto3" json:"deckPublicID,omitempty"`
	Frontside    string `protobuf:"bytes,3,opt,name=frontside,proto3" json:"frontside,omitempty"`
	Backside     string `protobuf:"bytes,4,opt,name=backside,proto3" json:"backside,omitempty"`
}

func (x *CreateCardRequest) Reset() {
	*x = CreateCardRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_carddeck_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateCardRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCardRequest) ProtoMessage() {}

func (x *CreateCardRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use CreateCardRequest.ProtoReflect.Descriptor instead.
func (*CreateCardRequest) Descriptor() ([]byte, []int) {
	return file_proto_carddeck_proto_rawDescGZIP(), []int{0}
}

func (x *CreateCardRequest) GetUserID() uint64 {
	if x != nil {
		return x.UserID
	}
	return 0
}

func (x *CreateCardRequest) GetDeckPublicID() string {
	if x != nil {
		return x.DeckPublicID
	}
	return ""
}

func (x *CreateCardRequest) GetFrontside() string {
	if x != nil {
		return x.Frontside
	}
	return ""
}

func (x *CreateCardRequest) GetBackside() string {
	if x != nil {
		return x.Backside
	}
	return ""
}

type CreateDeckRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID        uint64 `protobuf:"varint,1,opt,name=userID,proto3" json:"userID,omitempty"`
	GroupPublicID string `protobuf:"bytes,2,opt,name=groupPublicID,proto3" json:"groupPublicID,omitempty"`
	DeckName      string `protobuf:"bytes,3,opt,name=deckName,proto3" json:"deckName,omitempty"`
}

func (x *CreateDeckRequest) Reset() {
	*x = CreateDeckRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_carddeck_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateDeckRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateDeckRequest) ProtoMessage() {}

func (x *CreateDeckRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use CreateDeckRequest.ProtoReflect.Descriptor instead.
func (*CreateDeckRequest) Descriptor() ([]byte, []int) {
	return file_proto_carddeck_proto_rawDescGZIP(), []int{1}
}

func (x *CreateDeckRequest) GetUserID() uint64 {
	if x != nil {
		return x.UserID
	}
	return 0
}

func (x *CreateDeckRequest) GetGroupPublicID() string {
	if x != nil {
		return x.GroupPublicID
	}
	return ""
}

func (x *CreateDeckRequest) GetDeckName() string {
	if x != nil {
		return x.DeckName
	}
	return ""
}

type PublicIDResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PublicID string `protobuf:"bytes,1,opt,name=publicID,proto3" json:"publicID,omitempty"`
}

func (x *PublicIDResponse) Reset() {
	*x = PublicIDResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_carddeck_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PublicIDResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublicIDResponse) ProtoMessage() {}

func (x *PublicIDResponse) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use PublicIDResponse.ProtoReflect.Descriptor instead.
func (*PublicIDResponse) Descriptor() ([]byte, []int) {
	return file_proto_carddeck_proto_rawDescGZIP(), []int{2}
}

func (x *PublicIDResponse) GetPublicID() string {
	if x != nil {
		return x.PublicID
	}
	return ""
}

type DeckCardsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID       uint64 `protobuf:"varint,1,opt,name=userID,proto3" json:"userID,omitempty"`
	DeckPublicID string `protobuf:"bytes,2,opt,name=deckPublicID,proto3" json:"deckPublicID,omitempty"`
}

func (x *DeckCardsRequest) Reset() {
	*x = DeckCardsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_carddeck_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeckCardsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeckCardsRequest) ProtoMessage() {}

func (x *DeckCardsRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use DeckCardsRequest.ProtoReflect.Descriptor instead.
func (*DeckCardsRequest) Descriptor() ([]byte, []int) {
	return file_proto_carddeck_proto_rawDescGZIP(), []int{3}
}

func (x *DeckCardsRequest) GetUserID() uint64 {
	if x != nil {
		return x.UserID
	}
	return 0
}

func (x *DeckCardsRequest) GetDeckPublicID() string {
	if x != nil {
		return x.DeckPublicID
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
		mi := &file_proto_carddeck_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeckCardsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeckCardsResponse) ProtoMessage() {}

func (x *DeckCardsResponse) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use DeckCardsResponse.ProtoReflect.Descriptor instead.
func (*DeckCardsResponse) Descriptor() ([]byte, []int) {
	return file_proto_carddeck_proto_rawDescGZIP(), []int{4}
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

	CardPublicID string `protobuf:"bytes,1,opt,name=cardPublicID,proto3" json:"cardPublicID,omitempty"`
	Frontside    string `protobuf:"bytes,2,opt,name=frontside,proto3" json:"frontside,omitempty"`
	Backside     string `protobuf:"bytes,3,opt,name=backside,proto3" json:"backside,omitempty"`
}

func (x *Card) Reset() {
	*x = Card{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_carddeck_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Card) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Card) ProtoMessage() {}

func (x *Card) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use Card.ProtoReflect.Descriptor instead.
func (*Card) Descriptor() ([]byte, []int) {
	return file_proto_carddeck_proto_rawDescGZIP(), []int{5}
}

func (x *Card) GetCardPublicID() string {
	if x != nil {
		return x.CardPublicID
	}
	return ""
}

func (x *Card) GetFrontside() string {
	if x != nil {
		return x.Frontside
	}
	return ""
}

func (x *Card) GetBackside() string {
	if x != nil {
		return x.Backside
	}
	return ""
}

type GroupDecksRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID        uint64 `protobuf:"varint,1,opt,name=userID,proto3" json:"userID,omitempty"`
	GroupPublicID string `protobuf:"bytes,2,opt,name=groupPublicID,proto3" json:"groupPublicID,omitempty"`
}

func (x *GroupDecksRequest) Reset() {
	*x = GroupDecksRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_carddeck_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GroupDecksRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GroupDecksRequest) ProtoMessage() {}

func (x *GroupDecksRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use GroupDecksRequest.ProtoReflect.Descriptor instead.
func (*GroupDecksRequest) Descriptor() ([]byte, []int) {
	return file_proto_carddeck_proto_rawDescGZIP(), []int{6}
}

func (x *GroupDecksRequest) GetUserID() uint64 {
	if x != nil {
		return x.UserID
	}
	return 0
}

func (x *GroupDecksRequest) GetGroupPublicID() string {
	if x != nil {
		return x.GroupPublicID
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
		mi := &file_proto_carddeck_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GroupDecksResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GroupDecksResponse) ProtoMessage() {}

func (x *GroupDecksResponse) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use GroupDecksResponse.ProtoReflect.Descriptor instead.
func (*GroupDecksResponse) Descriptor() ([]byte, []int) {
	return file_proto_carddeck_proto_rawDescGZIP(), []int{7}
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

	DeckPublicID string `protobuf:"bytes,1,opt,name=deckPublicID,proto3" json:"deckPublicID,omitempty"`
	DeckName     string `protobuf:"bytes,2,opt,name=deckName,proto3" json:"deckName,omitempty"`
}

func (x *Deck) Reset() {
	*x = Deck{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_carddeck_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Deck) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Deck) ProtoMessage() {}

func (x *Deck) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use Deck.ProtoReflect.Descriptor instead.
func (*Deck) Descriptor() ([]byte, []int) {
	return file_proto_carddeck_proto_rawDescGZIP(), []int{8}
}

func (x *Deck) GetDeckPublicID() string {
	if x != nil {
		return x.DeckPublicID
	}
	return ""
}

func (x *Deck) GetDeckName() string {
	if x != nil {
		return x.DeckName
	}
	return ""
}

var File_proto_carddeck_proto protoreflect.FileDescriptor

var file_proto_carddeck_proto_rawDesc = []byte{
	0x0a, 0x14, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x61, 0x72, 0x64, 0x64, 0x65, 0x63, 0x6b,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x63, 0x61, 0x72, 0x64, 0x64, 0x65, 0x63, 0x6b,
	0x22, 0x89, 0x01, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x61, 0x72, 0x64, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x22,
	0x0a, 0x0c, 0x64, 0x65, 0x63, 0x6b, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x49, 0x44, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x64, 0x65, 0x63, 0x6b, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63,
	0x49, 0x44, 0x12, 0x1c, 0x0a, 0x09, 0x66, 0x72, 0x6f, 0x6e, 0x74, 0x73, 0x69, 0x64, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x72, 0x6f, 0x6e, 0x74, 0x73, 0x69, 0x64, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x62, 0x61, 0x63, 0x6b, 0x73, 0x69, 0x64, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x62, 0x61, 0x63, 0x6b, 0x73, 0x69, 0x64, 0x65, 0x22, 0x6d, 0x0a, 0x11,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x24, 0x0a, 0x0d, 0x67, 0x72, 0x6f,
	0x75, 0x70, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0d, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x49, 0x44, 0x12,
	0x1a, 0x0a, 0x08, 0x64, 0x65, 0x63, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x64, 0x65, 0x63, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x2e, 0x0a, 0x10, 0x50,
	0x75, 0x62, 0x6c, 0x69, 0x63, 0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x49, 0x44, 0x22, 0x4e, 0x0a, 0x10, 0x44,
	0x65, 0x63, 0x6b, 0x43, 0x61, 0x72, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x22, 0x0a, 0x0c, 0x64, 0x65, 0x63, 0x6b, 0x50,
	0x75, 0x62, 0x6c, 0x69, 0x63, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x64,
	0x65, 0x63, 0x6b, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x49, 0x44, 0x22, 0x39, 0x0a, 0x11, 0x44,
	0x65, 0x63, 0x6b, 0x43, 0x61, 0x72, 0x64, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x24, 0x0a, 0x05, 0x63, 0x61, 0x72, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x0e, 0x2e, 0x63, 0x61, 0x72, 0x64, 0x64, 0x65, 0x63, 0x6b, 0x2e, 0x43, 0x61, 0x72, 0x64, 0x52,
	0x05, 0x63, 0x61, 0x72, 0x64, 0x73, 0x22, 0x64, 0x0a, 0x04, 0x43, 0x61, 0x72, 0x64, 0x12, 0x22,
	0x0a, 0x0c, 0x63, 0x61, 0x72, 0x64, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x49, 0x44, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x61, 0x72, 0x64, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63,
	0x49, 0x44, 0x12, 0x1c, 0x0a, 0x09, 0x66, 0x72, 0x6f, 0x6e, 0x74, 0x73, 0x69, 0x64, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x72, 0x6f, 0x6e, 0x74, 0x73, 0x69, 0x64, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x62, 0x61, 0x63, 0x6b, 0x73, 0x69, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x62, 0x61, 0x63, 0x6b, 0x73, 0x69, 0x64, 0x65, 0x22, 0x51, 0x0a, 0x11,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x44, 0x65, 0x63, 0x6b, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x24, 0x0a, 0x0d, 0x67, 0x72, 0x6f,
	0x75, 0x70, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0d, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x49, 0x44, 0x22,
	0x3a, 0x0a, 0x12, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x44, 0x65, 0x63, 0x6b, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x24, 0x0a, 0x05, 0x64, 0x65, 0x63, 0x6b, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x63, 0x61, 0x72, 0x64, 0x64, 0x65, 0x63, 0x6b, 0x2e,
	0x44, 0x65, 0x63, 0x6b, 0x52, 0x05, 0x64, 0x65, 0x63, 0x6b, 0x73, 0x22, 0x46, 0x0a, 0x04, 0x44,
	0x65, 0x63, 0x6b, 0x12, 0x22, 0x0a, 0x0c, 0x64, 0x65, 0x63, 0x6b, 0x50, 0x75, 0x62, 0x6c, 0x69,
	0x63, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x64, 0x65, 0x63, 0x6b, 0x50,
	0x75, 0x62, 0x6c, 0x69, 0x63, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x65, 0x63, 0x6b, 0x4e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x65, 0x63, 0x6b, 0x4e,
	0x61, 0x6d, 0x65, 0x32, 0xb5, 0x02, 0x0a, 0x08, 0x43, 0x61, 0x72, 0x64, 0x64, 0x65, 0x63, 0x6b,
	0x12, 0x47, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x61, 0x72, 0x64, 0x12, 0x1b,
	0x2e, 0x63, 0x61, 0x72, 0x64, 0x64, 0x65, 0x63, 0x6b, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x43, 0x61, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x63, 0x61,
	0x72, 0x64, 0x64, 0x65, 0x63, 0x6b, 0x2e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x49, 0x44, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x47, 0x0a, 0x0a, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x44, 0x65, 0x63, 0x6b, 0x12, 0x1b, 0x2e, 0x63, 0x61, 0x72, 0x64, 0x64, 0x65,
	0x63, 0x6b, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x63, 0x61, 0x72, 0x64, 0x64, 0x65, 0x63, 0x6b, 0x2e,
	0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x49, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x44, 0x65, 0x63, 0x6b, 0x43, 0x61, 0x72,
	0x64, 0x73, 0x12, 0x1a, 0x2e, 0x63, 0x61, 0x72, 0x64, 0x64, 0x65, 0x63, 0x6b, 0x2e, 0x44, 0x65,
	0x63, 0x6b, 0x43, 0x61, 0x72, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b,
	0x2e, 0x63, 0x61, 0x72, 0x64, 0x64, 0x65, 0x63, 0x6b, 0x2e, 0x44, 0x65, 0x63, 0x6b, 0x43, 0x61,
	0x72, 0x64, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4c, 0x0a,
	0x0d, 0x47, 0x65, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x44, 0x65, 0x63, 0x6b, 0x73, 0x12, 0x1b,
	0x2e, 0x63, 0x61, 0x72, 0x64, 0x64, 0x65, 0x63, 0x6b, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x44,
	0x65, 0x63, 0x6b, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x63, 0x61,
	0x72, 0x64, 0x64, 0x65, 0x63, 0x6b, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x44, 0x65, 0x63, 0x6b,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x12, 0x5a, 0x10, 0x2e,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x63, 0x61, 0x72, 0x64, 0x64, 0x65, 0x63, 0x6b, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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

var file_proto_carddeck_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_proto_carddeck_proto_goTypes = []interface{}{
	(*CreateCardRequest)(nil),  // 0: carddeck.CreateCardRequest
	(*CreateDeckRequest)(nil),  // 1: carddeck.CreateDeckRequest
	(*PublicIDResponse)(nil),   // 2: carddeck.PublicIDResponse
	(*DeckCardsRequest)(nil),   // 3: carddeck.DeckCardsRequest
	(*DeckCardsResponse)(nil),  // 4: carddeck.DeckCardsResponse
	(*Card)(nil),               // 5: carddeck.Card
	(*GroupDecksRequest)(nil),  // 6: carddeck.GroupDecksRequest
	(*GroupDecksResponse)(nil), // 7: carddeck.GroupDecksResponse
	(*Deck)(nil),               // 8: carddeck.Deck
}
var file_proto_carddeck_proto_depIdxs = []int32{
	5, // 0: carddeck.DeckCardsResponse.cards:type_name -> carddeck.Card
	8, // 1: carddeck.GroupDecksResponse.decks:type_name -> carddeck.Deck
	0, // 2: carddeck.Carddeck.CreateCard:input_type -> carddeck.CreateCardRequest
	1, // 3: carddeck.Carddeck.CreateDeck:input_type -> carddeck.CreateDeckRequest
	3, // 4: carddeck.Carddeck.GetDeckCards:input_type -> carddeck.DeckCardsRequest
	6, // 5: carddeck.Carddeck.GetGroupDecks:input_type -> carddeck.GroupDecksRequest
	2, // 6: carddeck.Carddeck.CreateCard:output_type -> carddeck.PublicIDResponse
	2, // 7: carddeck.Carddeck.CreateDeck:output_type -> carddeck.PublicIDResponse
	4, // 8: carddeck.Carddeck.GetDeckCards:output_type -> carddeck.DeckCardsResponse
	7, // 9: carddeck.Carddeck.GetGroupDecks:output_type -> carddeck.GroupDecksResponse
	6, // [6:10] is the sub-list for method output_type
	2, // [2:6] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_proto_carddeck_proto_init() }
func file_proto_carddeck_proto_init() {
	if File_proto_carddeck_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_carddeck_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_proto_carddeck_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_proto_carddeck_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PublicIDResponse); i {
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
			switch v := v.(*DeckCardsRequest); i {
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
		file_proto_carddeck_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
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
		file_proto_carddeck_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
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
		file_proto_carddeck_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
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
		file_proto_carddeck_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
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
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_carddeck_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
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