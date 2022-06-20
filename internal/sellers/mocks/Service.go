package mocks

import (
	"github.com/natpapa17/MercadoFresco-ASociedadeGo/internal/sellers"
	"github.com/stretchr/testify/mock"
)

type Service struct {
	mock.Mock
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
func (_m *Service) GetById(id int) (sellers.Seller, error) {
	ret := _m.Called(id)

	var r0 sellers.Seller
	if rf, ok := ret.Get(0).(func(int) sellers.Seller); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(sellers.Seller)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (_m *Service) GetAll() ([]sellers.Seller, error) {
	ret := _m.Called()

	var r0 []sellers.Seller
	if rf, ok := ret.Get(0).(func() []sellers.Seller); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]sellers.Seller)
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