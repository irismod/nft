package types

import (
	"bytes"
	"fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// NewIDCollection creates a new IDCollection instance
func NewIDCollection(denom string, ids []string) IDCollection {
	return IDCollection{
		Denom: strings.TrimSpace(denom),
		IDs:   ids,
	}
}

// Supply return the amount of the denom
func (idc IDCollection) Supply() int {
	return len(idc.IDs)
}

// AddID adds an TokenID to the idCollection
func (idc IDCollection) AddID(id string) IDCollection {
	idc.IDs = append(idc.IDs, id)
	return idc
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
// IDCollections is an array of TokenID Collections
type IDCollections []IDCollection

// Add adds an TokenID to the idCollection
func (idcs IDCollections) Add(denom, id string) IDCollections {
	for i, idc := range idcs {
		if idc.Denom == denom {
			idcs[i] = idc.AddID(id)
			return idcs
		}
	}
	return append(idcs, IDCollection{
		Denom: denom,
		IDs:   []string{id},
	})
}

// String follows stringer interface
func (idcs IDCollections) String() string {
	if len(idcs) == 0 {
		return ""
	}

	var buf bytes.Buffer
	for _, idCollection := range idcs {
		if buf.Len() > 0 {
			buf.WriteString("\n")
		}
		buf.WriteString(idCollection.String())
	}
	return buf.String()
}

// Owner of non fungible tokens
//type Owner struct {
//	Address       sdk.AccAddress `json:"address" yaml:"address"`
//	IDCollections IDCollections  `json:"id_collections" yaml:"id_collections"`
//}

// NewOwner creates a new Owner
func NewOwner(owner sdk.AccAddress, idCollections ...IDCollection) Owner {
	return Owner{
		Address:       owner,
		IDCollections: idCollections,
	}
}

// String follows stringer interface
func (owner Owner) String() string {
	var buf bytes.Buffer
	for _, idCollection := range owner.IDCollections {
		if buf.Len() > 0 {
			buf.WriteString("\n")
		}
		buf.WriteString(idCollection.String())
	}
	return fmt.Sprintf(`
	Address: 				%s
	IDCollections:        	%s`,
		owner.Address,
		buf.String(),
	)
}

type Owners []Owner

// NewOwner creates a new Owner
func NewOwners(owner ...Owner) Owners {
	return append([]Owner{}, owner...)
}

// String follows stringer interface
func (owners Owners) String() string {
	var buf bytes.Buffer
	for _, owner := range owners {
		if buf.Len() > 0 {
			buf.WriteString("\n")
		}
		buf.WriteString(owner.String())
	}
	return buf.String()
}
