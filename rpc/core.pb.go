// Code generated by protoc-gen-go.
// source: core.proto
// DO NOT EDIT!

/*
Package ricochet is a generated protocol buffer package.

It is generated from these files:
	core.proto
	network.proto

It has these top-level messages:
	ServerStatusRequest
	ServerStatusReply
	MonitorNetworkRequest
	TorProcessStatus
	TorControlStatus
	TorConnectionStatus
	NetworkStatus
	StartNetworkRequest
	StopNetworkRequest
*/
package ricochet

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type ServerStatusRequest struct {
	RpcVersion int32 `protobuf:"varint,1,opt,name=rpcVersion" json:"rpcVersion,omitempty"`
}

func (m *ServerStatusRequest) Reset()                    { *m = ServerStatusRequest{} }
func (m *ServerStatusRequest) String() string            { return proto.CompactTextString(m) }
func (*ServerStatusRequest) ProtoMessage()               {}
func (*ServerStatusRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type ServerStatusReply struct {
	RpcVersion    int32  `protobuf:"varint,1,opt,name=rpcVersion" json:"rpcVersion,omitempty"`
	ServerVersion string `protobuf:"bytes,2,opt,name=serverVersion" json:"serverVersion,omitempty"`
}

func (m *ServerStatusReply) Reset()                    { *m = ServerStatusReply{} }
func (m *ServerStatusReply) String() string            { return proto.CompactTextString(m) }
func (*ServerStatusReply) ProtoMessage()               {}
func (*ServerStatusReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func init() {
	proto.RegisterType((*ServerStatusRequest)(nil), "ricochet.ServerStatusRequest")
	proto.RegisterType((*ServerStatusReply)(nil), "ricochet.ServerStatusReply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for RicochetCore service

type RicochetCoreClient interface {
	GetServerStatus(ctx context.Context, in *ServerStatusRequest, opts ...grpc.CallOption) (*ServerStatusReply, error)
	MonitorNetwork(ctx context.Context, in *MonitorNetworkRequest, opts ...grpc.CallOption) (RicochetCore_MonitorNetworkClient, error)
	StartNetwork(ctx context.Context, in *StartNetworkRequest, opts ...grpc.CallOption) (*NetworkStatus, error)
	StopNetwork(ctx context.Context, in *StopNetworkRequest, opts ...grpc.CallOption) (*NetworkStatus, error)
}

type ricochetCoreClient struct {
	cc *grpc.ClientConn
}

func NewRicochetCoreClient(cc *grpc.ClientConn) RicochetCoreClient {
	return &ricochetCoreClient{cc}
}

func (c *ricochetCoreClient) GetServerStatus(ctx context.Context, in *ServerStatusRequest, opts ...grpc.CallOption) (*ServerStatusReply, error) {
	out := new(ServerStatusReply)
	err := grpc.Invoke(ctx, "/ricochet.RicochetCore/GetServerStatus", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ricochetCoreClient) MonitorNetwork(ctx context.Context, in *MonitorNetworkRequest, opts ...grpc.CallOption) (RicochetCore_MonitorNetworkClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_RicochetCore_serviceDesc.Streams[0], c.cc, "/ricochet.RicochetCore/MonitorNetwork", opts...)
	if err != nil {
		return nil, err
	}
	x := &ricochetCoreMonitorNetworkClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type RicochetCore_MonitorNetworkClient interface {
	Recv() (*NetworkStatus, error)
	grpc.ClientStream
}

type ricochetCoreMonitorNetworkClient struct {
	grpc.ClientStream
}

func (x *ricochetCoreMonitorNetworkClient) Recv() (*NetworkStatus, error) {
	m := new(NetworkStatus)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *ricochetCoreClient) StartNetwork(ctx context.Context, in *StartNetworkRequest, opts ...grpc.CallOption) (*NetworkStatus, error) {
	out := new(NetworkStatus)
	err := grpc.Invoke(ctx, "/ricochet.RicochetCore/StartNetwork", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ricochetCoreClient) StopNetwork(ctx context.Context, in *StopNetworkRequest, opts ...grpc.CallOption) (*NetworkStatus, error) {
	out := new(NetworkStatus)
	err := grpc.Invoke(ctx, "/ricochet.RicochetCore/StopNetwork", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for RicochetCore service

type RicochetCoreServer interface {
	GetServerStatus(context.Context, *ServerStatusRequest) (*ServerStatusReply, error)
	MonitorNetwork(*MonitorNetworkRequest, RicochetCore_MonitorNetworkServer) error
	StartNetwork(context.Context, *StartNetworkRequest) (*NetworkStatus, error)
	StopNetwork(context.Context, *StopNetworkRequest) (*NetworkStatus, error)
}

func RegisterRicochetCoreServer(s *grpc.Server, srv RicochetCoreServer) {
	s.RegisterService(&_RicochetCore_serviceDesc, srv)
}

func _RicochetCore_GetServerStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ServerStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RicochetCoreServer).GetServerStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ricochet.RicochetCore/GetServerStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RicochetCoreServer).GetServerStatus(ctx, req.(*ServerStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RicochetCore_MonitorNetwork_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(MonitorNetworkRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(RicochetCoreServer).MonitorNetwork(m, &ricochetCoreMonitorNetworkServer{stream})
}

type RicochetCore_MonitorNetworkServer interface {
	Send(*NetworkStatus) error
	grpc.ServerStream
}

type ricochetCoreMonitorNetworkServer struct {
	grpc.ServerStream
}

func (x *ricochetCoreMonitorNetworkServer) Send(m *NetworkStatus) error {
	return x.ServerStream.SendMsg(m)
}

func _RicochetCore_StartNetwork_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartNetworkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RicochetCoreServer).StartNetwork(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ricochet.RicochetCore/StartNetwork",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RicochetCoreServer).StartNetwork(ctx, req.(*StartNetworkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RicochetCore_StopNetwork_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StopNetworkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RicochetCoreServer).StopNetwork(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ricochet.RicochetCore/StopNetwork",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RicochetCoreServer).StopNetwork(ctx, req.(*StopNetworkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _RicochetCore_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ricochet.RicochetCore",
	HandlerType: (*RicochetCoreServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetServerStatus",
			Handler:    _RicochetCore_GetServerStatus_Handler,
		},
		{
			MethodName: "StartNetwork",
			Handler:    _RicochetCore_StartNetwork_Handler,
		},
		{
			MethodName: "StopNetwork",
			Handler:    _RicochetCore_StopNetwork_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "MonitorNetwork",
			Handler:       _RicochetCore_MonitorNetwork_Handler,
			ServerStreams: true,
		},
	},
	Metadata: fileDescriptor0,
}

func init() { proto.RegisterFile("core.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 236 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x4a, 0xce, 0x2f, 0x4a,
	0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x28, 0xca, 0x4c, 0xce, 0x4f, 0xce, 0x48, 0x2d,
	0x91, 0xe2, 0xcd, 0x4b, 0x2d, 0x29, 0xcf, 0x2f, 0xca, 0x86, 0x48, 0x28, 0x99, 0x72, 0x09, 0x07,
	0xa7, 0x16, 0x95, 0xa5, 0x16, 0x05, 0x97, 0x24, 0x96, 0x94, 0x16, 0x07, 0xa5, 0x16, 0x96, 0xa6,
	0x16, 0x97, 0x08, 0xc9, 0x71, 0x71, 0x15, 0x15, 0x24, 0x87, 0xa5, 0x16, 0x15, 0x67, 0xe6, 0xe7,
	0x49, 0x30, 0x2a, 0x30, 0x6a, 0xb0, 0x06, 0x21, 0x89, 0x28, 0x45, 0x72, 0x09, 0xa2, 0x6a, 0x2b,
	0xc8, 0xa9, 0x24, 0xa4, 0x49, 0x48, 0x85, 0x8b, 0xb7, 0x18, 0xac, 0x09, 0xa6, 0x84, 0x09, 0xa8,
	0x84, 0x33, 0x08, 0x55, 0xd0, 0x68, 0x27, 0x13, 0x17, 0x4f, 0x10, 0xd4, 0xb5, 0xce, 0x40, 0x1f,
	0x08, 0xf9, 0x72, 0xf1, 0xbb, 0xa7, 0x96, 0x20, 0x5b, 0x27, 0x24, 0xab, 0x07, 0xf3, 0x8f, 0x1e,
	0x16, 0xd7, 0x4b, 0x49, 0xe3, 0x92, 0x06, 0xb9, 0xd2, 0x87, 0x8b, 0xcf, 0x37, 0x3f, 0x2f, 0xb3,
	0x24, 0xbf, 0xc8, 0x0f, 0x12, 0x12, 0x42, 0xf2, 0x08, 0xe5, 0xa8, 0x32, 0x30, 0xf3, 0xc4, 0x11,
	0x0a, 0xa0, 0x32, 0x10, 0x03, 0x0d, 0x18, 0x85, 0xdc, 0xb8, 0x78, 0x80, 0xec, 0xa2, 0x12, 0x98,
	0x59, 0xc8, 0x2e, 0x43, 0x12, 0x27, 0x64, 0x92, 0x90, 0x0b, 0x17, 0x77, 0x70, 0x49, 0x7e, 0x01,
	0xcc, 0x18, 0x19, 0x64, 0x63, 0xe0, 0xc2, 0x84, 0x4c, 0x49, 0x62, 0x03, 0x47, 0xaa, 0x31, 0x20,
	0x00, 0x00, 0xff, 0xff, 0x19, 0x90, 0x5f, 0x9e, 0xfb, 0x01, 0x00, 0x00,
}