// Code generated by mockery v1.0.0. DO NOT EDIT.

// Regenerate this file using `make store-mocks`.

package mocks

import mock "github.com/stretchr/testify/mock"
import model "github.com/mattermost/mattermost-server/model"
import store "github.com/mattermost/mattermost-server/store"

// SystemStore is an autogenerated mock type for the SystemStore type
type SystemStore struct {
	mock.Mock
}

// Get provides a mock function with given fields:
func (_m *SystemStore) Get() store.StoreChannel {
	ret := _m.Called()

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func() store.StoreChannel); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}

// GetByName provides a mock function with given fields: name
func (_m *SystemStore) GetByName(name string) store.StoreChannel {
	ret := _m.Called(name)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(string) store.StoreChannel); ok {
		r0 = rf(name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}

// PermanentDeleteByName provides a mock function with given fields: name
func (_m *SystemStore) PermanentDeleteByName(name string) store.StoreChannel {
	ret := _m.Called(name)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(string) store.StoreChannel); ok {
		r0 = rf(name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}

// Save provides a mock function with given fields: system
func (_m *SystemStore) Save(system *model.System) store.StoreChannel {
	ret := _m.Called(system)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(*model.System) store.StoreChannel); ok {
		r0 = rf(system)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}

// SaveOrUpdate provides a mock function with given fields: system
func (_m *SystemStore) SaveOrUpdate(system *model.System) store.StoreChannel {
	ret := _m.Called(system)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(*model.System) store.StoreChannel); ok {
		r0 = rf(system)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}

// Update provides a mock function with given fields: system
func (_m *SystemStore) Update(system *model.System) store.StoreChannel {
	ret := _m.Called(system)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(*model.System) store.StoreChannel); ok {
		r0 = rf(system)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}
