// Code generated by MockGen. DO NOT EDIT.
// Source: ./parking_lot_uc.go

// Package mockgen is a generated GoMock package.
package mockgen

import (
	reflect "reflect"

	usecase "github.com/CamiloLeonP/parking-radar/internal/app/usecase"
	gomock "github.com/golang/mock/gomock"
)

// MockIParkingLotUseCase is a mock of IParkingLotUseCase interface.
type MockIParkingLotUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockIParkingLotUseCaseMockRecorder
}

// MockIParkingLotUseCaseMockRecorder is the mock recorder for MockIParkingLotUseCase.
type MockIParkingLotUseCaseMockRecorder struct {
	mock *MockIParkingLotUseCase
}

// NewMockIParkingLotUseCase creates a new mock instance.
func NewMockIParkingLotUseCase(ctrl *gomock.Controller) *MockIParkingLotUseCase {
	mock := &MockIParkingLotUseCase{ctrl: ctrl}
	mock.recorder = &MockIParkingLotUseCaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIParkingLotUseCase) EXPECT() *MockIParkingLotUseCaseMockRecorder {
	return m.recorder
}

// CreateParkingLot mocks base method.
func (m *MockIParkingLotUseCase) CreateParkingLot(req usecase.CreateParkingLotRequest) (*usecase.ParkingLotResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateParkingLot", req)
	ret0, _ := ret[0].(*usecase.ParkingLotResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateParkingLot indicates an expected call of CreateParkingLot.
func (mr *MockIParkingLotUseCaseMockRecorder) CreateParkingLot(req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateParkingLot", reflect.TypeOf((*MockIParkingLotUseCase)(nil).CreateParkingLot), req)
}

// DeleteParkingLot mocks base method.
func (m *MockIParkingLotUseCase) DeleteParkingLot(parkingLotID uint) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteParkingLot", parkingLotID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteParkingLot indicates an expected call of DeleteParkingLot.
func (mr *MockIParkingLotUseCaseMockRecorder) DeleteParkingLot(parkingLotID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteParkingLot", reflect.TypeOf((*MockIParkingLotUseCase)(nil).DeleteParkingLot), parkingLotID)
}

// GetParkingLot mocks base method.
func (m *MockIParkingLotUseCase) GetParkingLot(parkingLotID uint) (*usecase.ParkingLotResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetParkingLot", parkingLotID)
	ret0, _ := ret[0].(*usecase.ParkingLotResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetParkingLot indicates an expected call of GetParkingLot.
func (mr *MockIParkingLotUseCaseMockRecorder) GetParkingLot(parkingLotID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetParkingLot", reflect.TypeOf((*MockIParkingLotUseCase)(nil).GetParkingLot), parkingLotID)
}

// ListParkingLots mocks base method.
func (m *MockIParkingLotUseCase) ListParkingLots() ([]usecase.ParkingLotResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListParkingLots")
	ret0, _ := ret[0].([]usecase.ParkingLotResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListParkingLots indicates an expected call of ListParkingLots.
func (mr *MockIParkingLotUseCaseMockRecorder) ListParkingLots() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListParkingLots", reflect.TypeOf((*MockIParkingLotUseCase)(nil).ListParkingLots))
}

// UpdateParkingLot mocks base method.
func (m *MockIParkingLotUseCase) UpdateParkingLot(parkingLotID uint, req usecase.UpdateParkingLotRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateParkingLot", parkingLotID, req)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateParkingLot indicates an expected call of UpdateParkingLot.
func (mr *MockIParkingLotUseCaseMockRecorder) UpdateParkingLot(parkingLotID, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateParkingLot", reflect.TypeOf((*MockIParkingLotUseCase)(nil).UpdateParkingLot), parkingLotID, req)
}
