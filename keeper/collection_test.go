package keeper_test

import (
	"github.com/irismod/nft/keeper"
	"github.com/irismod/nft/types"
)

func (suite *KeeperSuite) TestSetCollection() {
	nft := types.NewBaseNFT(id, address, tokenURI, metadata)
	// create a new NFT and add it to the collection created with the NFT mint
	nft2 := types.NewBaseNFT(id2, address, tokenURI, metadata)

	denomE := types.Denom{
		Name:    denom,
		Schema:  schema,
		Creator: address,
	}

	collection2 := types.Collection{
		Denom: denomE,
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
	_, err := suite.keeper.GetCollection(suite.ctx, denom)
	suite.Error(err)

	// MintNFT shouldn't fail when collection does not exist
	err = suite.keeper.MintNFT(suite.ctx, denom, id, tokenURI, metadata, address)
	suite.NoError(err)

	// collection should exist
	collection, err := suite.keeper.GetCollection(suite.ctx, denom)
	suite.NoError(err)
	suite.NotEmpty(collection)

	msg, fail := keeper.SupplyInvariant(suite.keeper)(suite.ctx)
	suite.False(fail, msg)
}

func (suite *KeeperSuite) TestGetCollections() {

	// MintNFT shouldn't fail when collection does not exist
	err := suite.keeper.MintNFT(suite.ctx, denom, id, tokenURI, metadata, address)
	suite.NoError(err)

	msg, fail := keeper.SupplyInvariant(suite.keeper)(suite.ctx)
	suite.False(fail, msg)
}

func (suite *KeeperSuite) TestGetSupply() {
	// MintNFT shouldn't fail when collection does not exist
	err := suite.keeper.MintNFT(suite.ctx, denom, id, tokenURI, metadata, address)
	suite.NoError(err)

	// MintNFT shouldn't fail when collection does not exist
	err = suite.keeper.MintNFT(suite.ctx, denom, id2, tokenURI, metadata, address2)
	suite.NoError(err)

	// MintNFT shouldn't fail when collection does not exist
	err = suite.keeper.MintNFT(suite.ctx, denom2, id, tokenURI, metadata, address2)
	suite.NoError(err)

	supply := suite.keeper.GetTotalSupply(suite.ctx, denom)
	suite.Equal(uint64(2), supply)

	supply = suite.keeper.GetTotalSupply(suite.ctx, denom2)
	suite.Equal(uint64(1), supply)

	supply = suite.keeper.GetTotalSupplyOfOwner(suite.ctx, denom, address)
	suite.Equal(uint64(1), supply)

	supply = suite.keeper.GetTotalSupplyOfOwner(suite.ctx, denom, address2)
	suite.Equal(uint64(1), supply)

	supply = suite.keeper.GetTotalSupply(suite.ctx, denom)
	suite.Equal(uint64(2), supply)

	supply = suite.keeper.GetTotalSupply(suite.ctx, denom2)
	suite.Equal(uint64(1), supply)

	//burn nft
	err = suite.keeper.BurnNFT(suite.ctx, denom, id, address)
	suite.NoError(err)

	supply = suite.keeper.GetTotalSupply(suite.ctx, denom)
	suite.Equal(uint64(1), supply)

	supply = suite.keeper.GetTotalSupply(suite.ctx, denom)
	suite.Equal(uint64(1), supply)

	//burn nft
	err = suite.keeper.BurnNFT(suite.ctx, denom, id2, address2)
	suite.NoError(err)

	supply = suite.keeper.GetTotalSupply(suite.ctx, denom)
	suite.Equal(uint64(0), supply)

	supply = suite.keeper.GetTotalSupply(suite.ctx, denom)
	suite.Equal(uint64(0), supply)
}
