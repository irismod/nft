package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/irismod/nft/exported"
	"github.com/irismod/nft/types"
)

// GetNFT gets the entire NFT tokenData struct
func (k Keeper) GetNFT(ctx sdk.Context, denom, tokenID string) (nft exported.NFT, err error) {
	store := ctx.KVStore(k.storeKey)

	bz := store.Get(types.KeyNFT(denom, tokenID))
	if bz == nil {
		return nil, sdkerrors.Wrapf(types.ErrUnknownCollection, "not found NFT: %s", denom)
	}

	var baseNFT types.BaseNFT
	k.cdc.MustUnmarshalBinaryBare(bz, &baseNFT)
	return baseNFT, nil
}

// GetNFTs return the all NFT by the specified denom
func (k Keeper) GetNFTs(ctx sdk.Context, denom string) (nfts []exported.NFT) {
	store := ctx.KVStore(k.storeKey)

	iterator := sdk.KVStorePrefixIterator(store, types.KeyNFT(denom, ""))
	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		var baseNFT types.BaseNFT
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &baseNFT)
		nfts = append(nfts, baseNFT)
	}
	return nfts
}

//Authorize check if the sender is the issuer of nft, if it returns nft, if not, return an error
func (k Keeper) Authorize(ctx sdk.Context,
	denom, tokenID string,
	owner sdk.AccAddress) (types.BaseNFT, error) {
	nft, err := k.GetNFT(ctx, denom, tokenID)
	if err != nil {
		return types.BaseNFT{}, err
	}

	if !owner.Equals(nft.GetOwner()) {
		return types.BaseNFT{}, sdkerrors.Wrap(types.ErrUnauthorized, owner.String())
	}
	return nft.(types.BaseNFT), nil
}

//HasNFT determine if nft exists
func (k Keeper) HasNFT(ctx sdk.Context, denom, tokenID string) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has(types.KeyNFT(denom, tokenID))
}

func (k Keeper) setNFT(ctx sdk.Context, denom string, nft types.BaseNFT) {
	store := ctx.KVStore(k.storeKey)

	bz := k.cdc.MustMarshalBinaryBare(&nft)
	store.Set(types.KeyNFT(denom, nft.GetID()), bz)
}

// deleteNFT deletes an existing NFT from store
func (k Keeper) deleteNFT(ctx sdk.Context, denom string, nft exported.NFT) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.KeyNFT(denom, nft.GetID()))
}
