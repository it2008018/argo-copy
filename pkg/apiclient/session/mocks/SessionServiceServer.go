// Code generated by mockery v2.15.0. DO NOT EDIT.

package mocks

import (
	context "context"

	session "github.com/argoproj/argo-cd/v2/pkg/apiclient/session"
	mock "github.com/stretchr/testify/mock"
)

// SessionServiceServer is an autogenerated mock type for the SessionServiceServer type
type SessionServiceServer struct {
	mock.Mock
}

// Create provides a mock function with given fields: _a0, _a1
func (_m *SessionServiceServer) Create(_a0 context.Context, _a1 *session.SessionCreateRequest) (*session.SessionResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *session.SessionResponse
	if rf, ok := ret.Get(0).(func(context.Context, *session.SessionCreateRequest) *session.SessionResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*session.SessionResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *session.SessionCreateRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: _a0, _a1
func (_m *SessionServiceServer) Delete(_a0 context.Context, _a1 *session.SessionDeleteRequest) (*session.SessionResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *session.SessionResponse
	if rf, ok := ret.Get(0).(func(context.Context, *session.SessionDeleteRequest) *session.SessionResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*session.SessionResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *session.SessionDeleteRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUserInfo provides a mock function with given fields: _a0, _a1
func (_m *SessionServiceServer) GetUserInfo(_a0 context.Context, _a1 *session.GetUserInfoRequest) (*session.GetUserInfoResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *session.GetUserInfoResponse
	if rf, ok := ret.Get(0).(func(context.Context, *session.GetUserInfoRequest) *session.GetUserInfoResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*session.GetUserInfoResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *session.GetUserInfoRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewSessionServiceServer interface {
	mock.TestingT
	Cleanup(func())
}

// NewSessionServiceServer creates a new instance of SessionServiceServer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewSessionServiceServer(t mockConstructorTestingTNewSessionServiceServer) *SessionServiceServer {
	mock := &SessionServiceServer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
