package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/irismod/nft/types"
)

// GetOwner gets all the TokenID Collections owned by an address and denom
func (k Keeper) GetOwner(ctx sdk.Context, address sdk.AccAddress, denom string) types.Owner {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.KeyOwner(address, denom, ""))
	defer iterator.Close()

	owner := types.Owner{
		Address:       address,
		IDCollections: types.IDCollections{},
	}
	idsMap := make(map[string][]string)

	for ; iterator.Valid(); iterator.Next() {
		_, denom, tokenID, _ := types.SplitKeyOwner(iterator.Key())
		if ids, ok := idsMap[denom]; ok {
			idsMap[denom] = append(ids, tokenID)
		} else {
			idsMap[denom] = []string{tokenID}
			owner.IDCollections = append(owner.IDCollections, types.IDCollection{
				Denom: denom,
			})
		}
	}

	for i := 0; i < len(owner.IDCollections); i++ {
		owner.IDCollections[i].IDs = idsMap[owner.IDCollections[i].Denom]
	}
	return owner
}

// GetOwner gets all the TokenID Collections
func (k Keeper) GetOwners(ctx sdk.Context) (owners types.Owners) {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStoreReversePrefixIterator(store, types.KeyOwner(nil, "", ""))
	defer iterator.Close()

	idcsMap := make(map[string]types.IDCollections)
	for ; iterator.Valid(); iterator.Next() {
		key := iterator.Key()
		address, denom, id, _ := types.SplitKeyOwner(key)
		if _, ok := idcsMap[address.String()]; !ok {
			idcsMap[address.String()] = types.IDCollections{}
			owners = append(owners, types.Owner{
				Address: address,
			})
		}
		idcs := idcsMap[address.String()]
		idcs = idcs.Add(denom, id)
		idcsMap[address.String()] = idcs
	}
	for i, owner := range owners {
		owners[i].IDCollections = idcsMap[owner.Address.String()]
	}
	return owners
}

func (k Keeper) deleteOwner(ctx sdk.Context,
	denom, tokenID string,
	owner sdk.AccAddress) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.KeyOwner(owner, denom, tokenID))
}

func (k Keeper) setOwner(ctx sdk.Context,
	denom, tokenID string,
	owner sdk.AccAddress) {
	store := ctx.KVStore(k.storeKey)
	bzID := k.cdc.MustMarshalBinaryLengthPrefixed(tokenID)
	store.Set(types.KeyOwner(owner, denom, tokenID), bzID)
}

func (k Keeper) swapOwner(ctx sdk.Context,
	denom, tokenID string,
	srcOwner, dstOwner sdk.AccAddress) {

	//delete old owner key
	k.deleteOwner(ctx, denom, tokenID, srcOwner)

	//set new owner key
	k.setOwner(ctx, denom, tokenID, dstOwner)
}
