package keeper

import (
	testUtil "github.com/G4AL-Entertainment/g4al-chain/x/game/testutil"
	"github.com/golang/mock/gomock"
	"testing"

	"github.com/G4AL-Entertainment/g4al-chain/x/game/keeper"
	"github.com/G4AL-Entertainment/g4al-chain/x/game/types"

	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/store"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	capabilitykeeper "github.com/cosmos/cosmos-sdk/x/capability/keeper"
	capabilitytypes "github.com/cosmos/cosmos-sdk/x/capability/types"
	typesparams "github.com/cosmos/cosmos-sdk/x/params/types"
	channeltypes "github.com/cosmos/ibc-go/v5/modules/core/04-channel/types"
	ibcexported "github.com/cosmos/ibc-go/v5/modules/core/exported"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/tendermint/libs/log"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	tmdb "github.com/tendermint/tm-db"
)

// gameChannelKeeper is a stub of cosmosibckeeper.ChannelKeeper.
type gameChannelKeeper struct{}

func (gameChannelKeeper) GetChannel(ctx sdk.Context, srcPort, srcChan string) (channel channeltypes.Channel, found bool) {
	return channeltypes.Channel{}, false
}
func (gameChannelKeeper) GetNextSequenceSend(ctx sdk.Context, portID, channelID string) (uint64, bool) {
	return 0, false
}
func (gameChannelKeeper) SendPacket(ctx sdk.Context, channelCap *capabilitytypes.Capability, packet ibcexported.PacketI) error {
	return nil
}
func (gameChannelKeeper) ChanCloseInit(ctx sdk.Context, portID, channelID string, chanCap *capabilitytypes.Capability) error {
	return nil
}

// gameportKeeper is a stub of cosmosibckeeper.PortKeeper
type gamePortKeeper struct{}

func (gamePortKeeper) BindPort(ctx sdk.Context, portID string) *capabilitytypes.Capability {
	return &capabilitytypes.Capability{}
}

func GameKeeper(t testing.TB) (*keeper.Keeper, sdk.Context) {
	logger := log.NewNopLogger()

	storeKey := sdk.NewKVStoreKey(types.StoreKey)
	memStoreKey := storetypes.NewMemoryStoreKey(types.MemStoreKey)

	db := tmdb.NewMemDB()
	stateStore := store.NewCommitMultiStore(db)
	stateStore.MountStoreWithDB(storeKey, storetypes.StoreTypeIAVL, db)
	stateStore.MountStoreWithDB(memStoreKey, storetypes.StoreTypeMemory, nil)
	require.NoError(t, stateStore.LoadLatestVersion())

	registry := codectypes.NewInterfaceRegistry()
	appCodec := codec.NewProtoCodec(registry)
	capabilityKeeper := capabilitykeeper.NewKeeper(appCodec, storeKey, memStoreKey)

	paramsSubspace := typesparams.NewSubspace(appCodec,
		types.Amino,
		storeKey,
		memStoreKey,
		"GameParams",
	)
	ctrl := gomock.NewController(t)
	accK := testUtil.NewMockAccountKeeper(ctrl)
	permK := testUtil.NewMockPermissionKeeper(ctrl)

	ctx := sdk.NewContext(stateStore, tmproto.Header{}, false, logger)

	// Account expected
	accK.EXPECT().GetAccount(ctx, gomock.Any()).AnyTimes()

	// Permission expected
	permK.EXPECT().ValidateDeveloper(ctx, gomock.Any()).AnyTimes()
	permK.EXPECT().GetDeveloper(ctx, gomock.Any()).AnyTimes()

	k := keeper.NewKeeper(
		appCodec,
		storeKey,
		memStoreKey,
		paramsSubspace,
		gameChannelKeeper{},
		gamePortKeeper{},
		capabilityKeeper.ScopeToModule("GameScopedKeeper"),
		accK,
		permK,
	)

	// Initialize params
	k.SetParams(ctx, types.DefaultParams())

	return k, ctx
}
