package keeper_test

import (
	"testing"

	testkeeper "github.com/mytestlab123/iochain/testutil/keeper"
	"github.com/mytestlab123/iochain/x/iochain/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.IochainKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
