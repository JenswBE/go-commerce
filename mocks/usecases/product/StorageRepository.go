// Code generated by mockery v2.39.1. DO NOT EDIT.

package product

import mock "github.com/stretchr/testify/mock"

// StorageRepository is an autogenerated mock type for the StorageRepository type
type StorageRepository struct {
	mock.Mock
}

type StorageRepository_Expecter struct {
	mock *mock.Mock
}

func (_m *StorageRepository) EXPECT() *StorageRepository_Expecter {
	return &StorageRepository_Expecter{mock: &_m.Mock}
}

// DeleteFile provides a mock function with given fields: filename
func (_m *StorageRepository) DeleteFile(filename string) error {
	ret := _m.Called(filename)

	if len(ret) == 0 {
		panic("no return value specified for DeleteFile")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(filename)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// StorageRepository_DeleteFile_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteFile'
type StorageRepository_DeleteFile_Call struct {
	*mock.Call
}

// DeleteFile is a helper method to define mock.On call
//   - filename string
func (_e *StorageRepository_Expecter) DeleteFile(filename interface{}) *StorageRepository_DeleteFile_Call {
	return &StorageRepository_DeleteFile_Call{Call: _e.mock.On("DeleteFile", filename)}
}

func (_c *StorageRepository_DeleteFile_Call) Run(run func(filename string)) *StorageRepository_DeleteFile_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *StorageRepository_DeleteFile_Call) Return(_a0 error) *StorageRepository_DeleteFile_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *StorageRepository_DeleteFile_Call) RunAndReturn(run func(string) error) *StorageRepository_DeleteFile_Call {
	_c.Call.Return(run)
	return _c
}

// SaveFile provides a mock function with given fields: filename, content
func (_m *StorageRepository) SaveFile(filename string, content []byte) error {
	ret := _m.Called(filename, content)

	if len(ret) == 0 {
		panic("no return value specified for SaveFile")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string, []byte) error); ok {
		r0 = rf(filename, content)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// StorageRepository_SaveFile_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SaveFile'
type StorageRepository_SaveFile_Call struct {
	*mock.Call
}

// SaveFile is a helper method to define mock.On call
//   - filename string
//   - content []byte
func (_e *StorageRepository_Expecter) SaveFile(filename interface{}, content interface{}) *StorageRepository_SaveFile_Call {
	return &StorageRepository_SaveFile_Call{Call: _e.mock.On("SaveFile", filename, content)}
}

func (_c *StorageRepository_SaveFile_Call) Run(run func(filename string, content []byte)) *StorageRepository_SaveFile_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].([]byte))
	})
	return _c
}

func (_c *StorageRepository_SaveFile_Call) Return(_a0 error) *StorageRepository_SaveFile_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *StorageRepository_SaveFile_Call) RunAndReturn(run func(string, []byte) error) *StorageRepository_SaveFile_Call {
	_c.Call.Return(run)
	return _c
}

// NewStorageRepository creates a new instance of StorageRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewStorageRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *StorageRepository {
	mock := &StorageRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
