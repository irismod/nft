package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/irismod/nft/types"
)

// GetOwner gets all the ID Collections owned by an address
func (k Keeper) GetOwner(ctx sdk.Context, address sdk.AccAddress) (owner types.Owner) {
	idCs := make(map[string]types.IDCollection)
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.KeyOwner(address, "", ""))
	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		_, denom, id, _ := types.SplitKeyOwner(iterator.Key())
		if idc, ok := idCs[denom]; ok {
			idc.IDs = append(idc.IDs, id)
		} else {
			idc := types.IDCollection{
				Denom: denom,
				IDs:   []string{id},
			}
			idCs[denom] = idc
		}
	}

	idCollections := make([]types.IDCollection, len(idCs))
	i := 0
	for _, idc := range idCs {
		idCollections[i] = idc
		i++
	}
	return types.NewOwner(address, idCollections...)
}

// GetOwner gets all the ID Collections
func (k Keeper) GetOwners(ctx sdk.Context) (owners []types.Owner) {
	ownerMap := make(map[string]types.Owner)

	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.KeyOwner(nil, "", ""))
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		key := iterator.Key()
		address, denom, id, _ := types.SplitKeyOwner(key)
		if _, ok := ownerMap[address.String()]; !ok {
			ownerMap[address.String()] = types.Owner{
				Address:       address,
				IDCollections: []types.IDCollection{},
			}
		}
		owner := ownerMap[address.String()]
		idCs := owner.IDCollections
		idCs = append(idCs, types.IDCollection{
			Denom: denom,
			IDs:   []string{id},
		})
		owner.IDCollections = idCs
		ownerMap[address.String()] = owner
	}
	for _, owner := range ownerMap {
		owners = append(owners, owner)
	}
	return owners
}

// GetOwnerByDenom gets the ID Collection owned by an address of a specific denom
func (k Keeper) GetOwnerByDenom(ctx sdk.Context, owner sdk.AccAddress, denom string) (idc types.IDCollection) {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.KeyOwner(owner, denom, ""))
	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		_, _, id, _ := types.SplitKeyOwner(iterator.Key())
		idc.IDs = append(idc.IDs, id)
	}
	idc.Denom = denom
	return idc
}
