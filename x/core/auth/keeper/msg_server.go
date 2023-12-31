package keeper

import (
	"context"

	"github.com/datachainlab/cross/x/core/auth/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/datachainlab/cross/x/core/auth/types"
	xcctypes "github.com/datachainlab/cross/x/core/xcc/types"
	"github.com/datachainlab/cross/x/packets"
)

var _ types.MsgServer = (*Keeper)(nil)

// SignTx defines a rpc handler method for MsgSignTx.
func (k Keeper) SignTx(goCtx context.Context, msg *types.MsgSignTx) (*types.MsgSignTxResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	var accounts []authtypes.Account
	for _, addr := range msg.Signers {
		accounts = append(accounts, authtypes.NewAccount(addr, authtypes.NewAuthTypeLocal()))
	}
	completed, err := k.Sign(ctx, msg.TxID, accounts)
	if err != nil {
		return nil, err
	}
	res := &types.MsgSignTxResponse{TxAuthCompleted: completed}
	if completed {
		if err := k.txManager.OnPostAuth(ctx, msg.TxID); err != nil {
			k.Logger(ctx).Error("failed to call PostAuth", "err", err)
			res.Log = err.Error()
		}
	}
	return res, nil
}

// IBCSignTx defines a rpc handler method for MsgIBCSignTx.
func (k Keeper) IBCSignTx(goCtx context.Context, msg *types.MsgIBCSignTx) (*types.MsgIBCSignTxResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	xcc, err := xcctypes.UnpackCrossChainChannel(k.m, *msg.CrossChainChannel)
	if err != nil {
		return nil, err
	}

	// Run packet middlewares

	ctx, ps, err := k.packetMiddleware.HandleMsg(ctx, msg, packets.NewBasicPacketSender(k.channelKeeper))
	if err != nil {
		return nil, err
	}

	err = k.SendIBCSignTx(
		ctx,
		ps,
		xcc,
		msg.TxID,
		msg.Signers,
		msg.TimeoutHeight,
		msg.TimeoutTimestamp,
	)
	if err != nil {
		return nil, err
	}
	return &types.MsgIBCSignTxResponse{}, nil
}

func (k Keeper) ExtSignTx(goCtx context.Context, msg *types.MsgExtSignTx) (*types.MsgExtSignTxResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	completed, err := k.Sign(ctx, msg.TxID, msg.GetSignerAccounts())
	if err != nil {
		return nil, err
	}
	if completed {
		if err := k.txManager.OnPostAuth(ctx, msg.TxID); err != nil {
			k.Logger(ctx).Error("failed to call PostAuth", "err", err)
		}
	}
	return &types.MsgExtSignTxResponse{}, nil
}
