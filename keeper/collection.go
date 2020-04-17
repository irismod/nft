package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/irismod/nft/types"
)

// SetCollection save all NFT and panic if existed
func (k Keeper) SetCollection(ctx sdk.Context, collection types.Collection) error {
	for _, nft := range collection.NFTs {
		if err := k.MintNFT(ctx,
			collection.Denom,
			nft.GetID(),
			nft.GetTokenURI(),
			nft.GetOwner(),
		); err != nil {
			return err
		}
	}
	return nil
}

// GetSupply returns the number of nft by the specified  collection
func (k Keeper) GetSupply(ctx sdk.Context, denom string) uint64 {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.KeyCollection(denom))

	var supply uint64
	k.cdc.MustUnmarshalBinaryLengthPrefixed(bz, &supply)
	return supply
}

// GetSupplyOf returns the number of nft by the specified conditions
func (k Keeper) GetSupplyOf(ctx sdk.Context, denom string, owner sdk.AccAddress) uint64 {
	if owner == nil {
		return k.GetSupply(ctx, denom)
	}

	idc := k.GetTokenIDsOfDenom(ctx, owner, denom)
	return uint64(len(idc.IDs))
}

// GetCollection returns the collection by the specified denom
func (k Keeper) GetCollection(ctx sdk.Context, denom string) (types.Collection, error) {
	nfts := k.GetNFTs(ctx, denom)
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

// IterateCollections iterate all the collection
func (k Keeper) IterateCollections(ctx sdk.Context, fn func(collection types.Collection)) {
	denoms := k.GetAllDenoms(ctx)
	for _, denom := range denoms {
		nfts := k.GetNFTs(ctx, denom)
		fn(types.Collection{
			Denom: denom,
			NFTs:  nfts,
		})
	}
}

// GetAllDenoms return the denoms of all the collection
func (k Keeper) GetAllDenoms(ctx sdk.Context) (denoms []string) {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, store.Get(types.KeyCollection("")))
	defer iterator.Close()

	var denomMap = make(map[string]int)
	for ; iterator.Valid(); iterator.Next() {
		key := iterator.Key()
		denom := types.SplitKeyCollection(key)
		denomMap[denom] = 1
	}
	for k, _ := range denomMap {
		denoms = append(denoms, k)
	}
	return denoms
}

func (k Keeper) increaseSupply(ctx sdk.Context, denom string) {
	nfts := k.GetNFTs(ctx, denom)
	supply := uint64(len(nfts)) + 1

	store := ctx.KVStore(k.storeKey)

	bzSupply := k.cdc.MustMarshalBinaryLengthPrefixed(supply)
	store.Set(types.KeyCollection(denom), bzSupply)
}

func (k Keeper) decreaseSupply(ctx sdk.Context, denom string) {
	nfts := k.GetNFTs(ctx, denom)
	supply := uint64(len(nfts)) - 1

	store := ctx.KVStore(k.storeKey)
	if supply <= 0 {
		store.Delete(types.KeyCollection(denom))
		return
	}

	bzSupply := k.cdc.MustMarshalBinaryLengthPrefixed(supply)
	store.Set(types.KeyCollection(denom), bzSupply)
}
