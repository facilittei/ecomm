// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	context "context"

	payment "github.com/facilittei/ecomm/internal/domains/payment"
	mock "github.com/stretchr/testify/mock"
)

// ChargeRepositoryMock is an autogenerated mock type for the Charge type
type ChargeRepositoryMock struct {
	mock.Mock
}

// Store provides a mock function with given fields: ctx, charge
func (_m *ChargeRepositoryMock) Store(ctx context.Context, charge payment.Charge) error {
	ret := _m.Called(ctx, charge)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, payment.Charge) error); ok {
		r0 = rf(ctx, charge)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
