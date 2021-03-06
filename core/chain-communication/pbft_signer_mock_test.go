// Copyright 2019, Keychain Foundation Ltd.
// This file is part of the dipperin-core library.
//
// The dipperin-core library is free software: you can redistribute
// it and/or modify it under the terms of the GNU Lesser General Public License
// as published by the Free Software Foundation, either version 3 of the
// License, or (at your option) any later version.
//
// The dipperin-core library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

// Automatically generated by MockGen. DO NOT EDIT!
// Source: github.com/caiqingfeng/dipperin-core/core/chain-communication (interfaces: PbftSigner)

package chain_communication

import (
	ecdsa "crypto/ecdsa"
	common "github.com/dipperin/dipperin-core/common"
	accounts "github.com/dipperin/dipperin-core/core/accounts"
	gomock "github.com/golang/mock/gomock"
)

// Mock of PbftSigner interface
type MockPbftSigner struct {
	ctrl     *gomock.Controller
	recorder *_MockPbftSignerRecorder
}

// Recorder for MockPbftSigner (not exported)
type _MockPbftSignerRecorder struct {
	mock *MockPbftSigner
}

func NewMockPbftSigner(ctrl *gomock.Controller) *MockPbftSigner {
	mock := &MockPbftSigner{ctrl: ctrl}
	mock.recorder = &_MockPbftSignerRecorder{mock}
	return mock
}

func (_m *MockPbftSigner) EXPECT() *_MockPbftSignerRecorder {
	return _m.recorder
}

func (_m *MockPbftSigner) Evaluate(_param0 accounts.Account, _param1 []byte) ([32]byte, []byte, error) {
	ret := _m.ctrl.Call(_m, "Evaluate", _param0, _param1)
	ret0, _ := ret[0].([32]byte)
	ret1, _ := ret[1].([]byte)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

func (_mr *_MockPbftSignerRecorder) Evaluate(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Evaluate", arg0, arg1)
}

func (_m *MockPbftSigner) GetAddress() common.Address {
	ret := _m.ctrl.Call(_m, "GetAddress")
	ret0, _ := ret[0].(common.Address)
	return ret0
}

func (_mr *_MockPbftSignerRecorder) GetAddress() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetAddress")
}

func (_m *MockPbftSigner) PublicKey() *ecdsa.PublicKey {
	ret := _m.ctrl.Call(_m, "PublicKey")
	ret0, _ := ret[0].(*ecdsa.PublicKey)
	return ret0
}

func (_mr *_MockPbftSignerRecorder) PublicKey() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "PublicKey")
}

func (_m *MockPbftSigner) SetBaseAddress(_param0 common.Address) {
	_m.ctrl.Call(_m, "SetBaseAddress", _param0)
}

func (_mr *_MockPbftSignerRecorder) SetBaseAddress(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SetBaseAddress", arg0)
}

func (_m *MockPbftSigner) SignHash(_param0 []byte) ([]byte, error) {
	ret := _m.ctrl.Call(_m, "SignHash", _param0)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockPbftSignerRecorder) SignHash(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SignHash", arg0)
}

func (_m *MockPbftSigner) ValidSign(_param0 []byte, _param1 []byte, _param2 []byte) error {
	ret := _m.ctrl.Call(_m, "ValidSign", _param0, _param1, _param2)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockPbftSignerRecorder) ValidSign(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ValidSign", arg0, arg1, arg2)
}
