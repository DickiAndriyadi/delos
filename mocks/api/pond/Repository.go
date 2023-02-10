// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	filter "delos/api/pond/filter"

	gorm "github.com/jinzhu/gorm"

	mock "github.com/stretchr/testify/mock"

	model "delos/api/model"

	payload "delos/api/model/payload"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// AddNewPond provides a mock function with given fields: db, store
func (_m *Repository) AddNewPond(db *gorm.DB, store *model.Pond) (*model.Pond, error) {
	ret := _m.Called(db, store)

	var r0 *model.Pond
	if rf, ok := ret.Get(0).(func(*gorm.DB, *model.Pond) *model.Pond); ok {
		r0 = rf(db, store)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Pond)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*gorm.DB, *model.Pond) error); ok {
		r1 = rf(db, store)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeletePond provides a mock function with given fields: db, deletePondMap
func (_m *Repository) DeletePond(db *gorm.DB, deletePondMap map[string]interface{}) error {
	ret := _m.Called(db, deletePondMap)

	var r0 error
	if rf, ok := ret.Get(0).(func(*gorm.DB, map[string]interface{}) error); ok {
		r0 = rf(db, deletePondMap)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DetailOfPond provides a mock function with given fields: db, _a1
func (_m *Repository) DetailOfPond(db *gorm.DB, _a1 map[string]interface{}) (*model.Pond, error) {
	ret := _m.Called(db, _a1)

	var r0 *model.Pond
	if rf, ok := ret.Get(0).(func(*gorm.DB, map[string]interface{}) *model.Pond); ok {
		r0 = rf(db, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Pond)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*gorm.DB, map[string]interface{}) error); ok {
		r1 = rf(db, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListOfPond provides a mock function with given fields: db, f
func (_m *Repository) ListOfPond(db *gorm.DB, f *filter.FilterListOfPond) (*model.Ponds, error) {
	ret := _m.Called(db, f)

	var r0 *model.Ponds
	if rf, ok := ret.Get(0).(func(*gorm.DB, *filter.FilterListOfPond) *model.Ponds); ok {
		r0 = rf(db, f)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Ponds)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*gorm.DB, *filter.FilterListOfPond) error); ok {
		r1 = rf(db, f)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdatePond provides a mock function with given fields: db, store, id
func (_m *Repository) UpdatePond(db *gorm.DB, store *payload.AddUpdatePond, id string) (*model.Pond, error) {
	ret := _m.Called(db, store, id)

	var r0 *model.Pond
	if rf, ok := ret.Get(0).(func(*gorm.DB, *payload.AddUpdatePond, string) *model.Pond); ok {
		r0 = rf(db, store, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Pond)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*gorm.DB, *payload.AddUpdatePond, string) error); ok {
		r1 = rf(db, store, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewRepository creates a new instance of Repository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewRepository(t mockConstructorTestingTNewRepository) *Repository {
	mock := &Repository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
