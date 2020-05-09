package cli

import (
	"bufio"
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/version"
	"github.com/cosmos/cosmos-sdk/x/auth"
	authclient "github.com/cosmos/cosmos-sdk/x/auth/client"

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
		GetCmdIssueDenom(cdc),
		GetCmdMintNFT(cdc),
		GetCmdEditNFT(cdc),
		GetCmdTransferNFT(cdc),
		GetCmdBurnNFT(cdc),
	)...)

	return txCmd
}

// GetCmdMintNFT is the CLI command for a MintNFT transaction
func GetCmdIssueDenom(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use: "issue [denom]",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Issue a new denom.
Example:
$ %s tx nft issue [denom] --from=<key-name> --metadata=<schema> --chain-id=<chain-id> --fees=<fee>`,
				version.ClientName,
			),
		),
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(authclient.GetTxEncoder(cdc))
			cliCtx := context.NewCLIContextWithInput(inBuf).WithCodec(cdc)

			denom := args[0]
			metadata := viper.GetString(FlagMetadata)

			msg := types.NewMsgIssueDenom(cliCtx.GetFromAddress(), denom, metadata)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return authclient.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
	cmd.Flags().AddFlagSet(FsMintNFT)
	return cmd
}

// GetCmdMintNFT is the CLI command for a MintNFT transaction
func GetCmdMintNFT(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use: "mint [denom] [tokenID]",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Mint an NFT and set the owner to the recipient.
Example:
$ %s tx nft mint [denom] [tokenID] --token-uri=<token-uri> --recipient=<recipient> --from=<key-name> --chain-id=<chain-id> --fees=<fee>`,
				version.ClientName,
			),
		),
		Args: cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(authclient.GetTxEncoder(cdc))
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
			return authclient.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
	cmd.Flags().AddFlagSet(FsMintNFT)
	return cmd
}

// GetCmdEditNFT is the CLI command for sending an MsgEditNFT transaction
func GetCmdEditNFT(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use: "edit [denom] [tokenID]",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Edit the metadata of an NFT.
Example:
$ %s tx nft edit [denom] [tokenID] --token-uri=<token-uri> --from=<key-name> --chain-id=<chain-id> --fees=<fee>`,
				version.ClientName,
			),
		),
		Args: cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(authclient.GetTxEncoder(cdc))
			cliCtx := context.NewCLIContextWithInput(inBuf).WithCodec(cdc)

			denom := args[0]
			tokenID := args[1]
			tokenURI := viper.GetString(FlagTokenURI)
			metadata := viper.GetString(FlagMetadata)

			msg := types.NewMsgEditNFT(cliCtx.GetFromAddress(), tokenID, denom, tokenURI, metadata)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return authclient.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
	cmd.Flags().AddFlagSet(FsEditNFT)
	return cmd
}

// GetCmdTransferNFT is the CLI command for sending a TransferNFT transaction
func GetCmdTransferNFT(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use: "transfer [recipient] [denom] [tokenID]",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Transfer a NFT to a recipient.
Example:
$ %s tx nft transfer [recipient] [denom] [tokenID] --token-uri=<token-uri> --from=<key-name> --chain-id=<chain-id> --fees=<fee>`,
				version.ClientName,
			),
		),
		Args: cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(authclient.GetTxEncoder(cdc))
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
			return authclient.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
	cmd.Flags().AddFlagSet(FsTransferNFT)
	return cmd
}

// GetCmdBurnNFT is the CLI command for sending a BurnNFT transaction
func GetCmdBurnNFT(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use: "burn [denom] [tokenID]",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Burn an NFT.
Example:
$ %s tx nft burn [denom] [tokenID] --from=<key-name> --chain-id=<chain-id> --fees=<fee>`,
				version.ClientName,
			),
		),
		Args: cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(authclient.GetTxEncoder(cdc))
			cliCtx := context.NewCLIContextWithInput(inBuf).WithCodec(cdc)

			denom := args[0]
			tokenID := args[1]

			msg := types.NewMsgBurnNFT(cliCtx.GetFromAddress(), tokenID, denom)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return authclient.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
	return cmd
}
