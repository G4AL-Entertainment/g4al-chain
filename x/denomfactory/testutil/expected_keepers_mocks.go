// Code generated by MockGen. DO NOT EDIT.
// Source: x/denomfactory//types/expected_keepers.go

// Package testutil is a generated GoMock package.
package testutil

import (
	reflect "reflect"

	types "github.com/G4AL-Entertainment/g4al-chain/x/game/types"
	types0 "github.com/G4AL-Entertainment/g4al-chain/x/permission/types"
	types1 "github.com/cosmos/cosmos-sdk/types"
	query "github.com/cosmos/cosmos-sdk/types/query"
	types2 "github.com/cosmos/cosmos-sdk/x/auth/types"
	keeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	types3 "github.com/cosmos/cosmos-sdk/x/bank/types"
	gomock "github.com/golang/mock/gomock"
)

// MockPermissionKeeper is a mock of PermissionKeeper interface.
type MockPermissionKeeper struct {
	ctrl     *gomock.Controller
	recorder *MockPermissionKeeperMockRecorder
}

// MockPermissionKeeperMockRecorder is the mock recorder for MockPermissionKeeper.
type MockPermissionKeeperMockRecorder struct {
	mock *MockPermissionKeeper
}

// NewMockPermissionKeeper creates a new mock instance.
func NewMockPermissionKeeper(ctrl *gomock.Controller) *MockPermissionKeeper {
	mock := &MockPermissionKeeper{ctrl: ctrl}
	mock.recorder = &MockPermissionKeeperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPermissionKeeper) EXPECT() *MockPermissionKeeperMockRecorder {
	return m.recorder
}

// GetDeveloper mocks base method.
func (m *MockPermissionKeeper) GetDeveloper(ctx types1.Context, address string) (types0.Developer, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDeveloper", ctx, address)
	ret0, _ := ret[0].(types0.Developer)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// GetDeveloper indicates an expected call of GetDeveloper.
func (mr *MockPermissionKeeperMockRecorder) GetDeveloper(ctx, address interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDeveloper", reflect.TypeOf((*MockPermissionKeeper)(nil).GetDeveloper), ctx, address)
}

// ValidateDeveloper mocks base method.
func (m *MockPermissionKeeper) ValidateDeveloper(ctx types1.Context, address string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidateDeveloper", ctx, address)
	ret0, _ := ret[0].(error)
	return ret0
}

// ValidateDeveloper indicates an expected call of ValidateDeveloper.
func (mr *MockPermissionKeeperMockRecorder) ValidateDeveloper(ctx, address interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidateDeveloper", reflect.TypeOf((*MockPermissionKeeper)(nil).ValidateDeveloper), ctx, address)
}

// MockGameKeeper is a mock of GameKeeper interface.
type MockGameKeeper struct {
	ctrl     *gomock.Controller
	recorder *MockGameKeeperMockRecorder
}

// MockGameKeeperMockRecorder is the mock recorder for MockGameKeeper.
type MockGameKeeperMockRecorder struct {
	mock *MockGameKeeper
}

// NewMockGameKeeper creates a new mock instance.
func NewMockGameKeeper(ctrl *gomock.Controller) *MockGameKeeper {
	mock := &MockGameKeeper{ctrl: ctrl}
	mock.recorder = &MockGameKeeperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGameKeeper) EXPECT() *MockGameKeeperMockRecorder {
	return m.recorder
}

// GetProject mocks base method.
func (m *MockGameKeeper) GetProject(ctx types1.Context, symbol string) (types.Project, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProject", ctx, symbol)
	ret0, _ := ret[0].(types.Project)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// GetProject indicates an expected call of GetProject.
func (mr *MockGameKeeperMockRecorder) GetProject(ctx, symbol interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProject", reflect.TypeOf((*MockGameKeeper)(nil).GetProject), ctx, symbol)
}

// ValidateProjectOwnershipOrDelegateByProject mocks base method.
func (m *MockGameKeeper) ValidateProjectOwnershipOrDelegateByProject(ctx types1.Context, creator, symbol string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidateProjectOwnershipOrDelegateByProject", ctx, creator, symbol)
	ret0, _ := ret[0].(error)
	return ret0
}

// ValidateProjectOwnershipOrDelegateByProject indicates an expected call of ValidateProjectOwnershipOrDelegateByProject.
func (mr *MockGameKeeperMockRecorder) ValidateProjectOwnershipOrDelegateByProject(ctx, creator, symbol interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidateProjectOwnershipOrDelegateByProject", reflect.TypeOf((*MockGameKeeper)(nil).ValidateProjectOwnershipOrDelegateByProject), ctx, creator, symbol)
}

// MockAccountKeeper is a mock of AccountKeeper interface.
type MockAccountKeeper struct {
	ctrl     *gomock.Controller
	recorder *MockAccountKeeperMockRecorder
}

// MockAccountKeeperMockRecorder is the mock recorder for MockAccountKeeper.
type MockAccountKeeperMockRecorder struct {
	mock *MockAccountKeeper
}

// NewMockAccountKeeper creates a new mock instance.
func NewMockAccountKeeper(ctrl *gomock.Controller) *MockAccountKeeper {
	mock := &MockAccountKeeper{ctrl: ctrl}
	mock.recorder = &MockAccountKeeperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAccountKeeper) EXPECT() *MockAccountKeeperMockRecorder {
	return m.recorder
}

// GetAccount mocks base method.
func (m *MockAccountKeeper) GetAccount(ctx types1.Context, addr types1.AccAddress) types2.AccountI {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAccount", ctx, addr)
	ret0, _ := ret[0].(types2.AccountI)
	return ret0
}

// GetAccount indicates an expected call of GetAccount.
func (mr *MockAccountKeeperMockRecorder) GetAccount(ctx, addr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccount", reflect.TypeOf((*MockAccountKeeper)(nil).GetAccount), ctx, addr)
}

// MockBankKeeper is a mock of BankKeeper interface.
type MockBankKeeper struct {
	ctrl     *gomock.Controller
	recorder *MockBankKeeperMockRecorder
}

// MockBankKeeperMockRecorder is the mock recorder for MockBankKeeper.
type MockBankKeeperMockRecorder struct {
	mock *MockBankKeeper
}

// NewMockBankKeeper creates a new mock instance.
func NewMockBankKeeper(ctrl *gomock.Controller) *MockBankKeeper {
	mock := &MockBankKeeper{ctrl: ctrl}
	mock.recorder = &MockBankKeeperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBankKeeper) EXPECT() *MockBankKeeperMockRecorder {
	return m.recorder
}

// BurnCoins mocks base method.
func (m *MockBankKeeper) BurnCoins(ctx types1.Context, moduleName string, amt types1.Coins) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BurnCoins", ctx, moduleName, amt)
	ret0, _ := ret[0].(error)
	return ret0
}

// BurnCoins indicates an expected call of BurnCoins.
func (mr *MockBankKeeperMockRecorder) BurnCoins(ctx, moduleName, amt interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BurnCoins", reflect.TypeOf((*MockBankKeeper)(nil).BurnCoins), ctx, moduleName, amt)
}

// DelegateCoins mocks base method.
func (m *MockBankKeeper) DelegateCoins(ctx types1.Context, delegatorAddr, moduleAccAddr types1.AccAddress, amt types1.Coins) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DelegateCoins", ctx, delegatorAddr, moduleAccAddr, amt)
	ret0, _ := ret[0].(error)
	return ret0
}

// DelegateCoins indicates an expected call of DelegateCoins.
func (mr *MockBankKeeperMockRecorder) DelegateCoins(ctx, delegatorAddr, moduleAccAddr, amt interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DelegateCoins", reflect.TypeOf((*MockBankKeeper)(nil).DelegateCoins), ctx, delegatorAddr, moduleAccAddr, amt)
}

// DelegateCoinsFromAccountToModule mocks base method.
func (m *MockBankKeeper) DelegateCoinsFromAccountToModule(ctx types1.Context, senderAddr types1.AccAddress, recipientModule string, amt types1.Coins) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DelegateCoinsFromAccountToModule", ctx, senderAddr, recipientModule, amt)
	ret0, _ := ret[0].(error)
	return ret0
}

// DelegateCoinsFromAccountToModule indicates an expected call of DelegateCoinsFromAccountToModule.
func (mr *MockBankKeeperMockRecorder) DelegateCoinsFromAccountToModule(ctx, senderAddr, recipientModule, amt interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DelegateCoinsFromAccountToModule", reflect.TypeOf((*MockBankKeeper)(nil).DelegateCoinsFromAccountToModule), ctx, senderAddr, recipientModule, amt)
}

// ExportGenesis mocks base method.
func (m *MockBankKeeper) ExportGenesis(arg0 types1.Context) *types3.GenesisState {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExportGenesis", arg0)
	ret0, _ := ret[0].(*types3.GenesisState)
	return ret0
}

// ExportGenesis indicates an expected call of ExportGenesis.
func (mr *MockBankKeeperMockRecorder) ExportGenesis(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExportGenesis", reflect.TypeOf((*MockBankKeeper)(nil).ExportGenesis), arg0)
}

// GetDenomMetaData mocks base method.
func (m *MockBankKeeper) GetDenomMetaData(ctx types1.Context, denom string) (types3.Metadata, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDenomMetaData", ctx, denom)
	ret0, _ := ret[0].(types3.Metadata)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// GetDenomMetaData indicates an expected call of GetDenomMetaData.
func (mr *MockBankKeeperMockRecorder) GetDenomMetaData(ctx, denom interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDenomMetaData", reflect.TypeOf((*MockBankKeeper)(nil).GetDenomMetaData), ctx, denom)
}

// GetPaginatedTotalSupply mocks base method.
func (m *MockBankKeeper) GetPaginatedTotalSupply(ctx types1.Context, pagination *query.PageRequest) (types1.Coins, *query.PageResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPaginatedTotalSupply", ctx, pagination)
	ret0, _ := ret[0].(types1.Coins)
	ret1, _ := ret[1].(*query.PageResponse)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetPaginatedTotalSupply indicates an expected call of GetPaginatedTotalSupply.
func (mr *MockBankKeeperMockRecorder) GetPaginatedTotalSupply(ctx, pagination interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPaginatedTotalSupply", reflect.TypeOf((*MockBankKeeper)(nil).GetPaginatedTotalSupply), ctx, pagination)
}

// GetSupply mocks base method.
func (m *MockBankKeeper) GetSupply(ctx types1.Context, denom string) types1.Coin {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSupply", ctx, denom)
	ret0, _ := ret[0].(types1.Coin)
	return ret0
}

// GetSupply indicates an expected call of GetSupply.
func (mr *MockBankKeeperMockRecorder) GetSupply(ctx, denom interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSupply", reflect.TypeOf((*MockBankKeeper)(nil).GetSupply), ctx, denom)
}

// HasDenomMetaData mocks base method.
func (m *MockBankKeeper) HasDenomMetaData(ctx types1.Context, denom string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HasDenomMetaData", ctx, denom)
	ret0, _ := ret[0].(bool)
	return ret0
}

// HasDenomMetaData indicates an expected call of HasDenomMetaData.
func (mr *MockBankKeeperMockRecorder) HasDenomMetaData(ctx, denom interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HasDenomMetaData", reflect.TypeOf((*MockBankKeeper)(nil).HasDenomMetaData), ctx, denom)
}

// HasSupply mocks base method.
func (m *MockBankKeeper) HasSupply(ctx types1.Context, denom string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HasSupply", ctx, denom)
	ret0, _ := ret[0].(bool)
	return ret0
}

// HasSupply indicates an expected call of HasSupply.
func (mr *MockBankKeeperMockRecorder) HasSupply(ctx, denom interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HasSupply", reflect.TypeOf((*MockBankKeeper)(nil).HasSupply), ctx, denom)
}

// InitGenesis mocks base method.
func (m *MockBankKeeper) InitGenesis(arg0 types1.Context, arg1 *types3.GenesisState) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "InitGenesis", arg0, arg1)
}

// InitGenesis indicates an expected call of InitGenesis.
func (mr *MockBankKeeperMockRecorder) InitGenesis(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InitGenesis", reflect.TypeOf((*MockBankKeeper)(nil).InitGenesis), arg0, arg1)
}

// IterateAllDenomMetaData mocks base method.
func (m *MockBankKeeper) IterateAllDenomMetaData(ctx types1.Context, cb func(types3.Metadata) bool) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "IterateAllDenomMetaData", ctx, cb)
}

// IterateAllDenomMetaData indicates an expected call of IterateAllDenomMetaData.
func (mr *MockBankKeeperMockRecorder) IterateAllDenomMetaData(ctx, cb interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IterateAllDenomMetaData", reflect.TypeOf((*MockBankKeeper)(nil).IterateAllDenomMetaData), ctx, cb)
}

// IterateTotalSupply mocks base method.
func (m *MockBankKeeper) IterateTotalSupply(ctx types1.Context, cb func(types1.Coin) bool) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "IterateTotalSupply", ctx, cb)
}

// IterateTotalSupply indicates an expected call of IterateTotalSupply.
func (mr *MockBankKeeperMockRecorder) IterateTotalSupply(ctx, cb interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IterateTotalSupply", reflect.TypeOf((*MockBankKeeper)(nil).IterateTotalSupply), ctx, cb)
}

// MintCoins mocks base method.
func (m *MockBankKeeper) MintCoins(ctx types1.Context, moduleName string, amt types1.Coins) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MintCoins", ctx, moduleName, amt)
	ret0, _ := ret[0].(error)
	return ret0
}

// MintCoins indicates an expected call of MintCoins.
func (mr *MockBankKeeperMockRecorder) MintCoins(ctx, moduleName, amt interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MintCoins", reflect.TypeOf((*MockBankKeeper)(nil).MintCoins), ctx, moduleName, amt)
}

// SendCoinsFromAccountToModule mocks base method.
func (m *MockBankKeeper) SendCoinsFromAccountToModule(ctx types1.Context, senderAddr types1.AccAddress, recipientModule string, amt types1.Coins) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendCoinsFromAccountToModule", ctx, senderAddr, recipientModule, amt)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendCoinsFromAccountToModule indicates an expected call of SendCoinsFromAccountToModule.
func (mr *MockBankKeeperMockRecorder) SendCoinsFromAccountToModule(ctx, senderAddr, recipientModule, amt interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendCoinsFromAccountToModule", reflect.TypeOf((*MockBankKeeper)(nil).SendCoinsFromAccountToModule), ctx, senderAddr, recipientModule, amt)
}

// SendCoinsFromModuleToAccount mocks base method.
func (m *MockBankKeeper) SendCoinsFromModuleToAccount(ctx types1.Context, senderModule string, recipientAddr types1.AccAddress, amt types1.Coins) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendCoinsFromModuleToAccount", ctx, senderModule, recipientAddr, amt)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendCoinsFromModuleToAccount indicates an expected call of SendCoinsFromModuleToAccount.
func (mr *MockBankKeeperMockRecorder) SendCoinsFromModuleToAccount(ctx, senderModule, recipientAddr, amt interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendCoinsFromModuleToAccount", reflect.TypeOf((*MockBankKeeper)(nil).SendCoinsFromModuleToAccount), ctx, senderModule, recipientAddr, amt)
}

// SendCoinsFromModuleToModule mocks base method.
func (m *MockBankKeeper) SendCoinsFromModuleToModule(ctx types1.Context, senderModule, recipientModule string, amt types1.Coins) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendCoinsFromModuleToModule", ctx, senderModule, recipientModule, amt)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendCoinsFromModuleToModule indicates an expected call of SendCoinsFromModuleToModule.
func (mr *MockBankKeeperMockRecorder) SendCoinsFromModuleToModule(ctx, senderModule, recipientModule, amt interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendCoinsFromModuleToModule", reflect.TypeOf((*MockBankKeeper)(nil).SendCoinsFromModuleToModule), ctx, senderModule, recipientModule, amt)
}

// SetDenomMetaData mocks base method.
func (m *MockBankKeeper) SetDenomMetaData(ctx types1.Context, denomMetaData types3.Metadata) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetDenomMetaData", ctx, denomMetaData)
}

// SetDenomMetaData indicates an expected call of SetDenomMetaData.
func (mr *MockBankKeeperMockRecorder) SetDenomMetaData(ctx, denomMetaData interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetDenomMetaData", reflect.TypeOf((*MockBankKeeper)(nil).SetDenomMetaData), ctx, denomMetaData)
}

// UndelegateCoins mocks base method.
func (m *MockBankKeeper) UndelegateCoins(ctx types1.Context, moduleAccAddr, delegatorAddr types1.AccAddress, amt types1.Coins) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UndelegateCoins", ctx, moduleAccAddr, delegatorAddr, amt)
	ret0, _ := ret[0].(error)
	return ret0
}

// UndelegateCoins indicates an expected call of UndelegateCoins.
func (mr *MockBankKeeperMockRecorder) UndelegateCoins(ctx, moduleAccAddr, delegatorAddr, amt interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UndelegateCoins", reflect.TypeOf((*MockBankKeeper)(nil).UndelegateCoins), ctx, moduleAccAddr, delegatorAddr, amt)
}

// UndelegateCoinsFromModuleToAccount mocks base method.
func (m *MockBankKeeper) UndelegateCoinsFromModuleToAccount(ctx types1.Context, senderModule string, recipientAddr types1.AccAddress, amt types1.Coins) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UndelegateCoinsFromModuleToAccount", ctx, senderModule, recipientAddr, amt)
	ret0, _ := ret[0].(error)
	return ret0
}

// UndelegateCoinsFromModuleToAccount indicates an expected call of UndelegateCoinsFromModuleToAccount.
func (mr *MockBankKeeperMockRecorder) UndelegateCoinsFromModuleToAccount(ctx, senderModule, recipientAddr, amt interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UndelegateCoinsFromModuleToAccount", reflect.TypeOf((*MockBankKeeper)(nil).UndelegateCoinsFromModuleToAccount), ctx, senderModule, recipientAddr, amt)
}

// WithMintCoinsRestriction mocks base method.
func (m *MockBankKeeper) WithMintCoinsRestriction(arg0 keeper.MintingRestrictionFn) keeper.BaseKeeper {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WithMintCoinsRestriction", arg0)
	ret0, _ := ret[0].(keeper.BaseKeeper)
	return ret0
}

// WithMintCoinsRestriction indicates an expected call of WithMintCoinsRestriction.
func (mr *MockBankKeeperMockRecorder) WithMintCoinsRestriction(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WithMintCoinsRestriction", reflect.TypeOf((*MockBankKeeper)(nil).WithMintCoinsRestriction), arg0)
}
