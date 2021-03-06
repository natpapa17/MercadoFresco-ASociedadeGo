// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	domain "github.com/natpapa17/MercadoFresco-ASociedadeGo/internal/purchase_orders/domain"
	mock "github.com/stretchr/testify/mock"
)

// PurchaseOrderService is an autogenerated mock type for the PurchaseOrderService type
type PurchaseOrderService struct {
	mock.Mock
}

// Create provides a mock function with given fields: orderNumber, orderDate, trackingCode, buyerId, productRecordId, orderStatusId
func (_m *PurchaseOrderService) Create(orderNumber string, orderDate string, trackingCode string, buyerId int, productRecordId int, orderStatusId int) (domain.Purchase_Order, error) {
	ret := _m.Called(orderNumber, orderDate, trackingCode, buyerId, productRecordId, orderStatusId)

	var r0 domain.Purchase_Order
	if rf, ok := ret.Get(0).(func(string, string, string, int, int, int) domain.Purchase_Order); ok {
		r0 = rf(orderNumber, orderDate, trackingCode, buyerId, productRecordId, orderStatusId)
	} else {
		r0 = ret.Get(0).(domain.Purchase_Order)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string, string, int, int, int) error); ok {
		r1 = rf(orderNumber, orderDate, trackingCode, buyerId, productRecordId, orderStatusId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewPurchaseOrderService interface {
	mock.TestingT
	Cleanup(func())
}

// NewPurchaseOrderService creates a new instance of PurchaseOrderService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewPurchaseOrderService(t mockConstructorTestingTNewPurchaseOrderService) *PurchaseOrderService {
	mock := &PurchaseOrderService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
