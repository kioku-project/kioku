// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.0
// source: services/notifications/proto/notifications.proto

package notifications

import (
	proto "github.com/kioku-project/kioku/pkg/proto"
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

type PushSubscriptionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID       string            `protobuf:"bytes,1,opt,name=userID,proto3" json:"userID,omitempty"`
	Subscription *PushSubscription `protobuf:"bytes,2,opt,name=subscription,proto3" json:"subscription,omitempty"`
}

func (x *PushSubscriptionRequest) Reset() {
	*x = PushSubscriptionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_notifications_proto_notifications_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PushSubscriptionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PushSubscriptionRequest) ProtoMessage() {}

func (x *PushSubscriptionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_services_notifications_proto_notifications_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PushSubscriptionRequest.ProtoReflect.Descriptor instead.
func (*PushSubscriptionRequest) Descriptor() ([]byte, []int) {
	return file_services_notifications_proto_notifications_proto_rawDescGZIP(), []int{0}
}

func (x *PushSubscriptionRequest) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

func (x *PushSubscriptionRequest) GetSubscription() *PushSubscription {
	if x != nil {
		return x.Subscription
	}
	return nil
}

type PushSubscription struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Endpoint string `protobuf:"bytes,1,opt,name=endpoint,proto3" json:"endpoint,omitempty"`
	Auth     string `protobuf:"bytes,2,opt,name=auth,proto3" json:"auth,omitempty"`
	P256Dh   string `protobuf:"bytes,3,opt,name=p256dh,proto3" json:"p256dh,omitempty"`
}

func (x *PushSubscription) Reset() {
	*x = PushSubscription{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_notifications_proto_notifications_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PushSubscription) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PushSubscription) ProtoMessage() {}

func (x *PushSubscription) ProtoReflect() protoreflect.Message {
	mi := &file_services_notifications_proto_notifications_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PushSubscription.ProtoReflect.Descriptor instead.
func (*PushSubscription) Descriptor() ([]byte, []int) {
	return file_services_notifications_proto_notifications_proto_rawDescGZIP(), []int{1}
}

func (x *PushSubscription) GetEndpoint() string {
	if x != nil {
		return x.Endpoint
	}
	return ""
}

func (x *PushSubscription) GetAuth() string {
	if x != nil {
		return x.Auth
	}
	return ""
}

func (x *PushSubscription) GetP256Dh() string {
	if x != nil {
		return x.P256Dh
	}
	return ""
}

var File_services_notifications_proto_notifications_proto protoreflect.FileDescriptor

var file_services_notifications_proto_notifications_proto_rawDesc = []byte{
	0x0a, 0x30, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x6e, 0x6f, 0x74, 0x69, 0x66,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6e,
	0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0d, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x1a, 0x16, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x76, 0x0a, 0x17, 0x50, 0x75, 0x73,
	0x68, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x43, 0x0a, 0x0c,
	0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x50, 0x75, 0x73, 0x68, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x0c, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x22, 0x5a, 0x0a, 0x10, 0x50, 0x75, 0x73, 0x68, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x61, 0x75, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x61, 0x75, 0x74, 0x68, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x32, 0x35, 0x36, 0x64, 0x68, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x32, 0x35, 0x36, 0x64, 0x68, 0x32, 0x54, 0x0a,
	0x0d, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x43,
	0x0a, 0x06, 0x45, 0x6e, 0x72, 0x6f, 0x6c, 0x6c, 0x12, 0x26, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x50, 0x75, 0x73, 0x68, 0x53, 0x75, 0x62,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x0f, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x22, 0x00, 0x42, 0x4b, 0x5a, 0x49, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x6b, 0x69, 0x6f, 0x6b, 0x75, 0x2d, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2f,
	0x6b, 0x69, 0x6f, 0x6b, 0x75, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x6e,
	0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x3b, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_services_notifications_proto_notifications_proto_rawDescOnce sync.Once
	file_services_notifications_proto_notifications_proto_rawDescData = file_services_notifications_proto_notifications_proto_rawDesc
)

func file_services_notifications_proto_notifications_proto_rawDescGZIP() []byte {
	file_services_notifications_proto_notifications_proto_rawDescOnce.Do(func() {
		file_services_notifications_proto_notifications_proto_rawDescData = protoimpl.X.CompressGZIP(file_services_notifications_proto_notifications_proto_rawDescData)
	})
	return file_services_notifications_proto_notifications_proto_rawDescData
}

var file_services_notifications_proto_notifications_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_services_notifications_proto_notifications_proto_goTypes = []interface{}{
	(*PushSubscriptionRequest)(nil), // 0: notifications.PushSubscriptionRequest
	(*PushSubscription)(nil),        // 1: notifications.PushSubscription
	(*proto.Success)(nil),           // 2: common.Success
}
var file_services_notifications_proto_notifications_proto_depIdxs = []int32{
	1, // 0: notifications.PushSubscriptionRequest.subscription:type_name -> notifications.PushSubscription
	0, // 1: notifications.Notifications.Enroll:input_type -> notifications.PushSubscriptionRequest
	2, // 2: notifications.Notifications.Enroll:output_type -> common.Success
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_services_notifications_proto_notifications_proto_init() }
func file_services_notifications_proto_notifications_proto_init() {
	if File_services_notifications_proto_notifications_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_services_notifications_proto_notifications_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PushSubscriptionRequest); i {
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
		file_services_notifications_proto_notifications_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PushSubscription); i {
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
			RawDescriptor: file_services_notifications_proto_notifications_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_services_notifications_proto_notifications_proto_goTypes,
		DependencyIndexes: file_services_notifications_proto_notifications_proto_depIdxs,
		MessageInfos:      file_services_notifications_proto_notifications_proto_msgTypes,
	}.Build()
	File_services_notifications_proto_notifications_proto = out.File
	file_services_notifications_proto_notifications_proto_rawDesc = nil
	file_services_notifications_proto_notifications_proto_goTypes = nil
	file_services_notifications_proto_notifications_proto_depIdxs = nil
}
