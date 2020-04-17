package keeper_test

import (
	"encoding/binary"

	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/irismod/nft/exported"
	keep "github.com/irismod/nft/keeper"
	"github.com/irismod/nft/types"
)

func (suite *KeeperSuite) TestNewQuerier() {
	querier := keep.NewQuerier(suite.keeper)
	query := abci.RequestQuery{
		Path: "",
		Data: []byte{},
	}
	_, err := querier(suite.ctx, []string{"foo", "bar"}, query)
	suite.Error(err)
}

func (suite *KeeperSuite) TestQuerySupply() {
	// MintNFT shouldn't fail when collection does not exist
	nft := types.NewBaseNFT(id, address, tokenURI)
	err := suite.keeper.MintNFT(suite.ctx, denom, &nft)
	suite.NoError(err)

	querier := keep.NewQuerier(suite.keeper)

	query := abci.RequestQuery{
		Path: "",
		Data: []byte{},
	}

	query.Path = "/custom/nft/supply"
	query.Data = []byte("?")

	res, err := querier(suite.ctx, []string{"supply"}, query)
	suite.Error(err)
	suite.Nil(res)

	queryCollectionParams := types.NewQueryCollectionParams(denom2)
	bz, errRes := suite.cdc.MarshalJSON(queryCollectionParams)
	suite.Nil(errRes)
	query.Data = bz
	res, err = querier(suite.ctx, []string{"supply"}, query)
	suite.Nil(errRes)
	supplyResp := binary.LittleEndian.Uint64(res)
	suite.Equal(0, int(supplyResp))

	queryCollectionParams = types.NewQueryCollectionParams(denom)
	bz, errRes = suite.cdc.MarshalJSON(queryCollectionParams)
	suite.Nil(errRes)
	query.Data = bz

	res, err = querier(suite.ctx, []string{"supply"}, query)
	suite.NoError(err)
	suite.NotNil(res)

	supplyResp = binary.LittleEndian.Uint64(res)
	suite.Equal(1, int(supplyResp))
}

func (suite *KeeperSuite) TestQueryCollection() {
	// MintNFT shouldn't fail when collection does not exist
	nft := types.NewBaseNFT(id, address, tokenURI)
	err := suite.keeper.MintNFT(suite.ctx, denom, &nft)
	suite.NoError(err)

	querier := keep.NewQuerier(suite.keeper)

	query := abci.RequestQuery{
		Path: "",
		Data: []byte{},
	}

	query.Path = "/custom/nft/collection"

	query.Data = []byte("?")
	res, err := querier(suite.ctx, []string{"collection"}, query)
	suite.Error(err)
	suite.Nil(res)

	queryCollectionParams := types.NewQueryCollectionParams(denom2)
	bz, errRes := suite.cdc.MarshalJSON(queryCollectionParams)
	suite.Nil(errRes)

	query.Data = bz
	res, err = querier(suite.ctx, []string{"collection"}, query)
	suite.Error(err)
	suite.Nil(res)

	queryCollectionParams = types.NewQueryCollectionParams(denom)
	bz, errRes = suite.cdc.MarshalJSON(queryCollectionParams)
	suite.Nil(errRes)

	query.Data = bz
	res, err = querier(suite.ctx, []string{"collection"}, query)
	suite.NoError(err)
	suite.NotNil(res)

	var collection types.Collection
	types.ModuleCdc.MustUnmarshalJSON(res, &collection)
	suite.Len(collection.NFTs, 1)
}

func (suite *KeeperSuite) TestQueryOwner() {
	// MintNFT shouldn't fail when collection does not exist
	nft := types.NewBaseNFT(id, address, tokenURI)
	err := suite.keeper.MintNFT(suite.ctx, denom, &nft)
	suite.NoError(err)

	denom2 := "test_denom2"
	err = suite.keeper.MintNFT(suite.ctx, denom2, &nft)
	suite.NoError(err)

	querier := keep.NewQuerier(suite.keeper)

	query := abci.RequestQuery{
		Path: "",
		Data: []byte{},
	}
	query.Path = "/custom/nft/ownerByDenom"

	query.Data = []byte("?")
	res, err := querier(suite.ctx, []string{"ownerByDenom"}, query)
	suite.Error(err)
	suite.Nil(res)

	// query the balance using the first denom
	params := types.NewQueryBalanceParams(address, denom)
	bz, err2 := suite.cdc.MarshalJSON(params)
	suite.Nil(err2)
	query.Data = bz

	res, err = querier(suite.ctx, []string{"ownerByDenom"}, query)
	suite.NoError(err)
	suite.NotNil(res)

	var out types.Owner
	suite.cdc.MustUnmarshalJSON(res, &out)

	// build the owner using only the first denom
	idCollection1 := types.NewIDCollection(denom, []string{id})
	owner := types.NewOwner(address, idCollection1)

	suite.Equal(out.String(), owner.String())

	// query the balance using no denom so that all denoms will be returns
	params = types.NewQueryBalanceParams(address, "")
	bz, err2 = suite.cdc.MarshalJSON(params)
	suite.Nil(err2)

	query.Path = "/custom/nft/owner"
	query.Data = []byte("?")
	_, err = querier(suite.ctx, []string{"owner"}, query)
	suite.Error(err)

	query.Data = bz
	res, err = querier(suite.ctx, []string{"owner"}, query)
	suite.NoError(err)
	suite.NotNil(res)

	suite.cdc.MustUnmarshalJSON(res, &out)

	// build the owner using both denoms
	idCollection2 := types.NewIDCollection(denom2, []string{id})
	owner = types.NewOwner(address, idCollection1, idCollection2)

	suite.Equal(out.String(), owner.String())
}

func (suite *KeeperSuite) TestQueryNFT() {
	// MintNFT shouldn't fail when collection does not exist
	nft := types.NewBaseNFT(id, address, tokenURI)
	err := suite.keeper.MintNFT(suite.ctx, denom, &nft)
	suite.NoError(err)

	querier := keep.NewQuerier(suite.keeper)

	query := abci.RequestQuery{
		Path: "",
		Data: []byte{},
	}
	query.Path = "/custom/nft/nft"
	var res []byte

	query.Data = []byte("?")
	res, err = querier(suite.ctx, []string{"nft"}, query)
	suite.Error(err)
	suite.Nil(res)

	params := types.NewQueryNFTParams(denom2, id2)
	bz, err2 := suite.cdc.MarshalJSON(params)
	suite.Nil(err2)

	query.Data = bz
	res, err = querier(suite.ctx, []string{"nft"}, query)
	suite.Error(err)
	suite.Nil(res)

	params = types.NewQueryNFTParams(denom, id)
	bz, err2 = suite.cdc.MarshalJSON(params)
	suite.Nil(err2)

	query.Data = bz
	res, err = querier(suite.ctx, []string{"nft"}, query)
	suite.NoError(err)
	suite.NotNil(res)

	var out exported.NFT
	suite.cdc.MustUnmarshalJSON(res, &out)

	suite.Equal(out.String(), nft.String())
}

func (suite *KeeperSuite) TestQueryDenoms() {
	// MintNFT shouldn't fail when collection does not exist
	nft := types.NewBaseNFT(id, address, tokenURI)
	err := suite.keeper.MintNFT(suite.ctx, denom, &nft)
	suite.NoError(err)

	err = suite.keeper.MintNFT(suite.ctx, denom2, &nft)
	suite.NoError(err)

	querier := keep.NewQuerier(suite.keeper)

	query := abci.RequestQuery{
		Path: "",
		Data: []byte{},
	}
	var res []byte
	query.Path = "/custom/nft/denoms"

	res, err = querier(suite.ctx, []string{"denoms"}, query)
	suite.NoError(err)
	suite.NotNil(res)

	denoms := []string{denom, denom2}

	var out []string
	suite.cdc.MustUnmarshalJSON(res, &out)

	for key, denomInQuestion := range out {
		suite.Equal(denomInQuestion, denoms[key])
	}
}
