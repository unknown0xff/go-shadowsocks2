// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.21.0
// 	protoc        v3.11.4
// source: api.proto

package api

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type StatsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Password string `protobuf:"bytes,1,opt,name=password,proto3" json:"password,omitempty"` //reserved for server api
}

func (x *StatsRequest) Reset() {
	*x = StatsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatsRequest) ProtoMessage() {}

func (x *StatsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatsRequest.ProtoReflect.Descriptor instead.
func (*StatsRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{0}
}

func (x *StatsRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type StatsReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UploadTraffic   uint64 `protobuf:"varint,1,opt,name=upload_traffic,json=uploadTraffic,proto3" json:"upload_traffic,omitempty"`
	DownloadTraffic uint64 `protobuf:"varint,2,opt,name=download_traffic,json=downloadTraffic,proto3" json:"download_traffic,omitempty"`
	UploadSpeed     uint64 `protobuf:"varint,3,opt,name=upload_speed,json=uploadSpeed,proto3" json:"upload_speed,omitempty"`
	DownloadSpeed   uint64 `protobuf:"varint,4,opt,name=download_speed,json=downloadSpeed,proto3" json:"download_speed,omitempty"`
}

func (x *StatsReply) Reset() {
	*x = StatsReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatsReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatsReply) ProtoMessage() {}

func (x *StatsReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatsReply.ProtoReflect.Descriptor instead.
func (*StatsReply) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{1}
}

func (x *StatsReply) GetUploadTraffic() uint64 {
	if x != nil {
		return x.UploadTraffic
	}
	return 0
}

func (x *StatsReply) GetDownloadTraffic() uint64 {
	if x != nil {
		return x.DownloadTraffic
	}
	return 0
}

func (x *StatsReply) GetUploadSpeed() uint64 {
	if x != nil {
		return x.UploadSpeed
	}
	return 0
}

func (x *StatsReply) GetDownloadSpeed() uint64 {
	if x != nil {
		return x.DownloadSpeed
	}
	return 0
}

var File_api_proto protoreflect.FileDescriptor

var file_api_proto_rawDesc = []byte{
	0x0a, 0x09, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x73, 0x68, 0x61,
	0x64, 0x6f, 0x77, 0x73, 0x6f, 0x63, 0x6b, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x22, 0x2a, 0x0a, 0x0c,
	0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08,
	0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0xa8, 0x01, 0x0a, 0x0a, 0x53, 0x74, 0x61,
	0x74, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x25, 0x0a, 0x0e, 0x75, 0x70, 0x6c, 0x6f, 0x61,
	0x64, 0x5f, 0x74, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x0d, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x54, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x12, 0x29,
	0x0a, 0x10, 0x64, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x74, 0x72, 0x61, 0x66, 0x66,
	0x69, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0f, 0x64, 0x6f, 0x77, 0x6e, 0x6c, 0x6f,
	0x61, 0x64, 0x54, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x12, 0x21, 0x0a, 0x0c, 0x75, 0x70, 0x6c,
	0x6f, 0x61, 0x64, 0x5f, 0x73, 0x70, 0x65, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x0b, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x53, 0x70, 0x65, 0x65, 0x64, 0x12, 0x25, 0x0a, 0x0e,
	0x64, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x73, 0x70, 0x65, 0x65, 0x64, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x0d, 0x64, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x53, 0x70,
	0x65, 0x65, 0x64, 0x32, 0x57, 0x0a, 0x09, 0x53, 0x53, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x4a, 0x0a, 0x0a, 0x51, 0x75, 0x65, 0x72, 0x79, 0x53, 0x74, 0x61, 0x74, 0x73, 0x12, 0x1d,
	0x2e, 0x73, 0x68, 0x61, 0x64, 0x6f, 0x77, 0x73, 0x6f, 0x63, 0x6b, 0x73, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e,
	0x73, 0x68, 0x61, 0x64, 0x6f, 0x77, 0x73, 0x6f, 0x63, 0x6b, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x42, 0x05, 0x5a, 0x03,
	0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_proto_rawDescOnce sync.Once
	file_api_proto_rawDescData = file_api_proto_rawDesc
)

func file_api_proto_rawDescGZIP() []byte {
	file_api_proto_rawDescOnce.Do(func() {
		file_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_proto_rawDescData)
	})
	return file_api_proto_rawDescData
}

var file_api_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_api_proto_goTypes = []interface{}{
	(*StatsRequest)(nil), // 0: shadowsocks.api.StatsRequest
	(*StatsReply)(nil),   // 1: shadowsocks.api.StatsReply
}
var file_api_proto_depIdxs = []int32{
	0, // 0: shadowsocks.api.SSService.QueryStats:input_type -> shadowsocks.api.StatsRequest
	1, // 1: shadowsocks.api.SSService.QueryStats:output_type -> shadowsocks.api.StatsReply
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_proto_init() }
func file_api_proto_init() {
	if File_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StatsRequest); i {
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
		file_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StatsReply); i {
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
			RawDescriptor: file_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_proto_goTypes,
		DependencyIndexes: file_api_proto_depIdxs,
		MessageInfos:      file_api_proto_msgTypes,
	}.Build()
	File_api_proto = out.File
	file_api_proto_rawDesc = nil
	file_api_proto_goTypes = nil
	file_api_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// SSServiceClient is the client API for SSService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SSServiceClient interface {
	QueryStats(ctx context.Context, in *StatsRequest, opts ...grpc.CallOption) (*StatsReply, error)
}

type sSServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSSServiceClient(cc grpc.ClientConnInterface) SSServiceClient {
	return &sSServiceClient{cc}
}

func (c *sSServiceClient) QueryStats(ctx context.Context, in *StatsRequest, opts ...grpc.CallOption) (*StatsReply, error) {
	out := new(StatsReply)
	err := c.cc.Invoke(ctx, "/shadowsocks.api.SSService/QueryStats", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SSServiceServer is the server API for SSService service.
type SSServiceServer interface {
	QueryStats(context.Context, *StatsRequest) (*StatsReply, error)
}

// UnimplementedSSServiceServer can be embedded to have forward compatible implementations.
type UnimplementedSSServiceServer struct {
}

func (*UnimplementedSSServiceServer) QueryStats(context.Context, *StatsRequest) (*StatsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryStats not implemented")
}

func RegisterSSServiceServer(s *grpc.Server, srv SSServiceServer) {
	s.RegisterService(&_SSService_serviceDesc, srv)
}

func _SSService_QueryStats_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StatsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SSServiceServer).QueryStats(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/shadowsocks.api.SSService/QueryStats",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SSServiceServer).QueryStats(ctx, req.(*StatsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _SSService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "shadowsocks.api.SSService",
	HandlerType: (*SSServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "QueryStats",
			Handler:    _SSService_QueryStats_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api.proto",
}
