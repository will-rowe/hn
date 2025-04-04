// Code generated by mockery v2.51.1. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// MockUnsafeReportServiceServer is an autogenerated mock type for the UnsafeReportServiceServer type
type MockUnsafeReportServiceServer struct {
	mock.Mock
}

type MockUnsafeReportServiceServer_Expecter struct {
	mock *mock.Mock
}

func (_m *MockUnsafeReportServiceServer) EXPECT() *MockUnsafeReportServiceServer_Expecter {
	return &MockUnsafeReportServiceServer_Expecter{mock: &_m.Mock}
}

// mustEmbedUnimplementedReportServiceServer provides a mock function with no fields
func (_m *MockUnsafeReportServiceServer) mustEmbedUnimplementedReportServiceServer() {
	_m.Called()
}

// MockUnsafeReportServiceServer_mustEmbedUnimplementedReportServiceServer_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'mustEmbedUnimplementedReportServiceServer'
type MockUnsafeReportServiceServer_mustEmbedUnimplementedReportServiceServer_Call struct {
	*mock.Call
}

// mustEmbedUnimplementedReportServiceServer is a helper method to define mock.On call
func (_e *MockUnsafeReportServiceServer_Expecter) mustEmbedUnimplementedReportServiceServer() *MockUnsafeReportServiceServer_mustEmbedUnimplementedReportServiceServer_Call {
	return &MockUnsafeReportServiceServer_mustEmbedUnimplementedReportServiceServer_Call{Call: _e.mock.On("mustEmbedUnimplementedReportServiceServer")}
}

func (_c *MockUnsafeReportServiceServer_mustEmbedUnimplementedReportServiceServer_Call) Run(run func()) *MockUnsafeReportServiceServer_mustEmbedUnimplementedReportServiceServer_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockUnsafeReportServiceServer_mustEmbedUnimplementedReportServiceServer_Call) Return() *MockUnsafeReportServiceServer_mustEmbedUnimplementedReportServiceServer_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockUnsafeReportServiceServer_mustEmbedUnimplementedReportServiceServer_Call) RunAndReturn(run func()) *MockUnsafeReportServiceServer_mustEmbedUnimplementedReportServiceServer_Call {
	_c.Run(run)
	return _c
}

// NewMockUnsafeReportServiceServer creates a new instance of MockUnsafeReportServiceServer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockUnsafeReportServiceServer(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockUnsafeReportServiceServer {
	mock := &MockUnsafeReportServiceServer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
