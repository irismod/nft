package nft

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// InitGenesis sets nft information for genesis.
func InitGenesis(ctx sdk.Context, k Keeper, data GenesisState) {
	for _, c := range data.Collections {
		if err := k.SetCollection(ctx, c); err != nil {
			panic(err)
		}
	}
}

// ExportGenesis returns a GenesisState for a given context and keeper.
func ExportGenesis(ctx sdk.Context, k Keeper) GenesisState {
	return NewGenesisState(k.GetOwners(ctx), k.GetCollections(ctx))
}

// DefaultGenesisState returns a default genesis state
func DefaultGenesisState() GenesisState {
	return NewGenesisState([]Owner{}, Collections{})
}

// ValidateGenesis performs basic validation of nfts genesis data returning an
// error for any failed validation criteria.
func ValidateGenesis(data GenesisState) error {
	return nil
}
