package nft_test

import (
	"fmt"

	abci "github.com/tendermint/tendermint/abci/types"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/irismod/nft"
	simapp "github.com/irismod/nft/app"
	"github.com/irismod/nft/types"
)

var (
	denom2    = "denom2"
	id        = "id1"
	id2       = "id2"
	id3       = "id3"
	address   = types.CreateTestAddrs(1)[0]
	address2  = types.CreateTestAddrs(2)[1]
	address3  = types.CreateTestAddrs(3)[2]
	tokenURI1 = "https://google.com/token-1.json"
	tokenURI2 = "https://google.com/token-2.json"
	tokenData = "{}"
)

func createTestApp(isCheckTx bool) (*simapp.SimApp, sdk.Context) {
	app := simapp.Setup(isCheckTx)
	ctx := app.BaseApp.NewContext(isCheckTx, abci.Header{})

	return app, ctx
}

// CheckInvariants checks the invariants
func CheckInvariants(k nft.Keeper, ctx sdk.Context) bool {
	collectionsSupply := make(map[string]int)
	ownersCollectionsSupply := make(map[string]int)

	for _, c := range k.GetCollections(ctx) {
		collectionsSupply[c.Denom.Name] = c.Supply()
	}

	owners := k.GetOwners(ctx)
	for _, owner := range owners {
		for _, idCollection := range owner.IDCollections {
			ownersCollectionsSupply[idCollection.Denom] += idCollection.Supply()
		}
	}

	for denom, supply := range collectionsSupply {
		if supply != ownersCollectionsSupply[denom] {
			fmt.Printf("denom is %s, supply is %d, ownerSupply is %d", denom, supply, ownersCollectionsSupply[denom])
			return false
		}
	}
	return true
}
