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

func (k Keeper) Authorize(ctx sdk.Context, denom, id string, owner sdk.AccAddress) (exported.NFT, error) {
	nft, err := k.GetNFT(ctx, denom, id)
	if err != nil {
		return nil, err
	}
	if !owner.Equals(nft.GetOwner()) {
		return nil, sdkerrors.Wrap(types.ErrUnauthorized, owner.String())
	}
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

		//denom may have the same prefix
		dstDenom, id, _ := types.SplitKeyNFT(iterator.Key())
		if dstDenom == denom && id == nft.GetID() {
			nfts = append(nfts, nft)
		}
	}
	return nfts
}

// UpdateNFT updates an already existing NFTs
func (k Keeper) UpdateNFT(ctx sdk.Context, denom string, nft exported.NFT) (err error) {
	oldNFT, err := k.GetNFT(ctx, denom, nft.GetID())
	if err != nil {
		return err
	}
	// if the owner changed then update the owners KVStore
	if !oldNFT.GetOwner().Equals(nft.GetOwner()) {
		k.removeOwner(ctx, denom, oldNFT)
	}
	k.SetNFT(ctx, denom, nft)
	return nil
}

// MintNFT mints an NFT and manages that NFTs existence within Collections and Owners
func (k Keeper) MintNFT(ctx sdk.Context, denom string, nft exported.NFT) error {
	if k.HasNFT(ctx, denom, nft.GetID()) {
		return sdkerrors.Wrapf(types.ErrNFTAlreadyExists, "NFT %s already exists in collection %s", nft.GetID(), denom)
	}
	k.SetNFT(ctx, denom, nft)
	return nil
}

func (k Keeper) SetNFT(ctx sdk.Context, denom string, nft exported.NFT) {
	store := ctx.KVStore(k.storeKey)

	bz := k.cdc.MustMarshalBinaryLengthPrefixed(nft)
	store.Set(types.KeyNFT(denom, nft.GetID()), bz)

	bzID := k.cdc.MustMarshalBinaryLengthPrefixed(nft.GetID())
	store.Set(types.KeyOwner(nft.GetOwner(), denom, nft.GetID()), bzID)
}

func (k Keeper) HasNFT(ctx sdk.Context, denom, id string) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has(types.KeyNFT(denom, id))
}

// DeleteNFT deletes an existing NFT from store
func (k Keeper) DeleteNFT(ctx sdk.Context, denom string, nft exported.NFT) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.KeyNFT(denom, nft.GetID()))
	k.removeOwner(ctx, denom, nft)
}

func (k Keeper) removeOwner(ctx sdk.Context, denom string, nft exported.NFT) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.KeyOwner(nft.GetOwner(), denom, nft.GetID()))
}
