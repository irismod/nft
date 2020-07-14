package keeper_test

import (
	"testing"

	"github.com/cosmos/cosmos-sdk/baseapp"

	"github.com/irismod/nft/types"

	"github.com/stretchr/testify/suite"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	abci "github.com/tendermint/tendermint/abci/types"

	simapp "github.com/irismod/nft/app"
	"github.com/irismod/nft/keeper"
)

var (
	denom     = "denom"
	schema    = "{a:a,b:b}"
	denom2    = "denom2"
	id        = "id1"
	id2       = "id2"
	id3       = "id3"
	address   = types.CreateTestAddrs(1)[0]
	address2  = types.CreateTestAddrs(2)[1]
	address3  = types.CreateTestAddrs(3)[2]
	tokenURI  = "https://google.com/token-1.json"
	tokenURI2 = "https://google.com/token-2.json"
	tokenData = "{a:a,b:b}"

	isCheckTx = false
)

type KeeperSuite struct {
	suite.Suite

	cdc    *codec.Codec
	ctx    sdk.Context
	keeper keeper.Keeper
	app    *simapp.SimApp

	queryClient types.QueryClient
}

func (suite *KeeperSuite) SetupTest() {

	app := simapp.Setup(isCheckTx)

	suite.app = app
	suite.cdc = app.Codec()
	suite.ctx = app.BaseApp.NewContext(isCheckTx, abci.Header{})
	suite.keeper = app.NFTKeeper

	queryHelper := baseapp.NewQueryServerTestHelper(suite.ctx, app.InterfaceRegistry())
	types.RegisterQueryServer(queryHelper, app.NFTKeeper)
	suite.queryClient = types.NewQueryClient(queryHelper)

	denomE := types.Denom{
		Name:    denom,
		Schema:  schema,
		Creator: address,
	}

	err := suite.keeper.SetDenom(suite.ctx, denomE)
	suite.Nil(err)

	denomE2 := types.Denom{
		Name:    denom2,
		Schema:  schema,
		Creator: address,
	}

	err = suite.keeper.SetDenom(suite.ctx, denomE2)
	suite.Nil(err)

	// collections should equal 1
	collections := suite.keeper.GetCollections(suite.ctx)
	suite.NotEmpty(collections)
	suite.Equal(len(collections), 2)
}

func TestKeeperSuite(t *testing.T) {
	suite.Run(t, new(KeeperSuite))
}

func (suite *KeeperSuite) TestMintNFT() {
	// MintNFT shouldn't fail when collection does not exist
	err := suite.keeper.MintNFT(suite.ctx, denom, id, tokenURI, tokenData, address)
	suite.NoError(err)

	// MintNFT shouldn't fail when collection exists
	err = suite.keeper.MintNFT(suite.ctx, denom, id2, tokenURI, tokenData, address)
	suite.NoError(err)
}

func (suite *KeeperSuite) TestUpdateNFT() {
	// EditNFT should fail when NFT doesn't exists
	err := suite.keeper.EditNFT(suite.ctx, denom, id, tokenURI, tokenData, address)
	suite.Error(err)

	// MintNFT shouldn't fail when collection does not exist
	err = suite.keeper.MintNFT(suite.ctx, denom, id, tokenURI, tokenData, address)
	suite.NoError(err)

	// EditNFT should fail when NFT doesn't exists
	err = suite.keeper.EditNFT(suite.ctx, denom, id2, tokenURI, tokenData, address)
	suite.Error(err)

	// EditNFT shouldn't fail when NFT exists
	err = suite.keeper.EditNFT(suite.ctx, denom, id, tokenURI2, tokenData, address)
	suite.NoError(err)

	// GetNFT should get the NFT with new tokenURI
	receivedNFT, err := suite.keeper.GetNFT(suite.ctx, denom, id)
	suite.NoError(err)
	suite.Equal(receivedNFT.GetTokenURI(), tokenURI2)

	// EditNFT shouldn't fail when NFT exists
	err = suite.keeper.EditNFT(suite.ctx, denom, id, tokenURI2, tokenData, address2)
	suite.Error(err)
}

func (suite *KeeperSuite) TestTransferOwner() {

	// MintNFT shouldn't fail when collection does not exist
	err := suite.keeper.MintNFT(suite.ctx, denom, id, tokenURI, tokenData, address)
	suite.NoError(err)

	//invalid owner
	err = suite.keeper.TransferOwner(suite.ctx, denom, id, tokenURI, tokenData, address2, address3)
	suite.Error(err)

	//right
	err = suite.keeper.TransferOwner(suite.ctx, denom, id, tokenURI2, tokenData, address, address2)
	suite.NoError(err)

	nft, err := suite.keeper.GetNFT(suite.ctx, denom, id)
	suite.NoError(err)
	suite.Equal(tokenURI2, nft.GetTokenURI())
}

func (suite *KeeperSuite) TestBurnNFT() {
	// MintNFT should not fail when collection does not exist
	err := suite.keeper.MintNFT(suite.ctx, denom, id, tokenURI, tokenData, address)
	suite.NoError(err)

	// BurnNFT should fail when NFT doesn't exist but collection does exist
	err = suite.keeper.BurnNFT(suite.ctx, denom, id, address)
	suite.NoError(err)

	// NFT should no longer exist
	isNFT := suite.keeper.HasNFT(suite.ctx, denom, id)
	suite.False(isNFT)

	msg, fail := keeper.SupplyInvariant(suite.keeper)(suite.ctx)
	suite.False(fail, msg)
}
