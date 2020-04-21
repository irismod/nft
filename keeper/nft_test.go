package keeper_test

import (
	"github.com/irismod/nft/keeper"
)

func (suite *KeeperSuite) TestGetNFT() {
	// MintNFT shouldn't fail when collection does not exist
	err := suite.keeper.MintNFT(suite.ctx, denom, id, tokenURI, address)
	suite.NoError(err)

	// GetNFT should get the NFT
	receivedNFT, err := suite.keeper.GetNFT(suite.ctx, denom, id)
	suite.NoError(err)
	suite.Equal(receivedNFT.GetID(), id)
	suite.True(receivedNFT.GetOwner().Equals(address))
	suite.Equal(receivedNFT.GetTokenURI(), tokenURI)

	// MintNFT shouldn't fail when collection exists
	err = suite.keeper.MintNFT(suite.ctx, denom, id2, tokenURI, address)
	suite.NoError(err)

	// GetNFT should get the NFT when collection exists
	receivedNFT2, err := suite.keeper.GetNFT(suite.ctx, denom, id2)
	suite.NoError(err)
	suite.Equal(receivedNFT2.GetID(), id2)
	suite.True(receivedNFT2.GetOwner().Equals(address))
	suite.Equal(receivedNFT2.GetTokenURI(), tokenURI)

	msg, fail := keeper.SupplyInvariant(suite.keeper)(suite.ctx)
	suite.False(fail, msg)
}

func (suite *KeeperSuite) TestGetNFTs() {
	err := suite.keeper.MintNFT(suite.ctx, denom2, id, tokenURI, address)
	suite.NoError(err)

	err = suite.keeper.MintNFT(suite.ctx, denom2, id2, tokenURI, address)
	suite.NoError(err)

	err = suite.keeper.MintNFT(suite.ctx, denom2, id3, tokenURI, address)
	suite.NoError(err)

	err = suite.keeper.MintNFT(suite.ctx, denom, id3, tokenURI, address)
	suite.NoError(err)

	nfts := suite.keeper.GetNFTs(suite.ctx, denom2)
	suite.Len(nfts, 3)
}

func (suite *KeeperSuite) TestAuthorize() {
	err := suite.keeper.MintNFT(suite.ctx, denom, id, tokenURI, address)
	suite.NoError(err)

	_, err = suite.keeper.Authorize(suite.ctx, denom, id, address2)
	suite.Error(err)

	_, err = suite.keeper.Authorize(suite.ctx, denom, id, address)
	suite.NoError(err)
}

func (suite *KeeperSuite) TestHasNFT() {
	// IsNFT should return false
	isNFT := suite.keeper.HasNFT(suite.ctx, denom, id)
	suite.False(isNFT)

	// MintNFT shouldn't fail when collection does not exist
	err := suite.keeper.MintNFT(suite.ctx, denom, id, tokenURI, address)
	suite.NoError(err)

	// IsNFT should return true
	isNFT = suite.keeper.HasNFT(suite.ctx, denom, id)
	suite.True(isNFT)
}
