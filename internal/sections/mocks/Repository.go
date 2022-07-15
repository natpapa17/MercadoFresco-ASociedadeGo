// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	context "context"

	sections "github.com/natpapa17/MercadoFresco-ASociedadeGo/internal/sections"
	mock "github.com/stretchr/testify/mock"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// Add provides a mock function with given fields: ctx, id, sectionNumber, currentTemperature, minimumTemprarature, currentCapacity, minimumCapacity, maximumCapacity, warehouseID, productTypeID
func (_m *Repository) Add(ctx context.Context, id int, sectionNumber int, currentTemperature float32, minimumTemprarature float32, currentCapacity int, minimumCapacity int, maximumCapacity int, warehouseID int, productTypeID int) (sections.Section, error) {
	ret := _m.Called(ctx, id, sectionNumber, currentTemperature, minimumTemprarature, currentCapacity, minimumCapacity, maximumCapacity, warehouseID, productTypeID)

	var r0 sections.Section
	if rf, ok := ret.Get(0).(func(context.Context, int, int, float32, float32, int, int, int, int, int) sections.Section); ok {
		r0 = rf(ctx, id, sectionNumber, currentTemperature, minimumTemprarature, currentCapacity, minimumCapacity, maximumCapacity, warehouseID, productTypeID)
	} else {
		r0 = ret.Get(0).(sections.Section)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int, int, float32, float32, int, int, int, int, int) error); ok {
		r1 = rf(ctx, id, sectionNumber, currentTemperature, minimumTemprarature, currentCapacity, minimumCapacity, maximumCapacity, warehouseID, productTypeID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: ctx, id
func (_m *Repository) Delete(ctx context.Context, id int) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAll provides a mock function with given fields: ctx
func (_m *Repository) GetAll(ctx context.Context) ([]sections.Section, error) {
	ret := _m.Called(ctx)

	var r0 []sections.Section
	if rf, ok := ret.Get(0).(func(context.Context) []sections.Section); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]sections.Section)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetById provides a mock function with given fields: ctx, id
func (_m *Repository) GetById(ctx context.Context, id int) (sections.Section, error) {
	ret := _m.Called(ctx, id)

	var r0 sections.Section
	if rf, ok := ret.Get(0).(func(context.Context, int) sections.Section); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(sections.Section)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// HasSectionNumber provides a mock function with given fields: ctx, number
func (_m *Repository) HasSectionNumber(ctx context.Context, number int) (bool, error) {
	ret := _m.Called(ctx, number)

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context, int) bool); ok {
		r0 = rf(ctx, number)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int) error); ok {
		r1 = rf(ctx, number)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// LastID provides a mock function with given fields: ctx
func (_m *Repository) LastID(ctx context.Context) (int, error) {
	ret := _m.Called(ctx)

	var r0 int
	if rf, ok := ret.Get(0).(func(context.Context) int); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateById provides a mock function with given fields: ctx, id, section
func (_m *Repository) UpdateById(ctx context.Context, id int, section sections.Section) (sections.Section, error) {
	ret := _m.Called(ctx, id, section)

	var r0 sections.Section
	if rf, ok := ret.Get(0).(func(context.Context, int, sections.Section) sections.Section); ok {
		r0 = rf(ctx, id, section)
	} else {
		r0 = ret.Get(0).(sections.Section)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int, sections.Section) error); ok {
		r1 = rf(ctx, id, section)
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
