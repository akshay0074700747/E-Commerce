// Code generated by mockery v2.36.0. DO NOT EDIT.

package mocks

import (
	helperstructs "ecommerce/web/helpers/helper_structs"

	mock "github.com/stretchr/testify/mock"

	responce "ecommerce/web/helpers/responce"
)

// SuAdminRepo is an autogenerated mock type for the SuAdminRepo type
type SuAdminRepo struct {
	mock.Mock
}

// BlockUser provides a mock function with given fields: blockreq
func (_m *SuAdminRepo) BlockUser(blockreq helperstructs.BlockReq) error {
	ret := _m.Called(blockreq)

	var r0 error
	if rf, ok := ret.Get(0).(func(helperstructs.BlockReq) error); ok {
		r0 = rf(blockreq)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateAdmin provides a mock function with given fields: admin
func (_m *SuAdminRepo) CreateAdmin(admin helperstructs.AdminReq) (responce.AdminData, error) {
	ret := _m.Called(admin)

	var r0 responce.AdminData
	var r1 error
	if rf, ok := ret.Get(0).(func(helperstructs.AdminReq) (responce.AdminData, error)); ok {
		return rf(admin)
	}
	if rf, ok := ret.Get(0).(func(helperstructs.AdminReq) responce.AdminData); ok {
		r0 = rf(admin)
	} else {
		r0 = ret.Get(0).(responce.AdminData)
	}

	if rf, ok := ret.Get(1).(func(helperstructs.AdminReq) error); ok {
		r1 = rf(admin)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAllAdmins provides a mock function with given fields:
func (_m *SuAdminRepo) GetAllAdmins() ([]responce.AdminData, error) {
	ret := _m.Called()

	var r0 []responce.AdminData
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]responce.AdminData, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []responce.AdminData); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]responce.AdminData)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAllUsers provides a mock function with given fields:
func (_m *SuAdminRepo) GetAllUsers() ([]responce.AdminsideUsersData, error) {
	ret := _m.Called()

	var r0 []responce.AdminsideUsersData
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]responce.AdminsideUsersData, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []responce.AdminsideUsersData); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]responce.AdminsideUsersData)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByEmail provides a mock function with given fields: suadmin
func (_m *SuAdminRepo) GetByEmail(suadmin helperstructs.SuAdminReq) (responce.SuAdminData, error) {
	ret := _m.Called(suadmin)

	var r0 responce.SuAdminData
	var r1 error
	if rf, ok := ret.Get(0).(func(helperstructs.SuAdminReq) (responce.SuAdminData, error)); ok {
		return rf(suadmin)
	}
	if rf, ok := ret.Get(0).(func(helperstructs.SuAdminReq) responce.SuAdminData); ok {
		r0 = rf(suadmin)
	} else {
		r0 = ret.Get(0).(responce.SuAdminData)
	}

	if rf, ok := ret.Get(1).(func(helperstructs.SuAdminReq) error); ok {
		r1 = rf(suadmin)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDetailedReport provides a mock function with given fields: email
func (_m *SuAdminRepo) GetDetailedReport(email string) (responce.DetailReportResponce, error) {
	ret := _m.Called(email)

	var r0 responce.DetailReportResponce
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (responce.DetailReportResponce, error)); ok {
		return rf(email)
	}
	if rf, ok := ret.Get(0).(func(string) responce.DetailReportResponce); ok {
		r0 = rf(email)
	} else {
		r0 = ret.Get(0).(responce.DetailReportResponce)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetReportes provides a mock function with given fields:
func (_m *SuAdminRepo) GetReportes() ([]responce.DetailReportResponce, error) {
	ret := _m.Called()

	var r0 []responce.DetailReportResponce
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]responce.DetailReportResponce, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []responce.DetailReportResponce); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]responce.DetailReportResponce)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetReports provides a mock function with given fields: email
func (_m *SuAdminRepo) GetReports(email string) (int, error) {
	ret := _m.Called(email)

	var r0 int
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (int, error)); ok {
		return rf(email)
	}
	if rf, ok := ret.Get(0).(func(string) int); ok {
		r0 = rf(email)
	} else {
		r0 = ret.Get(0).(int)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewSuAdminRepo creates a new instance of SuAdminRepo. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewSuAdminRepo(t interface {
	mock.TestingT
	Cleanup(func())
}) *SuAdminRepo {
	mock := &SuAdminRepo{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
