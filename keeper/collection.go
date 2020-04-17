package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/irismod/nft/types"
)

// SetCollection save all NFT and panic if existed
func (k Keeper) SetCollection(ctx sdk.Context, collection types.Collection) error {
	for _, nft := range collection.NFTs {
		if err := k.MintNFT(ctx, collection.Denom, nft); err != nil {
			return err
		}
	}
	return nil
}

// GetSupply returns the number of nft under collection
func (k Keeper) GetSupply(ctx sdk.Context, denom string) uint64 {
	nfts := k.GetNFTs(ctx, denom)
	return uint64(len(nfts))
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
	denoms := k.GetDenoms(ctx)
	for _, denom := range denoms {
		nfts := k.GetNFTs(ctx, denom)
		fn(types.Collection{
			Denom: denom,
			NFTs:  nfts,
		})
	}
}

// GetDenoms return the denoms of all the collection
func (k Keeper) GetDenoms(ctx sdk.Context) (denoms []string) {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.KeyNFT("", ""))
	defer iterator.Close()

	var denomMap = make(map[string]int)
	for ; iterator.Valid(); iterator.Next() {
		key := iterator.Key()
		denom, _, _ := types.SplitKeyNFT(key)
		denomMap[denom] = 1
	}
	for k, _ := range denomMap {
		denoms = append(denoms, k)
	}
	return denoms
}
