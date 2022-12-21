// Code generated by MockGen. DO NOT EDIT.
// Source: x/game//types/expected_keepers.go

// Package testutil is a generated GoMock package.
package testutil

import (
	reflect "reflect"

	types "github.com/G4AL-Entertainment/g4al-chain/x/permission/types"
	types0 "github.com/cosmos/cosmos-sdk/types"
	types1 "github.com/cosmos/cosmos-sdk/x/auth/types"
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
func (m *MockPermissionKeeper) GetDeveloper(ctx types0.Context, address string) (types.Developer, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDeveloper", ctx, address)
	ret0, _ := ret[0].(types.Developer)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// GetDeveloper indicates an expected call of GetDeveloper.
func (mr *MockPermissionKeeperMockRecorder) GetDeveloper(ctx, address interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDeveloper", reflect.TypeOf((*MockPermissionKeeper)(nil).GetDeveloper), ctx, address)
}

// ValidateDeveloper mocks base method.
func (m *MockPermissionKeeper) ValidateDeveloper(ctx types0.Context, address string) error {
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
func (m *MockAccountKeeper) GetAccount(ctx types0.Context, addr types0.AccAddress) types1.AccountI {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAccount", ctx, addr)
	ret0, _ := ret[0].(types1.AccountI)
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

// SpendableCoins mocks base method.
func (m *MockBankKeeper) SpendableCoins(ctx types0.Context, addr types0.AccAddress) types0.Coins {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SpendableCoins", ctx, addr)
	ret0, _ := ret[0].(types0.Coins)
	return ret0
}

// SpendableCoins indicates an expected call of SpendableCoins.
func (mr *MockBankKeeperMockRecorder) SpendableCoins(ctx, addr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SpendableCoins", reflect.TypeOf((*MockBankKeeper)(nil).SpendableCoins), ctx, addr)
}
