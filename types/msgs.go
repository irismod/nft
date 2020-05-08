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
	// IsAlphaNumeric only accepts alphanumeric characters
	IsAlphaNumeric   = regexp.MustCompile(`^[a-zA-Z0-9]+$`).MatchString
	IsBeginWithAlpha = regexp.MustCompile(`^[a-zA-Z].*`).MatchString
)

/* --------------------------------------------------------------------------- */
// MsgIssueDenom
/* --------------------------------------------------------------------------- */
type MsgIssueDenom struct {
	Sender sdk.AccAddress `json:"sender",yaml:"sender"`
	Denom  string         `json:"denom",yaml:"denom"`
	Schema string         `json:"schema" yaml:"schema"`
}

// NewMsgTransferNFT is a constructor function for MsgSetName
func NewMsgIssueDenom(sender sdk.AccAddress, denom, metadata string) MsgIssueDenom {
	return MsgIssueDenom{
		Sender: sender,
		Denom:  strings.TrimSpace(denom),
		Schema: strings.TrimSpace(metadata),
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

	//validate Denom
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

/* --------------------------------------------------------------------------- */
// MsgTransferNFT
/* --------------------------------------------------------------------------- */

// MsgTransferNFT defines a TransferNFT message
type MsgTransferNFT struct {
	Sender    sdk.AccAddress `json:"sender",yaml:"sender"`
	Recipient sdk.AccAddress `json:"recipient",yaml:"recipient"`
	Denom     string         `json:"denom",yaml:"denom"`
	ID        string         `json:"id",yaml:"id"`
	TokenURI  string         `json:"token_uri",yaml:"token_uri"`
	Metadata  string         `json:"metadata",yaml:"metadata"`
}

// NewMsgTransferNFT is a constructor function for MsgSetName
func NewMsgTransferNFT(sender, recipient sdk.AccAddress,
	denom, id, tokenURI, metadata string) MsgTransferNFT {
	return MsgTransferNFT{
		Sender:    sender,
		Recipient: recipient,
		Denom:     strings.TrimSpace(denom),
		ID:        strings.TrimSpace(id),
		TokenURI:  strings.TrimSpace(tokenURI),
		Metadata:  strings.TrimSpace(metadata),
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
	Sender   sdk.AccAddress `json:"sender",yaml:"sender"`
	ID       string         `json:"id",yaml:"id"`
	Denom    string         `json:"denom",yaml:"denom"`
	TokenURI string         `json:"token_uri",yaml:"token_uri"`
	Metadata string         `json:"metadata",yaml:"metadata"`
}

// NewMsgEditNFT is a constructor function for MsgSetName
func NewMsgEditNFT(sender sdk.AccAddress, id,
	denom, tokenURI, metadata string) MsgEditNFT {
	return MsgEditNFT{
		Sender:   sender,
		Denom:    strings.TrimSpace(denom),
		ID:       strings.TrimSpace(id),
		TokenURI: strings.TrimSpace(tokenURI),
		Metadata: strings.TrimSpace(metadata),
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

/* --------------------------------------------------------------------------- */
// MsgMintNFT
/* --------------------------------------------------------------------------- */

// MsgMintNFT defines a MintNFT message
type MsgMintNFT struct {
	Sender    sdk.AccAddress `json:"sender",yaml:"sender"`
	Recipient sdk.AccAddress `json:"recipient",yaml:"recipient"`
	Denom     string         `json:"denom",yaml:"denom"`
	ID        string         `json:"id",yaml:"id"`
	TokenURI  string         `json:"token_uri",yaml:"token_uri"`
	Metadata  string         `json:"metadata",yaml:"metadata"`
}

// NewMsgMintNFT is a constructor function for MsgMintNFT
func NewMsgMintNFT(sender, recipient sdk.AccAddress, id, denom, tokenURI, metadata string) MsgMintNFT {
	return MsgMintNFT{
		Sender:    sender,
		Recipient: recipient,
		Denom:     strings.TrimSpace(denom),
		ID:        strings.TrimSpace(id),
		TokenURI:  strings.TrimSpace(tokenURI),
		Metadata:  strings.TrimSpace(metadata),
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

/* --------------------------------------------------------------------------- */
// MsgBurnNFT
/* --------------------------------------------------------------------------- */

// MsgBurnNFT defines a BurnNFT message
type MsgBurnNFT struct {
	Sender sdk.AccAddress `json:"sender",yaml:"sender"`
	Denom  string         `json:"denom",yaml:"denom"`
	ID     string         `json:"id",yaml:"id"`
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
