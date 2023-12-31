// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: samplemod/auth.proto

package types

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type SampleAuthExtension struct {
}

func (m *SampleAuthExtension) Reset()         { *m = SampleAuthExtension{} }
func (m *SampleAuthExtension) String() string { return proto.CompactTextString(m) }
func (*SampleAuthExtension) ProtoMessage()    {}
func (*SampleAuthExtension) Descriptor() ([]byte, []int) {
	return fileDescriptor_542029488d515869, []int{0}
}
func (m *SampleAuthExtension) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SampleAuthExtension) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SampleAuthExtension.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SampleAuthExtension) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SampleAuthExtension.Merge(m, src)
}
func (m *SampleAuthExtension) XXX_Size() int {
	return m.Size()
}
func (m *SampleAuthExtension) XXX_DiscardUnknown() {
	xxx_messageInfo_SampleAuthExtension.DiscardUnknown(m)
}

var xxx_messageInfo_SampleAuthExtension proto.InternalMessageInfo

func init() {
	proto.RegisterType((*SampleAuthExtension)(nil), "samplemod.types.SampleAuthExtension")
}

func init() { proto.RegisterFile("samplemod/auth.proto", fileDescriptor_542029488d515869) }

var fileDescriptor_542029488d515869 = []byte{
	// 195 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x44, 0x8e, 0x31, 0x6a, 0xc3, 0x30,
	0x14, 0x86, 0xe5, 0xa5, 0x83, 0x97, 0x42, 0xeb, 0x2e, 0xa6, 0x68, 0xe8, 0x01, 0xfc, 0x86, 0xf6,
	0x02, 0x0d, 0xe4, 0x02, 0xc9, 0x96, 0xed, 0xd9, 0x16, 0x92, 0xc0, 0xd2, 0x13, 0xd6, 0x33, 0x24,
	0xb7, 0xc8, 0xb1, 0x3c, 0x7a, 0xcc, 0x98, 0xd8, 0x17, 0x09, 0x91, 0x21, 0xd9, 0x3e, 0x3e, 0x1e,
	0xef, 0xff, 0xf2, 0x22, 0xa2, 0x0b, 0x9d, 0x72, 0xd4, 0x02, 0x0e, 0x6c, 0xaa, 0xd0, 0x13, 0xd3,
	0xc7, 0xfb, 0xd3, 0x56, 0x7c, 0x0a, 0x2a, 0x96, 0xdf, 0x9a, 0x48, 0x77, 0x0a, 0x30, 0x58, 0x40,
	0xef, 0x89, 0x91, 0x2d, 0xf9, 0xb8, 0x9e, 0x97, 0x85, 0x26, 0x4d, 0x09, 0xe1, 0x41, 0xab, 0xfd,
	0xf9, 0xca, 0x3f, 0xf7, 0xe9, 0xcd, 0xff, 0xc0, 0x66, 0x7b, 0x64, 0xe5, 0xa3, 0x25, 0xbf, 0xd9,
	0x8d, 0x37, 0x29, 0xc6, 0x59, 0x66, 0xd3, 0x2c, 0xb3, 0xeb, 0x2c, 0xb3, 0xf3, 0x22, 0xc5, 0xb4,
	0x48, 0x71, 0x59, 0xa4, 0x38, 0xfc, 0x69, 0xcb, 0x66, 0xa8, 0xab, 0x86, 0x1c, 0xb4, 0xc8, 0xd8,
	0x18, 0xb4, 0xbe, 0xc3, 0x1a, 0x9a, 0x9e, 0x62, 0x84, 0x68, 0x1d, 0x86, 0x00, 0xaf, 0xe8, 0x94,
	0x57, 0xbf, 0xa5, 0xc5, 0xdf, 0x7b, 0x00, 0x00, 0x00, 0xff, 0xff, 0x22, 0x77, 0x9a, 0x58, 0xce,
	0x00, 0x00, 0x00,
}

func (m *SampleAuthExtension) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SampleAuthExtension) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SampleAuthExtension) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintAuth(dAtA []byte, offset int, v uint64) int {
	offset -= sovAuth(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *SampleAuthExtension) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovAuth(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozAuth(x uint64) (n int) {
	return sovAuth(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *SampleAuthExtension) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAuth
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
			return fmt.Errorf("proto: SampleAuthExtension: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SampleAuthExtension: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipAuth(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthAuth
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthAuth
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
func skipAuth(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowAuth
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
					return 0, ErrIntOverflowAuth
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
					return 0, ErrIntOverflowAuth
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
				return 0, ErrInvalidLengthAuth
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupAuth
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthAuth
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthAuth        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowAuth          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupAuth = fmt.Errorf("proto: unexpected end of group")
)
