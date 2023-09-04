// Code generated by mockery v2.33.0. DO NOT EDIT.

package dao

import mock "github.com/stretchr/testify/mock"

// MockDomainDao is an autogenerated mock type for the DomainDao type
type MockDomainDao struct {
	mock.Mock
}

// FetchOrCreateDomain provides a mock function with given fields: orgId
func (_m *MockDomainDao) FetchOrCreateDomain(orgId string) (string, error) {
	ret := _m.Called(orgId)

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (string, error)); ok {
		return rf(orgId)
	}
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(orgId)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(orgId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewMockDomainDao creates a new instance of MockDomainDao. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockDomainDao(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockDomainDao {
	mock := &MockDomainDao{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
