package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/pewt1403/nameservice/x/nameservice/types"
)

// Keeper of the nameservice store
type Keeper struct {
	CoinKeeper types.BankKeeper
	storeKey   sdk.StoreKey
	cdc        *codec.Codec
}

func (k Keeper) SetWhois(ctx sdk.Context, name string, whois types.Whois) {
	if whois.Owner.Empty() {
		return
	}
	store := ctx.KVStore(k.storeKey) //Get Store Object
	store.Set([]byte(name), k.cdc.MustMarshalBinaryBare(whois)) //INSERT map[name]Whois to the store which Whois is encoded with amino
}

func (k Keeper) GetWhois(ctx sdk.Context, name string) type.Whois{
	store := ctx.KVStore(k.storeKey)
	if !k.IsNamePresent(ctx, name){ //If name not exist => create new Whois with min price
		return types.NewWhois()
	}
	bz := store.Get([]byte(name)) // Cast name to byte
	var whois types.Whois
	k.cdc.MustUnmarshalBinaryBare(bz, &whois) //Unmarshal Whois
	return whois
}

func (k Keeper) DeleteWhois(ctx sdk.Context, name string){
	store := ctx.KVStore(k.storeKey)
	store.Delete([]byte(name))
}

func (k.Keeper) ResolveName(ctx sdk.Context, name string) string{
	return k.GetWhois(ctx, name).Value
}

func (k.Keeper) SetName(ctx sdk.Context, name string, value string){
	whois := k.GetWhois(ctx, name)
	whois.Value = value
	k.SetWhois(ctx, name, whois)
}

func (k.Keeper) HasOwner(ctx sdk.Context, name string) bool {
	return !k.GetWhois(ctx, name).Owner.Empty()
}

func (k.Keeper) SetOwner(ctx sdk.Context, name string, owner sdk.AccAddress){
	whois := k.GetWhois(ctx,name)
	whois.Owner = owner
	k.SetWhois(ctx, name, whois)
}
