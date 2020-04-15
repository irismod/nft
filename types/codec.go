package types

// DONTCOVER

import (
	"github.com/cosmos/cosmos-sdk/codec"

	"github.com/irismod/nft/exported"
)

// RegisterCodec concrete types on codec
func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterConcrete(MsgTransferNFT{}, "irismod/nft/MsgTransferNFT", nil)
	cdc.RegisterConcrete(MsgEditNFTMetadata{}, "irismod/nft/MsgEditNFTMetadata", nil)
	cdc.RegisterConcrete(MsgMintNFT{}, "irismod/nft/MsgMintNFT", nil)
	cdc.RegisterConcrete(MsgBurnNFT{}, "irismod/nft/MsgBurnNFT", nil)

	cdc.RegisterInterface((*exported.NFT)(nil), nil)
	cdc.RegisterConcrete(&BaseNFT{}, "irismod/nft/BaseNFT", nil)
	cdc.RegisterConcrete(&IDCollection{}, "irismod/nft/IDCollection", nil)
	cdc.RegisterConcrete(&Collection{}, "irismod/nft/Collection", nil)
	cdc.RegisterConcrete(&Owner{}, "irismod/nft/Owner", nil)

}

// ModuleCdc generic sealed codec to be used throughout this module
var ModuleCdc *codec.Codec

func init() {
	ModuleCdc = codec.New()
	codec.RegisterCrypto(ModuleCdc)
	RegisterCodec(ModuleCdc)
	ModuleCdc.Seal()
}
