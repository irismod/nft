package types

import (
	"strings"
)

// Collection of non fungible tokens
type Collection struct {
	Denom string `json:"denom,omitempty" yaml:"denom"` // name of the collection; not exported to clients
	NFTs  NFTs   `json:"nfts" yaml:"nfts"`             // NFTs that belong to a collection
}

// NewCollection creates a new NFT Collection
func NewCollection(denom string, nfts NFTs) Collection {
	return Collection{
		Denom: strings.TrimSpace(denom),
		NFTs:  nfts,
	}
}

func (c Collection) Supply() int {
	return len(c.NFTs)
}

//Collections define an array of Collection
type Collections []Collection
