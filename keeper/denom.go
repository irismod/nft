package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/irismod/nft/types"
)

// HasDenom returns whether the specified denom exists
func (k Keeper) HasDenom(ctx sdk.Context, denom string) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has(types.KeyDenom(denom))
}

// SetDenom is responsible for saving the definition of denom
func (k Keeper) SetDenom(ctx sdk.Context, denom types.Denom) error {
	if k.HasDenom(ctx, denom.Name) {
		return sdkerrors.Wrapf(types.ErrInvalidDenom, "denom %s has already exists", denom.Name)
	}

	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshalBinaryLengthPrefixed(denom)
	store.Set(types.KeyDenom(denom.Name), bz)
	return nil
}

// SetDenom is responsible for saving the definition of denom
func (k Keeper) GetDenom(ctx sdk.Context, name string) (denom types.Denom, err error) {
	store := ctx.KVStore(k.storeKey)

	bz := store.Get(types.KeyDenom(name))
	if bz == nil || len(bz) == 0 {
		return denom, sdkerrors.Wrapf(types.ErrInvalidDenom, "not found denom: %s", name)
	}

	k.cdc.MustUnmarshalBinaryLengthPrefixed(bz, &denom)
	return denom, nil
}

// GetDenoms return all the denoms
func (k Keeper) GetDenoms(ctx sdk.Context) (denoms []types.Denom) {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.KeyDenom(""))
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var denom types.Denom
		k.cdc.MustUnmarshalBinaryLengthPrefixed(iterator.Value(), &denom)
		denoms = append(denoms, denom)
	}
	return denoms
}
