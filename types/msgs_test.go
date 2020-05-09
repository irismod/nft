package types

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

// ---------------------------------------- Msgs ---------------------------------------------------

func TestNewMsgTransferNFT(t *testing.T) {
	newMsgTransferNFT := NewMsgTransferNFT(address, address2,
		fmt.Sprintf("     %s     ", denom),
		fmt.Sprintf("     %s     ", id), tokenURI, metadata)
	require.Equal(t, newMsgTransferNFT.Sender, address)
	require.Equal(t, newMsgTransferNFT.Recipient, address2)
	require.Equal(t, newMsgTransferNFT.Denom, denom)
	require.Equal(t, newMsgTransferNFT.ID, id)
}

func TestMsgTransferNFTValidateBasicMethod(t *testing.T) {
	newMsgTransferNFT := NewMsgTransferNFT(address, address2, "", id, tokenURI, metadata)
	err := newMsgTransferNFT.ValidateBasic()
	require.Error(t, err)

	newMsgTransferNFT = NewMsgTransferNFT(address, address2, denom, "", tokenURI, metadata)
	err = newMsgTransferNFT.ValidateBasic()
	require.Error(t, err)

	newMsgTransferNFT = NewMsgTransferNFT(nil, address2, denom, "", tokenURI, metadata)
	err = newMsgTransferNFT.ValidateBasic()
	require.Error(t, err)

	newMsgTransferNFT = NewMsgTransferNFT(address, nil, denom, "", tokenURI, metadata)
	err = newMsgTransferNFT.ValidateBasic()
	require.Error(t, err)

	newMsgTransferNFT = NewMsgTransferNFT(address, address2, denom, id, tokenURI, metadata)
	err = newMsgTransferNFT.ValidateBasic()
	require.NoError(t, err)
}

func TestMsgTransferNFTGetSignBytesMethod(t *testing.T) {
	newMsgTransferNFT := NewMsgTransferNFT(address, address2, denom, id, tokenURI, metadata)
	sortedBytes := newMsgTransferNFT.GetSignBytes()
	require.Equal(t, string(sortedBytes), `{"type":"irismod/nft/MsgTransferNFT","value":{"denom":"denom","id":"id1","metadata":"https://google.com/token-1.json","recipient":"cosmos15ky9du8a2wlstz6fpx3p4mqpjyrm5cgp0ctjdj","sender":"cosmos15ky9du8a2wlstz6fpx3p4mqpjyrm5cgqjwl8sq","token_uri":"https://google.com/token-1.json"}}`)
}

func TestMsgTransferNFTGetSignersMethod(t *testing.T) {
	newMsgTransferNFT := NewMsgTransferNFT(address, address2, denom, id, tokenURI, metadata)
	signers := newMsgTransferNFT.GetSigners()
	require.Equal(t, 1, len(signers))
	require.Equal(t, address.String(), signers[0].String())
}

func TestNewMsgEditNFTMetadata(t *testing.T) {
	newMsgEditNFTMetadata := NewMsgEditNFT(address,
		fmt.Sprintf("     %s     ", id),
		fmt.Sprintf("     %s     ", denom),
		fmt.Sprintf("     %s     ", tokenURI), metadata)

	require.Equal(t, newMsgEditNFTMetadata.Sender.String(), address.String())
	require.Equal(t, newMsgEditNFTMetadata.ID, id)
	require.Equal(t, newMsgEditNFTMetadata.Denom, denom)
	require.Equal(t, newMsgEditNFTMetadata.TokenURI, tokenURI)
}

func TestMsgEditNFTMetadataValidateBasicMethod(t *testing.T) {
	newMsgEditNFTMetadata := NewMsgEditNFT(nil, id, denom, tokenURI, metadata)

	err := newMsgEditNFTMetadata.ValidateBasic()
	require.Error(t, err)

	newMsgEditNFTMetadata = NewMsgEditNFT(address, "", denom, tokenURI, metadata)
	err = newMsgEditNFTMetadata.ValidateBasic()
	require.Error(t, err)

	newMsgEditNFTMetadata = NewMsgEditNFT(address, id, "", tokenURI, metadata)
	err = newMsgEditNFTMetadata.ValidateBasic()
	require.Error(t, err)

	newMsgEditNFTMetadata = NewMsgEditNFT(address, id, denom, tokenURI, metadata)
	err = newMsgEditNFTMetadata.ValidateBasic()
	require.NoError(t, err)
}

func TestMsgEditNFTMetadataGetSignBytesMethod(t *testing.T) {
	newMsgEditNFTMetadata := NewMsgEditNFT(address, id, denom, tokenURI, metadata)
	sortedBytes := newMsgEditNFTMetadata.GetSignBytes()
	require.Equal(t, string(sortedBytes), `{"type":"irismod/nft/MsgEditNFT","value":{"denom":"denom","id":"id1","metadata":"https://google.com/token-1.json","sender":"cosmos15ky9du8a2wlstz6fpx3p4mqpjyrm5cgqjwl8sq","token_uri":"https://google.com/token-1.json"}}`)
}

func TestMsgEditNFTMetadataGetSignersMethod(t *testing.T) {
	newMsgEditNFTMetadata := NewMsgEditNFT(address, id, denom, tokenURI, metadata)
	signers := newMsgEditNFTMetadata.GetSigners()
	require.Equal(t, 1, len(signers))
	require.Equal(t, address.String(), signers[0].String())
}

func TestNewMsgMintNFT(t *testing.T) {
	newMsgMintNFT := NewMsgMintNFT(address, address2,
		fmt.Sprintf("     %s     ", id),
		fmt.Sprintf("     %s     ", denom),
		fmt.Sprintf("     %s     ", tokenURI), metadata)

	require.Equal(t, newMsgMintNFT.Sender.String(), address.String())
	require.Equal(t, newMsgMintNFT.Recipient.String(), address2.String())
	require.Equal(t, newMsgMintNFT.ID, id)
	require.Equal(t, newMsgMintNFT.Denom, denom)
	require.Equal(t, newMsgMintNFT.TokenURI, tokenURI)
}

func TestMsgMsgMintNFTValidateBasicMethod(t *testing.T) {
	newMsgMintNFT := NewMsgMintNFT(nil, address2, id, denom, tokenURI, metadata)
	err := newMsgMintNFT.ValidateBasic()
	require.Error(t, err)

	newMsgMintNFT = NewMsgMintNFT(address, address2, "", denom, tokenURI, metadata)
	err = newMsgMintNFT.ValidateBasic()
	require.Error(t, err)

	newMsgMintNFT = NewMsgMintNFT(address, address2, id, "", tokenURI, metadata)
	err = newMsgMintNFT.ValidateBasic()
	require.Error(t, err)

	newMsgMintNFT = NewMsgMintNFT(address, address2, id, denom, tokenURI, metadata)
	err = newMsgMintNFT.ValidateBasic()
	require.NoError(t, err)
}

func TestMsgMintNFTGetSignBytesMethod(t *testing.T) {
	newMsgMintNFT := NewMsgMintNFT(address, address2, id, denom, tokenURI, metadata)
	sortedBytes := newMsgMintNFT.GetSignBytes()
	require.Equal(t, string(sortedBytes), `{"type":"irismod/nft/MsgMintNFT","value":{"denom":"denom","id":"id1","metadata":"https://google.com/token-1.json","recipient":"cosmos15ky9du8a2wlstz6fpx3p4mqpjyrm5cgp0ctjdj","sender":"cosmos15ky9du8a2wlstz6fpx3p4mqpjyrm5cgqjwl8sq","token_uri":"https://google.com/token-1.json"}}`)
}

func TestMsgMintNFTGetSignersMethod(t *testing.T) {
	newMsgMintNFT := NewMsgMintNFT(address, address2, id, denom, tokenURI, metadata)
	signers := newMsgMintNFT.GetSigners()
	require.Equal(t, 1, len(signers))
	require.Equal(t, address.String(), signers[0].String())
}

func TestNewMsgBurnNFT(t *testing.T) {
	newMsgBurnNFT := NewMsgBurnNFT(address,
		fmt.Sprintf("     %s     ", id),
		fmt.Sprintf("     %s     ", denom))

	require.Equal(t, newMsgBurnNFT.Sender.String(), address.String())
	require.Equal(t, newMsgBurnNFT.ID, id)
	require.Equal(t, newMsgBurnNFT.Denom, denom)
}

func TestMsgMsgBurnNFTValidateBasicMethod(t *testing.T) {
	newMsgBurnNFT := NewMsgBurnNFT(nil, id, denom)
	err := newMsgBurnNFT.ValidateBasic()
	require.Error(t, err)

	newMsgBurnNFT = NewMsgBurnNFT(address, "", denom)
	err = newMsgBurnNFT.ValidateBasic()
	require.Error(t, err)

	newMsgBurnNFT = NewMsgBurnNFT(address, id, "")
	err = newMsgBurnNFT.ValidateBasic()
	require.Error(t, err)

	newMsgBurnNFT = NewMsgBurnNFT(address, id, denom)
	err = newMsgBurnNFT.ValidateBasic()
	require.NoError(t, err)
}

func TestMsgBurnNFTGetSignBytesMethod(t *testing.T) {
	newMsgBurnNFT := NewMsgBurnNFT(address, id, denom)
	sortedBytes := newMsgBurnNFT.GetSignBytes()
	require.Equal(t, string(sortedBytes), fmt.Sprintf(`{"type":"irismod/nft/MsgBurnNFT","value":{"denom":"%s","id":"%s","sender":"%s"}}`,
		denom, id, address.String(),
	))
}

func TestMsgBurnNFTGetSignersMethod(t *testing.T) {
	newMsgBurnNFT := NewMsgBurnNFT(address, id, denom)
	signers := newMsgBurnNFT.GetSigners()
	require.Equal(t, 1, len(signers))
	require.Equal(t, address.String(), signers[0].String())
}
