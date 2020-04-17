package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/irismod/nft/types"
)

// GetOwner gets all the ID Collections owned by an address
func (k Keeper) GetOwner(ctx sdk.Context, address sdk.AccAddress) (owner types.Owner) {
	return k.GetOwnerOfDenom(ctx, address, "")
}

// GetOwner gets all the ID Collections owned by an address and denom
func (k Keeper) GetOwnerOfDenom(ctx sdk.Context, address sdk.AccAddress, denom string) types.Owner {
	idCs := make(map[string]types.IDCollection)
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.KeyOwner(address, denom, ""))
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
func (k Keeper) GetOwners(ctx sdk.Context) (owners types.Owners) {
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
		owner.IDCollections = owner.IDCollections.Add(denom, id)
		ownerMap[address.String()] = owner
	}
	for _, owner := range ownerMap {
		owners = append(owners, owner)
	}
	return owners
}

func (k Keeper) deleteOwner(ctx sdk.Context,
	denom, id string,
	owner sdk.AccAddress) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.KeyOwner(owner, denom, id))
}

func (k Keeper) setOwner(ctx sdk.Context,
	denom, id string,
	owner sdk.AccAddress) {
	store := ctx.KVStore(k.storeKey)
	bzID := k.cdc.MustMarshalBinaryLengthPrefixed(id)
	store.Set(types.KeyOwner(owner, denom, id), bzID)
}

func (k Keeper) swapOwner(ctx sdk.Context,
	denom, id string,
	srcOwner, dstOwner sdk.AccAddress) {

	//delete old owner key
	k.deleteOwner(ctx, denom, id, srcOwner)

	//set new owner key
	k.setOwner(ctx, denom, id, dstOwner)
}
