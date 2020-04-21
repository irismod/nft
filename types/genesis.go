package types

// GenesisState is the state that must be provided at genesis.
type GenesisState struct {
	Collections Collections `json:"collections"`
}

// NewGenesisState creates a new genesis state.
func NewGenesisState(collections Collections) GenesisState {
	return GenesisState{
		Collections: collections,
	}
}
