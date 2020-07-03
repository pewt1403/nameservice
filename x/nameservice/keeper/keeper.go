package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/pewt1403/cosmos_module/x/nameservice/types"
)

// Keeper of the nameservice store
type Keeper struct {
	CoinKeeper types.
	storeKey   sdk.StoreKey
	cdc        *codec.Codec
}

func (k Keeper) SetWhois(ctx sdk.Context, name string, whois types.Whois) {
	if whois.Owner.Empty() {
		return
	}
	store := ctx.KVStore(k.storeKey)
	store.Set([]byte(name), k.cdc.MustMarshalBinaryBare(whois))
}
