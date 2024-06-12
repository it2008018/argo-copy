// Code generated by mockery v2.40.2. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// Repos is an autogenerated mock type for the Repos type
type Repos struct {
	mock.Mock
}

// GetDirectories provides a mock function with given fields: ctx, repoURL, revision, noRevisionCache, verifyCommit
func (_m *Repos) GetDirectories(ctx context.Context, repoURL string, revision string, noRevisionCache bool, verifyCommit bool) ([]string, error) {
	ret := _m.Called(ctx, repoURL, revision, noRevisionCache, verifyCommit)

	if len(ret) == 0 {
		panic("no return value specified for GetDirectories")
	}

	var r0 []string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, bool, bool) ([]string, error)); ok {
		return rf(ctx, repoURL, revision, noRevisionCache, verifyCommit)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string, bool, bool) []string); ok {
		r0 = rf(ctx, repoURL, revision, noRevisionCache, verifyCommit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string, bool, bool) error); ok {
		r1 = rf(ctx, repoURL, revision, noRevisionCache, verifyCommit)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetFiles provides a mock function with given fields: ctx, repoURL, revision, pattern, noRevisionCache, verifyCommit
func (_m *Repos) GetFiles(ctx context.Context, repoURL string, revision string, pattern string, noRevisionCache bool, verifyCommit bool) (map[string][]byte, error) {
	ret := _m.Called(ctx, repoURL, revision, pattern, noRevisionCache, verifyCommit)

	if len(ret) == 0 {
		panic("no return value specified for GetFiles")
	}

	var r0 map[string][]byte
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string, bool, bool) (map[string][]byte, error)); ok {
		return rf(ctx, repoURL, revision, pattern, noRevisionCache, verifyCommit)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string, bool, bool) map[string][]byte); ok {
		r0 = rf(ctx, repoURL, revision, pattern, noRevisionCache, verifyCommit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string][]byte)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string, string, bool, bool) error); ok {
		r1 = rf(ctx, repoURL, revision, pattern, noRevisionCache, verifyCommit)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewRepos creates a new instance of Repos. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewRepos(t interface {
	mock.TestingT
	Cleanup(func())
}) *Repos {
	mock := &Repos{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
