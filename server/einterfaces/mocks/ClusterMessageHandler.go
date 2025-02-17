// Code generated by mockery v2.23.2. DO NOT EDIT.

// Regenerate this file using `make einterfaces-mocks`.

package mocks

import (
	model "github.com/mattermost/mattermost-server/server/public/model"
	mock "github.com/stretchr/testify/mock"
)

// ClusterMessageHandler is an autogenerated mock type for the ClusterMessageHandler type
type ClusterMessageHandler struct {
	mock.Mock
}

// Execute provides a mock function with given fields: msg
func (_m *ClusterMessageHandler) Execute(msg *model.ClusterMessage) {
	_m.Called(msg)
}

type mockConstructorTestingTNewClusterMessageHandler interface {
	mock.TestingT
	Cleanup(func())
}

// NewClusterMessageHandler creates a new instance of ClusterMessageHandler. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewClusterMessageHandler(t mockConstructorTestingTNewClusterMessageHandler) *ClusterMessageHandler {
	mock := &ClusterMessageHandler{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
