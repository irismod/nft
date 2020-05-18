package rest

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/gorilla/mux"

	"github.com/cosmos/cosmos-sdk/client/context"
)

// RegisterRoutes register distribution REST routes.
func RegisterRoutes(cliCtx context.CLIContext, r *mux.Router, queryRoute string) {
	registerQueryRoutes(cliCtx, r, queryRoute)
	registerTxRoutes(cliCtx, r, queryRoute)
}

const (
	RestParamDenom   = "denom"
	RestParamTokenID = "token-id"
	RestParamOwner   = "owner"
)

type issueDenomReq struct {
	BaseReq rest.BaseReq   `json:"base_req"`
	Owner   sdk.AccAddress `json:"owner"`
	Denom   string         `json:"denom"`
	Schema  string         `json:"schema"`
}

type mintNFTReq struct {
	BaseReq   rest.BaseReq   `json:"base_req"`
	Owner     sdk.AccAddress `json:"owner"`
	Recipient sdk.AccAddress `json:"recipient"`
	Denom     string         `json:"denom"`
	TokenID   string         `json:"token_id"`
	TokenURI  string         `json:"token_uri"`
	TokenData string         `json:"token_data"`
}

type editNFTReq struct {
	BaseReq   rest.BaseReq   `json:"base_req"`
	Owner     sdk.AccAddress `json:"owner"`
	TokenURI  string         `json:"token_uri"`
	TokenData string         `json:"token_data"`
}

type transferNFTReq struct {
	BaseReq   rest.BaseReq   `json:"base_req"`
	Owner     sdk.AccAddress `json:"owner"`
	Recipient string         `json:"recipient"`
	TokenURI  string         `json:"token_uri"`
	TokenData string         `json:"token_data"`
}

type burnNFTReq struct {
	BaseReq rest.BaseReq   `json:"base_req"`
	Owner   sdk.AccAddress `json:"owner"`
}
