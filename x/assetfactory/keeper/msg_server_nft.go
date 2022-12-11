package keeper

import (
	"context"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/nft"

	"github.com/G4AL-Entertainment/g4al-chain/x/assetfactory/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) MintNft(goCtx context.Context, msg *types.MsgMintNft) (*types.MsgMintNftResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Validation wrapper function
	err := validateMintNft(ctx, k, msg)
	if err != nil {
		return nil, err
	}

	// Validating receiver account as Bech32 address
	bech32, err := sdk.AccAddressFromBech32(msg.Receiver)
	if err != nil {
		return nil, err
	}

	toMint := nft.NFT{
		ClassId: msg.Symbol,
		Id:      string(k.nftKeeper.GetTotalSupply(ctx, msg.Symbol)), // check conversion
		Uri:     msg.Uri,
		UriHash: msg.UriHash,
		//Data: msg.Data, // TODO
	}

	err = k.nftKeeper.Mint(ctx, toMint, bech32)
	if err != nil {
		return nil, err
	}

	return &types.MsgMintNftResponse{}, nil
}

func (k msgServer) UpdateNft(goCtx context.Context, msg *types.MsgUpdateNft) (*types.MsgUpdateNftResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	err := validateUpdateNft(ctx, k, msg)
	if err != nil {
		return nil, err
	}

	toUpdate, found := k.nftKeeper.GetNFT(ctx, msg.Symbol, msg.Id)
	if !found {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "nftID not found (%s)", msg.Id)
	}

	toUpdate.Uri = msg.Uri
	toUpdate.UriHash = msg.UriHash
	//toUpdate.Data = msg.Data / /TODO string to types.Any

	err = k.nftKeeper.Update(ctx, toUpdate)
	if err != nil {
		return nil, err
	}

	return &types.MsgUpdateNftResponse{}, nil
}

func (k msgServer) BurnNft(goCtx context.Context, msg *types.MsgBurnNft) (*types.MsgBurnNftResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	err := validateBurnNft(ctx, k, msg)
	if err != nil {
		return nil, err
	}

	err = k.nftKeeper.Burn(ctx, msg.Symbol, msg.Id)
	if err != nil {
		return nil, err
	}

	return &types.MsgBurnNftResponse{}, nil
}

func (k msgServer) TransferNft(goCtx context.Context, msg *types.MsgTransferNft) (*types.MsgTransferNftResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	err := validateTransferNft(ctx, k, msg)
	if err != nil {
		return nil, err
	}

	bech32, err := sdk.AccAddressFromBech32(msg.Receiver)
	if err != nil {
		return nil, err
	}

	err = k.nftKeeper.Transfer(ctx, msg.Symbol, msg.Id, bech32)
	if err != nil {
		return nil, err
	}

	return &types.MsgTransferNftResponse{}, nil
}

func validateMintNft(ctx sdk.Context, k msgServer, msg *types.MsgMintNft) error {
	err := validateDeveloper(ctx, k, msg.Creator)
	if err != nil {
		return err
	}
	err = validateProjectOwnershipOrDelegateByClassId(ctx, k, msg.Creator, msg.Symbol)
	if err != nil {
		return err
	}
	// check on map to what project is associated with TODO this is repeated, remove
	class, found := k.GetClass(ctx, msg.Symbol)
	if !found {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "symbol of classID not found x/assetfactory (%s)", msg.Symbol)
	}
	// Validate maxSupply only if not 0
	if class.MaxSupply != 0 {
		// If the current total supply is already equal (or higher just in case, but should do not happen)
		if k.nftKeeper.GetTotalSupply(ctx, class.Symbol) >= uint64(class.MaxSupply) {
			return sdkerrors.Wrapf(sdkerrors.ErrUnauthorized, "max supply already reached for class (%s)", msg.Symbol)
		}
	}
	return nil
}

func validateBurnNft(ctx sdk.Context, k msgServer, msg *types.MsgBurnNft) error {
	err := validateDeveloper(ctx, k, msg.Creator)
	if err != nil {
		return err
	}
	err = validateProjectOwnershipOrDelegateByClassId(ctx, k, msg.Creator, msg.Symbol)
	if err != nil {
		return err
	}
	return nil // TODO
}

func validateUpdateNft(ctx sdk.Context, k msgServer, msg *types.MsgUpdateNft) error {
	err := validateDeveloper(ctx, k, msg.Creator)
	if err != nil {
		return err
	}
	err = validateProjectOwnershipOrDelegateByClassId(ctx, k, msg.Creator, msg.Symbol)
	if err != nil {
		return err
	}
	// TODO validate data and params
	return nil
}

func validateTransferNft(ctx sdk.Context, k msgServer, msg *types.MsgTransferNft) error {
	// TODO validate is owner of nft, not project
	// TODO validate allowance (see x/authz if can help with nfts?) (or maybe skip ownership if developer. but not delegate?)
	return nil // TODO
}

func validateDeveloper(ctx sdk.Context, k msgServer, creator string) error {
	// Checking developer role
	val, found := k.permissionKeeper.GetDeveloper(ctx, creator)
	if !found {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "creator invalid developer address (%s)", creator)
	}
	if val.Blocked {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "creator developer address blocked (%s)", creator)
	}
	return nil
}

func validateProjectOwnershipOrDelegateByClassId(ctx sdk.Context, k msgServer, creator string, symbol string) error {
	// Checking if symbol/classID exists via x/nft
	classFound, found := k.nftKeeper.GetClass(ctx, symbol)
	if !found {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "symbol of classID not found x/nft (%s)", symbol)
	}
	// check on map to what project is associated with
	classMapFound, found := k.GetClass(ctx, classFound.Symbol)
	if !found {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "symbol of classID not found x/assetfactory (%s)", symbol)
	}
	// Checking project existing and related to this game developer or delegate
	project, found := k.gameKeeper.GetProject(ctx, classMapFound.Project)
	if !found {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "project invalid symbol (%s)", classMapFound.Project)
	}
	// Check if msg.Creator included in valFound.Delegate
	isDelegate := false
	for _, del := range project.Delegate {
		if del == creator {
			isDelegate = true
			break
		}
	}
	// Checks if the msg creator is the same as the current owner
	if creator != project.Creator && !isDelegate {
		return sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner nor delegate address")
	}
	return nil
}