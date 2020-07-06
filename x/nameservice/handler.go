package nameservice

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	k "github.com/pewt1403/nameservice/x/nameservice/keeper"
	"github.com/pewt1403/nameservice/x/nameservice/types"
	loc "github.com/pewt1403/nameservice/x/nameservice/types"
)

// NewHandler creates an sdk.Handler for all the nameservice type messages
func NewHandler(keeper k.Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		ctx = ctx.WithEventManager(sdk.NewEventManager())
		switch msg := msg.(type) {
		// TODO: Define your msg cases
		//
		//Example:
		// case Msg<Action>:
		// 	return handleMsg<Action>(ctx, k, msg)
		// They all do Basicvalidation
		case loc.MsgSetName:
			return handleMsgSetName(ctx, keeper, msg)
		case loc.MsgBuyName:
			return handleMsgBuyName(ctx, keeper, msg)
		case loc.MsgDeleteName:
			return handleMsgDeleteName(ctx, keeper, msg)
		default:
			errMsg := fmt.Sprintf("unrecognized %s message type: %T", ModuleName, msg)
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, errMsg)
		}
	}
}
func handleMsgSetName(ctx sdk.Context, keeper k.Keeper, msg loc.MsgSetName) (*sdk.Result, error) {
	if !msg.Owner.Equals(keeper.GetOwner(ctx, msg.Name)) { // Checks if the the msg sender is the same as the current owner
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Incorrect Owner") // If not, throw an error
	}
	keeper.SetName(ctx, msg.Name, msg.Value) // If so, set the name to the value specified in the msg.
	return &sdk.Result{}, nil                // return
}
func handleMsgBuyName(ctx sdk.Context, keeper k.Keeper, msg loc.MsgBuyName) (*sdk.Result, error) {
	if keeper.GetPrice(ctx, msg.Name).IsAllGT(msg.Bid) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInsufficientFunds, "Bid not high enough")
	}
	// If err != nil => throw error => revert
	if keeper.HasOwner(ctx, msg.Name) {
		err := keeper.CoinKeeper.SendCoins(ctx, msg.Buyer, keeper.GetOwner(ctx, msg.Name), msg.Bid)
		if err != nil {
			return nil, err
		}
	} else {
		_, err := keeper.CoinKeeper.SubtractCoins(ctx, msg.Buyer, msg.Bid)
		if err != nil {
			return nil, err
		}
	}
	keeper.SetOwner(ctx, msg.Name, msg.Buyer)
	keeper.SetPrice(ctx, msg.Name, msg.Bid)
	return &sdk.Result{}, nil
}
func handleMsgDeleteName(ctx sdk.Context, keeper k.Keeper, msg loc.MsgDeleteName) (*sdk.Result, error) {
	if !keeper.IsNamePresent(ctx, msg.Name) {
		return nil, sdkerrors.Wrap(types.ErrNameDoesNotExist, msg.Name)
	}
	if !msg.Owner.Equals(keeper.GetOwner(ctx, msg.Name)) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Incorrect Owner")
	}
	keeper.DeleteWhois(ctx, msg.Name)
	return &sdk.Result{}, nil
}

// handle<Action> does x

/*
func handleMsg<Action>(ctx sdk.Context, k Keeper, msg Msg<Action>) (*sdk.Result, error) {
	err := k.<Action>(ctx, msg.ValidatorAddr)
	if err != nil {
		return nil, err
	}

	// TODO: Define your msg events
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.ValidatorAddr.String()),
		),
	)

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
*/
