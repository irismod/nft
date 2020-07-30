package keeper

import (
	"encoding/binary"
	"strings"

	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/irismod/nft/types"
)

// NewQuerier is the module level router for state queries
func NewQuerier(k Keeper) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) (res []byte, err error) {
		switch path[0] {
		case types.QuerySupply:
			return querySupply(ctx, req, k)
		case types.QueryOwner:
			return queryOwner(ctx, req, k)
		case types.QueryCollection:
			return queryCollection(ctx, req, k)
		case types.QueryDenom:
			return queryDenom(ctx, req, k)
		case types.QueryDenoms:
			return queryDenoms(ctx, req, k)
		case types.QueryNFT:
			return queryNFT(ctx, req, k)
		default:
			return nil, sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "unknown query path: %s", path[0])
		}
	}
}

func querySupply(ctx sdk.Context, req abci.RequestQuery, k Keeper) ([]byte, error) {
	var params types.QuerySupplyParams

	err := types.ModuleCdc.UnmarshalJSON(req.Data, &params)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}

	denom := strings.ToLower(strings.TrimSpace(params.Denom))

	var supply uint64
	switch {
	case params.Owner.Empty() && len(denom) > 0:
		supply = k.GetTotalSupply(ctx, denom)
	default:
		supply = k.GetTotalSupplyOfOwner(ctx, denom, params.Owner)
	}

	bz := make([]byte, 8)
	binary.LittleEndian.PutUint64(bz, supply)
	return bz, nil
}

func queryOwner(ctx sdk.Context, req abci.RequestQuery, k Keeper) ([]byte, error) {
	var params types.QueryOwnerParams

	err := types.ModuleCdc.UnmarshalJSON(req.Data, &params)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}

	owner := k.GetOwner(ctx, params.Owner, params.Denom)
	bz, err := codec.MarshalJSONIndent(k.cdc, owner)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return bz, nil
}

func queryCollection(ctx sdk.Context, req abci.RequestQuery, k Keeper) ([]byte, error) {
	var params types.QueryCollectionParams

	err := types.ModuleCdc.UnmarshalJSON(req.Data, &params)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}

	denom := strings.ToLower(strings.TrimSpace(params.Denom))
	collection, err := k.GetCollection(ctx, denom)
	if err != nil {
		return nil, err
	}

	bz, err := codec.MarshalJSONIndent(k.cdc, collection)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return bz, nil
}

func queryDenom(ctx sdk.Context, req abci.RequestQuery, k Keeper) ([]byte, error) {
	var params types.QueryDenomParams

	err := types.ModuleCdc.UnmarshalJSON(req.Data, &params)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}

	denom, err := k.GetDenom(ctx, params.Denom)

	bz, err := codec.MarshalJSONIndent(k.cdc, denom)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return bz, nil
}

func queryDenoms(ctx sdk.Context, req abci.RequestQuery, k Keeper) ([]byte, error) {
	denoms := k.GetDenoms(ctx)

	bz, err := codec.MarshalJSONIndent(k.cdc, denoms)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return bz, nil
}

func queryNFT(ctx sdk.Context, req abci.RequestQuery, k Keeper) ([]byte, error) {
	var params types.QueryNFTParams

	err := types.ModuleCdc.UnmarshalJSON(req.Data, &params)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}

	denom := strings.ToLower(strings.TrimSpace(params.Denom))
	tokenID := strings.ToLower(strings.TrimSpace(params.TokenID))
	nft, err := k.GetNFT(ctx, denom, tokenID)
	if err != nil {
		return nil, sdkerrors.Wrapf(types.ErrUnknownNFT, "invalid NFT %s from collection %s", params.TokenID, params.Denom)
	}

	bz, err := codec.MarshalJSONIndent(k.cdc, nft)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return bz, nil
}