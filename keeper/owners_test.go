package keeper_test

import (
	"testing"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simapp "github.com/irismod/nft/app"
	"github.com/stretchr/testify/suite"
	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/irismod/nft/keeper"
	"github.com/irismod/nft/types"
)

type KeeperSuite struct {
	suite.Suite

	cdc    *codec.Codec
	ctx    sdk.Context
	keeper keeper.Keeper
}

func (suite *KeeperSuite) SetupTest() {

	app := simapp.Setup(isCheckTx)

	suite.cdc = app.Codec()
	suite.ctx = app.BaseApp.NewContext(isCheckTx, abci.Header{})
	suite.keeper = app.NFTKeeper
}

func TestKeeperSuite(t *testing.T) {
	suite.Run(t, new(KeeperSuite))
}

func (suite *KeeperSuite) TestGetOwners() {

	nft := types.NewBaseNFT(id, address, tokenURI)
	err := suite.keeper.MintNFT(suite.ctx, denom, &nft)
	suite.NoError(err)

	nft2 := types.NewBaseNFT(id2, address2, tokenURI)
	err = suite.keeper.MintNFT(suite.ctx, denom, &nft2)
	suite.NoError(err)

	nft3 := types.NewBaseNFT(id3, address3, tokenURI)
	err = suite.keeper.MintNFT(suite.ctx, denom, &nft3)
	suite.NoError(err)

	owners := suite.keeper.GetOwners(suite.ctx)
	suite.Equal(3, len(owners))

	nft = types.NewBaseNFT(id, address, tokenURI)
	err = suite.keeper.MintNFT(suite.ctx, denom2, &nft)
	suite.NoError(err)

	nft2 = types.NewBaseNFT(id2, address2, tokenURI)
	err = suite.keeper.MintNFT(suite.ctx, denom2, &nft2)
	suite.NoError(err)

	nft3 = types.NewBaseNFT(id3, address3, tokenURI)
	err = suite.keeper.MintNFT(suite.ctx, denom2, &nft3)
	suite.NoError(err)

	owners = suite.keeper.GetOwners(suite.ctx)
	suite.Equal(3, len(owners))

	msg, fail := keeper.SupplyInvariant(suite.keeper)(suite.ctx)
	suite.False(fail, msg)
}

func (suite *KeeperSuite) TestGetOwner() {
	nft := types.NewBaseNFT(id, address, tokenURI)
	err := suite.keeper.MintNFT(suite.ctx, denom, &nft)
	suite.NoError(err)

	owner := suite.keeper.GetOwner(suite.ctx, address)
	suite.Len(owner.IDCollections, 1)
	suite.Len(owner.IDCollections[0].IDs, 1)
	suite.Equal(owner.IDCollections[0].IDs[0], nft.ID)
	suite.Equal(owner.Address, address)
	suite.Equal(owner.Address, nft.Owner)
}
