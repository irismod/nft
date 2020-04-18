package types

import (
	"bytes"
	"fmt"
	"github.com/irismod/nft/exported"
	"sort"
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
func (cs Collections) Len() int           { return len(cs) }
func (cs Collections) Swap(i, j int)      { cs[i], cs[j] = cs[j], cs[i] }
func (cs Collections) Less(i, j int) bool { return cs[i].Denom < cs[j].Denom }
func (cs Collections) Asc()               { sort.Sort(cs) }
func (cs Collections) Dsc()               { sort.Sort(sort.Reverse(cs)) }

// String follows stringer interface
func (cs Collections) String() string {
	if len(cs) == 0 {
		return ""
	}
	var buf bytes.Buffer
	for _, collection := range cs {
		buf.WriteString(collection.String())
		buf.WriteString("\n")
	}
	return buf.String()
}
