package nft_test

import (
	"fmt"
	"testing"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/stretchr/testify/suite"
	abci "github.com/tendermint/tendermint/abci/types"

	simapp "github.com/irismod/nft/app"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/irismod/nft"
	"github.com/irismod/nft/types"
)

const (
	module    = "module"
	denom     = "denom"
	schema    = "{}"
	nftID     = "token-id"
	sender    = "sender"
	recipient = "recipient"
	tokenURI  = "token-uri"
)

type HandlerSuite struct {
	suite.Suite

	cdc     *codec.Codec
	ctx     sdk.Context
	app     *simapp.SimApp
	handler sdk.Handler
}

func (suite *HandlerSuite) SetupTest() {
	app := simapp.Setup(false)

	suite.cdc = app.Codec()
	suite.ctx = app.BaseApp.NewContext(false, abci.Header{})
	suite.app = app
	suite.handler = nft.NewHandler(app.NFTKeeper)

	issueDenomMsg := types.NewMsgIssueDenom(address, denom, schema)
	_, err := suite.handler(suite.ctx, &issueDenomMsg)
	suite.NoError(err)

	issueDenom2Msg := types.NewMsgIssueDenom(address, denom2, schema)
	_, err = suite.handler(suite.ctx, &issueDenom2Msg)
	suite.NoError(err)
}

func TestHandlerSuite(t *testing.T) {
	suite.Run(t, new(HandlerSuite))
}

func (suite *HandlerSuite) TestTransferNFTMsg() {
	// Define MsgTransferNft
	transferNftMsg := types.NewMsgTransferNFT(address, address2, denom, id, tokenURI, tokenData)

	// handle should fail trying to transfer NFT that doesn't exist
	res, err := suite.handler(suite.ctx, &transferNftMsg)
	suite.Error(err)
	suite.Nil(res)

	// Create token (collection and owner)
	err = suite.app.NFTKeeper.MintNFT(suite.ctx, denom, id, tokenURI, tokenData, address)
	suite.Nil(err)
	suite.True(CheckInvariants(suite.app.NFTKeeper, suite.ctx))

	// handle should succeed when nft exists and is transferred by owner
	res, err = suite.handler(suite.ctx, &transferNftMsg)
	suite.NoError(err)
	suite.NotNil(res)
	suite.True(CheckInvariants(suite.app.NFTKeeper, suite.ctx))

	// event events should be emitted correctly
	for _, event := range res.Events {
		for _, attribute := range event.Attributes {
			value := string(attribute.Value)
			switch key := string(attribute.Key); key {
			case module:
				suite.Equal(value, types.ModuleName)
			case denom:
				suite.Equal(value, denom)
			case nftID:
				suite.Equal(value, id)
			case sender:
				suite.Equal(value, address.String())
			case recipient:
				suite.Equal(value, address2.String())
			default:
				suite.Fail(fmt.Sprintf("unrecognized event %s", key))
			}
		}
	}

	// nft should have been transferred as a result of the message
	nftAfterwards, err := suite.app.NFTKeeper.GetNFT(suite.ctx, denom, id)
	suite.NoError(err)
	suite.True(nftAfterwards.GetOwner().Equals(address2))

	transferNftMsg = types.NewMsgTransferNFT(address2, address3, denom, id, tokenURI, tokenData)

	// handle should succeed when nft exists and is transferred by owner
	res, err = suite.handler(suite.ctx, &transferNftMsg)
	suite.NoError(err)
	suite.NotNil(res)
	suite.True(CheckInvariants(suite.app.NFTKeeper, suite.ctx))

	// Create token (collection and owner)
	err = suite.app.NFTKeeper.MintNFT(suite.ctx, denom2, id, tokenURI, tokenData, address)
	suite.Nil(err)
	suite.True(CheckInvariants(suite.app.NFTKeeper, suite.ctx))

	transferNftMsg = types.NewMsgTransferNFT(address2, address3, denom2, id, tokenURI, tokenData)

	// handle should fail when nft exists and is not transferred by owner
	res, err = suite.handler(suite.ctx, &transferNftMsg)
	suite.Error(err)
	suite.Nil(res)
	suite.True(CheckInvariants(suite.app.NFTKeeper, suite.ctx))
}

func (suite *HandlerSuite) TestEditNFTMsg() {
	// Create token (collection and address)
	err := suite.app.NFTKeeper.MintNFT(suite.ctx, denom, id, tokenURI, tokenData, address)
	suite.Nil(err)

	// Define MsgTransferNft
	failingEditNFT := types.NewMsgEditNFT(address, id, denom2, tokenURI2, tokenData)

	res, err := suite.handler(suite.ctx, &failingEditNFT)
	suite.Error(err)
	suite.Nil(res)

	// Define MsgTransferNft
	editNFT := types.NewMsgEditNFT(address, id, denom, tokenURI2, tokenData)

	res, err = suite.handler(suite.ctx, &editNFT)
	suite.NoError(err)
	suite.NotNil(res)

	// event events should be emitted correctly
	for _, event := range res.Events {
		for _, attribute := range event.Attributes {
			value := string(attribute.Value)
			switch key := string(attribute.Key); key {
			case module:
				suite.Equal(value, types.ModuleName)
			case denom:
				suite.Equal(value, denom)
			case nftID:
				suite.Equal(value, id)
			case sender:
				suite.Equal(value, address.String())
			case tokenURI:
				suite.Equal(value, tokenURI2)
			default:
				suite.Fail(fmt.Sprintf("unrecognized event %s", key))
			}
		}
	}

	nftAfterwards, err := suite.app.NFTKeeper.GetNFT(suite.ctx, denom, id)
	suite.NoError(err)
	suite.Equal(tokenURI2, nftAfterwards.GetTokenURI())
}

func (suite *HandlerSuite) TestMintNFTMsg() {
	// Define MsgMintNFT
	mintNFT := types.NewMsgMintNFT(address, address, id, denom, tokenURI, tokenData)

	// minting a token should succeed
	res, err := suite.handler(suite.ctx, &mintNFT)
	suite.NoError(err)
	suite.NotNil(res)

	// event events should be emitted correctly
	for _, event := range res.Events {
		for _, attribute := range event.Attributes {
			value := string(attribute.Value)
			switch key := string(attribute.Key); key {
			case module:
				suite.Equal(value, types.ModuleName)
			case denom:
				suite.Equal(value, denom)
			case nftID:
				suite.Equal(value, id)
			case sender:
				suite.Equal(value, address.String())
			case recipient:
				suite.Equal(value, address.String())
			case tokenURI:
				suite.Equal(value, tokenURI)
			default:
				suite.Fail(fmt.Sprintf("unrecognized event %s", key))
			}
		}
	}

	nftAfterwards, err := suite.app.NFTKeeper.GetNFT(suite.ctx, denom, id)

	suite.NoError(err)
	suite.Equal(tokenURI, nftAfterwards.GetTokenURI())

	// minting the same token should fail
	res, err = suite.handler(suite.ctx, &mintNFT)
	suite.Error(err)
	suite.Nil(res)

	suite.True(CheckInvariants(suite.app.NFTKeeper, suite.ctx))
}

func (suite *HandlerSuite) TestBurnNFTMsg() {
	// Create token (collection and address)
	err := suite.app.NFTKeeper.MintNFT(suite.ctx, denom, id, tokenURI, tokenData, address)
	suite.Nil(err)

	exists := suite.app.NFTKeeper.HasNFT(suite.ctx, denom, id)
	suite.True(exists)

	// burning a non-existent NFT should fail
	failBurnNFT := types.NewMsgBurnNFT(address, id2, denom)
	res, err := suite.handler(suite.ctx, &failBurnNFT)
	suite.Error(err)
	suite.Nil(res)

	// NFT should still exist
	exists = suite.app.NFTKeeper.HasNFT(suite.ctx, denom, id)
	suite.True(exists)

	// burning the NFt should succeed
	burnNFT := types.NewMsgBurnNFT(address, id, denom)

	res, err = suite.handler(suite.ctx, &burnNFT)
	suite.NoError(err)
	suite.NotNil(res)

	// event events should be emitted correctly
	for _, event := range res.Events {
		for _, attribute := range event.Attributes {
			value := string(attribute.Value)
			switch key := string(attribute.Key); key {
			case module:
				suite.Equal(value, types.ModuleName)
			case denom:
				suite.Equal(value, denom)
			case nftID:
				suite.Equal(value, id)
			case sender:
				suite.Equal(value, address.String())
			default:
				suite.Fail(fmt.Sprintf("unrecognized event %s", key))
			}
		}
	}

	// the NFT should not exist after burn
	exists = suite.app.NFTKeeper.HasNFT(suite.ctx, denom, id)
	suite.False(exists)

	ownerReturned := suite.app.NFTKeeper.GetOwner(suite.ctx, address, "")
	suite.Equal(0, len(ownerReturned.IDCollections))

	suite.True(CheckInvariants(suite.app.NFTKeeper, suite.ctx))
}
