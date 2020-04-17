package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/irismod/nft/exported"
	"github.com/irismod/nft/types"
)

// GetNFT gets the entire NFT metadata struct for a uint64
func (k Keeper) GetNFT(ctx sdk.Context, denom, id string) (nft exported.NFT, err error) {
	store := ctx.KVStore(k.storeKey)

	bz := store.Get(types.KeyNFT(denom, id))
	if bz == nil || len(bz) == 0 {
		return nil, sdkerrors.Wrapf(types.ErrUnknownCollection, "not found NFT: %s", denom)
	}

	k.cdc.MustUnmarshalBinaryLengthPrefixed(bz, &nft)
	return nft, nil
}

// GetNFTs return the all NFT by the specified denom
func (k Keeper) GetNFTs(ctx sdk.Context, denom string) (nfts []exported.NFT) {
	store := ctx.KVStore(k.storeKey)

	iterator := sdk.KVStorePrefixIterator(store, types.KeyNFT(denom, ""))
	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		var nft exported.NFT
		k.cdc.MustUnmarshalBinaryLengthPrefixed(iterator.Value(), &nft)
		nfts = append(nfts, nft)
	}
	return nfts
}

//Authorize check if the sender is the issuer of nft, if it returns nft, if not, return an error
func (k Keeper) Authorize(ctx sdk.Context,
	denom, id string,
	owner sdk.AccAddress) (exported.NFT, error) {
	nft, err := k.GetNFT(ctx, denom, id)
	if err != nil {
		return nil, err
	}

	if !owner.Equals(nft.GetOwner()) {
		return nil, sdkerrors.Wrap(types.ErrUnauthorized, owner.String())
	}
	return nft, nil
}

//HasNFT determine if nft exists
func (k Keeper) HasNFT(ctx sdk.Context, denom, id string) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has(types.KeyNFT(denom, id))
}

func (k Keeper) setNFT(ctx sdk.Context, denom string, nft exported.NFT) {
	store := ctx.KVStore(k.storeKey)

	bz := k.cdc.MustMarshalBinaryLengthPrefixed(nft)
	store.Set(types.KeyNFT(denom, nft.GetID()), bz)
}

// deleteNFT deletes an existing NFT from store
func (k Keeper) deleteNFT(ctx sdk.Context, denom string, nft exported.NFT) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.KeyNFT(denom, nft.GetID()))
}
