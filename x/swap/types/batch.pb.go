// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: swap/batch.proto

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

type Batch struct {
	Id        uint64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Signature string   `protobuf:"bytes,2,opt,name=signature,proto3" json:"signature,omitempty"`
	TxIds     []uint64 `protobuf:"varint,3,rep,packed,name=txIds,proto3" json:"txIds,omitempty"`
	Status    string   `protobuf:"bytes,4,opt,name=status,proto3" json:"status,omitempty"`
	TxHash    string   `protobuf:"bytes,6,opt,name=txHash,proto3" json:"txHash,omitempty"`
	Network   string   `protobuf:"bytes,7,opt,name=network,proto3" json:"network,omitempty"`
}

func (m *Batch) Reset()         { *m = Batch{} }
func (m *Batch) String() string { return proto.CompactTextString(m) }
func (*Batch) ProtoMessage()    {}
func (*Batch) Descriptor() ([]byte, []int) {
	return fileDescriptor_0fefcd8eb88292cc, []int{0}
}
func (m *Batch) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Batch) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Batch.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Batch) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Batch.Merge(m, src)
}
func (m *Batch) XXX_Size() int {
	return m.Size()
}
func (m *Batch) XXX_DiscardUnknown() {
	xxx_messageInfo_Batch.DiscardUnknown(m)
}

var xxx_messageInfo_Batch proto.InternalMessageInfo

func (m *Batch) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Batch) GetSignature() string {
	if m != nil {
		return m.Signature
	}
	return ""
}

func (m *Batch) GetTxIds() []uint64 {
	if m != nil {
		return m.TxIds
	}
	return nil
}

func (m *Batch) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *Batch) GetTxHash() string {
	if m != nil {
		return m.TxHash
	}
	return ""
}

func (m *Batch) GetNetwork() string {
	if m != nil {
		return m.Network
	}
	return ""
}

func init() {
	proto.RegisterType((*Batch)(nil), "shareledger.swap.Batch")
}

func init() { proto.RegisterFile("swap/batch.proto", fileDescriptor_0fefcd8eb88292cc) }

var fileDescriptor_0fefcd8eb88292cc = []byte{
	// 233 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x8f, 0x31, 0x4e, 0xc3, 0x30,
	0x14, 0x86, 0xe3, 0x34, 0x4d, 0x55, 0x0f, 0xa8, 0xb2, 0x10, 0xf2, 0x80, 0xac, 0x88, 0x29, 0x0b,
	0xf1, 0xc0, 0x0d, 0xba, 0x00, 0x6b, 0x46, 0x36, 0xa7, 0xb1, 0x62, 0x0b, 0x88, 0x23, 0xbf, 0x17,
	0x35, 0xdc, 0x82, 0x85, 0x3b, 0x31, 0x76, 0x64, 0x44, 0xc9, 0x45, 0x50, 0x1c, 0x10, 0x1d, 0xbf,
	0xef, 0x7f, 0x6f, 0xf8, 0xe8, 0x0e, 0x8e, 0xaa, 0x93, 0x95, 0xc2, 0x83, 0x29, 0x3a, 0xef, 0xd0,
	0xb1, 0x1d, 0x18, 0xe5, 0xf5, 0x8b, 0xae, 0x1b, 0xed, 0x8b, 0x79, 0xbd, 0xf9, 0x20, 0x74, 0xbd,
	0x9f, 0x2f, 0xd8, 0x05, 0x8d, 0x6d, 0xcd, 0x49, 0x46, 0xf2, 0xa4, 0x8c, 0x6d, 0xcd, 0xae, 0xe9,
	0x16, 0x6c, 0xd3, 0x2a, 0xec, 0xbd, 0xe6, 0x71, 0x46, 0xf2, 0x6d, 0xf9, 0x2f, 0xd8, 0x25, 0x5d,
	0xe3, 0xf0, 0x58, 0x03, 0x5f, 0x65, 0xab, 0x3c, 0x29, 0x17, 0x60, 0x57, 0x34, 0x05, 0x54, 0xd8,
	0x03, 0x4f, 0xc2, 0xc3, 0x2f, 0xcd, 0x1e, 0x87, 0x07, 0x05, 0x86, 0xa7, 0x8b, 0x5f, 0x88, 0x71,
	0xba, 0x69, 0x35, 0x1e, 0x9d, 0x7f, 0xe6, 0x9b, 0x30, 0xfc, 0xe1, 0xfe, 0xfe, 0x73, 0x14, 0xe4,
	0x34, 0x0a, 0xf2, 0x3d, 0x0a, 0xf2, 0x3e, 0x89, 0xe8, 0x34, 0x89, 0xe8, 0x6b, 0x12, 0xd1, 0xd3,
	0x6d, 0x63, 0xd1, 0xf4, 0x55, 0x71, 0x70, 0xaf, 0x32, 0xe4, 0x78, 0xdb, 0x36, 0xf2, 0x2c, 0x4c,
	0x0e, 0x32, 0x84, 0xe3, 0x5b, 0xa7, 0xa1, 0x4a, 0x43, 0xf9, 0xdd, 0x4f, 0x00, 0x00, 0x00, 0xff,
	0xff, 0xb5, 0xc7, 0x8c, 0xfa, 0x0d, 0x01, 0x00, 0x00,
}

func (m *Batch) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Batch) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Batch) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Network) > 0 {
		i -= len(m.Network)
		copy(dAtA[i:], m.Network)
		i = encodeVarintBatch(dAtA, i, uint64(len(m.Network)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.TxHash) > 0 {
		i -= len(m.TxHash)
		copy(dAtA[i:], m.TxHash)
		i = encodeVarintBatch(dAtA, i, uint64(len(m.TxHash)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.Status) > 0 {
		i -= len(m.Status)
		copy(dAtA[i:], m.Status)
		i = encodeVarintBatch(dAtA, i, uint64(len(m.Status)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.TxIds) > 0 {
		dAtA2 := make([]byte, len(m.TxIds)*10)
		var j1 int
		for _, num := range m.TxIds {
			for num >= 1<<7 {
				dAtA2[j1] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j1++
			}
			dAtA2[j1] = uint8(num)
			j1++
		}
		i -= j1
		copy(dAtA[i:], dAtA2[:j1])
		i = encodeVarintBatch(dAtA, i, uint64(j1))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Signature) > 0 {
		i -= len(m.Signature)
		copy(dAtA[i:], m.Signature)
		i = encodeVarintBatch(dAtA, i, uint64(len(m.Signature)))
		i--
		dAtA[i] = 0x12
	}
	if m.Id != 0 {
		i = encodeVarintBatch(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintBatch(dAtA []byte, offset int, v uint64) int {
	offset -= sovBatch(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Batch) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovBatch(uint64(m.Id))
	}
	l = len(m.Signature)
	if l > 0 {
		n += 1 + l + sovBatch(uint64(l))
	}
	if len(m.TxIds) > 0 {
		l = 0
		for _, e := range m.TxIds {
			l += sovBatch(uint64(e))
		}
		n += 1 + sovBatch(uint64(l)) + l
	}
	l = len(m.Status)
	if l > 0 {
		n += 1 + l + sovBatch(uint64(l))
	}
	l = len(m.TxHash)
	if l > 0 {
		n += 1 + l + sovBatch(uint64(l))
	}
	l = len(m.Network)
	if l > 0 {
		n += 1 + l + sovBatch(uint64(l))
	}
	return n
}

func sovBatch(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozBatch(x uint64) (n int) {
	return sovBatch(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Batch) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBatch
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
			return fmt.Errorf("proto: Batch: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Batch: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBatch
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Signature", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBatch
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
				return ErrInvalidLengthBatch
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBatch
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Signature = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType == 0 {
				var v uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowBatch
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.TxIds = append(m.TxIds, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowBatch
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthBatch
				}
				postIndex := iNdEx + packedLen
				if postIndex < 0 {
					return ErrInvalidLengthBatch
				}
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				var elementCount int
				var count int
				for _, integer := range dAtA[iNdEx:postIndex] {
					if integer < 128 {
						count++
					}
				}
				elementCount = count
				if elementCount != 0 && len(m.TxIds) == 0 {
					m.TxIds = make([]uint64, 0, elementCount)
				}
				for iNdEx < postIndex {
					var v uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowBatch
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.TxIds = append(m.TxIds, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field TxIds", wireType)
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBatch
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
				return ErrInvalidLengthBatch
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBatch
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Status = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TxHash", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBatch
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
				return ErrInvalidLengthBatch
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBatch
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TxHash = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Network", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBatch
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
				return ErrInvalidLengthBatch
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBatch
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Network = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipBatch(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthBatch
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
func skipBatch(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowBatch
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
					return 0, ErrIntOverflowBatch
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
					return 0, ErrIntOverflowBatch
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
				return 0, ErrInvalidLengthBatch
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupBatch
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthBatch
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthBatch        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowBatch          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupBatch = fmt.Errorf("proto: unexpected end of group")
)
