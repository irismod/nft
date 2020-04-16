package types

import (
	"fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// IDCollection defines a set of nft ids that belong to a specific
// collection
type IDCollection struct {
	Denom string   `json:"denom" yaml:"denom"`
	IDs   []string `json:"ids" yaml:"ids"`
}

// NewIDCollection creates a new IDCollection instance
func NewIDCollection(denom string, ids []string) IDCollection {
	return IDCollection{
		Denom: strings.TrimSpace(denom),
		IDs:   ids,
	}
}

func (idc IDCollection) Supply() int {
	return len(idc.IDs)
}

// String follows stringer interface
func (idc IDCollection) String() string {
	return fmt.Sprintf(`Denom: 			%s
IDs:        	%s`,
		idc.Denom,
		strings.Join(idc.IDs, ","),
	)
}

// ----------------------------------------------------------------------------
// Owners

// IDCollections is an array of ID Collections whose sole purpose is for find
type IDCollections []IDCollection

// String follows stringer interface
func (idCollections IDCollections) String() string {
	if len(idCollections) == 0 {
		return ""
	}

	out := ""
	for _, idCollection := range idCollections {
		out += fmt.Sprintf("%v\n", idCollection.String())
	}
	return out[:len(out)-1]
}

// Owner of non fungible tokens
type Owner struct {
	Address       sdk.AccAddress `json:"address" yaml:"address"`
	IDCollections IDCollections  `json:"idCollections" yaml:"idCollections"`
}

// NewOwner creates a new Owner
func NewOwner(owner sdk.AccAddress, idCollections ...IDCollection) Owner {
	return Owner{
		Address:       owner,
		IDCollections: idCollections,
	}
}

// String follows stringer interface
func (owner Owner) String() string {
	return fmt.Sprintf(`
	Address: 				%s
	IDCollections:        	%s`,
		owner.Address,
		owner.IDCollections.String(),
	)
}
