// Code generated by mockery v2.16.0. DO NOT EDIT.

package mocks

import (
	events "opensource.turistikrota.com/shared/events"

	mock "github.com/stretchr/testify/mock"
)

// Subscriber is an autogenerated mock type for the Subscriber type
type Subscriber struct {
	mock.Mock
}

// Subscribe provides a mock function with given fields: event, handler
func (_m *Subscriber) Subscribe(event string, handler events.Handler) error {
	ret := _m.Called(event, handler)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, events.Handler) error); ok {
		r0 = rf(event, handler)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Unmarshal provides a mock function with given fields: data, v
func (_m *Subscriber) Unmarshal(data []byte, v interface{}) error {
	ret := _m.Called(data, v)

	var r0 error
	if rf, ok := ret.Get(0).(func([]byte, interface{}) error); ok {
		r0 = rf(data, v)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Unsubscribe provides a mock function with given fields: event, handler
func (_m *Subscriber) Unsubscribe(event string, handler events.Handler) error {
	ret := _m.Called(event, handler)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, events.Handler) error); ok {
		r0 = rf(event, handler)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewSubscriber interface {
	mock.TestingT
	Cleanup(func())
}

// NewSubscriber creates a new instance of Subscriber. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewSubscriber(t mockConstructorTestingTNewSubscriber) *Subscriber {
	mock := &Subscriber{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
