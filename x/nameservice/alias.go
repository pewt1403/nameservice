package nameservice

import (
	"github.com/pewt1403/nameservice/x/nameservice/keeper"
	"github.com/pewt1403/nameservice/x/nameservice/types"
)

const (
	// TODO: define constants that you would like exposed from your module
	ModuleName   = types.ModuleName
	RouterKey    = types.RouterKey
	StoreKey     = types.StoreKey
	QuerierRoute = types.QuerierRoute
)

var (
	// functions aliases
	NewKeeper     = keeper.NewKeeper
	NewQuerier    = keeper.NewQuerier
	RegisterCodec = types.RegisterCodec
	// TODO: Fill out function aliases
	NewMsgBuyName    = types.NewMsgBuyName
	NewMsgSetName    = types.NewMsgSetName
	NewMsgDeleteName = types.NewMsgDeleteName
	NewWhois         = types.NewWhois
	// variable aliases
	ModuleCdc = types.ModuleCdc
	// TODO: Fill out variable aliases
)

type (
	Keeper = keeper.Keeper
	// TODO: Fill out module types
	MsgSetName      = types.MsgSetName
	MsgBuyName      = types.MsgBuyName
	MsgDeleteName   = types.MsgDeleteName
	Whois           = types.Whois
	QueryResResolve = types.QueryResResolve
	QueryResList    = types.QueryResList
)
