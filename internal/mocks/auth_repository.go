// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// AuthRepositoryhMock is an autogenerated mock type for the Auth type
type AuthRepositoryhMock struct {
	mock.Mock
}

// Get provides a mock function with given fields: ctx
func (_m *AuthRepositoryhMock) Get(ctx context.Context) (string, error) {
	ret := _m.Called(ctx)

	var r0 string
	if rf, ok := ret.Get(0).(func(context.Context) string); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Store provides a mock function with given fields: ctx, token
func (_m *AuthRepositoryhMock) Store(ctx context.Context, token string) error {
	ret := _m.Called(ctx, token)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, token)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
