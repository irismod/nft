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
	MaxDenomLen = 8

	MaxTokenURILen = 256
)

var (
	IsAlphaNumeric   = regexp.MustCompile(`^[a-zA-Z0-9]+$`).MatchString // only accepts alphanumeric characters
	IsBeginWithAlpha = regexp.MustCompile(`^[a-zA-Z].*`).MatchString
)

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
func (msg MsgMintNFT) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// GetSigners Implements Msg.
func (msg MsgMintNFT) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Sender}
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

func ValidateDenom(denom string) error {
	denom = strings.TrimSpace(denom)
	if len(denom) < MinDenomLen || len(denom) > MaxDenomLen {
		return sdkerrors.Wrapf(ErrInvalidDenom, "invalid denom %s, only accepts value [%d, %d]", denom, MinDenomLen, MaxDenomLen)
	}
	if !IsBeginWithAlpha(denom) || !IsAlphaNumeric(denom) {
		return sdkerrors.Wrapf(ErrInvalidDenom, "invalid denom %s, only accepts alphanumeric characters,and begin with an english letter", denom)
	}
	return nil
}

func ValidateTokenID(tokenID string) error {
	tokenID = strings.TrimSpace(tokenID)
	if len(tokenID) < MinDenomLen || len(tokenID) > MaxDenomLen {
		return sdkerrors.Wrapf(ErrInvalidTokenID, "invalid tokenID %s, only accepts value [%d, %d]", denom, MinDenomLen, MaxDenomLen)
	}
	if !IsBeginWithAlpha(tokenID) || !IsAlphaNumeric(tokenID) {
		return sdkerrors.Wrapf(ErrInvalidTokenID, "invalid tokenID %s, only accepts alphanumeric characters,and begin with an english letter", denom)
	}
	return nil
}

func ValidateTokenURI(tokenURI string) error {
	if len(tokenURI) > MaxTokenURILen {
		return sdkerrors.Wrapf(ErrInvalidTokenURI, "invalid tokenURI %s, only accepts value [0, %d]", denom, MaxTokenURILen)
	}
	return nil
}
