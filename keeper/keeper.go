package keeper

import (
	"fmt"
	"strings"

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

func (k Keeper) IssueDenom(ctx sdk.Context, name, schema string, creator sdk.AccAddress) error {
	name = strings.ToLower(strings.TrimSpace(name))
	schema = strings.ToLower(strings.TrimSpace(schema))

	denom := types.NewDenom(name, schema, creator)
	return k.SetDenom(ctx, denom)
}

// MintNFT mints an NFT and manages that NFTs existence within Collections and Owners
func (k Keeper) MintNFT(ctx sdk.Context,
	denom, tokenID, tokenURI,metadata string,
	owner sdk.AccAddress) error {
	denom = strings.ToLower(strings.TrimSpace(denom))
	tokenID = strings.ToLower(strings.TrimSpace(tokenID))
	tokenURI = strings.TrimSpace(tokenURI)

	if !k.HasDenom(ctx, denom) {
		return sdkerrors.Wrapf(types.ErrInvalidDenom, "denom %s not exists", denom)
	}

	if k.HasNFT(ctx, denom, tokenID) {
		return sdkerrors.Wrapf(types.ErrNFTAlreadyExists, "NFT %s already exists in collection %s", tokenID, denom)
	}
	nft := types.NewBaseNFT(tokenID, owner, tokenURI,metadata)
	k.setNFT(ctx, denom, &nft)
	k.setOwner(ctx, denom, tokenID, owner)
	k.increaseSupply(ctx, denom)
	return nil
}

// EditNFT updates an already existing NFTs
func (k Keeper) EditNFT(ctx sdk.Context,
	denom, tokenID, tokenURI,metadata string,
	owner sdk.AccAddress) error {
	denom = strings.ToLower(strings.TrimSpace(denom))
	tokenID = strings.ToLower(strings.TrimSpace(tokenID))
	tokenURI = strings.TrimSpace(tokenURI)

	if !k.HasDenom(ctx, denom) {
		return sdkerrors.Wrapf(types.ErrInvalidDenom, "denom %s not exists", denom)
	}

	nft, err := k.Authorize(ctx, denom, tokenID, owner)
	if err != nil {
		return err
	}

	nft.SetMetadata(metadata)
	nft.SetTokenURI(tokenURI)
	k.setNFT(ctx, denom, nft)
	return nil
}

// TransferOwner gets all the TokenID Collections owned by an address
func (k Keeper) TransferOwner(ctx sdk.Context,
	denom, tokenID, tokenURI,metadata string,
	srcOwner, dstOwner sdk.AccAddress) error {
	denom = strings.ToLower(strings.TrimSpace(denom))
	tokenID = strings.ToLower(strings.TrimSpace(tokenID))
	tokenURI = strings.TrimSpace(tokenURI)

	if !k.HasDenom(ctx, denom) {
		return sdkerrors.Wrapf(types.ErrInvalidDenom, "denom %s not exists", denom)
	}

	nft, err := k.Authorize(ctx, denom, tokenID, srcOwner)
	if err != nil {
		return err
	}

	nft.SetOwner(dstOwner)
	if tokenURI != types.DoNotModify {
		nft.SetTokenURI(tokenURI)
	}
	if metadata != types.DoNotModify {
		nft.SetMetadata(metadata)
	}

	k.setNFT(ctx, denom, nft)
	k.swapOwner(ctx, denom, tokenID, srcOwner, dstOwner)
	return nil
}

// BurnNFT delete a specified nft
func (k Keeper) BurnNFT(ctx sdk.Context,
	denom, tokenID string,
	owner sdk.AccAddress) error {
	denom = strings.ToLower(strings.TrimSpace(denom))
	tokenID = strings.ToLower(strings.TrimSpace(tokenID))

	if !k.HasDenom(ctx, denom) {
		return sdkerrors.Wrapf(types.ErrInvalidDenom, "denom %s not exists", denom)
	}

	nft, err := k.Authorize(ctx, denom, tokenID, owner)
	if err != nil {
		return err
	}

	k.deleteNFT(ctx, denom, nft)
	k.deleteOwner(ctx, denom, tokenID, owner)
	k.decreaseSupply(ctx, denom)
	return nil
}
