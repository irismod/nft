package cli

import (
	"context"
	"fmt"

	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/version"

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
$ %s q nft supply [denom]`, version.AppName)),
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := client.NewContext().WithCodec(cdc).WithJSONMarshaler(cdc)

			var owner sdk.AccAddress
			var err error

			ownerStr := strings.TrimSpace(viper.GetString(FlagOwner))
			if len(ownerStr) > 0 {
				owner, err = sdk.AccAddressFromBech32(ownerStr)
				if err != nil {
					return err
				}
			}

			denom := strings.TrimSpace(args[0])
			if err := types.ValidateDenomID(denom); err != nil {
				return err
			}

			queryClient := types.NewQueryClient(cliCtx)
			resp, err := queryClient.Supply(context.Background(), &types.QuerySupplyRequest{
				Denom: denom,
				Owner: owner,
			})
			if err != nil {
				return err
			}
			return cliCtx.PrintOutput(resp.Amount)
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
$ %s q nft owner <address> --denom=<denom>`, version.AppName)),
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := client.NewContext().WithCodec(cdc).WithJSONMarshaler(cdc)

			address, err := sdk.AccAddressFromBech32(args[0])
			if err != nil {
				return err
			}

			denom := viper.GetString(FlagDenom)
			queryClient := types.NewQueryClient(cliCtx)
			resp, err := queryClient.Owner(context.Background(), &types.QueryOwnerRequest{
				Denom: denom,
				Owner: address,
			})
			if err != nil {
				return err
			}

			return cliCtx.PrintOutput(resp.Owner)
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
$ %s q nft collection <denom>`, version.AppName)),
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := client.NewContext().WithCodec(cdc).WithJSONMarshaler(cdc)

			denom := strings.TrimSpace(args[0])
			if err := types.ValidateDenomID(denom); err != nil {
				return err
			}

			queryClient := types.NewQueryClient(cliCtx)
			resp, err := queryClient.Collection(context.Background(), &types.QueryCollectionRequest{
				Denom: denom,
			})
			if err != nil {
				return err
			}
			return cliCtx.PrintOutput(resp.Collection)
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
$ %s q nft denoms`, version.AppName)),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := client.NewContext().WithCodec(cdc).WithJSONMarshaler(cdc)

			queryClient := types.NewQueryClient(cliCtx)
			resp, err := queryClient.Denoms(context.Background(), &types.QueryDenomsRequest{})
			if err != nil {
				return err
			}
			return cliCtx.PrintOutput(resp.Denoms)
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
$ %s q nft denom <denom>`, version.AppName)),
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := client.NewContext().WithCodec(cdc).WithJSONMarshaler(cdc)

			denom := strings.TrimSpace(args[0])
			if err := types.ValidateDenomID(denom); err != nil {
				return err
			}

			queryClient := types.NewQueryClient(cliCtx)
			resp, err := queryClient.Denom(context.Background(), &types.QueryDenomRequest{
				Denom: denom,
			})
			if err != nil {
				return err
			}
			return cliCtx.PrintOutput(resp.Denom)
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
$ %s q nft token <denom> <tokenID>`, version.AppName)),
		Args: cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := client.NewContext().WithCodec(cdc).WithJSONMarshaler(cdc)

			denom := strings.TrimSpace(args[0])
			if err := types.ValidateDenomID(denom); err != nil {
				return err
			}

			tokenID := strings.TrimSpace(args[1])
			if err := types.ValidateTokenID(tokenID); err != nil {
				return err
			}

			queryClient := types.NewQueryClient(cliCtx)
			resp, err := queryClient.NFT(context.Background(), &types.QueryNFTRequest{
				Denom:   denom,
				TokenID: tokenID,
			})
			if err != nil {
				return err
			}
			return cliCtx.PrintOutput(resp.NFT)
		},
	}
	return cmd
}
