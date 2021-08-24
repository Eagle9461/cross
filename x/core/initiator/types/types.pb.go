// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cross/core/initiator/types.proto

package types

import (
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/codec/types"
	github_com_datachainlab_cross_x_core_auth_types "github.com/datachainlab/cross/x/core/auth/types"
	github_com_datachainlab_cross_x_core_tx_types "github.com/datachainlab/cross/x/core/tx/types"
	types1 "github.com/datachainlab/cross/x/core/tx/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/tendermint/tendermint/abci/types"
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

type ContractTransaction struct {
	CrossChainChannel *types.Any                                                     `protobuf:"bytes,1,opt,name=cross_chain_channel,json=crossChainChannel,proto3" json:"cross_chain_channel,omitempty"`
	Signers           []github_com_datachainlab_cross_x_core_auth_types.AccountID    `protobuf:"bytes,2,rep,name=signers,proto3,casttype=github.com/datachainlab/cross/x/core/auth/types.AccountID" json:"signers,omitempty"`
	CallInfo          github_com_datachainlab_cross_x_core_tx_types.ContractCallInfo `protobuf:"bytes,3,opt,name=call_info,json=callInfo,proto3,casttype=github.com/datachainlab/cross/x/core/tx/types.ContractCallInfo" json:"call_info,omitempty"`
	ReturnValue       *types1.ReturnValue                                            `protobuf:"bytes,4,opt,name=return_value,json=returnValue,proto3" json:"return_value,omitempty"`
	Links             []Link                                                         `protobuf:"bytes,5,rep,name=links,proto3" json:"links"`
}

func (m *ContractTransaction) Reset()         { *m = ContractTransaction{} }
func (m *ContractTransaction) String() string { return proto.CompactTextString(m) }
func (*ContractTransaction) ProtoMessage()    {}
func (*ContractTransaction) Descriptor() ([]byte, []int) {
	return fileDescriptor_8a6f064a72728169, []int{0}
}
func (m *ContractTransaction) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ContractTransaction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ContractTransaction.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ContractTransaction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContractTransaction.Merge(m, src)
}
func (m *ContractTransaction) XXX_Size() int {
	return m.Size()
}
func (m *ContractTransaction) XXX_DiscardUnknown() {
	xxx_messageInfo_ContractTransaction.DiscardUnknown(m)
}

var xxx_messageInfo_ContractTransaction proto.InternalMessageInfo

type Link struct {
	SrcIndex uint32 `protobuf:"varint,1,opt,name=src_index,json=srcIndex,proto3" json:"src_index,omitempty"`
}

func (m *Link) Reset()         { *m = Link{} }
func (m *Link) String() string { return proto.CompactTextString(m) }
func (*Link) ProtoMessage()    {}
func (*Link) Descriptor() ([]byte, []int) {
	return fileDescriptor_8a6f064a72728169, []int{1}
}
func (m *Link) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Link) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Link.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Link) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Link.Merge(m, src)
}
func (m *Link) XXX_Size() int {
	return m.Size()
}
func (m *Link) XXX_DiscardUnknown() {
	xxx_messageInfo_Link.DiscardUnknown(m)
}

var xxx_messageInfo_Link proto.InternalMessageInfo

type GenesisState struct {
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_8a6f064a72728169, []int{2}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func init() {
	proto.RegisterType((*ContractTransaction)(nil), "cross.core.initiator.ContractTransaction")
	proto.RegisterType((*Link)(nil), "cross.core.initiator.Link")
	proto.RegisterType((*GenesisState)(nil), "cross.core.initiator.GenesisState")
}

func init() { proto.RegisterFile("cross/core/initiator/types.proto", fileDescriptor_8a6f064a72728169) }

var fileDescriptor_8a6f064a72728169 = []byte{
	// 456 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0x41, 0x8b, 0xd3, 0x40,
	0x14, 0xc7, 0x13, 0xdb, 0xd5, 0xdd, 0xd9, 0x2a, 0x98, 0xed, 0x21, 0xdb, 0x85, 0x34, 0xd4, 0x4b,
	0x4f, 0x13, 0x50, 0x11, 0x14, 0x56, 0xd8, 0x54, 0x90, 0x8a, 0xa7, 0x28, 0x0a, 0x5e, 0xc2, 0x64,
	0x3a, 0x4d, 0x87, 0xcd, 0xbe, 0x59, 0x66, 0x5e, 0xa4, 0xfd, 0x16, 0x1e, 0x3d, 0xfa, 0x71, 0x7a,
	0xdc, 0xa3, 0xa7, 0xa2, 0xed, 0xc5, 0xcf, 0xb0, 0x27, 0xc9, 0xa4, 0xed, 0x96, 0x3d, 0xf5, 0x12,
	0xe6, 0xf1, 0x7e, 0xff, 0x99, 0xff, 0x7b, 0xff, 0x90, 0x90, 0x6b, 0x65, 0x4c, 0xc4, 0x95, 0x16,
	0x91, 0x04, 0x89, 0x92, 0xa1, 0xd2, 0x11, 0xce, 0xae, 0x85, 0xa1, 0xd7, 0x5a, 0xa1, 0xf2, 0xda,
	0x96, 0xa0, 0x15, 0x41, 0xb7, 0x44, 0xa7, 0x9d, 0xab, 0x5c, 0x59, 0x20, 0xaa, 0x4e, 0x35, 0xdb,
	0x39, 0xcd, 0x95, 0xca, 0x0b, 0x11, 0xd9, 0x2a, 0x2b, 0xc7, 0x11, 0x83, 0xd9, 0xba, 0x75, 0x86,
	0x02, 0x46, 0x42, 0x5f, 0x49, 0xc0, 0x88, 0x65, 0x5c, 0xee, 0xbe, 0xd1, 0x39, 0xdd, 0x71, 0x81,
	0xd3, 0xdd, 0x56, 0xef, 0x67, 0x83, 0x9c, 0x0c, 0x14, 0xa0, 0x66, 0x1c, 0x3f, 0x6b, 0x06, 0x86,
	0x71, 0x94, 0x0a, 0xbc, 0x0f, 0xe4, 0xc4, 0x8a, 0x52, 0x3e, 0x61, 0x12, 0xaa, 0x2f, 0x80, 0x28,
	0x7c, 0x37, 0x74, 0xfb, 0xc7, 0xcf, 0xdb, 0xb4, 0x36, 0x42, 0x37, 0x46, 0xe8, 0x05, 0xcc, 0xe2,
	0xe6, 0x7c, 0xd1, 0x75, 0x93, 0xa7, 0x56, 0x36, 0xa8, 0x54, 0x83, 0x5a, 0xe4, 0x7d, 0x25, 0x8f,
	0x8c, 0xcc, 0x41, 0x68, 0xe3, 0x3f, 0x08, 0x1b, 0xfd, 0x56, 0x7c, 0x7e, 0xbb, 0xe8, 0xbe, 0xce,
	0x25, 0x4e, 0xca, 0x8c, 0x72, 0x75, 0x15, 0x8d, 0x18, 0x32, 0xfb, 0x4e, 0xc1, 0xb2, 0xa8, 0xf6,
	0x3a, 0xad, 0xdd, 0xb2, 0x12, 0x27, 0x6b, 0xbf, 0x17, 0x9c, 0xab, 0x12, 0x70, 0xf8, 0x2e, 0xd9,
	0xdc, 0xe6, 0xa5, 0xe4, 0x88, 0xb3, 0xa2, 0x48, 0x25, 0x8c, 0x95, 0xdf, 0x08, 0xdd, 0x7e, 0x2b,
	0x8e, 0x6f, 0x17, 0xdd, 0xb7, 0x7b, 0x5d, 0xbd, 0x5d, 0xc4, 0x66, 0xfc, 0x01, 0x2b, 0x8a, 0x21,
	0x8c, 0x55, 0x72, 0xc8, 0xd7, 0x27, 0xef, 0x9c, 0xb4, 0xb4, 0xc0, 0x52, 0x43, 0xfa, 0x9d, 0x15,
	0xa5, 0xf0, 0x9b, 0x76, 0xfc, 0x0e, 0xdd, 0xc9, 0x0c, 0xa7, 0x34, 0xb1, 0xc8, 0x97, 0x8a, 0x48,
	0x8e, 0xf5, 0x5d, 0xe1, 0xbd, 0x22, 0x07, 0x85, 0x84, 0x4b, 0xe3, 0x1f, 0x84, 0x8d, 0xfb, 0xba,
	0x6d, 0xd6, 0xf4, 0xa3, 0x84, 0x4b, 0xbb, 0x3c, 0x27, 0xa9, 0xf1, 0x37, 0xcd, 0x7f, 0xbf, 0xba,
	0x4e, 0xef, 0x19, 0x69, 0x56, 0x2d, 0xef, 0x8c, 0x1c, 0x19, 0xcd, 0x53, 0x09, 0x23, 0x31, 0xb5,
	0x01, 0x3c, 0x4e, 0x0e, 0x8d, 0xe6, 0xc3, 0xaa, 0xee, 0x3d, 0x21, 0xad, 0xf7, 0x02, 0x84, 0x91,
	0xe6, 0x13, 0x32, 0x14, 0x71, 0x32, 0xff, 0x1b, 0x38, 0xf3, 0x65, 0xe0, 0xde, 0x2c, 0x03, 0xf7,
	0xcf, 0x32, 0x70, 0x7f, 0xac, 0x02, 0xe7, 0x66, 0x15, 0x38, 0xbf, 0x57, 0x81, 0xf3, 0xed, 0xe5,
	0x5e, 0x9b, 0xb9, 0xf7, 0xa3, 0x66, 0x0f, 0x6d, 0xcc, 0x2f, 0xfe, 0x07, 0x00, 0x00, 0xff, 0xff,
	0xf7, 0x43, 0x7b, 0x28, 0xcd, 0x02, 0x00, 0x00,
}

func (m *ContractTransaction) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ContractTransaction) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ContractTransaction) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Links) > 0 {
		for iNdEx := len(m.Links) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Links[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTypes(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x2a
		}
	}
	if m.ReturnValue != nil {
		{
			size, err := m.ReturnValue.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTypes(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	if len(m.CallInfo) > 0 {
		i -= len(m.CallInfo)
		copy(dAtA[i:], m.CallInfo)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.CallInfo)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Signers) > 0 {
		for iNdEx := len(m.Signers) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Signers[iNdEx])
			copy(dAtA[i:], m.Signers[iNdEx])
			i = encodeVarintTypes(dAtA, i, uint64(len(m.Signers[iNdEx])))
			i--
			dAtA[i] = 0x12
		}
	}
	if m.CrossChainChannel != nil {
		{
			size, err := m.CrossChainChannel.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTypes(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Link) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Link) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Link) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.SrcIndex != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.SrcIndex))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintTypes(dAtA []byte, offset int, v uint64) int {
	offset -= sovTypes(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ContractTransaction) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.CrossChainChannel != nil {
		l = m.CrossChainChannel.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	if len(m.Signers) > 0 {
		for _, b := range m.Signers {
			l = len(b)
			n += 1 + l + sovTypes(uint64(l))
		}
	}
	l = len(m.CallInfo)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	if m.ReturnValue != nil {
		l = m.ReturnValue.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	if len(m.Links) > 0 {
		for _, e := range m.Links {
			l = e.Size()
			n += 1 + l + sovTypes(uint64(l))
		}
	}
	return n
}

func (m *Link) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.SrcIndex != 0 {
		n += 1 + sovTypes(uint64(m.SrcIndex))
	}
	return n
}

func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovTypes(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTypes(x uint64) (n int) {
	return sovTypes(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ContractTransaction) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: ContractTransaction: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ContractTransaction: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CrossChainChannel", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.CrossChainChannel == nil {
				m.CrossChainChannel = &types.Any{}
			}
			if err := m.CrossChainChannel.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Signers", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Signers = append(m.Signers, make([]byte, postIndex-iNdEx))
			copy(m.Signers[len(m.Signers)-1], dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CallInfo", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CallInfo = append(m.CallInfo[:0], dAtA[iNdEx:postIndex]...)
			if m.CallInfo == nil {
				m.CallInfo = []byte{}
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ReturnValue", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ReturnValue == nil {
				m.ReturnValue = &types1.ReturnValue{}
			}
			if err := m.ReturnValue.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Links", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Links = append(m.Links, Link{})
			if err := m.Links[len(m.Links)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTypes
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
func (m *Link) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: Link: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Link: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SrcIndex", wireType)
			}
			m.SrcIndex = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SrcIndex |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTypes
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
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTypes
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
func skipTypes(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTypes
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
					return 0, ErrIntOverflowTypes
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
					return 0, ErrIntOverflowTypes
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
				return 0, ErrInvalidLengthTypes
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTypes
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTypes
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTypes        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTypes          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTypes = fmt.Errorf("proto: unexpected end of group")
)
