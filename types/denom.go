package types

import (
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// Denom defines the data structure of the nft category
type Denom struct {
	Name    string         `json:"name"`
	Schema  string         `json:"schema"`
	Creator sdk.AccAddress `json:"creator"`
}

// NewDenom return a new denom
func NewDenom(name, schema string, creator sdk.AccAddress) Denom {
	return Denom{
		Name:    strings.TrimSpace(name),
		Schema:  strings.TrimSpace(schema),
		Creator: creator,
	}
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
