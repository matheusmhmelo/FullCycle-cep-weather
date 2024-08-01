// Code generated by MockGen. DO NOT EDIT.
// Source: ./internal/infra/gateway/gateway.go

// Package mock_gateway is a generated GoMock package.
package mock_gateway

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockWeatherGatewayInterface is a mock of WeatherGatewayInterface interface.
type MockWeatherGatewayInterface struct {
	ctrl     *gomock.Controller
	recorder *MockWeatherGatewayInterfaceMockRecorder
}

// MockWeatherGatewayInterfaceMockRecorder is the mock recorder for MockWeatherGatewayInterface.
type MockWeatherGatewayInterfaceMockRecorder struct {
	mock *MockWeatherGatewayInterface
}

// NewMockWeatherGatewayInterface creates a new mock instance.
func NewMockWeatherGatewayInterface(ctrl *gomock.Controller) *MockWeatherGatewayInterface {
	mock := &MockWeatherGatewayInterface{ctrl: ctrl}
	mock.recorder = &MockWeatherGatewayInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockWeatherGatewayInterface) EXPECT() *MockWeatherGatewayInterfaceMockRecorder {
	return m.recorder
}

// GetWeather mocks base method.
func (m *MockWeatherGatewayInterface) GetWeather() (float64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetWeather")
	ret0, _ := ret[0].(float64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetWeather indicates an expected call of GetWeather.
func (mr *MockWeatherGatewayInterfaceMockRecorder) GetWeather() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWeather", reflect.TypeOf((*MockWeatherGatewayInterface)(nil).GetWeather))
}

// ValidateLocation mocks base method.
func (m *MockWeatherGatewayInterface) ValidateLocation(cep string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidateLocation", cep)
	ret0, _ := ret[0].(error)
	return ret0
}

// ValidateLocation indicates an expected call of ValidateLocation.
func (mr *MockWeatherGatewayInterfaceMockRecorder) ValidateLocation(cep interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidateLocation", reflect.TypeOf((*MockWeatherGatewayInterface)(nil).ValidateLocation), cep)
}
