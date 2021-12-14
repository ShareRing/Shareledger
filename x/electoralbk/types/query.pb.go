// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: electoralbk/query.proto

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

type QueryVoterRequest struct {
	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
}

func (m *QueryVoterRequest) Reset()         { *m = QueryVoterRequest{} }
func (m *QueryVoterRequest) String() string { return proto.CompactTextString(m) }
func (*QueryVoterRequest) ProtoMessage()    {}
func (*QueryVoterRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_19740b5b1696327f, []int{0}
}
func (m *QueryVoterRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryVoterRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryVoterRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryVoterRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryVoterRequest.Merge(m, src)
}
func (m *QueryVoterRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryVoterRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryVoterRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryVoterRequest proto.InternalMessageInfo

type QueryVoterResponse struct {
	Voter *Voter `protobuf:"bytes,1,opt,name=voter,proto3" json:"voter,omitempty"`
}

func (m *QueryVoterResponse) Reset()         { *m = QueryVoterResponse{} }
func (m *QueryVoterResponse) String() string { return proto.CompactTextString(m) }
func (*QueryVoterResponse) ProtoMessage()    {}
func (*QueryVoterResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_19740b5b1696327f, []int{1}
}
func (m *QueryVoterResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryVoterResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryVoterResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryVoterResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryVoterResponse.Merge(m, src)
}
func (m *QueryVoterResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryVoterResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryVoterResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryVoterResponse proto.InternalMessageInfo

func (m *QueryVoterResponse) GetVoter() *Voter {
	if m != nil {
		return m.Voter
	}
	return nil
}

func init() {
	proto.RegisterType((*QueryVoterRequest)(nil), "ShareRing.shareledger.electoralbk.QueryVoterRequest")
	proto.RegisterType((*QueryVoterResponse)(nil), "ShareRing.shareledger.electoralbk.QueryVoterResponse")
}

func init() { proto.RegisterFile("electoralbk/query.proto", fileDescriptor_19740b5b1696327f) }

var fileDescriptor_19740b5b1696327f = []byte{
	// 310 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4f, 0xcd, 0x49, 0x4d,
	0x2e, 0xc9, 0x2f, 0x4a, 0xcc, 0x49, 0xca, 0xd6, 0x2f, 0x2c, 0x4d, 0x2d, 0xaa, 0xd4, 0x2b, 0x28,
	0xca, 0x2f, 0xc9, 0x17, 0x52, 0x0c, 0xce, 0x48, 0x2c, 0x4a, 0x0d, 0xca, 0xcc, 0x4b, 0xd7, 0x2b,
	0x06, 0xb1, 0x72, 0x52, 0x53, 0xd2, 0x53, 0x8b, 0xf4, 0x90, 0x94, 0x4b, 0xc9, 0xa4, 0xe7, 0xe7,
	0xa7, 0xe7, 0xa4, 0xea, 0x27, 0x16, 0x64, 0xea, 0x27, 0xe6, 0xe5, 0xe5, 0x97, 0x24, 0x96, 0x64,
	0xe6, 0xe7, 0x15, 0x43, 0x0c, 0x90, 0x12, 0x49, 0xcf, 0x4f, 0xcf, 0x07, 0x33, 0xf5, 0x41, 0x2c,
	0x98, 0x28, 0xb2, 0x7d, 0x25, 0x15, 0x10, 0x51, 0x25, 0x73, 0x2e, 0xc1, 0x40, 0x90, 0xdd, 0x61,
	0xf9, 0x25, 0xa9, 0x45, 0x41, 0xa9, 0x85, 0xa5, 0xa9, 0xc5, 0x25, 0x42, 0x12, 0x5c, 0xec, 0x89,
	0x29, 0x29, 0x45, 0xa9, 0xc5, 0xc5, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x30, 0xae, 0x15,
	0x47, 0xc7, 0x02, 0x79, 0x86, 0x17, 0x0b, 0xe4, 0x19, 0x94, 0x42, 0xb8, 0x84, 0x90, 0x35, 0x16,
	0x17, 0xe4, 0xe7, 0x15, 0xa7, 0x0a, 0xd9, 0x71, 0xb1, 0x96, 0x81, 0x04, 0xc0, 0xfa, 0xb8, 0x8d,
	0x34, 0xf4, 0x08, 0xfa, 0x45, 0x0f, 0x62, 0x00, 0x44, 0x9b, 0xd1, 0x12, 0x46, 0x2e, 0x56, 0xb0,
	0xb1, 0x42, 0xb3, 0x18, 0xb9, 0x58, 0xc1, 0x52, 0x42, 0x26, 0x44, 0x18, 0x82, 0xe1, 0x07, 0x29,
	0x53, 0x12, 0x75, 0x41, 0x3c, 0xa0, 0xa4, 0xd2, 0x74, 0xf9, 0xc9, 0x64, 0x26, 0x39, 0x21, 0x19,
	0x7d, 0x24, 0x4d, 0xfa, 0x60, 0xc7, 0xe9, 0x57, 0x43, 0x43, 0xa1, 0xd6, 0xc9, 0xef, 0xc4, 0x23,
	0x39, 0xc6, 0x0b, 0x8f, 0xe4, 0x18, 0x1f, 0x3c, 0x92, 0x63, 0x9c, 0xf0, 0x58, 0x8e, 0xe1, 0xc2,
	0x63, 0x39, 0x86, 0x1b, 0x8f, 0xe5, 0x18, 0xa2, 0x4c, 0xd2, 0x33, 0x4b, 0x32, 0x4a, 0x93, 0xf4,
	0x92, 0xf3, 0x73, 0xf5, 0xe1, 0x0e, 0x80, 0xb0, 0xa0, 0x66, 0x55, 0xe8, 0xa3, 0x44, 0x44, 0x65,
	0x41, 0x6a, 0x71, 0x12, 0x1b, 0x38, 0x32, 0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0xe4, 0x84,
	0x8c, 0x2e, 0x14, 0x02, 0x00, 0x00,
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
	// Balance queries the balance of a single coin for a single account.
	Voter(ctx context.Context, in *QueryVoterRequest, opts ...grpc.CallOption) (*QueryVoterResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Voter(ctx context.Context, in *QueryVoterRequest, opts ...grpc.CallOption) (*QueryVoterResponse, error) {
	out := new(QueryVoterResponse)
	err := c.cc.Invoke(ctx, "/ShareRing.shareledger.electoralbk.Query/Voter", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// Balance queries the balance of a single coin for a single account.
	Voter(context.Context, *QueryVoterRequest) (*QueryVoterResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) Voter(ctx context.Context, req *QueryVoterRequest) (*QueryVoterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Voter not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_Voter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryVoterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Voter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ShareRing.shareledger.electoralbk.Query/Voter",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Voter(ctx, req.(*QueryVoterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ShareRing.shareledger.electoralbk.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Voter",
			Handler:    _Query_Voter_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "electoralbk/query.proto",
}

func (m *QueryVoterRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryVoterRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryVoterRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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

func (m *QueryVoterResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryVoterResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryVoterResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Voter != nil {
		{
			size, err := m.Voter.MarshalToSizedBuffer(dAtA[:i])
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
func (m *QueryVoterRequest) Size() (n int) {
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

func (m *QueryVoterResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Voter != nil {
		l = m.Voter.Size()
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
func (m *QueryVoterRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryVoterRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryVoterRequest: illegal tag %d (wire type %d)", fieldNum, wire)
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
func (m *QueryVoterResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryVoterResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryVoterResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Voter", wireType)
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
			if m.Voter == nil {
				m.Voter = &Voter{}
			}
			if err := m.Voter.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
