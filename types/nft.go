package types

import (
	"bytes"
	"fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/irismod/nft/exported"
)

var _ exported.NFT = (*BaseNFT)(nil)

// BaseNFT non fungible token definition
type BaseNFT struct {
	ID       string         `json:"id" yaml:"id"`               // id of the token; not exported to clients
	Owner    sdk.AccAddress `json:"owner" yaml:"owner"`         // account address that owns the NFT
	TokenURI string         `json:"token_uri" yaml:"token_uri"` // optional extra properties available for querying
	Metadata string         `json:"metadata" yaml:"metadata"`
}

// NewBaseNFT creates a new NFT instance
func NewBaseNFT(id string, owner sdk.AccAddress, tokenURI,schema string) BaseNFT {
	return BaseNFT{
		ID:       strings.ToLower(strings.TrimSpace(id)),
		Owner:    owner,
		TokenURI: strings.TrimSpace(tokenURI),
		Metadata: strings.TrimSpace(schema),
	}
}

// GetID returns the TokenID of the token
func (bnft BaseNFT) GetID() string { return bnft.ID }

// GetOwner returns the account address that owns the NFT
func (bnft BaseNFT) GetOwner() sdk.AccAddress { return bnft.Owner }

// SetOwner updates the owner address of the NFT
func (bnft *BaseNFT) SetOwner(address sdk.AccAddress) {
	bnft.Owner = address
}

// GetTokenURI returns the path to optional extra properties
func (bnft BaseNFT) GetTokenURI() string { return bnft.TokenURI }

// SetTokenURI edits metadata of an nft
func (bnft *BaseNFT) SetTokenURI(tokenURI string) {
	bnft.TokenURI = tokenURI
}

func (bnft BaseNFT) GetMetadata() string {
	return bnft.Metadata
}

func (bnft *BaseNFT) SetMetadata(metadata string) {
	bnft.Metadata = metadata
}

func (bnft BaseNFT) String() string {
	return fmt.Sprintf(`TokenID:				%s
Owner:			%s
TokenURI:		%s`,
		bnft.ID,
		bnft.Owner,
		bnft.TokenURI,
	)
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
