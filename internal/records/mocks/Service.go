// Code generated by mockery v2.13.1. DO NOT EDIT.

package mocks

import (
	records "github.com/natpapa17/MercadoFresco-ASociedadeGo/internal/records"
	mock "github.com/stretchr/testify/mock"
)

// Service is an autogenerated mock type for the Service type
type Service struct {
	mock.Mock
}

// Create provides a mock function with given fields: last_update_date, purchase_price, sale_price, product_id
func (_m *Service) Create(last_update_date string, purchase_price int, sale_price int, product_id int) (records.Records, error) {
	ret := _m.Called(last_update_date, purchase_price, sale_price, product_id)

	var r0 records.Records
	if rf, ok := ret.Get(0).(func(string, int, int, int) records.Records); ok {
		r0 = rf(last_update_date, purchase_price, sale_price, product_id)
	} else {
		r0 = ret.Get(0).(records.Records)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, int, int, int) error); ok {
		r1 = rf(last_update_date, purchase_price, sale_price, product_id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAll provides a mock function with given fields:
func (_m *Service) GetAll() ([]records.Records, error) {
	ret := _m.Called()

	var r0 []records.Records
	if rf, ok := ret.Get(0).(func() []records.Records); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]records.Records)
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

// GetById provides a mock function with given fields: id
func (_m *Service) GetById(id int) (records.Records, error) {
	ret := _m.Called(id)

	var r0 records.Records
	if rf, ok := ret.Get(0).(func(int) records.Records); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(records.Records)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewService interface {
	mock.TestingT
	Cleanup(func())
}

// NewService creates a new instance of Service. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewService(t mockConstructorTestingTNewService) *Service {
	mock := &Service{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}