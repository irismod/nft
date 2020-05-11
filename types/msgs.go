package types

import (
	"regexp"
	"strings"

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
func NewMsgIssueDenom(sender sdk.AccAddress, denom, schema string) MsgIssueDenom {
	return MsgIssueDenom{
		Sender: sender,
		Denom:  strings.TrimSpace(denom),
		Schema: strings.TrimSpace(schema),
	}
}

// Route Implements Msg
func (msg MsgIssueDenom) Route() string { return RouterKey }

// Type Implements Msg
func (msg MsgIssueDenom) Type() string { return "issue_denom" }

// ValidateBasic Implements Msg.
func (msg MsgIssueDenom) ValidateBasic() error {
	if err := ValidateDenom(msg.Denom); err != nil {
		return err
	}

	if msg.Sender.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "missing sender address")
	}

	return nil
}

// GetSignBytes Implements Msg.
func (msg MsgIssueDenom) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// GetSigners Implements Msg.
func (msg MsgIssueDenom) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Sender}
}

// NewMsgTransferNFT is a constructor function for MsgSetName
func NewMsgTransferNFT(sender, recipient sdk.AccAddress,
	denom, id, tokenURI, tokenData string) MsgTransferNFT {
	return MsgTransferNFT{
		Sender:    sender,
		Recipient: recipient,
		Denom:     strings.TrimSpace(denom),
		ID:        strings.TrimSpace(id),
		TokenURI:  strings.TrimSpace(tokenURI),
		TokenData: strings.TrimSpace(tokenData),
	}
}

// Route Implements Msg
func (msg MsgTransferNFT) Route() string { return RouterKey }

// Type Implements Msg
func (msg MsgTransferNFT) Type() string { return "transfer_nft" }

// ValidateBasic Implements Msg.
func (msg MsgTransferNFT) ValidateBasic() error {
	if err := ValidateDenom(msg.Denom); err != nil {
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
func NewMsgEditNFT(sender sdk.AccAddress, id,
	denom, tokenURI, tokenData string) MsgEditNFT {
	return MsgEditNFT{
		Sender:    sender,
		Denom:     strings.TrimSpace(denom),
		ID:        strings.TrimSpace(id),
		TokenURI:  strings.TrimSpace(tokenURI),
		TokenData: strings.TrimSpace(tokenData),
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

	if err := ValidateDenom(msg.Denom); err != nil {
		return err
	}

	if err := ValidateTokenURI(msg.TokenURI); err != nil {
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
func NewMsgMintNFT(sender, recipient sdk.AccAddress, id, denom, tokenURI, tokenData string) MsgMintNFT {
	return MsgMintNFT{
		Sender:    sender,
		Recipient: recipient,
		Denom:     strings.TrimSpace(denom),
		ID:        strings.TrimSpace(id),
		TokenURI:  strings.TrimSpace(tokenURI),
		TokenData: strings.TrimSpace(tokenData),
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
	if err := ValidateDenom(msg.Denom); err != nil {
		return err
	}

	if err := ValidateTokenURI(msg.TokenURI); err != nil {
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
	if msg.Sender.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "missing sender address")
	}

	if err := ValidateDenom(msg.Denom); err != nil {
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
