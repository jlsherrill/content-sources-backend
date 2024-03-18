// Code generated by mockery v2.36.1. DO NOT EDIT.

package pulp_client

import (
	context "context"

	zest "github.com/content-services/zest/release/v2024"
	mock "github.com/stretchr/testify/mock"
)

// MockPulpClient is an autogenerated mock type for the PulpClient type
type MockPulpClient struct {
	mock.Mock
}

// CancelTask provides a mock function with given fields: ctx, taskHref
func (_m *MockPulpClient) CancelTask(ctx context.Context, taskHref string) (zest.TaskResponse, error) {
	ret := _m.Called(ctx, taskHref)

	var r0 zest.TaskResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (zest.TaskResponse, error)); ok {
		return rf(ctx, taskHref)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) zest.TaskResponse); ok {
		r0 = rf(ctx, taskHref)
	} else {
		r0 = ret.Get(0).(zest.TaskResponse)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, taskHref)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateOrUpdateGuardsForOrg provides a mock function with given fields: ctx, orgId
func (_m *MockPulpClient) CreateOrUpdateGuardsForOrg(ctx context.Context, orgId string) (string, error) {
	ret := _m.Called(ctx, orgId)

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (string, error)); ok {
		return rf(ctx, orgId)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) string); ok {
		r0 = rf(ctx, orgId)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, orgId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateRpmDistribution provides a mock function with given fields: ctx, publicationHref, name, basePath, contentGuardHref
func (_m *MockPulpClient) CreateRpmDistribution(ctx context.Context, publicationHref string, name string, basePath string, contentGuardHref *string) (*string, error) {
	ret := _m.Called(ctx, publicationHref, name, basePath, contentGuardHref)

	var r0 *string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string, *string) (*string, error)); ok {
		return rf(ctx, publicationHref, name, basePath, contentGuardHref)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string, *string) *string); ok {
		r0 = rf(ctx, publicationHref, name, basePath, contentGuardHref)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*string)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string, string, *string) error); ok {
		r1 = rf(ctx, publicationHref, name, basePath, contentGuardHref)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateRpmPublication provides a mock function with given fields: ctx, versionHref
func (_m *MockPulpClient) CreateRpmPublication(ctx context.Context, versionHref string) (*string, error) {
	ret := _m.Called(ctx, versionHref)

	var r0 *string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*string, error)); ok {
		return rf(ctx, versionHref)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *string); ok {
		r0 = rf(ctx, versionHref)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*string)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, versionHref)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateRpmRemote provides a mock function with given fields: ctx, name, url, clientCert, clientKey, caCert
func (_m *MockPulpClient) CreateRpmRemote(ctx context.Context, name string, url string, clientCert *string, clientKey *string, caCert *string) (*zest.RpmRpmRemoteResponse, error) {
	ret := _m.Called(ctx, name, url, clientCert, clientKey, caCert)

	var r0 *zest.RpmRpmRemoteResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, *string, *string, *string) (*zest.RpmRpmRemoteResponse, error)); ok {
		return rf(ctx, name, url, clientCert, clientKey, caCert)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string, *string, *string, *string) *zest.RpmRpmRemoteResponse); ok {
		r0 = rf(ctx, name, url, clientCert, clientKey, caCert)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*zest.RpmRpmRemoteResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string, *string, *string, *string) error); ok {
		r1 = rf(ctx, name, url, clientCert, clientKey, caCert)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateRpmRepository provides a mock function with given fields: ctx, uuid, rpmRemotePulpRef
func (_m *MockPulpClient) CreateRpmRepository(ctx context.Context, uuid string, rpmRemotePulpRef *string) (*zest.RpmRpmRepositoryResponse, error) {
	ret := _m.Called(ctx, uuid, rpmRemotePulpRef)

	var r0 *zest.RpmRpmRepositoryResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, *string) (*zest.RpmRpmRepositoryResponse, error)); ok {
		return rf(ctx, uuid, rpmRemotePulpRef)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, *string) *zest.RpmRpmRepositoryResponse); ok {
		r0 = rf(ctx, uuid, rpmRemotePulpRef)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*zest.RpmRpmRepositoryResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, *string) error); ok {
		r1 = rf(ctx, uuid, rpmRemotePulpRef)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteRpmDistribution provides a mock function with given fields: ctx, rpmDistributionHref
func (_m *MockPulpClient) DeleteRpmDistribution(ctx context.Context, rpmDistributionHref string) (string, error) {
	ret := _m.Called(ctx, rpmDistributionHref)

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (string, error)); ok {
		return rf(ctx, rpmDistributionHref)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) string); ok {
		r0 = rf(ctx, rpmDistributionHref)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, rpmDistributionHref)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteRpmRemote provides a mock function with given fields: ctx, pulpHref
func (_m *MockPulpClient) DeleteRpmRemote(ctx context.Context, pulpHref string) (string, error) {
	ret := _m.Called(ctx, pulpHref)

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (string, error)); ok {
		return rf(ctx, pulpHref)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) string); ok {
		r0 = rf(ctx, pulpHref)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, pulpHref)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteRpmRepository provides a mock function with given fields: ctx, rpmRepositoryHref
func (_m *MockPulpClient) DeleteRpmRepository(ctx context.Context, rpmRepositoryHref string) (string, error) {
	ret := _m.Called(ctx, rpmRepositoryHref)

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (string, error)); ok {
		return rf(ctx, rpmRepositoryHref)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) string); ok {
		r0 = rf(ctx, rpmRepositoryHref)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, rpmRepositoryHref)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteRpmRepositoryVersion provides a mock function with given fields: ctx, href
func (_m *MockPulpClient) DeleteRpmRepositoryVersion(ctx context.Context, href string) (string, error) {
	ret := _m.Called(ctx, href)

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (string, error)); ok {
		return rf(ctx, href)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) string); ok {
		r0 = rf(ctx, href)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, href)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindDistributionByPath provides a mock function with given fields: ctx, path
func (_m *MockPulpClient) FindDistributionByPath(ctx context.Context, path string) (*zest.RpmRpmDistributionResponse, error) {
	ret := _m.Called(ctx, path)

	var r0 *zest.RpmRpmDistributionResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*zest.RpmRpmDistributionResponse, error)); ok {
		return rf(ctx, path)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *zest.RpmRpmDistributionResponse); ok {
		r0 = rf(ctx, path)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*zest.RpmRpmDistributionResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, path)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindRpmPublicationByVersion provides a mock function with given fields: ctx, versionHref
func (_m *MockPulpClient) FindRpmPublicationByVersion(ctx context.Context, versionHref string) (*zest.RpmRpmPublicationResponse, error) {
	ret := _m.Called(ctx, versionHref)

	var r0 *zest.RpmRpmPublicationResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*zest.RpmRpmPublicationResponse, error)); ok {
		return rf(ctx, versionHref)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *zest.RpmRpmPublicationResponse); ok {
		r0 = rf(ctx, versionHref)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*zest.RpmRpmPublicationResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, versionHref)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetContentPath provides a mock function with given fields: ctx
func (_m *MockPulpClient) GetContentPath(ctx context.Context) (string, error) {
	ret := _m.Called(ctx)

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (string, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) string); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetRpmRemoteByName provides a mock function with given fields: ctx, name
func (_m *MockPulpClient) GetRpmRemoteByName(ctx context.Context, name string) (*zest.RpmRpmRemoteResponse, error) {
	ret := _m.Called(ctx, name)

	var r0 *zest.RpmRpmRemoteResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*zest.RpmRpmRemoteResponse, error)); ok {
		return rf(ctx, name)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *zest.RpmRpmRemoteResponse); ok {
		r0 = rf(ctx, name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*zest.RpmRpmRemoteResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetRpmRemoteList provides a mock function with given fields: ctx
func (_m *MockPulpClient) GetRpmRemoteList(ctx context.Context) ([]zest.RpmRpmRemoteResponse, error) {
	ret := _m.Called(ctx)

	var r0 []zest.RpmRpmRemoteResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]zest.RpmRpmRemoteResponse, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []zest.RpmRpmRemoteResponse); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]zest.RpmRpmRemoteResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetRpmRepositoryByName provides a mock function with given fields: ctx, name
func (_m *MockPulpClient) GetRpmRepositoryByName(ctx context.Context, name string) (*zest.RpmRpmRepositoryResponse, error) {
	ret := _m.Called(ctx, name)

	var r0 *zest.RpmRpmRepositoryResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*zest.RpmRpmRepositoryResponse, error)); ok {
		return rf(ctx, name)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *zest.RpmRpmRepositoryResponse); ok {
		r0 = rf(ctx, name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*zest.RpmRpmRepositoryResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetRpmRepositoryByRemote provides a mock function with given fields: ctx, pulpHref
func (_m *MockPulpClient) GetRpmRepositoryByRemote(ctx context.Context, pulpHref string) (*zest.RpmRpmRepositoryResponse, error) {
	ret := _m.Called(ctx, pulpHref)

	var r0 *zest.RpmRpmRepositoryResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*zest.RpmRpmRepositoryResponse, error)); ok {
		return rf(ctx, pulpHref)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *zest.RpmRpmRepositoryResponse); ok {
		r0 = rf(ctx, pulpHref)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*zest.RpmRpmRepositoryResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, pulpHref)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetRpmRepositoryVersion provides a mock function with given fields: ctx, href
func (_m *MockPulpClient) GetRpmRepositoryVersion(ctx context.Context, href string) (*zest.RepositoryVersionResponse, error) {
	ret := _m.Called(ctx, href)

	var r0 *zest.RepositoryVersionResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*zest.RepositoryVersionResponse, error)); ok {
		return rf(ctx, href)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *zest.RepositoryVersionResponse); ok {
		r0 = rf(ctx, href)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*zest.RepositoryVersionResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, href)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTask provides a mock function with given fields: ctx, taskHref
func (_m *MockPulpClient) GetTask(ctx context.Context, taskHref string) (zest.TaskResponse, error) {
	ret := _m.Called(ctx, taskHref)

	var r0 zest.TaskResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (zest.TaskResponse, error)); ok {
		return rf(ctx, taskHref)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) zest.TaskResponse); ok {
		r0 = rf(ctx, taskHref)
	} else {
		r0 = ret.Get(0).(zest.TaskResponse)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, taskHref)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// LookupDomain provides a mock function with given fields: ctx, name
func (_m *MockPulpClient) LookupDomain(ctx context.Context, name string) (string, error) {
	ret := _m.Called(ctx, name)

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (string, error)); ok {
		return rf(ctx, name)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) string); ok {
		r0 = rf(ctx, name)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// LookupOrCreateDomain provides a mock function with given fields: ctx, name
func (_m *MockPulpClient) LookupOrCreateDomain(ctx context.Context, name string) (string, error) {
	ret := _m.Called(ctx, name)

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (string, error)); ok {
		return rf(ctx, name)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) string); ok {
		r0 = rf(ctx, name)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// OrphanCleanup provides a mock function with given fields: ctx
func (_m *MockPulpClient) OrphanCleanup(ctx context.Context) (string, error) {
	ret := _m.Called(ctx)

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (string, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) string); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PollTask provides a mock function with given fields: ctx, taskHref
func (_m *MockPulpClient) PollTask(ctx context.Context, taskHref string) (*zest.TaskResponse, error) {
	ret := _m.Called(ctx, taskHref)

	var r0 *zest.TaskResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*zest.TaskResponse, error)); ok {
		return rf(ctx, taskHref)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *zest.TaskResponse); ok {
		r0 = rf(ctx, taskHref)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*zest.TaskResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, taskHref)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RepairRpmRepositoryVersion provides a mock function with given fields: ctx, href
func (_m *MockPulpClient) RepairRpmRepositoryVersion(ctx context.Context, href string) (string, error) {
	ret := _m.Called(ctx, href)

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (string, error)); ok {
		return rf(ctx, href)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) string); ok {
		r0 = rf(ctx, href)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, href)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Status provides a mock function with given fields: ctx
func (_m *MockPulpClient) Status(ctx context.Context) (*zest.StatusResponse, error) {
	ret := _m.Called(ctx)

	var r0 *zest.StatusResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (*zest.StatusResponse, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) *zest.StatusResponse); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*zest.StatusResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SyncRpmRepository provides a mock function with given fields: ctx, rpmRpmRepositoryHref, remoteHref
func (_m *MockPulpClient) SyncRpmRepository(ctx context.Context, rpmRpmRepositoryHref string, remoteHref *string) (string, error) {
	ret := _m.Called(ctx, rpmRpmRepositoryHref, remoteHref)

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, *string) (string, error)); ok {
		return rf(ctx, rpmRpmRepositoryHref, remoteHref)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, *string) string); ok {
		r0 = rf(ctx, rpmRpmRepositoryHref, remoteHref)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, *string) error); ok {
		r1 = rf(ctx, rpmRpmRepositoryHref, remoteHref)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateDomainIfNeeded provides a mock function with given fields: ctx, name
func (_m *MockPulpClient) UpdateDomainIfNeeded(ctx context.Context, name string) error {
	ret := _m.Called(ctx, name)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, name)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateRpmRemote provides a mock function with given fields: ctx, pulpHref, url, clientCert, clientKey, caCert
func (_m *MockPulpClient) UpdateRpmRemote(ctx context.Context, pulpHref string, url string, clientCert *string, clientKey *string, caCert *string) (string, error) {
	ret := _m.Called(ctx, pulpHref, url, clientCert, clientKey, caCert)

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, *string, *string, *string) (string, error)); ok {
		return rf(ctx, pulpHref, url, clientCert, clientKey, caCert)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string, *string, *string, *string) string); ok {
		r0 = rf(ctx, pulpHref, url, clientCert, clientKey, caCert)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string, *string, *string, *string) error); ok {
		r1 = rf(ctx, pulpHref, url, clientCert, clientKey, caCert)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// WithDomain provides a mock function with given fields: domainName
func (_m *MockPulpClient) WithDomain(domainName string) PulpClient {
	ret := _m.Called(domainName)

	var r0 PulpClient
	if rf, ok := ret.Get(0).(func(string) PulpClient); ok {
		r0 = rf(domainName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(PulpClient)
		}
	}

	return r0
}

// NewMockPulpClient creates a new instance of MockPulpClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockPulpClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockPulpClient {
	mock := &MockPulpClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
