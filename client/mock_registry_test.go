// Code generated by mockery v1.0.0. DO NOT EDIT.

package client

import (
	agentpb "github.com/percona/pmm/api/agentpb"
	mock "github.com/stretchr/testify/mock"
)

// mockRegistry is an autogenerated mock type for the registry type
type mockRegistry struct {
	mock.Mock
}

// SetState provides a mock function with given fields: _a0
func (_m *mockRegistry) SetState(_a0 map[string]*agentpb.SetStateRequest_Tunnel) {
	_m.Called(_a0)
}
