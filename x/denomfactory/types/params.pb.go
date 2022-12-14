// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: g4alchain/denomfactory/params.proto

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

// Params defines the parameters for the module.
type Params struct {
	DenomSymbol      string `protobuf:"bytes,1,opt,name=denom_symbol,json=denomSymbol,proto3" json:"denom_symbol,omitempty" yaml:"genesis_denom_symbol"`
	DenomName        string `protobuf:"bytes,2,opt,name=denom_name,json=denomName,proto3" json:"denom_name,omitempty" yaml:"genesis_denom_name"`
	DenomDescription string `protobuf:"bytes,3,opt,name=denom_description,json=denomDescription,proto3" json:"denom_description,omitempty" yaml:"genesis_denom_description"`
}

func (m *Params) Reset()      { *m = Params{} }
func (*Params) ProtoMessage() {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_2e2b212449c9fc37, []int{0}
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

func (m *Params) GetDenomSymbol() string {
	if m != nil {
		return m.DenomSymbol
	}
	return ""
}

func (m *Params) GetDenomName() string {
	if m != nil {
		return m.DenomName
	}
	return ""
}

func (m *Params) GetDenomDescription() string {
	if m != nil {
		return m.DenomDescription
	}
	return ""
}

func init() {
	proto.RegisterType((*Params)(nil), "g4alentertainment.g4alchain.denomfactory.Params")
}

func init() {
	proto.RegisterFile("g4alchain/denomfactory/params.proto", fileDescriptor_2e2b212449c9fc37)
}

var fileDescriptor_2e2b212449c9fc37 = []byte{
	// 301 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x4e, 0x37, 0x49, 0xcc,
	0x49, 0xce, 0x48, 0xcc, 0xcc, 0xd3, 0x4f, 0x49, 0xcd, 0xcb, 0xcf, 0x4d, 0x4b, 0x4c, 0x2e, 0xc9,
	0x2f, 0xaa, 0xd4, 0x2f, 0x48, 0x2c, 0x4a, 0xcc, 0x2d, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17,
	0xd2, 0x00, 0x29, 0x4a, 0xcd, 0x2b, 0x49, 0x2d, 0x2a, 0x49, 0xcc, 0xcc, 0xcb, 0x4d, 0xcd, 0x2b,
	0xd1, 0x83, 0x6b, 0xd3, 0x43, 0xd6, 0x26, 0x25, 0x92, 0x9e, 0x9f, 0x9e, 0x0f, 0xd6, 0xa4, 0x0f,
	0x62, 0x41, 0xf4, 0x2b, 0x3d, 0x66, 0xe4, 0x62, 0x0b, 0x00, 0x1b, 0x28, 0xe4, 0xc4, 0xc5, 0x03,
	0xd6, 0x10, 0x5f, 0x5c, 0x99, 0x9b, 0x94, 0x9f, 0x23, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0xe9, 0x24,
	0xff, 0xe9, 0x9e, 0xbc, 0x74, 0x65, 0x62, 0x6e, 0x8e, 0x95, 0x52, 0x7a, 0x6a, 0x5e, 0x6a, 0x71,
	0x66, 0x71, 0x3c, 0xb2, 0x2a, 0xa5, 0x20, 0x6e, 0x30, 0x37, 0x18, 0xcc, 0x13, 0xb2, 0xe1, 0xe2,
	0x82, 0xc8, 0xe6, 0x25, 0xe6, 0xa6, 0x4a, 0x30, 0x81, 0x4d, 0x90, 0xfd, 0x74, 0x4f, 0x5e, 0x12,
	0x9b, 0x09, 0x20, 0x35, 0x4a, 0x41, 0x9c, 0x60, 0x8e, 0x5f, 0x62, 0x6e, 0xaa, 0x50, 0x20, 0x97,
	0x20, 0x44, 0x26, 0x25, 0xb5, 0x38, 0xb9, 0x28, 0xb3, 0xa0, 0x24, 0x33, 0x3f, 0x4f, 0x82, 0x19,
	0x6c, 0x88, 0xca, 0xa7, 0x7b, 0xf2, 0x0a, 0xd8, 0x0c, 0x41, 0x52, 0xaa, 0x14, 0x24, 0x00, 0x16,
	0x73, 0x41, 0x08, 0x59, 0xb1, 0xcc, 0x58, 0x20, 0xcf, 0xe0, 0x14, 0x7e, 0xe2, 0x91, 0x1c, 0xe3,
	0x85, 0x47, 0x72, 0x8c, 0x0f, 0x1e, 0xc9, 0x31, 0x4e, 0x78, 0x2c, 0xc7, 0x70, 0xe1, 0xb1, 0x1c,
	0xc3, 0x8d, 0xc7, 0x72, 0x0c, 0x51, 0xb6, 0xe9, 0x99, 0x25, 0x19, 0xa5, 0x49, 0x7a, 0xc9, 0xf9,
	0xb9, 0xfa, 0xee, 0x26, 0x8e, 0x3e, 0xba, 0xae, 0xc8, 0x61, 0xa9, 0x0f, 0x0a, 0x4b, 0x5d, 0x48,
	0x1c, 0x54, 0xa0, 0xc6, 0x42, 0x49, 0x65, 0x41, 0x6a, 0x71, 0x12, 0x1b, 0x38, 0x14, 0x8d, 0x01,
	0x01, 0x00, 0x00, 0xff, 0xff, 0x02, 0x2a, 0xaf, 0xe6, 0xac, 0x01, 0x00, 0x00,
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
	if len(m.DenomDescription) > 0 {
		i -= len(m.DenomDescription)
		copy(dAtA[i:], m.DenomDescription)
		i = encodeVarintParams(dAtA, i, uint64(len(m.DenomDescription)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.DenomName) > 0 {
		i -= len(m.DenomName)
		copy(dAtA[i:], m.DenomName)
		i = encodeVarintParams(dAtA, i, uint64(len(m.DenomName)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.DenomSymbol) > 0 {
		i -= len(m.DenomSymbol)
		copy(dAtA[i:], m.DenomSymbol)
		i = encodeVarintParams(dAtA, i, uint64(len(m.DenomSymbol)))
		i--
		dAtA[i] = 0xa
	}
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
	l = len(m.DenomSymbol)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	l = len(m.DenomName)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	l = len(m.DenomDescription)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
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
				return fmt.Errorf("proto: wrong wireType = %d for field DenomSymbol", wireType)
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
			m.DenomSymbol = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DenomName", wireType)
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
			m.DenomName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DenomDescription", wireType)
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
			m.DenomDescription = string(dAtA[iNdEx:postIndex])
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
