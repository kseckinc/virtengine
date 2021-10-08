// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: virtengine/base/v1beta1/endpoint.proto

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

// This describes how the endpoint is implemented when the lease is deployed
type Endpoint_Kind int32

const (
	// Describes an endpoint that becomes a Kubernetes Ingress
	Endpoint_SHARED_HTTP Endpoint_Kind = 0
	// Describes an endpoint that becomes a Kubernetes NodePort
	Endpoint_RANDOM_PORT Endpoint_Kind = 1
)

var Endpoint_Kind_name = map[int32]string{
	0: "SHARED_HTTP",
	1: "RANDOM_PORT",
}

var Endpoint_Kind_value = map[string]int32{
	"SHARED_HTTP": 0,
	"RANDOM_PORT": 1,
}

func (x Endpoint_Kind) String() string {
	return proto.EnumName(Endpoint_Kind_name, int32(x))
}

func (Endpoint_Kind) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_ddf973b5d8483608, []int{0, 0}
}

// Endpoint describes a publicly accessible IP service
type Endpoint struct {
	Kind Endpoint_Kind `protobuf:"varint,1,opt,name=kind,proto3,enum=virtengine.base.v1beta1.Endpoint_Kind" json:"kind,omitempty"`
}

func (m *Endpoint) Reset()         { *m = Endpoint{} }
func (m *Endpoint) String() string { return proto.CompactTextString(m) }
func (*Endpoint) ProtoMessage()    {}
func (*Endpoint) Descriptor() ([]byte, []int) {
	return fileDescriptor_ddf973b5d8483608, []int{0}
}
func (m *Endpoint) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Endpoint) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Endpoint.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Endpoint) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Endpoint.Merge(m, src)
}
func (m *Endpoint) XXX_Size() int {
	return m.Size()
}
func (m *Endpoint) XXX_DiscardUnknown() {
	xxx_messageInfo_Endpoint.DiscardUnknown(m)
}

var xxx_messageInfo_Endpoint proto.InternalMessageInfo

func (m *Endpoint) GetKind() Endpoint_Kind {
	if m != nil {
		return m.Kind
	}
	return Endpoint_SHARED_HTTP
}

func init() {
	proto.RegisterEnum("virtengine.base.v1beta1.Endpoint_Kind", Endpoint_Kind_name, Endpoint_Kind_value)
	proto.RegisterType((*Endpoint)(nil), "virtengine.base.v1beta1.Endpoint")
}

func init() {
	proto.RegisterFile("virtengine/base/v1beta1/endpoint.proto", fileDescriptor_ddf973b5d8483608)
}

var fileDescriptor_ddf973b5d8483608 = []byte{
	// 223 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x2b, 0xcb, 0x2c, 0x2a,
	0x49, 0xcd, 0x4b, 0xcf, 0xcc, 0x4b, 0xd5, 0x4f, 0x4a, 0x2c, 0x4e, 0xd5, 0x2f, 0x33, 0x4c, 0x4a,
	0x2d, 0x49, 0x34, 0xd4, 0x4f, 0xcd, 0x4b, 0x29, 0xc8, 0xcf, 0xcc, 0x2b, 0xd1, 0x2b, 0x28, 0xca,
	0x2f, 0xc9, 0x17, 0x12, 0x47, 0xa8, 0xd3, 0x03, 0xa9, 0xd3, 0x83, 0xaa, 0x93, 0x12, 0x49, 0xcf,
	0x4f, 0xcf, 0x07, 0xab, 0xd1, 0x07, 0xb1, 0x20, 0xca, 0x95, 0xca, 0xb8, 0x38, 0x5c, 0xa1, 0x06,
	0x08, 0x59, 0x71, 0xb1, 0x64, 0x67, 0xe6, 0xa5, 0x48, 0x30, 0x2a, 0x30, 0x6a, 0xf0, 0x19, 0xa9,
	0xe9, 0xe1, 0x30, 0x49, 0x0f, 0xa6, 0x41, 0xcf, 0x3b, 0x33, 0x2f, 0x25, 0x08, 0xac, 0x47, 0x49,
	0x83, 0x8b, 0x05, 0xc4, 0x13, 0xe2, 0xe7, 0xe2, 0x0e, 0xf6, 0x70, 0x0c, 0x72, 0x75, 0x89, 0xf7,
	0x08, 0x09, 0x09, 0x10, 0x60, 0x00, 0x09, 0x04, 0x39, 0xfa, 0xb9, 0xf8, 0xfb, 0xc6, 0x07, 0xf8,
	0x07, 0x85, 0x08, 0x30, 0x5a, 0xb1, 0xbc, 0x58, 0x20, 0xcf, 0xe8, 0xe4, 0x70, 0xe2, 0x91, 0x1c,
	0xe3, 0x85, 0x47, 0x72, 0x8c, 0x0f, 0x1e, 0xc9, 0x31, 0x4e, 0x78, 0x2c, 0xc7, 0x70, 0xe1, 0xb1,
	0x1c, 0xc3, 0x8d, 0xc7, 0x72, 0x0c, 0x51, 0x6a, 0xe9, 0x99, 0x25, 0x19, 0xa5, 0x49, 0x7a, 0xc9,
	0xf9, 0xb9, 0xfa, 0x48, 0x7e, 0x46, 0x62, 0x96, 0x54, 0x16, 0xa4, 0x16, 0x27, 0xb1, 0x81, 0x3d,
	0x60, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x8a, 0x1d, 0x37, 0x34, 0x19, 0x01, 0x00, 0x00,
}

func (this *Endpoint) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Endpoint)
	if !ok {
		that2, ok := that.(Endpoint)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Kind != that1.Kind {
		return false
	}
	return true
}
func (m *Endpoint) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Endpoint) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Endpoint) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Kind != 0 {
		i = encodeVarintEndpoint(dAtA, i, uint64(m.Kind))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintEndpoint(dAtA []byte, offset int, v uint64) int {
	offset -= sovEndpoint(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Endpoint) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Kind != 0 {
		n += 1 + sovEndpoint(uint64(m.Kind))
	}
	return n
}

func sovEndpoint(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozEndpoint(x uint64) (n int) {
	return sovEndpoint(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Endpoint) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEndpoint
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
			return fmt.Errorf("proto: Endpoint: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Endpoint: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Kind", wireType)
			}
			m.Kind = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEndpoint
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Kind |= Endpoint_Kind(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipEndpoint(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEndpoint
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
func skipEndpoint(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowEndpoint
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
					return 0, ErrIntOverflowEndpoint
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
					return 0, ErrIntOverflowEndpoint
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
				return 0, ErrInvalidLengthEndpoint
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupEndpoint
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthEndpoint
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthEndpoint        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowEndpoint          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupEndpoint = fmt.Errorf("proto: unexpected end of group")
)
