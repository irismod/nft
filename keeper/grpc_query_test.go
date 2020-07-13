package keeper_test

import (
	gocontext "context"
	"github.com/irismod/nft/types"
)

func (suite *KeeperSuite) TestSupply() {
	err := suite.keeper.MintNFT(suite.ctx, denom, id, tokenURI, tokenData, address)
	suite.NoError(err)

	response, err := suite.queryClient.Supply(gocontext.Background(), &types.QuerySupplyRequest{
		Denom: denom,
		Owner: address,
	})

	suite.NoError(err)
	suite.Equal(1, int(response.Amount))
}

func (suite *KeeperSuite) TestOwner() {
	err := suite.keeper.MintNFT(suite.ctx, denom, id, tokenURI, tokenData, address)
	suite.NoError(err)

	response, err := suite.queryClient.Owner(gocontext.Background(), &types.QueryOwnerRequest{
		Denom: denom,
		Owner: nil,
	})

	suite.NoError(err)
	suite.NotNil(response.Owner)
	suite.Contains(response.Owner.IDCollections[0].IDs, id)
}

func (suite *KeeperSuite) TestCollection() {
	err := suite.keeper.MintNFT(suite.ctx, denom, id, tokenURI, tokenData, address)
	suite.NoError(err)

	response, err := suite.queryClient.Collection(gocontext.Background(), &types.QueryCollectionRequest{
		Denom: denom,
	})

	suite.NoError(err)
	suite.NotNil(response.Collection)
	suite.Len(response.Collection.NFTs, 1)
	suite.Equal(response.Collection.NFTs[0].ID, id)
}

func (suite *KeeperSuite) TestDenom() {
	err := suite.keeper.MintNFT(suite.ctx, denom, id, tokenURI, tokenData, address)
	suite.NoError(err)

	response, err := suite.queryClient.Denom(gocontext.Background(), &types.QueryDenomRequest{
		Denom: denom,
	})

	suite.NoError(err)
	suite.NotNil(response.Denom)
	suite.Equal(response.Denom.Name, denom)
}

func (suite *KeeperSuite) TestDenoms() {
	err := suite.keeper.MintNFT(suite.ctx, denom, id, tokenURI, tokenData, address)
	suite.NoError(err)

	response, err := suite.queryClient.Denoms(gocontext.Background(), &types.QueryDenomsRequest{})

	suite.NoError(err)
	suite.NotEmpty(response.Denoms)
	suite.Equal(response.Denoms[0].Name, denom)
}

func (suite *KeeperSuite) TestNFT() {
	err := suite.keeper.MintNFT(suite.ctx, denom, id, tokenURI, tokenData, address)
	suite.NoError(err)

	response, err := suite.queryClient.NFT(gocontext.Background(), &types.QueryNFTRequest{
		Denom:   denom,
		TokenID: id,
	})

	suite.NoError(err)
	suite.NotEmpty(response.NFT)
	suite.Equal(response.NFT.ID, id)
}
