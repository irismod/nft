package keeper

import (
	"fmt"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/tendermint/tendermint/libs/log"

	"github.com/cosmos/cosmos-sdk/codec"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/irismod/nft/types"
)

// Keeper maintains the link to data storage and exposes getter/setter methods for the various parts of the state machine
type Keeper struct {
	storeKey sdk.StoreKey // Unexposed key to access store from sdk.Context

	cdc *codec.Codec // The amino codec for binary encoding/decoding.
}

// NewKeeper creates new instances of the nft Keeper
func NewKeeper(cdc *codec.Codec, storeKey sdk.StoreKey) Keeper {
	return Keeper{
		storeKey: storeKey,
		cdc:      cdc,
	}
}

// Logger returns a module-specific logger.
func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("irismod/%s", types.ModuleName))
}

// MintNFT mints an NFT and manages that NFTs existence within Collections and Owners
func (k Keeper) MintNFT(ctx sdk.Context,
	denom, id, tokenURI string,
	owner sdk.AccAddress) error {
	if k.HasNFT(ctx, denom, id) {
		return sdkerrors.Wrapf(types.ErrNFTAlreadyExists, "NFT %s already exists in collection %s", id, denom)
	}
	nft := types.NewBaseNFT(id, owner, tokenURI)
	k.setNFT(ctx, denom, &nft)
	k.setOwner(ctx, denom, id, owner)
	k.increaseSupply(ctx, denom)
	return nil
}

// EditNFT updates an already existing NFTs
func (k Keeper) EditNFT(ctx sdk.Context,
	denom, id, tokenURI string,
	owner sdk.AccAddress) error {
	nft, err := k.Authorize(ctx, denom, id, owner)
	if err != nil {
		return err
	}

	nft.SetTokenURI(tokenURI)
	k.setNFT(ctx, denom, nft)
	return nil
}

// TransferOwner gets all the ID Collections owned by an address
func (k Keeper) TransferOwner(ctx sdk.Context,
	denom, id, tokenURI string,
	srcOwner, dstOwner sdk.AccAddress) error {
	nft, err := k.Authorize(ctx, denom, id, srcOwner)
	if err != nil {
		return err
	}

	nft.SetOwner(dstOwner)
	if tokenURI != types.DoNotModify {
		nft.SetTokenURI(tokenURI)
	}

	k.setNFT(ctx, denom, nft)
	k.swapOwner(ctx, denom, id, srcOwner, dstOwner)
	return nil
}

// BurnNFT delete a specified nft
func (k Keeper) BurnNFT(ctx sdk.Context,
	denom, id string,
	owner sdk.AccAddress) error {
	nft, err := k.Authorize(ctx, denom, id, owner)
	if err != nil {
		return err
	}

	k.deleteNFT(ctx, denom, nft)
	k.deleteOwner(ctx, denom, id, owner)
	k.decreaseSupply(ctx, denom)
	return nil
}
