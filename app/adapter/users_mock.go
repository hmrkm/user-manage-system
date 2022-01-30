// Code generated by MockGen. DO NOT EDIT.
// Source: users.go

// Package adapter is a generated GoMock package.
package adapter

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockUsersAdapter is a mock of UsersAdapter interface.
type MockUsersAdapter struct {
	ctrl     *gomock.Controller
	recorder *MockUsersAdapterMockRecorder
}

// MockUsersAdapterMockRecorder is the mock recorder for MockUsersAdapter.
type MockUsersAdapterMockRecorder struct {
	mock *MockUsersAdapter
}

// NewMockUsersAdapter creates a new mock instance.
func NewMockUsersAdapter(ctrl *gomock.Controller) *MockUsersAdapter {
	mock := &MockUsersAdapter{ctrl: ctrl}
	mock.recorder = &MockUsersAdapterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUsersAdapter) EXPECT() *MockUsersAdapterMockRecorder {
	return m.recorder
}

// List mocks base method.
func (m *MockUsersAdapter) List(req RequestUsersList) (ResponseUsersList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", req)
	ret0, _ := ret[0].(ResponseUsersList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockUsersAdapterMockRecorder) List(req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockUsersAdapter)(nil).List), req)
}
