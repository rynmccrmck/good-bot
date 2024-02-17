// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/rynmccrmck/goodbot (interfaces: NetworkUtils)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockNetworkUtils is a mock of NetworkUtils interface.
type MockNetworkUtils struct {
	ctrl     *gomock.Controller
	recorder *MockNetworkUtilsMockRecorder
}

// MockNetworkUtilsMockRecorder is the mock recorder for MockNetworkUtils.
type MockNetworkUtilsMockRecorder struct {
	mock *MockNetworkUtils
}

// NewMockNetworkUtils creates a new mock instance.
func NewMockNetworkUtils(ctrl *gomock.Controller) *MockNetworkUtils {
	mock := &MockNetworkUtils{ctrl: ctrl}
	mock.recorder = &MockNetworkUtilsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockNetworkUtils) EXPECT() *MockNetworkUtilsMockRecorder {
	return m.recorder
}

// GetASN mocks base method.
func (m *MockNetworkUtils) GetASN(arg0 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetASN", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetASN indicates an expected call of GetASN.
func (mr *MockNetworkUtilsMockRecorder) GetASN(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetASN", reflect.TypeOf((*MockNetworkUtils)(nil).GetASN), arg0)
}

// GetDomainName mocks base method.
func (m *MockNetworkUtils) GetDomainName(arg0 string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDomainName", arg0)
	ret0, _ := ret[0].(string)
	return ret0
}

// GetDomainName indicates an expected call of GetDomainName.
func (mr *MockNetworkUtilsMockRecorder) GetDomainName(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDomainName", reflect.TypeOf((*MockNetworkUtils)(nil).GetDomainName), arg0)
}
