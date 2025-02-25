package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"

	pkitypes "github.com/zigbee-alliance/distributed-compliance-ledger/types/pki"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/types"
)

// SetNocRootCertificatesByVidAndSkid set a specific nocRootCertificatesByVidAndSkid in the store from its index.
func (k Keeper) SetNocRootCertificatesByVidAndSkid(ctx sdk.Context, nocRootCertificatesByVidAndSkid types.NocRootCertificatesByVidAndSkid) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), pkitypes.KeyPrefix(types.NocRootCertificatesByVidAndSkidKeyPrefix))
	b := k.cdc.MustMarshal(&nocRootCertificatesByVidAndSkid)
	store.Set(types.NocRootCertificatesByVidAndSkidKey(
		nocRootCertificatesByVidAndSkid.Vid,
		nocRootCertificatesByVidAndSkid.SubjectKeyId,
	), b)
}

// GetNocRootCertificatesByVidAndSkid returns a nocRootCertificatesByVidAndSkid from its index.
func (k Keeper) GetNocRootCertificatesByVidAndSkid(
	ctx sdk.Context,
	vid int32,
	subjectKeyID string,

) (val types.NocRootCertificatesByVidAndSkid, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), pkitypes.KeyPrefix(types.NocRootCertificatesByVidAndSkidKeyPrefix))

	b := store.Get(types.NocRootCertificatesByVidAndSkidKey(
		vid,
		subjectKeyID,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)

	return val, true
}

// Add a NOC root certificate to the list of NOC root certificates for the VID map.
func (k Keeper) AddNocRootCertificatesByVidAndSkid(ctx sdk.Context, nocCertificate types.Certificate) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), pkitypes.KeyPrefix(types.NocRootCertificatesByVidAndSkidKeyPrefix))

	nocRootCertificatesByVidAndSkidKeyBytes := store.Get(types.NocRootCertificatesByVidAndSkidKey(nocCertificate.Vid, nocCertificate.SubjectKeyId))
	var nocRootCertificatesByVidAndSkid types.NocRootCertificatesByVidAndSkid

	if nocRootCertificatesByVidAndSkidKeyBytes == nil {
		nocRootCertificatesByVidAndSkid = types.NocRootCertificatesByVidAndSkid{
			Vid:          nocCertificate.Vid,
			SubjectKeyId: nocCertificate.SubjectKeyId,
			Certs:        []*types.Certificate{},
			Tq:           1,
		}
	} else {
		k.cdc.MustUnmarshal(nocRootCertificatesByVidAndSkidKeyBytes, &nocRootCertificatesByVidAndSkid)
	}

	nocRootCertificatesByVidAndSkid.Certs = append(nocRootCertificatesByVidAndSkid.Certs, &nocCertificate)

	b := k.cdc.MustMarshal(&nocRootCertificatesByVidAndSkid)
	store.Set(types.NocRootCertificatesByVidAndSkidKey(nocCertificate.Vid, nocCertificate.SubjectKeyId), b)
}

// RemoveNocRootCertificatesByVidAndSkid removes a nocRootCertificatesByVidAndSkid from the store.
func (k Keeper) RemoveNocRootCertificatesByVidAndSkid(
	ctx sdk.Context,
	vid int32,
	subjectKeyID string,
) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), pkitypes.KeyPrefix(types.NocRootCertificatesByVidAndSkidKeyPrefix))
	store.Delete(types.NocRootCertificatesByVidAndSkidKey(
		vid,
		subjectKeyID,
	))
}

// RemoveNocRootCertificateByVidSkidSubjectAndSerialNumber removes root certificate with specified subject from the list.
func (k Keeper) RemoveNocRootCertificateByVidSubjectAndSkid(
	ctx sdk.Context,
	vid int32,
	subject string,
	subjectKeyID string,
) {
	k._filterAndSetNocRootCertificateByVidAndSkid(
		ctx,
		vid,
		subjectKeyID,
		func(cert *types.Certificate) bool {
			return cert.Subject != subject
		},
	)
}

// RemoveNocRootCertificateByVidSkidSubjectAndSerialNumber removes root certificate with specified subject and serial number from the list.
func (k Keeper) RemoveNocRootCertificateByVidSubjectSkidAndSerialNumber(
	ctx sdk.Context,
	vid int32,
	subject string,
	subjectKeyID string,
	serialNumber string,
) {
	k._filterAndSetNocRootCertificateByVidAndSkid(
		ctx,
		vid,
		subjectKeyID,
		func(cert *types.Certificate) bool {
			return !(cert.Subject == subject && cert.SerialNumber == serialNumber)
		},
	)
}

// RemoveNocRootCertificateByVidSkidSubjectAndSerialNumber removes root certificate with specified subject and serial number from the list.
func (k Keeper) _filterAndSetNocRootCertificateByVidAndSkid(
	ctx sdk.Context,
	vid int32,
	subjectKeyID string,
	predicate CertificatePredicate,
) {
	nocCertificates, _ := k.GetNocRootCertificatesByVidAndSkid(ctx, vid, subjectKeyID)
	filteredCertificates := filterCertificates(&nocCertificates.Certs, predicate)

	if len(filteredCertificates) > 0 {
		nocCertificates.Certs = filteredCertificates
		k.SetNocRootCertificatesByVidAndSkid(ctx, nocCertificates)
	} else {
		k.RemoveNocRootCertificatesByVidAndSkid(ctx, vid, subjectKeyID)
	}
}

// GetAllNocRootCertificatesByVidAndSkid returns all nocRootCertificatesByVidAndSkid.
func (k Keeper) GetAllNocRootCertificatesByVidAndSkid(ctx sdk.Context) (list []types.NocRootCertificatesByVidAndSkid) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), pkitypes.KeyPrefix(types.NocRootCertificatesByVidAndSkidKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.NocRootCertificatesByVidAndSkid
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
