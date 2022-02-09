// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: gentlemint/exchange_rate.proto

package types

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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

type ExchangeRate struct {
	ShrpToShr string `protobuf:"bytes,1,opt,name=shrpToShr,proto3" json:"shrpToShr,omitempty"`
}

func (m *ExchangeRate) Reset()         { *m = ExchangeRate{} }
func (m *ExchangeRate) String() string { return proto.CompactTextString(m) }
func (*ExchangeRate) ProtoMessage()    {}
func (*ExchangeRate) Descriptor() ([]byte, []int) {
	return fileDescriptor_e32923d681230718, []int{0}
}
func (m *ExchangeRate) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ExchangeRate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ExchangeRate.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ExchangeRate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExchangeRate.Merge(m, src)
}
func (m *ExchangeRate) XXX_Size() int {
	return m.Size()
}
func (m *ExchangeRate) XXX_DiscardUnknown() {
	xxx_messageInfo_ExchangeRate.DiscardUnknown(m)
}

var xxx_messageInfo_ExchangeRate proto.InternalMessageInfo

func (m *ExchangeRate) GetShrpToShr() string {
	if m != nil {
		return m.ShrpToShr
	}
	return ""
}

func init() {
	proto.RegisterType((*ExchangeRate)(nil), "shareledger.gentlemint.ExchangeRate")
}

func init() { proto.RegisterFile("gentlemint/exchange_rate.proto", fileDescriptor_e32923d681230718) }

var fileDescriptor_e32923d681230718 = []byte{
	// 184 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4b, 0x4f, 0xcd, 0x2b,
	0xc9, 0x49, 0xcd, 0xcd, 0xcc, 0x2b, 0xd1, 0x4f, 0xad, 0x48, 0xce, 0x48, 0xcc, 0x4b, 0x4f, 0x8d,
	0x2f, 0x4a, 0x2c, 0x49, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x2b, 0xce, 0x48, 0x2c,
	0x4a, 0xcd, 0x49, 0x4d, 0x49, 0x4f, 0x2d, 0xd2, 0x43, 0xa8, 0x95, 0x12, 0x49, 0xcf, 0x4f, 0xcf,
	0x07, 0x2b, 0xd1, 0x07, 0xb1, 0x20, 0xaa, 0x95, 0x74, 0xb8, 0x78, 0x5c, 0xa1, 0x86, 0x04, 0x25,
	0x96, 0xa4, 0x0a, 0xc9, 0x70, 0x71, 0x16, 0x67, 0x14, 0x15, 0x84, 0xe4, 0x07, 0x67, 0x14, 0x49,
	0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0x21, 0x04, 0x9c, 0x7c, 0x4f, 0x3c, 0x92, 0x63, 0xbc, 0xf0,
	0x48, 0x8e, 0xf1, 0xc1, 0x23, 0x39, 0xc6, 0x09, 0x8f, 0xe5, 0x18, 0x2e, 0x3c, 0x96, 0x63, 0xb8,
	0xf1, 0x58, 0x8e, 0x21, 0xca, 0x38, 0x3d, 0xb3, 0x24, 0xa3, 0x34, 0x49, 0x2f, 0x39, 0x3f, 0x57,
	0x1f, 0xec, 0x80, 0xa2, 0xcc, 0xbc, 0x74, 0x7d, 0x24, 0xa7, 0xe8, 0x57, 0xe8, 0x23, 0x39, 0xbc,
	0xa4, 0xb2, 0x20, 0xb5, 0x38, 0x89, 0x0d, 0xec, 0x06, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x6a, 0x6f, 0x25, 0x74, 0xd3, 0x00, 0x00, 0x00,
}

func (m *ExchangeRate) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ExchangeRate) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ExchangeRate) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ShrpToShr) > 0 {
		i -= len(m.ShrpToShr)
		copy(dAtA[i:], m.ShrpToShr)
		i = encodeVarintExchangeRate(dAtA, i, uint64(len(m.ShrpToShr)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintExchangeRate(dAtA []byte, offset int, v uint64) int {
	offset -= sovExchangeRate(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ExchangeRate) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ShrpToShr)
	if l > 0 {
		n += 1 + l + sovExchangeRate(uint64(l))
	}
	return n
}

func sovExchangeRate(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozExchangeRate(x uint64) (n int) {
	return sovExchangeRate(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ExchangeRate) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowExchangeRate
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
			return fmt.Errorf("proto: ExchangeRate: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ExchangeRate: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ShrpToShr", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExchangeRate
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthExchangeRate
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthExchangeRate
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ShrpToShr = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipExchangeRate(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthExchangeRate
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
func skipExchangeRate(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowExchangeRate
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
					return 0, ErrIntOverflowExchangeRate
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
					return 0, ErrIntOverflowExchangeRate
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
				return 0, ErrInvalidLengthExchangeRate
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupExchangeRate
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthExchangeRate
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthExchangeRate        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowExchangeRate          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupExchangeRate = fmt.Errorf("proto: unexpected end of group")
)
