// Copyright 2020 DSR Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cli

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store/iavl"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/zigbee-alliance/distributed-compliance-ledger/utils/bytes"
)

const (
	FlagPreviousHeight      = "prev-height"
	FlagPreviousHeightUsage = "Query data from previous height to avoid delay linked to state proof verification"
)

type ReadResult struct {
	Result json.RawMessage `json:"result"`
	Height int64           `json:"height"`
}

func NewReadResult(result json.RawMessage, height int64) ReadResult {
	return ReadResult{
		Result: result,
		Height: height,
	}
}

type RangeResult struct {
	Total int               `json:"total"`
	Items []json.RawMessage `json:"items"`
}

// Implement fmt.Stringer.
func (n ReadResult) String() string {
	res, err := json.Marshal(n)
	if err != nil {
		panic(err)
	}

	return string(res)
}

type CliContext struct {
	context client.CLIContext
}

func NewCLIContext() CliContext {
	return CliContext{
		context: context.NewCLIContext(),
	}
}

func (ctx CliContext) Context() client.CLIContext {
	return ctx.context
}

func (ctx CliContext) Codec() *codec.Codec {
	return ctx.context.Codec
}

func (ctx CliContext) FromAddress() sdk.AccAddress {
	return ctx.context.GetFromAddress()
}

func (ctx CliContext) WithCodec(cdc *codec.Codec) CliContext {
	ctx.context = ctx.context.WithCodec(cdc)

	return ctx
}

func (ctx CliContext) WithHeight(height int64) CliContext {
	ctx.context = ctx.context.WithHeight(height)

	return ctx
}

func (ctx CliContext) FormerHeight() (int64, error) {
	node, err := ctx.context.GetNode()
	if err != nil {
		return 0, err
	}

	status, err := node.Status()
	if err != nil {
		return 0, err
	}

	return status.SyncInfo.LatestBlockHeight - 1, nil
}

func (ctx CliContext) WithFormerHeight() (CliContext, error) {
	height, err := ctx.FormerHeight()
	if err != nil {
		return CliContext{}, err
	}

	ctx.context = ctx.context.WithHeight(height)

	return ctx, nil
}

func (ctx CliContext) WithHeightFromFlag() (CliContext, error) {
	if viper.GetBool(FlagPreviousHeight) {
		return ctx.WithFormerHeight()
	}

	return ctx.WithHeight(0), nil
}

func (ctx CliContext) QueryStore(key []byte, storeName string) ([]byte, int64, error) {
	ctx, err := ctx.WithHeightFromFlag()
	if err != nil {
		return nil, 0, err
	}

	return ctx.context.QueryStore(key, storeName)
}

func (ctx CliContext) QueryRange(startKey []byte, endKey []byte, limit int,
	storeName string) (iavl.RangeRes, int64, error) {
	ctx, err := ctx.WithHeightFromFlag()
	if err != nil {
		return iavl.RangeRes{}, 0, err
	}

	req := iavl.RangeReq{
		StartKey: startKey,
		EndKey:   endKey,
		Limit:    limit,
	}

	return ctx.context.QueryRange(req, storeName)
}

func (ctx CliContext) QueryWithData(path string, data interface{}) ([]byte, int64, error) {
	return ctx.context.QueryWithData(path, ctx.context.Codec.MustMarshalJSON(data))
}

func (ctx CliContext) QueryList(path string, params interface{}) error {
	res, height, err := ctx.QueryWithData(path, params)
	if err != nil {
		return sdk.ErrInternal(fmt.Sprintf("Could not get data: %s\n", err))
	}

	return ctx.PrintWithHeight(res, height)
}

func (ctx CliContext) QueryAllWithProof(storeKey string,
	prefix []byte, totalKey []byte, valueUnmarshaler func([]byte) json.RawMessage) error {
	// Query range (all items)
	rangeRes, rangeHeight, err := ctx.QueryRange(prefix, bytes.CpIncr(prefix), 0, storeKey)
	if err != nil {
		return sdk.ErrInternal(fmt.Sprintf("Could not get data: %s\n", err))
	}

	// Compare values length with the total number of items at the same height
	totalRes, totalHeight, err := ctx.context.WithHeight(rangeHeight).QueryStore(totalKey, storeKey)
	if err != nil {
		return sdk.ErrInternal(fmt.Sprintf("Could not get data: %s\n", err))
	}

	if rangeHeight != totalHeight {
		panic("should not happen")
	}

	var total int
	if totalRes == nil {
		total = 0
	} else {
		ctx.Codec().MustUnmarshalBinaryLengthPrefixed(totalRes, &total)
	}

	if len(rangeRes.Values) != total {
		return sdk.ErrInternal(fmt.Sprintf("Response length doesn't match value stored in totoal key: %s\n", err))
	}

	// Convert result to json
	result := RangeResult{
		total,
		make([]json.RawMessage, 0, len(rangeRes.Values)),
	}

	for _, valueBytes := range rangeRes.Values {
		result.Items = append(result.Items, valueUnmarshaler(valueBytes))
	}

	return ctx.PrintWithHeight(ctx.Codec().MustMarshalJSON(result), rangeHeight)
}

func (ctx CliContext) HandleWriteMessage(msg sdk.Msg) error {
	err := msg.ValidateBasic()
	if err != nil {
		return err
	}

	txBldr := auth.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(ctx.context.Codec))

	return utils.GenerateOrBroadcastMsgs(ctx.context, txBldr, []sdk.Msg{msg})
}

func (ctx CliContext) EncodeAndPrintWithHeight(data interface{}, height int64) (err error) {
	out, err := json.Marshal(data)
	if err != nil {
		return sdk.ErrInternal(fmt.Sprintf("Could not encode result: %v", err))
	}

	return ctx.PrintWithHeight(out, height)
}

func (ctx CliContext) PrintWithHeight(out []byte, height int64) (err error) {
	var value json.RawMessage

	ctx.context.Codec.MustUnmarshalJSON(out, &value)

	return ctx.context.PrintOutput(NewReadResult(value, height))
}

func (ctx CliContext) ReadFromFile(target string) (string, error) {
	if _, err := os.Stat(target); err == nil { // check whether it is a path
		bytes, err := ioutil.ReadFile(target)
		if err != nil {
			return "", err
		}

		return string(bytes), nil
	} else { // else return as is
		return target, nil
	}
}

func SignedCommands(cmds ...*cobra.Command) []*cobra.Command {
	for _, c := range cmds {
		_ = c.MarkFlagRequired(flags.FlagFrom)
	}

	return cmds
}
