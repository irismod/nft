package types

import (
	"bytes"
	"fmt"
	"sort"
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

// AddID adds an ID to the idCollection
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
// Owners

// IDCollections is an array of ID Collections whose sole purpose is for find
type IDCollections []IDCollection

// AddID adds an ID to the idCollection
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

func (idcs IDCollections) Len() int           { return len(idcs) }
func (idcs IDCollections) Swap(i, j int)      { idcs[i], idcs[j] = idcs[j], idcs[i] }
func (idcs IDCollections) Less(i, j int) bool { return idcs[i].Denom < idcs[j].Denom }
func (idcs IDCollections) Asc()               { sort.Sort(idcs) }
func (idcs IDCollections) Dsc()               { sort.Sort(sort.Reverse(idcs)) }

// String follows stringer interface
func (idcs IDCollections) String() string {
	if len(idcs) == 0 {
		return ""
	}

	out := ""
	sort.Sort(idcs)
	for _, idCollection := range idcs {
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
	owner.IDCollections.Asc()
	return fmt.Sprintf(`
	Address: 				%s
	IDCollections:        	%s`,
		owner.Address,
		owner.IDCollections.String(),
	)
}

type Owners []Owner

// NewOwner creates a new Owner
func NewOwners(owner ...Owner) Owners {
	return append([]Owner{}, owner...)
}

func (owners Owners) Len() int      { return len(owners) }
func (owners Owners) Swap(i, j int) { owners[i], owners[j] = owners[j], owners[i] }
func (owners Owners) Less(i, j int) bool {
	return owners[i].Address.String() < owners[j].Address.String()
}
func (owners Owners) Asc() { sort.Sort(owners) }
func (owners Owners) Dsc() { sort.Sort(sort.Reverse(owners)) }

// String follows stringer interface
func (owners Owners) String() string {
	owners.Asc()
	var buf bytes.Buffer
	for _, owner := range owners {
		buf.WriteString(owner.String())
		buf.WriteString("\n")
	}
	return buf.String()
}
