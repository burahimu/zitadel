// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/zitadel/zitadel/internal/auth_request/repository (interfaces: AuthRequestCache)
//
// Generated by this command:
//
//	mockgen -package mock -destination ./mock/repository.mock.go github.com/zitadel/zitadel/internal/auth_request/repository AuthRequestCache
//

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	domain "github.com/zitadel/zitadel/internal/domain"
	gomock "go.uber.org/mock/gomock"
)

// MockAuthRequestCache is a mock of AuthRequestCache interface.
type MockAuthRequestCache struct {
	ctrl     *gomock.Controller
	recorder *MockAuthRequestCacheMockRecorder
}

// MockAuthRequestCacheMockRecorder is the mock recorder for MockAuthRequestCache.
type MockAuthRequestCacheMockRecorder struct {
	mock *MockAuthRequestCache
}

// NewMockAuthRequestCache creates a new mock instance.
func NewMockAuthRequestCache(ctrl *gomock.Controller) *MockAuthRequestCache {
	mock := &MockAuthRequestCache{ctrl: ctrl}
	mock.recorder = &MockAuthRequestCacheMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuthRequestCache) EXPECT() *MockAuthRequestCacheMockRecorder {
	return m.recorder
}

// CacheAuthRequest mocks base method.
func (m *MockAuthRequestCache) CacheAuthRequest(arg0 context.Context, arg1 *domain.AuthRequest) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "CacheAuthRequest", arg0, arg1)
}

// CacheAuthRequest indicates an expected call of CacheAuthRequest.
func (mr *MockAuthRequestCacheMockRecorder) CacheAuthRequest(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CacheAuthRequest", reflect.TypeOf((*MockAuthRequestCache)(nil).CacheAuthRequest), arg0, arg1)
}

// DeleteAuthRequest mocks base method.
func (m *MockAuthRequestCache) DeleteAuthRequest(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAuthRequest", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteAuthRequest indicates an expected call of DeleteAuthRequest.
func (mr *MockAuthRequestCacheMockRecorder) DeleteAuthRequest(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAuthRequest", reflect.TypeOf((*MockAuthRequestCache)(nil).DeleteAuthRequest), arg0, arg1)
}

// GetAuthRequestByCode mocks base method.
func (m *MockAuthRequestCache) GetAuthRequestByCode(arg0 context.Context, arg1 string) (*domain.AuthRequest, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAuthRequestByCode", arg0, arg1)
	ret0, _ := ret[0].(*domain.AuthRequest)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAuthRequestByCode indicates an expected call of GetAuthRequestByCode.
func (mr *MockAuthRequestCacheMockRecorder) GetAuthRequestByCode(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAuthRequestByCode", reflect.TypeOf((*MockAuthRequestCache)(nil).GetAuthRequestByCode), arg0, arg1)
}

// GetAuthRequestByID mocks base method.
func (m *MockAuthRequestCache) GetAuthRequestByID(arg0 context.Context, arg1 string) (*domain.AuthRequest, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAuthRequestByID", arg0, arg1)
	ret0, _ := ret[0].(*domain.AuthRequest)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAuthRequestByID indicates an expected call of GetAuthRequestByID.
func (mr *MockAuthRequestCacheMockRecorder) GetAuthRequestByID(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAuthRequestByID", reflect.TypeOf((*MockAuthRequestCache)(nil).GetAuthRequestByID), arg0, arg1)
}

// Health mocks base method.
func (m *MockAuthRequestCache) Health(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Health", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Health indicates an expected call of Health.
func (mr *MockAuthRequestCacheMockRecorder) Health(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Health", reflect.TypeOf((*MockAuthRequestCache)(nil).Health), arg0)
}

// SaveAuthRequest mocks base method.
func (m *MockAuthRequestCache) SaveAuthRequest(arg0 context.Context, arg1 *domain.AuthRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveAuthRequest", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveAuthRequest indicates an expected call of SaveAuthRequest.
func (mr *MockAuthRequestCacheMockRecorder) SaveAuthRequest(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveAuthRequest", reflect.TypeOf((*MockAuthRequestCache)(nil).SaveAuthRequest), arg0, arg1)
}

// UpdateAuthRequest mocks base method.
func (m *MockAuthRequestCache) UpdateAuthRequest(arg0 context.Context, arg1 *domain.AuthRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAuthRequest", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateAuthRequest indicates an expected call of UpdateAuthRequest.
func (mr *MockAuthRequestCacheMockRecorder) UpdateAuthRequest(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAuthRequest", reflect.TypeOf((*MockAuthRequestCache)(nil).UpdateAuthRequest), arg0, arg1)
}
