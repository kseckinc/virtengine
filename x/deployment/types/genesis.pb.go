// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: virtengine/deployment/v1beta1/genesis.proto

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

// GenesisDeployment defines the basic genesis state used by deployment module
type GenesisDeployment struct {
	Deployment Deployment `protobuf:"bytes,1,opt,name=deployment,proto3" json:"deployment" yaml:"deployment"`
	Groups     []Group    `protobuf:"bytes,2,rep,name=groups,proto3" json:"groups" yaml:"groups"`
}

func (m *GenesisDeployment) Reset()         { *m = GenesisDeployment{} }
func (m *GenesisDeployment) String() string { return proto.CompactTextString(m) }
func (*GenesisDeployment) ProtoMessage()    {}
func (*GenesisDeployment) Descriptor() ([]byte, []int) {
	return fileDescriptor_cbb2cbae4c442342, []int{0}
}
func (m *GenesisDeployment) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisDeployment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisDeployment.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisDeployment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisDeployment.Merge(m, src)
}
func (m *GenesisDeployment) XXX_Size() int {
	return m.Size()
}
func (m *GenesisDeployment) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisDeployment.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisDeployment proto.InternalMessageInfo

func (m *GenesisDeployment) GetDeployment() Deployment {
	if m != nil {
		return m.Deployment
	}
	return Deployment{}
}

func (m *GenesisDeployment) GetGroups() []Group {
	if m != nil {
		return m.Groups
	}
	return nil
}

// GenesisState stores slice of genesis deployment instance
type GenesisState struct {
	Deployments []GenesisDeployment `protobuf:"bytes,1,rep,name=deployments,proto3" json:"deployments" yaml:"deployments"`
	Params      Params              `protobuf:"bytes,2,opt,name=params,proto3" json:"params" yaml:"params"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_cbb2cbae4c442342, []int{1}
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

func (m *GenesisState) GetDeployments() []GenesisDeployment {
	if m != nil {
		return m.Deployments
	}
	return nil
}

func (m *GenesisState) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

func init() {
	proto.RegisterType((*GenesisDeployment)(nil), "virtengine.deployment.v1beta1.GenesisDeployment")
	proto.RegisterType((*GenesisState)(nil), "virtengine.deployment.v1beta1.GenesisState")
}

func init() {
	proto.RegisterFile("virtengine/deployment/v1beta1/genesis.proto", fileDescriptor_cbb2cbae4c442342)
}

var fileDescriptor_cbb2cbae4c442342 = []byte{
	// 354 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0x41, 0x4b, 0xc3, 0x30,
	0x14, 0xc7, 0x9b, 0x09, 0x3b, 0x64, 0x7a, 0x58, 0xf1, 0x50, 0x06, 0xb6, 0x23, 0x28, 0x6e, 0x0a,
	0xa9, 0xdb, 0x6e, 0x1e, 0x8b, 0xb0, 0x93, 0x20, 0xf5, 0x26, 0x28, 0x74, 0x1a, 0x6a, 0x61, 0x6d,
	0x4a, 0x93, 0x0d, 0x77, 0xf0, 0x3b, 0xf8, 0xb1, 0x76, 0xdc, 0xd1, 0x53, 0x91, 0x0d, 0x3c, 0x08,
	0x5e, 0xf6, 0x09, 0x64, 0x49, 0x30, 0x61, 0x87, 0xf6, 0x96, 0x34, 0xbf, 0xff, 0xfb, 0xbf, 0xff,
	0xeb, 0x83, 0x97, 0xf3, 0xa4, 0xe0, 0x24, 0x8b, 0x93, 0x8c, 0xf8, 0x2f, 0x24, 0x9f, 0xd2, 0x45,
	0x4a, 0x32, 0xee, 0xcf, 0x07, 0x13, 0xc2, 0xa3, 0x81, 0x1f, 0x93, 0x8c, 0xb0, 0x84, 0xe1, 0xbc,
	0xa0, 0x9c, 0xda, 0x27, 0x1a, 0xc6, 0x1a, 0xc6, 0x0a, 0xee, 0x1c, 0xc7, 0x34, 0xa6, 0x82, 0xf4,
	0x77, 0x27, 0x29, 0xea, 0xe0, 0x6a, 0x07, 0xa3, 0x8e, 0xe4, 0xfb, 0x35, 0x1d, 0x15, 0x74, 0x96,
	0x2b, 0xf4, 0xa2, 0x1a, 0xcd, 0xa3, 0x22, 0x4a, 0x55, 0xef, 0xe8, 0x1b, 0xc0, 0xf6, 0x58, 0xa6,
	0xb9, 0xf9, 0x47, 0xed, 0x02, 0x42, 0x2d, 0x74, 0x40, 0x17, 0xf4, 0x5a, 0xc3, 0x3e, 0xae, 0x8c,
	0x89, 0xb5, 0x3c, 0x38, 0x5f, 0x96, 0x9e, 0xf5, 0x53, 0x7a, 0x46, 0x91, 0x6d, 0xe9, 0xb5, 0x17,
	0x51, 0x3a, 0xbd, 0x46, 0xfa, 0x1b, 0x0a, 0x0d, 0xc0, 0x7e, 0x84, 0x4d, 0x11, 0x82, 0x39, 0x8d,
	0xee, 0x41, 0xaf, 0x35, 0x3c, 0xad, 0xf1, 0x1b, 0xef, 0xe0, 0xc0, 0x53, 0x56, 0x4a, 0xbb, 0x2d,
	0xbd, 0x23, 0x69, 0x23, 0xef, 0x28, 0x54, 0x0f, 0xe8, 0x17, 0xc0, 0x43, 0x15, 0xf4, 0x9e, 0x47,
	0x9c, 0xd8, 0xef, 0xb0, 0xa5, 0xab, 0x32, 0x07, 0x08, 0xd3, 0xab, 0x3a, 0xd3, 0xfd, 0x51, 0x05,
	0x7d, 0xd5, 0x80, 0x59, 0x6c, 0x5b, 0x7a, 0xf6, 0x7e, 0x58, 0x86, 0x42, 0x13, 0xb1, 0x9f, 0x60,
	0x53, 0xfe, 0x08, 0xa7, 0x21, 0xc6, 0x7b, 0x56, 0xe3, 0x7c, 0x27, 0x60, 0x9d, 0x57, 0x8a, 0x75,
	0x5e, 0x79, 0x47, 0xa1, 0x7a, 0x08, 0x6e, 0x97, 0x6b, 0x17, 0xac, 0xd6, 0x2e, 0xf8, 0x5a, 0xbb,
	0xe0, 0x63, 0xe3, 0x5a, 0xab, 0x8d, 0x6b, 0x7d, 0x6e, 0x5c, 0xeb, 0x61, 0x14, 0x27, 0xfc, 0x75,
	0x36, 0xc1, 0xcf, 0x34, 0xf5, 0x8d, 0x4d, 0x31, 0x8e, 0x6f, 0xe6, 0xda, 0xf0, 0x45, 0x4e, 0xd8,
	0xa4, 0x29, 0xd6, 0x65, 0xf4, 0x17, 0x00, 0x00, 0xff, 0xff, 0x1b, 0x5c, 0xda, 0x44, 0x19, 0x03,
	0x00, 0x00,
}

func (m *GenesisDeployment) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisDeployment) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisDeployment) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Groups) > 0 {
		for iNdEx := len(m.Groups) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Groups[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
	{
		size, err := m.Deployment.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
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
	dAtA[i] = 0x12
	if len(m.Deployments) > 0 {
		for iNdEx := len(m.Deployments) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Deployments[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
func (m *GenesisDeployment) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Deployment.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if len(m.Groups) > 0 {
		for _, e := range m.Groups {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Deployments) > 0 {
		for _, e := range m.Deployments {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GenesisDeployment) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: GenesisDeployment: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisDeployment: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Deployment", wireType)
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
			if err := m.Deployment.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Groups", wireType)
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
			m.Groups = append(m.Groups, Group{})
			if err := m.Groups[len(m.Groups)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
				return fmt.Errorf("proto: wrong wireType = %d for field Deployments", wireType)
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
			m.Deployments = append(m.Deployments, GenesisDeployment{})
			if err := m.Deployments[len(m.Deployments)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
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