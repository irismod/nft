package types

import (
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// constant used to indicate that some field should not be updated
const DoNotModify = "[do-not-modify]"

/* --------------------------------------------------------------------------- */
// MsgTransferNFT
/* --------------------------------------------------------------------------- */

// MsgTransferNFT defines a TransferNFT message
type MsgTransferNFT struct {
	Sender    sdk.AccAddress
	Recipient sdk.AccAddress
	TokenURI  string
	Denom     string
	ID        string
}

// NewMsgTransferNFT is a constructor function for MsgSetName
func NewMsgTransferNFT(sender, recipient sdk.AccAddress, denom, id, tokenURI string) MsgTransferNFT {
	return MsgTransferNFT{
		Sender:    sender,
		Recipient: recipient,
		Denom:     strings.TrimSpace(denom),
		ID:        strings.TrimSpace(id),
		TokenURI:  strings.TrimSpace(tokenURI),
	}
}

// Route Implements Msg
func (msg MsgTransferNFT) Route() string { return RouterKey }

// Type Implements Msg
func (msg MsgTransferNFT) Type() string { return "transfer_nft" }

// ValidateBasic Implements Msg.
func (msg MsgTransferNFT) ValidateBasic() error {
	if strings.TrimSpace(msg.Denom) == "" {
		return ErrInvalidCollection
	}
	if msg.Sender.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "missing sender address")
	}
	if msg.Recipient.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "missing recipient address")
	}
	if strings.TrimSpace(msg.ID) == "" {
		return ErrInvalidNFT
	}

	return nil
}

// GetSignBytes Implements Msg.
func (msg MsgTransferNFT) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// GetSigners Implements Msg.
func (msg MsgTransferNFT) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Sender}
}

// Format Implements Msg.
func (msg *MsgTransferNFT) Format() {
	msg.Denom = strings.ToLower(strings.TrimSpace(msg.Denom))
	msg.ID = strings.ToLower(strings.TrimSpace(msg.ID))
	msg.TokenURI = strings.ToLower(strings.TrimSpace(msg.TokenURI))
}

/* --------------------------------------------------------------------------- */
// MsgEditNFT
/* --------------------------------------------------------------------------- */

// MsgEditNFT edits an NFT's metadata
type MsgEditNFT struct {
	Sender   sdk.AccAddress
	ID       string
	Denom    string
	TokenURI string
}

// NewMsgEditNFT is a constructor function for MsgSetName
func NewMsgEditNFT(sender sdk.AccAddress, id,
	denom, tokenURI string,
) MsgEditNFT {
	return MsgEditNFT{
		Sender:   sender,
		Denom:    strings.TrimSpace(denom),
		ID:       strings.TrimSpace(id),
		TokenURI: strings.TrimSpace(tokenURI),
	}
}

// Route Implements Msg
func (msg MsgEditNFT) Route() string { return RouterKey }

// Type Implements Msg
func (msg MsgEditNFT) Type() string { return "edit_nft_metadata" }

// ValidateBasic Implements Msg.
func (msg MsgEditNFT) ValidateBasic() error {
	if msg.Sender.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "missing sender address")
	}
	if strings.TrimSpace(msg.ID) == "" {
		return ErrInvalidNFT
	}
	if strings.TrimSpace(msg.Denom) == "" {
		return ErrInvalidNFT
	}
	return nil
}

// GetSignBytes Implements Msg.
func (msg MsgEditNFT) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// GetSigners Implements Msg.
func (msg MsgEditNFT) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Sender}
}

// Format Implements Msg.
func (msg *MsgEditNFT) Format() {
	msg.Denom = strings.ToLower(strings.TrimSpace(msg.Denom))
	msg.ID = strings.ToLower(strings.TrimSpace(msg.ID))
	msg.TokenURI = strings.ToLower(strings.TrimSpace(msg.TokenURI))
}

/* --------------------------------------------------------------------------- */
// MsgMintNFT
/* --------------------------------------------------------------------------- */

// MsgMintNFT defines a MintNFT message
type MsgMintNFT struct {
	Sender    sdk.AccAddress
	Recipient sdk.AccAddress
	ID        string
	Denom     string
	TokenURI  string
}

// NewMsgMintNFT is a constructor function for MsgMintNFT
func NewMsgMintNFT(sender, recipient sdk.AccAddress, id, denom, tokenURI string) MsgMintNFT {
	return MsgMintNFT{
		Sender:    sender,
		Recipient: recipient,
		Denom:     strings.TrimSpace(denom),
		ID:        strings.TrimSpace(id),
		TokenURI:  strings.TrimSpace(tokenURI),
	}
}

// Route Implements Msg
func (msg MsgMintNFT) Route() string { return RouterKey }

// Type Implements Msg
func (msg MsgMintNFT) Type() string { return "mint_nft" }

// ValidateBasic Implements Msg.
func (msg MsgMintNFT) ValidateBasic() error {
	if strings.TrimSpace(msg.Denom) == "" {
		return ErrInvalidNFT
	}
	if strings.TrimSpace(msg.ID) == "" {
		return ErrInvalidNFT
	}
	if msg.Sender.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "missing sender address")
	}
	if msg.Recipient.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "missing recipient address")
	}
	return nil
}

// GetSignBytes Implements Msg.
func (msg MsgMintNFT) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// GetSigners Implements Msg.
func (msg MsgMintNFT) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Sender}
}

// Format Implements Msg.
func (msg *MsgMintNFT) Format() {
	msg.Denom = strings.ToLower(strings.TrimSpace(msg.Denom))
	msg.ID = strings.ToLower(strings.TrimSpace(msg.ID))
	msg.TokenURI = strings.ToLower(strings.TrimSpace(msg.TokenURI))
}

/* --------------------------------------------------------------------------- */
// MsgBurnNFT
/* --------------------------------------------------------------------------- */

// MsgBurnNFT defines a BurnNFT message
type MsgBurnNFT struct {
	Sender sdk.AccAddress
	ID     string
	Denom  string
}

// NewMsgBurnNFT is a constructor function for MsgBurnNFT
func NewMsgBurnNFT(sender sdk.AccAddress, id string, denom string) MsgBurnNFT {
	return MsgBurnNFT{
		Sender: sender,
		Denom:  strings.TrimSpace(denom),
		ID:     strings.TrimSpace(id),
	}
}

// Route Implements Msg
func (msg MsgBurnNFT) Route() string { return RouterKey }

// Type Implements Msg
func (msg MsgBurnNFT) Type() string { return "burn_nft" }

// ValidateBasic Implements Msg.
func (msg MsgBurnNFT) ValidateBasic() error {
	if strings.TrimSpace(msg.ID) == "" {
		return ErrInvalidNFT
	}
	if strings.TrimSpace(msg.Denom) == "" {
		return ErrInvalidNFT
	}
	if msg.Sender.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "missing sender address")
	}
	return nil
}

// GetSignBytes Implements Msg.
func (msg MsgBurnNFT) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// GetSigners Implements Msg.
func (msg MsgBurnNFT) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Sender}
}

// Format Implements Msg.
func (msg *MsgBurnNFT) Format() {
	msg.Denom = strings.ToLower(strings.TrimSpace(msg.Denom))
	msg.ID = strings.ToLower(strings.TrimSpace(msg.ID))
}
