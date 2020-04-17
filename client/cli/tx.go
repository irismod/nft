package cli

import (
	"bufio"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"

	"github.com/irismod/nft/types"
)

// Edit metadata flags
const (
	flagTokenURI  = "token-uri"
	flagRecipient = "recipient"
	flagOwner     = "owner"
	flagDenom     = "denom"
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
		GetCmdTransferNFT(cdc),
		GetCmdEditNFT(cdc),
		GetCmdMintNFT(cdc),
		GetCmdBurnNFT(cdc),
	)...)

	return txCmd
}

// GetCmdTransferNFT is the CLI command for sending a TransferNFT transaction
func GetCmdTransferNFT(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "transfer [sender] [recipient] [denom] [tokenID]",
		Short:   "transfer a NFT to a recipient",
		Example: "nft transfer [sender] [recipient] [denom] [tokenID] --token-uri=<token-uri> --from=<key-name> --chain-id=<chain-id> --fee=<fee>",
		Args:    cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) error {
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))
			cliCtx := context.NewCLIContextWithInput(inBuf).WithCodec(cdc)

			sender, err := sdk.AccAddressFromBech32(args[0])
			if err != nil {
				return err
			}

			recipient, err := sdk.AccAddressFromBech32(args[1])
			if err != nil {
				return err
			}

			denom := args[2]
			tokenID := args[3]
			tokenURI := viper.GetString(flagTokenURI)

			msg := types.NewMsgTransferNFT(sender, recipient, denom, tokenID, tokenURI)
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
	cmd.Flags().String(flagTokenURI, "[do-not-modify]", "Extra properties available for querying")
	return cmd
}

// GetCmdEditNFT is the CLI command for sending an MsgEditNFT transaction
func GetCmdEditNFT(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "edit [denom] [tokenID]",
		Short:   "edit the metadata of an NFT",
		Example: "nft edit [denom] [tokenID] --token-uri=<token-uri> --from=<key-name> --chain-id=<chain-id> --fee=<fee>",
		Args:    cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))
			cliCtx := context.NewCLIContextWithInput(inBuf).WithCodec(cdc)

			denom := args[0]
			tokenID := args[1]
			tokenURI := viper.GetString(flagTokenURI)

			msg := types.NewMsgEditNFT(cliCtx.GetFromAddress(), tokenID, denom, tokenURI)
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}

	cmd.Flags().String(flagTokenURI, "", "Extra properties available for querying")
	return cmd
}

// GetCmdMintNFT is the CLI command for a MintNFT transaction
func GetCmdMintNFT(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "mint [denom] [tokenID]",
		Short:   "mint an NFT and set the owner to the recipient",
		Example: "nft mint [denom] [tokenID] --token-uri=<token-uri> --recipient=<recipient> --from=<key-name> --chain-id=<chain-id> --fee=<fee>",
		Args:    cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))
			cliCtx := context.NewCLIContextWithInput(inBuf).WithCodec(cdc)

			denom := args[0]
			tokenID := args[1]

			recipient, err := sdk.AccAddressFromBech32(args[2])
			if err != nil {
				return err
			}

			tokenURI := viper.GetString(flagTokenURI)

			msg := types.NewMsgMintNFT(cliCtx.GetFromAddress(), recipient, tokenID, denom, tokenURI)
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}

	cmd.Flags().String(flagTokenURI, "", "URI for supplemental off-chain metadata (should return a JSON object)")
	cmd.Flags().String(flagRecipient, "", "Receiver of the nft)")
	return cmd
}

// GetCmdBurnNFT is the CLI command for sending a BurnNFT transaction
func GetCmdBurnNFT(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "burn [denom] [tokenID]",
		Short:   "burn an NFT",
		Example: "nft mint [denom] [tokenID] --from=<key-name> --chain-id=<chain-id> --fee=<fee>",
		Args:    cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))
			cliCtx := context.NewCLIContextWithInput(inBuf).WithCodec(cdc)

			denom := args[0]
			tokenID := args[1]

			msg := types.NewMsgBurnNFT(cliCtx.GetFromAddress(), tokenID, denom)
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}

	cmd = flags.PostCommands(cmd)[0]
	return cmd
}
