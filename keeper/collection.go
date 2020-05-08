package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/irismod/nft/types"
)

// SetCollection save all NFT and return error if existed
func (k Keeper) SetCollection(ctx sdk.Context, collection types.Collection) error {
	for _, nft := range collection.NFTs {
		if err := k.MintNFT(ctx,
			collection.Denom.Name,
			nft.GetID(),
			nft.GetTokenURI(),
			nft.GetMetadata(),
			nft.GetOwner(),
		); err != nil {
			return err
		}
	}
	return nil
}

// GetCollection returns the collection by the specified denom
func (k Keeper) GetCollection(ctx sdk.Context, denomNm string) (types.Collection, error) {
	denom, err := k.GetDenom(ctx, denomNm)
	if err != nil {
		return types.Collection{}, sdkerrors.Wrapf(types.ErrInvalidDenom, "denom %s not existed ", denom)
	}

	nfts := k.GetNFTs(ctx, denomNm)
	if len(nfts) == 0 {
		return types.Collection{}, sdkerrors.Wrapf(types.ErrUnknownCollection, "collection %s not existed ", denom)
	}

	return types.NewCollection(denom, nfts), nil
}

// GetCollections returns all the collection
func (k Keeper) GetCollections(ctx sdk.Context) (cs types.Collections) {
	k.IterateCollections(ctx, func(collection types.Collection) {
		cs = append(cs, collection)
	})
	return cs
}

// GetTotalSupply returns the number of nft by the specified denom
func (k Keeper) GetTotalSupply(ctx sdk.Context, denom string) uint64 {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.KeyCollection(denom))
	if len(bz) == 0 {
		return 0
	}

	var supply uint64
	k.cdc.MustUnmarshalBinaryLengthPrefixed(bz, &supply)
	return supply
}

// GetTotalSupplyOfOwner returns the amount of nft by the specified conditions
func (k Keeper) GetTotalSupplyOfOwner(ctx sdk.Context, denom string, owner sdk.AccAddress) (supply uint64) {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.KeyOwner(owner, denom, ""))
	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		supply++
	}
	return supply
}

// IterateCollections iterate all the collection
func (k Keeper) IterateCollections(ctx sdk.Context, fn func(collection types.Collection)) {
	for _, denom := range k.GetDenoms(ctx) {
		nfts := k.GetNFTs(ctx, denom.Name)
		fn(types.Collection{
			Denom: denom.Name,
			NFTs:  nfts,
		})
	}
}

// GetDenoms return the denoms of all the collection
//func (k Keeper) GetDenoms(ctx sdk.Context) (denoms []string) {
//	store := ctx.KVStore(k.storeKey)
//	iterator := sdk.KVStorePrefixIterator(store, types.KeyCollection(""))
//	defer iterator.Close()
//
//	var denomMap = make(map[string]int)
//	for ; iterator.Valid(); iterator.Next() {
//		key := iterator.Key()
//		denom := types.SplitKeyCollection(key)
//		if _, existed := denomMap[denom]; !existed {
//			denoms = append(denoms, denom)
//			denomMap[denom] = 1
//		}
//	}
//	return denoms
//}

func (k Keeper) increaseSupply(ctx sdk.Context, denom string) {
	supply := k.GetTotalSupply(ctx, denom)
	supply++

	bzSupply := k.cdc.MustMarshalBinaryLengthPrefixed(supply)
	store := ctx.KVStore(k.storeKey)
	store.Set(types.KeyCollection(denom), bzSupply)
}

func (k Keeper) decreaseSupply(ctx sdk.Context, denom string) {
	supply := k.GetTotalSupply(ctx, denom)
	supply--

	store := ctx.KVStore(k.storeKey)
	if supply == 0 {
		store.Delete(types.KeyCollection(denom))
		return
	}

	bzSupply := k.cdc.MustMarshalBinaryLengthPrefixed(supply)
	store.Set(types.KeyCollection(denom), bzSupply)
}
