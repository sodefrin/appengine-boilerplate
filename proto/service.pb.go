// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/service.proto

package proto

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type HealthzResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HealthzResponse) Reset()         { *m = HealthzResponse{} }
func (m *HealthzResponse) String() string { return proto.CompactTextString(m) }
func (*HealthzResponse) ProtoMessage()    {}
func (*HealthzResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c33392ef2c1961ba, []int{0}
}

func (m *HealthzResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HealthzResponse.Unmarshal(m, b)
}
func (m *HealthzResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HealthzResponse.Marshal(b, m, deterministic)
}
func (m *HealthzResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HealthzResponse.Merge(m, src)
}
func (m *HealthzResponse) XXX_Size() int {
	return xxx_messageInfo_HealthzResponse.Size(m)
}
func (m *HealthzResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_HealthzResponse.DiscardUnknown(m)
}

var xxx_messageInfo_HealthzResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*HealthzResponse)(nil), "proto.HealthzResponse")
}

func init() { proto.RegisterFile("proto/service.proto", fileDescriptor_c33392ef2c1961ba) }

var fileDescriptor_c33392ef2c1961ba = []byte{
	// 180 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2e, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0x2f, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0xd5, 0x03, 0xf3, 0x84, 0x58, 0xc1, 0x94,
	0x94, 0x4c, 0x7a, 0x7e, 0x7e, 0x7a, 0x4e, 0xaa, 0x7e, 0x62, 0x41, 0xa6, 0x7e, 0x62, 0x5e, 0x5e,
	0x7e, 0x49, 0x62, 0x49, 0x66, 0x7e, 0x5e, 0x31, 0x44, 0x91, 0x94, 0x34, 0x54, 0x16, 0xcc, 0x4b,
	0x2a, 0x4d, 0xd3, 0x4f, 0xcd, 0x2d, 0x28, 0xa9, 0x84, 0x48, 0x2a, 0x09, 0x72, 0xf1, 0x7b, 0xa4,
	0x26, 0xe6, 0x94, 0x64, 0x54, 0x05, 0xa5, 0x16, 0x17, 0xe4, 0xe7, 0x15, 0xa7, 0x1a, 0x65, 0x71,
	0x89, 0x38, 0x16, 0x14, 0xa4, 0xe6, 0xa5, 0x67, 0xe6, 0xa5, 0x3a, 0xe5, 0x67, 0xe6, 0xa4, 0x16,
	0x15, 0xe4, 0x24, 0x96, 0xa4, 0x0a, 0x05, 0x71, 0xb1, 0x43, 0x95, 0x0a, 0x89, 0xe9, 0x41, 0xcc,
	0xd4, 0x83, 0x99, 0xa9, 0xe7, 0x0a, 0x32, 0x53, 0x4a, 0x0c, 0x22, 0xa0, 0x87, 0x66, 0xa4, 0x92,
	0x78, 0xd3, 0xe5, 0x27, 0x93, 0x99, 0x04, 0x85, 0xf8, 0xc1, 0x4e, 0x2c, 0x33, 0xd4, 0xcf, 0x80,
	0x28, 0x48, 0x62, 0x03, 0xab, 0x37, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x59, 0x84, 0x8e, 0x74,
	0xde, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AppengineBoilerplateClient is the client API for AppengineBoilerplate service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AppengineBoilerplateClient interface {
	Healthz(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*HealthzResponse, error)
}

type appengineBoilerplateClient struct {
	cc *grpc.ClientConn
}

func NewAppengineBoilerplateClient(cc *grpc.ClientConn) AppengineBoilerplateClient {
	return &appengineBoilerplateClient{cc}
}

func (c *appengineBoilerplateClient) Healthz(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*HealthzResponse, error) {
	out := new(HealthzResponse)
	err := c.cc.Invoke(ctx, "/proto.AppengineBoilerplate/Healthz", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AppengineBoilerplateServer is the server API for AppengineBoilerplate service.
type AppengineBoilerplateServer interface {
	Healthz(context.Context, *empty.Empty) (*HealthzResponse, error)
}

// UnimplementedAppengineBoilerplateServer can be embedded to have forward compatible implementations.
type UnimplementedAppengineBoilerplateServer struct {
}

func (*UnimplementedAppengineBoilerplateServer) Healthz(ctx context.Context, req *empty.Empty) (*HealthzResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Healthz not implemented")
}

func RegisterAppengineBoilerplateServer(s *grpc.Server, srv AppengineBoilerplateServer) {
	s.RegisterService(&_AppengineBoilerplate_serviceDesc, srv)
}

func _AppengineBoilerplate_Healthz_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppengineBoilerplateServer).Healthz(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.AppengineBoilerplate/Healthz",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppengineBoilerplateServer).Healthz(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _AppengineBoilerplate_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.AppengineBoilerplate",
	HandlerType: (*AppengineBoilerplateServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Healthz",
			Handler:    _AppengineBoilerplate_Healthz_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/service.proto",
}
