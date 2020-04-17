package types

// GenesisState is the state that must be provided at genesis.
type GenesisState struct {
	Owners      Owners      `json:"owners"`
	Collections Collections `json:"collections"`
}

// NewGenesisState creates a new genesis state.
func NewGenesisState(owners []Owner, collections Collections) GenesisState {
	return GenesisState{
		Owners:      owners,
		Collections: collections,
	}
}
