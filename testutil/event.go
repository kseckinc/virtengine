package testutil

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/virtengine/virtengine/sdkutil"
	dtypes "github.com/virtengine/virtengine/x/deployment/types"
	mtypes "github.com/virtengine/virtengine/x/market/types"
	ptypes "github.com/virtengine/virtengine/x/provider/types"
	"github.com/stretchr/testify/require"
	abci "github.com/tendermint/tendermint/abci/types"
)

func ParseEvent(t testing.TB, events []abci.Event) sdkutil.Event {
	t.Helper()

	require.Equal(t, 1, len(events))

	sev := sdk.StringifyEvent(events[0])
	ev, err := sdkutil.ParseEvent(sev)

	require.NoError(t, err)

	return ev
}

func ParseDeploymentEvent(t testing.TB, events []abci.Event) sdkutil.ModuleEvent {
	t.Helper()

	uev := ParseEvent(t, events)

	iev, err := dtypes.ParseEvent(uev)
	require.NoError(t, err)

	return iev
}

func ParseMarketEvent(t testing.TB, events []abci.Event) sdkutil.ModuleEvent {
	t.Helper()

	uev := ParseEvent(t, events)

	iev, err := mtypes.ParseEvent(uev)
	require.NoError(t, err)

	return iev
}

func ParseProviderEvent(t testing.TB, events []abci.Event) sdkutil.ModuleEvent {
	t.Helper()

	uev := ParseEvent(t, events)

	iev, err := ptypes.ParseEvent(uev)
	require.NoError(t, err)

	return iev
}
