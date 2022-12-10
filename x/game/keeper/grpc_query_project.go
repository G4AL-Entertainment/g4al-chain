package keeper

import (
	"context"

	"github.com/G4AL-Entertainment/g4al-chain/x/game/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) ProjectAll(c context.Context, req *types.QueryAllProjectRequest) (*types.QueryAllProjectResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var projects []types.Project
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	projectStore := prefix.NewStore(store, types.KeyPrefix(types.ProjectKeyPrefix))

	pageRes, err := query.Paginate(projectStore, req.Pagination, func(key []byte, value []byte) error {
		var project types.Project
		if err := k.cdc.Unmarshal(value, &project); err != nil {
			return err
		}

		projects = append(projects, project)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllProjectResponse{Project: projects, Pagination: pageRes}, nil
}

func (k Keeper) Project(c context.Context, req *types.QueryGetProjectRequest) (*types.QueryGetProjectResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetProject(
		ctx,
		req.Symbol,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetProjectResponse{Project: val}, nil
}