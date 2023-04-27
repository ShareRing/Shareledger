// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: shareledger/distributionx/params.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
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

// Params defines the parameters for the module.
type Params struct {
	ConfigPercent  *Params_ConfigPercent `protobuf:"bytes,1,opt,name=config_percent,json=configPercent,proto3" json:"config_percent,omitempty" yaml:"config_percent"`
	BuilderWindows uint32                `protobuf:"varint,2,opt,name=builder_windows,json=builderWindows,proto3" json:"builder_windows,omitempty" yaml:"builder_windows"`
	TxThreshold    uint32                `protobuf:"varint,3,opt,name=tx_threshold,json=txThreshold,proto3" json:"tx_threshold,omitempty" yaml:"tx_threshold"`
	DevPoolAccount string                `protobuf:"bytes,4,opt,name=dev_pool_account,json=devPoolAccount,proto3" json:"dev_pool_account,omitempty" yaml:"dev_pool_account"`
}

func (m *Params) Reset()      { *m = Params{} }
func (*Params) ProtoMessage() {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_8f3bdd2d46050e6d, []int{0}
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

func (m *Params) GetConfigPercent() *Params_ConfigPercent {
	if m != nil {
		return m.ConfigPercent
	}
	return nil
}

func (m *Params) GetBuilderWindows() uint32 {
	if m != nil {
		return m.BuilderWindows
	}
	return 0
}

func (m *Params) GetTxThreshold() uint32 {
	if m != nil {
		return m.TxThreshold
	}
	return 0
}

func (m *Params) GetDevPoolAccount() string {
	if m != nil {
		return m.DevPoolAccount
	}
	return ""
}

type Params_ConfigPercent struct {
	WasmMasterBuilder github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,1,opt,name=wasm_master_builder,json=wasmMasterBuilder,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"wasm_master_builder"`
	WasmContractAdmin github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,2,opt,name=wasm_contract_admin,json=wasmContractAdmin,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"wasm_contract_admin"`
	WasmDevelopment   github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,3,opt,name=wasm_development,json=wasmDevelopment,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"wasm_development"`
	WasmValidator     github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,4,opt,name=wasm_validator,json=wasmValidator,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"wasm_validator"`
	NativeValidator   github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,5,opt,name=native_validator,json=nativeValidator,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"native_validator"`
	NativeDevelopment github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,6,opt,name=native_development,json=nativeDevelopment,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"native_development"`
}

func (m *Params_ConfigPercent) Reset()         { *m = Params_ConfigPercent{} }
func (m *Params_ConfigPercent) String() string { return proto.CompactTextString(m) }
func (*Params_ConfigPercent) ProtoMessage()    {}
func (*Params_ConfigPercent) Descriptor() ([]byte, []int) {
	return fileDescriptor_8f3bdd2d46050e6d, []int{0, 0}
}
func (m *Params_ConfigPercent) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params_ConfigPercent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params_ConfigPercent.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params_ConfigPercent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params_ConfigPercent.Merge(m, src)
}
func (m *Params_ConfigPercent) XXX_Size() int {
	return m.Size()
}
func (m *Params_ConfigPercent) XXX_DiscardUnknown() {
	xxx_messageInfo_Params_ConfigPercent.DiscardUnknown(m)
}

var xxx_messageInfo_Params_ConfigPercent proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Params)(nil), "sharering.shareledger.distributionx.Params")
	proto.RegisterType((*Params_ConfigPercent)(nil), "sharering.shareledger.distributionx.Params.ConfigPercent")
}

func init() {
	proto.RegisterFile("shareledger/distributionx/params.proto", fileDescriptor_8f3bdd2d46050e6d)
}

var fileDescriptor_8f3bdd2d46050e6d = []byte{
	// 537 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x94, 0xc1, 0x6e, 0xd3, 0x30,
	0x18, 0x80, 0x1b, 0xb6, 0x15, 0xcd, 0xa3, 0xdd, 0xc8, 0x80, 0x75, 0x45, 0x4a, 0xaa, 0x20, 0x4d,
	0xbd, 0x2c, 0x95, 0x40, 0x42, 0xa2, 0xe2, 0xb2, 0xb4, 0x1c, 0x91, 0xaa, 0x0a, 0x81, 0xc4, 0x25,
	0x72, 0x6d, 0x93, 0x5a, 0x4b, 0xe2, 0xc8, 0x76, 0xd3, 0x4e, 0x3c, 0x04, 0x1c, 0x39, 0xf2, 0x10,
	0x3c, 0xc4, 0x8e, 0x13, 0x27, 0xc4, 0x21, 0x42, 0xad, 0x78, 0x81, 0x3e, 0x01, 0x8a, 0x9d, 0x6d,
	0x69, 0x4f, 0x48, 0xeb, 0x29, 0xf1, 0x9f, 0xff, 0xff, 0xbe, 0xf8, 0xff, 0x13, 0x83, 0x13, 0x31,
	0x86, 0x9c, 0x84, 0x04, 0x07, 0x84, 0x77, 0x30, 0x15, 0x92, 0xd3, 0xd1, 0x44, 0x52, 0x16, 0xcf,
	0x3a, 0x09, 0xe4, 0x30, 0x12, 0x6e, 0xc2, 0x99, 0x64, 0xe6, 0x33, 0x95, 0xc7, 0x69, 0x1c, 0xb8,
	0xa5, 0x0a, 0x77, 0xa5, 0xa2, 0x79, 0x8c, 0x98, 0x88, 0x98, 0xf0, 0x55, 0x49, 0x47, 0x2f, 0x74,
	0x7d, 0xf3, 0x51, 0xc0, 0x02, 0xa6, 0xe3, 0xf9, 0x9d, 0x8e, 0x3a, 0x7f, 0xef, 0x83, 0xea, 0x40,
	0x69, 0xcc, 0xcf, 0xa0, 0x8e, 0x58, 0xfc, 0x89, 0x06, 0x7e, 0x42, 0x38, 0x22, 0xb1, 0x6c, 0x18,
	0x2d, 0xa3, 0xbd, 0xf7, 0xfc, 0x95, 0xfb, 0x1f, 0x66, 0x57, 0x43, 0xdc, 0x9e, 0x22, 0x0c, 0x34,
	0xc0, 0x3b, 0x5e, 0x66, 0xf6, 0xe3, 0x0b, 0x18, 0x85, 0x5d, 0x67, 0x15, 0xed, 0x0c, 0x6b, 0xa8,
	0x9c, 0x69, 0xf6, 0xc0, 0xfe, 0x68, 0x42, 0x43, 0x4c, 0xb8, 0x3f, 0xa5, 0x31, 0x66, 0x53, 0xd1,
	0xb8, 0xd7, 0x32, 0xda, 0x35, 0xaf, 0xb9, 0xcc, 0xec, 0x27, 0x1a, 0xb1, 0x96, 0xe0, 0x0c, 0xeb,
	0x45, 0xe4, 0x83, 0x0e, 0x98, 0x5d, 0xf0, 0x40, 0xce, 0x7c, 0x39, 0xe6, 0x44, 0x8c, 0x59, 0x88,
	0x1b, 0x5b, 0x8a, 0x70, 0xb4, 0xcc, 0xec, 0x43, 0x4d, 0x28, 0x3f, 0x75, 0x86, 0x7b, 0x72, 0xf6,
	0xee, 0x7a, 0x65, 0xbe, 0x01, 0x07, 0x98, 0xa4, 0x7e, 0xc2, 0x58, 0xe8, 0x43, 0x84, 0xd8, 0x24,
	0x96, 0x8d, 0xed, 0x96, 0xd1, 0xde, 0xf5, 0x9e, 0x2e, 0x33, 0xfb, 0x48, 0xd7, 0xaf, 0x67, 0x38,
	0xc3, 0x3a, 0x26, 0xe9, 0x80, 0xb1, 0xf0, 0x4c, 0x07, 0x9a, 0x5f, 0x76, 0x40, 0x6d, 0xa5, 0x07,
	0x66, 0x08, 0x0e, 0xa7, 0x50, 0x44, 0x7e, 0x04, 0x85, 0x24, 0xdc, 0x2f, 0x5e, 0x59, 0xf5, 0x76,
	0xd7, 0x7b, 0x7d, 0x99, 0xd9, 0x95, 0xdf, 0x99, 0x7d, 0x12, 0x50, 0x39, 0x9e, 0x8c, 0x5c, 0xc4,
	0xa2, 0x62, 0x6a, 0xc5, 0xe5, 0x54, 0xe0, 0xf3, 0x8e, 0xbc, 0x48, 0x88, 0x70, 0xfb, 0x04, 0xfd,
	0xfc, 0x71, 0x0a, 0x8a, 0xa1, 0xf6, 0x09, 0x1a, 0x3e, 0xcc, 0xc1, 0x6f, 0x15, 0xd7, 0xd3, 0xd8,
	0x1b, 0x1b, 0x62, 0xb1, 0xe4, 0x10, 0x49, 0x1f, 0xe2, 0x88, 0xc6, 0xaa, 0x97, 0x1b, 0xb1, 0xf5,
	0x0a, 0xee, 0x59, 0x8e, 0x35, 0x03, 0x70, 0xa0, 0x6c, 0x98, 0xa4, 0x24, 0x64, 0x49, 0x94, 0x7f,
	0x34, 0x5b, 0x1b, 0x50, 0xed, 0xe7, 0xd4, 0xfe, 0x2d, 0xd4, 0x44, 0xa0, 0xae, 0x44, 0x29, 0x0c,
	0x29, 0x86, 0x92, 0xf1, 0x62, 0x36, 0x77, 0xd3, 0xd4, 0x72, 0xe6, 0xfb, 0x6b, 0x64, 0xbe, 0x9b,
	0x18, 0x4a, 0x9a, 0x92, 0x92, 0x66, 0x67, 0x13, 0xbb, 0xd1, 0xd4, 0x5b, 0xd1, 0x39, 0x30, 0x0b,
	0x51, 0xb9, 0x71, 0xd5, 0x4d, 0xcc, 0x48, 0x73, 0x4b, 0xad, 0xeb, 0x6e, 0x7f, 0xfb, 0x6e, 0x57,
	0xbc, 0xc1, 0xe5, 0xdc, 0x32, 0xae, 0xe6, 0x96, 0xf1, 0x67, 0x6e, 0x19, 0x5f, 0x17, 0x56, 0xe5,
	0x6a, 0x61, 0x55, 0x7e, 0x2d, 0xac, 0xca, 0xc7, 0x97, 0x25, 0xd1, 0xcd, 0x8f, 0xde, 0x29, 0x1f,
	0x4a, 0xb3, 0xb5, 0x63, 0x49, 0xc9, 0x47, 0x55, 0x75, 0x80, 0xbc, 0xf8, 0x17, 0x00, 0x00, 0xff,
	0xff, 0x4e, 0xe0, 0xd6, 0x6c, 0xc0, 0x04, 0x00, 0x00,
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
	if len(m.DevPoolAccount) > 0 {
		i -= len(m.DevPoolAccount)
		copy(dAtA[i:], m.DevPoolAccount)
		i = encodeVarintParams(dAtA, i, uint64(len(m.DevPoolAccount)))
		i--
		dAtA[i] = 0x22
	}
	if m.TxThreshold != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.TxThreshold))
		i--
		dAtA[i] = 0x18
	}
	if m.BuilderWindows != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.BuilderWindows))
		i--
		dAtA[i] = 0x10
	}
	if m.ConfigPercent != nil {
		{
			size, err := m.ConfigPercent.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintParams(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Params_ConfigPercent) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params_ConfigPercent) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params_ConfigPercent) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.NativeDevelopment.Size()
		i -= size
		if _, err := m.NativeDevelopment.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x32
	{
		size := m.NativeValidator.Size()
		i -= size
		if _, err := m.NativeValidator.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	{
		size := m.WasmValidator.Size()
		i -= size
		if _, err := m.WasmValidator.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	{
		size := m.WasmDevelopment.Size()
		i -= size
		if _, err := m.WasmDevelopment.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size := m.WasmContractAdmin.Size()
		i -= size
		if _, err := m.WasmContractAdmin.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size := m.WasmMasterBuilder.Size()
		i -= size
		if _, err := m.WasmMasterBuilder.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintParams(dAtA []byte, offset int, v uint64) int {
	offset -= sovParams(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ConfigPercent != nil {
		l = m.ConfigPercent.Size()
		n += 1 + l + sovParams(uint64(l))
	}
	if m.BuilderWindows != 0 {
		n += 1 + sovParams(uint64(m.BuilderWindows))
	}
	if m.TxThreshold != 0 {
		n += 1 + sovParams(uint64(m.TxThreshold))
	}
	l = len(m.DevPoolAccount)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	return n
}

func (m *Params_ConfigPercent) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.WasmMasterBuilder.Size()
	n += 1 + l + sovParams(uint64(l))
	l = m.WasmContractAdmin.Size()
	n += 1 + l + sovParams(uint64(l))
	l = m.WasmDevelopment.Size()
	n += 1 + l + sovParams(uint64(l))
	l = m.WasmValidator.Size()
	n += 1 + l + sovParams(uint64(l))
	l = m.NativeValidator.Size()
	n += 1 + l + sovParams(uint64(l))
	l = m.NativeDevelopment.Size()
	n += 1 + l + sovParams(uint64(l))
	return n
}

func sovParams(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozParams(x uint64) (n int) {
	return sovParams(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowParams
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
				return fmt.Errorf("proto: wrong wireType = %d for field ConfigPercent", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ConfigPercent == nil {
				m.ConfigPercent = &Params_ConfigPercent{}
			}
			if err := m.ConfigPercent.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BuilderWindows", wireType)
			}
			m.BuilderWindows = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BuilderWindows |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TxThreshold", wireType)
			}
			m.TxThreshold = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TxThreshold |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DevPoolAccount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DevPoolAccount = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipParams(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthParams
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
func (m *Params_ConfigPercent) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowParams
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
			return fmt.Errorf("proto: ConfigPercent: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ConfigPercent: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field WasmMasterBuilder", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.WasmMasterBuilder.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field WasmContractAdmin", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.WasmContractAdmin.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field WasmDevelopment", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.WasmDevelopment.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field WasmValidator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.WasmValidator.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NativeValidator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.NativeValidator.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NativeDevelopment", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.NativeDevelopment.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipParams(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthParams
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
func skipParams(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
				return 0, ErrInvalidLengthParams
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupParams
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthParams
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthParams        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowParams          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupParams = fmt.Errorf("proto: unexpected end of group")
)
