// Code generated by mockery v2.13.1. DO NOT EDIT.

package mocks

import (
	domain "github.com/natpapa17/MercadoFresco-ASociedadeGo/internal/localities/domain"
	mock "github.com/stretchr/testify/mock"

	repository "github.com/natpapa17/MercadoFresco-ASociedadeGo/internal/localities/repository"
)

// LocalityRepository is an autogenerated mock type for the LocalityRepository type
type LocalityRepository struct {
	mock.Mock
}

// Create provides a mock function with given fields: name, province_id
func (_m *LocalityRepository) Create(name string, province_id int) (domain.Locality, error) {
	ret := _m.Called(name, province_id)

	var r0 domain.Locality
	if rf, ok := ret.Get(0).(func(string, int) domain.Locality); ok {
		r0 = rf(name, province_id)
	} else {
		r0 = ret.Get(0).(domain.Locality)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, int) error); ok {
		r1 = rf(name, province_id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAll provides a mock function with given fields:
func (_m *LocalityRepository) GetAll() ([]domain.Locality, error) {
	ret := _m.Called()

	var r0 []domain.Locality
	if rf, ok := ret.Get(0).(func() []domain.Locality); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.Locality)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ReportAll provides a mock function with given fields:
func (_m *LocalityRepository) ReportAll() ([]repository.LocalityReport, error) {
	ret := _m.Called()

	var r0 []repository.LocalityReport
	if rf, ok := ret.Get(0).(func() []repository.LocalityReport); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]repository.LocalityReport)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ReportById provides a mock function with given fields: id
func (_m *LocalityRepository) ReportById(id int) (repository.LocalityReport, error) {
	ret := _m.Called(id)

	var r0 repository.LocalityReport
	if rf, ok := ret.Get(0).(func(int) repository.LocalityReport); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(repository.LocalityReport)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewLocalityRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewLocalityRepository creates a new instance of LocalityRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewLocalityRepository(t mockConstructorTestingTNewLocalityRepository) *LocalityRepository {
	mock := &LocalityRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
