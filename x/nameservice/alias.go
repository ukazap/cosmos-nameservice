package nameservice

import (
	"github.com/ukazap/cosmos-nameservice/x/nameservice/internal/keeper"
	"github.com/ukazap/cosmos-nameservice/x/nameservice/internal/types"
)

const (
	// TODO: define constants that you would like exposed from the internal package

	ModuleName        = types.ModuleName
	RouterKey         = types.RouterKey
	StoreKey          = types.StoreKey
	DefaultParamspace = types.DefaultParamspace
	QueryParams       = types.QueryParams
	QuerierRoute      = types.QuerierRoute
)

var (
	// functions aliases
    NewKeeper           = keeper.NewKeeper
    NewQuerier          = keeper.NewQuerier
    RegisterCodec       = types.RegisterCodec
    NewGenesisState     = types.NewGenesisState
    DefaultGenesisState = types.DefaultGenesisState
    ValidateGenesis     = types.ValidateGenesis
	NewMsgBuyName       = types.NewMsgBuyName
    NewMsgSetName       = types.NewMsgSetName
    NewMsgDeleteName    = types.NewMsgDeleteName
    NewWhois            = types.NewWhois

	// variable aliases
    ModuleCdc           = types.ModuleCdc
)

type (
	Keeper       = keeper.Keeper
	GenesisState = types.GenesisState
	Params       = types.Params

	MsgSetName      = types.MsgSetName
    MsgBuyName      = types.MsgBuyName
    MsgDeleteName   = types.MsgDeleteName
    QueryResResolve = types.QueryResResolve
    QueryResNames   = types.QueryResNames
    Whois           = types.Whois
)
