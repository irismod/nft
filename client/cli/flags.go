package cli

import (
	flag "github.com/spf13/pflag"
)

const (
	FlagTokenName  = "token-name"
	FlagTokenURI  = "token-uri"
	FlagTokenData = "token-data"
	FlagRecipient = "recipient"
	FlagOwner     = "owner"
	FlagDenom     = "denom"
	FlagSchema    = "schema"
)

var (
	FsIssueDenom  = flag.NewFlagSet("", flag.ContinueOnError)
	FsMintNFT     = flag.NewFlagSet("", flag.ContinueOnError)
	FsEditNFT     = flag.NewFlagSet("", flag.ContinueOnError)
	FsTransferNFT = flag.NewFlagSet("", flag.ContinueOnError)
	FsQuerySupply = flag.NewFlagSet("", flag.ContinueOnError)
	FsQueryOwner  = flag.NewFlagSet("", flag.ContinueOnError)
)

func init() {
	FsIssueDenom.String(FlagSchema, "", "denom data structure definition")

	FsMintNFT.String(FlagTokenURI, "", "URI for supplemental off-chain tokenData (should return a JSON object)")
	FsMintNFT.String(FlagRecipient, "", "Receiver of the nft, if not filled, the default is the sender of the transaction")
	FsMintNFT.String(FlagTokenData, "", "the origin data of nft")
	FsMintNFT.String(FlagTokenName, "", "the name of nft")

	FsEditNFT.String(FlagTokenURI, "[do-not-modify]", "URI for supplemental off-chain tokenData (should return a JSON object)")
	FsEditNFT.String(FlagTokenData, "[do-not-modify]", "the tokenData of nft")
	FsEditNFT.String(FlagTokenName, "", "the name of nft")

	FsTransferNFT.String(FlagTokenURI, "[do-not-modify]", "URI for supplemental off-chain tokenData (should return a JSON object)")
	FsTransferNFT.String(FlagTokenData, "[do-not-modify]", "the tokenData of nft")
	FsTransferNFT.String(FlagTokenName, "", "the name of nft")

	FsQuerySupply.String(FlagOwner, "", "the owner of a nft")

	FsQueryOwner.String(FlagDenom, "", "the name of a collection")
}
