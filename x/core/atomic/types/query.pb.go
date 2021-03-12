// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cross/core/atomic/query.proto

package types

import (
	context "context"
	fmt "fmt"
	github_com_datachainlab_cross_x_core_tx_types "github.com/datachainlab/cross/x/core/tx/types"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type QueryCoordinatorStateRequest struct {
	TxId github_com_datachainlab_cross_x_core_tx_types.TxID `protobuf:"bytes,1,opt,name=tx_id,json=txId,proto3,casttype=github.com/datachainlab/cross/x/core/tx/types.TxID" json:"tx_id,omitempty"`
}

func (m *QueryCoordinatorStateRequest) Reset()         { *m = QueryCoordinatorStateRequest{} }
func (m *QueryCoordinatorStateRequest) String() string { return proto.CompactTextString(m) }
func (*QueryCoordinatorStateRequest) ProtoMessage()    {}
func (*QueryCoordinatorStateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_54c8df456f56ba8a, []int{0}
}
func (m *QueryCoordinatorStateRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryCoordinatorStateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryCoordinatorStateRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryCoordinatorStateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryCoordinatorStateRequest.Merge(m, src)
}
func (m *QueryCoordinatorStateRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryCoordinatorStateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryCoordinatorStateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryCoordinatorStateRequest proto.InternalMessageInfo

type QueryCoordinatorStateResponse struct {
	CoodinatorState CoordinatorState `protobuf:"bytes,1,opt,name=coodinator_state,json=coodinatorState,proto3" json:"coodinator_state"`
}

func (m *QueryCoordinatorStateResponse) Reset()         { *m = QueryCoordinatorStateResponse{} }
func (m *QueryCoordinatorStateResponse) String() string { return proto.CompactTextString(m) }
func (*QueryCoordinatorStateResponse) ProtoMessage()    {}
func (*QueryCoordinatorStateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_54c8df456f56ba8a, []int{1}
}
func (m *QueryCoordinatorStateResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryCoordinatorStateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryCoordinatorStateResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryCoordinatorStateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryCoordinatorStateResponse.Merge(m, src)
}
func (m *QueryCoordinatorStateResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryCoordinatorStateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryCoordinatorStateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryCoordinatorStateResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*QueryCoordinatorStateRequest)(nil), "cross.core.atomic.QueryCoordinatorStateRequest")
	proto.RegisterType((*QueryCoordinatorStateResponse)(nil), "cross.core.atomic.QueryCoordinatorStateResponse")
}

func init() { proto.RegisterFile("cross/core/atomic/query.proto", fileDescriptor_54c8df456f56ba8a) }

var fileDescriptor_54c8df456f56ba8a = []byte{
	// 349 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4d, 0x2e, 0xca, 0x2f,
	0x2e, 0xd6, 0x4f, 0xce, 0x2f, 0x4a, 0xd5, 0x4f, 0x2c, 0xc9, 0xcf, 0xcd, 0x4c, 0xd6, 0x2f, 0x2c,
	0x4d, 0x2d, 0xaa, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x04, 0x4b, 0xeb, 0x81, 0xa4,
	0xf5, 0x20, 0xd2, 0x52, 0x22, 0xe9, 0xf9, 0xe9, 0xf9, 0x60, 0x59, 0x7d, 0x10, 0x0b, 0xa2, 0x50,
	0x0a, 0x8b, 0x39, 0x25, 0x95, 0x05, 0xa9, 0xc5, 0x50, 0x69, 0x99, 0xf4, 0xfc, 0xfc, 0xf4, 0x9c,
	0x54, 0xfd, 0xc4, 0x82, 0x4c, 0xfd, 0xc4, 0xbc, 0xbc, 0xfc, 0x92, 0xc4, 0x92, 0xcc, 0xfc, 0x3c,
	0xa8, 0xac, 0x52, 0x36, 0x97, 0x4c, 0x20, 0xc8, 0x52, 0xe7, 0xfc, 0xfc, 0xa2, 0x94, 0xcc, 0xbc,
	0xc4, 0x92, 0xfc, 0xa2, 0xe0, 0x92, 0xc4, 0x92, 0xd4, 0xa0, 0xd4, 0xc2, 0xd2, 0xd4, 0xe2, 0x12,
	0x21, 0x6f, 0x2e, 0xd6, 0x92, 0x8a, 0xf8, 0xcc, 0x14, 0x09, 0x46, 0x05, 0x46, 0x0d, 0x1e, 0x27,
	0xb3, 0x5f, 0xf7, 0xe4, 0x8d, 0xd2, 0x33, 0x4b, 0x32, 0x4a, 0x93, 0xf4, 0x92, 0xf3, 0x73, 0xf5,
	0x53, 0x12, 0x4b, 0x12, 0x93, 0x33, 0x12, 0x33, 0xf3, 0x72, 0x12, 0x93, 0xf4, 0x21, 0xee, 0xa8,
	0x80, 0xb8, 0xa4, 0xa4, 0x02, 0xea, 0x8a, 0x90, 0x0a, 0x4f, 0x97, 0x20, 0x96, 0x92, 0x0a, 0xcf,
	0x14, 0xa5, 0x52, 0x2e, 0x59, 0x1c, 0x96, 0x15, 0x17, 0xe4, 0xe7, 0x15, 0xa7, 0x0a, 0x85, 0x70,
	0x09, 0x24, 0xe7, 0xe7, 0x43, 0xa5, 0xe2, 0x8b, 0x41, 0x72, 0x60, 0x8b, 0xb9, 0x8d, 0x94, 0xf5,
	0x30, 0x82, 0x43, 0x0f, 0xdd, 0x18, 0x27, 0x96, 0x13, 0xf7, 0xe4, 0x19, 0x82, 0xf8, 0x11, 0x46,
	0x80, 0x85, 0x8d, 0xd6, 0x32, 0x72, 0xb1, 0x82, 0xed, 0x15, 0x5a, 0xcc, 0xc8, 0x25, 0x80, 0xae,
	0x4b, 0x48, 0x1f, 0x8b, 0xd1, 0xf8, 0xc2, 0x44, 0xca, 0x80, 0x78, 0x0d, 0x10, 0x7f, 0x29, 0xe9,
	0x34, 0x5d, 0x7e, 0x32, 0x99, 0x49, 0x4d, 0x48, 0x45, 0x1f, 0x33, 0xae, 0x92, 0x11, 0x9a, 0x74,
	0xc1, 0x3e, 0x76, 0xf2, 0x3f, 0xf1, 0x50, 0x8e, 0xe1, 0xc4, 0x23, 0x39, 0xc6, 0x0b, 0x8f, 0xe4,
	0x18, 0x1f, 0x3c, 0x92, 0x63, 0x9c, 0xf0, 0x58, 0x8e, 0xe1, 0xc2, 0x63, 0x39, 0x86, 0x1b, 0x8f,
	0xe5, 0x18, 0xa2, 0x0c, 0x89, 0x0a, 0x7e, 0xe4, 0x84, 0x90, 0xc4, 0x06, 0x8e, 0x6b, 0x63, 0x40,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x20, 0x9f, 0x11, 0x3f, 0x72, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
	CoordinatorState(ctx context.Context, in *QueryCoordinatorStateRequest, opts ...grpc.CallOption) (*QueryCoordinatorStateResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) CoordinatorState(ctx context.Context, in *QueryCoordinatorStateRequest, opts ...grpc.CallOption) (*QueryCoordinatorStateResponse, error) {
	out := new(QueryCoordinatorStateResponse)
	err := c.cc.Invoke(ctx, "/cross.core.atomic.Query/CoordinatorState", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	CoordinatorState(context.Context, *QueryCoordinatorStateRequest) (*QueryCoordinatorStateResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) CoordinatorState(ctx context.Context, req *QueryCoordinatorStateRequest) (*QueryCoordinatorStateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CoordinatorState not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_CoordinatorState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryCoordinatorStateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).CoordinatorState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cross.core.atomic.Query/CoordinatorState",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).CoordinatorState(ctx, req.(*QueryCoordinatorStateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cross.core.atomic.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CoordinatorState",
			Handler:    _Query_CoordinatorState_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cross/core/atomic/query.proto",
}

func (m *QueryCoordinatorStateRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryCoordinatorStateRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryCoordinatorStateRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.TxId) > 0 {
		i -= len(m.TxId)
		copy(dAtA[i:], m.TxId)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.TxId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryCoordinatorStateResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryCoordinatorStateResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryCoordinatorStateResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.CoodinatorState.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintQuery(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *QueryCoordinatorStateRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.TxId)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryCoordinatorStateResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.CoodinatorState.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryCoordinatorStateRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryCoordinatorStateRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryCoordinatorStateRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TxId", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TxId = append(m.TxId[:0], dAtA[iNdEx:postIndex]...)
			if m.TxId == nil {
				m.TxId = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryCoordinatorStateResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryCoordinatorStateResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryCoordinatorStateResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CoodinatorState", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.CoodinatorState.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthQuery
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
func skipQuery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
				return 0, ErrInvalidLengthQuery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuery = fmt.Errorf("proto: unexpected end of group")
)