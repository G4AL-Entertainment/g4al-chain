package keeper

import (
	"github.com/golang/mock/gomock"
	"testing"

	"github.com/G4AL-Entertainment/g4al-chain/x/assetfactory/keeper"
	testUtil "github.com/G4AL-Entertainment/g4al-chain/x/assetfactory/testutil"
	"github.com/G4AL-Entertainment/g4al-chain/x/assetfactory/types"

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

// assetfactoryChannelKeeper is a stub of cosmosibckeeper.ChannelKeeper.
type assetfactoryChannelKeeper struct{}

func (assetfactoryChannelKeeper) GetChannel(ctx sdk.Context, srcPort, srcChan string) (channel channeltypes.Channel, found bool) {
	return channeltypes.Channel{}, false
}
func (assetfactoryChannelKeeper) GetNextSequenceSend(ctx sdk.Context, portID, channelID string) (uint64, bool) {
	return 0, false
}
func (assetfactoryChannelKeeper) SendPacket(ctx sdk.Context, channelCap *capabilitytypes.Capability, packet ibcexported.PacketI) error {
	return nil
}
func (assetfactoryChannelKeeper) ChanCloseInit(ctx sdk.Context, portID, channelID string, chanCap *capabilitytypes.Capability) error {
	return nil
}

// assetfactoryportKeeper is a stub of cosmosibckeeper.PortKeeper
type assetfactoryPortKeeper struct{}

func (assetfactoryPortKeeper) BindPort(ctx sdk.Context, portID string) *capabilitytypes.Capability {
	return &capabilitytypes.Capability{}
}

func AssetfactoryKeeper(t testing.TB) (*keeper.Keeper, sdk.Context) {
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
		"AssetfactoryParams",
	)
	ctrl := gomock.NewController(t)
	accK := testUtil.NewMockAccountKeeper(ctrl)
	permK := testUtil.NewMockPermissionKeeper(ctrl)
	gameK := testUtil.NewMockGameKeeper(ctrl)
	nftK := testUtil.NewMockNftKeeper(ctrl)

	ctx := sdk.NewContext(stateStore, tmproto.Header{}, false, logger)

	// Account expected
	accK.EXPECT().GetAccount(ctx, gomock.Any()).AnyTimes()

	// Permission expected
	permK.EXPECT().ValidateDeveloper(ctx, gomock.Any()).AnyTimes()
	permK.EXPECT().GetDeveloper(ctx, gomock.Any()).AnyTimes()

	// Game expected
	gameK.EXPECT().GetProject(ctx, gomock.Any()).AnyTimes()
	gameK.EXPECT().ValidateProjectOwnershipOrDelegateByProject(ctx, gomock.Any(), gomock.Any()).AnyTimes()

	// Nft expected
	nftK.EXPECT().Mint(ctx, gomock.Any(), gomock.Any()).AnyTimes()
	nftK.EXPECT().Burn(ctx, gomock.Any(), gomock.Any()).AnyTimes()
	nftK.EXPECT().Transfer(ctx, gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes()
	nftK.EXPECT().SaveClass(ctx, gomock.Any()).AnyTimes()
	nftK.EXPECT().UpdateClass(ctx, gomock.Any()).AnyTimes()
	nftK.EXPECT().Update(ctx, gomock.Any()).AnyTimes()
	// Getters
	nftK.EXPECT().GetNFT(ctx, gomock.Any(), gomock.Any()).AnyTimes()
	nftK.EXPECT().GetClass(ctx, gomock.Any()).AnyTimes()
	nftK.EXPECT().GetClasses(ctx).AnyTimes()
	nftK.EXPECT().GetBalance(ctx, gomock.Any(), gomock.Any()).AnyTimes()
	nftK.EXPECT().GetNFTsOfClass(ctx, gomock.Any()).AnyTimes().AnyTimes()
	nftK.EXPECT().GetNFTsOfClassByOwner(ctx, gomock.Any(), gomock.Any()).AnyTimes()
	nftK.EXPECT().GetOwner(ctx, gomock.Any(), gomock.Any()).AnyTimes()
	nftK.EXPECT().GetTotalSupply(ctx, gomock.Any()).AnyTimes()

	k := keeper.NewKeeper(
		appCodec,
		storeKey,
		memStoreKey,
		paramsSubspace,
		assetfactoryChannelKeeper{},
		assetfactoryPortKeeper{},
		capabilityKeeper.ScopeToModule("AssetfactoryScopedKeeper"),
		accK,
		permK,
		gameK,
		nftK,
	)

	// Initialize params
	k.SetParams(ctx, types.DefaultParams())

	return k, ctx
}
