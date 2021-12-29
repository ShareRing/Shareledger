// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: gentlemint/query.proto

package types

import (
	context "context"
	encoding_binary "encoding/binary"
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

type QueryGetExchangeRateRequest struct {
}

func (m *QueryGetExchangeRateRequest) Reset()         { *m = QueryGetExchangeRateRequest{} }
func (m *QueryGetExchangeRateRequest) String() string { return proto.CompactTextString(m) }
func (*QueryGetExchangeRateRequest) ProtoMessage()    {}
func (*QueryGetExchangeRateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c3912bb3197ea16, []int{0}
}
func (m *QueryGetExchangeRateRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryGetExchangeRateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryGetExchangeRateRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryGetExchangeRateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryGetExchangeRateRequest.Merge(m, src)
}
func (m *QueryGetExchangeRateRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryGetExchangeRateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryGetExchangeRateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryGetExchangeRateRequest proto.InternalMessageInfo

type QueryGetExchangeRateResponse struct {
	Rate float64 `protobuf:"fixed64,1,opt,name=rate,proto3" json:"rate,omitempty"`
}

func (m *QueryGetExchangeRateResponse) Reset()         { *m = QueryGetExchangeRateResponse{} }
func (m *QueryGetExchangeRateResponse) String() string { return proto.CompactTextString(m) }
func (*QueryGetExchangeRateResponse) ProtoMessage()    {}
func (*QueryGetExchangeRateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c3912bb3197ea16, []int{1}
}
func (m *QueryGetExchangeRateResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryGetExchangeRateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryGetExchangeRateResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryGetExchangeRateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryGetExchangeRateResponse.Merge(m, src)
}
func (m *QueryGetExchangeRateResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryGetExchangeRateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryGetExchangeRateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryGetExchangeRateResponse proto.InternalMessageInfo

func (m *QueryGetExchangeRateResponse) GetRate() float64 {
	if m != nil {
		return m.Rate
	}
	return 0
}

func init() {
	proto.RegisterType((*QueryGetExchangeRateRequest)(nil), "shareledger.gentlemint.QueryGetExchangeRateRequest")
	proto.RegisterType((*QueryGetExchangeRateResponse)(nil), "shareledger.gentlemint.QueryGetExchangeRateResponse")
}

func init() { proto.RegisterFile("gentlemint/query.proto", fileDescriptor_4c3912bb3197ea16) }

var fileDescriptor_4c3912bb3197ea16 = []byte{
	// 311 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x90, 0xbf, 0x4a, 0x03, 0x41,
	0x10, 0xc6, 0xb3, 0xa2, 0x16, 0x8b, 0xd5, 0x22, 0x41, 0x62, 0x5c, 0x24, 0x8d, 0x7f, 0x8a, 0x5b,
	0x92, 0xf8, 0x04, 0x82, 0x58, 0x59, 0x78, 0x76, 0x36, 0xb2, 0x17, 0x87, 0xcd, 0x42, 0x6e, 0xe7,
	0x72, 0x3b, 0x27, 0x49, 0xeb, 0x13, 0x08, 0xbe, 0x85, 0xb5, 0x0f, 0x61, 0x19, 0xb0, 0xb1, 0x94,
	0x3b, 0x1f, 0x44, 0xee, 0x0f, 0x78, 0x42, 0x10, 0xec, 0x3e, 0x66, 0xbe, 0xdf, 0x37, 0xc3, 0xc7,
	0xbb, 0x06, 0x1c, 0xcd, 0x20, 0xb6, 0x8e, 0xd4, 0x3c, 0x83, 0x74, 0x19, 0x24, 0x29, 0x12, 0x8a,
	0xae, 0x9f, 0xea, 0x14, 0x66, 0x70, 0x6f, 0x20, 0x0d, 0x7e, 0x3c, 0xbd, 0xbe, 0x41, 0x34, 0x33,
	0x50, 0x3a, 0xb1, 0x4a, 0x3b, 0x87, 0xa4, 0xc9, 0xa2, 0xf3, 0x35, 0xd5, 0x3b, 0x9d, 0xa0, 0x8f,
	0xd1, 0xab, 0x48, 0x7b, 0xa8, 0xe3, 0xd4, 0xc3, 0x30, 0x02, 0xd2, 0x43, 0x95, 0x68, 0x63, 0x5d,
	0x65, 0x6e, 0xbc, 0xb2, 0x75, 0x19, 0x16, 0x93, 0xa9, 0x76, 0x06, 0xee, 0x52, 0x4d, 0xd0, 0xec,
	0x77, 0x0d, 0x1a, 0xac, 0xa4, 0x2a, 0x55, 0x3d, 0x1d, 0x1c, 0xf0, 0xfd, 0xeb, 0x32, 0xf7, 0x12,
	0xe8, 0xa2, 0x81, 0x42, 0x4d, 0x10, 0xc2, 0x3c, 0x03, 0x4f, 0x83, 0x11, 0xef, 0xaf, 0x5f, 0xfb,
	0x04, 0x9d, 0x07, 0x21, 0xf8, 0x66, 0x79, 0x62, 0x8f, 0x1d, 0xb2, 0x63, 0x16, 0x56, 0x7a, 0xf4,
	0xca, 0xf8, 0x56, 0x05, 0x89, 0x17, 0xc6, 0x77, 0xda, 0x98, 0x18, 0x07, 0xeb, 0x6b, 0x08, 0xfe,
	0xf8, 0xa1, 0x77, 0xf6, 0x3f, 0xa8, 0xfe, 0x6c, 0xa0, 0x1e, 0xdf, 0xbf, 0x9e, 0x37, 0x4e, 0xc4,
	0x91, 0x6a, 0xd1, 0xaa, 0xd5, 0x91, 0xf9, 0x0d, 0x9e, 0x5f, 0xbd, 0xe5, 0x92, 0xad, 0x72, 0xc9,
	0x3e, 0x73, 0xc9, 0x9e, 0x0a, 0xd9, 0x59, 0x15, 0xb2, 0xf3, 0x51, 0xc8, 0xce, 0xed, 0xd8, 0x58,
	0x9a, 0x66, 0x51, 0x30, 0xc1, 0x58, 0xdd, 0x94, 0x61, 0xa1, 0x75, 0xa6, 0x56, 0x4d, 0xec, 0xa2,
	0x1d, 0x4c, 0xcb, 0x04, 0x7c, 0xb4, 0x5d, 0xf5, 0x3b, 0xfe, 0x0e, 0x00, 0x00, 0xff, 0xff, 0x37,
	0xa0, 0x33, 0x7f, 0x11, 0x02, 0x00, 0x00,
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
	// Queries a exchangeRate by index.
	ExchangeRate(ctx context.Context, in *QueryGetExchangeRateRequest, opts ...grpc.CallOption) (*QueryGetExchangeRateResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) ExchangeRate(ctx context.Context, in *QueryGetExchangeRateRequest, opts ...grpc.CallOption) (*QueryGetExchangeRateResponse, error) {
	out := new(QueryGetExchangeRateResponse)
	err := c.cc.Invoke(ctx, "/shareledger.gentlemint.Query/ExchangeRate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// Queries a exchangeRate by index.
	ExchangeRate(context.Context, *QueryGetExchangeRateRequest) (*QueryGetExchangeRateResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) ExchangeRate(ctx context.Context, req *QueryGetExchangeRateRequest) (*QueryGetExchangeRateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExchangeRate not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_ExchangeRate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryGetExchangeRateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).ExchangeRate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/shareledger.gentlemint.Query/ExchangeRate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).ExchangeRate(ctx, req.(*QueryGetExchangeRateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "shareledger.gentlemint.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ExchangeRate",
			Handler:    _Query_ExchangeRate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "gentlemint/query.proto",
}

func (m *QueryGetExchangeRateRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryGetExchangeRateRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryGetExchangeRateRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryGetExchangeRateResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryGetExchangeRateResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryGetExchangeRateResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Rate != 0 {
		i -= 8
		encoding_binary.LittleEndian.PutUint64(dAtA[i:], uint64(math.Float64bits(float64(m.Rate))))
		i--
		dAtA[i] = 0x9
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
func (m *QueryGetExchangeRateRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryGetExchangeRateResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Rate != 0 {
		n += 9
	}
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryGetExchangeRateRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryGetExchangeRateRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryGetExchangeRateRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
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
func (m *QueryGetExchangeRateResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryGetExchangeRateResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryGetExchangeRateResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field Rate", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint64(encoding_binary.LittleEndian.Uint64(dAtA[iNdEx:]))
			iNdEx += 8
			m.Rate = float64(math.Float64frombits(v))
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
