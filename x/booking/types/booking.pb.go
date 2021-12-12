// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: booking/booking.proto

package types

import (
	fmt "fmt"
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

type Booking struct {
	BookID      string `protobuf:"bytes,1,opt,name=bookID,proto3" json:"bookID,omitempty"`
	Booker      string `protobuf:"bytes,2,opt,name=booker,proto3" json:"booker,omitempty"`
	UUID        string `protobuf:"bytes,3,opt,name=UUID,proto3" json:"UUID,omitempty"`
	Duration    int64  `protobuf:"varint,4,opt,name=duration,proto3" json:"duration,omitempty"`
	IsCompleted bool   `protobuf:"varint,5,opt,name=isCompleted,proto3" json:"isCompleted,omitempty"`
}

func (m *Booking) Reset()         { *m = Booking{} }
func (m *Booking) String() string { return proto.CompactTextString(m) }
func (*Booking) ProtoMessage()    {}
func (*Booking) Descriptor() ([]byte, []int) {
	return fileDescriptor_00ffdf80b5823d65, []int{0}
}
func (m *Booking) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Booking) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Booking.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Booking) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Booking.Merge(m, src)
}
func (m *Booking) XXX_Size() int {
	return m.Size()
}
func (m *Booking) XXX_DiscardUnknown() {
	xxx_messageInfo_Booking.DiscardUnknown(m)
}

var xxx_messageInfo_Booking proto.InternalMessageInfo

func (m *Booking) GetBookID() string {
	if m != nil {
		return m.BookID
	}
	return ""
}

func (m *Booking) GetBooker() string {
	if m != nil {
		return m.Booker
	}
	return ""
}

func (m *Booking) GetUUID() string {
	if m != nil {
		return m.UUID
	}
	return ""
}

func (m *Booking) GetDuration() int64 {
	if m != nil {
		return m.Duration
	}
	return 0
}

func (m *Booking) GetIsCompleted() bool {
	if m != nil {
		return m.IsCompleted
	}
	return false
}

func init() {
	proto.RegisterType((*Booking)(nil), "ShareRing.shareledger.booking.Booking")
}

func init() { proto.RegisterFile("booking/booking.proto", fileDescriptor_00ffdf80b5823d65) }

var fileDescriptor_00ffdf80b5823d65 = []byte{
	// 218 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4d, 0xca, 0xcf, 0xcf,
	0xce, 0xcc, 0x4b, 0xd7, 0x87, 0xd2, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42, 0xb2, 0xc1, 0x19,
	0x89, 0x45, 0xa9, 0x41, 0x20, 0x81, 0x62, 0x10, 0x2b, 0x27, 0x35, 0x25, 0x3d, 0xb5, 0x48, 0x0f,
	0xaa, 0x48, 0xa9, 0x9b, 0x91, 0x8b, 0xdd, 0x09, 0xc2, 0x16, 0x12, 0xe3, 0x62, 0x03, 0x09, 0x7b,
	0xba, 0x48, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0x41, 0x79, 0x30, 0xf1, 0xd4, 0x22, 0x09, 0x26,
	0x84, 0x78, 0x6a, 0x91, 0x90, 0x10, 0x17, 0x4b, 0x68, 0xa8, 0xa7, 0x8b, 0x04, 0x33, 0x58, 0x14,
	0xcc, 0x16, 0x92, 0xe2, 0xe2, 0x48, 0x29, 0x2d, 0x4a, 0x2c, 0xc9, 0xcc, 0xcf, 0x93, 0x60, 0x51,
	0x60, 0xd4, 0x60, 0x0e, 0x82, 0xf3, 0x85, 0x14, 0xb8, 0xb8, 0x33, 0x8b, 0x9d, 0xf3, 0x73, 0x0b,
	0x72, 0x52, 0x4b, 0x52, 0x53, 0x24, 0x58, 0x15, 0x18, 0x35, 0x38, 0x82, 0x90, 0x85, 0x9c, 0xbc,
	0x4e, 0x3c, 0x92, 0x63, 0xbc, 0xf0, 0x48, 0x8e, 0xf1, 0xc1, 0x23, 0x39, 0xc6, 0x09, 0x8f, 0xe5,
	0x18, 0x2e, 0x3c, 0x96, 0x63, 0xb8, 0xf1, 0x58, 0x8e, 0x21, 0xca, 0x20, 0x3d, 0xb3, 0x24, 0xa3,
	0x34, 0x49, 0x2f, 0x39, 0x3f, 0x57, 0x1f, 0xee, 0x23, 0x08, 0x0b, 0xe2, 0x23, 0xfd, 0x0a, 0x98,
	0xc7, 0xf5, 0x4b, 0x2a, 0x0b, 0x52, 0x8b, 0x93, 0xd8, 0xc0, 0xfe, 0x37, 0x06, 0x04, 0x00, 0x00,
	0xff, 0xff, 0xd2, 0x0f, 0x5f, 0x68, 0x18, 0x01, 0x00, 0x00,
}

func (m *Booking) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Booking) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Booking) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.IsCompleted {
		i--
		if m.IsCompleted {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x28
	}
	if m.Duration != 0 {
		i = encodeVarintBooking(dAtA, i, uint64(m.Duration))
		i--
		dAtA[i] = 0x20
	}
	if len(m.UUID) > 0 {
		i -= len(m.UUID)
		copy(dAtA[i:], m.UUID)
		i = encodeVarintBooking(dAtA, i, uint64(len(m.UUID)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Booker) > 0 {
		i -= len(m.Booker)
		copy(dAtA[i:], m.Booker)
		i = encodeVarintBooking(dAtA, i, uint64(len(m.Booker)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.BookID) > 0 {
		i -= len(m.BookID)
		copy(dAtA[i:], m.BookID)
		i = encodeVarintBooking(dAtA, i, uint64(len(m.BookID)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintBooking(dAtA []byte, offset int, v uint64) int {
	offset -= sovBooking(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Booking) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.BookID)
	if l > 0 {
		n += 1 + l + sovBooking(uint64(l))
	}
	l = len(m.Booker)
	if l > 0 {
		n += 1 + l + sovBooking(uint64(l))
	}
	l = len(m.UUID)
	if l > 0 {
		n += 1 + l + sovBooking(uint64(l))
	}
	if m.Duration != 0 {
		n += 1 + sovBooking(uint64(m.Duration))
	}
	if m.IsCompleted {
		n += 2
	}
	return n
}

func sovBooking(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozBooking(x uint64) (n int) {
	return sovBooking(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Booking) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBooking
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
			return fmt.Errorf("proto: Booking: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Booking: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BookID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBooking
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
				return ErrInvalidLengthBooking
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBooking
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BookID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Booker", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBooking
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
				return ErrInvalidLengthBooking
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBooking
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Booker = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UUID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBooking
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
				return ErrInvalidLengthBooking
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBooking
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.UUID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Duration", wireType)
			}
			m.Duration = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBooking
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Duration |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IsCompleted", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBooking
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
			m.IsCompleted = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipBooking(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthBooking
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
func skipBooking(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowBooking
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
					return 0, ErrIntOverflowBooking
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
					return 0, ErrIntOverflowBooking
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
				return 0, ErrInvalidLengthBooking
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupBooking
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthBooking
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthBooking        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowBooking          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupBooking = fmt.Errorf("proto: unexpected end of group")
)