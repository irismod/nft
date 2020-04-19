package cli

import (
	"encoding/binary"
	"fmt"
	"github.com/spf13/viper"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
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
		GetCmdQuerySupply(queryRoute, cdc),
		GetCmdQueryOwner(queryRoute, cdc),
		GetCmdQueryCollection(queryRoute, cdc),
		GetCmdQueryDenoms(queryRoute, cdc),
		GetCmdQueryNFT(queryRoute, cdc),
	)...)

	return queryCmd
}

// GetCmdQuerySupply queries the supply of a nft collection
func GetCmdQuerySupply(queryRoute string, cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "supply",
		Short:   "total supply of a collection or owner of NFTs",
		Example: "nft supply",
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			ownerStr := viper.GetString(FlagOwner)
			owner, err := sdk.AccAddressFromBech32(ownerStr)
			if err != nil {
				return err
			}

			denom := viper.GetString(FlagDenom)
			params := types.NewQuerySupplyParams(denom, owner)
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
		Use:     "owner [address]",
		Short:   "get the NFTs owned by an account address",
		Example: "nft owner <address> --denom=<denom>",
		Args:    cobra.ExactArgs(1),
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
		Use:     "collection [denom]",
		Short:   "get all the NFTs from a given collection",
		Example: "nft collection <denom>",
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			params := types.NewQueryCollectionParams(args[0])
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
		Use:     "denoms",
		Short:   "queries all denominations of all collections of NFTs",
		Example: "nft denoms",
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			res, _, err := cliCtx.QueryWithData(
				fmt.Sprintf("custom/%s/%s", queryRoute, types.QueryDenoms), nil)
			if err != nil {
				return err
			}

			var out []string
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
		Use:     "token [denom] [ID]",
		Short:   "query a single NFT from a collection",
		Example: "nft token <denom> [ID]",
		Args:    cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			params := types.NewQueryNFTParams(args[0], args[1])
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
