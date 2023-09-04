// Code generated by mockery v2.33.0. DO NOT EDIT.

package dao

import (
	api "github.com/content-services/content-sources-backend/pkg/api"
	mock "github.com/stretchr/testify/mock"
)

// MockTaskInfoDao is an autogenerated mock type for the TaskInfoDao type
type MockTaskInfoDao struct {
	mock.Mock
}

// Fetch provides a mock function with given fields: OrgID, id
func (_m *MockTaskInfoDao) Fetch(OrgID string, id string) (api.TaskInfoResponse, error) {
	ret := _m.Called(OrgID, id)

	var r0 api.TaskInfoResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string) (api.TaskInfoResponse, error)); ok {
		return rf(OrgID, id)
	}
	if rf, ok := ret.Get(0).(func(string, string) api.TaskInfoResponse); ok {
		r0 = rf(OrgID, id)
	} else {
		r0 = ret.Get(0).(api.TaskInfoResponse)
	}

	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(OrgID, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IsSnapshotInProgress provides a mock function with given fields: orgID, repoUUID
func (_m *MockTaskInfoDao) IsSnapshotInProgress(orgID string, repoUUID string) (bool, error) {
	ret := _m.Called(orgID, repoUUID)

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string) (bool, error)); ok {
		return rf(orgID, repoUUID)
	}
	if rf, ok := ret.Get(0).(func(string, string) bool); ok {
		r0 = rf(orgID, repoUUID)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(orgID, repoUUID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// List provides a mock function with given fields: OrgID, pageData, filterData
func (_m *MockTaskInfoDao) List(OrgID string, pageData api.PaginationData, filterData api.TaskInfoFilterData) (api.TaskInfoCollectionResponse, int64, error) {
	ret := _m.Called(OrgID, pageData, filterData)

	var r0 api.TaskInfoCollectionResponse
	var r1 int64
	var r2 error
	if rf, ok := ret.Get(0).(func(string, api.PaginationData, api.TaskInfoFilterData) (api.TaskInfoCollectionResponse, int64, error)); ok {
		return rf(OrgID, pageData, filterData)
	}
	if rf, ok := ret.Get(0).(func(string, api.PaginationData, api.TaskInfoFilterData) api.TaskInfoCollectionResponse); ok {
		r0 = rf(OrgID, pageData, filterData)
	} else {
		r0 = ret.Get(0).(api.TaskInfoCollectionResponse)
	}

	if rf, ok := ret.Get(1).(func(string, api.PaginationData, api.TaskInfoFilterData) int64); ok {
		r1 = rf(OrgID, pageData, filterData)
	} else {
		r1 = ret.Get(1).(int64)
	}

	if rf, ok := ret.Get(2).(func(string, api.PaginationData, api.TaskInfoFilterData) error); ok {
		r2 = rf(OrgID, pageData, filterData)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// NewMockTaskInfoDao creates a new instance of MockTaskInfoDao. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockTaskInfoDao(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockTaskInfoDao {
	mock := &MockTaskInfoDao{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
