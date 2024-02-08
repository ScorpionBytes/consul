// Code generated by mockery v2.37.1. DO NOT EDIT.

package indexersmock

import (
	resource "github.com/hashicorp/consul/internal/resource"
	pbresource "github.com/hashicorp/consul/proto-public/pbresource/v1"
	mock "github.com/stretchr/testify/mock"
)

// GetSingleRefOrID is an autogenerated mock type for the GetSingleRefOrID type
type GetSingleRefOrID struct {
	mock.Mock
}

type GetSingleRefOrID_Expecter struct {
	mock *mock.Mock
}

func (_m *GetSingleRefOrID) EXPECT() *GetSingleRefOrID_Expecter {
	return &GetSingleRefOrID_Expecter{mock: &_m.Mock}
}

// Execute provides a mock function with given fields: _a0
func (_m *GetSingleRefOrID) Execute(_a0 *pbresource.Resource) resource.ReferenceOrID {
	ret := _m.Called(_a0)

	var r0 resource.ReferenceOrID
	if rf, ok := ret.Get(0).(func(*pbresource.Resource) resource.ReferenceOrID); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(resource.ReferenceOrID)
		}
	}

	return r0
}

// GetSingleRefOrID_Execute_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Execute'
type GetSingleRefOrID_Execute_Call struct {
	*mock.Call
}

// Execute is a helper method to define mock.On call
//   - _a0 *pbresource.Resource
func (_e *GetSingleRefOrID_Expecter) Execute(_a0 interface{}) *GetSingleRefOrID_Execute_Call {
	return &GetSingleRefOrID_Execute_Call{Call: _e.mock.On("Execute", _a0)}
}

func (_c *GetSingleRefOrID_Execute_Call) Run(run func(_a0 *pbresource.Resource)) *GetSingleRefOrID_Execute_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*pbresource.Resource))
	})
	return _c
}

func (_c *GetSingleRefOrID_Execute_Call) Return(_a0 resource.ReferenceOrID) *GetSingleRefOrID_Execute_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *GetSingleRefOrID_Execute_Call) RunAndReturn(run func(*pbresource.Resource) resource.ReferenceOrID) *GetSingleRefOrID_Execute_Call {
	_c.Call.Return(run)
	return _c
}

// NewGetSingleRefOrID creates a new instance of GetSingleRefOrID. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewGetSingleRefOrID(t interface {
	mock.TestingT
	Cleanup(func())
}) *GetSingleRefOrID {
	mock := &GetSingleRefOrID{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
