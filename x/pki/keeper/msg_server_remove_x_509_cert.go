package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	pkitypes "github.com/zigbee-alliance/distributed-compliance-ledger/types/pki"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/types"
)

func (k msgServer) RemoveX509Cert(goCtx context.Context, msg *types.MsgRemoveX509Cert) (*types.MsgRemoveX509CertResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	aprCerts, foundApproved := k.GetApprovedCertificates(ctx, msg.Subject, msg.SubjectKeyId)
	revCerts, foundRevoked := k.GetRevokedCertificates(ctx, msg.Subject, msg.SubjectKeyId)
	if !foundApproved && !foundRevoked {
		return nil, pkitypes.NewErrCertificateDoesNotExist(msg.Subject, msg.SubjectKeyId)
	}

	var certificates []*types.Certificate
	if foundApproved {
		certificates = aprCerts.Certs
	} else {
		certificates = revCerts.Certs
	}

	if len(certificates) == 0 {
		return nil, pkitypes.NewErrCertificateDoesNotExist(msg.Subject, msg.SubjectKeyId)
	}

	if certificates[0].IsRoot {
		return nil, pkitypes.NewErrMessageRemoveRoot(msg.Subject, msg.SubjectKeyId)
	}

	if msg.Signer != certificates[0].Owner {
		return nil, pkitypes.NewErrMessageOnlyOwnerCanExecute("REVOKE_X509_CERT")
	}

	certID := types.CertificateIdentifier{
		Subject:      msg.Subject,
		SubjectKeyId: msg.SubjectKeyId,
	}

	if msg.SerialNumber != "" {
		certBySerialNumber, found := findCertificate(msg.SerialNumber, &certificates)
		if !found {
			return nil, pkitypes.NewErrCertificateBySerialNumberDoesNotExist(msg.Subject, msg.SubjectKeyId, msg.SerialNumber)
		}

		// remove from subject with serialNumber map
		k.RemoveUniqueCertificate(ctx, certBySerialNumber.Issuer, certBySerialNumber.SerialNumber)

		certs := types.ApprovedCertificates{
			Subject:      msg.Subject,
			SubjectKeyId: msg.SubjectKeyId,
			Certs:        certificates,
		}
		k.removeCertFromList(certBySerialNumber.Issuer, certBySerialNumber.SerialNumber, &certs)
		k._removeX509Cert(ctx, certID, certs, foundRevoked)
	} else {
		k.RemoveApprovedCertificates(ctx, certID.Subject, certID.SubjectKeyId)
		// remove from subject -> subject key ID map
		k.RemoveApprovedCertificateBySubject(ctx, certID.Subject, certID.SubjectKeyId)
		// remove from subject key ID -> certificates map
		k.RemoveApprovedCertificatesBySubjectKeyID(ctx, certID.Subject, certID.SubjectKeyId)
		// remove from revoked list
		k.RemoveRevokedCertificates(ctx, certID.Subject, certID.SubjectKeyId)
		// remove from subject with serialNumber map
		for _, cert := range certificates {
			k.RemoveUniqueCertificate(ctx, cert.Issuer, cert.SerialNumber)
		}
	}

	return &types.MsgRemoveX509CertResponse{}, nil
}

func (k msgServer) _removeX509Cert(ctx sdk.Context, certID types.CertificateIdentifier, certificates types.ApprovedCertificates, isRevoked bool) {
	if isRevoked { //nolint:nestif
		if len(certificates.Certs) == 0 {
			k.RemoveRevokedCertificates(ctx, certID.Subject, certID.SubjectKeyId)
		} else {
			k.SetRevokedCertificates(
				ctx,
				types.RevokedCertificates{
					Subject:      certID.Subject,
					SubjectKeyId: certID.SubjectKeyId,
					Certs:        certificates.Certs,
				},
			)
		}
	} else {
		if len(certificates.Certs) == 0 {
			k.RemoveApprovedCertificates(ctx, certID.Subject, certID.SubjectKeyId)
			k.RemoveApprovedCertificateBySubject(ctx, certID.Subject, certID.SubjectKeyId)
			k.RemoveApprovedCertificatesBySubjectKeyID(ctx, certID.Subject, certID.SubjectKeyId)
		} else {
			k.SetApprovedCertificates(ctx, certificates)
			k.SetApprovedCertificatesBySubjectKeyID(
				ctx,
				types.ApprovedCertificatesBySubjectKeyId{SubjectKeyId: certID.SubjectKeyId, Certs: certificates.Certs},
			)
		}
	}
}
