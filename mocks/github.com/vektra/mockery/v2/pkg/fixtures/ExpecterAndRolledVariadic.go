// Code generated by mockery. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// ExpecterAndRolledVariadic is an autogenerated mock type for the Expecter type
type ExpecterAndRolledVariadic struct {
	mock.Mock
}

type ExpecterAndRolledVariadic_Expecter struct {
	mock *mock.Mock
}

func (_m *ExpecterAndRolledVariadic) EXPECT() *ExpecterAndRolledVariadic_Expecter {
	return &ExpecterAndRolledVariadic_Expecter{mock: &_m.Mock}
}

// ManyArgsReturns provides a mock function with given fields: str, i
func (_m *ExpecterAndRolledVariadic) ManyArgsReturns(str string, i int) ([]string, error) {
	ret := _m.Called(str, i)

	var r0 []string
	var r1 error
	if rf, ok := ret.Get(0).(func(string, int) ([]string, error)); ok {
		return rf(str, i)
	}
	if rf, ok := ret.Get(0).(func(string, int) []string); ok {
		r0 = rf(str, i)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	if rf, ok := ret.Get(1).(func(string, int) error); ok {
		r1 = rf(str, i)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ExpecterAndRolledVariadic_ManyArgsReturns_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ManyArgsReturns'
type ExpecterAndRolledVariadic_ManyArgsReturns_Call struct {
	*mock.Call
}

// ManyArgsReturns is a helper method to define mock.On call
//   - str string
//   - i int
func (_e *ExpecterAndRolledVariadic_Expecter) ManyArgsReturns(str interface{}, i interface{}) *ExpecterAndRolledVariadic_ManyArgsReturns_Call {
	return &ExpecterAndRolledVariadic_ManyArgsReturns_Call{Call: _e.mock.On("ManyArgsReturns", str, i)}
}

func (_c *ExpecterAndRolledVariadic_ManyArgsReturns_Call) Run(run func(str string, i int)) *ExpecterAndRolledVariadic_ManyArgsReturns_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(int))
	})
	return _c
}

func (_c *ExpecterAndRolledVariadic_ManyArgsReturns_Call) Return(strs []string, err error) *ExpecterAndRolledVariadic_ManyArgsReturns_Call {
	_c.Call.Return(strs, err)
	return _c
}

func (_c *ExpecterAndRolledVariadic_ManyArgsReturns_Call) RunAndReturn(run func(string, int) ([]string, error)) *ExpecterAndRolledVariadic_ManyArgsReturns_Call {
	_c.Call.Return(run)
	return _c
}

// NoArg provides a mock function with given fields:
func (_m *ExpecterAndRolledVariadic) NoArg() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// ExpecterAndRolledVariadic_NoArg_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'NoArg'
type ExpecterAndRolledVariadic_NoArg_Call struct {
	*mock.Call
}

// NoArg is a helper method to define mock.On call
func (_e *ExpecterAndRolledVariadic_Expecter) NoArg() *ExpecterAndRolledVariadic_NoArg_Call {
	return &ExpecterAndRolledVariadic_NoArg_Call{Call: _e.mock.On("NoArg")}
}

func (_c *ExpecterAndRolledVariadic_NoArg_Call) Run(run func()) *ExpecterAndRolledVariadic_NoArg_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *ExpecterAndRolledVariadic_NoArg_Call) Return(_a0 string) *ExpecterAndRolledVariadic_NoArg_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ExpecterAndRolledVariadic_NoArg_Call) RunAndReturn(run func() string) *ExpecterAndRolledVariadic_NoArg_Call {
	_c.Call.Return(run)
	return _c
}

// NoReturn provides a mock function with given fields: str
func (_m *ExpecterAndRolledVariadic) NoReturn(str string) {
	_m.Called(str)
}

// ExpecterAndRolledVariadic_NoReturn_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'NoReturn'
type ExpecterAndRolledVariadic_NoReturn_Call struct {
	*mock.Call
}

// NoReturn is a helper method to define mock.On call
//   - str string
func (_e *ExpecterAndRolledVariadic_Expecter) NoReturn(str interface{}) *ExpecterAndRolledVariadic_NoReturn_Call {
	return &ExpecterAndRolledVariadic_NoReturn_Call{Call: _e.mock.On("NoReturn", str)}
}

func (_c *ExpecterAndRolledVariadic_NoReturn_Call) Run(run func(str string)) *ExpecterAndRolledVariadic_NoReturn_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *ExpecterAndRolledVariadic_NoReturn_Call) Return() *ExpecterAndRolledVariadic_NoReturn_Call {
	_c.Call.Return()
	return _c
}

func (_c *ExpecterAndRolledVariadic_NoReturn_Call) RunAndReturn(run func(string)) *ExpecterAndRolledVariadic_NoReturn_Call {
	_c.Call.Return(run)
	return _c
}

// Variadic provides a mock function with given fields: ints
func (_m *ExpecterAndRolledVariadic) Variadic(ints ...int) error {
	var tmpRet mock.Arguments
	if len(ints) > 0 {
		tmpRet = _m.Called(ints)
	} else {
		tmpRet = _m.Called()
	}

	ret := tmpRet

	var r0 error
	if rf, ok := ret.Get(0).(func(...int) error); ok {
		r0 = rf(ints...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ExpecterAndRolledVariadic_Variadic_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Variadic'
type ExpecterAndRolledVariadic_Variadic_Call struct {
	*mock.Call
}

// Variadic is a helper method to define mock.On call
//   - ints ...int
func (_e *ExpecterAndRolledVariadic_Expecter) Variadic(ints ...interface{}) *ExpecterAndRolledVariadic_Variadic_Call {
	return &ExpecterAndRolledVariadic_Variadic_Call{Call: _e.mock.On("Variadic",
		append([]interface{}{}, ints...)...)}
}

func (_c *ExpecterAndRolledVariadic_Variadic_Call) Run(run func(ints ...int)) *ExpecterAndRolledVariadic_Variadic_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]int, len(args)-0)
		for i, a := range args[0:] {
			if a != nil {
				variadicArgs[i] = a.(int)
			}
		}
		run(variadicArgs...)
	})
	return _c
}

func (_c *ExpecterAndRolledVariadic_Variadic_Call) Return(_a0 error) *ExpecterAndRolledVariadic_Variadic_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ExpecterAndRolledVariadic_Variadic_Call) RunAndReturn(run func(...int) error) *ExpecterAndRolledVariadic_Variadic_Call {
	_c.Call.Return(run)
	return _c
}

// VariadicMany provides a mock function with given fields: i, a, intfs
func (_m *ExpecterAndRolledVariadic) VariadicMany(i int, a string, intfs ...interface{}) error {
	var tmpRet mock.Arguments
	if len(intfs) > 0 {
		tmpRet = _m.Called(i, a, intfs)
	} else {
		tmpRet = _m.Called(i, a)
	}

	ret := tmpRet

	var r0 error
	if rf, ok := ret.Get(0).(func(int, string, ...interface{}) error); ok {
		r0 = rf(i, a, intfs...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ExpecterAndRolledVariadic_VariadicMany_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'VariadicMany'
type ExpecterAndRolledVariadic_VariadicMany_Call struct {
	*mock.Call
}

// VariadicMany is a helper method to define mock.On call
//   - i int
//   - a string
//   - intfs ...interface{}
func (_e *ExpecterAndRolledVariadic_Expecter) VariadicMany(i interface{}, a interface{}, intfs ...interface{}) *ExpecterAndRolledVariadic_VariadicMany_Call {
	return &ExpecterAndRolledVariadic_VariadicMany_Call{Call: _e.mock.On("VariadicMany",
		append([]interface{}{i, a}, intfs...)...)}
}

func (_c *ExpecterAndRolledVariadic_VariadicMany_Call) Run(run func(i int, a string, intfs ...interface{})) *ExpecterAndRolledVariadic_VariadicMany_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]interface{}, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(interface{})
			}
		}
		run(args[0].(int), args[1].(string), variadicArgs...)
	})
	return _c
}

func (_c *ExpecterAndRolledVariadic_VariadicMany_Call) Return(_a0 error) *ExpecterAndRolledVariadic_VariadicMany_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ExpecterAndRolledVariadic_VariadicMany_Call) RunAndReturn(run func(int, string, ...interface{}) error) *ExpecterAndRolledVariadic_VariadicMany_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewExpecterAndRolledVariadic interface {
	mock.TestingT
	Cleanup(func())
}

// NewExpecterAndRolledVariadic creates a new instance of ExpecterAndRolledVariadic. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewExpecterAndRolledVariadic(t mockConstructorTestingTNewExpecterAndRolledVariadic) *ExpecterAndRolledVariadic {
	mock := &ExpecterAndRolledVariadic{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
