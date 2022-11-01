package iochain_test

import (
	"testing"

	keepertest "github.com/mytestlab123/iochain/testutil/keeper"
	"github.com/mytestlab123/iochain/testutil/nullify"
	"github.com/mytestlab123/iochain/x/iochain"
	"github.com/mytestlab123/iochain/x/iochain/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.IochainKeeper(t)
	iochain.InitGenesis(ctx, *k, genesisState)
	got := iochain.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
