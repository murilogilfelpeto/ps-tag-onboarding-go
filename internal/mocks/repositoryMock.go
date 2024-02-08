// Code generated by mockery v2.40.1. DO NOT EDIT.

package mocks

import (
	context "context"

	models "github.com/murilogilfelpeto/ps-tag-onboarding-go/internal/service/models"
	mock "github.com/stretchr/testify/mock"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// GetUserByFullName provides a mock function with given fields: ctx, firstName, lastName
func (_m *Repository) GetUserByFullName(ctx context.Context, firstName string, lastName string) (*models.User, error) {
	ret := _m.Called(ctx, firstName, lastName)

	if len(ret) == 0 {
		panic("no return value specified for GetUserByFullName")
	}

	var r0 *models.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) (*models.User, error)); ok {
		return rf(ctx, firstName, lastName)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string) *models.User); ok {
		r0 = rf(ctx, firstName, lastName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.User)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, firstName, lastName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUserById provides a mock function with given fields: ctx, id
func (_m *Repository) GetUserById(ctx context.Context, id string) (*models.User, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for GetUserById")
	}

	var r0 *models.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*models.User, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *models.User); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.User)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Save provides a mock function with given fields: ctx, user
func (_m *Repository) Save(ctx context.Context, user models.User) (*models.User, error) {
	ret := _m.Called(ctx, user)

	if len(ret) == 0 {
		panic("no return value specified for Save")
	}

	var r0 *models.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, models.User) (*models.User, error)); ok {
		return rf(ctx, user)
	}
	if rf, ok := ret.Get(0).(func(context.Context, models.User) *models.User); ok {
		r0 = rf(ctx, user)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.User)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, models.User) error); ok {
		r1 = rf(ctx, user)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewRepository creates a new instance of Repository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *Repository {
	mock := &Repository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
