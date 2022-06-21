package mocks

import (
	"github.com/natpapa17/MercadoFresco-ASociedadeGo/internal/buyers"
	buyers "github.com/natpapa17/MercadoFresco-ASociedadeGo/internal/buyers"
	mock "github.com/stretchr/testify/mock"
)

type Service struct {
	mock.Mock
}

func (_m *Service) Create(firstName string, lastName string, address string, documentNumber string) (buyers.Buyer, error) {
	ret := _m.Called(firstName, lastName, address, documentNumber)

	var r0 buyers.Buyer
	if rf, ok := ret.Get(0).(func(string, string, string, string) buyers.Buyer); ok {
		r0 = rf(firstName, lastName, address, documentNumber)
	} else {
		r0 = ret.Get(0).(buyers.Buyer)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string, string, string) error); ok {
		r1 = rf(firstName, lastName, address, documentNumber)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (_m *Service) DeleteById(id int) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(int) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

func (_m *Service) GetAll() ([]buyers.Buyer, error) {
	ret := _m.Called()

	var r0 []buyers.Buyer
	if rf, ok := ret.Get(0).(func() []buyers.Buyer); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]buyers.Buyer)
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
func (_m *Service) GetById(id int) (buyers.Buyer, error) {
	ret := _m.Called(id)

	var r0 buyers.Buyer
	if rf, ok := ret.Get(0).(func(int) buyers.Buyer); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(buyers.Buyer)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (_m *Service) UpdateById(id int, firstName string, lastName string, address string, documentNumber string) (buyers.Buyer, error) {
	ret := _m.Called(id, firstName, lastName, address, documentNumber)

	var r0 buyers.Buyer
	if rf, ok := ret.Get(0).(func(int, string, string, string, string) buyers.Buyer); ok {
		r0 = rf(id, firstName, lastName, address, documentNumber)
	} else {
		r0 = ret.Get(0).(buyers.Buyer)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int, string, string, string, string) error); ok {
		r1 = rf(id, firstName, lastName, address, documentNumber)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewService interface {
	mock.TestingT
	Cleanup(func())
}

func NewService(t mockConstructorTestingTNewService) *Service {
	mock := &Service{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
