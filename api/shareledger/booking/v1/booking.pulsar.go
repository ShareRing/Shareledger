// Code generated by protoc-gen-go-pulsar. DO NOT EDIT.
package v1

import (
	fmt "fmt"
	runtime "github.com/cosmos/cosmos-proto/runtime"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoiface "google.golang.org/protobuf/runtime/protoiface"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	io "io"
	reflect "reflect"
	sync "sync"
)

var (
	md_Booking             protoreflect.MessageDescriptor
	fd_Booking_bookID      protoreflect.FieldDescriptor
	fd_Booking_booker      protoreflect.FieldDescriptor
	fd_Booking_UUID        protoreflect.FieldDescriptor
	fd_Booking_duration    protoreflect.FieldDescriptor
	fd_Booking_isCompleted protoreflect.FieldDescriptor
)

func init() {
	file_shareledger_booking_v1_booking_proto_init()
	md_Booking = File_shareledger_booking_v1_booking_proto.Messages().ByName("Booking")
	fd_Booking_bookID = md_Booking.Fields().ByName("bookID")
	fd_Booking_booker = md_Booking.Fields().ByName("booker")
	fd_Booking_UUID = md_Booking.Fields().ByName("UUID")
	fd_Booking_duration = md_Booking.Fields().ByName("duration")
	fd_Booking_isCompleted = md_Booking.Fields().ByName("isCompleted")
}

var _ protoreflect.Message = (*fastReflection_Booking)(nil)

type fastReflection_Booking Booking

func (x *Booking) ProtoReflect() protoreflect.Message {
	return (*fastReflection_Booking)(x)
}

func (x *Booking) slowProtoReflect() protoreflect.Message {
	mi := &file_shareledger_booking_v1_booking_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_Booking_messageType fastReflection_Booking_messageType
var _ protoreflect.MessageType = fastReflection_Booking_messageType{}

type fastReflection_Booking_messageType struct{}

func (x fastReflection_Booking_messageType) Zero() protoreflect.Message {
	return (*fastReflection_Booking)(nil)
}
func (x fastReflection_Booking_messageType) New() protoreflect.Message {
	return new(fastReflection_Booking)
}
func (x fastReflection_Booking_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_Booking
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_Booking) Descriptor() protoreflect.MessageDescriptor {
	return md_Booking
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_Booking) Type() protoreflect.MessageType {
	return _fastReflection_Booking_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_Booking) New() protoreflect.Message {
	return new(fastReflection_Booking)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_Booking) Interface() protoreflect.ProtoMessage {
	return (*Booking)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_Booking) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.BookID != "" {
		value := protoreflect.ValueOfString(x.BookID)
		if !f(fd_Booking_bookID, value) {
			return
		}
	}
	if x.Booker != "" {
		value := protoreflect.ValueOfString(x.Booker)
		if !f(fd_Booking_booker, value) {
			return
		}
	}
	if x.UUID != "" {
		value := protoreflect.ValueOfString(x.UUID)
		if !f(fd_Booking_UUID, value) {
			return
		}
	}
	if x.Duration != int64(0) {
		value := protoreflect.ValueOfInt64(x.Duration)
		if !f(fd_Booking_duration, value) {
			return
		}
	}
	if x.IsCompleted != false {
		value := protoreflect.ValueOfBool(x.IsCompleted)
		if !f(fd_Booking_isCompleted, value) {
			return
		}
	}
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_Booking) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "shareledger.booking.Booking.bookID":
		return x.BookID != ""
	case "shareledger.booking.Booking.booker":
		return x.Booker != ""
	case "shareledger.booking.Booking.UUID":
		return x.UUID != ""
	case "shareledger.booking.Booking.duration":
		return x.Duration != int64(0)
	case "shareledger.booking.Booking.isCompleted":
		return x.IsCompleted != false
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: shareledger.booking.Booking"))
		}
		panic(fmt.Errorf("message shareledger.booking.Booking does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Booking) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "shareledger.booking.Booking.bookID":
		x.BookID = ""
	case "shareledger.booking.Booking.booker":
		x.Booker = ""
	case "shareledger.booking.Booking.UUID":
		x.UUID = ""
	case "shareledger.booking.Booking.duration":
		x.Duration = int64(0)
	case "shareledger.booking.Booking.isCompleted":
		x.IsCompleted = false
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: shareledger.booking.Booking"))
		}
		panic(fmt.Errorf("message shareledger.booking.Booking does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_Booking) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "shareledger.booking.Booking.bookID":
		value := x.BookID
		return protoreflect.ValueOfString(value)
	case "shareledger.booking.Booking.booker":
		value := x.Booker
		return protoreflect.ValueOfString(value)
	case "shareledger.booking.Booking.UUID":
		value := x.UUID
		return protoreflect.ValueOfString(value)
	case "shareledger.booking.Booking.duration":
		value := x.Duration
		return protoreflect.ValueOfInt64(value)
	case "shareledger.booking.Booking.isCompleted":
		value := x.IsCompleted
		return protoreflect.ValueOfBool(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: shareledger.booking.Booking"))
		}
		panic(fmt.Errorf("message shareledger.booking.Booking does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Booking) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "shareledger.booking.Booking.bookID":
		x.BookID = value.Interface().(string)
	case "shareledger.booking.Booking.booker":
		x.Booker = value.Interface().(string)
	case "shareledger.booking.Booking.UUID":
		x.UUID = value.Interface().(string)
	case "shareledger.booking.Booking.duration":
		x.Duration = value.Int()
	case "shareledger.booking.Booking.isCompleted":
		x.IsCompleted = value.Bool()
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: shareledger.booking.Booking"))
		}
		panic(fmt.Errorf("message shareledger.booking.Booking does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Booking) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "shareledger.booking.Booking.bookID":
		panic(fmt.Errorf("field bookID of message shareledger.booking.Booking is not mutable"))
	case "shareledger.booking.Booking.booker":
		panic(fmt.Errorf("field booker of message shareledger.booking.Booking is not mutable"))
	case "shareledger.booking.Booking.UUID":
		panic(fmt.Errorf("field UUID of message shareledger.booking.Booking is not mutable"))
	case "shareledger.booking.Booking.duration":
		panic(fmt.Errorf("field duration of message shareledger.booking.Booking is not mutable"))
	case "shareledger.booking.Booking.isCompleted":
		panic(fmt.Errorf("field isCompleted of message shareledger.booking.Booking is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: shareledger.booking.Booking"))
		}
		panic(fmt.Errorf("message shareledger.booking.Booking does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_Booking) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "shareledger.booking.Booking.bookID":
		return protoreflect.ValueOfString("")
	case "shareledger.booking.Booking.booker":
		return protoreflect.ValueOfString("")
	case "shareledger.booking.Booking.UUID":
		return protoreflect.ValueOfString("")
	case "shareledger.booking.Booking.duration":
		return protoreflect.ValueOfInt64(int64(0))
	case "shareledger.booking.Booking.isCompleted":
		return protoreflect.ValueOfBool(false)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: shareledger.booking.Booking"))
		}
		panic(fmt.Errorf("message shareledger.booking.Booking does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_Booking) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in shareledger.booking.Booking", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_Booking) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Booking) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_Booking) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_Booking) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*Booking)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		l = len(x.BookID)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.Booker)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.UUID)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.Duration != 0 {
			n += 1 + runtime.Sov(uint64(x.Duration))
		}
		if x.IsCompleted {
			n += 2
		}
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*Booking)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if x.IsCompleted {
			i--
			if x.IsCompleted {
				dAtA[i] = 1
			} else {
				dAtA[i] = 0
			}
			i--
			dAtA[i] = 0x28
		}
		if x.Duration != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.Duration))
			i--
			dAtA[i] = 0x20
		}
		if len(x.UUID) > 0 {
			i -= len(x.UUID)
			copy(dAtA[i:], x.UUID)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.UUID)))
			i--
			dAtA[i] = 0x1a
		}
		if len(x.Booker) > 0 {
			i -= len(x.Booker)
			copy(dAtA[i:], x.Booker)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Booker)))
			i--
			dAtA[i] = 0x12
		}
		if len(x.BookID) > 0 {
			i -= len(x.BookID)
			copy(dAtA[i:], x.BookID)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.BookID)))
			i--
			dAtA[i] = 0xa
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*Booking)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: Booking: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: Booking: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field BookID", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
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
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.BookID = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Booker", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
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
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Booker = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 3:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field UUID", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
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
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.UUID = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 4:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Duration", wireType)
				}
				x.Duration = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.Duration |= int64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			case 5:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field IsCompleted", wireType)
				}
				var v int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				x.IsCompleted = bool(v != 0)
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.0
// 	protoc        (unknown)
// source: shareledger/booking/v1/booking.proto

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Booking struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BookID      string `protobuf:"bytes,1,opt,name=bookID,proto3" json:"bookID,omitempty"`
	Booker      string `protobuf:"bytes,2,opt,name=booker,proto3" json:"booker,omitempty"`
	UUID        string `protobuf:"bytes,3,opt,name=UUID,proto3" json:"UUID,omitempty"`
	Duration    int64  `protobuf:"varint,4,opt,name=duration,proto3" json:"duration,omitempty"`
	IsCompleted bool   `protobuf:"varint,5,opt,name=isCompleted,proto3" json:"isCompleted,omitempty"`
}

func (x *Booking) Reset() {
	*x = Booking{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shareledger_booking_v1_booking_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Booking) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Booking) ProtoMessage() {}

// Deprecated: Use Booking.ProtoReflect.Descriptor instead.
func (*Booking) Descriptor() ([]byte, []int) {
	return file_shareledger_booking_v1_booking_proto_rawDescGZIP(), []int{0}
}

func (x *Booking) GetBookID() string {
	if x != nil {
		return x.BookID
	}
	return ""
}

func (x *Booking) GetBooker() string {
	if x != nil {
		return x.Booker
	}
	return ""
}

func (x *Booking) GetUUID() string {
	if x != nil {
		return x.UUID
	}
	return ""
}

func (x *Booking) GetDuration() int64 {
	if x != nil {
		return x.Duration
	}
	return 0
}

func (x *Booking) GetIsCompleted() bool {
	if x != nil {
		return x.IsCompleted
	}
	return false
}

var File_shareledger_booking_v1_booking_proto protoreflect.FileDescriptor

var file_shareledger_booking_v1_booking_proto_rawDesc = []byte{
	0x0a, 0x24, 0x73, 0x68, 0x61, 0x72, 0x65, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2f, 0x62, 0x6f,
	0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31, 0x2f, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x73, 0x68, 0x61, 0x72, 0x65, 0x6c, 0x65, 0x64,
	0x67, 0x65, 0x72, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x22, 0x8b, 0x01, 0x0a, 0x07,
	0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x12, 0x16, 0x0a, 0x06, 0x62, 0x6f, 0x6f, 0x6b, 0x49,
	0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x62, 0x6f, 0x6f, 0x6b, 0x49, 0x44, 0x12,
	0x16, 0x0a, 0x06, 0x62, 0x6f, 0x6f, 0x6b, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x62, 0x6f, 0x6f, 0x6b, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x55, 0x55, 0x49, 0x44, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x55, 0x55, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08, 0x64,
	0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x64,
	0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x20, 0x0a, 0x0b, 0x69, 0x73, 0x43, 0x6f, 0x6d,
	0x70, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x69, 0x73,
	0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x42, 0xbd, 0x01, 0x0a, 0x17, 0x63, 0x6f,
	0x6d, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x65, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2e, 0x62, 0x6f,
	0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x42, 0x0c, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x27, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x73, 0x64, 0x6b,
	0x2e, 0x69, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x6c, 0x65, 0x64,
	0x67, 0x65, 0x72, 0x2f, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31, 0xa2, 0x02,
	0x03, 0x53, 0x42, 0x58, 0xaa, 0x02, 0x13, 0x53, 0x68, 0x61, 0x72, 0x65, 0x6c, 0x65, 0x64, 0x67,
	0x65, 0x72, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0xca, 0x02, 0x13, 0x53, 0x68, 0x61,
	0x72, 0x65, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x5c, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67,
	0xe2, 0x02, 0x1f, 0x53, 0x68, 0x61, 0x72, 0x65, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x5c, 0x42,
	0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0xea, 0x02, 0x14, 0x53, 0x68, 0x61, 0x72, 0x65, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72,
	0x3a, 0x3a, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_shareledger_booking_v1_booking_proto_rawDescOnce sync.Once
	file_shareledger_booking_v1_booking_proto_rawDescData = file_shareledger_booking_v1_booking_proto_rawDesc
)

func file_shareledger_booking_v1_booking_proto_rawDescGZIP() []byte {
	file_shareledger_booking_v1_booking_proto_rawDescOnce.Do(func() {
		file_shareledger_booking_v1_booking_proto_rawDescData = protoimpl.X.CompressGZIP(file_shareledger_booking_v1_booking_proto_rawDescData)
	})
	return file_shareledger_booking_v1_booking_proto_rawDescData
}

var file_shareledger_booking_v1_booking_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_shareledger_booking_v1_booking_proto_goTypes = []interface{}{
	(*Booking)(nil), // 0: shareledger.booking.Booking
}
var file_shareledger_booking_v1_booking_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_shareledger_booking_v1_booking_proto_init() }
func file_shareledger_booking_v1_booking_proto_init() {
	if File_shareledger_booking_v1_booking_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_shareledger_booking_v1_booking_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Booking); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_shareledger_booking_v1_booking_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_shareledger_booking_v1_booking_proto_goTypes,
		DependencyIndexes: file_shareledger_booking_v1_booking_proto_depIdxs,
		MessageInfos:      file_shareledger_booking_v1_booking_proto_msgTypes,
	}.Build()
	File_shareledger_booking_v1_booking_proto = out.File
	file_shareledger_booking_v1_booking_proto_rawDesc = nil
	file_shareledger_booking_v1_booking_proto_goTypes = nil
	file_shareledger_booking_v1_booking_proto_depIdxs = nil
}
