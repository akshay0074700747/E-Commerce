// Code generated by mockery v2.36.0. DO NOT EDIT.

package mocks

import (
	context "context"
	helperstructs "ecommerce/web/helpers/helper_structs"

	mock "github.com/stretchr/testify/mock"

	responce "ecommerce/web/helpers/responce"
)

// DiscountUsecaseInterface is an autogenerated mock type for the DiscountUsecaseInterface type
type DiscountUsecaseInterface struct {
	mock.Mock
}

// AddDiscount provides a mock function with given fields: ctx, discountreq
func (_m *DiscountUsecaseInterface) AddDiscount(ctx context.Context, discountreq helperstructs.DiscountReq) (responce.DiscountData, error) {
	ret := _m.Called(ctx, discountreq)

	var r0 responce.DiscountData
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, helperstructs.DiscountReq) (responce.DiscountData, error)); ok {
		return rf(ctx, discountreq)
	}
	if rf, ok := ret.Get(0).(func(context.Context, helperstructs.DiscountReq) responce.DiscountData); ok {
		r0 = rf(ctx, discountreq)
	} else {
		r0 = ret.Get(0).(responce.DiscountData)
	}

	if rf, ok := ret.Get(1).(func(context.Context, helperstructs.DiscountReq) error); ok {
		r1 = rf(ctx, discountreq)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteDiscount provides a mock function with given fields: ctx, id
func (_m *DiscountUsecaseInterface) DeleteDiscount(ctx context.Context, id string) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAllDiscounts provides a mock function with given fields: ctx
func (_m *DiscountUsecaseInterface) GetAllDiscounts(ctx context.Context) ([]responce.DiscountData, error) {
	ret := _m.Called(ctx)

	var r0 []responce.DiscountData
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]responce.DiscountData, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []responce.DiscountData); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]responce.DiscountData)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByID provides a mock function with given fields: ctx, category_id
func (_m *DiscountUsecaseInterface) GetByID(ctx context.Context, category_id string) (responce.DiscountData, error) {
	ret := _m.Called(ctx, category_id)

	var r0 responce.DiscountData
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (responce.DiscountData, error)); ok {
		return rf(ctx, category_id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) responce.DiscountData); ok {
		r0 = rf(ctx, category_id)
	} else {
		r0 = ret.Get(0).(responce.DiscountData)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, category_id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateDiscount provides a mock function with given fields: ctx, discountreq
func (_m *DiscountUsecaseInterface) UpdateDiscount(ctx context.Context, discountreq helperstructs.DiscountReq) (responce.DiscountData, error) {
	ret := _m.Called(ctx, discountreq)

	var r0 responce.DiscountData
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, helperstructs.DiscountReq) (responce.DiscountData, error)); ok {
		return rf(ctx, discountreq)
	}
	if rf, ok := ret.Get(0).(func(context.Context, helperstructs.DiscountReq) responce.DiscountData); ok {
		r0 = rf(ctx, discountreq)
	} else {
		r0 = ret.Get(0).(responce.DiscountData)
	}

	if rf, ok := ret.Get(1).(func(context.Context, helperstructs.DiscountReq) error); ok {
		r1 = rf(ctx, discountreq)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewDiscountUsecaseInterface creates a new instance of DiscountUsecaseInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewDiscountUsecaseInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *DiscountUsecaseInterface {
	mock := &DiscountUsecaseInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
