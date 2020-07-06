package cli

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/codec"

	"github.com/pewt1403/nameservice/x/nameservice/types"
)

// GetQueryCmd returns the cli query commands for this module
func GetQueryCmd(queryRoute string, cdc *codec.Codec) *cobra.Command {
	// Group nameservice queries under a subcommand
	nameserviceQueryCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("Querying commands for the %s module", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	nameserviceQueryCmd.AddCommand(
		flags.GetCommands(
			// TODO: Add query Cmds
			GetCmdResolveName(types.StoreKey, cdc),
			GetCmdWhois(types.StoreKey, cdc),
			GetCmdNames(types.StoreKey, cdc),
		)...,
	)

	return nameserviceQueryCmd
}

// TODO: Add Query Commands
func GetCmdResolveName(queryRoute string, cdc *codec.Codec) *cobra.Command{
	return &cobra.Command{
		Use: "resolve [name]",
		Short: "resolve name",
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			name := args[0]

			res, _, err := cliCtx.QueryWithData()
		}
	}
}