// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: zigbeealliance/distributedcomplianceledger/dclauth/revoked_account.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
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

type RevokedAccount_Reason int32

const (
	RevokedAccount_TrusteeVoting      RevokedAccount_Reason = 0
	RevokedAccount_MaliciousValidator RevokedAccount_Reason = 1
)

var RevokedAccount_Reason_name = map[int32]string{
	0: "TrusteeVoting",
	1: "MaliciousValidator",
}

var RevokedAccount_Reason_value = map[string]int32{
	"TrusteeVoting":      0,
	"MaliciousValidator": 1,
}

func (x RevokedAccount_Reason) String() string {
	return proto.EnumName(RevokedAccount_Reason_name, int32(x))
}

func (RevokedAccount_Reason) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_0da4e0e6766e981b, []int{0, 0}
}

type RevokedAccount struct {
	*Account                    `protobuf:"bytes,1,opt,name=account,proto3,embedded=account" json:"account,omitempty"`
	RevokeApprovals             []*Grant              `protobuf:"bytes,2,rep,name=revokeApprovals,proto3" json:"revokeApprovals,omitempty"`
	Reason                      RevokedAccount_Reason `protobuf:"varint,3,opt,name=reason,proto3,enum=zigbeealliance.distributedcomplianceledger.dclauth.RevokedAccount_Reason" json:"reason,omitempty"`
	RevokedAccountSchemaVersion uint32                `protobuf:"varint,4,opt,name=revokedAccountSchemaVersion,proto3" json:"revokedAccountSchemaVersion,omitempty"`
}

func (m *RevokedAccount) Reset()         { *m = RevokedAccount{} }
func (m *RevokedAccount) String() string { return proto.CompactTextString(m) }
func (*RevokedAccount) ProtoMessage()    {}
func (*RevokedAccount) Descriptor() ([]byte, []int) {
	return fileDescriptor_0da4e0e6766e981b, []int{0}
}
func (m *RevokedAccount) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RevokedAccount) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RevokedAccount.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RevokedAccount) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RevokedAccount.Merge(m, src)
}
func (m *RevokedAccount) XXX_Size() int {
	return m.Size()
}
func (m *RevokedAccount) XXX_DiscardUnknown() {
	xxx_messageInfo_RevokedAccount.DiscardUnknown(m)
}

var xxx_messageInfo_RevokedAccount proto.InternalMessageInfo

func (m *RevokedAccount) GetRevokeApprovals() []*Grant {
	if m != nil {
		return m.RevokeApprovals
	}
	return nil
}

func (m *RevokedAccount) GetReason() RevokedAccount_Reason {
	if m != nil {
		return m.Reason
	}
	return RevokedAccount_TrusteeVoting
}

func (m *RevokedAccount) GetRevokedAccountSchemaVersion() uint32 {
	if m != nil {
		return m.RevokedAccountSchemaVersion
	}
	return 0
}

func init() {
	proto.RegisterEnum("zigbeealliance.distributedcomplianceledger.dclauth.RevokedAccount_Reason", RevokedAccount_Reason_name, RevokedAccount_Reason_value)
	proto.RegisterType((*RevokedAccount)(nil), "zigbeealliance.distributedcomplianceledger.dclauth.RevokedAccount")
}

func init() {
	proto.RegisterFile("zigbeealliance/distributedcomplianceledger/dclauth/revoked_account.proto", fileDescriptor_0da4e0e6766e981b)
}

var fileDescriptor_0da4e0e6766e981b = []byte{
	// 374 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0xd2, 0xc1, 0x4a, 0xe3, 0x40,
	0x18, 0x07, 0xf0, 0xcc, 0xb6, 0x74, 0x61, 0x4a, 0xbb, 0xdd, 0x61, 0x59, 0x42, 0x85, 0x18, 0x7a,
	0xca, 0xa5, 0x09, 0xa4, 0x27, 0x11, 0xa4, 0xed, 0xc5, 0x7a, 0xf0, 0x12, 0xa5, 0x07, 0x3d, 0xc8,
	0x64, 0x32, 0xa4, 0x83, 0x69, 0xbe, 0x30, 0x99, 0x14, 0xf5, 0x29, 0x7c, 0x10, 0x1f, 0xc4, 0x63,
	0x8f, 0x9e, 0x44, 0xda, 0x17, 0x91, 0x36, 0xa9, 0x12, 0x91, 0x82, 0xb9, 0x85, 0x30, 0xf3, 0x9b,
	0xef, 0xff, 0xf1, 0xc7, 0x93, 0x07, 0x11, 0xfa, 0x9c, 0xd3, 0x28, 0x12, 0x34, 0x66, 0xdc, 0x09,
	0x44, 0xaa, 0xa4, 0xf0, 0x33, 0xc5, 0x03, 0x06, 0xf3, 0x24, 0xff, 0x1b, 0xf1, 0x20, 0xe4, 0xd2,
	0x09, 0x58, 0x44, 0x33, 0x35, 0x73, 0x24, 0x5f, 0xc0, 0x2d, 0x0f, 0x6e, 0x28, 0x63, 0x90, 0xc5,
	0xca, 0x4e, 0x24, 0x28, 0x20, 0x6e, 0x59, 0xb2, 0xf7, 0x48, 0x76, 0x21, 0x75, 0xff, 0x85, 0x10,
	0xc2, 0xf6, 0xba, 0xb3, 0xf9, 0xca, 0xa5, 0xee, 0xb0, 0xc2, 0x4c, 0xa5, 0x59, 0xba, 0x27, 0x15,
	0x84, 0x50, 0xd2, 0xdd, 0xfd, 0xde, 0x53, 0x0d, 0xb7, 0xbd, 0x3c, 0xe5, 0x28, 0x87, 0xc9, 0x35,
	0xfe, 0x5d, 0xbc, 0xa1, 0x23, 0x13, 0x59, 0x4d, 0xf7, 0xd8, 0xfe, 0x79, 0x60, 0xbb, 0xd0, 0xc6,
	0xf5, 0xe5, 0xeb, 0x21, 0xf2, 0x76, 0x22, 0x61, 0xf8, 0x4f, 0xbe, 0xd4, 0x51, 0x92, 0x48, 0x58,
	0xd0, 0x28, 0xd5, 0x7f, 0x99, 0x35, 0xab, 0xe9, 0x1e, 0x55, 0x79, 0xe4, 0x74, 0x93, 0xc4, 0xfb,
	0x2a, 0x12, 0x8a, 0x1b, 0x92, 0xd3, 0x14, 0x62, 0xbd, 0x66, 0x22, 0xab, 0xed, 0x9e, 0x55, 0xb1,
	0xcb, 0x5b, 0xb1, 0xbd, 0x2d, 0xe8, 0x15, 0x30, 0x19, 0xe2, 0x03, 0x59, 0x3a, 0x70, 0xc1, 0x66,
	0x7c, 0x4e, 0xa7, 0x5c, 0xa6, 0x02, 0x62, 0xbd, 0x6e, 0x22, 0xab, 0xe5, 0xed, 0x3b, 0xd2, 0x1b,
	0xe0, 0x46, 0x6e, 0x92, 0xbf, 0xb8, 0x75, 0x29, 0xb3, 0x54, 0x71, 0x3e, 0x05, 0x25, 0xe2, 0xb0,
	0xa3, 0x91, 0xff, 0x98, 0x9c, 0xd3, 0x48, 0x30, 0x01, 0x59, 0x3a, 0xa5, 0x91, 0x08, 0xa8, 0x02,
	0xd9, 0x41, 0x63, 0xff, 0x79, 0x65, 0xa0, 0xe5, 0xca, 0x40, 0x6f, 0x2b, 0x03, 0x3d, 0xae, 0x0d,
	0x6d, 0xb9, 0x36, 0xb4, 0x97, 0xb5, 0xa1, 0x5d, 0x4d, 0x42, 0xa1, 0x66, 0x99, 0x6f, 0x33, 0x98,
	0x3b, 0x79, 0xda, 0xfe, 0x77, 0xa5, 0xe8, 0x7f, 0xe6, 0xed, 0x17, 0xb5, 0xb8, 0xfb, 0x28, 0x86,
	0xba, 0x4f, 0x78, 0xea, 0x37, 0xb6, 0xcd, 0x18, 0xbc, 0x07, 0x00, 0x00, 0xff, 0xff, 0xd5, 0x03,
	0xd3, 0x47, 0x31, 0x03, 0x00, 0x00,
}

func (m *RevokedAccount) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RevokedAccount) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RevokedAccount) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.RevokedAccountSchemaVersion != 0 {
		i = encodeVarintRevokedAccount(dAtA, i, uint64(m.RevokedAccountSchemaVersion))
		i--
		dAtA[i] = 0x20
	}
	if m.Reason != 0 {
		i = encodeVarintRevokedAccount(dAtA, i, uint64(m.Reason))
		i--
		dAtA[i] = 0x18
	}
	if len(m.RevokeApprovals) > 0 {
		for iNdEx := len(m.RevokeApprovals) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.RevokeApprovals[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintRevokedAccount(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if m.Account != nil {
		{
			size, err := m.Account.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintRevokedAccount(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintRevokedAccount(dAtA []byte, offset int, v uint64) int {
	offset -= sovRevokedAccount(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *RevokedAccount) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Account != nil {
		l = m.Account.Size()
		n += 1 + l + sovRevokedAccount(uint64(l))
	}
	if len(m.RevokeApprovals) > 0 {
		for _, e := range m.RevokeApprovals {
			l = e.Size()
			n += 1 + l + sovRevokedAccount(uint64(l))
		}
	}
	if m.Reason != 0 {
		n += 1 + sovRevokedAccount(uint64(m.Reason))
	}
	if m.RevokedAccountSchemaVersion != 0 {
		n += 1 + sovRevokedAccount(uint64(m.RevokedAccountSchemaVersion))
	}
	return n
}

func sovRevokedAccount(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozRevokedAccount(x uint64) (n int) {
	return sovRevokedAccount(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *RevokedAccount) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRevokedAccount
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
			return fmt.Errorf("proto: RevokedAccount: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RevokedAccount: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Account", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRevokedAccount
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
				return ErrInvalidLengthRevokedAccount
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthRevokedAccount
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Account == nil {
				m.Account = &Account{}
			}
			if err := m.Account.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RevokeApprovals", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRevokedAccount
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
				return ErrInvalidLengthRevokedAccount
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthRevokedAccount
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RevokeApprovals = append(m.RevokeApprovals, &Grant{})
			if err := m.RevokeApprovals[len(m.RevokeApprovals)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Reason", wireType)
			}
			m.Reason = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRevokedAccount
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Reason |= RevokedAccount_Reason(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RevokedAccountSchemaVersion", wireType)
			}
			m.RevokedAccountSchemaVersion = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRevokedAccount
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RevokedAccountSchemaVersion |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipRevokedAccount(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRevokedAccount
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
func skipRevokedAccount(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowRevokedAccount
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
					return 0, ErrIntOverflowRevokedAccount
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
					return 0, ErrIntOverflowRevokedAccount
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
				return 0, ErrInvalidLengthRevokedAccount
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupRevokedAccount
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthRevokedAccount
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthRevokedAccount        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowRevokedAccount          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupRevokedAccount = fmt.Errorf("proto: unexpected end of group")
)
