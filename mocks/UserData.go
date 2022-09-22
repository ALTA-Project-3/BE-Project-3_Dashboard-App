// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	user "project/dashboard/features/user"

	mock "github.com/stretchr/testify/mock"
)

// UserData is an autogenerated mock type for the DataInterface type
type UserData struct {
	mock.Mock
}

// DelData provides a mock function with given fields: id
func (_m *UserData) DelData(id int) int {
	ret := _m.Called(id)

	var r0 int
	if rf, ok := ret.Get(0).(func(int) int); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// InsertData provides a mock function with given fields: data
func (_m *UserData) InsertData(data user.Core) int {
	ret := _m.Called(data)

	var r0 int
	if rf, ok := ret.Get(0).(func(user.Core) int); ok {
		r0 = rf(data)
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// SelectAll provides a mock function with given fields: page, token
func (_m *UserData) SelectAll(page int, token int) ([]user.Core, error) {
	ret := _m.Called(page, token)

	var r0 []user.Core
	if rf, ok := ret.Get(0).(func(int, int) []user.Core); ok {
		r0 = rf(page, token)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]user.Core)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int, int) error); ok {
		r1 = rf(page, token)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SelectProfile provides a mock function with given fields: id
func (_m *UserData) SelectProfile(id int) (user.Core, user.DashBoard, error) {
	ret := _m.Called(id)

	var r0 user.Core
	if rf, ok := ret.Get(0).(func(int) user.Core); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(user.Core)
	}

	var r1 user.DashBoard
	if rf, ok := ret.Get(1).(func(int) user.DashBoard); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Get(1).(user.DashBoard)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(int) error); ok {
		r2 = rf(id)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// UpdateData provides a mock function with given fields: data
func (_m *UserData) UpdateData(data user.Core) int {
	ret := _m.Called(data)

	var r0 int
	if rf, ok := ret.Get(0).(func(user.Core) int); ok {
		r0 = rf(data)
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

type mockConstructorTestingTNewUserData interface {
	mock.TestingT
	Cleanup(func())
}

// NewUserData creates a new instance of UserData. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewUserData(t mockConstructorTestingTNewUserData) *UserData {
	mock := &UserData{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}