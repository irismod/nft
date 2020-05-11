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
			nft.GetTokenData(),
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
	for _, denom := range k.GetDenoms(ctx) {
		nfts := k.GetNFTs(ctx, denom.Name)
		cs = append(cs, types.Collection{
			Denom: denom,
			NFTs:  nfts,
		})
	}
	return cs
}

// GetTotalSupply returns the number of nft by the specified denom
func (k Keeper) GetTotalSupply(ctx sdk.Context, denom string) uint64 {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.KeyCollection(denom))
	if len(bz) == 0 {
		return 0
	}
	return types.MustUnMarshalSupply(k.cdc, bz)
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

func (k Keeper) increaseSupply(ctx sdk.Context, denom string) {
	supply := k.GetTotalSupply(ctx, denom)
	supply++

	store := ctx.KVStore(k.storeKey)
	bz := types.MustMarshalSupply(k.cdc, supply)
	store.Set(types.KeyCollection(denom), bz)
}

func (k Keeper) decreaseSupply(ctx sdk.Context, denom string) {
	supply := k.GetTotalSupply(ctx, denom)
	supply--

	store := ctx.KVStore(k.storeKey)
	if supply == 0 {
		store.Delete(types.KeyCollection(denom))
		return
	}

	bz := types.MustMarshalSupply(k.cdc, supply)
	store.Set(types.KeyCollection(denom), bz)
}
