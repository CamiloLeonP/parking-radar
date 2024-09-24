// Code generated by MockGen. DO NOT EDIT.
// Source: ./sensor_repository.go

// Package mockgen is a generated GoMock package.
package mockgen

import (
	reflect "reflect"

	domain "github.com/CamiloLeonP/parking-radar/internal/app/domain"
	gomock "github.com/golang/mock/gomock"
)

// MockISensorRepository is a mock of ISensorRepository interface.
type MockISensorRepository struct {
	ctrl     *gomock.Controller
	recorder *MockISensorRepositoryMockRecorder
}

// MockISensorRepositoryMockRecorder is the mock recorder for MockISensorRepository.
type MockISensorRepositoryMockRecorder struct {
	mock *MockISensorRepository
}

// NewMockISensorRepository creates a new mock instance.
func NewMockISensorRepository(ctrl *gomock.Controller) *MockISensorRepository {
	mock := &MockISensorRepository{ctrl: ctrl}
	mock.recorder = &MockISensorRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockISensorRepository) EXPECT() *MockISensorRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockISensorRepository) Create(sensor *domain.Sensor) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", sensor)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockISensorRepositoryMockRecorder) Create(sensor interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockISensorRepository)(nil).Create), sensor)
}

// Delete mocks base method.
func (m *MockISensorRepository) Delete(id uint) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockISensorRepositoryMockRecorder) Delete(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockISensorRepository)(nil).Delete), id)
}

// GetByDeviceAndNumber mocks base method.
func (m *MockISensorRepository) GetByDeviceAndNumber(deviceIdentifier string, sensorNumber int) (*domain.Sensor, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByDeviceAndNumber", deviceIdentifier, sensorNumber)
	ret0, _ := ret[0].(*domain.Sensor)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByDeviceAndNumber indicates an expected call of GetByDeviceAndNumber.
func (mr *MockISensorRepositoryMockRecorder) GetByDeviceAndNumber(deviceIdentifier, sensorNumber interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByDeviceAndNumber", reflect.TypeOf((*MockISensorRepository)(nil).GetByDeviceAndNumber), deviceIdentifier, sensorNumber)
}

// GetByID mocks base method.
func (m *MockISensorRepository) GetByID(id uint) (*domain.Sensor, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", id)
	ret0, _ := ret[0].(*domain.Sensor)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID.
func (mr *MockISensorRepositoryMockRecorder) GetByID(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockISensorRepository)(nil).GetByID), id)
}

// ListByEsp32DeviceID mocks base method.
func (m *MockISensorRepository) ListByEsp32DeviceID(esp32DeviceID uint64) ([]domain.Sensor, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListByEsp32DeviceID", esp32DeviceID)
	ret0, _ := ret[0].([]domain.Sensor)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListByEsp32DeviceID indicates an expected call of ListByEsp32DeviceID.
func (mr *MockISensorRepositoryMockRecorder) ListByEsp32DeviceID(esp32DeviceID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListByEsp32DeviceID", reflect.TypeOf((*MockISensorRepository)(nil).ListByEsp32DeviceID), esp32DeviceID)
}

// ListByParkingLot mocks base method.
func (m *MockISensorRepository) ListByParkingLot(parkingLotID uint) ([]domain.Sensor, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListByParkingLot", parkingLotID)
	ret0, _ := ret[0].([]domain.Sensor)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListByParkingLot indicates an expected call of ListByParkingLot.
func (mr *MockISensorRepositoryMockRecorder) ListByParkingLot(parkingLotID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListByParkingLot", reflect.TypeOf((*MockISensorRepository)(nil).ListByParkingLot), parkingLotID)
}

// Update mocks base method.
func (m *MockISensorRepository) Update(sensor *domain.Sensor) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", sensor)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockISensorRepositoryMockRecorder) Update(sensor interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockISensorRepository)(nil).Update), sensor)
}
