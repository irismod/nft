package types

// DONTCOVER

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	gogotypes "github.com/gogo/protobuf/types"

	"github.com/irismod/nft/exported"
)

// RegisterCodec concrete types on codec
func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterConcrete(MsgIssueDenom{}, "irismod/nft/MsgIssueDenom", nil)
	cdc.RegisterConcrete(MsgTransferNFT{}, "irismod/nft/MsgTransferNFT", nil)
	cdc.RegisterConcrete(MsgEditNFT{}, "irismod/nft/MsgEditNFT", nil)
	cdc.RegisterConcrete(MsgMintNFT{}, "irismod/nft/MsgMintNFT", nil)
	cdc.RegisterConcrete(MsgBurnNFT{}, "irismod/nft/MsgBurnNFT", nil)

	cdc.RegisterInterface((*exported.NFT)(nil), nil)
	cdc.RegisterConcrete(&BaseNFT{}, "irismod/nft/BaseNFT", nil)
}

var (
	amino = codec.New()

	// ModuleCdc references the global x/gov module codec. Note, the codec should
	// ONLY be used in certain instances of tests and for JSON encoding as Amino is
	// still used for that purpose.
	//
	// The actual codec used for serialization should be provided to x/gov and
	// defined at the application level.
	ModuleCdc = codec.NewHybridCodec(amino, types.NewInterfaceRegistry())
)

func init() {
	codec.RegisterCrypto(amino)
	RegisterCodec(amino)
	amino.Seal()
}

// return supply protobuf code
func MustMarshalSupply(cdc codec.Marshaler, supply uint64) []byte {
	supplyWrap := gogotypes.UInt64Value{Value: supply}
	return cdc.MustMarshalBinaryBare(&supplyWrap)
}

// return th supply
func MustUnMarshalSupply(cdc codec.Marshaler, value []byte) uint64 {
	var supplyWrap gogotypes.UInt64Value
	cdc.MustUnmarshalBinaryBare(value, &supplyWrap)
	return supplyWrap.Value
}

// return the tokenID protobuf code
func MustMarshalTokenID(cdc codec.Marshaler, tokenID string) []byte {
	tokenIDWrap := gogotypes.StringValue{Value: tokenID}
	return cdc.MustMarshalBinaryBare(&tokenIDWrap)
}

// return th tokenID
func MustUnMarshalTokenID(cdc codec.Marshaler, value []byte) string {
	var tokenIDWrap gogotypes.StringValue
	cdc.MustUnmarshalBinaryBare(value, &tokenIDWrap)
	return tokenIDWrap.Value
}
