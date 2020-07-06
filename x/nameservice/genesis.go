package nameservice

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

type GenesisState struct {
	WhoisRecords []Whois `json:"whois_records"`
}

func NewGenesisState(whoIsRecords []Whois) GenesisState {
	return GenesisState{WhoisRecords: nil}
}

func ValidateGenesis(data GenesisState) error {
	for _, record := range data.WhoisRecords {
		if record.Owner == nil {
			return fmt.Errorf("invalid WhoisRecord: Value: %s. Error: Missing Owner", record.Value)
		}
		if record.Value == "" {
			return fmt.Errorf("invalid WhoisRecord: Owner: %s. Error: Missing Value", record.Owner)
		}
		if record.Price == nil {
			return fmt.Errorf("invalid WhoisRecord: Value: %s. Error: Missing Price", record.Value)
		}
	}
	return nil
}
func DefaultGenesisState() GenesisState {
	return GenesisState{
		WhoisRecords: []Whois{},
	}
}

// InitGenesis initialize default parameters
// and the keeper's address to pubkey map
// Import to keeper when chain start
func InitGenesis(ctx sdk.Context, keeper Keeper /* TODO: Define what keepers the module needs */, data GenesisState) {
	// TODO: Define logic for when you would like to initalize a new genesis
	for _, record := range data.WhoisRecords {
		keeper.SetWhois(ctx, record.Value, record)
	}
}

// ExportGenesis writes the current store values
// to a genesis file, which can be imported again
// with InitGenesis
// Call adter stopped chain
func ExportGenesis(ctx sdk.Context, k Keeper) (data GenesisState) {
	// TODO: Define logic for exporting state
	//return NewGenesisState()
	var records []Whois
	iterator := k.GetNamesIterator(ctx)
	for ; iterator.Valid(); iterator.Next() {
		name := string(iterator.Key())
		whois := k.GetWhois(ctx, name)
		records = append(records, whois)
	}
	return GenesisState{WhoisRecords: records}
}
