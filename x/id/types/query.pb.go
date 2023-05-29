// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: shareledger/id/v1/query.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/types/query"
	_ "github.com/cosmos/gogoproto/gogoproto"
	grpc1 "github.com/cosmos/gogoproto/grpc"
	proto "github.com/cosmos/gogoproto/proto"
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

type QueryIdByAddressRequest struct {
	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
}

func (m *QueryIdByAddressRequest) Reset()         { *m = QueryIdByAddressRequest{} }
func (m *QueryIdByAddressRequest) String() string { return proto.CompactTextString(m) }
func (*QueryIdByAddressRequest) ProtoMessage()    {}
func (*QueryIdByAddressRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_fd17cba8d5de8407, []int{0}
}
func (m *QueryIdByAddressRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryIdByAddressRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryIdByAddressRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryIdByAddressRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryIdByAddressRequest.Merge(m, src)
}
func (m *QueryIdByAddressRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryIdByAddressRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryIdByAddressRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryIdByAddressRequest proto.InternalMessageInfo

type QueryIdByAddressResponse struct {
	Id *Id `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (m *QueryIdByAddressResponse) Reset()         { *m = QueryIdByAddressResponse{} }
func (m *QueryIdByAddressResponse) String() string { return proto.CompactTextString(m) }
func (*QueryIdByAddressResponse) ProtoMessage()    {}
func (*QueryIdByAddressResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_fd17cba8d5de8407, []int{1}
}
func (m *QueryIdByAddressResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryIdByAddressResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryIdByAddressResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryIdByAddressResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryIdByAddressResponse.Merge(m, src)
}
func (m *QueryIdByAddressResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryIdByAddressResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryIdByAddressResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryIdByAddressResponse proto.InternalMessageInfo

func (m *QueryIdByAddressResponse) GetId() *Id {
	if m != nil {
		return m.Id
	}
	return nil
}

type QueryIdByIdRequest struct {
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (m *QueryIdByIdRequest) Reset()         { *m = QueryIdByIdRequest{} }
func (m *QueryIdByIdRequest) String() string { return proto.CompactTextString(m) }
func (*QueryIdByIdRequest) ProtoMessage()    {}
func (*QueryIdByIdRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_fd17cba8d5de8407, []int{2}
}
func (m *QueryIdByIdRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryIdByIdRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryIdByIdRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryIdByIdRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryIdByIdRequest.Merge(m, src)
}
func (m *QueryIdByIdRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryIdByIdRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryIdByIdRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryIdByIdRequest proto.InternalMessageInfo

type QueryIdByIdResponse struct {
	Id *Id `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (m *QueryIdByIdResponse) Reset()         { *m = QueryIdByIdResponse{} }
func (m *QueryIdByIdResponse) String() string { return proto.CompactTextString(m) }
func (*QueryIdByIdResponse) ProtoMessage()    {}
func (*QueryIdByIdResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_fd17cba8d5de8407, []int{3}
}
func (m *QueryIdByIdResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryIdByIdResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryIdByIdResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryIdByIdResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryIdByIdResponse.Merge(m, src)
}
func (m *QueryIdByIdResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryIdByIdResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryIdByIdResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryIdByIdResponse proto.InternalMessageInfo

func (m *QueryIdByIdResponse) GetId() *Id {
	if m != nil {
		return m.Id
	}
	return nil
}

func init() {
	proto.RegisterType((*QueryIdByAddressRequest)(nil), "shareledger.id.QueryIdByAddressRequest")
	proto.RegisterType((*QueryIdByAddressResponse)(nil), "shareledger.id.QueryIdByAddressResponse")
	proto.RegisterType((*QueryIdByIdRequest)(nil), "shareledger.id.QueryIdByIdRequest")
	proto.RegisterType((*QueryIdByIdResponse)(nil), "shareledger.id.QueryIdByIdResponse")
}

func init() { proto.RegisterFile("shareledger/id/v1/query.proto", fileDescriptor_fd17cba8d5de8407) }

var fileDescriptor_fd17cba8d5de8407 = []byte{
	// 391 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0x4f, 0x4b, 0x32, 0x41,
	0x18, 0xdf, 0x5d, 0x78, 0x7d, 0x5f, 0x47, 0xf0, 0x30, 0x6f, 0xa0, 0x2d, 0xb5, 0xd6, 0x7a, 0xc8,
	0x0a, 0x76, 0xd0, 0x4e, 0x05, 0x05, 0x09, 0x1d, 0x3c, 0xe6, 0xb1, 0xdb, 0xac, 0x33, 0x8c, 0x03,
	0xba, 0xb3, 0xee, 0xac, 0x92, 0x88, 0x97, 0x4e, 0x45, 0x97, 0xa0, 0x2f, 0xe0, 0xc7, 0xe9, 0x28,
	0x74, 0xe9, 0x18, 0xda, 0xa1, 0x8f, 0x11, 0xbb, 0xb3, 0x8a, 0x7f, 0xb0, 0xba, 0x3d, 0xf3, 0x3c,
	0xbf, 0x3f, 0xcf, 0xf3, 0x63, 0xc0, 0xae, 0x6c, 0xe2, 0x80, 0xb6, 0x28, 0x61, 0x34, 0x40, 0x9c,
	0xa0, 0x5e, 0x19, 0x75, 0xba, 0x34, 0xe8, 0x3b, 0x7e, 0x20, 0x42, 0x01, 0xb3, 0x0b, 0x63, 0x87,
	0x13, 0x73, 0x87, 0x09, 0xc1, 0x5a, 0x14, 0x61, 0x9f, 0x23, 0xec, 0x79, 0x22, 0xc4, 0x21, 0x17,
	0x9e, 0x54, 0x68, 0xf3, 0xa8, 0x21, 0x64, 0x5b, 0x48, 0xe4, 0x62, 0x49, 0x95, 0x0c, 0xea, 0x95,
	0x5d, 0x1a, 0xe2, 0x32, 0xf2, 0x31, 0xe3, 0x5e, 0x0c, 0x4e, 0xb0, 0x5b, 0x4c, 0x30, 0x11, 0x97,
	0x28, 0xaa, 0x92, 0xae, 0xb9, 0xbe, 0x0e, 0x27, 0x6a, 0x66, 0x9f, 0x83, 0xdc, 0x75, 0xa4, 0x59,
	0x23, 0xd5, 0xfe, 0x25, 0x21, 0x01, 0x95, 0xb2, 0x4e, 0x3b, 0x5d, 0x2a, 0x43, 0x98, 0x07, 0x7f,
	0xb1, 0xea, 0xe4, 0xf5, 0x3d, 0xbd, 0x94, 0xae, 0xcf, 0x9e, 0x67, 0xff, 0xee, 0x47, 0x05, 0xed,
	0x73, 0x54, 0xd0, 0xec, 0x0b, 0x90, 0x5f, 0xa7, 0x4b, 0x5f, 0x78, 0x92, 0x42, 0x1b, 0x18, 0x9c,
	0xc4, 0xd4, 0x4c, 0x05, 0x3a, 0xcb, 0x37, 0x3b, 0x35, 0x52, 0x37, 0x38, 0xb1, 0x1d, 0x00, 0xe7,
	0xfc, 0x1a, 0x99, 0x39, 0x67, 0xe7, 0xcc, 0x74, 0x84, 0x5a, 0xf0, 0x3b, 0x05, 0xff, 0x97, 0xf0,
	0xbf, 0xb7, 0xaa, 0x3c, 0x18, 0xe0, 0x4f, 0xcc, 0x85, 0x8f, 0x3a, 0xc8, 0x2c, 0x2c, 0x0c, 0x0f,
	0x56, 0x19, 0x1b, 0x12, 0x31, 0x4b, 0x3f, 0x03, 0xd5, 0x42, 0xf6, 0xe1, 0xdd, 0xeb, 0xc7, 0xb3,
	0x51, 0x84, 0xfb, 0x68, 0x25, 0xfb, 0x24, 0x42, 0x34, 0x48, 0x8a, 0x21, 0x0c, 0x40, 0x4a, 0x5d,
	0x03, 0xed, 0x8d, 0xf2, 0xf3, 0x68, 0xcc, 0xe2, 0xb7, 0x98, 0xc4, 0xbd, 0x10, 0xbb, 0x6f, 0xc3,
	0xdc, 0xaa, 0x3b, 0x27, 0x68, 0xc0, 0xc9, 0xb0, 0x7a, 0xf5, 0x32, 0xb1, 0xf4, 0xf1, 0xc4, 0xd2,
	0xdf, 0x27, 0x96, 0xfe, 0x34, 0xb5, 0xb4, 0xf1, 0xd4, 0xd2, 0xde, 0xa6, 0x96, 0x76, 0x73, 0xcc,
	0x78, 0xd8, 0xec, 0xba, 0x4e, 0x43, 0xb4, 0x15, 0x39, 0xe0, 0x1e, 0x5b, 0x92, 0xb9, 0x8d, 0x54,
	0xc2, 0xbe, 0x4f, 0xa5, 0x9b, 0x8a, 0xff, 0xd0, 0xc9, 0x57, 0x00, 0x00, 0x00, 0xff, 0xff, 0xd3,
	0x8b, 0xdc, 0x8e, 0xf0, 0x02, 0x00, 0x00,
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
	// Queries a list of idByAddress items.
	IdByAddress(ctx context.Context, in *QueryIdByAddressRequest, opts ...grpc.CallOption) (*QueryIdByAddressResponse, error)
	// Queries a list of idById items.
	IdById(ctx context.Context, in *QueryIdByIdRequest, opts ...grpc.CallOption) (*QueryIdByIdResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) IdByAddress(ctx context.Context, in *QueryIdByAddressRequest, opts ...grpc.CallOption) (*QueryIdByAddressResponse, error) {
	out := new(QueryIdByAddressResponse)
	err := c.cc.Invoke(ctx, "/shareledger.id.Query/IdByAddress", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) IdById(ctx context.Context, in *QueryIdByIdRequest, opts ...grpc.CallOption) (*QueryIdByIdResponse, error) {
	out := new(QueryIdByIdResponse)
	err := c.cc.Invoke(ctx, "/shareledger.id.Query/IdById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// Queries a list of idByAddress items.
	IdByAddress(context.Context, *QueryIdByAddressRequest) (*QueryIdByAddressResponse, error)
	// Queries a list of idById items.
	IdById(context.Context, *QueryIdByIdRequest) (*QueryIdByIdResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) IdByAddress(ctx context.Context, req *QueryIdByAddressRequest) (*QueryIdByAddressResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IdByAddress not implemented")
}
func (*UnimplementedQueryServer) IdById(ctx context.Context, req *QueryIdByIdRequest) (*QueryIdByIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IdById not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_IdByAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryIdByAddressRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).IdByAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/shareledger.id.Query/IdByAddress",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).IdByAddress(ctx, req.(*QueryIdByAddressRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_IdById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryIdByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).IdById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/shareledger.id.Query/IdById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).IdById(ctx, req.(*QueryIdByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "shareledger.id.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "IdByAddress",
			Handler:    _Query_IdByAddress_Handler,
		},
		{
			MethodName: "IdById",
			Handler:    _Query_IdById_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "shareledger/id/v1/query.proto",
}

func (m *QueryIdByAddressRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryIdByAddressRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryIdByAddressRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryIdByAddressResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryIdByAddressResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryIdByAddressResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Id != nil {
		{
			size, err := m.Id.MarshalToSizedBuffer(dAtA[:i])
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

func (m *QueryIdByIdRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryIdByIdRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryIdByIdRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryIdByIdResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryIdByIdResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryIdByIdResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Id != nil {
		{
			size, err := m.Id.MarshalToSizedBuffer(dAtA[:i])
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
func (m *QueryIdByAddressRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryIdByAddressResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != nil {
		l = m.Id.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryIdByIdRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryIdByIdResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != nil {
		l = m.Id.Size()
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
func (m *QueryIdByAddressRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryIdByAddressRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryIdByAddressRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
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
			m.Address = string(dAtA[iNdEx:postIndex])
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
func (m *QueryIdByAddressResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryIdByAddressResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryIdByAddressResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
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
			if m.Id == nil {
				m.Id = &Id{}
			}
			if err := m.Id.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *QueryIdByIdRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryIdByIdRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryIdByIdRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
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
			m.Id = string(dAtA[iNdEx:postIndex])
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
func (m *QueryIdByIdResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryIdByIdResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryIdByIdResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
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
			if m.Id == nil {
				m.Id = &Id{}
			}
			if err := m.Id.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
