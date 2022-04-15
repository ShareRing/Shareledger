// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: swap/query.proto

package types

import (
	context "context"
	fmt "fmt"
	query "github.com/cosmos/cosmos-sdk/types/query"
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

// QueryParamsRequest is request type for the Query/Params RPC method.
type QueryParamsRequest struct {
}

func (m *QueryParamsRequest) Reset()         { *m = QueryParamsRequest{} }
func (m *QueryParamsRequest) String() string { return proto.CompactTextString(m) }
func (*QueryParamsRequest) ProtoMessage()    {}
func (*QueryParamsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ba9b806334123578, []int{0}
}
func (m *QueryParamsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryParamsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryParamsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryParamsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryParamsRequest.Merge(m, src)
}
func (m *QueryParamsRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryParamsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryParamsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryParamsRequest proto.InternalMessageInfo

// QueryParamsResponse is response type for the Query/Params RPC method.
type QueryParamsResponse struct {
	// params holds all the parameters of this module.
	Params Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
}

func (m *QueryParamsResponse) Reset()         { *m = QueryParamsResponse{} }
func (m *QueryParamsResponse) String() string { return proto.CompactTextString(m) }
func (*QueryParamsResponse) ProtoMessage()    {}
func (*QueryParamsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ba9b806334123578, []int{1}
}
func (m *QueryParamsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryParamsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryParamsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryParamsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryParamsResponse.Merge(m, src)
}
func (m *QueryParamsResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryParamsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryParamsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryParamsResponse proto.InternalMessageInfo

func (m *QueryParamsResponse) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

type QuerySearchRequest struct {
	Status      string             `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Id          uint64             `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	SrcAddr     string             `protobuf:"bytes,3,opt,name=srcAddr,proto3" json:"srcAddr,omitempty"`
	DestAddr    string             `protobuf:"bytes,4,opt,name=destAddr,proto3" json:"destAddr,omitempty"`
	SrcNetwork  string             `protobuf:"bytes,5,opt,name=srcNetwork,proto3" json:"srcNetwork,omitempty"`
	DestNetwork string             `protobuf:"bytes,6,opt,name=destNetwork,proto3" json:"destNetwork,omitempty"`
	Pagination  *query.PageRequest `protobuf:"bytes,7,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QuerySearchRequest) Reset()         { *m = QuerySearchRequest{} }
func (m *QuerySearchRequest) String() string { return proto.CompactTextString(m) }
func (*QuerySearchRequest) ProtoMessage()    {}
func (*QuerySearchRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ba9b806334123578, []int{2}
}
func (m *QuerySearchRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QuerySearchRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QuerySearchRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QuerySearchRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QuerySearchRequest.Merge(m, src)
}
func (m *QuerySearchRequest) XXX_Size() int {
	return m.Size()
}
func (m *QuerySearchRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QuerySearchRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QuerySearchRequest proto.InternalMessageInfo

func (m *QuerySearchRequest) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *QuerySearchRequest) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *QuerySearchRequest) GetSrcAddr() string {
	if m != nil {
		return m.SrcAddr
	}
	return ""
}

func (m *QuerySearchRequest) GetDestAddr() string {
	if m != nil {
		return m.DestAddr
	}
	return ""
}

func (m *QuerySearchRequest) GetSrcNetwork() string {
	if m != nil {
		return m.SrcNetwork
	}
	return ""
}

func (m *QuerySearchRequest) GetDestNetwork() string {
	if m != nil {
		return m.DestNetwork
	}
	return ""
}

func (m *QuerySearchRequest) GetPagination() *query.PageRequest {
	if m != nil {
		return m.Pagination
	}
	return nil
}

type QuerySearchResponse struct {
	Request    []Request           `protobuf:"bytes,1,rep,name=Request,proto3" json:"Request"`
	Pagination *query.PageResponse `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QuerySearchResponse) Reset()         { *m = QuerySearchResponse{} }
func (m *QuerySearchResponse) String() string { return proto.CompactTextString(m) }
func (*QuerySearchResponse) ProtoMessage()    {}
func (*QuerySearchResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ba9b806334123578, []int{3}
}
func (m *QuerySearchResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QuerySearchResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QuerySearchResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QuerySearchResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QuerySearchResponse.Merge(m, src)
}
func (m *QuerySearchResponse) XXX_Size() int {
	return m.Size()
}
func (m *QuerySearchResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QuerySearchResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QuerySearchResponse proto.InternalMessageInfo

func (m *QuerySearchResponse) GetRequest() []Request {
	if m != nil {
		return m.Request
	}
	return nil
}

func (m *QuerySearchResponse) GetPagination() *query.PageResponse {
	if m != nil {
		return m.Pagination
	}
	return nil
}

func init() {
	proto.RegisterType((*QueryParamsRequest)(nil), "shareledger.swap.QueryParamsRequest")
	proto.RegisterType((*QueryParamsResponse)(nil), "shareledger.swap.QueryParamsResponse")
	proto.RegisterType((*QuerySearchRequest)(nil), "shareledger.swap.QuerySearchRequest")
	proto.RegisterType((*QuerySearchResponse)(nil), "shareledger.swap.QuerySearchResponse")
}

func init() { proto.RegisterFile("swap/query.proto", fileDescriptor_ba9b806334123578) }

var fileDescriptor_ba9b806334123578 = []byte{
	// 538 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x53, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0xcd, 0xba, 0xa9, 0x03, 0x1b, 0x81, 0xca, 0x52, 0x21, 0x13, 0x21, 0x13, 0x59, 0x05, 0xa2,
	0x4a, 0x78, 0xd5, 0x20, 0x21, 0x71, 0xa4, 0x07, 0x7a, 0x02, 0x15, 0x73, 0xe3, 0xc4, 0xda, 0x5e,
	0x39, 0x16, 0x8d, 0xd7, 0xdd, 0xdd, 0x50, 0xaa, 0x28, 0x07, 0xf8, 0x02, 0x24, 0x4e, 0x7c, 0x0d,
	0xd7, 0x5e, 0x90, 0x2a, 0x71, 0xe1, 0x84, 0x50, 0xc2, 0x67, 0x70, 0x40, 0x9e, 0x5d, 0x83, 0x43,
	0x14, 0x7a, 0xf3, 0xcc, 0xbc, 0xd9, 0x79, 0xf3, 0xe6, 0x19, 0x6f, 0xa9, 0x13, 0x56, 0xd2, 0xe3,
	0x09, 0x97, 0xa7, 0x61, 0x29, 0x85, 0x16, 0x64, 0x4b, 0x8d, 0x98, 0xe4, 0x47, 0x3c, 0xcd, 0xb8,
	0x0c, 0xab, 0x6a, 0x6f, 0x3b, 0x13, 0x99, 0x80, 0x22, 0xad, 0xbe, 0x0c, 0xae, 0x77, 0x2b, 0x13,
	0x22, 0x3b, 0xe2, 0x94, 0x95, 0x39, 0x65, 0x45, 0x21, 0x34, 0xd3, 0xb9, 0x28, 0x94, 0xad, 0xee,
	0x26, 0x42, 0x8d, 0x85, 0xa2, 0x31, 0x53, 0xdc, 0x3c, 0x4f, 0xdf, 0xec, 0xc5, 0x5c, 0xb3, 0x3d,
	0x5a, 0xb2, 0x2c, 0x2f, 0x00, 0x6c, 0xb1, 0xd7, 0x80, 0x43, 0xc9, 0x24, 0x1b, 0xd7, 0xed, 0x57,
	0x20, 0x95, 0xa7, 0x36, 0x24, 0x10, 0x4a, 0x7e, 0x3c, 0xe1, 0x4a, 0x9b, 0x5c, 0xb0, 0x8d, 0xc9,
	0xf3, 0xea, 0xdd, 0x43, 0xe8, 0x8b, 0x4c, 0x2d, 0x78, 0x8a, 0xaf, 0x2f, 0x65, 0x55, 0x29, 0x0a,
	0xc5, 0xc9, 0x43, 0xec, 0x9a, 0xf7, 0x3d, 0xd4, 0x47, 0x83, 0xee, 0xd0, 0x0b, 0xff, 0xdd, 0x32,
	0x34, 0x1d, 0xfb, 0xed, 0xb3, 0xef, 0xb7, 0x5b, 0x91, 0x45, 0x07, 0xbf, 0x90, 0x9d, 0xf2, 0x82,
	0x33, 0x99, 0x8c, 0xec, 0x14, 0x72, 0x03, 0xbb, 0x4a, 0x33, 0x3d, 0x31, 0xcf, 0x5d, 0x8e, 0x6c,
	0x44, 0xae, 0x62, 0x27, 0x4f, 0x3d, 0xa7, 0x8f, 0x06, 0xed, 0xc8, 0xc9, 0x53, 0xe2, 0xe1, 0x8e,
	0x92, 0xc9, 0xe3, 0x34, 0x95, 0xde, 0x06, 0x00, 0xeb, 0x90, 0xf4, 0xf0, 0xa5, 0x94, 0x2b, 0x0d,
	0xa5, 0x36, 0x94, 0xfe, 0xc4, 0xc4, 0xc7, 0x58, 0xc9, 0xe4, 0x19, 0xd7, 0x27, 0x42, 0xbe, 0xf6,
	0x36, 0xa1, 0xda, 0xc8, 0x90, 0x3e, 0xee, 0x56, 0xd8, 0x1a, 0xe0, 0x02, 0xa0, 0x99, 0x22, 0x4f,
	0x30, 0xfe, 0xab, 0xb2, 0xd7, 0x81, 0x95, 0xef, 0x86, 0xe6, 0x24, 0x61, 0x75, 0x92, 0xd0, 0x5c,
	0xdc, 0x9e, 0x24, 0x3c, 0x64, 0x19, 0xb7, 0xbb, 0x45, 0x8d, 0xce, 0xe0, 0x13, 0xb2, 0x72, 0xd6,
	0xeb, 0x5b, 0x39, 0x1f, 0xe1, 0x8e, 0x85, 0x7b, 0xa8, 0xbf, 0x31, 0xe8, 0x0e, 0x6f, 0xae, 0xea,
	0x69, 0x01, 0x56, 0xd0, 0x1a, 0x4f, 0x0e, 0x96, 0xa8, 0x39, 0x40, 0xed, 0xde, 0x85, 0xd4, 0xcc,
	0xdc, 0x26, 0xb7, 0xe1, 0x17, 0x07, 0x6f, 0x02, 0x37, 0xf2, 0x0e, 0x61, 0xd7, 0x5c, 0x8f, 0xec,
	0xac, 0xf2, 0x58, 0x35, 0x49, 0xef, 0xce, 0x05, 0x28, 0x33, 0x2d, 0xd8, 0x7d, 0xff, 0xf5, 0xe7,
	0x47, 0x67, 0x87, 0x04, 0x14, 0xe0, 0x32, 0x2f, 0x32, 0xda, 0x68, 0xa4, 0x0d, 0xdb, 0x92, 0xcf,
	0x08, 0xbb, 0x46, 0xa4, 0xb5, 0x1c, 0x96, 0x2c, 0xb4, 0x96, 0xc3, 0xb2, 0xd2, 0xc1, 0x08, 0x38,
	0xc4, 0xe4, 0xd5, 0xff, 0x38, 0x28, 0xe8, 0xa1, 0x53, 0xe3, 0xc2, 0x19, 0x9d, 0xe6, 0xe9, 0x8c,
	0x4e, 0xad, 0xd5, 0x66, 0x74, 0x5a, 0x3b, 0xcb, 0x24, 0xad, 0x49, 0x6c, 0xbe, 0x8e, 0xf6, 0x0f,
	0xce, 0xe6, 0x3e, 0x3a, 0x9f, 0xfb, 0xe8, 0xc7, 0xdc, 0x47, 0x1f, 0x16, 0x7e, 0xeb, 0x7c, 0xe1,
	0xb7, 0xbe, 0x2d, 0xfc, 0xd6, 0xcb, 0xfb, 0x59, 0xae, 0x47, 0x93, 0x38, 0x4c, 0xc4, 0x78, 0x0d,
	0x8b, 0xb7, 0x86, 0x87, 0x3e, 0x2d, 0xb9, 0x8a, 0x5d, 0xf8, 0x3f, 0x1f, 0xfc, 0x0e, 0x00, 0x00,
	0xff, 0xff, 0x4d, 0x8d, 0x41, 0xe9, 0x5b, 0x04, 0x00, 0x00,
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
	// Parameters queries the parameters of the module.
	Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error)
	// Queries a list of Search items.
	Search(ctx context.Context, in *QuerySearchRequest, opts ...grpc.CallOption) (*QuerySearchResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error) {
	out := new(QueryParamsResponse)
	err := c.cc.Invoke(ctx, "/shareledger.swap.Query/Params", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Search(ctx context.Context, in *QuerySearchRequest, opts ...grpc.CallOption) (*QuerySearchResponse, error) {
	out := new(QuerySearchResponse)
	err := c.cc.Invoke(ctx, "/shareledger.swap.Query/Search", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// Parameters queries the parameters of the module.
	Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error)
	// Queries a list of Search items.
	Search(context.Context, *QuerySearchRequest) (*QuerySearchResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) Params(ctx context.Context, req *QueryParamsRequest) (*QueryParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Params not implemented")
}
func (*UnimplementedQueryServer) Search(ctx context.Context, req *QuerySearchRequest) (*QuerySearchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Search not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_Params_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryParamsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Params(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/shareledger.swap.Query/Params",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Params(ctx, req.(*QueryParamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Search_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QuerySearchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Search(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/shareledger.swap.Query/Search",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Search(ctx, req.(*QuerySearchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "shareledger.swap.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Params",
			Handler:    _Query_Params_Handler,
		},
		{
			MethodName: "Search",
			Handler:    _Query_Search_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "swap/query.proto",
}

func (m *QueryParamsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryParamsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryParamsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryParamsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryParamsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryParamsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *QuerySearchRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QuerySearchRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QuerySearchRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x3a
	}
	if len(m.DestNetwork) > 0 {
		i -= len(m.DestNetwork)
		copy(dAtA[i:], m.DestNetwork)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.DestNetwork)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.SrcNetwork) > 0 {
		i -= len(m.SrcNetwork)
		copy(dAtA[i:], m.SrcNetwork)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.SrcNetwork)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.DestAddr) > 0 {
		i -= len(m.DestAddr)
		copy(dAtA[i:], m.DestAddr)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.DestAddr)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.SrcAddr) > 0 {
		i -= len(m.SrcAddr)
		copy(dAtA[i:], m.SrcAddr)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.SrcAddr)))
		i--
		dAtA[i] = 0x1a
	}
	if m.Id != 0 {
		i = encodeVarintQuery(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Status) > 0 {
		i -= len(m.Status)
		copy(dAtA[i:], m.Status)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Status)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QuerySearchResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QuerySearchResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QuerySearchResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.Request) > 0 {
		for iNdEx := len(m.Request) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Request[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintQuery(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
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
func (m *QueryParamsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryParamsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func (m *QuerySearchRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Status)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	if m.Id != 0 {
		n += 1 + sovQuery(uint64(m.Id))
	}
	l = len(m.SrcAddr)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	l = len(m.DestAddr)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	l = len(m.SrcNetwork)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	l = len(m.DestNetwork)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QuerySearchResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Request) > 0 {
		for _, e := range m.Request {
			l = e.Size()
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	if m.Pagination != nil {
		l = m.Pagination.Size()
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
func (m *QueryParamsRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryParamsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryParamsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
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
func (m *QueryParamsResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryParamsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryParamsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
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
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *QuerySearchRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QuerySearchRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QuerySearchRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
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
			m.Status = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SrcAddr", wireType)
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
			m.SrcAddr = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DestAddr", wireType)
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
			m.DestAddr = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SrcNetwork", wireType)
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
			m.SrcNetwork = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DestNetwork", wireType)
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
			m.DestNetwork = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
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
			if m.Pagination == nil {
				m.Pagination = &query.PageRequest{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *QuerySearchResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QuerySearchResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QuerySearchResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Request", wireType)
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
			m.Request = append(m.Request, Request{})
			if err := m.Request[len(m.Request)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
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
			if m.Pagination == nil {
				m.Pagination = &query.PageResponse{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
