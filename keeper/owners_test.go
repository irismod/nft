package keeper_test

import (
	"github.com/irismod/nft/keeper"
)

func (suite *KeeperSuite) TestGetOwners() {

	err := suite.keeper.MintNFT(suite.ctx, denom, id, tokenURI, address)
	suite.NoError(err)

	err = suite.keeper.MintNFT(suite.ctx, denom, id2, tokenURI, address2)
	suite.NoError(err)

	err = suite.keeper.MintNFT(suite.ctx, denom, id3, tokenURI, address3)
	suite.NoError(err)

	owners := suite.keeper.GetOwners(suite.ctx)
	suite.Equal(3, len(owners))

	err = suite.keeper.MintNFT(suite.ctx, denom2, id, tokenURI, address)
	suite.NoError(err)

	err = suite.keeper.MintNFT(suite.ctx, denom2, id2, tokenURI, address2)
	suite.NoError(err)

	err = suite.keeper.MintNFT(suite.ctx, denom2, id3, tokenURI, address3)
	suite.NoError(err)

	owners = suite.keeper.GetOwners(suite.ctx)
	suite.Equal(3, len(owners))

	msg, fail := keeper.SupplyInvariant(suite.keeper)(suite.ctx)
	suite.False(fail, msg)
}

func (suite *KeeperSuite) TestGetOwner() {
	err := suite.keeper.MintNFT(suite.ctx, denom, id, tokenURI, address)
	suite.NoError(err)

	owner := suite.keeper.GetOwner(suite.ctx, address)
	suite.Len(owner.IDCollections, 1)
	suite.Len(owner.IDCollections[0].IDs, 1)
	suite.Equal(owner.IDCollections[0].IDs[0], id)
	suite.Equal(owner.Address, address)
}
