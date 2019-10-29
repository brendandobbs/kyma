// Code generated by mockery v1.0.0. DO NOT EDIT.
package mocks

import (
	apperrors "github.com/kyma-project/kyma/components/application-registry/internal/apperrors"
	mock "github.com/stretchr/testify/mock"

	model "github.com/kyma-project/kyma/components/application-registry/internal/metadata/model"
)

// Service is an autogenerated mock type for the Service type
type Service struct {
	mock.Mock
}

// GetSpec provides a mock function with given fields: id
func (_m *Service) GetSpec(id string) ([]byte, []byte, []byte, apperrors.AppError) {
	ret := _m.Called(id)

	var r0 []byte
	if rf, ok := ret.Get(0).(func(string) []byte); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	var r1 []byte
	if rf, ok := ret.Get(1).(func(string) []byte); ok {
		r1 = rf(id)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).([]byte)
		}
	}

	var r2 []byte
	if rf, ok := ret.Get(2).(func(string) []byte); ok {
		r2 = rf(id)
	} else {
		if ret.Get(2) != nil {
			r2 = ret.Get(2).([]byte)
		}
	}

	var r3 apperrors.AppError
	if rf, ok := ret.Get(3).(func(string) apperrors.AppError); ok {
		r3 = rf(id)
	} else {
		if ret.Get(3) != nil {
			r3 = ret.Get(3).(apperrors.AppError)
		}
	}

	return r0, r1, r2, r3
}

// PutSpec provides a mock function with given fields: serviceDef, gatewayUrl
func (_m *Service) PutSpec(serviceDef *model.ServiceDefinition, gatewayUrl string) apperrors.AppError {
	ret := _m.Called(serviceDef, gatewayUrl)

	var r0 apperrors.AppError
	if rf, ok := ret.Get(0).(func(*model.ServiceDefinition, string) apperrors.AppError); ok {
		r0 = rf(serviceDef, gatewayUrl)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(apperrors.AppError)
		}
	}

	return r0
}

// RemoveSpec provides a mock function with given fields: id
func (_m *Service) RemoveSpec(id string) apperrors.AppError {
	ret := _m.Called(id)

	var r0 apperrors.AppError
	if rf, ok := ret.Get(0).(func(string) apperrors.AppError); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(apperrors.AppError)
		}
	}

	return r0
}
