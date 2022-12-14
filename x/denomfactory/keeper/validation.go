package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k Keeper) ValidateArgsDenom(symbol string, description string, name string) error {
	if len(symbol) < SymbolMinLength {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "symbol is needed and must contain at least %d characters", SymbolMinLength)
	}
	if len(symbol) > SymbolMaxLength {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "symbol is needed and can contain at most %d characters", SymbolMaxLength)
	}
	if len(name) < NameMinLength {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "name is needed and must contain at least %d characters", NameMinLength)
	}
	if len(name) > NameMaxLength {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "name is needed and can contain at most %d characters", NameMaxLength)
	}
	if len(description) < DescriptionMinLength {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "description is needed and must contain at least %d characters", DescriptionMinLength)
	}
	if len(description) > DescriptionMaxLength {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "description is needed and must contain at most %d characters", DescriptionMaxLength)
	}
	return nil
}

func (k Keeper) ValidateProjectOwnershipOrDelegateByDenom(ctx sdk.Context, creator string, symbol string) error {
	// Checking if symbol/classID exists via x/nft
	denomFound, found := k.bankKeeper.GetDenomMetaData(ctx, symbol)
	if !found {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "symbol of denom not found x/bank (%s)", symbol)
	}
	// check on map to what project is associated with
	denomMapFound, found := k.GetDenom(ctx, denomFound.Base)
	if !found {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "symbol of denom not found x/assetfactory (%s)", symbol)
	}
	// Check delegate
	err := k.gameKeeper.ValidateProjectOwnershipOrDelegateByProject(ctx, creator, denomMapFound.Project)
	if err != nil {
		return err
	}
	return nil
}
