// Code generated by mockery v2.13.1. DO NOT EDIT.

package mocks

import (
	context "context"

	product_batch "github.com/natpapa17/MercadoFresco-ASociedadeGo/internal/product_batch"
	mock "github.com/stretchr/testify/mock"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// Add provides a mock function with given fields: ctx, id, batchNumber, currentQuantity, currentTemperature, dueDate, initialQuantity, manufacturingDate, manufacturingHour, minimumTemperature, productID, sectionID
func (_m *Repository) Add(ctx context.Context, id int, batchNumber int, currentQuantity int, currentTemperature int, dueDate string, initialQuantity int, manufacturingDate string, manufacturingHour int, minimumTemperature int, productID int, sectionID int) (product_batch.ProductBatch, error) {
	ret := _m.Called(ctx, id, batchNumber, currentQuantity, currentTemperature, dueDate, initialQuantity, manufacturingDate, manufacturingHour, minimumTemperature, productID, sectionID)

	var r0 product_batch.ProductBatch
	if rf, ok := ret.Get(0).(func(context.Context, int, int, int, int, string, int, string, int, int, int, int) product_batch.ProductBatch); ok {
		r0 = rf(ctx, id, batchNumber, currentQuantity, currentTemperature, dueDate, initialQuantity, manufacturingDate, manufacturingHour, minimumTemperature, productID, sectionID)
	} else {
		r0 = ret.Get(0).(product_batch.ProductBatch)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int, int, int, int, string, int, string, int, int, int, int) error); ok {
		r1 = rf(ctx, id, batchNumber, currentQuantity, currentTemperature, dueDate, initialQuantity, manufacturingDate, manufacturingHour, minimumTemperature, productID, sectionID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetById provides a mock function with given fields: ctx, id
func (_m *Repository) GetById(ctx context.Context, id int) ([]product_batch.ProductsReport, error) {
	ret := _m.Called(ctx, id)

	var r0 []product_batch.ProductsReport
	if rf, ok := ret.Get(0).(func(context.Context, int) []product_batch.ProductsReport); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]product_batch.ProductsReport)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// HasBatchNumber provides a mock function with given fields: ctx, number
func (_m *Repository) HasBatchNumber(ctx context.Context, number int) (bool, error) {
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