package keeper

import (
	"github.com/golang/mock/gomock"
	"testing"

	"github.com/G4AL-Entertainment/g4al-chain/x/denomfactory/keeper"
	testUtil "github.com/G4AL-Entertainment/g4al-chain/x/denomfactory/testutil"
	"github.com/G4AL-Entertainment/g4al-chain/x/denomfactory/types"

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

// denomfactoryChannelKeeper is a stub of cosmosibckeeper.ChannelKeeper.
type denomfactoryChannelKeeper struct{}

func (denomfactoryChannelKeeper) GetChannel(ctx sdk.Context, srcPort, srcChan string) (channel channeltypes.Channel, found bool) {
	return channeltypes.Channel{}, false
}
func (denomfactoryChannelKeeper) GetNextSequenceSend(ctx sdk.Context, portID, channelID string) (uint64, bool) {
	return 0, false
}
func (denomfactoryChannelKeeper) SendPacket(ctx sdk.Context, channelCap *capabilitytypes.Capability, packet ibcexported.PacketI) error {
	return nil
}
func (denomfactoryChannelKeeper) ChanCloseInit(ctx sdk.Context, portID, channelID string, chanCap *capabilitytypes.Capability) error {
	return nil
}

// denomfactoryportKeeper is a stub of cosmosibckeeper.PortKeeper
type denomfactoryPortKeeper struct{}

func (denomfactoryPortKeeper) BindPort(ctx sdk.Context, portID string) *capabilitytypes.Capability {
	return &capabilitytypes.Capability{}
}

func DenomfactoryKeeper(t testing.TB) (*keeper.Keeper, sdk.Context) {
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
		"DenomfactoryParams",
	)

	ctrl := gomock.NewController(t)
	accK := testUtil.NewMockAccountKeeper(ctrl)
	permK := testUtil.NewMockPermissionKeeper(ctrl)
	gameK := testUtil.NewMockGameKeeper(ctrl)
	banK := testUtil.NewMockBankKeeper(ctrl)

	ctx := sdk.NewContext(stateStore, tmproto.Header{}, false, logger)

	// Account expected
	accK.EXPECT().GetAccount(ctx, gomock.Any()).AnyTimes()

	// Permission expected
	permK.EXPECT().ValidateDeveloper(ctx, gomock.Any()).AnyTimes()
	permK.EXPECT().GetDeveloper(ctx, gomock.Any()).AnyTimes()

	// Game expected
	gameK.EXPECT().GetProject(ctx, gomock.Any()).AnyTimes()
	gameK.EXPECT().ValidateProjectOwnershipOrDelegateByProject(ctx, gomock.Any(), gomock.Any()).AnyTimes()

	// Bank expected
	banK.EXPECT().BurnCoins(ctx, gomock.Any(), gomock.Any()).AnyTimes()
	banK.EXPECT().DelegateCoins(ctx, gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes()
	banK.EXPECT().DelegateCoinsFromAccountToModule(ctx, gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes()
	banK.EXPECT().ExportGenesis(gomock.Any()).AnyTimes()
	banK.EXPECT().GetDenomMetaData(ctx, gomock.Any()).AnyTimes()
	banK.EXPECT().GetPaginatedTotalSupply(ctx, gomock.Any()).AnyTimes()
	banK.EXPECT().GetSupply(ctx, gomock.Any()).AnyTimes()
	banK.EXPECT().HasDenomMetaData(ctx, gomock.Any()).AnyTimes()
	banK.EXPECT().HasSupply(ctx, gomock.Any()).AnyTimes()
	banK.EXPECT().InitGenesis(gomock.Any(), gomock.Any()).AnyTimes()
	banK.EXPECT().IterateAllDenomMetaData(ctx, gomock.Any()).AnyTimes()
	banK.EXPECT().IterateTotalSupply(ctx, gomock.Any()).AnyTimes()
	banK.EXPECT().MintCoins(ctx, gomock.Any(), gomock.Any()).AnyTimes()
	banK.EXPECT().SendCoinsFromAccountToModule(ctx, gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes()
	banK.EXPECT().SendCoinsFromModuleToAccount(ctx, gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes()
	banK.EXPECT().SendCoinsFromModuleToModule(ctx, gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes()
	banK.EXPECT().SetDenomMetaData(ctx, gomock.Any()).AnyTimes()
	banK.EXPECT().UndelegateCoins(ctx, gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes()
	banK.EXPECT().UndelegateCoinsFromModuleToAccount(ctx, gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes()
	banK.EXPECT().WithMintCoinsRestriction(gomock.Any()).AnyTimes()

	k := keeper.NewKeeper(
		appCodec,
		storeKey,
		memStoreKey,
		paramsSubspace,
		denomfactoryChannelKeeper{},
		denomfactoryPortKeeper{},
		capabilityKeeper.ScopeToModule("DenomfactoryScopedKeeper"),
		accK,
		permK,
		gameK,
		nil,
	)

	// Initialize params
	k.SetParams(ctx, types.DefaultParams())

	return k, ctx
}
