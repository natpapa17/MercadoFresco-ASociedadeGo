package mocks

import (
	buyers "github.com/natpapa17/MercadoFresco-ASociedadeGo/internal/buyers"
	mock "github.com/stretchr/testify/mock"
)

type Repository struct {
	mock.Mock
}

func (_m *Repository) Create(firstName string, lastName string, address string, documentNumber string) (buyers.Buyer, error) {
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

func (_m *Repository) DeleteBuyerById(id int) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(int) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

func (_m *Repository) GetAll() ([]buyers.Buyer, error) {
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

func (_m *Repository) GetBuyerById(id int) (buyers.Buyer, error) {
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

func (_m *Repository) GetByBuyer(code string) (buyers.Buyer, error) {
	ret := _m.Called(code)

	var r0 buyers.Buyer
	if rf, ok := ret.Get(0).(func(string) buyers.Buyer); ok {
		r0 = rf(code)
	} else {
		r0 = ret.Get(0).(buyers.Buyer)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(code)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (_m *Repository) UpdateBuyerById(id int, firstName string, lastName string, address string, documentNumber string) (buyers.Buyer, error) {
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

type mockConstructorTestingTNewRepository interface {
	mock.TestingT
	Cleanup(func())
}

func NewRepository(t mockConstructorTestingTNewRepository) *Repository {
	mock := &Repository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
