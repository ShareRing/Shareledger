// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: shareledger/swap/v1/past_tx_event.proto

package types

import (
	fmt "fmt"
	proto "github.com/cosmos/gogoproto/proto"
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

// RequestedIns
type PastTxEvent struct {
	SrcAddr  string `protobuf:"bytes,1,opt,name=srcAddr,proto3" json:"srcAddr,omitempty"`
	DestAddr string `protobuf:"bytes,2,opt,name=destAddr,proto3" json:"destAddr,omitempty"`
}

func (m *PastTxEvent) Reset()         { *m = PastTxEvent{} }
func (m *PastTxEvent) String() string { return proto.CompactTextString(m) }
func (*PastTxEvent) ProtoMessage()    {}
func (*PastTxEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_256afd1be8870ba5, []int{0}
}
func (m *PastTxEvent) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PastTxEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PastTxEvent.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PastTxEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PastTxEvent.Merge(m, src)
}
func (m *PastTxEvent) XXX_Size() int {
	return m.Size()
}
func (m *PastTxEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_PastTxEvent.DiscardUnknown(m)
}

var xxx_messageInfo_PastTxEvent proto.InternalMessageInfo

func (m *PastTxEvent) GetSrcAddr() string {
	if m != nil {
		return m.SrcAddr
	}
	return ""
}

func (m *PastTxEvent) GetDestAddr() string {
	if m != nil {
		return m.DestAddr
	}
	return ""
}

type PastTxEventGenesis struct {
	SrcAddr  string `protobuf:"bytes,1,opt,name=srcAddr,proto3" json:"srcAddr,omitempty"`
	DestAddr string `protobuf:"bytes,2,opt,name=destAddr,proto3" json:"destAddr,omitempty"`
	TxHash   string `protobuf:"bytes,3,opt,name=txHash,proto3" json:"txHash,omitempty"`
	LogIndex uint64 `protobuf:"varint,4,opt,name=logIndex,proto3" json:"logIndex,omitempty"`
}

func (m *PastTxEventGenesis) Reset()         { *m = PastTxEventGenesis{} }
func (m *PastTxEventGenesis) String() string { return proto.CompactTextString(m) }
func (*PastTxEventGenesis) ProtoMessage()    {}
func (*PastTxEventGenesis) Descriptor() ([]byte, []int) {
	return fileDescriptor_256afd1be8870ba5, []int{1}
}
func (m *PastTxEventGenesis) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PastTxEventGenesis) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PastTxEventGenesis.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PastTxEventGenesis) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PastTxEventGenesis.Merge(m, src)
}
func (m *PastTxEventGenesis) XXX_Size() int {
	return m.Size()
}
func (m *PastTxEventGenesis) XXX_DiscardUnknown() {
	xxx_messageInfo_PastTxEventGenesis.DiscardUnknown(m)
}

var xxx_messageInfo_PastTxEventGenesis proto.InternalMessageInfo

func (m *PastTxEventGenesis) GetSrcAddr() string {
	if m != nil {
		return m.SrcAddr
	}
	return ""
}

func (m *PastTxEventGenesis) GetDestAddr() string {
	if m != nil {
		return m.DestAddr
	}
	return ""
}

func (m *PastTxEventGenesis) GetTxHash() string {
	if m != nil {
		return m.TxHash
	}
	return ""
}

func (m *PastTxEventGenesis) GetLogIndex() uint64 {
	if m != nil {
		return m.LogIndex
	}
	return 0
}

func init() {
	proto.RegisterType((*PastTxEvent)(nil), "shareledger.swap.PastTxEvent")
	proto.RegisterType((*PastTxEventGenesis)(nil), "shareledger.swap.PastTxEventGenesis")
}

func init() {
	proto.RegisterFile("shareledger/swap/v1/past_tx_event.proto", fileDescriptor_256afd1be8870ba5)
}

var fileDescriptor_256afd1be8870ba5 = []byte{
	// 234 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x2f, 0xce, 0x48, 0x2c,
	0x4a, 0xcd, 0x49, 0x4d, 0x49, 0x4f, 0x2d, 0xd2, 0x2f, 0x2e, 0x4f, 0x2c, 0xd0, 0x2f, 0x33, 0xd4,
	0x2f, 0x48, 0x2c, 0x2e, 0x89, 0x2f, 0xa9, 0x88, 0x4f, 0x2d, 0x4b, 0xcd, 0x2b, 0xd1, 0x2b, 0x28,
	0xca, 0x2f, 0xc9, 0x17, 0x12, 0x40, 0x52, 0xa8, 0x07, 0x52, 0xa8, 0xe4, 0xcc, 0xc5, 0x1d, 0x90,
	0x58, 0x5c, 0x12, 0x52, 0xe1, 0x0a, 0x52, 0x26, 0x24, 0xc1, 0xc5, 0x5e, 0x5c, 0x94, 0xec, 0x98,
	0x92, 0x52, 0x24, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0x19, 0x04, 0xe3, 0x0a, 0x49, 0x71, 0x71, 0xa4,
	0xa4, 0x16, 0x97, 0x80, 0xa5, 0x98, 0xc0, 0x52, 0x70, 0xbe, 0x52, 0x1d, 0x97, 0x10, 0x92, 0x21,
	0xee, 0xa9, 0x79, 0xa9, 0xc5, 0x99, 0xc5, 0xe4, 0x99, 0x25, 0x24, 0xc6, 0xc5, 0x56, 0x52, 0xe1,
	0x91, 0x58, 0x9c, 0x21, 0xc1, 0x0c, 0x96, 0x81, 0xf2, 0x40, 0x7a, 0x72, 0xf2, 0xd3, 0x3d, 0xf3,
	0x52, 0x52, 0x2b, 0x24, 0x58, 0x14, 0x18, 0x35, 0x58, 0x82, 0xe0, 0x7c, 0x27, 0xf7, 0x13, 0x8f,
	0xe4, 0x18, 0x2f, 0x3c, 0x92, 0x63, 0x7c, 0xf0, 0x48, 0x8e, 0x71, 0xc2, 0x63, 0x39, 0x86, 0x0b,
	0x8f, 0xe5, 0x18, 0x6e, 0x3c, 0x96, 0x63, 0x88, 0xd2, 0x4d, 0xcf, 0x2c, 0xc9, 0x28, 0x4d, 0xd2,
	0x4b, 0xce, 0xcf, 0xd5, 0x07, 0xfb, 0xbd, 0x28, 0x33, 0x2f, 0x5d, 0x1f, 0x39, 0xb8, 0x2a, 0x20,
	0x01, 0x56, 0x52, 0x59, 0x90, 0x5a, 0x9c, 0xc4, 0x06, 0x0e, 0x26, 0x63, 0x40, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x94, 0x9b, 0xa6, 0xf1, 0x51, 0x01, 0x00, 0x00,
}

func (m *PastTxEvent) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PastTxEvent) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PastTxEvent) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.DestAddr) > 0 {
		i -= len(m.DestAddr)
		copy(dAtA[i:], m.DestAddr)
		i = encodeVarintPastTxEvent(dAtA, i, uint64(len(m.DestAddr)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.SrcAddr) > 0 {
		i -= len(m.SrcAddr)
		copy(dAtA[i:], m.SrcAddr)
		i = encodeVarintPastTxEvent(dAtA, i, uint64(len(m.SrcAddr)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *PastTxEventGenesis) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PastTxEventGenesis) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PastTxEventGenesis) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.LogIndex != 0 {
		i = encodeVarintPastTxEvent(dAtA, i, uint64(m.LogIndex))
		i--
		dAtA[i] = 0x20
	}
	if len(m.TxHash) > 0 {
		i -= len(m.TxHash)
		copy(dAtA[i:], m.TxHash)
		i = encodeVarintPastTxEvent(dAtA, i, uint64(len(m.TxHash)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.DestAddr) > 0 {
		i -= len(m.DestAddr)
		copy(dAtA[i:], m.DestAddr)
		i = encodeVarintPastTxEvent(dAtA, i, uint64(len(m.DestAddr)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.SrcAddr) > 0 {
		i -= len(m.SrcAddr)
		copy(dAtA[i:], m.SrcAddr)
		i = encodeVarintPastTxEvent(dAtA, i, uint64(len(m.SrcAddr)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintPastTxEvent(dAtA []byte, offset int, v uint64) int {
	offset -= sovPastTxEvent(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *PastTxEvent) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.SrcAddr)
	if l > 0 {
		n += 1 + l + sovPastTxEvent(uint64(l))
	}
	l = len(m.DestAddr)
	if l > 0 {
		n += 1 + l + sovPastTxEvent(uint64(l))
	}
	return n
}

func (m *PastTxEventGenesis) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.SrcAddr)
	if l > 0 {
		n += 1 + l + sovPastTxEvent(uint64(l))
	}
	l = len(m.DestAddr)
	if l > 0 {
		n += 1 + l + sovPastTxEvent(uint64(l))
	}
	l = len(m.TxHash)
	if l > 0 {
		n += 1 + l + sovPastTxEvent(uint64(l))
	}
	if m.LogIndex != 0 {
		n += 1 + sovPastTxEvent(uint64(m.LogIndex))
	}
	return n
}

func sovPastTxEvent(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozPastTxEvent(x uint64) (n int) {
	return sovPastTxEvent(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *PastTxEvent) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPastTxEvent
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
			return fmt.Errorf("proto: PastTxEvent: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PastTxEvent: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SrcAddr", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPastTxEvent
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
				return ErrInvalidLengthPastTxEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPastTxEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SrcAddr = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DestAddr", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPastTxEvent
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
				return ErrInvalidLengthPastTxEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPastTxEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DestAddr = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPastTxEvent(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPastTxEvent
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
func (m *PastTxEventGenesis) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPastTxEvent
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
			return fmt.Errorf("proto: PastTxEventGenesis: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PastTxEventGenesis: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SrcAddr", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPastTxEvent
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
				return ErrInvalidLengthPastTxEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPastTxEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SrcAddr = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DestAddr", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPastTxEvent
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
				return ErrInvalidLengthPastTxEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPastTxEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DestAddr = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TxHash", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPastTxEvent
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
				return ErrInvalidLengthPastTxEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPastTxEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TxHash = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LogIndex", wireType)
			}
			m.LogIndex = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPastTxEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LogIndex |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipPastTxEvent(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPastTxEvent
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
func skipPastTxEvent(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPastTxEvent
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
					return 0, ErrIntOverflowPastTxEvent
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
					return 0, ErrIntOverflowPastTxEvent
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
				return 0, ErrInvalidLengthPastTxEvent
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupPastTxEvent
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthPastTxEvent
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthPastTxEvent        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPastTxEvent          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupPastTxEvent = fmt.Errorf("proto: unexpected end of group")
)
