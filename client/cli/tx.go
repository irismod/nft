package cli

import (
	"bufio"
	"strings"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/irismod/nft/types"
)

// GetTxCmd returns the transaction commands for this module
func GetTxCmd(storeKey string, cdc *codec.Codec) *cobra.Command {
	txCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "NFT transactions subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	txCmd.AddCommand(flags.PostCommands(
		GetCmdMintNFT(cdc),
		GetCmdEditNFT(cdc),
		GetCmdTransferNFT(cdc),
		GetCmdBurnNFT(cdc),
	)...)

	return txCmd
}

// GetCmdMintNFT is the CLI command for a MintNFT transaction
func GetCmdMintNFT(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "mint [denom] [tokenID]",
		Short:   "mint an NFT and set the owner to the recipient",
		Example: "nft mint [denom] [tokenID] --token-uri=<token-uri> --recipient=<recipient> --from=<key-name> --chain-id=<chain-id> --fees=<fee>",
		Args:    cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))
			cliCtx := context.NewCLIContextWithInput(inBuf).WithCodec(cdc)

			denom := args[0]
			tokenID := args[1]
			tokenURI := viper.GetString(FlagTokenURI)
			metadata := viper.GetString(FlagMetadata)

			var recipient = cliCtx.GetFromAddress()
			var err error
			recipientStr := strings.TrimSpace(viper.GetString(FlagRecipient))
			if len(recipientStr) > 0 {
				recipient, err = sdk.AccAddressFromBech32(recipientStr)
				if err != nil {
					return err
				}
			}

			msg := types.NewMsgMintNFT(cliCtx.GetFromAddress(), recipient, tokenID, denom, tokenURI, metadata)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
	cmd.Flags().AddFlagSet(FsMintNFT)
	return cmd
}

// GetCmdEditNFT is the CLI command for sending an MsgEditNFT transaction
func GetCmdEditNFT(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "edit [denom] [tokenID]",
		Short:   "edit the metadata of an NFT",
		Example: "nft edit [denom] [tokenID] --token-uri=<token-uri> --from=<key-name> --chain-id=<chain-id> --fees=<fee>",
		Args:    cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))
			cliCtx := context.NewCLIContextWithInput(inBuf).WithCodec(cdc)

			denom := args[0]
			tokenID := args[1]
			tokenURI := viper.GetString(FlagTokenURI)
			metadata := viper.GetString(FlagMetadata)

			msg := types.NewMsgEditNFT(cliCtx.GetFromAddress(), tokenID, denom, tokenURI, metadata)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
	cmd.Flags().AddFlagSet(FsEditNFT)
	return cmd
}

// GetCmdTransferNFT is the CLI command for sending a TransferNFT transaction
func GetCmdTransferNFT(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "transfer [recipient] [denom] [tokenID]",
		Short:   "transfer a NFT to a recipient",
		Example: "nft transfer [recipient] [denom] [tokenID] --token-uri=<token-uri> --from=<key-name> --chain-id=<chain-id> --fees=<fee>",
		Args:    cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))
			cliCtx := context.NewCLIContextWithInput(inBuf).WithCodec(cdc)

			sender := cliCtx.GetFromAddress()
			recipient, err := sdk.AccAddressFromBech32(args[0])
			if err != nil {
				return err
			}

			denom := args[1]
			tokenID := args[2]
			tokenURI := viper.GetString(FlagTokenURI)
			metadata := viper.GetString(FlagMetadata)

			msg := types.NewMsgTransferNFT(sender, recipient, denom, tokenID, tokenURI, metadata)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
	cmd.Flags().AddFlagSet(FsTransferNFT)
	return cmd
}

// GetCmdBurnNFT is the CLI command for sending a BurnNFT transaction
func GetCmdBurnNFT(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "burn [denom] [tokenID]",
		Short:   "burn an NFT",
		Example: "nft burn [denom] [tokenID] --from=<key-name> --chain-id=<chain-id> --fees=<fee>",
		Args:    cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))
			cliCtx := context.NewCLIContextWithInput(inBuf).WithCodec(cdc)

			denom := args[0]
			tokenID := args[1]

			msg := types.NewMsgBurnNFT(cliCtx.GetFromAddress(), tokenID, denom)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
	return cmd
}
