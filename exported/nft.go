package exported

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// NFT non fungible token interface
type NFT interface {
	GetID() string
	GetOwner() sdk.AccAddress
	GetTokenURI() string
	GetTokenData() string
	String() string
}
