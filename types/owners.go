package types

import (
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

// AddID adds an ID to the idCollection
func (idc IDCollection) AddID(id string) IDCollection {
	idc.IDs = append(idc.IDs, id)
	return idc
}

// ----------------------------------------------------------------------------
// IDCollections is an array of ID Collections
type IDCollections []IDCollection

// Add adds an ID to the idCollection
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

// NewOwner creates a new Owner
func NewOwner(owner sdk.AccAddress, idCollections ...IDCollection) Owner {
	return Owner{
		Address:       owner,
		IDCollections: idCollections,
	}
}

type Owners []Owner

// NewOwner creates a new Owner
func NewOwners(owner ...Owner) Owners {
	return append([]Owner{}, owner...)
}