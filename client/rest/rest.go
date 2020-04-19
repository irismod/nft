package rest

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/gorilla/mux"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
)

// RegisterRoutes register distribution REST routes.
func RegisterRoutes(cliCtx context.CLIContext, r *mux.Router, cdc *codec.Codec, queryRoute string) {
	registerQueryRoutes(cliCtx, r, cdc, queryRoute)
	registerTxRoutes(cliCtx, r, cdc, queryRoute)
}

const (
	RestParamDenom = "denom"
	RestParamID    = "id"
	RestParamOwner = "owner"
)

type mintNFTReq struct {
	BaseTx    rest.BaseReq   `json:"base_tx"`
	Recipient sdk.AccAddress `json:"recipient"`
	Denom     string         `json:"denom"`
	ID        string         `json:"id"`
	TokenURI  string         `json:"tokenURI"`
}

type editNFTReq struct {
	BaseTx   rest.BaseReq `json:"base_tx"`
	TokenURI string       `json:"tokenURI"`
}

type transferNFTReq struct {
	BaseTx    rest.BaseReq `json:"base_tx"`
	Recipient string       `json:"recipient"`
	TokenURI  string       `json:"tokenURI"`
}

type burnNFTReq struct {
	BaseTx rest.BaseReq `json:"base_tx"`
}
