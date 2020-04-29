package cli

import (
	flag "github.com/spf13/pflag"
)

const (
	FlagTokenURI  = "token-uri"
	FlagRecipient = "recipient"
	FlagOwner     = "owner"
	FlagDenom     = "denom"
)

var (
	FsMintNFT     = flag.NewFlagSet("", flag.ContinueOnError)
	FsEditNFT     = flag.NewFlagSet("", flag.ContinueOnError)
	FsTransferNFT = flag.NewFlagSet("", flag.ContinueOnError)
	FsQuerySupply = flag.NewFlagSet("", flag.ContinueOnError)
	FsQueryOwner  = flag.NewFlagSet("", flag.ContinueOnError)
)

func init() {
	FsMintNFT.String(FlagTokenURI, "[do-not-modify]", "URI for supplemental off-chain metadata (should return a JSON object)")
	FsMintNFT.String(FlagRecipient, "", "Receiver of the nft,if not filled, the default is the sender of the transaction")

	FsEditNFT.String(FlagTokenURI, "[do-not-modify]", "URI for supplemental off-chain metadata (should return a JSON object)")

	FsTransferNFT.String(FlagTokenURI, "[do-not-modify]", "URI for supplemental off-chain metadata (should return a JSON object)")

	FsQuerySupply.String(FlagOwner, "", "the owner of a nft")

	FsQueryOwner.String(FlagDenom, "", "the name of a collection")
}
