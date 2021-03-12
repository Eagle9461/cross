// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cross/core/auth/msgs.proto

package types

import (
	context "context"
	fmt "fmt"
	types1 "github.com/cosmos/cosmos-sdk/codec/types"
	types "github.com/cosmos/cosmos-sdk/x/ibc/core/02-client/types"
	github_com_datachainlab_cross_x_core_account_types "github.com/datachainlab/cross/x/core/account/types"
	github_com_datachainlab_cross_x_core_tx_types "github.com/datachainlab/cross/x/core/tx/types"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type MsgSignTx struct {
	TxID    github_com_datachainlab_cross_x_core_tx_types.TxID             `protobuf:"bytes,1,opt,name=txID,proto3,casttype=github.com/datachainlab/cross/x/core/tx/types.TxID" json:"txID,omitempty"`
	Signers []github_com_datachainlab_cross_x_core_account_types.AccountID `protobuf:"bytes,2,rep,name=signers,proto3,casttype=github.com/datachainlab/cross/x/core/account/types.AccountID" json:"signers,omitempty"`
	// Timeout height relative to the current block height.
	// The timeout is disabled when set to 0.
	TimeoutHeight types.Height `protobuf:"bytes,3,opt,name=timeout_height,json=timeoutHeight,proto3" json:"timeout_height" yaml:"timeout_height"`
	// Timeout timestamp (in nanoseconds) relative to the current block timestamp.
	// The timeout is disabled when set to 0.
	TimeoutTimestamp uint64 `protobuf:"varint,4,opt,name=timeout_timestamp,json=timeoutTimestamp,proto3" json:"timeout_timestamp,omitempty" yaml:"timeout_timestamp"`
}

func (m *MsgSignTx) Reset()         { *m = MsgSignTx{} }
func (m *MsgSignTx) String() string { return proto.CompactTextString(m) }
func (*MsgSignTx) ProtoMessage()    {}
func (*MsgSignTx) Descriptor() ([]byte, []int) {
	return fileDescriptor_bca20369ddda3126, []int{0}
}
func (m *MsgSignTx) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSignTx) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSignTx.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSignTx) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSignTx.Merge(m, src)
}
func (m *MsgSignTx) XXX_Size() int {
	return m.Size()
}
func (m *MsgSignTx) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSignTx.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSignTx proto.InternalMessageInfo

// MsgSignTxResponse defines the Msg/SignTx response type.
type MsgSignTxResponse struct {
	TxAuthCompleted bool `protobuf:"varint,1,opt,name=tx_auth_completed,json=txAuthCompleted,proto3" json:"tx_auth_completed,omitempty"`
}

func (m *MsgSignTxResponse) Reset()         { *m = MsgSignTxResponse{} }
func (m *MsgSignTxResponse) String() string { return proto.CompactTextString(m) }
func (*MsgSignTxResponse) ProtoMessage()    {}
func (*MsgSignTxResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_bca20369ddda3126, []int{1}
}
func (m *MsgSignTxResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSignTxResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSignTxResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSignTxResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSignTxResponse.Merge(m, src)
}
func (m *MsgSignTxResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgSignTxResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSignTxResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSignTxResponse proto.InternalMessageInfo

type MsgIBCSignTx struct {
	CrossChainChannel *types1.Any                                                    `protobuf:"bytes,1,opt,name=cross_chain_channel,json=crossChainChannel,proto3" json:"cross_chain_channel,omitempty"`
	TxID              github_com_datachainlab_cross_x_core_tx_types.TxID             `protobuf:"bytes,2,opt,name=txID,proto3,casttype=github.com/datachainlab/cross/x/core/tx/types.TxID" json:"txID,omitempty"`
	Signers           []github_com_datachainlab_cross_x_core_account_types.AccountID `protobuf:"bytes,3,rep,name=signers,proto3,casttype=github.com/datachainlab/cross/x/core/account/types.AccountID" json:"signers,omitempty"`
	// Timeout height relative to the current block height.
	// The timeout is disabled when set to 0.
	TimeoutHeight types.Height `protobuf:"bytes,4,opt,name=timeout_height,json=timeoutHeight,proto3" json:"timeout_height" yaml:"timeout_height"`
	// Timeout timestamp (in nanoseconds) relative to the current block timestamp.
	// The timeout is disabled when set to 0.
	TimeoutTimestamp uint64 `protobuf:"varint,5,opt,name=timeout_timestamp,json=timeoutTimestamp,proto3" json:"timeout_timestamp,omitempty" yaml:"timeout_timestamp"`
}

func (m *MsgIBCSignTx) Reset()         { *m = MsgIBCSignTx{} }
func (m *MsgIBCSignTx) String() string { return proto.CompactTextString(m) }
func (*MsgIBCSignTx) ProtoMessage()    {}
func (*MsgIBCSignTx) Descriptor() ([]byte, []int) {
	return fileDescriptor_bca20369ddda3126, []int{2}
}
func (m *MsgIBCSignTx) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgIBCSignTx) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgIBCSignTx.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgIBCSignTx) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgIBCSignTx.Merge(m, src)
}
func (m *MsgIBCSignTx) XXX_Size() int {
	return m.Size()
}
func (m *MsgIBCSignTx) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgIBCSignTx.DiscardUnknown(m)
}

var xxx_messageInfo_MsgIBCSignTx proto.InternalMessageInfo

// MsgIBCSignTxResponse defines the Msg/IBCSignTx response type.
type MsgIBCSignTxResponse struct {
}

func (m *MsgIBCSignTxResponse) Reset()         { *m = MsgIBCSignTxResponse{} }
func (m *MsgIBCSignTxResponse) String() string { return proto.CompactTextString(m) }
func (*MsgIBCSignTxResponse) ProtoMessage()    {}
func (*MsgIBCSignTxResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_bca20369ddda3126, []int{3}
}
func (m *MsgIBCSignTxResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgIBCSignTxResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgIBCSignTxResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgIBCSignTxResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgIBCSignTxResponse.Merge(m, src)
}
func (m *MsgIBCSignTxResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgIBCSignTxResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgIBCSignTxResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgIBCSignTxResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgSignTx)(nil), "cross.core.auth.MsgSignTx")
	proto.RegisterType((*MsgSignTxResponse)(nil), "cross.core.auth.MsgSignTxResponse")
	proto.RegisterType((*MsgIBCSignTx)(nil), "cross.core.auth.MsgIBCSignTx")
	proto.RegisterType((*MsgIBCSignTxResponse)(nil), "cross.core.auth.MsgIBCSignTxResponse")
}

func init() { proto.RegisterFile("cross/core/auth/msgs.proto", fileDescriptor_bca20369ddda3126) }

var fileDescriptor_bca20369ddda3126 = []byte{
	// 559 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x54, 0x3d, 0x6f, 0xd3, 0x40,
	0x18, 0xb6, 0x9b, 0x50, 0xda, 0x6b, 0xa1, 0xc4, 0x14, 0x29, 0x58, 0xd4, 0x8e, 0x2c, 0x21, 0x45,
	0x0c, 0x77, 0x6a, 0x90, 0x18, 0x2a, 0x06, 0xe2, 0x64, 0xa8, 0x91, 0x32, 0x60, 0x32, 0x75, 0x09,
	0xb6, 0x7b, 0x9c, 0x2d, 0xd9, 0xbe, 0x28, 0x77, 0xae, 0x9c, 0x7f, 0xc0, 0xc8, 0x4f, 0xa8, 0xf8,
	0x17, 0xfc, 0x83, 0x8c, 0x1d, 0x99, 0x2c, 0x48, 0x16, 0xc4, 0xd8, 0xb1, 0x13, 0xf2, 0xf9, 0xa3,
	0x2d, 0x02, 0x54, 0x89, 0x8f, 0xc5, 0x7e, 0xfd, 0x3e, 0xcf, 0xfb, 0xc8, 0xf7, 0x3e, 0x8f, 0x0e,
	0xa8, 0xde, 0x8c, 0x32, 0x86, 0x3c, 0x3a, 0xc3, 0xc8, 0x49, 0xb8, 0x8f, 0x22, 0x46, 0x18, 0x9c,
	0xce, 0x28, 0xa7, 0xca, 0x8e, 0xc0, 0x60, 0x8e, 0xc1, 0x1c, 0x53, 0x1f, 0x12, 0x4a, 0x49, 0x88,
	0x91, 0x80, 0xdd, 0xe4, 0x2d, 0x72, 0xe2, 0x79, 0xc1, 0x55, 0x77, 0x09, 0x25, 0x54, 0x94, 0x28,
	0xaf, 0xca, 0xae, 0x1e, 0xb8, 0x5e, 0xa1, 0xed, 0x85, 0x01, 0x8e, 0x39, 0x3a, 0xd9, 0x2f, 0xab,
	0x82, 0x60, 0x7c, 0x5b, 0x03, 0x9b, 0x23, 0x46, 0x5e, 0x07, 0x24, 0x1e, 0xa7, 0xca, 0x4b, 0xd0,
	0xe4, 0xa9, 0x35, 0x6c, 0xcb, 0x1d, 0xb9, 0xbb, 0x6d, 0x3e, 0xbb, 0xc8, 0xf4, 0x1e, 0x09, 0xb8,
	0x9f, 0xb8, 0xd0, 0xa3, 0x11, 0x3a, 0x76, 0xb8, 0xe3, 0xf9, 0x4e, 0x10, 0x87, 0x8e, 0x8b, 0x8a,
	0xdf, 0x4e, 0x0b, 0x71, 0x9e, 0x22, 0x3e, 0x9f, 0x62, 0x06, 0xc7, 0xa9, 0x35, 0xb4, 0x85, 0x86,
	0x72, 0x04, 0x6e, 0xb3, 0x80, 0xc4, 0x78, 0xc6, 0xda, 0x6b, 0x9d, 0x46, 0x77, 0xdb, 0x7c, 0x71,
	0x91, 0xe9, 0xcf, 0x6f, 0x24, 0xe7, 0x78, 0x1e, 0x4d, 0x62, 0x5e, 0x6a, 0xf6, 0x8b, 0x2f, 0x6b,
	0x68, 0x57, 0x82, 0xca, 0x1b, 0x70, 0x97, 0x07, 0x11, 0xa6, 0x09, 0x9f, 0xf8, 0x38, 0x20, 0x3e,
	0x6f, 0x37, 0x3a, 0x72, 0x77, 0xab, 0xa7, 0xc2, 0xc0, 0xf5, 0x8a, 0x7d, 0x95, 0xa7, 0x3c, 0xd9,
	0x87, 0x87, 0x82, 0x61, 0xee, 0x2d, 0x32, 0x5d, 0x3a, 0xcf, 0xf4, 0x07, 0x73, 0x27, 0x0a, 0x0f,
	0x8c, 0xeb, 0xf3, 0x86, 0x7d, 0xa7, 0x6c, 0x14, 0x6c, 0xc5, 0x02, 0xad, 0x8a, 0x91, 0xbf, 0x19,
	0x77, 0xa2, 0x69, 0xbb, 0xd9, 0x91, 0xbb, 0x4d, 0xf3, 0xd1, 0x79, 0xa6, 0xb7, 0xaf, 0x8b, 0xd4,
	0x14, 0xc3, 0xbe, 0x57, 0xf6, 0xc6, 0x55, 0xeb, 0x60, 0xe3, 0xdd, 0xa9, 0x2e, 0x7d, 0x3d, 0xd5,
	0x25, 0xc3, 0x02, 0xad, 0x7a, 0xd7, 0x36, 0x66, 0x53, 0x1a, 0x33, 0xac, 0x3c, 0x01, 0x2d, 0x9e,
	0x4e, 0x72, 0x7b, 0x27, 0x1e, 0x8d, 0xa6, 0x21, 0xe6, 0xf8, 0x58, 0x18, 0xb0, 0x61, 0xef, 0xf0,
	0xb4, 0x9f, 0x70, 0x7f, 0x50, 0xb5, 0xaf, 0x48, 0x7d, 0x6c, 0x80, 0xed, 0x11, 0x23, 0x96, 0x39,
	0xa8, 0xad, 0xbb, 0x2f, 0x76, 0x38, 0x11, 0x2b, 0xcd, 0x9f, 0x71, 0x8c, 0x43, 0x21, 0xb4, 0xd5,
	0xdb, 0x85, 0x45, 0x70, 0x60, 0x15, 0x1c, 0xd8, 0x8f, 0xe7, 0x66, 0x73, 0x91, 0xe9, 0xb2, 0xdd,
	0x12, 0x63, 0x83, 0x7c, 0x6a, 0x50, 0x0c, 0xd5, 0x31, 0x58, 0xfb, 0xbb, 0x31, 0x68, 0xfc, 0xfb,
	0x18, 0x34, 0xff, 0x47, 0x0c, 0x6e, 0xfd, 0x61, 0x0c, 0x3a, 0x60, 0xf7, 0xaa, 0x75, 0x55, 0x12,
	0x2e, 0x19, 0xbd, 0x0f, 0x32, 0x68, 0x8c, 0x18, 0x51, 0x0e, 0xc1, 0x7a, 0x69, 0xaf, 0x0a, 0x7f,
	0xb8, 0x0b, 0x60, 0x9d, 0x24, 0xd5, 0xf8, 0x35, 0x56, 0xa7, 0xec, 0x15, 0xd8, 0xbc, 0xcc, 0xca,
	0xde, 0xcf, 0x06, 0x6a, 0x58, 0x7d, 0xfc, 0x5b, 0xb8, 0x92, 0x34, 0x47, 0x8b, 0x2f, 0x9a, 0xb4,
	0x58, 0x6a, 0xf2, 0xd9, 0x52, 0x93, 0x3f, 0x2f, 0x35, 0xf9, 0xfd, 0x4a, 0x93, 0xce, 0x56, 0x9a,
	0xf4, 0x69, 0xa5, 0x49, 0x47, 0xe8, 0x66, 0x16, 0xe7, 0x37, 0x9e, 0xf0, 0xd7, 0x5d, 0x17, 0xd9,
	0x7c, 0xfa, 0x3d, 0x00, 0x00, 0xff, 0xff, 0x8b, 0x96, 0x75, 0xe7, 0x11, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	// SignTx defines a rpc handler method for MsgSignTx.
	SignTx(ctx context.Context, in *MsgSignTx, opts ...grpc.CallOption) (*MsgSignTxResponse, error)
	// IBCSignTx defines a rpc handler method for MsgIBCSignTx.
	IBCSignTx(ctx context.Context, in *MsgIBCSignTx, opts ...grpc.CallOption) (*MsgIBCSignTxResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) SignTx(ctx context.Context, in *MsgSignTx, opts ...grpc.CallOption) (*MsgSignTxResponse, error) {
	out := new(MsgSignTxResponse)
	err := c.cc.Invoke(ctx, "/cross.core.auth.Msg/SignTx", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) IBCSignTx(ctx context.Context, in *MsgIBCSignTx, opts ...grpc.CallOption) (*MsgIBCSignTxResponse, error) {
	out := new(MsgIBCSignTxResponse)
	err := c.cc.Invoke(ctx, "/cross.core.auth.Msg/IBCSignTx", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	// SignTx defines a rpc handler method for MsgSignTx.
	SignTx(context.Context, *MsgSignTx) (*MsgSignTxResponse, error)
	// IBCSignTx defines a rpc handler method for MsgIBCSignTx.
	IBCSignTx(context.Context, *MsgIBCSignTx) (*MsgIBCSignTxResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) SignTx(ctx context.Context, req *MsgSignTx) (*MsgSignTxResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SignTx not implemented")
}
func (*UnimplementedMsgServer) IBCSignTx(ctx context.Context, req *MsgIBCSignTx) (*MsgIBCSignTxResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IBCSignTx not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_SignTx_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgSignTx)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).SignTx(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cross.core.auth.Msg/SignTx",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).SignTx(ctx, req.(*MsgSignTx))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_IBCSignTx_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgIBCSignTx)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).IBCSignTx(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cross.core.auth.Msg/IBCSignTx",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).IBCSignTx(ctx, req.(*MsgIBCSignTx))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cross.core.auth.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SignTx",
			Handler:    _Msg_SignTx_Handler,
		},
		{
			MethodName: "IBCSignTx",
			Handler:    _Msg_IBCSignTx_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cross/core/auth/msgs.proto",
}

func (m *MsgSignTx) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSignTx) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSignTx) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.TimeoutTimestamp != 0 {
		i = encodeVarintMsgs(dAtA, i, uint64(m.TimeoutTimestamp))
		i--
		dAtA[i] = 0x20
	}
	{
		size, err := m.TimeoutHeight.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintMsgs(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if len(m.Signers) > 0 {
		for iNdEx := len(m.Signers) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Signers[iNdEx])
			copy(dAtA[i:], m.Signers[iNdEx])
			i = encodeVarintMsgs(dAtA, i, uint64(len(m.Signers[iNdEx])))
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.TxID) > 0 {
		i -= len(m.TxID)
		copy(dAtA[i:], m.TxID)
		i = encodeVarintMsgs(dAtA, i, uint64(len(m.TxID)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgSignTxResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSignTxResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSignTxResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.TxAuthCompleted {
		i--
		if m.TxAuthCompleted {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *MsgIBCSignTx) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgIBCSignTx) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgIBCSignTx) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.TimeoutTimestamp != 0 {
		i = encodeVarintMsgs(dAtA, i, uint64(m.TimeoutTimestamp))
		i--
		dAtA[i] = 0x28
	}
	{
		size, err := m.TimeoutHeight.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintMsgs(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	if len(m.Signers) > 0 {
		for iNdEx := len(m.Signers) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Signers[iNdEx])
			copy(dAtA[i:], m.Signers[iNdEx])
			i = encodeVarintMsgs(dAtA, i, uint64(len(m.Signers[iNdEx])))
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.TxID) > 0 {
		i -= len(m.TxID)
		copy(dAtA[i:], m.TxID)
		i = encodeVarintMsgs(dAtA, i, uint64(len(m.TxID)))
		i--
		dAtA[i] = 0x12
	}
	if m.CrossChainChannel != nil {
		{
			size, err := m.CrossChainChannel.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintMsgs(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgIBCSignTxResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgIBCSignTxResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgIBCSignTxResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintMsgs(dAtA []byte, offset int, v uint64) int {
	offset -= sovMsgs(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgSignTx) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.TxID)
	if l > 0 {
		n += 1 + l + sovMsgs(uint64(l))
	}
	if len(m.Signers) > 0 {
		for _, b := range m.Signers {
			l = len(b)
			n += 1 + l + sovMsgs(uint64(l))
		}
	}
	l = m.TimeoutHeight.Size()
	n += 1 + l + sovMsgs(uint64(l))
	if m.TimeoutTimestamp != 0 {
		n += 1 + sovMsgs(uint64(m.TimeoutTimestamp))
	}
	return n
}

func (m *MsgSignTxResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.TxAuthCompleted {
		n += 2
	}
	return n
}

func (m *MsgIBCSignTx) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.CrossChainChannel != nil {
		l = m.CrossChainChannel.Size()
		n += 1 + l + sovMsgs(uint64(l))
	}
	l = len(m.TxID)
	if l > 0 {
		n += 1 + l + sovMsgs(uint64(l))
	}
	if len(m.Signers) > 0 {
		for _, b := range m.Signers {
			l = len(b)
			n += 1 + l + sovMsgs(uint64(l))
		}
	}
	l = m.TimeoutHeight.Size()
	n += 1 + l + sovMsgs(uint64(l))
	if m.TimeoutTimestamp != 0 {
		n += 1 + sovMsgs(uint64(m.TimeoutTimestamp))
	}
	return n
}

func (m *MsgIBCSignTxResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovMsgs(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMsgs(x uint64) (n int) {
	return sovMsgs(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgSignTx) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMsgs
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgSignTx: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSignTx: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TxID", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgs
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthMsgs
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthMsgs
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TxID = append(m.TxID[:0], dAtA[iNdEx:postIndex]...)
			if m.TxID == nil {
				m.TxID = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Signers", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgs
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthMsgs
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthMsgs
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Signers = append(m.Signers, make([]byte, postIndex-iNdEx))
			copy(m.Signers[len(m.Signers)-1], dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TimeoutHeight", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgs
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthMsgs
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMsgs
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.TimeoutHeight.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TimeoutTimestamp", wireType)
			}
			m.TimeoutTimestamp = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgs
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TimeoutTimestamp |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipMsgs(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMsgs
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthMsgs
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgSignTxResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMsgs
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgSignTxResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSignTxResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TxAuthCompleted", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgs
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.TxAuthCompleted = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipMsgs(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMsgs
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthMsgs
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgIBCSignTx) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMsgs
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgIBCSignTx: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgIBCSignTx: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CrossChainChannel", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgs
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthMsgs
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMsgs
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.CrossChainChannel == nil {
				m.CrossChainChannel = &types1.Any{}
			}
			if err := m.CrossChainChannel.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TxID", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgs
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthMsgs
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthMsgs
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TxID = append(m.TxID[:0], dAtA[iNdEx:postIndex]...)
			if m.TxID == nil {
				m.TxID = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Signers", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgs
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthMsgs
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthMsgs
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Signers = append(m.Signers, make([]byte, postIndex-iNdEx))
			copy(m.Signers[len(m.Signers)-1], dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TimeoutHeight", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgs
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthMsgs
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMsgs
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.TimeoutHeight.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TimeoutTimestamp", wireType)
			}
			m.TimeoutTimestamp = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgs
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TimeoutTimestamp |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipMsgs(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMsgs
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthMsgs
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgIBCSignTxResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMsgs
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgIBCSignTxResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgIBCSignTxResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipMsgs(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMsgs
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthMsgs
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipMsgs(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMsgs
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowMsgs
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowMsgs
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthMsgs
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupMsgs
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthMsgs
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthMsgs        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMsgs          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupMsgs = fmt.Errorf("proto: unexpected end of group")
)