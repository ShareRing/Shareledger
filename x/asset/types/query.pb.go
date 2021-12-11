// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: asset/query.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/types/query"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type QueryAssetByUUIDRequest struct {
	Uuid string `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
}

func (m *QueryAssetByUUIDRequest) Reset()         { *m = QueryAssetByUUIDRequest{} }
func (m *QueryAssetByUUIDRequest) String() string { return proto.CompactTextString(m) }
func (*QueryAssetByUUIDRequest) ProtoMessage()    {}
func (*QueryAssetByUUIDRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4785ee6dd031b7c8, []int{0}
}
func (m *QueryAssetByUUIDRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryAssetByUUIDRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryAssetByUUIDRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryAssetByUUIDRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryAssetByUUIDRequest.Merge(m, src)
}
func (m *QueryAssetByUUIDRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryAssetByUUIDRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryAssetByUUIDRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryAssetByUUIDRequest proto.InternalMessageInfo

type QueryAssetByUUIDResponse struct {
	Asset *Asset `protobuf:"bytes,1,opt,name=asset,proto3" json:"asset,omitempty"`
}

func (m *QueryAssetByUUIDResponse) Reset()         { *m = QueryAssetByUUIDResponse{} }
func (m *QueryAssetByUUIDResponse) String() string { return proto.CompactTextString(m) }
func (*QueryAssetByUUIDResponse) ProtoMessage()    {}
func (*QueryAssetByUUIDResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_4785ee6dd031b7c8, []int{1}
}
func (m *QueryAssetByUUIDResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryAssetByUUIDResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryAssetByUUIDResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryAssetByUUIDResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryAssetByUUIDResponse.Merge(m, src)
}
func (m *QueryAssetByUUIDResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryAssetByUUIDResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryAssetByUUIDResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryAssetByUUIDResponse proto.InternalMessageInfo

func (m *QueryAssetByUUIDResponse) GetAsset() *Asset {
	if m != nil {
		return m.Asset
	}
	return nil
}

func init() {
	proto.RegisterType((*QueryAssetByUUIDRequest)(nil), "ShareRing.shareledger.asset.QueryAssetByUUIDRequest")
	proto.RegisterType((*QueryAssetByUUIDResponse)(nil), "ShareRing.shareledger.asset.QueryAssetByUUIDResponse")
}

func init() { proto.RegisterFile("asset/query.proto", fileDescriptor_4785ee6dd031b7c8) }

var fileDescriptor_4785ee6dd031b7c8 = []byte{
	// 336 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4c, 0x2c, 0x2e, 0x4e,
	0x2d, 0xd1, 0x2f, 0x2c, 0x4d, 0x2d, 0xaa, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x92, 0x0e,
	0xce, 0x48, 0x2c, 0x4a, 0x0d, 0xca, 0xcc, 0x4b, 0xd7, 0x2b, 0x06, 0xb1, 0x72, 0x52, 0x53, 0xd2,
	0x53, 0x8b, 0xf4, 0xc0, 0x0a, 0xa5, 0x44, 0xd2, 0xf3, 0xd3, 0xf3, 0xc1, 0xea, 0xf4, 0x41, 0x2c,
	0x88, 0x16, 0x29, 0x99, 0xf4, 0xfc, 0xfc, 0xf4, 0x9c, 0x54, 0xfd, 0xc4, 0x82, 0x4c, 0xfd, 0xc4,
	0xbc, 0xbc, 0xfc, 0x92, 0xc4, 0x92, 0xcc, 0xfc, 0xbc, 0x62, 0xa8, 0xac, 0x56, 0x72, 0x7e, 0x71,
	0x6e, 0x7e, 0xb1, 0x7e, 0x52, 0x62, 0x71, 0x2a, 0xc4, 0x26, 0xfd, 0x32, 0xc3, 0xa4, 0xd4, 0x92,
	0x44, 0x43, 0xfd, 0x82, 0xc4, 0xf4, 0xcc, 0x3c, 0xb0, 0x62, 0xa8, 0x5a, 0xa8, 0x7b, 0xc0, 0x24,
	0x44, 0x48, 0xc9, 0x9c, 0x4b, 0x3c, 0x10, 0xa4, 0xc9, 0x11, 0x24, 0xe6, 0x54, 0x19, 0x1a, 0xea,
	0xe9, 0x12, 0x94, 0x5a, 0x58, 0x9a, 0x5a, 0x5c, 0x22, 0x24, 0xc4, 0xc5, 0x52, 0x5a, 0x9a, 0x99,
	0x22, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0x19, 0x04, 0x66, 0x5b, 0x71, 0x74, 0x2c, 0x90, 0x67, 0x78,
	0xb1, 0x40, 0x9e, 0x41, 0x29, 0x84, 0x4b, 0x02, 0x53, 0x63, 0x71, 0x41, 0x7e, 0x5e, 0x71, 0xaa,
	0x90, 0x05, 0x17, 0x2b, 0xd8, 0x0e, 0xb0, 0x56, 0x6e, 0x23, 0x25, 0x3d, 0x3c, 0x9e, 0xd6, 0x03,
	0x1b, 0x10, 0x04, 0xd1, 0x60, 0xb4, 0x9c, 0x91, 0x8b, 0x15, 0x6c, 0xac, 0xd0, 0x5c, 0x46, 0x2e,
	0x6e, 0x24, 0xb3, 0x85, 0x4c, 0xf0, 0x1a, 0x82, 0xc3, 0x0f, 0x52, 0xa6, 0x24, 0xea, 0x82, 0x78,
	0x40, 0x49, 0xb1, 0xe9, 0xf2, 0x93, 0xc9, 0x4c, 0xd2, 0x42, 0x92, 0xfa, 0x48, 0x9a, 0x20, 0xe1,
	0xa6, 0x5f, 0x0d, 0x0a, 0x88, 0x5a, 0x27, 0x8f, 0x13, 0x8f, 0xe4, 0x18, 0x2f, 0x3c, 0x92, 0x63,
	0x7c, 0xf0, 0x48, 0x8e, 0x71, 0xc2, 0x63, 0x39, 0x86, 0x0b, 0x8f, 0xe5, 0x18, 0x6e, 0x3c, 0x96,
	0x63, 0x88, 0xd2, 0x4b, 0xcf, 0x2c, 0xc9, 0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x87, 0xdb,
	0x0e, 0x61, 0x41, 0x0d, 0xaa, 0x80, 0x1a, 0x55, 0x52, 0x59, 0x90, 0x5a, 0x9c, 0xc4, 0x06, 0x8e,
	0x09, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0xa4, 0xaa, 0x04, 0x18, 0x2e, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
	// Queries a list of assetByUUID items.
	AssetByUUID(ctx context.Context, in *QueryAssetByUUIDRequest, opts ...grpc.CallOption) (*QueryAssetByUUIDResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) AssetByUUID(ctx context.Context, in *QueryAssetByUUIDRequest, opts ...grpc.CallOption) (*QueryAssetByUUIDResponse, error) {
	out := new(QueryAssetByUUIDResponse)
	err := c.cc.Invoke(ctx, "/ShareRing.shareledger.asset.Query/AssetByUUID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// Queries a list of assetByUUID items.
	AssetByUUID(context.Context, *QueryAssetByUUIDRequest) (*QueryAssetByUUIDResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) AssetByUUID(ctx context.Context, req *QueryAssetByUUIDRequest) (*QueryAssetByUUIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AssetByUUID not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_AssetByUUID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryAssetByUUIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).AssetByUUID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ShareRing.shareledger.asset.Query/AssetByUUID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).AssetByUUID(ctx, req.(*QueryAssetByUUIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ShareRing.shareledger.asset.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AssetByUUID",
			Handler:    _Query_AssetByUUID_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "asset/query.proto",
}

func (m *QueryAssetByUUIDRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryAssetByUUIDRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryAssetByUUIDRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Uuid) > 0 {
		i -= len(m.Uuid)
		copy(dAtA[i:], m.Uuid)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Uuid)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryAssetByUUIDResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryAssetByUUIDResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryAssetByUUIDResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Asset != nil {
		{
			size, err := m.Asset.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintQuery(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *QueryAssetByUUIDRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Uuid)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryAssetByUUIDResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Asset != nil {
		l = m.Asset.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryAssetByUUIDRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryAssetByUUIDRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryAssetByUUIDRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Uuid", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Uuid = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryAssetByUUIDResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryAssetByUUIDResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryAssetByUUIDResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Asset", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Asset == nil {
				m.Asset = &Asset{}
			}
			if err := m.Asset.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func skipQuery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
				return 0, ErrInvalidLengthQuery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuery = fmt.Errorf("proto: unexpected end of group")
)