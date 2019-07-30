// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/dipperin/dipperin-core/core/vm/resolver (interfaces: ContractService)

// Package resolver is a generated GoMock package.
package resolver

import (
	common "github.com/dipperin/dipperin-core/common"
	gomock "github.com/golang/mock/gomock"
	big "math/big"
	reflect "reflect"
)

// MockContractService is a mock of ContractService interface
type MockContractService struct {
	ctrl     *gomock.Controller
	recorder *MockContractServiceMockRecorder
}

// MockContractServiceMockRecorder is the mock recorder for MockContractService
type MockContractServiceMockRecorder struct {
	mock *MockContractService
}

// NewMockContractService creates a new mock instance
func NewMockContractService(ctrl *gomock.Controller) *MockContractService {
	mock := &MockContractService{ctrl: ctrl}
	mock.recorder = &MockContractServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockContractService) EXPECT() *MockContractServiceMockRecorder {
	return m.recorder
}

// Address mocks base method
func (m *MockContractService) Address() common.Address {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Address")
	ret0, _ := ret[0].(common.Address)
	return ret0
}

// Address indicates an expected call of Address
func (mr *MockContractServiceMockRecorder) Address() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Address", reflect.TypeOf((*MockContractService)(nil).Address))
}

// CallValue mocks base method
func (m *MockContractService) CallValue() *big.Int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CallValue")
	ret0, _ := ret[0].(*big.Int)
	return ret0
}

// CallValue indicates an expected call of CallValue
func (mr *MockContractServiceMockRecorder) CallValue() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CallValue", reflect.TypeOf((*MockContractService)(nil).CallValue))
}

// Caller mocks base method
func (m *MockContractService) Caller() common.Address {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Caller")
	ret0, _ := ret[0].(common.Address)
	return ret0
}

// Caller indicates an expected call of Caller
func (mr *MockContractServiceMockRecorder) Caller() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Caller", reflect.TypeOf((*MockContractService)(nil).Caller))
}

// GetGas mocks base method
func (m *MockContractService) GetGas() uint64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetGas")
	ret0, _ := ret[0].(uint64)
	return ret0
}

// GetGas indicates an expected call of GetGas
func (mr *MockContractServiceMockRecorder) GetGas() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGas", reflect.TypeOf((*MockContractService)(nil).GetGas))
}