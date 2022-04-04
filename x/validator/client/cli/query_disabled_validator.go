package cli

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"
	"github.com/zigbee-alliance/distributed-compliance-ledger/utils/cli"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/validator/types"
)

func CmdListDisabledValidator() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "all-disabled-validators",
		Short: "Query the list of all disabled validators",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllDisabledValidatorRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.DisabledValidatorAll(context.Background(), params)
			if cli.IsKeyNotFoundRpcError(err) {
				return clientCtx.PrintString(cli.LightClientProxyForListQueries)
			}
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddPaginationFlagsToCmd(cmd, cmd.Use)
	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdShowDisabledValidator() *cobra.Command {
	var address string

	cmd := &cobra.Command{
		Use:   "disabled-validator --address [address]",
		Short: "Query disabled validator by address",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)

			var res types.DisabledValidator

			addr, err := sdk.ValAddressFromBech32(address)
			if err != nil {
				addr2, err2 := sdk.AccAddressFromBech32(address)
				if err2 != nil {
					return err2
				}
				addr = sdk.ValAddress(addr2)
			}

			return cli.QueryWithProof(
				clientCtx,
				types.StoreKey,
				types.DisabledValidatorKeyPrefix,
				types.DisabledValidatorKey(addr.String()),
				&res,
			)
		},
	}

	cmd.Flags().StringVar(&address, FlagAddress, "", "Validator address")

	_ = cmd.MarkFlagRequired(FlagAddress)

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
