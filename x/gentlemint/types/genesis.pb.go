// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: shareledger/gentlemint/genesis.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
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
	ExchangeRate       *ExchangeRate    `protobuf:"bytes,1,opt,name=exchangeRate,proto3" json:"exchangeRate,omitempty"`
	LevelFeeList       []LevelFee       `protobuf:"bytes,2,rep,name=levelFeeList,proto3" json:"levelFeeList"`
	ActionLevelFeeList []ActionLevelFee `protobuf:"bytes,3,rep,name=actionLevelFeeList,proto3" json:"actionLevelFeeList"`
	Params             Params           `protobuf:"bytes,4,opt,name=params,proto3" json:"params,omitempty"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_94bc731420dd597d, []int{0}
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

func (m *GenesisState) GetExchangeRate() *ExchangeRate {
	if m != nil {
		return m.ExchangeRate
	}
	return nil
}

func (m *GenesisState) GetLevelFeeList() []LevelFee {
	if m != nil {
		return m.LevelFeeList
	}
	return nil
}

func (m *GenesisState) GetActionLevelFeeList() []ActionLevelFee {
	if m != nil {
		return m.ActionLevelFeeList
	}
	return nil
}

func (m *GenesisState) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

// Params defines the set of module parameters.
type Params struct {
	// Minimum stores the minimum gas price(s) for all TX on the chain.
	// When multiple coins are defined then they are accepted alternatively.
	// The list must be sorted by denoms asc. No duplicate denoms or zero amount
	// values allowed. For more information see
	// https://docs.cosmos.network/main/modules/auth#concepts
	MinimumGasPrices github_com_cosmos_cosmos_sdk_types.DecCoins `protobuf:"bytes,1,rep,name=minimum_gas_prices,json=minimumGasPrices,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.DecCoins" json:"minimum_gas_prices,omitempty" yaml:"minimum_gas_prices"`
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_94bc731420dd597d, []int{1}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

func (m *Params) GetMinimumGasPrices() github_com_cosmos_cosmos_sdk_types.DecCoins {
	if m != nil {
		return m.MinimumGasPrices
	}
	return nil
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "shareledger.gentlemint.GenesisState")
	proto.RegisterType((*Params)(nil), "shareledger.gentlemint.Params")
}

func init() {
	proto.RegisterFile("shareledger/gentlemint/genesis.proto", fileDescriptor_94bc731420dd597d)
}

var fileDescriptor_94bc731420dd597d = []byte{
	// 451 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0x3f, 0x6b, 0x1b, 0x31,
	0x1c, 0x86, 0x7d, 0x49, 0xf0, 0xa0, 0x78, 0x08, 0xa2, 0x94, 0x6b, 0x08, 0xb2, 0x31, 0x21, 0x84,
	0xb6, 0x91, 0x48, 0xb2, 0x75, 0xab, 0xfb, 0x27, 0xa5, 0xb8, 0x60, 0xae, 0x5b, 0x29, 0x1c, 0xf2,
	0xe5, 0x57, 0x59, 0xf4, 0x24, 0x1d, 0x27, 0x25, 0xc4, 0xdf, 0xa2, 0x9f, 0xa3, 0x9f, 0xa1, 0x7b,
	0x33, 0x66, 0xe8, 0xd0, 0xc9, 0x2d, 0xf6, 0xd6, 0xb1, 0x9f, 0xa0, 0x58, 0x27, 0x93, 0x73, 0x93,
	0x9b, 0x2c, 0xcc, 0xf3, 0x3e, 0xef, 0x49, 0xbc, 0x68, 0xdf, 0x4e, 0x78, 0x09, 0x39, 0x9c, 0x0b,
	0x28, 0x99, 0x00, 0xed, 0x72, 0x50, 0x52, 0xbb, 0xe5, 0x11, 0xac, 0xb4, 0xb4, 0x28, 0x8d, 0x33,
	0xf8, 0x61, 0x8d, 0xa2, 0xb7, 0xd4, 0x2e, 0xc9, 0x8c, 0x55, 0xc6, 0xb2, 0x31, 0xb7, 0xc0, 0x2e,
	0x8f, 0xc7, 0xe0, 0xf8, 0x31, 0xcb, 0x8c, 0xd4, 0x55, 0x6e, 0xf7, 0x81, 0x30, 0xc2, 0xf8, 0x23,
	0x5b, 0x9e, 0xc2, 0xbf, 0x47, 0x0d, 0x9d, 0x3c, 0x73, 0xd2, 0xe8, 0x34, 0x87, 0x4b, 0xc8, 0xd3,
	0x4f, 0x00, 0x01, 0x7f, 0xdc, 0x80, 0xc3, 0x55, 0x36, 0xe1, 0x5a, 0x40, 0x5a, 0x72, 0xb7, 0x62,
	0x0f, 0x1a, 0xd8, 0xff, 0x9c, 0xfd, 0x1f, 0x1b, 0xa8, 0x73, 0x56, 0x5d, 0xf1, 0xbd, 0xe3, 0x0e,
	0xf0, 0x1b, 0xd4, 0x59, 0xf9, 0x12, 0xee, 0x20, 0x8e, 0x7a, 0xd1, 0xe1, 0xf6, 0xc9, 0x3e, 0xbd,
	0xff, 0xe2, 0xf4, 0x55, 0x8d, 0x4d, 0xd6, 0x92, 0xf8, 0x2d, 0xea, 0xf8, 0xb6, 0xd7, 0x00, 0x43,
	0x69, 0x5d, 0xbc, 0xd1, 0xdb, 0x3c, 0xdc, 0x3e, 0xe9, 0x35, 0x99, 0x86, 0x81, 0x1d, 0x6c, 0x5d,
	0xcf, 0xba, 0xad, 0x64, 0x2d, 0x8b, 0x3f, 0x22, 0x5c, 0x3d, 0xca, 0xb0, 0x6e, 0xdc, 0xf4, 0xc6,
	0x83, 0x26, 0xe3, 0xf3, 0xb5, 0x44, 0xf0, 0xde, 0xe3, 0xc1, 0x23, 0xd4, 0x2e, 0x78, 0xc9, 0x95,
	0x8d, 0xb7, 0xfc, 0x6d, 0x49, 0x93, 0x71, 0xe4, 0xa9, 0x41, 0xbc, 0x34, 0xfd, 0x99, 0x75, 0x77,
	0xaa, 0xd4, 0x53, 0xa3, 0xa4, 0x03, 0x55, 0xb8, 0x69, 0x12, 0x3c, 0xfd, 0xef, 0x11, 0x6a, 0x57,
	0x30, 0xfe, 0x16, 0x21, 0xac, 0xa4, 0x96, 0xea, 0x42, 0xa5, 0x82, 0xdb, 0xb4, 0x28, 0x65, 0x06,
	0x36, 0x8e, 0xfc, 0xb7, 0xef, 0xd1, 0x6a, 0x38, 0x74, 0x39, 0x1c, 0x1a, 0x86, 0x43, 0x5f, 0x42,
	0xf6, 0xc2, 0x48, 0x3d, 0x28, 0x42, 0xcf, 0xde, 0xdd, 0xfc, 0x6d, 0xe7, 0xdf, 0x59, 0xf7, 0xd1,
	0x94, 0xab, 0xfc, 0x59, 0xff, 0x2e, 0xd5, 0xff, 0xfa, 0xab, 0xfb, 0x44, 0x48, 0x37, 0xb9, 0x18,
	0xd3, 0xcc, 0x28, 0x16, 0x56, 0x5a, 0xfd, 0x1c, 0xd9, 0xf3, 0xcf, 0xcc, 0x4d, 0x0b, 0xb0, 0xab,
	0x42, 0x9b, 0xec, 0x04, 0xc7, 0x19, 0xb7, 0x23, 0x6f, 0x18, 0xbc, 0xbb, 0x9e, 0x93, 0xe8, 0x66,
	0x4e, 0xa2, 0xdf, 0x73, 0x12, 0x7d, 0x59, 0x90, 0xd6, 0xcd, 0x82, 0xb4, 0x7e, 0x2e, 0x48, 0xeb,
	0xc3, 0x69, 0x4d, 0xec, 0xdf, 0xab, 0x94, 0x5a, 0xb0, 0xfa, 0xee, 0xae, 0xea, 0xcb, 0xf3, 0x4d,
	0xe3, 0xb6, 0x9f, 0xdd, 0xe9, 0xbf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xe3, 0xd1, 0xee, 0xbc, 0x6f,
	0x03, 0x00, 0x00,
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
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	if len(m.ActionLevelFeeList) > 0 {
		for iNdEx := len(m.ActionLevelFeeList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ActionLevelFeeList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.LevelFeeList) > 0 {
		for iNdEx := len(m.LevelFeeList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.LevelFeeList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
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
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.MinimumGasPrices) > 0 {
		for iNdEx := len(m.MinimumGasPrices) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.MinimumGasPrices[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
	if m.ExchangeRate != nil {
		l = m.ExchangeRate.Size()
		n += 1 + l + sovGenesis(uint64(l))
	}
	if len(m.LevelFeeList) > 0 {
		for _, e := range m.LevelFeeList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.ActionLevelFeeList) > 0 {
		for _, e := range m.ActionLevelFeeList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	return n
}

func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.MinimumGasPrices) > 0 {
		for _, e := range m.MinimumGasPrices {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
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
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LevelFeeList", wireType)
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
			m.LevelFeeList = append(m.LevelFeeList, LevelFee{})
			if err := m.LevelFeeList[len(m.LevelFeeList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ActionLevelFeeList", wireType)
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
			m.ActionLevelFeeList = append(m.ActionLevelFeeList, ActionLevelFee{})
			if err := m.ActionLevelFeeList[len(m.ActionLevelFeeList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
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
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *Params) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinimumGasPrices", wireType)
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
			m.MinimumGasPrices = append(m.MinimumGasPrices, types.DecCoin{})
			if err := m.MinimumGasPrices[len(m.MinimumGasPrices)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
