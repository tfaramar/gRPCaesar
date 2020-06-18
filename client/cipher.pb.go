// Code generated by protoc-gen-go. DO NOT EDIT.
// source: cipher.proto

package main

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type Input struct {
	Text                 string   `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
	Shift                int32    `protobuf:"varint,2,opt,name=shift,proto3" json:"shift,omitempty"`
	Encode               bool     `protobuf:"varint,3,opt,name=encode,proto3" json:"encode,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Input) Reset()         { *m = Input{} }
func (m *Input) String() string { return proto.CompactTextString(m) }
func (*Input) ProtoMessage()    {}
func (*Input) Descriptor() ([]byte, []int) {
	return fileDescriptor_11bb0febb557427c, []int{0}
}

func (m *Input) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Input.Unmarshal(m, b)
}
func (m *Input) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Input.Marshal(b, m, deterministic)
}
func (m *Input) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Input.Merge(m, src)
}
func (m *Input) XXX_Size() int {
	return xxx_messageInfo_Input.Size(m)
}
func (m *Input) XXX_DiscardUnknown() {
	xxx_messageInfo_Input.DiscardUnknown(m)
}

var xxx_messageInfo_Input proto.InternalMessageInfo

func (m *Input) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

func (m *Input) GetShift() int32 {
	if m != nil {
		return m.Shift
	}
	return 0
}

func (m *Input) GetEncode() bool {
	if m != nil {
		return m.Encode
	}
	return false
}

type Message struct {
	Msg                  string   `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_11bb0febb557427c, []int{1}
}

func (m *Message) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Message.Unmarshal(m, b)
}
func (m *Message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Message.Marshal(b, m, deterministic)
}
func (m *Message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message.Merge(m, src)
}
func (m *Message) XXX_Size() int {
	return xxx_messageInfo_Message.Size(m)
}
func (m *Message) XXX_DiscardUnknown() {
	xxx_messageInfo_Message.DiscardUnknown(m)
}

var xxx_messageInfo_Message proto.InternalMessageInfo

func (m *Message) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func init() {
	proto.RegisterType((*Input)(nil), "cipher.Input")
	proto.RegisterType((*Message)(nil), "cipher.Message")
}

func init() {
	proto.RegisterFile("cipher.proto", fileDescriptor_11bb0febb557427c)
}

var fileDescriptor_11bb0febb557427c = []byte{
	// 172 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x49, 0xce, 0x2c, 0xc8,
	0x48, 0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x83, 0xf0, 0x94, 0x3c, 0xb9, 0x58,
	0x3d, 0xf3, 0x0a, 0x4a, 0x4b, 0x84, 0x84, 0xb8, 0x58, 0x4a, 0x52, 0x2b, 0x4a, 0x24, 0x18, 0x15,
	0x18, 0x35, 0x38, 0x83, 0xc0, 0x6c, 0x21, 0x11, 0x2e, 0xd6, 0xe2, 0x8c, 0xcc, 0xb4, 0x12, 0x09,
	0x26, 0x05, 0x46, 0x0d, 0xd6, 0x20, 0x08, 0x47, 0x48, 0x8c, 0x8b, 0x2d, 0x35, 0x2f, 0x39, 0x3f,
	0x25, 0x55, 0x82, 0x59, 0x81, 0x51, 0x83, 0x23, 0x08, 0xca, 0x53, 0x92, 0xe6, 0x62, 0xf7, 0x4d,
	0x2d, 0x2e, 0x4e, 0x4c, 0x4f, 0x15, 0x12, 0xe0, 0x62, 0xce, 0x2d, 0x4e, 0x87, 0x9a, 0x05, 0x62,
	0x1a, 0x39, 0x73, 0xf1, 0x3a, 0x83, 0x6d, 0x0c, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0x15, 0x32,
	0xe2, 0xe2, 0x0b, 0x28, 0xca, 0x4f, 0x4e, 0x2d, 0x2e, 0x86, 0x69, 0xe2, 0xd5, 0x83, 0xba, 0x10,
	0xec, 0x20, 0x29, 0x7e, 0x18, 0x17, 0x2a, 0xaf, 0xc4, 0x90, 0xc4, 0x06, 0x76, 0xbb, 0x31, 0x20,
	0x00, 0x00, 0xff, 0xff, 0xf5, 0xde, 0x1c, 0x89, 0xcb, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CipherServiceClient is the client API for CipherService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CipherServiceClient interface {
	ProcessMessage(ctx context.Context, in *Input, opts ...grpc.CallOption) (*Message, error)
}

type cipherServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCipherServiceClient(cc grpc.ClientConnInterface) CipherServiceClient {
	return &cipherServiceClient{cc}
}

func (c *cipherServiceClient) ProcessMessage(ctx context.Context, in *Input, opts ...grpc.CallOption) (*Message, error) {
	out := new(Message)
	err := c.cc.Invoke(ctx, "/cipher.CipherService/ProcessMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CipherServiceServer is the server API for CipherService service.
type CipherServiceServer interface {
	ProcessMessage(context.Context, *Input) (*Message, error)
}

// UnimplementedCipherServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCipherServiceServer struct {
}

func (*UnimplementedCipherServiceServer) ProcessMessage(ctx context.Context, req *Input) (*Message, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProcessMessage not implemented")
}

func RegisterCipherServiceServer(s *grpc.Server, srv CipherServiceServer) {
	s.RegisterService(&_CipherService_serviceDesc, srv)
}

func _CipherService_ProcessMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Input)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CipherServiceServer).ProcessMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cipher.CipherService/ProcessMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CipherServiceServer).ProcessMessage(ctx, req.(*Input))
	}
	return interceptor(ctx, in, info, handler)
}

var _CipherService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cipher.CipherService",
	HandlerType: (*CipherServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ProcessMessage",
			Handler:    _CipherService_ProcessMessage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cipher.proto",
}
