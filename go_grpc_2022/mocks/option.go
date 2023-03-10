// Code generated by mockery v2.16.0. DO NOT EDIT.

package mocks

import (
	gorm "gorm.io/gorm"

	mock "github.com/stretchr/testify/mock"
)

// Option is an autogenerated mock type for the Option type
type Option struct {
	mock.Mock
}

// Apply provides a mock function with given fields: db
func (_m *Option) Apply(db *gorm.DB) *gorm.DB {
	ret := _m.Called(db)

	var r0 *gorm.DB
	if rf, ok := ret.Get(0).(func(*gorm.DB) *gorm.DB); ok {
		r0 = rf(db)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gorm.DB)
		}
	}

	return r0
}

type mockConstructorTestingTNewOption interface {
	mock.TestingT
	Cleanup(func())
}

// NewOption creates a new instance of Option. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewOption(t mockConstructorTestingTNewOption) *Option {
	mock := &Option{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
