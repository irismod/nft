package cli

import (
	"fmt"

	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/version"

	"github.com/irismod/nft/types"
)

// GetTxCmd returns the transaction commands for this module
func GetTxCmd(clientCtx client.Context) *cobra.Command {
	txCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "NFT transactions subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	txCmd.AddCommand(flags.PostCommands(
		GetCmdIssueDenom(clientCtx),
		GetCmdMintNFT(clientCtx),
		GetCmdEditNFT(clientCtx),
		GetCmdTransferNFT(clientCtx),
		GetCmdBurnNFT(clientCtx),
	)...)

	return txCmd
}

// GetCmdMintNFT is the CLI command for a MintNFT transaction
func GetCmdIssueDenom(clientCtx client.Context) *cobra.Command {
	cmd := &cobra.Command{
		Use: "issue [denom]",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Issue a new denom.
Example:
$ %s tx nft issue [denomID] --from=<key-name> --name=<name> --schema=<schema> --name=<name> --chain-id=<chain-id> --fees=<fee>`,
				version.AppName,
			),
		),
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := clientCtx.InitWithInput(cmd.InOrStdin())
			msg := types.NewMsgIssueDenom(args[0],
				viper.GetString(FlagDenomName),
				viper.GetString(FlagSchema),
				cliCtx.GetFromAddress(),
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTx(clientCtx, &msg)
		},
	}
	cmd.Flags().AddFlagSet(FsIssueDenom)
	return cmd
}

// GetCmdMintNFT is the CLI command for a MintNFT transaction
func GetCmdMintNFT(clientCtx client.Context) *cobra.Command {
	cmd := &cobra.Command{
		Use: "mint [denomID] [tokenID]",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Mint an NFT and set the owner to the recipient.
Example:
$ %s tx nft mint [denomID] [tokenID] --uri=<uri> --recipient=<recipient> --from=<key-name> --chain-id=<chain-id> --fees=<fee>`,
				version.AppName,
			),
		),
		Args: cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := clientCtx.InitWithInput(cmd.InOrStdin())
			var recipient = cliCtx.GetFromAddress()
			var err error
			recipientStr := strings.TrimSpace(viper.GetString(FlagRecipient))
			if len(recipientStr) > 0 {
				recipient, err = sdk.AccAddressFromBech32(recipientStr)
				if err != nil {
					return err
				}
			}

			msg := types.NewMsgMintNFT(
				args[1],
				args[0],
				viper.GetString(FlagTokenName),
				viper.GetString(FlagTokenURI),
				viper.GetString(FlagTokenData),
				cliCtx.GetFromAddress(),
				recipient,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTx(clientCtx, &msg)
		},
	}
	cmd.Flags().AddFlagSet(FsMintNFT)
	return cmd
}

// GetCmdEditNFT is the CLI command for sending an MsgEditNFT transaction
func GetCmdEditNFT(clientCtx client.Context) *cobra.Command {
	cmd := &cobra.Command{
		Use: "edit [denomID] [tokenID]",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Edit the tokenData of an NFT.
Example:
$ %s tx nft edit [denomID] [tokenID] --uri=<uri> --from=<key-name> --chain-id=<chain-id> --fees=<fee>`,
				version.AppName,
			),
		),
		Args: cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := clientCtx.InitWithInput(cmd.InOrStdin())
			msg := types.NewMsgEditNFT(
				args[1],
				args[0],
				viper.GetString(FlagTokenName),
				viper.GetString(FlagTokenURI),
				viper.GetString(FlagTokenData),
				cliCtx.GetFromAddress(),
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTx(clientCtx, &msg)
		},
	}
	cmd.Flags().AddFlagSet(FsEditNFT)
	return cmd
}

// GetCmdTransferNFT is the CLI command for sending a TransferNFT transaction
func GetCmdTransferNFT(clientCtx client.Context) *cobra.Command {
	cmd := &cobra.Command{
		Use: "transfer [recipient] [denomID] [tokenID]",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Transfer a NFT to a recipient.
Example:
$ %s tx nft transfer [recipient] [denomID] [tokenID] --uri=<uri> --from=<key-name> --chain-id=<chain-id> --fees=<fee>`,
				version.AppName,
			),
		),
		Args: cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := clientCtx.InitWithInput(cmd.InOrStdin())
			recipient, err := sdk.AccAddressFromBech32(args[0])
			if err != nil {
				return err
			}

			msg := types.NewMsgTransferNFT(
				args[2],
				args[1],
				viper.GetString(FlagTokenName),
				viper.GetString(FlagTokenURI),
				viper.GetString(FlagTokenData),
				cliCtx.GetFromAddress(),
				recipient,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTx(clientCtx, &msg)
		},
	}
	cmd.Flags().AddFlagSet(FsTransferNFT)
	return cmd
}

// GetCmdBurnNFT is the CLI command for sending a BurnNFT transaction
func GetCmdBurnNFT(clientCtx client.Context) *cobra.Command {
	cmd := &cobra.Command{
		Use: "burn [denomID] [tokenID]",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Burn an NFT.
Example:
$ %s tx nft burn [denomID] [tokenID] --from=<key-name> --chain-id=<chain-id> --fees=<fee>`,
				version.AppName,
			),
		),
		Args: cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := clientCtx.InitWithInput(cmd.InOrStdin())
			msg := types.NewMsgBurnNFT(cliCtx.GetFromAddress(), args[1], args[0])
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTx(clientCtx, &msg)
		},
	}
	return cmd
}
