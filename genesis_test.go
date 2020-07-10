package nft_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/irismod/nft"
	"github.com/irismod/nft/types"
)

func TestInitGenesis(t *testing.T) {
	app, ctx := createTestApp(false)
	genesisState := nft.DefaultGenesisState()
	require.Equal(t, 0, len(genesisState.Collections))

	nft1 := types.NewBaseNFT(id, address, tokenURI1, tokenData)
	nft2 := types.NewBaseNFT(id2, address, tokenURI1, tokenData)
	nft3 := types.NewBaseNFT(id3, address, tokenURI1, tokenData)
	nfts := types.NewNFTs(&nft1, &nft2, &nft3)
	collection := types.NewCollection(types.Denom{
		Name:    denom,
		Schema:  "",
		Creator: address,
	}, nfts)

	nftx := types.NewBaseNFT(id, address2, tokenURI1, tokenData)
	nft2x := types.NewBaseNFT(id2, address2, tokenURI1, tokenData)
	nft3x := types.NewBaseNFT(id3, address2, tokenURI1, tokenData)
	nftsx := types.NewNFTs(&nftx, &nft2x, &nft3x)
	collection2 := types.NewCollection(types.Denom{
		Name:    denom2,
		Schema:  "",
		Creator: address,
	}, nftsx)

	collections := types.Collections{
		collection, collection2,
	}

	genesisState = types.NewGenesisState(collections)

	nft.InitGenesis(ctx, app.NFTKeeper, genesisState)

	returnedOwners := app.NFTKeeper.GetOwners(ctx)
	require.Equal(t, 2, len(returnedOwners))

	returnedCollections := app.NFTKeeper.GetCollections(ctx)
	require.Equal(t, 2, len(returnedCollections))
	require.Equal(t, returnedCollections.String(), collections.String())

	exportedGenesisState := nft.ExportGenesis(ctx, app.NFTKeeper)
	require.Equal(t, len(genesisState.Collections), len(exportedGenesisState.Collections))
	require.Equal(t, genesisState.Collections.String(), exportedGenesisState.Collections.String())
}
