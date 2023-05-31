package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	pkitypes "github.com/zigbee-alliance/distributed-compliance-ledger/types/pki"
	dclauthtypes "github.com/zigbee-alliance/distributed-compliance-ledger/x/dclauth/types"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/types"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/x509"
)

func (k msgServer) UpdatePkiRevocationDistributionPoint(goCtx context.Context, msg *types.MsgUpdatePkiRevocationDistributionPoint) (*types.MsgUpdatePkiRevocationDistributionPointResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	signerAddr, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "Invalid Address: (%s)", err)
	}

	signerAccount, _ := k.dclauthKeeper.GetAccountO(ctx, signerAddr)

	// check if signer has vendor role
	if !k.dclauthKeeper.HasRole(ctx, signerAddr, dclauthtypes.Vendor) {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrUnauthorized,
			"MsgUpdatePkiRevocationDistributionPoint transaction should be signed by an account with the \"%s\" role",
			dclauthtypes.Vendor,
		)
	}

	pkiRevocationDistributionPoint, isFound := k.GetPkiRevocationDistributionPoint(ctx, msg.Vid, msg.Label, msg.IssuerSubjectKeyID)
	if !isFound {
		return nil, pkitypes.NewErrPkiRevocationDistributionPointDoesNotExists("PKI revocation distribution point does not exist")
	}

	crlSignerCertificate, err := x509.DecodeX509Certificate(pkiRevocationDistributionPoint.CrlSignerCertificate)
	if err != nil {
		return nil, pkitypes.NewErrInvalidCertificate(err)
	}

	if pkiRevocationDistributionPoint.Vid != signerAccount.VendorID {
		return nil, sdkerrors.Wrap(pkitypes.ErrCRLSignerCertificateVidNotEqualAccountVid,
			"MsgUpdatePkiRevocationDistributionPoint signer must have the same vid as provided in an existing certificate from the revocation point",
		)
	}

	if msg.CrlSignerCertificate != "" {
		updatedCrlSignerCertificate, err := x509.DecodeX509Certificate(msg.CrlSignerCertificate)
		if err != nil {
			return nil, pkitypes.NewErrInvalidCertificate(err)
		}

		if crlSignerCertificate.IsSelfSigned() {
			if !updatedCrlSignerCertificate.IsSelfSigned() {
				return nil, pkitypes.NewErrRootCertificateIsNotSelfSigned("Updated CRL signer certificate must be self-signed since old one was self-signed")
			}

			vid, err := x509.GetVidFromSubject(updatedCrlSignerCertificate.SubjectAsText)

			if err != nil {
				return nil, sdkerrors.Wrapf(pkitypes.ErrInvalidVidFormat, "Could not parse vid: %s", err)
			}

			if vid == 0 {
				return nil, pkitypes.NewErrUnsupportedOperation("publishing a revocation point for non-VID scoped root certificates is currently not supported")
			}

			if vid != pkiRevocationDistributionPoint.Vid {
				return nil, pkitypes.NewErrCRLSignerCertificateVidNotEqualMsgVid("CRL Signer Certificate's vid must be equal to the provided vid in the message")
			}
		} else {
			if updatedCrlSignerCertificate.IsSelfSigned() {
				return nil, pkitypes.NewErrNonRootCertificateSelfSigned("Updated CRL signer certificate must not be self-signed since old one was not self-signed")
			}

			vid, err := x509.GetVidFromSubject(updatedCrlSignerCertificate.SubjectAsText)

			if err != nil {
				return nil, sdkerrors.Wrapf(pkitypes.ErrInvalidVidFormat, "Could not parse vid: %s", err)
			}

			if vid == 0 {
				return nil, pkitypes.NewErrVidNotFound("vid must be present in updated non-root CRL signer certificate")
			}

			if vid != msg.Vid {
				return nil, pkitypes.NewErrCRLSignerCertificateVidNotEqualMsgVid("CRL Signer Certificate's vid must be equal to the provided vid in the message")
			}

			pid, err := x509.GetPidFromSubject(updatedCrlSignerCertificate.SubjectAsText)

			if err != nil {
				return nil, sdkerrors.Wrapf(pkitypes.ErrInvalidPidFormat, "Could not parse pid: %s", err)
			}

			if pid != 0 && pid != pkiRevocationDistributionPoint.Pid {
				return nil, pkitypes.NewErrCRLSignerCertificateVidNotEqualMsgVid("PEM value of CRL signer certificate and certificate found by its Subject and SubjectKeyID on the ledger is not equal")
			}
		}
	}
	if msg.CrlSignerCertificate != "" {
		pkiRevocationDistributionPoint.CrlSignerCertificate = msg.CrlSignerCertificate
	}

	if msg.DataUrl != "" {
		pkiRevocationDistributionPoint.DataUrl = msg.DataUrl
	}

	if msg.DataFileSize != 0 {
		pkiRevocationDistributionPoint.DataFileSize = msg.DataFileSize
	}

	if msg.DataDigest != "" {
		pkiRevocationDistributionPoint.DataDigest = msg.DataDigest
	}

	if msg.DataDigestType != 0 {
		pkiRevocationDistributionPoint.DataDigestType = msg.DataDigestType
	}

	k.SetPkiRevocationDistributionPoint(ctx, pkiRevocationDistributionPoint)

	return &types.MsgUpdatePkiRevocationDistributionPointResponse{}, nil
}
