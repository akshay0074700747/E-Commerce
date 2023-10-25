// Code generated by mockery v2.36.0. DO NOT EDIT.

package mocks

import (
	context "context"
	helperstructs "ecommerce/web/helpers/helper_structs"

	mock "github.com/stretchr/testify/mock"

	responce "ecommerce/web/helpers/responce"
)

// ProductUsecaseInterface is an autogenerated mock type for the ProductUsecaseInterface type
type ProductUsecaseInterface struct {
	mock.Mock
}

// AddProduct provides a mock function with given fields: ctx, productreq
func (_m *ProductUsecaseInterface) AddProduct(ctx context.Context, productreq helperstructs.ProductReq) (responce.ProuctData, error) {
	ret := _m.Called(ctx, productreq)

	var r0 responce.ProuctData
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, helperstructs.ProductReq) (responce.ProuctData, error)); ok {
		return rf(ctx, productreq)
	}
	if rf, ok := ret.Get(0).(func(context.Context, helperstructs.ProductReq) responce.ProuctData); ok {
		r0 = rf(ctx, productreq)
	} else {
		r0 = ret.Get(0).(responce.ProuctData)
	}

	if rf, ok := ret.Get(1).(func(context.Context, helperstructs.ProductReq) error); ok {
		r1 = rf(ctx, productreq)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteProduct provides a mock function with given fields: ctx, product_id
func (_m *ProductUsecaseInterface) DeleteProduct(ctx context.Context, product_id uint) error {
	ret := _m.Called(ctx, product_id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uint) error); ok {
		r0 = rf(ctx, product_id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetProducts provides a mock function with given fields: ctx
func (_m *ProductUsecaseInterface) GetProducts(ctx context.Context) ([]responce.ProuctData, error) {
	ret := _m.Called(ctx)

	var r0 []responce.ProuctData
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]responce.ProuctData, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []responce.ProuctData); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]responce.ProuctData)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateProduct provides a mock function with given fields: ctx, productreq
func (_m *ProductUsecaseInterface) UpdateProduct(ctx context.Context, productreq helperstructs.ProductReq) (responce.ProuctData, error) {
	ret := _m.Called(ctx, productreq)

	var r0 responce.ProuctData
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, helperstructs.ProductReq) (responce.ProuctData, error)); ok {
		return rf(ctx, productreq)
	}
	if rf, ok := ret.Get(0).(func(context.Context, helperstructs.ProductReq) responce.ProuctData); ok {
		r0 = rf(ctx, productreq)
	} else {
		r0 = ret.Get(0).(responce.ProuctData)
	}

	if rf, ok := ret.Get(1).(func(context.Context, helperstructs.ProductReq) error); ok {
		r1 = rf(ctx, productreq)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewProductUsecaseInterface creates a new instance of ProductUsecaseInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewProductUsecaseInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *ProductUsecaseInterface {
	mock := &ProductUsecaseInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
