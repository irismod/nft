package rest

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"

	"github.com/irismod/nft/types"
)

func registerQueryRoutes(cliCtx context.CLIContext, r *mux.Router, cdc *codec.Codec, queryRoute string) {
	// Query a single NFT
	r.HandleFunc(
		fmt.Sprintf("/%s/{%s}/%s", types.ModuleName, RestParamDenom, RestParamTokenID),
		getNFT(cdc, cliCtx, queryRoute),
	).Methods("GET")
	// Get the total supply of a collection
	r.HandleFunc(
		fmt.Sprintf("/%s/supply", types.ModuleName),
		getSupply(cdc, cliCtx, queryRoute),
	).Methods("GET")

	// Get the collections of NFTs owned by an address
	r.HandleFunc(
		fmt.Sprintf("/%s/owner/%s", types.ModuleName, RestParamOwner),
		getOwner(cdc, cliCtx, queryRoute),
	).Methods("GET")

	// Get all the NFT from a given collection
	r.HandleFunc(
		fmt.Sprintf("/%s/collection/%s", types.ModuleName, RestParamDenom),
		getCollection(cdc, cliCtx, queryRoute),
	).Methods("GET")

	// Query all denoms
	r.HandleFunc(
		fmt.Sprintf("/%s/denoms", types.ModuleName),
		getDenoms(cliCtx, queryRoute),
	).Methods("GET")
}

func getNFT(cdc *codec.Codec, cliCtx context.CLIContext, queryRoute string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		denom := vars[RestParamDenom]
		id := vars[RestParamTokenID]

		params := types.NewQueryNFTParams(denom, id)
		bz, err := cdc.MarshalJSON(params)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		res, _, err := cliCtx.QueryWithData(
			fmt.Sprintf("custom/%s/%s", queryRoute, types.QueryNFT), bz)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		rest.PostProcessResponse(w, cliCtx, res)
	}
}

func getSupply(cdc *codec.Codec, cliCtx context.CLIContext, queryRoute string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		denom := r.FormValue(RestParamDenom)
		ownerStr := r.FormValue(RestParamOwner)

		owner, err := sdk.AccAddressFromBech32(ownerStr)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		params := types.NewQuerySupplyParams(denom, owner)
		bz, err := cdc.MarshalJSON(params)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		res, _, err := cliCtx.QueryWithData(
			fmt.Sprintf("custom/%s/%s", queryRoute, types.QuerySupply), bz)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		rest.PostProcessResponse(w, cliCtx, res)
	}
}

func getOwner(cdc *codec.Codec, cliCtx context.CLIContext, queryRoute string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		address, err := sdk.AccAddressFromBech32(mux.Vars(r)[RestParamOwner])
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		denom := r.FormValue(RestParamDenom)
		params := types.NewQuerySupplyParams(denom, address)
		bz, err := cdc.MarshalJSON(params)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		res, _, err := cliCtx.QueryWithData(
			fmt.Sprintf("custom/%s/%s", queryRoute, types.QueryOwner), bz)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		rest.PostProcessResponse(w, cliCtx, res)
	}
}

func getCollection(cdc *codec.Codec, cliCtx context.CLIContext, queryRoute string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		denom := mux.Vars(r)[RestParamDenom]

		params := types.NewQuerySupplyParams(denom, nil)
		bz, err := cdc.MarshalJSON(params)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		res, _, err := cliCtx.QueryWithData(
			fmt.Sprintf("custom/%s/%s", queryRoute, types.QueryCollection), bz)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		rest.PostProcessResponse(w, cliCtx, res)
	}
}

func getDenoms(cliCtx context.CLIContext, queryRoute string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res, _, err := cliCtx.QueryWithData(
			fmt.Sprintf("custom/%s/%s", queryRoute, types.QueryDenoms), nil)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		rest.PostProcessResponse(w, cliCtx, res)
	}
}
