// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pb/heartbeat.proto

/*
Package proto is a generated protocol buffer package.

It is generated from these files:
	pb/heartbeat.proto

It has these top-level messages:
	HeartbeatRequest
	HeartbeatResponse
*/
package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto1.ProtoPackageIsVersion2 // please upgrade the proto package

type HeartbeatRequest struct {
	AppID  string `protobuf:"bytes,1,opt,name=appID" json:"appID,omitempty"`
	NodeID int32  `protobuf:"varint,2,opt,name=nodeID" json:"nodeID,omitempty"`
}

func (m *HeartbeatRequest) Reset()                    { *m = HeartbeatRequest{} }
func (m *HeartbeatRequest) String() string            { return proto1.CompactTextString(m) }
func (*HeartbeatRequest) ProtoMessage()               {}
func (*HeartbeatRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *HeartbeatRequest) GetAppID() string {
	if m != nil {
		return m.AppID
	}
	return ""
}

func (m *HeartbeatRequest) GetNodeID() int32 {
	if m != nil {
		return m.NodeID
	}
	return 0
}

type HeartbeatResponse struct {
	AppID  string `protobuf:"bytes,1,opt,name=appID" json:"appID,omitempty"`
	NodeID int32  `protobuf:"varint,2,opt,name=nodeID" json:"nodeID,omitempty"`
}

func (m *HeartbeatResponse) Reset()                    { *m = HeartbeatResponse{} }
func (m *HeartbeatResponse) String() string            { return proto1.CompactTextString(m) }
func (*HeartbeatResponse) ProtoMessage()               {}
func (*HeartbeatResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *HeartbeatResponse) GetAppID() string {
	if m != nil {
		return m.AppID
	}
	return ""
}

func (m *HeartbeatResponse) GetNodeID() int32 {
	if m != nil {
		return m.NodeID
	}
	return 0
}

func init() {
	proto1.RegisterType((*HeartbeatRequest)(nil), "proto.HeartbeatRequest")
	proto1.RegisterType((*HeartbeatResponse)(nil), "proto.HeartbeatResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for HeartbeatService service

type HeartbeatServiceClient interface {
	Heartbeat(ctx context.Context, in *HeartbeatRequest, opts ...grpc.CallOption) (*HeartbeatResponse, error)
}

type heartbeatServiceClient struct {
	cc *grpc.ClientConn
}

func NewHeartbeatServiceClient(cc *grpc.ClientConn) HeartbeatServiceClient {
	return &heartbeatServiceClient{cc}
}

func (c *heartbeatServiceClient) Heartbeat(ctx context.Context, in *HeartbeatRequest, opts ...grpc.CallOption) (*HeartbeatResponse, error) {
	out := new(HeartbeatResponse)
	err := grpc.Invoke(ctx, "/proto.HeartbeatService/Heartbeat", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for HeartbeatService service

type HeartbeatServiceServer interface {
	Heartbeat(context.Context, *HeartbeatRequest) (*HeartbeatResponse, error)
}

func RegisterHeartbeatServiceServer(s *grpc.Server, srv HeartbeatServiceServer) {
	s.RegisterService(&_HeartbeatService_serviceDesc, srv)
}

func _HeartbeatService_Heartbeat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HeartbeatRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HeartbeatServiceServer).Heartbeat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.HeartbeatService/Heartbeat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HeartbeatServiceServer).Heartbeat(ctx, req.(*HeartbeatRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _HeartbeatService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.HeartbeatService",
	HandlerType: (*HeartbeatServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Heartbeat",
			Handler:    _HeartbeatService_Heartbeat_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pb/heartbeat.proto",
}

func init() { proto1.RegisterFile("pb/heartbeat.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 151 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2a, 0x48, 0xd2, 0xcf,
	0x48, 0x4d, 0x2c, 0x2a, 0x49, 0x4a, 0x4d, 0x2c, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62,
	0x05, 0x53, 0x4a, 0x0e, 0x5c, 0x02, 0x1e, 0x30, 0x99, 0xa0, 0xd4, 0xc2, 0xd2, 0xd4, 0xe2, 0x12,
	0x21, 0x11, 0x2e, 0xd6, 0xc4, 0x82, 0x02, 0x4f, 0x17, 0x09, 0x46, 0x05, 0x46, 0x0d, 0xce, 0x20,
	0x08, 0x47, 0x48, 0x8c, 0x8b, 0x2d, 0x2f, 0x3f, 0x25, 0xd5, 0xd3, 0x45, 0x82, 0x49, 0x81, 0x51,
	0x83, 0x35, 0x08, 0xca, 0x53, 0x72, 0xe4, 0x12, 0x44, 0x32, 0xa1, 0xb8, 0x20, 0x3f, 0xaf, 0x38,
	0x95, 0x34, 0x23, 0x8c, 0x42, 0x90, 0x1c, 0x11, 0x9c, 0x5a, 0x54, 0x96, 0x99, 0x9c, 0x2a, 0xe4,
	0xc0, 0xc5, 0x09, 0x17, 0x13, 0x12, 0x87, 0x38, 0x5a, 0x0f, 0xdd, 0xa9, 0x52, 0x12, 0x98, 0x12,
	0x10, 0x17, 0x28, 0x31, 0x24, 0xb1, 0x81, 0xa5, 0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0xb7,
	0xad, 0xbc, 0x5e, 0xfe, 0x00, 0x00, 0x00,
}
