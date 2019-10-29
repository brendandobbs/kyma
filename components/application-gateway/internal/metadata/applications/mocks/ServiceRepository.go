// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	applications "github.com/kyma-project/kyma/components/application-gateway/internal/metadata/applications"
	apperrors "github.com/kyma-project/kyma/components/application-gateway/pkg/apperrors"

	mock "github.com/stretchr/testify/mock"
)

// ServiceRepository is an autogenerated mock type for the ServiceRepository type
type ServiceRepository struct {
	mock.Mock
}

// Get provides a mock function with given fields: id
func (_m *ServiceRepository) Get(id string) (applications.Service, apperrors.AppError) {
	ret := _m.Called(id)

	var r0 applications.Service
	if rf, ok := ret.Get(0).(func(string) applications.Service); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(applications.Service)
	}

	var r1 apperrors.AppError
	if rf, ok := ret.Get(1).(func(string) apperrors.AppError); ok {
		r1 = rf(id)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(apperrors.AppError)
		}
	}

	return r0, r1
}
