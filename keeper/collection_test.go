package keeper_test

import (
	"github.com/irismod/nft/keeper"
	"github.com/irismod/nft/types"
)

func (suite *KeeperSuite) TestSetCollection() {
	nft := types.NewBaseNFT(id, address, tokenURI)
	// create a new NFT and add it to the collection created with the NFT mint
	nft2 := types.NewBaseNFT(id2, address, tokenURI)

	collection2 := types.Collection{
		Denom: denom,
		NFTs:  types.NFTs{&nft2, &nft},
	}
	err := suite.keeper.SetCollection(suite.ctx, collection2)
	suite.Nil(err)

	collection2, err = suite.keeper.GetCollection(suite.ctx, denom)
	suite.NoError(err)
	suite.Len(collection2.NFTs, 2)

	msg, fail := keeper.SupplyInvariant(suite.keeper)(suite.ctx)
	suite.False(fail, msg)
}

func (suite *KeeperSuite) TestGetCollection() {
	collection, err := suite.keeper.GetCollection(suite.ctx, denom)
	suite.Error(err)

	// MintNFT shouldn't fail when collection does not exist
	err = suite.keeper.MintNFT(suite.ctx, denom, id, tokenURI, address)
	suite.NoError(err)

	// collection should exist
	collection, err = suite.keeper.GetCollection(suite.ctx, denom)
	suite.NoError(err)
	suite.NotEmpty(collection)

	msg, fail := keeper.SupplyInvariant(suite.keeper)(suite.ctx)
	suite.False(fail, msg)
}

func (suite *KeeperSuite) TestGetCollections() {

	// collections should be empty
	collections := suite.keeper.GetCollections(suite.ctx)
	suite.Empty(collections)

	// MintNFT shouldn't fail when collection does not exist
	err := suite.keeper.MintNFT(suite.ctx, denom, id, tokenURI, address)
	suite.NoError(err)

	// collections should equal 1
	collections = suite.keeper.GetCollections(suite.ctx)
	suite.NotEmpty(collections)
	suite.Equal(len(collections), 1)

	msg, fail := keeper.SupplyInvariant(suite.keeper)(suite.ctx)
	suite.False(fail, msg)
}
