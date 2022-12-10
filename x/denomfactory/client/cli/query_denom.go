package cli

import (
	"context"

	"github.com/G4AL-Entertainment/g4al-chain/x/denomfactory/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
)

func CmdListDenom() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-denom",
		Short: "list all Denom",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllDenomRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.DenomAll(context.Background(), params)
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

func CmdShowDenom() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-denom [symbol]",
		Short: "shows a Denom",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argSymbol := args[0]

			params := &types.QueryGetDenomRequest{
				Symbol: argSymbol,
			}

			res, err := queryClient.Denom(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
