// Code generated by mockery v2.16.0. DO NOT EDIT.

package mocks

import (
	context "context"

	gorm "gorm.io/gorm"

	mock "github.com/stretchr/testify/mock"
)

// AbstractRepository is an autogenerated mock type for the AbstractRepository type
type AbstractRepository struct {
	mock.Mock
}

// NewQuery provides a mock function with given fields: ctx
func (_m *AbstractRepository) NewQuery(ctx context.Context) *gorm.DB {
	ret := _m.Called(ctx)

	var r0 *gorm.DB
	if rf, ok := ret.Get(0).(func(context.Context) *gorm.DB); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gorm.DB)
		}
	}

	return r0
}

// RunInTransaction provides a mock function with given fields: ctx, fn
func (_m *AbstractRepository) RunInTransaction(ctx context.Context, fn func(context.Context) error) error {
	ret := _m.Called(ctx, fn)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, func(context.Context) error) error); ok {
		r0 = rf(ctx, fn)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewAbstractRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewAbstractRepository creates a new instance of AbstractRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewAbstractRepository(t mockConstructorTestingTNewAbstractRepository) *AbstractRepository {
	mock := &AbstractRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
