// Code generated by protoc-gen-go. DO NOT EDIT.
// source: cpe_demo.proto

package cpedemo

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

type CpeInfo struct {
	LocalIp              string   `protobuf:"bytes,1,opt,name=localIp,proto3" json:"localIp,omitempty"`
	ExternalIp           string   `protobuf:"bytes,2,opt,name=externalIp,proto3" json:"externalIp,omitempty"`
	CpeName              string   `protobuf:"bytes,3,opt,name=cpeName,proto3" json:"cpeName,omitempty"`
	CpeUuid              string   `protobuf:"bytes,4,opt,name=cpeUuid,proto3" json:"cpeUuid,omitempty"`
	Other                string   `protobuf:"bytes,5,opt,name=other,proto3" json:"other,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CpeInfo) Reset()         { *m = CpeInfo{} }
func (m *CpeInfo) String() string { return proto.CompactTextString(m) }
func (*CpeInfo) ProtoMessage()    {}
func (*CpeInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_cpe_demo_52f7d02de55c3c45, []int{0}
}
func (m *CpeInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CpeInfo.Unmarshal(m, b)
}
func (m *CpeInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CpeInfo.Marshal(b, m, deterministic)
}
func (dst *CpeInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CpeInfo.Merge(dst, src)
}
func (m *CpeInfo) XXX_Size() int {
	return xxx_messageInfo_CpeInfo.Size(m)
}
func (m *CpeInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_CpeInfo.DiscardUnknown(m)
}

var xxx_messageInfo_CpeInfo proto.InternalMessageInfo

func (m *CpeInfo) GetLocalIp() string {
	if m != nil {
		return m.LocalIp
	}
	return ""
}

func (m *CpeInfo) GetExternalIp() string {
	if m != nil {
		return m.ExternalIp
	}
	return ""
}

func (m *CpeInfo) GetCpeName() string {
	if m != nil {
		return m.CpeName
	}
	return ""
}

func (m *CpeInfo) GetCpeUuid() string {
	if m != nil {
		return m.CpeUuid
	}
	return ""
}

func (m *CpeInfo) GetOther() string {
	if m != nil {
		return m.Other
	}
	return ""
}

type MsgStruc struct {
	MsgId                int32    `protobuf:"varint,1,opt,name=msgId,proto3" json:"msgId,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MsgStruc) Reset()         { *m = MsgStruc{} }
func (m *MsgStruc) String() string { return proto.CompactTextString(m) }
func (*MsgStruc) ProtoMessage()    {}
func (*MsgStruc) Descriptor() ([]byte, []int) {
	return fileDescriptor_cpe_demo_52f7d02de55c3c45, []int{1}
}
func (m *MsgStruc) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MsgStruc.Unmarshal(m, b)
}
func (m *MsgStruc) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MsgStruc.Marshal(b, m, deterministic)
}
func (dst *MsgStruc) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgStruc.Merge(dst, src)
}
func (m *MsgStruc) XXX_Size() int {
	return xxx_messageInfo_MsgStruc.Size(m)
}
func (m *MsgStruc) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgStruc.DiscardUnknown(m)
}

var xxx_messageInfo_MsgStruc proto.InternalMessageInfo

func (m *MsgStruc) GetMsgId() int32 {
	if m != nil {
		return m.MsgId
	}
	return 0
}

func (m *MsgStruc) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type Cpe2Cm struct {
	CpeUuid              string    `protobuf:"bytes,1,opt,name=cpeUuid,proto3" json:"cpeUuid,omitempty"`
	ResInfo              *MsgStruc `protobuf:"bytes,2,opt,name=resInfo,proto3" json:"resInfo,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Cpe2Cm) Reset()         { *m = Cpe2Cm{} }
func (m *Cpe2Cm) String() string { return proto.CompactTextString(m) }
func (*Cpe2Cm) ProtoMessage()    {}
func (*Cpe2Cm) Descriptor() ([]byte, []int) {
	return fileDescriptor_cpe_demo_52f7d02de55c3c45, []int{2}
}
func (m *Cpe2Cm) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Cpe2Cm.Unmarshal(m, b)
}
func (m *Cpe2Cm) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Cpe2Cm.Marshal(b, m, deterministic)
}
func (dst *Cpe2Cm) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Cpe2Cm.Merge(dst, src)
}
func (m *Cpe2Cm) XXX_Size() int {
	return xxx_messageInfo_Cpe2Cm.Size(m)
}
func (m *Cpe2Cm) XXX_DiscardUnknown() {
	xxx_messageInfo_Cpe2Cm.DiscardUnknown(m)
}

var xxx_messageInfo_Cpe2Cm proto.InternalMessageInfo

func (m *Cpe2Cm) GetCpeUuid() string {
	if m != nil {
		return m.CpeUuid
	}
	return ""
}

func (m *Cpe2Cm) GetResInfo() *MsgStruc {
	if m != nil {
		return m.ResInfo
	}
	return nil
}

type Cm2Cpe struct {
	CpeUuid              string    `protobuf:"bytes,1,opt,name=cpeUuid,proto3" json:"cpeUuid,omitempty"`
	ConfigInfo           *MsgStruc `protobuf:"bytes,2,opt,name=configInfo,proto3" json:"configInfo,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Cm2Cpe) Reset()         { *m = Cm2Cpe{} }
func (m *Cm2Cpe) String() string { return proto.CompactTextString(m) }
func (*Cm2Cpe) ProtoMessage()    {}
func (*Cm2Cpe) Descriptor() ([]byte, []int) {
	return fileDescriptor_cpe_demo_52f7d02de55c3c45, []int{3}
}
func (m *Cm2Cpe) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Cm2Cpe.Unmarshal(m, b)
}
func (m *Cm2Cpe) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Cm2Cpe.Marshal(b, m, deterministic)
}
func (dst *Cm2Cpe) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Cm2Cpe.Merge(dst, src)
}
func (m *Cm2Cpe) XXX_Size() int {
	return xxx_messageInfo_Cm2Cpe.Size(m)
}
func (m *Cm2Cpe) XXX_DiscardUnknown() {
	xxx_messageInfo_Cm2Cpe.DiscardUnknown(m)
}

var xxx_messageInfo_Cm2Cpe proto.InternalMessageInfo

func (m *Cm2Cpe) GetCpeUuid() string {
	if m != nil {
		return m.CpeUuid
	}
	return ""
}

func (m *Cm2Cpe) GetConfigInfo() *MsgStruc {
	if m != nil {
		return m.ConfigInfo
	}
	return nil
}

func init() {
	proto.RegisterType((*CpeInfo)(nil), "cpedemo.CpeInfo")
	proto.RegisterType((*MsgStruc)(nil), "cpedemo.MsgStruc")
	proto.RegisterType((*Cpe2Cm)(nil), "cpedemo.Cpe2Cm")
	proto.RegisterType((*Cm2Cpe)(nil), "cpedemo.Cm2Cpe")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CpeDemoClient is the client API for CpeDemo service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CpeDemoClient interface {
	Register(ctx context.Context, in *CpeInfo, opts ...grpc.CallOption) (*MsgStruc, error)
	Subscribe(ctx context.Context, in *CpeInfo, opts ...grpc.CallOption) (CpeDemo_SubscribeClient, error)
	Report(ctx context.Context, in *Cpe2Cm, opts ...grpc.CallOption) (*MsgStruc, error)
}

type cpeDemoClient struct {
	cc *grpc.ClientConn
}

func NewCpeDemoClient(cc *grpc.ClientConn) CpeDemoClient {
	return &cpeDemoClient{cc}
}

func (c *cpeDemoClient) Register(ctx context.Context, in *CpeInfo, opts ...grpc.CallOption) (*MsgStruc, error) {
	out := new(MsgStruc)
	err := c.cc.Invoke(ctx, "/cpedemo.CpeDemo/Register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cpeDemoClient) Subscribe(ctx context.Context, in *CpeInfo, opts ...grpc.CallOption) (CpeDemo_SubscribeClient, error) {
	stream, err := c.cc.NewStream(ctx, &_CpeDemo_serviceDesc.Streams[0], "/cpedemo.CpeDemo/Subscribe", opts...)
	if err != nil {
		return nil, err
	}
	x := &cpeDemoSubscribeClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type CpeDemo_SubscribeClient interface {
	Recv() (*MsgStruc, error)
	grpc.ClientStream
}

type cpeDemoSubscribeClient struct {
	grpc.ClientStream
}

func (x *cpeDemoSubscribeClient) Recv() (*MsgStruc, error) {
	m := new(MsgStruc)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *cpeDemoClient) Report(ctx context.Context, in *Cpe2Cm, opts ...grpc.CallOption) (*MsgStruc, error) {
	out := new(MsgStruc)
	err := c.cc.Invoke(ctx, "/cpedemo.CpeDemo/Report", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CpeDemoServer is the server API for CpeDemo service.
type CpeDemoServer interface {
	Register(context.Context, *CpeInfo) (*MsgStruc, error)
	Subscribe(*CpeInfo, CpeDemo_SubscribeServer) error
	Report(context.Context, *Cpe2Cm) (*MsgStruc, error)
}

func RegisterCpeDemoServer(s *grpc.Server, srv CpeDemoServer) {
	s.RegisterService(&_CpeDemo_serviceDesc, srv)
}

func _CpeDemo_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CpeInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CpeDemoServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cpedemo.CpeDemo/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CpeDemoServer).Register(ctx, req.(*CpeInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _CpeDemo_Subscribe_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(CpeInfo)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CpeDemoServer).Subscribe(m, &cpeDemoSubscribeServer{stream})
}

type CpeDemo_SubscribeServer interface {
	Send(*MsgStruc) error
	grpc.ServerStream
}

type cpeDemoSubscribeServer struct {
	grpc.ServerStream
}

func (x *cpeDemoSubscribeServer) Send(m *MsgStruc) error {
	return x.ServerStream.SendMsg(m)
}

func _CpeDemo_Report_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Cpe2Cm)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CpeDemoServer).Report(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cpedemo.CpeDemo/Report",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CpeDemoServer).Report(ctx, req.(*Cpe2Cm))
	}
	return interceptor(ctx, in, info, handler)
}

var _CpeDemo_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cpedemo.CpeDemo",
	HandlerType: (*CpeDemoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _CpeDemo_Register_Handler,
		},
		{
			MethodName: "Report",
			Handler:    _CpeDemo_Report_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Subscribe",
			Handler:       _CpeDemo_Subscribe_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "cpe_demo.proto",
}

func init() { proto.RegisterFile("cpe_demo.proto", fileDescriptor_cpe_demo_52f7d02de55c3c45) }

var fileDescriptor_cpe_demo_52f7d02de55c3c45 = []byte{
	// 325 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x92, 0xcd, 0x4e, 0xeb, 0x30,
	0x10, 0x85, 0x6f, 0x2e, 0xb4, 0x69, 0x07, 0xc4, 0x8f, 0xc5, 0x22, 0x62, 0x81, 0x50, 0x56, 0x08,
	0x24, 0x0b, 0x02, 0x2b, 0x96, 0x2d, 0x9b, 0x2e, 0x80, 0x2a, 0x55, 0xd7, 0x28, 0x75, 0xa7, 0x26,
	0x52, 0x5d, 0x5b, 0x76, 0x2a, 0xf5, 0x25, 0x78, 0x09, 0x9e, 0x14, 0xdb, 0x71, 0xaa, 0x20, 0x50,
	0xc5, 0xf2, 0x78, 0xce, 0x77, 0x66, 0x3c, 0x36, 0x1c, 0x31, 0x85, 0x6f, 0x73, 0x14, 0x92, 0x2a,
	0x2d, 0x2b, 0x49, 0x62, 0xab, 0x9d, 0x4c, 0x3f, 0x22, 0x88, 0x87, 0x0a, 0x47, 0xab, 0x85, 0x24,
	0x09, 0xc4, 0x4b, 0xc9, 0x8a, 0xe5, 0x48, 0x25, 0xd1, 0x65, 0x74, 0xd5, 0xcf, 0x1b, 0x49, 0x2e,
	0x00, 0x70, 0x53, 0xa1, 0x5e, 0xf9, 0xe2, 0x7f, 0x5f, 0x6c, 0x9d, 0x38, 0xd2, 0x06, 0xbe, 0x14,
	0x02, 0x93, 0xbd, 0x9a, 0x0c, 0x32, 0x54, 0xa6, 0xeb, 0x72, 0x9e, 0xec, 0x6f, 0x2b, 0x4e, 0x92,
	0x33, 0xe8, 0xc8, 0xea, 0x1d, 0x75, 0xd2, 0xf1, 0xe7, 0xb5, 0x48, 0x1f, 0xa1, 0xf7, 0x6c, 0xf8,
	0xa4, 0xd2, 0x6b, 0xe6, 0x1c, 0xc2, 0xf0, 0xd1, 0xdc, 0x4f, 0xd3, 0xc9, 0x6b, 0xe1, 0x12, 0x05,
	0x1a, 0x53, 0x70, 0x0c, 0x83, 0x34, 0x32, 0x7d, 0x85, 0xae, 0xbd, 0x4a, 0x36, 0x14, 0xed, 0xae,
	0xd1, 0xf7, 0xae, 0x37, 0x10, 0x6b, 0x34, 0xee, 0xba, 0x9e, 0x3e, 0xc8, 0x4e, 0x69, 0x58, 0x05,
	0x6d, 0xfa, 0xe6, 0x8d, 0x23, 0x9d, 0xda, 0x40, 0x91, 0xd9, 0xcc, 0x1d, 0x81, 0x77, 0x00, 0x4c,
	0xae, 0x16, 0x25, 0xdf, 0x9d, 0xd9, 0x32, 0x65, 0x9f, 0xf5, 0xce, 0x9f, 0xac, 0xc1, 0xe2, 0xbd,
	0x1c, 0x79, 0x69, 0xec, 0x2a, 0xc9, 0xc9, 0x16, 0x0b, 0x2f, 0x72, 0xfe, 0x33, 0x28, 0xfd, 0x47,
	0x1e, 0xa0, 0x3f, 0x59, 0xcf, 0x0c, 0xd3, 0xe5, 0x0c, 0xff, 0xc8, 0xdc, 0x46, 0x84, 0x42, 0x37,
	0x47, 0x25, 0x75, 0x45, 0x8e, 0xdb, 0x88, 0xdd, 0xd6, 0xaf, 0xc4, 0xe0, 0x1a, 0x92, 0x52, 0x52,
	0xae, 0x15, 0xa3, 0xb8, 0x29, 0x84, 0x5a, 0xa2, 0x69, 0x6c, 0x83, 0xc3, 0x30, 0xfd, 0xd8, 0xfd,
	0xa5, 0x71, 0x34, 0xeb, 0xfa, 0x4f, 0x75, 0xff, 0x15, 0x00, 0x00, 0xff, 0xff, 0x28, 0xf2, 0xaf,
	0x98, 0x66, 0x02, 0x00, 0x00,
}
