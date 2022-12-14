package keeper

import (
	"github.com/G4AL-Entertainment/g4al-chain/x/denomfactory/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
)

const (
	SymbolMinLength      = 3
	SymbolMaxLength      = 20
	NameMinLength        = 8
	NameMaxLength        = 30
	DescriptionMinLength = 10
	DescriptionMaxLength = 250
)

// SetDenom set a specific denom in the store from its index
func (k Keeper) SetDenom(ctx sdk.Context, denom types.Denom) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DenomKeyPrefix))
	b := k.cdc.MustMarshal(&denom)
	store.Set(types.DenomKey(
		denom.Symbol,
	), b)
}

// GetDenom returns a denom from its index
func (k Keeper) GetDenom(
	ctx sdk.Context,
	symbol string,

) (val types.Denom, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DenomKeyPrefix))

	b := store.Get(types.DenomKey(
		symbol,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveDenom removes a denom from the store
func (k Keeper) RemoveDenom(
	ctx sdk.Context,
	symbol string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DenomKeyPrefix))
	store.Delete(types.DenomKey(
		symbol,
	))
}

// GetAllDenom returns all denom
func (k Keeper) GetAllDenom(ctx sdk.Context) (list []types.Denom) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DenomKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Denom
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

func (k Keeper) SetMetadata(ctx sdk.Context, symbol string, name string, description string) {
	// Creating Metadata object
	var metadata = banktypes.Metadata{
		Base:        SymbolToMicroUnit(symbol),
		Display:     symbol,
		Symbol:      symbol,
		Name:        name,
		Description: description,
		//URI: "",
		//URIHash: "",
	}

	// Preparing DenomUnits

	// microUnit
	var microDenomUnit = banktypes.DenomUnit{
		Denom:    SymbolToMicroUnit(symbol),
		Exponent: 0,
	}
	microDenomUnit.Aliases = append(microDenomUnit.Aliases, "micro"+symbol)
	// milliUnit
	var milliDenomUnit = banktypes.DenomUnit{
		Denom:    SymbolToMilliUnit(symbol),
		Exponent: 3,
	}
	milliDenomUnit.Aliases = append(milliDenomUnit.Aliases, "milli"+symbol)
	// baseUnit (no alias)
	var baseDenomUnit = banktypes.DenomUnit{
		Denom:    symbol,
		Exponent: 6,
	}

	// Pushing denomUnits to Metadata
	metadata.DenomUnits = append(metadata.DenomUnits,
		&microDenomUnit,
		&milliDenomUnit,
		&baseDenomUnit,
	)

	// Set metadata object
	k.bankKeeper.SetDenomMetaData(ctx, metadata)
}

func SymbolToMicroUnit(symbol string) string {
	return "u" + symbol
}

func SymbolToMilliUnit(symbol string) string {
	return "m" + symbol
}
