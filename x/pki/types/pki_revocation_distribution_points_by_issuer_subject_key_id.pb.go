// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: zigbeealliance/distributedcomplianceledger/pki/pki_revocation_distribution_points_by_issuer_subject_key_id.proto

package types

import (
	fmt "fmt"
	proto "github.com/cosmos/gogoproto/proto"
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

type PkiRevocationDistributionPointsByIssuerSubjectKeyID struct {
	IssuerSubjectKeyID string                            `protobuf:"bytes,1,opt,name=issuerSubjectKeyID,proto3" json:"issuerSubjectKeyID,omitempty"`
	Points             []*PkiRevocationDistributionPoint `protobuf:"bytes,2,rep,name=points,proto3" json:"points,omitempty"`
	SchemaVersion      uint32                            `protobuf:"varint,3,opt,name=schemaVersion,proto3" json:"schemaVersion,omitempty"`
}

func (m *PkiRevocationDistributionPointsByIssuerSubjectKeyID) Reset() {
	*m = PkiRevocationDistributionPointsByIssuerSubjectKeyID{}
}
func (m *PkiRevocationDistributionPointsByIssuerSubjectKeyID) String() string {
	return proto.CompactTextString(m)
}
func (*PkiRevocationDistributionPointsByIssuerSubjectKeyID) ProtoMessage() {}
func (*PkiRevocationDistributionPointsByIssuerSubjectKeyID) Descriptor() ([]byte, []int) {
	return fileDescriptor_0b1940ba04139c1b, []int{0}
}
func (m *PkiRevocationDistributionPointsByIssuerSubjectKeyID) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PkiRevocationDistributionPointsByIssuerSubjectKeyID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PkiRevocationDistributionPointsByIssuerSubjectKeyID.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PkiRevocationDistributionPointsByIssuerSubjectKeyID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PkiRevocationDistributionPointsByIssuerSubjectKeyID.Merge(m, src)
}
func (m *PkiRevocationDistributionPointsByIssuerSubjectKeyID) XXX_Size() int {
	return m.Size()
}
func (m *PkiRevocationDistributionPointsByIssuerSubjectKeyID) XXX_DiscardUnknown() {
	xxx_messageInfo_PkiRevocationDistributionPointsByIssuerSubjectKeyID.DiscardUnknown(m)
}

var xxx_messageInfo_PkiRevocationDistributionPointsByIssuerSubjectKeyID proto.InternalMessageInfo

func (m *PkiRevocationDistributionPointsByIssuerSubjectKeyID) GetIssuerSubjectKeyID() string {
	if m != nil {
		return m.IssuerSubjectKeyID
	}
	return ""
}

func (m *PkiRevocationDistributionPointsByIssuerSubjectKeyID) GetPoints() []*PkiRevocationDistributionPoint {
	if m != nil {
		return m.Points
	}
	return nil
}

func (m *PkiRevocationDistributionPointsByIssuerSubjectKeyID) GetSchemaVersion() uint32 {
	if m != nil {
		return m.SchemaVersion
	}
	return 0
}

func init() {
	proto.RegisterType((*PkiRevocationDistributionPointsByIssuerSubjectKeyID)(nil), "zigbeealliance.distributedcomplianceledger.pki.PkiRevocationDistributionPointsByIssuerSubjectKeyID")
}

func init() {
	proto.RegisterFile("zigbeealliance/distributedcomplianceledger/pki/pki_revocation_distribution_points_by_issuer_subject_key_id.proto", fileDescriptor_0b1940ba04139c1b)
}

var fileDescriptor_0b1940ba04139c1b = []byte{
	// 304 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x90, 0x4f, 0x4a, 0x03, 0x31,
	0x14, 0xc6, 0x1b, 0x0b, 0x05, 0x47, 0xba, 0x99, 0x55, 0x71, 0x11, 0x8a, 0xb8, 0xe8, 0xa6, 0x29,
	0xd8, 0x1b, 0x94, 0x6e, 0x8a, 0x20, 0x65, 0x84, 0x2e, 0x5c, 0x18, 0x26, 0x99, 0x67, 0xfb, 0x9c,
	0x3f, 0x09, 0x49, 0x46, 0x1c, 0x4f, 0xe1, 0xb1, 0x5c, 0x76, 0xe9, 0x52, 0x66, 0x8e, 0xe0, 0x05,
	0xa4, 0x33, 0xa5, 0x5a, 0x2c, 0x05, 0xc1, 0x45, 0x16, 0x79, 0xf9, 0xc8, 0xfb, 0x7d, 0x3f, 0x4f,
	0xbf, 0xe0, 0x52, 0x00, 0x84, 0x49, 0x82, 0x61, 0x26, 0x61, 0x14, 0xa1, 0x75, 0x06, 0x45, 0xee,
	0x20, 0x92, 0x2a, 0xd5, 0xcd, 0x34, 0x81, 0x68, 0x09, 0x66, 0xa4, 0x63, 0xdc, 0x1c, 0x6e, 0xe0,
	0x49, 0xc9, 0xd0, 0xa1, 0xca, 0xf8, 0x2e, 0xbe, 0xb9, 0x68, 0x85, 0x99, 0xb3, 0x5c, 0x14, 0x1c,
	0xad, 0xcd, 0xc1, 0x70, 0x9b, 0x8b, 0x47, 0x90, 0x8e, 0xc7, 0x50, 0x70, 0x8c, 0x98, 0x36, 0xca,
	0x29, 0x9f, 0xed, 0x6f, 0x64, 0x47, 0x36, 0x32, 0x1d, 0xe3, 0xf9, 0xe2, 0xbf, 0x09, 0x1b, 0x8e,
	0x8b, 0x4f, 0xe2, 0x8d, 0xe7, 0x31, 0x06, 0xbb, 0xe8, 0xf4, 0x47, 0x72, 0x5e, 0x57, 0x99, 0x14,
	0xb3, 0xba, 0xc8, 0x6d, 0xd3, 0xe3, 0x1a, 0x8a, 0xd9, 0xd4, 0x67, 0x9e, 0x8f, 0xbf, 0xa6, 0x3d,
	0xd2, 0x27, 0x83, 0xd3, 0xe0, 0xc0, 0x8b, 0xff, 0xe0, 0x75, 0x1a, 0x31, 0xbd, 0x93, 0x7e, 0x7b,
	0x70, 0x76, 0x75, 0xf3, 0x47, 0x01, 0xec, 0x38, 0x64, 0xb0, 0xfd, 0xdd, 0xbf, 0xf4, 0xba, 0x56,
	0xae, 0x20, 0x0d, 0x17, 0x60, 0x2c, 0xaa, 0xac, 0xd7, 0xee, 0x93, 0x41, 0x37, 0xd8, 0x1f, 0x4e,
	0xee, 0xdf, 0x4a, 0x4a, 0xd6, 0x25, 0x25, 0x1f, 0x25, 0x25, 0xaf, 0x15, 0x6d, 0xad, 0x2b, 0xda,
	0x7a, 0xaf, 0x68, 0xeb, 0x6e, 0xba, 0x44, 0xb7, 0xca, 0x05, 0x93, 0x2a, 0x1d, 0x35, 0x84, 0xc3,
	0x43, 0xce, 0x87, 0xdf, 0x8c, 0xc3, 0xad, 0xf5, 0xe7, 0xda, 0xbb, 0x2b, 0x34, 0x58, 0xd1, 0xa9,
	0xe5, 0x8e, 0xbf, 0x02, 0x00, 0x00, 0xff, 0xff, 0x18, 0x8d, 0x09, 0xb4, 0x58, 0x02, 0x00, 0x00,
}

func (m *PkiRevocationDistributionPointsByIssuerSubjectKeyID) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PkiRevocationDistributionPointsByIssuerSubjectKeyID) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PkiRevocationDistributionPointsByIssuerSubjectKeyID) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.SchemaVersion != 0 {
		i = encodeVarintPkiRevocationDistributionPointsByIssuerSubjectKeyId(dAtA, i, uint64(m.SchemaVersion))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Points) > 0 {
		for iNdEx := len(m.Points) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Points[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintPkiRevocationDistributionPointsByIssuerSubjectKeyId(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.IssuerSubjectKeyID) > 0 {
		i -= len(m.IssuerSubjectKeyID)
		copy(dAtA[i:], m.IssuerSubjectKeyID)
		i = encodeVarintPkiRevocationDistributionPointsByIssuerSubjectKeyId(dAtA, i, uint64(len(m.IssuerSubjectKeyID)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintPkiRevocationDistributionPointsByIssuerSubjectKeyId(dAtA []byte, offset int, v uint64) int {
	offset -= sovPkiRevocationDistributionPointsByIssuerSubjectKeyId(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *PkiRevocationDistributionPointsByIssuerSubjectKeyID) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.IssuerSubjectKeyID)
	if l > 0 {
		n += 1 + l + sovPkiRevocationDistributionPointsByIssuerSubjectKeyId(uint64(l))
	}
	if len(m.Points) > 0 {
		for _, e := range m.Points {
			l = e.Size()
			n += 1 + l + sovPkiRevocationDistributionPointsByIssuerSubjectKeyId(uint64(l))
		}
	}
	if m.SchemaVersion != 0 {
		n += 1 + sovPkiRevocationDistributionPointsByIssuerSubjectKeyId(uint64(m.SchemaVersion))
	}
	return n
}

func sovPkiRevocationDistributionPointsByIssuerSubjectKeyId(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozPkiRevocationDistributionPointsByIssuerSubjectKeyId(x uint64) (n int) {
	return sovPkiRevocationDistributionPointsByIssuerSubjectKeyId(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *PkiRevocationDistributionPointsByIssuerSubjectKeyID) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPkiRevocationDistributionPointsByIssuerSubjectKeyId
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
			return fmt.Errorf("proto: PkiRevocationDistributionPointsByIssuerSubjectKeyID: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PkiRevocationDistributionPointsByIssuerSubjectKeyID: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IssuerSubjectKeyID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPkiRevocationDistributionPointsByIssuerSubjectKeyId
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
				return ErrInvalidLengthPkiRevocationDistributionPointsByIssuerSubjectKeyId
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPkiRevocationDistributionPointsByIssuerSubjectKeyId
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.IssuerSubjectKeyID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Points", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPkiRevocationDistributionPointsByIssuerSubjectKeyId
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
				return ErrInvalidLengthPkiRevocationDistributionPointsByIssuerSubjectKeyId
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPkiRevocationDistributionPointsByIssuerSubjectKeyId
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Points = append(m.Points, &PkiRevocationDistributionPoint{})
			if err := m.Points[len(m.Points)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SchemaVersion", wireType)
			}
			m.SchemaVersion = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPkiRevocationDistributionPointsByIssuerSubjectKeyId
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SchemaVersion |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipPkiRevocationDistributionPointsByIssuerSubjectKeyId(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPkiRevocationDistributionPointsByIssuerSubjectKeyId
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
func skipPkiRevocationDistributionPointsByIssuerSubjectKeyId(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPkiRevocationDistributionPointsByIssuerSubjectKeyId
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
					return 0, ErrIntOverflowPkiRevocationDistributionPointsByIssuerSubjectKeyId
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
					return 0, ErrIntOverflowPkiRevocationDistributionPointsByIssuerSubjectKeyId
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
				return 0, ErrInvalidLengthPkiRevocationDistributionPointsByIssuerSubjectKeyId
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupPkiRevocationDistributionPointsByIssuerSubjectKeyId
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthPkiRevocationDistributionPointsByIssuerSubjectKeyId
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthPkiRevocationDistributionPointsByIssuerSubjectKeyId        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPkiRevocationDistributionPointsByIssuerSubjectKeyId          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupPkiRevocationDistributionPointsByIssuerSubjectKeyId = fmt.Errorf("proto: unexpected end of group")
)
