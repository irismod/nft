package types

import (
	"bytes"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/irismod/nft/exported"
)

var _ exported.NFT = (*BaseNFT)(nil)

// NewBaseNFT creates a new NFT instance
func NewBaseNFT(id string, owner sdk.AccAddress, tokenURI, metadata string) BaseNFT {
	return BaseNFT{
		ID:       strings.ToLower(strings.TrimSpace(id)),
		Owner:    owner,
		TokenURI: strings.TrimSpace(tokenURI),
		Metadata: strings.TrimSpace(metadata),
	}
}

// SetOwner updates the owner address of the NFT
func (bnft *BaseNFT) SetOwner(address sdk.AccAddress) {
	bnft.Owner = address
}

// SetTokenURI edits metadata of an nft
func (bnft *BaseNFT) SetTokenURI(tokenURI string) {
	bnft.TokenURI = tokenURI
}

func (bnft *BaseNFT) SetMetadata(metadata string) {
	bnft.Metadata = metadata
}

// ----------------------------------------------------------------------------
// NFT

// NFTs define a list of NFT
type NFTs []exported.NFT

// NewNFTs creates a new set of NFTs
func NewNFTs(nfts ...exported.NFT) NFTs {
	if len(nfts) == 0 {
		return NFTs{}
	}
	return NFTs(nfts)
}

// String follows stringer interface
func (nfts NFTs) String() string {
	if len(nfts) == 0 {
		return ""
	}

	var buf bytes.Buffer
	for _, nft := range nfts {
		if buf.Len() > 0 {
			buf.WriteString("\n")
		}
		buf.WriteString(nft.String())
	}
	return buf.String()
}

func ValidateTokenID(tokenID string) error {
	tokenID = strings.TrimSpace(tokenID)
	if len(tokenID) < MinDenomLen || len(tokenID) > MaxDenomLen {
		return sdkerrors.Wrapf(ErrInvalidTokenID, "invalid tokenID %s, only accepts value [%d, %d]", denom, MinDenomLen, MaxDenomLen)
	}
	if !IsBeginWithAlpha(tokenID) || !IsAlphaNumeric(tokenID) {
		return sdkerrors.Wrapf(ErrInvalidTokenID, "invalid tokenID %s, only accepts alphanumeric characters,and begin with an english letter", denom)
	}
	return nil
}

func ValidateTokenURI(tokenURI string) error {
	if len(tokenURI) > MaxTokenURILen {
		return sdkerrors.Wrapf(ErrInvalidTokenURI, "invalid tokenURI %s, only accepts value [0, %d]", denom, MaxTokenURILen)
	}
	return nil
}
