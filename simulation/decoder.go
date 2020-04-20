package simulation

import (
	"bytes"
	"fmt"

	kv "github.com/tendermint/tendermint/libs/kv"

	"github.com/cosmos/cosmos-sdk/codec"

	"github.com/irismod/nft/types"
)

// DecodeStore unmarshals the KVPair's Value to the corresponding gov type
func DecodeStore(cdc *codec.Codec, kvA, kvB kv.Pair) string {
	switch {
	case bytes.Equal(kvA.Key[:1], types.PrefixNFT):
		var nftA, nftB types.BaseNFT
		cdc.MustUnmarshalBinaryLengthPrefixed(kvA.Value, &nftA)
		cdc.MustUnmarshalBinaryLengthPrefixed(kvB.Value, &nftB)
		return fmt.Sprintf("%v\n%v", nftA, nftB)

	case bytes.Equal(kvA.Key[:1], types.PrefixOwners):
		var idA, idB string
		cdc.MustUnmarshalBinaryLengthPrefixed(kvA.Value, &idA)
		cdc.MustUnmarshalBinaryLengthPrefixed(kvB.Value, &idB)
		return fmt.Sprintf("%v\n%v", idA, idB)
	case bytes.Equal(kvA.Key[:1], types.PrefixCollection):
		var supplyA, supplyB uint64
		cdc.MustUnmarshalBinaryLengthPrefixed(kvA.Value, &supplyA)
		cdc.MustUnmarshalBinaryLengthPrefixed(kvB.Value, &supplyB)
		return fmt.Sprintf("%d\n%d", supplyA, supplyB)

	default:
		panic(fmt.Sprintf("invalid %s key prefix %X", types.ModuleName, kvA.Key[:1]))
	}
}
