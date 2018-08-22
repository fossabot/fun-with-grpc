// Code generated by protoc-gen-go. DO NOT EDIT.
// source: comm.proto

package proto

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

type Msg struct {
	Val                  string   `protobuf:"bytes,1,opt,name=val,proto3" json:"val,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Msg) Reset()         { *m = Msg{} }
func (m *Msg) String() string { return proto.CompactTextString(m) }
func (*Msg) ProtoMessage()    {}
func (*Msg) Descriptor() ([]byte, []int) {
	return fileDescriptor_comm_2f8aae18f31ec1df, []int{0}
}
func (m *Msg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Msg.Unmarshal(m, b)
}
func (m *Msg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Msg.Marshal(b, m, deterministic)
}
func (dst *Msg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Msg.Merge(dst, src)
}
func (m *Msg) XXX_Size() int {
	return xxx_messageInfo_Msg.Size(m)
}
func (m *Msg) XXX_DiscardUnknown() {
	xxx_messageInfo_Msg.DiscardUnknown(m)
}

var xxx_messageInfo_Msg proto.InternalMessageInfo

func (m *Msg) GetVal() string {
	if m != nil {
		return m.Val
	}
	return ""
}

func init() {
	proto.RegisterType((*Msg)(nil), "proto.Msg")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CommClient is the client API for Comm service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CommClient interface {
	Ping(ctx context.Context, opts ...grpc.CallOption) (Comm_PingClient, error)
}

type commClient struct {
	cc *grpc.ClientConn
}

func NewCommClient(cc *grpc.ClientConn) CommClient {
	return &commClient{cc}
}

func (c *commClient) Ping(ctx context.Context, opts ...grpc.CallOption) (Comm_PingClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Comm_serviceDesc.Streams[0], "/proto.Comm/Ping", opts...)
	if err != nil {
		return nil, err
	}
	x := &commPingClient{stream}
	return x, nil
}

type Comm_PingClient interface {
	Send(*Msg) error
	Recv() (*Msg, error)
	grpc.ClientStream
}

type commPingClient struct {
	grpc.ClientStream
}

func (x *commPingClient) Send(m *Msg) error {
	return x.ClientStream.SendMsg(m)
}

func (x *commPingClient) Recv() (*Msg, error) {
	m := new(Msg)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// CommServer is the server API for Comm service.
type CommServer interface {
	Ping(Comm_PingServer) error
}

func RegisterCommServer(s *grpc.Server, srv CommServer) {
	s.RegisterService(&_Comm_serviceDesc, srv)
}

func _Comm_Ping_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CommServer).Ping(&commPingServer{stream})
}

type Comm_PingServer interface {
	Send(*Msg) error
	Recv() (*Msg, error)
	grpc.ServerStream
}

type commPingServer struct {
	grpc.ServerStream
}

func (x *commPingServer) Send(m *Msg) error {
	return x.ServerStream.SendMsg(m)
}

func (x *commPingServer) Recv() (*Msg, error) {
	m := new(Msg)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Comm_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Comm",
	HandlerType: (*CommServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Ping",
			Handler:       _Comm_Ping_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "comm.proto",
}

func init() { proto.RegisterFile("comm.proto", fileDescriptor_comm_2f8aae18f31ec1df) }

var fileDescriptor_comm_2f8aae18f31ec1df = []byte{
	// 100 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4a, 0xce, 0xcf, 0xcd,
	0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0x53, 0x4a, 0xe2, 0x5c, 0xcc, 0xbe, 0xc5,
	0xe9, 0x42, 0x02, 0x5c, 0xcc, 0x65, 0x89, 0x39, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x20,
	0xa6, 0x91, 0x0e, 0x17, 0x8b, 0x73, 0x7e, 0x6e, 0xae, 0x90, 0x0a, 0x17, 0x4b, 0x40, 0x66, 0x5e,
	0xba, 0x10, 0x17, 0x44, 0x9f, 0x9e, 0x6f, 0x71, 0xba, 0x14, 0x12, 0x5b, 0x89, 0x41, 0x83, 0xd1,
	0x80, 0x31, 0x89, 0x0d, 0x2c, 0x60, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0xdd, 0xc7, 0x4a, 0x03,
	0x62, 0x00, 0x00, 0x00,
}