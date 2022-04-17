// Code generated by mockery v2.10.6. DO NOT EDIT.

package service

import mock "github.com/stretchr/testify/mock"

// MockEmailSender is an autogenerated mock type for the EmailSender type
type MockEmailSender struct {
	mock.Mock
}

// Send provides a mock function with given fields: subject, body, to
func (_m *MockEmailSender) Send(subject string, body string, to []string) error {
	ret := _m.Called(subject, body, to)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, []string) error); ok {
		r0 = rf(subject, body, to)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
