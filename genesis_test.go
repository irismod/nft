package nft_test

import (
	"testing"

	"github.com/irismod/nft/types"

	"github.com/stretchr/testify/require"

	"github.com/irismod/nft"
)

func TestInitGenesis(t *testing.T) {
	app, ctx := createTestApp(false)
	genesisState := nft.DefaultGenesisState()
	require.Equal(t, 0, len(genesisState.Collections))

	nft1 := nft.NewBaseNFT(id, address, tokenURI1, metadata)
	nft2 := nft.NewBaseNFT(id2, address, tokenURI1, metadata)
	nft3 := nft.NewBaseNFT(id3, address, tokenURI1, metadata)
	nfts := nft.NewNFTs(&nft1, &nft2, &nft3)
	collection := nft.NewCollection(types.Denom{
		Name:    denom,
		Schema:  "",
		Creator: address,
	}, nfts)

	nftx := nft.NewBaseNFT(id, address2, tokenURI1, metadata)
	nft2x := nft.NewBaseNFT(id2, address2, tokenURI1, metadata)
	nft3x := nft.NewBaseNFT(id3, address2, tokenURI1, metadata)
	nftsx := nft.NewNFTs(&nftx, &nft2x, &nft3x)
	collection2 := nft.NewCollection(types.Denom{
		Name:    denom2,
		Schema:  "",
		Creator: address,
	}, nftsx)

	collections := nft.Collections{
		collection, collection2,
	}

	genesisState = nft.NewGenesisState(collections)

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
