// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        (unknown)
// source: grpc/multiplexer/v1/discovery.proto

package discovery

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

// List of possible event of WatchGRPCServerResponse
type WatchGRPCServerResponse_Event int32

const (
	// Unspecified event, probably an issue with the request or the gRPC multiplexer.
	WatchGRPCServerResponse_EVENT_UNSPECIFIED WatchGRPCServerResponse_Event = 0
	// The multiplexer client has connected a new gRPC server.
	WatchGRPCServerResponse_EVENT_CONNECTED WatchGRPCServerResponse_Event = 1
	// The multiplexer client has disconnected an existing gRPC server.
	WatchGRPCServerResponse_EVENT_DISCONNECTED WatchGRPCServerResponse_Event = 2
)

// Enum value maps for WatchGRPCServerResponse_Event.
var (
	WatchGRPCServerResponse_Event_name = map[int32]string{
		0: "EVENT_UNSPECIFIED",
		1: "EVENT_CONNECTED",
		2: "EVENT_DISCONNECTED",
	}
	WatchGRPCServerResponse_Event_value = map[string]int32{
		"EVENT_UNSPECIFIED":  0,
		"EVENT_CONNECTED":    1,
		"EVENT_DISCONNECTED": 2,
	}
)

func (x WatchGRPCServerResponse_Event) Enum() *WatchGRPCServerResponse_Event {
	p := new(WatchGRPCServerResponse_Event)
	*p = x
	return p
}

func (x WatchGRPCServerResponse_Event) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (WatchGRPCServerResponse_Event) Descriptor() protoreflect.EnumDescriptor {
	return file_grpc_multiplexer_v1_discovery_proto_enumTypes[0].Descriptor()
}

func (WatchGRPCServerResponse_Event) Type() protoreflect.EnumType {
	return &file_grpc_multiplexer_v1_discovery_proto_enumTypes[0]
}

func (x WatchGRPCServerResponse_Event) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use WatchGRPCServerResponse_Event.Descriptor instead.
func (WatchGRPCServerResponse_Event) EnumDescriptor() ([]byte, []int) {
	return file_grpc_multiplexer_v1_discovery_proto_rawDescGZIP(), []int{3, 0}
}

// ListGRPCServerRequest is the request message for ListGRPCServer method.
type ListGRPCServerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListGRPCServerRequest) Reset() {
	*x = ListGRPCServerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_multiplexer_v1_discovery_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListGRPCServerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListGRPCServerRequest) ProtoMessage() {}

func (x *ListGRPCServerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_multiplexer_v1_discovery_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListGRPCServerRequest.ProtoReflect.Descriptor instead.
func (*ListGRPCServerRequest) Descriptor() ([]byte, []int) {
	return file_grpc_multiplexer_v1_discovery_proto_rawDescGZIP(), []int{0}
}

// ListGRPCServerResponse is the response message for ListGRPCServer method.
type ListGRPCServerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The list of gRPC server client ID (represented by their multiplexer clients).
	Servers []string `protobuf:"bytes,1,rep,name=servers,proto3" json:"servers,omitempty"`
}

func (x *ListGRPCServerResponse) Reset() {
	*x = ListGRPCServerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_multiplexer_v1_discovery_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListGRPCServerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListGRPCServerResponse) ProtoMessage() {}

func (x *ListGRPCServerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_multiplexer_v1_discovery_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListGRPCServerResponse.ProtoReflect.Descriptor instead.
func (*ListGRPCServerResponse) Descriptor() ([]byte, []int) {
	return file_grpc_multiplexer_v1_discovery_proto_rawDescGZIP(), []int{1}
}

func (x *ListGRPCServerResponse) GetServers() []string {
	if x != nil {
		return x.Servers
	}
	return nil
}

// WatchGRPCServerRequest is the request message for WatchGRPCServer method.
type WatchGRPCServerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *WatchGRPCServerRequest) Reset() {
	*x = WatchGRPCServerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_multiplexer_v1_discovery_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WatchGRPCServerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WatchGRPCServerRequest) ProtoMessage() {}

func (x *WatchGRPCServerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_multiplexer_v1_discovery_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WatchGRPCServerRequest.ProtoReflect.Descriptor instead.
func (*WatchGRPCServerRequest) Descriptor() ([]byte, []int) {
	return file_grpc_multiplexer_v1_discovery_proto_rawDescGZIP(), []int{2}
}

// WatchGRPCServerResponse is the response message for WatchGRPCServer method.
type WatchGRPCServerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The client ID that represents the gRPC server.
	Server string `protobuf:"bytes,1,opt,name=server,proto3" json:"server,omitempty"`
	// Action event related to the gRPC server.
	Event WatchGRPCServerResponse_Event `protobuf:"varint,2,opt,name=event,proto3,enum=grpc.multiplexer.v1.WatchGRPCServerResponse_Event" json:"event,omitempty"`
}

func (x *WatchGRPCServerResponse) Reset() {
	*x = WatchGRPCServerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_multiplexer_v1_discovery_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WatchGRPCServerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WatchGRPCServerResponse) ProtoMessage() {}

func (x *WatchGRPCServerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_multiplexer_v1_discovery_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WatchGRPCServerResponse.ProtoReflect.Descriptor instead.
func (*WatchGRPCServerResponse) Descriptor() ([]byte, []int) {
	return file_grpc_multiplexer_v1_discovery_proto_rawDescGZIP(), []int{3}
}

func (x *WatchGRPCServerResponse) GetServer() string {
	if x != nil {
		return x.Server
	}
	return ""
}

func (x *WatchGRPCServerResponse) GetEvent() WatchGRPCServerResponse_Event {
	if x != nil {
		return x.Event
	}
	return WatchGRPCServerResponse_EVENT_UNSPECIFIED
}

var File_grpc_multiplexer_v1_discovery_proto protoreflect.FileDescriptor

var file_grpc_multiplexer_v1_discovery_proto_rawDesc = []byte{
	0x0a, 0x23, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x65, 0x78,
	0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x6d, 0x75, 0x6c, 0x74,
	0x69, 0x70, 0x6c, 0x65, 0x78, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x22, 0x17, 0x0a, 0x15, 0x4c, 0x69,
	0x73, 0x74, 0x47, 0x52, 0x50, 0x43, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x22, 0x32, 0x0a, 0x16, 0x4c, 0x69, 0x73, 0x74, 0x47, 0x52, 0x50, 0x43, 0x53,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x73, 0x22, 0x18, 0x0a, 0x16, 0x57, 0x61, 0x74, 0x63, 0x68,
	0x47, 0x52, 0x50, 0x43, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x22, 0xc8, 0x01, 0x0a, 0x17, 0x57, 0x61, 0x74, 0x63, 0x68, 0x47, 0x52, 0x50, 0x43, 0x53,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x48, 0x0a, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x32, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x6d, 0x75, 0x6c, 0x74,
	0x69, 0x70, 0x6c, 0x65, 0x78, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x57, 0x61, 0x74, 0x63, 0x68,
	0x47, 0x52, 0x50, 0x43, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x22,
	0x4b, 0x0a, 0x05, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x15, 0x0a, 0x11, 0x45, 0x56, 0x45, 0x4e,
	0x54, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12,
	0x13, 0x0a, 0x0f, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x43, 0x4f, 0x4e, 0x4e, 0x45, 0x43, 0x54,
	0x45, 0x44, 0x10, 0x01, 0x12, 0x16, 0x0a, 0x12, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x44, 0x49,
	0x53, 0x43, 0x4f, 0x4e, 0x4e, 0x45, 0x43, 0x54, 0x45, 0x44, 0x10, 0x02, 0x32, 0xed, 0x01, 0x0a,
	0x10, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x69, 0x0a, 0x0e, 0x4c, 0x69, 0x73, 0x74, 0x47, 0x52, 0x50, 0x43, 0x53, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x12, 0x2a, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x6d, 0x75, 0x6c, 0x74, 0x69,
	0x70, 0x6c, 0x65, 0x78, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x47, 0x52,
	0x50, 0x43, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x2b, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x65, 0x78,
	0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x47, 0x52, 0x50, 0x43, 0x53, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x6e, 0x0a, 0x0f,
	0x57, 0x61, 0x74, 0x63, 0x68, 0x47, 0x52, 0x50, 0x43, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12,
	0x2b, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x65, 0x78,
	0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x57, 0x61, 0x74, 0x63, 0x68, 0x47, 0x52, 0x50, 0x43, 0x53,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2c, 0x2e, 0x67,
	0x72, 0x70, 0x63, 0x2e, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x65, 0x78, 0x65, 0x72, 0x2e,
	0x76, 0x31, 0x2e, 0x57, 0x61, 0x74, 0x63, 0x68, 0x47, 0x52, 0x50, 0x43, 0x53, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x30, 0x01, 0x42, 0x44, 0x5a, 0x2f,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x66, 0x75, 0x68, 0x72, 0x6d,
	0x61, 0x6e, 0x6e, 0x62, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x70,
	0x6c, 0x65, 0x78, 0x65, 0x72, 0x2f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0xaa,
	0x02, 0x10, 0x47, 0x52, 0x50, 0x43, 0x2e, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x65, 0x78,
	0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_grpc_multiplexer_v1_discovery_proto_rawDescOnce sync.Once
	file_grpc_multiplexer_v1_discovery_proto_rawDescData = file_grpc_multiplexer_v1_discovery_proto_rawDesc
)

func file_grpc_multiplexer_v1_discovery_proto_rawDescGZIP() []byte {
	file_grpc_multiplexer_v1_discovery_proto_rawDescOnce.Do(func() {
		file_grpc_multiplexer_v1_discovery_proto_rawDescData = protoimpl.X.CompressGZIP(file_grpc_multiplexer_v1_discovery_proto_rawDescData)
	})
	return file_grpc_multiplexer_v1_discovery_proto_rawDescData
}

var file_grpc_multiplexer_v1_discovery_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_grpc_multiplexer_v1_discovery_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_grpc_multiplexer_v1_discovery_proto_goTypes = []interface{}{
	(WatchGRPCServerResponse_Event)(0), // 0: grpc.multiplexer.v1.WatchGRPCServerResponse.Event
	(*ListGRPCServerRequest)(nil),      // 1: grpc.multiplexer.v1.ListGRPCServerRequest
	(*ListGRPCServerResponse)(nil),     // 2: grpc.multiplexer.v1.ListGRPCServerResponse
	(*WatchGRPCServerRequest)(nil),     // 3: grpc.multiplexer.v1.WatchGRPCServerRequest
	(*WatchGRPCServerResponse)(nil),    // 4: grpc.multiplexer.v1.WatchGRPCServerResponse
}
var file_grpc_multiplexer_v1_discovery_proto_depIdxs = []int32{
	0, // 0: grpc.multiplexer.v1.WatchGRPCServerResponse.event:type_name -> grpc.multiplexer.v1.WatchGRPCServerResponse.Event
	1, // 1: grpc.multiplexer.v1.DiscoveryService.ListGRPCServer:input_type -> grpc.multiplexer.v1.ListGRPCServerRequest
	3, // 2: grpc.multiplexer.v1.DiscoveryService.WatchGRPCServer:input_type -> grpc.multiplexer.v1.WatchGRPCServerRequest
	2, // 3: grpc.multiplexer.v1.DiscoveryService.ListGRPCServer:output_type -> grpc.multiplexer.v1.ListGRPCServerResponse
	4, // 4: grpc.multiplexer.v1.DiscoveryService.WatchGRPCServer:output_type -> grpc.multiplexer.v1.WatchGRPCServerResponse
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_grpc_multiplexer_v1_discovery_proto_init() }
func file_grpc_multiplexer_v1_discovery_proto_init() {
	if File_grpc_multiplexer_v1_discovery_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_grpc_multiplexer_v1_discovery_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListGRPCServerRequest); i {
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
		file_grpc_multiplexer_v1_discovery_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListGRPCServerResponse); i {
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
		file_grpc_multiplexer_v1_discovery_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WatchGRPCServerRequest); i {
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
		file_grpc_multiplexer_v1_discovery_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WatchGRPCServerResponse); i {
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
			RawDescriptor: file_grpc_multiplexer_v1_discovery_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_grpc_multiplexer_v1_discovery_proto_goTypes,
		DependencyIndexes: file_grpc_multiplexer_v1_discovery_proto_depIdxs,
		EnumInfos:         file_grpc_multiplexer_v1_discovery_proto_enumTypes,
		MessageInfos:      file_grpc_multiplexer_v1_discovery_proto_msgTypes,
	}.Build()
	File_grpc_multiplexer_v1_discovery_proto = out.File
	file_grpc_multiplexer_v1_discovery_proto_rawDesc = nil
	file_grpc_multiplexer_v1_discovery_proto_goTypes = nil
	file_grpc_multiplexer_v1_discovery_proto_depIdxs = nil
}
