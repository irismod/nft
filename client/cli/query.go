package cli

import (
	"encoding/binary"
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/version"

	"github.com/irismod/nft/exported"
	"github.com/irismod/nft/types"
)

// GetQueryCmd returns the cli query commands for this module
func GetQueryCmd(queryRoute string, cdc *codec.Codec) *cobra.Command {
	queryCmd := &cobra.Command{
		Use:                types.ModuleName,
		Short:              "Querying commands for the NFT module",
		DisableFlagParsing: true,
	}

	queryCmd.AddCommand(flags.GetCommands(
		GetCmdQueryDenom(queryRoute, cdc),
		GetCmdQueryDenoms(queryRoute, cdc),
		GetCmdQueryCollection(queryRoute, cdc),
		GetCmdQuerySupply(queryRoute, cdc),
		GetCmdQueryOwner(queryRoute, cdc),
		GetCmdQueryNFT(queryRoute, cdc),
	)...)

	return queryCmd
}

// GetCmdQuerySupply queries the supply of a nft collection
func GetCmdQuerySupply(queryRoute string, cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use: "supply [denomID]",
		Long: strings.TrimSpace(
			fmt.Sprintf(`total supply of a collection or owner of NFTs.
Example:
$ %s q nft supply [denomID]`, version.ClientName)),
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			var owner sdk.AccAddress
			var err error

			ownerStr := strings.TrimSpace(viper.GetString(FlagOwner))
			if len(ownerStr) > 0 {
				owner, err = sdk.AccAddressFromBech32(ownerStr)
				if err != nil {
					return err
				}
			}

			denomID := strings.TrimSpace(args[0])
			if err := types.ValidateDenomID(denomID); err != nil {
				return err
			}

			params := types.NewQuerySupplyParams(denomID, owner)
			bz, err := cdc.MarshalJSON(params)
			if err != nil {
				return err
			}

			res, _, err := cliCtx.QueryWithData(
				fmt.Sprintf("custom/%s/%s", queryRoute, types.QuerySupply), bz)
			if err != nil {
				return err
			}

			out := binary.LittleEndian.Uint64(res)
			return cliCtx.PrintOutput(out)
		},
	}
	cmd.Flags().AddFlagSet(FsQuerySupply)
	return cmd
}

// GetCmdQueryOwner queries all the NFTs owned by an account
func GetCmdQueryOwner(queryRoute string, cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use: "owner [address]",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Get the NFTs owned by an account address
Example:
$ %s q nft owner <address> --denom=<denom>`, version.ClientName)),
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			address, err := sdk.AccAddressFromBech32(args[0])
			if err != nil {
				return err
			}

			denom := viper.GetString(FlagDenom)
			params := types.NewQueryOwnerParams(denom, address)
			bz, err := cdc.MarshalJSON(params)
			if err != nil {
				return err
			}

			res, _, err := cliCtx.QueryWithData(
				fmt.Sprintf("custom/%s/%s", queryRoute, types.QueryOwner), bz)
			if err != nil {
				return err
			}

			var out types.Owner
			err = cdc.UnmarshalJSON(res, &out)
			if err != nil {
				return err
			}

			return cliCtx.PrintOutput(out)
		},
	}
	cmd.Flags().AddFlagSet(FsQueryOwner)
	return cmd
}

// GetCmdQueryCollection queries all the NFTs from a collection
func GetCmdQueryCollection(queryRoute string, cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use: "collection [denomID]",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Get all the NFTs from a given collection
Example:
$ %s q nft collection <denom>`, version.ClientName)),
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			denomID := strings.TrimSpace(args[0])
			if err := types.ValidateDenomID(denomID); err != nil {
				return err
			}

			params := types.NewQueryCollectionParams(denomID)
			bz, err := cdc.MarshalJSON(params)
			if err != nil {
				return err
			}

			res, _, err := cliCtx.QueryWithData(
				fmt.Sprintf("custom/%s/%s", queryRoute, types.QueryCollection), bz)
			if err != nil {
				return err
			}

			var out types.Collection
			if err = cdc.UnmarshalJSON(res, &out); err != nil {
				return err
			}
			return cliCtx.PrintOutput(out)
		},
	}
	return cmd
}

// GetCmdQueryDenoms queries all denoms
func GetCmdQueryDenoms(queryRoute string, cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use: "denoms",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Query all denominations of all collections of NFTs
Example:
$ %s q nft denoms`, version.ClientName)),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			res, _, err := cliCtx.QueryWithData(
				fmt.Sprintf("custom/%s/%s", queryRoute, types.QueryDenoms), nil)
			if err != nil {
				return err
			}

			var out []types.Denom
			if err = cdc.UnmarshalJSON(res, &out); err != nil {
				return err
			}
			return cliCtx.PrintOutput(out)
		},
	}
	return cmd
}

// GetCmdQueryDenoms queries the specified denoms
func GetCmdQueryDenom(queryRoute string, cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use: "denom [denomID]",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Query the denominations by the specified denmo name
Example:
$ %s q nft denom <denomID>`, version.ClientName)),
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			denom := strings.TrimSpace(args[0])
			if err := types.ValidateDenomID(denom); err != nil {
				return err
			}

			params := types.NewQueryDenomParams(denom)
			bz, err := cdc.MarshalJSON(params)
			if err != nil {
				return err
			}

			res, _, err := cliCtx.QueryWithData(
				fmt.Sprintf("custom/%s/%s", queryRoute, types.QueryDenom), bz)
			if err != nil {
				return err
			}

			var out types.Denom
			if err = cdc.UnmarshalJSON(res, &out); err != nil {
				return err
			}
			return cliCtx.PrintOutput(out)
		},
	}
	return cmd
}

// GetCmdQueryNFT queries a single NFTs from a collection
func GetCmdQueryNFT(queryRoute string, cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use: "token [denomID] [tokenID]",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Query a single NFT from a collection
Example:
$ %s q nft token <denomID> <tokenID>`, version.ClientName)),
		Args: cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			denomID := strings.TrimSpace(args[0])
			if err := types.ValidateDenomID(denomID); err != nil {
				return err
			}

			tokenID := strings.TrimSpace(args[1])
			if err := types.ValidateTokenID(tokenID); err != nil {
				return err
			}

			params := types.NewQueryNFTParams(denomID, tokenID)
			bz, err := cdc.MarshalJSON(params)
			if err != nil {
				return err
			}

			res, _, err := cliCtx.QueryWithData(
				fmt.Sprintf("custom/%s/%s", queryRoute, types.QueryNFT), bz)
			if err != nil {
				return err
			}

			var out exported.NFT
			if err = cdc.UnmarshalJSON(res, &out); err != nil {
				return err
			}
			return cliCtx.PrintOutput(out)
		},
	}
	return cmd
}
