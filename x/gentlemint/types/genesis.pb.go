// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: gentlemint/genesis.proto

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

// GenesisState defines the gentlemint module's genesis state.
type GenesisState struct {
	AccStateList []AccState    `protobuf:"bytes,1,rep,name=accStateList,proto3" json:"accStateList"`
	Authority    *Authority    `protobuf:"bytes,2,opt,name=authority,proto3" json:"authority,omitempty"`
	Treasurer    *Treasurer    `protobuf:"bytes,3,opt,name=treasurer,proto3" json:"treasurer,omitempty"`
	ExchangeRate *ExchangeRate `protobuf:"bytes,4,opt,name=exchangeRate,proto3" json:"exchangeRate,omitempty"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_49e389fa8edb1952, []int{0}
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

func (m *GenesisState) GetAccStateList() []AccState {
	if m != nil {
		return m.AccStateList
	}
	return nil
}

func (m *GenesisState) GetAuthority() *Authority {
	if m != nil {
		return m.Authority
	}
	return nil
}

func (m *GenesisState) GetTreasurer() *Treasurer {
	if m != nil {
		return m.Treasurer
	}
	return nil
}

func (m *GenesisState) GetExchangeRate() *ExchangeRate {
	if m != nil {
		return m.ExchangeRate
	}
	return nil
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "ShareRing.shareledger.gentlemint.GenesisState")
}

func init() { proto.RegisterFile("gentlemint/genesis.proto", fileDescriptor_49e389fa8edb1952) }

var fileDescriptor_49e389fa8edb1952 = []byte{
	// 306 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0x4f, 0x6b, 0xc2, 0x30,
	0x18, 0xc6, 0x5b, 0x95, 0xc1, 0x6a, 0x4f, 0x65, 0x87, 0xe2, 0x21, 0x2b, 0x3b, 0xc9, 0x06, 0x29,
	0xe8, 0x27, 0x98, 0x30, 0xc6, 0x60, 0xbb, 0x44, 0x4f, 0xbb, 0x48, 0xcc, 0x5e, 0xd2, 0x80, 0x36,
	0x92, 0xa4, 0xa0, 0xdf, 0x62, 0x1f, 0xcb, 0xd3, 0xf0, 0xb8, 0xd3, 0x18, 0xed, 0x17, 0x19, 0x6d,
	0xed, 0x3f, 0x18, 0xe8, 0xed, 0xa1, 0x4f, 0x7f, 0xbf, 0xe4, 0xcd, 0xeb, 0xf8, 0x1c, 0x62, 0xb3,
	0x86, 0x8d, 0x88, 0x4d, 0xc8, 0x21, 0x06, 0x2d, 0x34, 0xde, 0x2a, 0x69, 0xa4, 0x17, 0xcc, 0x23,
	0xaa, 0x80, 0x88, 0x98, 0x63, 0x9d, 0xa7, 0x35, 0x7c, 0x70, 0x50, 0xb8, 0xf9, 0x7f, 0x34, 0x6a,
	0xb1, 0x94, 0xb1, 0xa5, 0x36, 0xd4, 0x40, 0x49, 0x77, 0xbb, 0xc4, 0x44, 0x52, 0x09, 0xb3, 0xff,
	0xa7, 0x33, 0x0a, 0xa8, 0x4e, 0x14, 0xa8, 0x53, 0x87, 0x5a, 0x1d, 0xec, 0x58, 0x44, 0x63, 0x0e,
	0x4b, 0xd5, 0x78, 0x6f, 0xb8, 0xe4, 0xb2, 0x88, 0x61, 0x9e, 0xca, 0xaf, 0x77, 0x5f, 0x3d, 0xc7,
	0x7d, 0x2e, 0x6f, 0x3f, 0xcf, 0x2f, 0xe1, 0x2d, 0x1c, 0x97, 0x32, 0x56, 0xe4, 0x57, 0xa1, 0x8d,
	0x6f, 0x07, 0xfd, 0xf1, 0x70, 0x72, 0x8f, 0xcf, 0xcd, 0x84, 0x1f, 0x4f, 0xd4, 0x6c, 0x70, 0xf8,
	0xb9, 0xb5, 0x48, 0xc7, 0xe2, 0xbd, 0x38, 0xd7, 0xf5, 0x2c, 0x7e, 0x2f, 0xb0, 0xc7, 0xc3, 0xc9,
	0xc3, 0x05, 0xca, 0x0a, 0x21, 0x0d, 0x9d, 0xab, 0xea, 0xd1, 0xfd, 0xfe, 0xa5, 0xaa, 0x45, 0x85,
	0x90, 0x86, 0xf6, 0x88, 0xe3, 0x56, 0x2f, 0x45, 0xa8, 0x01, 0x7f, 0x50, 0xd8, 0xf0, 0x79, 0xdb,
	0x53, 0x8b, 0x22, 0x1d, 0xc7, 0xec, 0xed, 0x90, 0x22, 0xfb, 0x98, 0x22, 0xfb, 0x37, 0x45, 0xf6,
	0x67, 0x86, 0xac, 0x63, 0x86, 0xac, 0xef, 0x0c, 0x59, 0xef, 0x53, 0x2e, 0x4c, 0x94, 0xac, 0x30,
	0x93, 0x9b, 0xb0, 0x3e, 0xa1, 0x4c, 0xe5, 0x09, 0xe1, 0x2e, 0x6c, 0xef, 0x77, 0xbf, 0x05, 0xbd,
	0xba, 0x2a, 0xd6, 0x34, 0xfd, 0x0b, 0x00, 0x00, 0xff, 0xff, 0x3b, 0x01, 0x2e, 0xd7, 0x6e, 0x02,
	0x00, 0x00,
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
	if m.ExchangeRate != nil {
		{
			size, err := m.ExchangeRate.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGenesis(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	if m.Treasurer != nil {
		{
			size, err := m.Treasurer.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGenesis(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if m.Authority != nil {
		{
			size, err := m.Authority.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGenesis(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.AccStateList) > 0 {
		for iNdEx := len(m.AccStateList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.AccStateList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.AccStateList) > 0 {
		for _, e := range m.AccStateList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if m.Authority != nil {
		l = m.Authority.Size()
		n += 1 + l + sovGenesis(uint64(l))
	}
	if m.Treasurer != nil {
		l = m.Treasurer.Size()
		n += 1 + l + sovGenesis(uint64(l))
	}
	if m.ExchangeRate != nil {
		l = m.ExchangeRate.Size()
		n += 1 + l + sovGenesis(uint64(l))
	}
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AccStateList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AccStateList = append(m.AccStateList, AccState{})
			if err := m.AccStateList[len(m.AccStateList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Authority", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Authority == nil {
				m.Authority = &Authority{}
			}
			if err := m.Authority.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Treasurer", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Treasurer == nil {
				m.Treasurer = &Treasurer{}
			}
			if err := m.Treasurer.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExchangeRate", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ExchangeRate == nil {
				m.ExchangeRate = &ExchangeRate{}
			}
			if err := m.ExchangeRate.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)