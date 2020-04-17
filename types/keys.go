package types

import (
	"bytes"
	"errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	// ModuleName is the name of the module
	ModuleName = "nft"

	// StoreKey is the default store key for NFT
	StoreKey = ModuleName

	// QuerierRoute is the querier route for the NFT store.
	QuerierRoute = ModuleName

	// RouterKey is the message route for the NFT module
	RouterKey = ModuleName
)

var (
	PrefixNFT    = []byte{0x01}
	PrefixOwners = []byte{0x02} // key for balance of NFTs held by an address

	delimiter = []byte("/")
)

func SplitKeyOwner(key []byte) (address sdk.AccAddress, denom, id string, err error) {
	key = key[len(PrefixOwners)+len(delimiter):]
	keys := bytes.Split(key, delimiter)

	switch len(keys) {
	case 3:
		address = sdk.AccAddress(keys[0])
		denom = string(keys[1])
		id = string(keys[2])
		return
	case 2:
		address = sdk.AccAddress(keys[0])
		denom = string(keys[1])
		return
	case 1:
		address = sdk.AccAddress(keys[0])
		return
	default:
		return address, denom, id, errors.New("wrong KeyOwner")
	}
}

// KeyOwner gets the key of a collection owned by an account address
func KeyOwner(address sdk.AccAddress, denom, id string) []byte {
	key := append(PrefixOwners, delimiter...)
	if address != nil {
		key = append(key, []byte(address.String())...)
	}

	if address != nil && len(denom) > 0 {
		key = append(key, delimiter...)
		key = append(key, []byte(denom)...)
	}

	if address != nil && len(denom) > 0 && len(id) > 0 {
		key = append(key, delimiter...)
		key = append(key, []byte(id)...)
	}
	return key
}

// KeyNFT gets the NFT by an ID
func KeyNFT(denom, id string) []byte {
	key := append(PrefixNFT, delimiter...)
	if len(denom) > 0 {
		key = append(key, []byte(denom)...)
	}

	if len(denom) > 0 && len(id) > 0 {
		key = append(key, delimiter...)
		key = append(key, []byte(id)...)
	}
	return key
}
func SplitKeyNFT(key []byte) (denom, id string, err error) {
	key = key[len(PrefixNFT)+len(delimiter):]
	keys := bytes.Split(key, delimiter)

	switch len(keys) {
	case 2:
		denom = string(keys[0])
		id = string(keys[1])
		return
	case 1:
		denom = string(keys[0])
		return
	default:
		return denom, id, errors.New("wrong KeyNFT")
	}
}
