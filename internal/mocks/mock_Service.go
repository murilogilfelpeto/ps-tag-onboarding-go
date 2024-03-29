// Code generated by mockery v2.40.1. DO NOT EDIT.

package mocks

import (
	context "context"

	models "github.com/murilogilfelpeto/ps-tag-onboarding-go/internal/service/models"
	mock "github.com/stretchr/testify/mock"
)

// Service is an autogenerated mock type for the Service type
type Service struct {
	mock.Mock
}

type Service_Expecter struct {
	mock *mock.Mock
}

func (_m *Service) EXPECT() *Service_Expecter {
	return &Service_Expecter{mock: &_m.Mock}
}

// GetUserById provides a mock function with given fields: ctx, id
func (_m *Service) GetUserById(ctx context.Context, id string) (*models.User, error) {
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

// Service_GetUserById_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetUserById'
type Service_GetUserById_Call struct {
	*mock.Call
}

// GetUserById is a helper method to define mock.On call
//   - ctx context.Context
//   - id string
func (_e *Service_Expecter) GetUserById(ctx interface{}, id interface{}) *Service_GetUserById_Call {
	return &Service_GetUserById_Call{Call: _e.mock.On("GetUserById", ctx, id)}
}

func (_c *Service_GetUserById_Call) Run(run func(ctx context.Context, id string)) *Service_GetUserById_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *Service_GetUserById_Call) Return(_a0 *models.User, _a1 error) *Service_GetUserById_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Service_GetUserById_Call) RunAndReturn(run func(context.Context, string) (*models.User, error)) *Service_GetUserById_Call {
	_c.Call.Return(run)
	return _c
}

// SaveUser provides a mock function with given fields: ctx, user
func (_m *Service) SaveUser(ctx context.Context, user models.User) (*models.User, error) {
	ret := _m.Called(ctx, user)

	if len(ret) == 0 {
		panic("no return value specified for SaveUser")
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

// Service_SaveUser_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SaveUser'
type Service_SaveUser_Call struct {
	*mock.Call
}

// SaveUser is a helper method to define mock.On call
//   - ctx context.Context
//   - user models.User
func (_e *Service_Expecter) SaveUser(ctx interface{}, user interface{}) *Service_SaveUser_Call {
	return &Service_SaveUser_Call{Call: _e.mock.On("SaveUser", ctx, user)}
}

func (_c *Service_SaveUser_Call) Run(run func(ctx context.Context, user models.User)) *Service_SaveUser_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(models.User))
	})
	return _c
}

func (_c *Service_SaveUser_Call) Return(_a0 *models.User, _a1 error) *Service_SaveUser_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Service_SaveUser_Call) RunAndReturn(run func(context.Context, models.User) (*models.User, error)) *Service_SaveUser_Call {
	_c.Call.Return(run)
	return _c
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
