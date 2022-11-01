package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/mytestlab123/iochain/testutil/keeper"
	"github.com/mytestlab123/iochain/x/iochain/keeper"
	"github.com/mytestlab123/iochain/x/iochain/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.IochainKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
