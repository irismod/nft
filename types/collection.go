package types

import (
	"bytes"
	"fmt"

	"github.com/irismod/nft/exported"
)

// Collection of non fungible tokens
type Collection struct {
	Denom Denom `json:"denom" yaml:"denom"` // name of the collection; not exported to clients
	NFTs  NFTs  `json:"nfts" yaml:"nfts"`   // NFTs that belong to a collection
}

// NewCollection creates a new NFT Collection
func NewCollection(denom Denom, nfts NFTs) Collection {
	return Collection{
		Denom: denom,
		NFTs:  nfts,
	}
}

// AddNFT adds an NFT to the collection
func (c Collection) AddNFT(nft exported.NFT) Collection {
	c.NFTs = append(c.NFTs, nft)
	return c
}

func (c Collection) Supply() int {
	return len(c.NFTs)
}

// String follows stringer interface
func (c Collection) String() string {
	return fmt.Sprintf(`Denom: 				%s
NFTs:

%s`,
		c.Denom,
		c.NFTs.String(),
	)
}

//Collections define an array of Collection
type Collections []Collection

// NewCollection creates a new NFT Collection
func NewCollections(c ...Collection) Collections {
	return append([]Collection{}, c...)
}

// String follows stringer interface
func (cs Collections) String() string {
	if len(cs) == 0 {
		return ""
	}
	var buf bytes.Buffer
	for _, collection := range cs {
		if buf.Len() > 0 {
			buf.WriteString("\n")
		}
		buf.WriteString(collection.String())
	}
	return buf.String()
}
