// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	gorm "github.com/jinzhu/gorm"
	mock "github.com/stretchr/testify/mock"
)

// DatabaseManager is an autogenerated mock type for the DatabaseManager type
type DatabaseManager struct {
	mock.Mock
}

// GetDB provides a mock function with given fields:
func (_m *DatabaseManager) GetDB() *gorm.DB {
	ret := _m.Called()

	var r0 *gorm.DB
	if rf, ok := ret.Get(0).(func() *gorm.DB); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gorm.DB)
		}
	}

	return r0
}

// Initialize provides a mock function with given fields: dsn, connection
func (_m *DatabaseManager) Initialize(dsn string, connection string) error {
	ret := _m.Called(dsn, connection)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(dsn, connection)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewDatabaseManager interface {
	mock.TestingT
	Cleanup(func())
}

// NewDatabaseManager creates a new instance of DatabaseManager. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewDatabaseManager(t mockConstructorTestingTNewDatabaseManager) *DatabaseManager {
	mock := &DatabaseManager{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
