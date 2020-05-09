package keeper_test

import (
	"github.com/irismod/nft/keeper"
)

func (suite *KeeperSuite) TestGetOwners() {

	err := suite.keeper.MintNFT(suite.ctx, denom, id, tokenURI, metadata, address)
	suite.NoError(err)

	err = suite.keeper.MintNFT(suite.ctx, denom, id2, tokenURI, metadata, address2)
	suite.NoError(err)

	err = suite.keeper.MintNFT(suite.ctx, denom, id3, tokenURI, metadata, address3)
	suite.NoError(err)

	owners := suite.keeper.GetOwners(suite.ctx)
	suite.Equal(3, len(owners))

	err = suite.keeper.MintNFT(suite.ctx, denom2, id, tokenURI, metadata, address)
	suite.NoError(err)

	err = suite.keeper.MintNFT(suite.ctx, denom2, id2, tokenURI, metadata, address2)
	suite.NoError(err)

	err = suite.keeper.MintNFT(suite.ctx, denom2, id3, tokenURI, metadata, address3)
	suite.NoError(err)

	owners = suite.keeper.GetOwners(suite.ctx)
	suite.Equal(3, len(owners))

	msg, fail := keeper.SupplyInvariant(suite.keeper)(suite.ctx)
	suite.False(fail, msg)
}
