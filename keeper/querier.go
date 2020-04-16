package keeper

import (
	"encoding/binary"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/irismod/nft/types"

	abci "github.com/tendermint/tendermint/abci/types"
)

// query endpoints supported by the NFT Querier
const (
	QuerySupply       = "supply"
	QueryOwner        = "owner"
	QueryOwnerByDenom = "ownerByDenom"
	QueryCollection   = "collection"
	QueryDenoms       = "denoms"
	QueryNFT          = "nft"
)

// NewQuerier is the module level router for state queries
func NewQuerier(k Keeper) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) (res []byte, err error) {
		switch path[0] {
		case QuerySupply:
			return querySupply(ctx, req, k)
		case QueryOwner:
			return queryOwner(ctx, req, k)
		case QueryOwnerByDenom:
			return queryOwnerByDenom(ctx, req, k)
		case QueryCollection:
			return queryCollection(ctx, req, k)
		case QueryDenoms:
			return queryDenoms(ctx, req, k)
		case QueryNFT:
			return queryNFT(ctx, req, k)
		default:
			return nil, sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "unknown query path: %s", path[0])
		}
	}
}

func querySupply(ctx sdk.Context, req abci.RequestQuery, k Keeper) ([]byte, error) {
	var params types.QueryCollectionParams

	err := types.ModuleCdc.UnmarshalJSON(req.Data, &params)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}
	bz := make([]byte, 8)
	binary.LittleEndian.PutUint64(bz, k.GetSupply(ctx, params.Denom))
	return bz, nil
}

func queryOwner(ctx sdk.Context, req abci.RequestQuery, k Keeper) ([]byte, error) {
	var params types.QueryBalanceParams

	err := types.ModuleCdc.UnmarshalJSON(req.Data, &params)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}

	owner := k.GetOwner(ctx, params.Owner)
	bz, err := types.ModuleCdc.MarshalJSON(owner)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return bz, nil
}

func queryOwnerByDenom(ctx sdk.Context, req abci.RequestQuery, k Keeper) ([]byte, error) {
	var params types.QueryBalanceParams

	err := types.ModuleCdc.UnmarshalJSON(req.Data, &params)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}

	idCollection := k.GetOwnerByDenom(ctx, params.Owner, params.Denom)
	owner := types.Owner{
		Address:       params.Owner,
		IDCollections: types.IDCollections{idCollection},
	}

	bz, err := types.ModuleCdc.MarshalJSON(owner)
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

	collection, err := k.GetCollection(ctx, params.Denom)
	if err != nil {
		return nil, err
	}

	bz, err := types.ModuleCdc.MarshalJSON(collection)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return bz, nil
}

func queryDenoms(ctx sdk.Context, req abci.RequestQuery, k Keeper) ([]byte, error) {
	denoms := k.GetDenoms(ctx)

	bz, err := types.ModuleCdc.MarshalJSON(denoms)
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

	nft, err := k.GetNFT(ctx, params.Denom, params.TokenID)
	if err != nil {
		return nil, sdkerrors.Wrapf(types.ErrUnknownNFT, "invalid NFT %s from collection %s", params.TokenID, params.Denom)
	}

	bz, err := types.ModuleCdc.MarshalJSON(nft)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return bz, nil
}
