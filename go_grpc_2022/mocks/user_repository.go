// Code generated by mockery v2.16.0. DO NOT EDIT.

package mocks

import (
	entity "challenge/internal/entity"
	context "context"

	gorm "gorm.io/gorm"

	gormprovider "challenge/pkg/gormprovider"

	mock "github.com/stretchr/testify/mock"
)

// UserRepository is an autogenerated mock type for the UserRepository type
type UserRepository struct {
	mock.Mock
}

// CreateUser provides a mock function with given fields: ctx, user, password
func (_m *UserRepository) CreateUser(ctx context.Context, user entity.User, password string) (entity.User, error) {
	ret := _m.Called(ctx, user, password)

	var r0 entity.User
	if rf, ok := ret.Get(0).(func(context.Context, entity.User, string) entity.User); ok {
		r0 = rf(ctx, user, password)
	} else {
		r0 = ret.Get(0).(entity.User)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, entity.User, string) error); ok {
		r1 = rf(ctx, user, password)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteUser provides a mock function with given fields: ctx, id
func (_m *UserRepository) DeleteUser(ctx context.Context, id string) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ListUsers provides a mock function with given fields: ctx, paginationToken, pageSize, opts
func (_m *UserRepository) ListUsers(ctx context.Context, paginationToken string, pageSize uint, opts ...gormprovider.Option) ([]entity.User, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, paginationToken, pageSize)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 []entity.User
	if rf, ok := ret.Get(0).(func(context.Context, string, uint, ...gormprovider.Option) []entity.User); ok {
		r0 = rf(ctx, paginationToken, pageSize, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entity.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, uint, ...gormprovider.Option) error); ok {
		r1 = rf(ctx, paginationToken, pageSize, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewQuery provides a mock function with given fields: ctx
func (_m *UserRepository) NewQuery(ctx context.Context) *gorm.DB {
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
func (_m *UserRepository) RunInTransaction(ctx context.Context, fn func(context.Context) error) error {
	ret := _m.Called(ctx, fn)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, func(context.Context) error) error); ok {
		r0 = rf(ctx, fn)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateUser provides a mock function with given fields: ctx, id, userUpdates, passwordUpdate
func (_m *UserRepository) UpdateUser(ctx context.Context, id string, userUpdates entity.User, passwordUpdate string) (entity.User, error) {
	ret := _m.Called(ctx, id, userUpdates, passwordUpdate)

	var r0 entity.User
	if rf, ok := ret.Get(0).(func(context.Context, string, entity.User, string) entity.User); ok {
		r0 = rf(ctx, id, userUpdates, passwordUpdate)
	} else {
		r0 = ret.Get(0).(entity.User)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, entity.User, string) error); ok {
		r1 = rf(ctx, id, userUpdates, passwordUpdate)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewUserRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewUserRepository creates a new instance of UserRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewUserRepository(t mockConstructorTestingTNewUserRepository) *UserRepository {
	mock := &UserRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
