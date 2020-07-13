package types

import (
	"bytes"
	"fmt"

	"github.com/irismod/nft/exported"
)

// NewCollection creates a new NFT Collection
func NewCollection(denom Denom, nfts []exported.NFT) (c Collection) {
	c.Denom = denom
	for _, nft := range nfts {
		c = c.AddNFT(nft.(BaseNFT))
	}
	return c
}

// AddNFT adds an NFT to the collection
func (c Collection) AddNFT(nft BaseNFT) Collection {
	c.NFTs = append(c.NFTs, nft)
	return c
}

func (c Collection) Supply() int {
	return len(c.NFTs)
}

// String follows stringer interface
func (c Collection) String() string {
	var buf bytes.Buffer
	for _, nft := range c.NFTs {
		if buf.Len() > 0 {
			buf.WriteString("\n")
		}
		buf.WriteString(nft.String())
	}
	return fmt.Sprintf(`Denom: 				%s
NFTs:

%s`,
		c.Denom,
		buf.String(),
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
