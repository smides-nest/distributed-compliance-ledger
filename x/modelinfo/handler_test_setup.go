package modelinfo

//nolint:goimports
import (
	"git.dsr-corporation.com/zb-ledger/zb-ledger/integration_tests/constants"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/auth"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/go-amino"
	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/libs/log"
	dbm "github.com/tendermint/tm-db"
)

type TestSetup struct {
	Cdc             *amino.Codec
	Ctx             sdk.Context
	ModelinfoKeeper Keeper
	authKeeper      auth.Keeper
	Handler         sdk.Handler
	Querier         sdk.Querier
	Vendor          sdk.AccAddress
}

func Setup() TestSetup {
	// Init Codec
	cdc := codec.New()
	sdk.RegisterCodec(cdc)

	// Init KVSore
	db := dbm.NewMemDB()

	dbStore := store.NewCommitMultiStore(db)

	modelinfoKey := sdk.NewKVStoreKey(StoreKey)
	dbStore.MountStoreWithDB(modelinfoKey, sdk.StoreTypeIAVL, db)

	authKey := sdk.NewKVStoreKey(auth.StoreKey)
	dbStore.MountStoreWithDB(authKey, sdk.StoreTypeIAVL, db)

	_ = dbStore.LoadLatestVersion()

	// Init Keepers
	modelinfoKeeper := NewKeeper(modelinfoKey, cdc)
	authKeeper := auth.NewKeeper(authKey, cdc)

	// Create context
	ctx := sdk.NewContext(dbStore, abci.Header{ChainID: testconstants.ChainID}, false, log.NewNopLogger())

	// Create Handler and Querier
	querier := NewQuerier(modelinfoKeeper)
	handler := NewHandler(modelinfoKeeper, authKeeper)

	account := testconstants.Address1
	authKeeper.AssignRole(ctx, account, auth.Vendor)

	setup := TestSetup{
		Cdc:             cdc,
		Ctx:             ctx,
		ModelinfoKeeper: modelinfoKeeper,
		authKeeper:      authKeeper,
		Handler:         handler,
		Querier:         querier,
		Vendor:          account,
	}

	return setup
}

func TestMsgAddModelInfo(signer sdk.AccAddress) MsgAddModelInfo {
	return MsgAddModelInfo{
		VID:                      testconstants.VID,
		PID:                      testconstants.PID,
		CID:                      testconstants.CID,
		Name:                     testconstants.Name,
		Description:              testconstants.Description,
		SKU:                      testconstants.Sku,
		FirmwareVersion:          testconstants.FirmwareVersion,
		HardwareVersion:          testconstants.HardwareVersion,
		Custom:                   testconstants.Custom,
		TisOrTrpTestingCompleted: testconstants.TisOrTrpTestingCompleted,
		Signer:                   signer,
	}
}

func TestMsgUpdatedModelInfo(signer sdk.AccAddress) MsgUpdateModelInfo {
	return MsgUpdateModelInfo{
		VID:                      testconstants.VID,
		PID:                      testconstants.PID,
		CID:                      testconstants.CID + 1,
		Description:              "New Description",
		Custom:                   "New Custom Data",
		TisOrTrpTestingCompleted: testconstants.TisOrTrpTestingCompleted,
		Signer:                   signer,
	}
}
