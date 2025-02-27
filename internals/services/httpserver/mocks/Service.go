// Code generated by mockery v2.46.3. DO NOT EDIT.

package mocks

import (
	context "context"

	constants "github.com/gndw/starting-golang/internals/constants"

	mock "github.com/stretchr/testify/mock"
)

// Service is an autogenerated mock type for the Service type
type Service struct {
	mock.Mock
}

// RegisterEndpoint provides a mock function with given fields: ctx, method, path, f
func (_m *Service) RegisterEndpoint(ctx context.Context, method string, path string, f constants.HttpFunction) error {
	ret := _m.Called(ctx, method, path, f)

	if len(ret) == 0 {
		panic("no return value specified for RegisterEndpoint")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, constants.HttpFunction) error); ok {
		r0 = rf(ctx, method, path, f)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Serve provides a mock function with given fields: ctx
func (_m *Service) Serve(ctx context.Context) error {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for Serve")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewService creates a new instance of Service. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewService(t interface {
	mock.TestingT
	Cleanup(func())
}) *Service {
	mock := &Service{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
