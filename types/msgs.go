package types

import (
	"regexp"
	"strings"
	"unicode/utf8"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// constant used to indicate that some field should not be updated
const (
	DoNotModify = "[do-not-modify]"
	MinDenomLen = 3
	MaxDenomLen = 64

	MaxTokenURILen = 256
)

var (
	// IsAlphaNumeric only accepts alphanumeric characters
	IsAlphaNumeric   = regexp.MustCompile(`^[a-zA-Z0-9]+$`).MatchString
	IsBeginWithAlpha = regexp.MustCompile(`^[a-zA-Z].*`).MatchString
)

// NewMsgIssueDenom is a constructor function for MsgSetName
func NewMsgIssueDenom(id, denom, schema string, sender sdk.AccAddress) MsgIssueDenom {
	return MsgIssueDenom{
		Sender: sender,
		ID:     strings.ToLower(strings.TrimSpace(id)),
		Name:   strings.TrimSpace(denom),
		Schema: strings.TrimSpace(schema),
	}
}

// Route Implements Msg
func (m MsgIssueDenom) Route() string { return RouterKey }

// Type Implements Msg
func (m MsgIssueDenom) Type() string { return "issue_denom" }

// ValidateBasic Implements Msg.
func (m MsgIssueDenom) ValidateBasic() error {
	if err := ValidateDenomID(m.ID); err != nil {
		return err
	}

	name := strings.TrimSpace(m.Name)
	if len(name) == 0 {
		return sdkerrors.Wrap(ErrInvalidDenom, "denom name should not be empty")
	}

	if !utf8.ValidString(name) {
		return sdkerrors.Wrap(ErrInvalidDenom, "denom name is invalid")
	}

	if m.Sender.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "missing sender address")
	}
	return nil
}

// GetSignBytes Implements Msg.
func (m MsgIssueDenom) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(m)
	return sdk.MustSortJSON(bz)
}

// GetSigners Implements Msg.
func (m MsgIssueDenom) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{m.Sender}
}

// NewMsgTransferNFT is a constructor function for MsgSetName
func NewMsgTransferNFT(
	id, denom, name, tokenURI, tokenData string,
	sender, recipient sdk.AccAddress) MsgTransferNFT {
	return MsgTransferNFT{
		ID:        strings.ToLower(strings.TrimSpace(id)),
		Denom:     strings.TrimSpace(denom),
		Name:      strings.TrimSpace(name),
		URI:       strings.TrimSpace(tokenURI),
		Data:      strings.TrimSpace(tokenData),
		Sender:    sender,
		Recipient: recipient,
	}
}

// Route Implements Msg
func (msg MsgTransferNFT) Route() string { return RouterKey }

// Type Implements Msg
func (msg MsgTransferNFT) Type() string { return "transfer_nft" }

// ValidateBasic Implements Msg.
func (msg MsgTransferNFT) ValidateBasic() error {
	if err := ValidateDenomID(msg.Denom); err != nil {
		return err
	}

	if msg.Sender.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "missing sender address")
	}

	if msg.Recipient.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "missing recipient address")
	}
	return ValidateTokenID(msg.ID)
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

// NewMsgEditNFT is a constructor function for MsgSetName
func NewMsgEditNFT(
	id, denom, name, tokenURI, tokenData string, sender sdk.AccAddress) MsgEditNFT {
	return MsgEditNFT{
		ID:     strings.ToLower(strings.TrimSpace(id)),
		Denom:  strings.TrimSpace(denom),
		Name:   strings.TrimSpace(name),
		URI:    strings.TrimSpace(tokenURI),
		Data:   strings.TrimSpace(tokenData),
		Sender: sender,
	}
}

// Route Implements Msg
func (msg MsgEditNFT) Route() string { return RouterKey }

// Type Implements Msg
func (msg MsgEditNFT) Type() string { return "edit_nft" }

// ValidateBasic Implements Msg.
func (msg MsgEditNFT) ValidateBasic() error {
	if msg.Sender.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "missing sender address")
	}

	if err := ValidateDenomID(msg.Denom); err != nil {
		return err
	}

	if err := ValidateTokenURI(msg.URI); err != nil {
		return err
	}
	return ValidateTokenID(msg.ID)
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

// NewMsgMintNFT is a constructor function for MsgMintNFT
func NewMsgMintNFT(
	id, denom, name, tokenURI, tokenData string,
	sender, recipient sdk.AccAddress) MsgMintNFT {
	return MsgMintNFT{
		ID:        strings.ToLower(strings.TrimSpace(id)),
		Denom:     strings.TrimSpace(denom),
		Name:      strings.TrimSpace(name),
		URI:       strings.TrimSpace(tokenURI),
		Data:      strings.TrimSpace(tokenData),
		Sender:    sender,
		Recipient: recipient,
	}
}

// Route Implements Msg
func (msg MsgMintNFT) Route() string { return RouterKey }

// Type Implements Msg
func (msg MsgMintNFT) Type() string { return "mint_nft" }

// ValidateBasic Implements Msg.
func (msg MsgMintNFT) ValidateBasic() error {
	if msg.Sender.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "missing sender address")
	}
	if msg.Recipient.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "missing receipt address")
	}
	if err := ValidateDenomID(msg.Denom); err != nil {
		return err
	}

	if err := ValidateTokenURI(msg.URI); err != nil {
		return err
	}
	return ValidateTokenID(msg.ID)
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

// NewMsgBurnNFT is a constructor function for MsgBurnNFT
func NewMsgBurnNFT(sender sdk.AccAddress, id string, denom string) MsgBurnNFT {
	return MsgBurnNFT{
		Sender: sender,
		ID:     strings.ToLower(strings.TrimSpace(id)),
		Denom:  strings.TrimSpace(denom),
	}
}

// Route Implements Msg
func (msg MsgBurnNFT) Route() string { return RouterKey }

// Type Implements Msg
func (msg MsgBurnNFT) Type() string { return "burn_nft" }

// ValidateBasic Implements Msg.
func (msg MsgBurnNFT) ValidateBasic() error {
	if msg.Sender.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "missing sender address")
	}

	if err := ValidateDenomID(msg.Denom); err != nil {
		return err
	}
	return ValidateTokenID(msg.ID)
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
