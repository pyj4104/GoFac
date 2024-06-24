// Code generated by mockery v2.43.2. DO NOT EDIT.

package samplestructs

import mock "github.com/stretchr/testify/mock"

// IStructRelyingOnIndependentStruct is an autogenerated mock type for the IStructRelyingOnIndependentStruct type
type IStructRelyingOnIndependentStruct struct {
	mock.Mock
}

type IStructRelyingOnIndependentStruct_Expecter struct {
	mock *mock.Mock
}

func (_m *IStructRelyingOnIndependentStruct) EXPECT() *IStructRelyingOnIndependentStruct_Expecter {
	return &IStructRelyingOnIndependentStruct_Expecter{mock: &_m.Mock}
}

// ReturnNameStructRelyingOnIndependentStruct provides a mock function with given fields:
func (_m *IStructRelyingOnIndependentStruct) ReturnNameStructRelyingOnIndependentStruct() string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for ReturnNameStructRelyingOnIndependentStruct")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// IStructRelyingOnIndependentStruct_ReturnNameStructRelyingOnIndependentStruct_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ReturnNameStructRelyingOnIndependentStruct'
type IStructRelyingOnIndependentStruct_ReturnNameStructRelyingOnIndependentStruct_Call struct {
	*mock.Call
}

// ReturnNameStructRelyingOnIndependentStruct is a helper method to define mock.On call
func (_e *IStructRelyingOnIndependentStruct_Expecter) ReturnNameStructRelyingOnIndependentStruct() *IStructRelyingOnIndependentStruct_ReturnNameStructRelyingOnIndependentStruct_Call {
	return &IStructRelyingOnIndependentStruct_ReturnNameStructRelyingOnIndependentStruct_Call{Call: _e.mock.On("ReturnNameStructRelyingOnIndependentStruct")}
}

func (_c *IStructRelyingOnIndependentStruct_ReturnNameStructRelyingOnIndependentStruct_Call) Run(run func()) *IStructRelyingOnIndependentStruct_ReturnNameStructRelyingOnIndependentStruct_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *IStructRelyingOnIndependentStruct_ReturnNameStructRelyingOnIndependentStruct_Call) Return(_a0 string) *IStructRelyingOnIndependentStruct_ReturnNameStructRelyingOnIndependentStruct_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *IStructRelyingOnIndependentStruct_ReturnNameStructRelyingOnIndependentStruct_Call) RunAndReturn(run func() string) *IStructRelyingOnIndependentStruct_ReturnNameStructRelyingOnIndependentStruct_Call {
	_c.Call.Return(run)
	return _c
}

// ReturnSubStructName provides a mock function with given fields:
func (_m *IStructRelyingOnIndependentStruct) ReturnSubStructName() string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for ReturnSubStructName")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// IStructRelyingOnIndependentStruct_ReturnSubStructName_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ReturnSubStructName'
type IStructRelyingOnIndependentStruct_ReturnSubStructName_Call struct {
	*mock.Call
}

// ReturnSubStructName is a helper method to define mock.On call
func (_e *IStructRelyingOnIndependentStruct_Expecter) ReturnSubStructName() *IStructRelyingOnIndependentStruct_ReturnSubStructName_Call {
	return &IStructRelyingOnIndependentStruct_ReturnSubStructName_Call{Call: _e.mock.On("ReturnSubStructName")}
}

func (_c *IStructRelyingOnIndependentStruct_ReturnSubStructName_Call) Run(run func()) *IStructRelyingOnIndependentStruct_ReturnSubStructName_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *IStructRelyingOnIndependentStruct_ReturnSubStructName_Call) Return(_a0 string) *IStructRelyingOnIndependentStruct_ReturnSubStructName_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *IStructRelyingOnIndependentStruct_ReturnSubStructName_Call) RunAndReturn(run func() string) *IStructRelyingOnIndependentStruct_ReturnSubStructName_Call {
	_c.Call.Return(run)
	return _c
}

// NewIStructRelyingOnIndependentStruct creates a new instance of IStructRelyingOnIndependentStruct. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewIStructRelyingOnIndependentStruct(t interface {
	mock.TestingT
	Cleanup(func())
}) *IStructRelyingOnIndependentStruct {
	mock := &IStructRelyingOnIndependentStruct{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}