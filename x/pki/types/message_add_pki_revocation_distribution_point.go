package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/zigbee-alliance/distributed-compliance-ledger/utils/validator"
)

const TypeMsgAddPkiRevocationDistributionPoint = "add_pki_revocation_distribution_point"

var _ sdk.Msg = &MsgAddPkiRevocationDistributionPoint{}

func NewMsgAddPkiRevocationDistributionPoint(signer string, vid int32, pid int32, isPAA bool, label string, crlSignerCertificate string, issuerSubjectKeyID string, dataUrl string, dataFileSize uint64, dataDigest string, dataDigestType uint32, revocationType uint64) *MsgAddPkiRevocationDistributionPoint {
	return &MsgAddPkiRevocationDistributionPoint{
		Signer:               signer,
		Vid:                  vid,
		Pid:                  pid,
		IsPAA:                isPAA,
		Label:                label,
		CrlSignerCertificate: crlSignerCertificate,
		IssuerSubjectKeyID:   issuerSubjectKeyID,
		DataUrl:              dataUrl,
		DataFileSize:         dataFileSize,
		DataDigest:           dataDigest,
		DataDigestType:       dataDigestType,
		RevocationType:       revocationType,
	}
}

func (msg *MsgAddPkiRevocationDistributionPoint) Route() string {
	return RouterKey
}

func (msg *MsgAddPkiRevocationDistributionPoint) Type() string {
	return TypeMsgAddPkiRevocationDistributionPoint
}

func (msg *MsgAddPkiRevocationDistributionPoint) GetSigners() []sdk.AccAddress {
	signer, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{signer}
}

func (msg *MsgAddPkiRevocationDistributionPoint) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgAddPkiRevocationDistributionPoint) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid signer address (%s)", err)
	}

	err = validator.Validate(msg)
	if err != nil {
		return err
	}

	allowedDataDigestTypes := [6]uint32{1, 7, 8, 10, 11, 12}
	allowedRevocationType := uint64(1)

	isDataDigestInTypes := false
	for _, digestType := range allowedDataDigestTypes {
		if digestType == msg.DataDigestType {
			isDataDigestInTypes = true

			break
		}
	}

	if !isDataDigestInTypes {
		return DataDigestNotInTypes
	}

	if msg.RevocationType != allowedRevocationType {
		return RevocationNotInTypes
	}

	if msg.IsPAA {
		if msg.Pid == nil {
			return EmptyPid
		}

	}

	if msg.DataFileSize == nil && msg.DataDigest != nil {
		return EmptyDataFileSize
	}

	if msg.DataFileSize != nil && msg.DataDigest == nil {
		return EmptyDataDigest
	}

	if msg.DataDigest == nil && msg.DataDigestType != nil {
		return EmptyDataDigest
	}

	if msg.DataDigest != nil && msg.DataDigestType == nil {
		return emptyDataDigestType
	}

	if msg.RevocationType == 1 && (msg.DataFileSize != nil || msg.DataDigest != nil || msg.dataDigestType != nil) {

	}

	return nil
}
