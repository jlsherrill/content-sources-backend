// Code generated by mockery v2.36.1. DO NOT EDIT.

package dao

import (
	api "github.com/content-services/content-sources-backend/pkg/api"
	mock "github.com/stretchr/testify/mock"

	yum "github.com/content-services/yummy/pkg/yum"
)

// MockEnvironmentDao is an autogenerated mock type for the EnvironmentDao type
type MockEnvironmentDao struct {
	mock.Mock
}

// InsertForRepository provides a mock function with given fields: repoUuid, environments
func (_m *MockEnvironmentDao) InsertForRepository(repoUuid string, environments []yum.Environment) (int64, error) {
	ret := _m.Called(repoUuid, environments)

	var r0 int64
	var r1 error
	if rf, ok := ret.Get(0).(func(string, []yum.Environment) (int64, error)); ok {
		return rf(repoUuid, environments)
	}
	if rf, ok := ret.Get(0).(func(string, []yum.Environment) int64); ok {
		r0 = rf(repoUuid, environments)
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func(string, []yum.Environment) error); ok {
		r1 = rf(repoUuid, environments)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// List provides a mock function with given fields: orgID, uuidRepo, limit, offset, search, sortBy
func (_m *MockEnvironmentDao) List(orgID string, uuidRepo string, limit int, offset int, search string, sortBy string) (api.RepositoryEnvironmentCollectionResponse, int64, error) {
	ret := _m.Called(orgID, uuidRepo, limit, offset, search, sortBy)

	var r0 api.RepositoryEnvironmentCollectionResponse
	var r1 int64
	var r2 error
	if rf, ok := ret.Get(0).(func(string, string, int, int, string, string) (api.RepositoryEnvironmentCollectionResponse, int64, error)); ok {
		return rf(orgID, uuidRepo, limit, offset, search, sortBy)
	}
	if rf, ok := ret.Get(0).(func(string, string, int, int, string, string) api.RepositoryEnvironmentCollectionResponse); ok {
		r0 = rf(orgID, uuidRepo, limit, offset, search, sortBy)
	} else {
		r0 = ret.Get(0).(api.RepositoryEnvironmentCollectionResponse)
	}

	if rf, ok := ret.Get(1).(func(string, string, int, int, string, string) int64); ok {
		r1 = rf(orgID, uuidRepo, limit, offset, search, sortBy)
	} else {
		r1 = ret.Get(1).(int64)
	}

	if rf, ok := ret.Get(2).(func(string, string, int, int, string, string) error); ok {
		r2 = rf(orgID, uuidRepo, limit, offset, search, sortBy)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// OrphanCleanup provides a mock function with given fields:
func (_m *MockEnvironmentDao) OrphanCleanup() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Search provides a mock function with given fields: orgID, request
func (_m *MockEnvironmentDao) Search(orgID string, request api.ContentUnitSearchRequest) ([]api.SearchEnvironmentResponse, error) {
	ret := _m.Called(orgID, request)

	var r0 []api.SearchEnvironmentResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(string, api.ContentUnitSearchRequest) ([]api.SearchEnvironmentResponse, error)); ok {
		return rf(orgID, request)
	}
	if rf, ok := ret.Get(0).(func(string, api.ContentUnitSearchRequest) []api.SearchEnvironmentResponse); ok {
		r0 = rf(orgID, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]api.SearchEnvironmentResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(string, api.ContentUnitSearchRequest) error); ok {
		r1 = rf(orgID, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewMockEnvironmentDao creates a new instance of MockEnvironmentDao. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockEnvironmentDao(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockEnvironmentDao {
	mock := &MockEnvironmentDao{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
