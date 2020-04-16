package keeper_test

import (
	"github.com/irismod/nft/keeper"
	"github.com/irismod/nft/types"
)

func (suite *KeeperSuite) TestMintNFT() {
	// MintNFT shouldn't fail when collection does not exist
	nft := types.NewBaseNFT(id, address, tokenURI)
	err := suite.keeper.MintNFT(suite.ctx, denom, &nft)
	suite.NoError(err)

	// MintNFT shouldn't fail when collection exists
	nft2 := types.NewBaseNFT(id2, address, tokenURI)
	err = suite.keeper.MintNFT(suite.ctx, denom, &nft2)
	suite.NoError(err)
}

func (suite *KeeperSuite) TestGetNFT() {
	// MintNFT shouldn't fail when collection does not exist
	nft := types.NewBaseNFT(id, address, tokenURI)
	err := suite.keeper.MintNFT(suite.ctx, denom, &nft)
	suite.NoError(err)

	// GetNFT should get the NFT
	receivedNFT, err := suite.keeper.GetNFT(suite.ctx, denom, id)
	suite.NoError(err)
	suite.Equal(receivedNFT.GetID(), id)
	suite.True(receivedNFT.GetOwner().Equals(address))
	suite.Equal(receivedNFT.GetTokenURI(), tokenURI)

	// MintNFT shouldn't fail when collection exists
	nft2 := types.NewBaseNFT(id2, address, tokenURI)
	err = suite.keeper.MintNFT(suite.ctx, denom, &nft2)
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

func (suite *KeeperSuite) TestUpdateNFT() {
	nft := types.NewBaseNFT(id, address, tokenURI)

	// UpdateNFT should fail when NFT doesn't exists
	err := suite.keeper.UpdateNFT(suite.ctx, denom, &nft)
	suite.Error(err)

	// MintNFT shouldn't fail when collection does not exist
	err = suite.keeper.MintNFT(suite.ctx, denom, &nft)
	suite.NoError(err)

	nonnft := types.NewBaseNFT(id2, address, tokenURI)
	// UpdateNFT should fail when NFT doesn't exists
	err = suite.keeper.UpdateNFT(suite.ctx, denom, &nonnft)
	suite.Error(err)

	// UpdateNFT shouldn't fail when NFT exists
	nft2 := types.NewBaseNFT(id, address, tokenURI2)
	err = suite.keeper.UpdateNFT(suite.ctx, denom, &nft2)
	suite.NoError(err)

	// UpdateNFT shouldn't fail when NFT exists
	nft2 = types.NewBaseNFT(id, address2, tokenURI2)
	err = suite.keeper.UpdateNFT(suite.ctx, denom, &nft2)
	suite.NoError(err)

	// GetNFT should get the NFT with new tokenURI
	receivedNFT, err := suite.keeper.GetNFT(suite.ctx, denom, id)
	suite.NoError(err)
	suite.Equal(receivedNFT.GetTokenURI(), tokenURI2)
}

func (suite *KeeperSuite) TestDeleteNFT() {
	nft := types.NewBaseNFT(id, address, tokenURI)

	// MintNFT should not fail when collection does not exist
	err := suite.keeper.MintNFT(suite.ctx, denom, &nft)
	suite.NoError(err)

	// DeleteNFT should fail when NFT doesn't exist but collection does exist
	suite.keeper.DeleteNFT(suite.ctx, denom, &nft)

	// NFT should no longer exist
	isNFT := suite.keeper.HasNFT(suite.ctx, denom, id)
	suite.False(isNFT)

	msg, fail := keeper.SupplyInvariant(suite.keeper)(suite.ctx)
	suite.False(fail, msg)
}

func (suite *KeeperSuite) TestHasNFT() {
	// IsNFT should return false
	isNFT := suite.keeper.HasNFT(suite.ctx, denom, id)
	suite.False(isNFT)

	// MintNFT shouldn't fail when collection does not exist
	nft := types.NewBaseNFT(id, address, tokenURI)
	err := suite.keeper.MintNFT(suite.ctx, denom, &nft)
	suite.NoError(err)

	// IsNFT should return true
	isNFT = suite.keeper.HasNFT(suite.ctx, denom, id)
	suite.True(isNFT)
}
