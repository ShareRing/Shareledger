// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: document/query.proto

package types

import (
	context "context"
	fmt "fmt"
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

type QueryDocumentByProofRequest struct {
	Proof string `protobuf:"bytes,1,opt,name=proof,proto3" json:"proof,omitempty"`
}

func (m *QueryDocumentByProofRequest) Reset()         { *m = QueryDocumentByProofRequest{} }
func (m *QueryDocumentByProofRequest) String() string { return proto.CompactTextString(m) }
func (*QueryDocumentByProofRequest) ProtoMessage()    {}
func (*QueryDocumentByProofRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d6f5760b19c0d3bb, []int{0}
}
func (m *QueryDocumentByProofRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryDocumentByProofRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryDocumentByProofRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryDocumentByProofRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryDocumentByProofRequest.Merge(m, src)
}
func (m *QueryDocumentByProofRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryDocumentByProofRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryDocumentByProofRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryDocumentByProofRequest proto.InternalMessageInfo

type QueryDocumentByProofResponse struct {
	Document *Document `protobuf:"bytes,1,opt,name=document,proto3" json:"document,omitempty"`
}

func (m *QueryDocumentByProofResponse) Reset()         { *m = QueryDocumentByProofResponse{} }
func (m *QueryDocumentByProofResponse) String() string { return proto.CompactTextString(m) }
func (*QueryDocumentByProofResponse) ProtoMessage()    {}
func (*QueryDocumentByProofResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d6f5760b19c0d3bb, []int{1}
}
func (m *QueryDocumentByProofResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryDocumentByProofResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryDocumentByProofResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryDocumentByProofResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryDocumentByProofResponse.Merge(m, src)
}
func (m *QueryDocumentByProofResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryDocumentByProofResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryDocumentByProofResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryDocumentByProofResponse proto.InternalMessageInfo

func (m *QueryDocumentByProofResponse) GetDocument() *Document {
	if m != nil {
		return m.Document
	}
	return nil
}

type QueryDocumentByHolderIdRequest struct {
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (m *QueryDocumentByHolderIdRequest) Reset()         { *m = QueryDocumentByHolderIdRequest{} }
func (m *QueryDocumentByHolderIdRequest) String() string { return proto.CompactTextString(m) }
func (*QueryDocumentByHolderIdRequest) ProtoMessage()    {}
func (*QueryDocumentByHolderIdRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d6f5760b19c0d3bb, []int{2}
}
func (m *QueryDocumentByHolderIdRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryDocumentByHolderIdRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryDocumentByHolderIdRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryDocumentByHolderIdRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryDocumentByHolderIdRequest.Merge(m, src)
}
func (m *QueryDocumentByHolderIdRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryDocumentByHolderIdRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryDocumentByHolderIdRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryDocumentByHolderIdRequest proto.InternalMessageInfo

type QueryDocumentByHolderIdResponse struct {
	Document []*Document `protobuf:"bytes,1,rep,name=document,proto3" json:"document,omitempty"`
}

func (m *QueryDocumentByHolderIdResponse) Reset()         { *m = QueryDocumentByHolderIdResponse{} }
func (m *QueryDocumentByHolderIdResponse) String() string { return proto.CompactTextString(m) }
func (*QueryDocumentByHolderIdResponse) ProtoMessage()    {}
func (*QueryDocumentByHolderIdResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d6f5760b19c0d3bb, []int{3}
}
func (m *QueryDocumentByHolderIdResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryDocumentByHolderIdResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryDocumentByHolderIdResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryDocumentByHolderIdResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryDocumentByHolderIdResponse.Merge(m, src)
}
func (m *QueryDocumentByHolderIdResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryDocumentByHolderIdResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryDocumentByHolderIdResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryDocumentByHolderIdResponse proto.InternalMessageInfo

func (m *QueryDocumentByHolderIdResponse) GetDocument() []*Document {
	if m != nil {
		return m.Document
	}
	return nil
}

func init() {
	proto.RegisterType((*QueryDocumentByProofRequest)(nil), "ShareRing.shareledger.document.QueryDocumentByProofRequest")
	proto.RegisterType((*QueryDocumentByProofResponse)(nil), "ShareRing.shareledger.document.QueryDocumentByProofResponse")
	proto.RegisterType((*QueryDocumentByHolderIdRequest)(nil), "ShareRing.shareledger.document.QueryDocumentByHolderIdRequest")
	proto.RegisterType((*QueryDocumentByHolderIdResponse)(nil), "ShareRing.shareledger.document.QueryDocumentByHolderIdResponse")
}

func init() { proto.RegisterFile("document/query.proto", fileDescriptor_d6f5760b19c0d3bb) }

var fileDescriptor_d6f5760b19c0d3bb = []byte{
	// 390 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x49, 0xc9, 0x4f, 0x2e,
	0xcd, 0x4d, 0xcd, 0x2b, 0xd1, 0x2f, 0x2c, 0x4d, 0x2d, 0xaa, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9,
	0x17, 0x92, 0x0b, 0xce, 0x48, 0x2c, 0x4a, 0x0d, 0xca, 0xcc, 0x4b, 0xd7, 0x2b, 0x06, 0xb1, 0x72,
	0x52, 0x53, 0xd2, 0x53, 0x8b, 0xf4, 0x60, 0x6a, 0xa5, 0x44, 0xd2, 0xf3, 0xd3, 0xf3, 0xc1, 0x4a,
	0xf5, 0x41, 0x2c, 0x88, 0x2e, 0x29, 0x99, 0xf4, 0xfc, 0xfc, 0xf4, 0x9c, 0x54, 0xfd, 0xc4, 0x82,
	0x4c, 0xfd, 0xc4, 0xbc, 0xbc, 0xfc, 0x92, 0xc4, 0x92, 0xcc, 0xfc, 0xbc, 0x62, 0xa8, 0xac, 0x20,
	0xdc, 0xa6, 0x92, 0x0a, 0x88, 0x90, 0x92, 0x2d, 0x97, 0x74, 0x20, 0xc8, 0x56, 0x17, 0xa8, 0x8c,
	0x53, 0x65, 0x40, 0x51, 0x7e, 0x7e, 0x5a, 0x50, 0x6a, 0x61, 0x69, 0x6a, 0x71, 0x89, 0x90, 0x08,
	0x17, 0x6b, 0x01, 0x88, 0x2f, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0x19, 0x04, 0xe1, 0x58, 0x71, 0x74,
	0x2c, 0x90, 0x67, 0x78, 0xb1, 0x40, 0x9e, 0x41, 0x29, 0x85, 0x4b, 0x06, 0xbb, 0xf6, 0xe2, 0x82,
	0xfc, 0xbc, 0xe2, 0x54, 0x21, 0x17, 0x2e, 0x0e, 0x98, 0x9d, 0x60, 0x23, 0xb8, 0x8d, 0x34, 0xf4,
	0xf0, 0x7b, 0x4c, 0x0f, 0x66, 0x54, 0x10, 0x5c, 0xa7, 0x92, 0x15, 0x97, 0x1c, 0x9a, 0x2d, 0x1e,
	0xf9, 0x39, 0x29, 0xa9, 0x45, 0x9e, 0x29, 0x30, 0x77, 0xf2, 0x71, 0x31, 0x65, 0xa6, 0x40, 0x1d,
	0xc9, 0x94, 0x99, 0x82, 0xe4, 0xc2, 0x74, 0x2e, 0x79, 0x9c, 0x7a, 0xb1, 0x3a, 0x92, 0x99, 0x3c,
	0x47, 0x1a, 0xb5, 0x31, 0x73, 0xb1, 0x82, 0x6d, 0x12, 0xda, 0xc9, 0xc8, 0xc5, 0x8f, 0x16, 0x20,
	0x42, 0xd6, 0x84, 0x4c, 0xc4, 0x13, 0x0b, 0x52, 0x36, 0xe4, 0x69, 0x86, 0x78, 0x4f, 0x49, 0xbb,
	0xe9, 0xf2, 0x93, 0xc9, 0x4c, 0xaa, 0x42, 0xca, 0xfa, 0x48, 0x7a, 0xf5, 0xe1, 0x49, 0x01, 0x1c,
	0xa5, 0xfa, 0xd5, 0x60, 0xaa, 0x56, 0xe8, 0x00, 0x23, 0x97, 0x10, 0x66, 0x50, 0x09, 0xd9, 0x91,
	0xe8, 0x02, 0xb4, 0xf8, 0x91, 0xb2, 0x27, 0x5b, 0x3f, 0xd4, 0x13, 0x9a, 0x60, 0x4f, 0x28, 0x0b,
	0x29, 0x62, 0xf7, 0x44, 0x06, 0x58, 0xbd, 0x7e, 0x75, 0x66, 0x4a, 0xad, 0x93, 0xf7, 0x89, 0x47,
	0x72, 0x8c, 0x17, 0x1e, 0xc9, 0x31, 0x3e, 0x78, 0x24, 0xc7, 0x38, 0xe1, 0xb1, 0x1c, 0xc3, 0x85,
	0xc7, 0x72, 0x0c, 0x37, 0x1e, 0xcb, 0x31, 0x44, 0x19, 0xa6, 0x67, 0x96, 0x64, 0x94, 0x26, 0xe9,
	0x25, 0xe7, 0xe7, 0xea, 0xc3, 0xdd, 0x03, 0x61, 0x41, 0x0d, 0xac, 0x40, 0x18, 0x59, 0x52, 0x59,
	0x90, 0x5a, 0x9c, 0xc4, 0x06, 0xce, 0x26, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0xa7, 0x9e,
	0xfb, 0x29, 0xa5, 0x03, 0x00, 0x00,
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
	DocumentByProof(ctx context.Context, in *QueryDocumentByProofRequest, opts ...grpc.CallOption) (*QueryDocumentByProofResponse, error)
	DocumentByHolderId(ctx context.Context, in *QueryDocumentByHolderIdRequest, opts ...grpc.CallOption) (*QueryDocumentByHolderIdResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) DocumentByProof(ctx context.Context, in *QueryDocumentByProofRequest, opts ...grpc.CallOption) (*QueryDocumentByProofResponse, error) {
	out := new(QueryDocumentByProofResponse)
	err := c.cc.Invoke(ctx, "/ShareRing.shareledger.document.Query/DocumentByProof", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) DocumentByHolderId(ctx context.Context, in *QueryDocumentByHolderIdRequest, opts ...grpc.CallOption) (*QueryDocumentByHolderIdResponse, error) {
	out := new(QueryDocumentByHolderIdResponse)
	err := c.cc.Invoke(ctx, "/ShareRing.shareledger.document.Query/DocumentByHolderId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	DocumentByProof(context.Context, *QueryDocumentByProofRequest) (*QueryDocumentByProofResponse, error)
	DocumentByHolderId(context.Context, *QueryDocumentByHolderIdRequest) (*QueryDocumentByHolderIdResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) DocumentByProof(ctx context.Context, req *QueryDocumentByProofRequest) (*QueryDocumentByProofResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DocumentByProof not implemented")
}
func (*UnimplementedQueryServer) DocumentByHolderId(ctx context.Context, req *QueryDocumentByHolderIdRequest) (*QueryDocumentByHolderIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DocumentByHolderId not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_DocumentByProof_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryDocumentByProofRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).DocumentByProof(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ShareRing.shareledger.document.Query/DocumentByProof",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).DocumentByProof(ctx, req.(*QueryDocumentByProofRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_DocumentByHolderId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryDocumentByHolderIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).DocumentByHolderId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ShareRing.shareledger.document.Query/DocumentByHolderId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).DocumentByHolderId(ctx, req.(*QueryDocumentByHolderIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ShareRing.shareledger.document.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DocumentByProof",
			Handler:    _Query_DocumentByProof_Handler,
		},
		{
			MethodName: "DocumentByHolderId",
			Handler:    _Query_DocumentByHolderId_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "document/query.proto",
}

func (m *QueryDocumentByProofRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryDocumentByProofRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryDocumentByProofRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Proof) > 0 {
		i -= len(m.Proof)
		copy(dAtA[i:], m.Proof)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Proof)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryDocumentByProofResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryDocumentByProofResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryDocumentByProofResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Document != nil {
		{
			size, err := m.Document.MarshalToSizedBuffer(dAtA[:i])
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

func (m *QueryDocumentByHolderIdRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryDocumentByHolderIdRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryDocumentByHolderIdRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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

func (m *QueryDocumentByHolderIdResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryDocumentByHolderIdResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryDocumentByHolderIdResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Document) > 0 {
		for iNdEx := len(m.Document) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Document[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
func (m *QueryDocumentByProofRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Proof)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryDocumentByProofResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Document != nil {
		l = m.Document.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryDocumentByHolderIdRequest) Size() (n int) {
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

func (m *QueryDocumentByHolderIdResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Document) > 0 {
		for _, e := range m.Document {
			l = e.Size()
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryDocumentByProofRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryDocumentByProofRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryDocumentByProofRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Proof", wireType)
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
			m.Proof = string(dAtA[iNdEx:postIndex])
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
func (m *QueryDocumentByProofResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryDocumentByProofResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryDocumentByProofResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Document", wireType)
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
			if m.Document == nil {
				m.Document = &Document{}
			}
			if err := m.Document.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *QueryDocumentByHolderIdRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryDocumentByHolderIdRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryDocumentByHolderIdRequest: illegal tag %d (wire type %d)", fieldNum, wire)
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
func (m *QueryDocumentByHolderIdResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryDocumentByHolderIdResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryDocumentByHolderIdResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Document", wireType)
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
			m.Document = append(m.Document, &Document{})
			if err := m.Document[len(m.Document)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
