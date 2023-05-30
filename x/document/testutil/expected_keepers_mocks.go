// Code generated by MockGen. DO NOT EDIT.
// Source: x/document/types/expected_keepers.go

// Package testutil is a generated GoMock package.
package testutil

import (
	reflect "reflect"

	types "github.com/cosmos/cosmos-sdk/types"
	gomock "github.com/golang/mock/gomock"
	types0 "github.com/sharering/shareledger/x/id/types"
)

// MockIDKeeper is a mock of IDKeeper interface.
type MockIDKeeper struct {
	ctrl     *gomock.Controller
	recorder *MockIDKeeperMockRecorder
}

// MockIDKeeperMockRecorder is the mock recorder for MockIDKeeper.
type MockIDKeeperMockRecorder struct {
	mock *MockIDKeeper
}

// NewMockIDKeeper creates a new mock instance.
func NewMockIDKeeper(ctrl *gomock.Controller) *MockIDKeeper {
	mock := &MockIDKeeper{ctrl: ctrl}
	mock.recorder = &MockIDKeeperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIDKeeper) EXPECT() *MockIDKeeperMockRecorder {
	return m.recorder
}

// GetFullIDByIDString mocks base method.
func (m *MockIDKeeper) GetFullIDByIDString(ctx types.Context, id string) (*types0.Id, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFullIDByIDString", ctx, id)
	ret0, _ := ret[0].(*types0.Id)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// GetFullIDByIDString indicates an expected call of GetFullIDByIDString.
func (mr *MockIDKeeperMockRecorder) GetFullIDByIDString(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFullIDByIDString", reflect.TypeOf((*MockIDKeeper)(nil).GetFullIDByIDString), ctx, id)
}
