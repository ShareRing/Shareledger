// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: shareledger/id/v1/base_id.proto

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

type BaseID struct {
	IssuerAddress string `protobuf:"bytes,1,opt,name=issuerAddress,proto3" json:"issuerAddress,omitempty"`
	BackupAddress string `protobuf:"bytes,2,opt,name=backupAddress,proto3" json:"backupAddress,omitempty"`
	OwnerAddress  string `protobuf:"bytes,3,opt,name=ownerAddress,proto3" json:"ownerAddress,omitempty"`
	ExtraData     string `protobuf:"bytes,4,opt,name=extraData,proto3" json:"extraData,omitempty"`
}

func (m *BaseID) Reset()         { *m = BaseID{} }
func (m *BaseID) String() string { return proto.CompactTextString(m) }
func (*BaseID) ProtoMessage()    {}
func (*BaseID) Descriptor() ([]byte, []int) {
	return fileDescriptor_d28abe44a1b31d68, []int{0}
}
func (m *BaseID) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BaseID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BaseID.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BaseID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BaseID.Merge(m, src)
}
func (m *BaseID) XXX_Size() int {
	return m.Size()
}
func (m *BaseID) XXX_DiscardUnknown() {
	xxx_messageInfo_BaseID.DiscardUnknown(m)
}

var xxx_messageInfo_BaseID proto.InternalMessageInfo

func (m *BaseID) GetIssuerAddress() string {
	if m != nil {
		return m.IssuerAddress
	}
	return ""
}

func (m *BaseID) GetBackupAddress() string {
	if m != nil {
		return m.BackupAddress
	}
	return ""
}

func (m *BaseID) GetOwnerAddress() string {
	if m != nil {
		return m.OwnerAddress
	}
	return ""
}

func (m *BaseID) GetExtraData() string {
	if m != nil {
		return m.ExtraData
	}
	return ""
}

func init() {
	proto.RegisterType((*BaseID)(nil), "shareledger.id.BaseID")
}

func init() { proto.RegisterFile("shareledger/id/v1/base_id.proto", fileDescriptor_d28abe44a1b31d68) }

var fileDescriptor_d28abe44a1b31d68 = []byte{
	// 213 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2f, 0xce, 0x48, 0x2c,
	0x4a, 0xcd, 0x49, 0x4d, 0x49, 0x4f, 0x2d, 0xd2, 0xcf, 0x4c, 0xd1, 0x2f, 0x33, 0xd4, 0x4f, 0x4a,
	0x2c, 0x4e, 0x8d, 0xcf, 0x4c, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x43, 0x52, 0xa0,
	0x97, 0x99, 0xa2, 0x34, 0x8d, 0x91, 0x8b, 0xcd, 0x29, 0xb1, 0x38, 0xd5, 0xd3, 0x45, 0x48, 0x85,
	0x8b, 0x37, 0xb3, 0xb8, 0xb8, 0x34, 0xb5, 0xc8, 0x31, 0x25, 0xa5, 0x28, 0xb5, 0xb8, 0x58, 0x82,
	0x51, 0x81, 0x51, 0x83, 0x33, 0x08, 0x55, 0x10, 0xa4, 0x2a, 0x29, 0x31, 0x39, 0xbb, 0xb4, 0x00,
	0xa6, 0x8a, 0x09, 0xa2, 0x0a, 0x45, 0x50, 0x48, 0x89, 0x8b, 0x27, 0xbf, 0x3c, 0x0f, 0x61, 0x14,
	0x33, 0x58, 0x11, 0x8a, 0x98, 0x90, 0x0c, 0x17, 0x67, 0x6a, 0x45, 0x49, 0x51, 0xa2, 0x4b, 0x62,
	0x49, 0xa2, 0x04, 0x0b, 0x58, 0x01, 0x42, 0xc0, 0xc9, 0xf5, 0xc4, 0x23, 0x39, 0xc6, 0x0b, 0x8f,
	0xe4, 0x18, 0x1f, 0x3c, 0x92, 0x63, 0x9c, 0xf0, 0x58, 0x8e, 0xe1, 0xc2, 0x63, 0x39, 0x86, 0x1b,
	0x8f, 0xe5, 0x18, 0xa2, 0xb4, 0xd3, 0x33, 0x4b, 0x32, 0x4a, 0x93, 0xf4, 0x92, 0xf3, 0x73, 0xf5,
	0xc1, 0xbe, 0x29, 0xca, 0xcc, 0x4b, 0xd7, 0x47, 0xf6, 0x78, 0x05, 0xc8, 0xeb, 0x25, 0x95, 0x05,
	0xa9, 0xc5, 0x49, 0x6c, 0x60, 0x6f, 0x1b, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x9f, 0x20, 0x4a,
	0x9a, 0x19, 0x01, 0x00, 0x00,
}

func (m *BaseID) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BaseID) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BaseID) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ExtraData) > 0 {
		i -= len(m.ExtraData)
		copy(dAtA[i:], m.ExtraData)
		i = encodeVarintBaseId(dAtA, i, uint64(len(m.ExtraData)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.OwnerAddress) > 0 {
		i -= len(m.OwnerAddress)
		copy(dAtA[i:], m.OwnerAddress)
		i = encodeVarintBaseId(dAtA, i, uint64(len(m.OwnerAddress)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.BackupAddress) > 0 {
		i -= len(m.BackupAddress)
		copy(dAtA[i:], m.BackupAddress)
		i = encodeVarintBaseId(dAtA, i, uint64(len(m.BackupAddress)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.IssuerAddress) > 0 {
		i -= len(m.IssuerAddress)
		copy(dAtA[i:], m.IssuerAddress)
		i = encodeVarintBaseId(dAtA, i, uint64(len(m.IssuerAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintBaseId(dAtA []byte, offset int, v uint64) int {
	offset -= sovBaseId(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *BaseID) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.IssuerAddress)
	if l > 0 {
		n += 1 + l + sovBaseId(uint64(l))
	}
	l = len(m.BackupAddress)
	if l > 0 {
		n += 1 + l + sovBaseId(uint64(l))
	}
	l = len(m.OwnerAddress)
	if l > 0 {
		n += 1 + l + sovBaseId(uint64(l))
	}
	l = len(m.ExtraData)
	if l > 0 {
		n += 1 + l + sovBaseId(uint64(l))
	}
	return n
}

func sovBaseId(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozBaseId(x uint64) (n int) {
	return sovBaseId(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *BaseID) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBaseId
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
			return fmt.Errorf("proto: BaseID: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BaseID: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IssuerAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBaseId
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
				return ErrInvalidLengthBaseId
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBaseId
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.IssuerAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BackupAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBaseId
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
				return ErrInvalidLengthBaseId
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBaseId
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BackupAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OwnerAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBaseId
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
				return ErrInvalidLengthBaseId
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBaseId
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OwnerAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExtraData", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBaseId
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
				return ErrInvalidLengthBaseId
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBaseId
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ExtraData = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipBaseId(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthBaseId
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
func skipBaseId(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowBaseId
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
					return 0, ErrIntOverflowBaseId
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
					return 0, ErrIntOverflowBaseId
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
				return 0, ErrInvalidLengthBaseId
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupBaseId
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthBaseId
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthBaseId        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowBaseId          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupBaseId = fmt.Errorf("proto: unexpected end of group")
)
