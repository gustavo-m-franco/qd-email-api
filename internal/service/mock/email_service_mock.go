// Code generated by MockGen. DO NOT EDIT.
// Source: ./email_service.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockEmailServicer is a mock of EmailServicer interface.
type MockEmailServicer struct {
	ctrl     *gomock.Controller
	recorder *MockEmailServicerMockRecorder
}

// MockEmailServicerMockRecorder is the mock recorder for MockEmailServicer.
type MockEmailServicerMockRecorder struct {
	mock *MockEmailServicer
}

// NewMockEmailServicer creates a new mock instance.
func NewMockEmailServicer(ctrl *gomock.Controller) *MockEmailServicer {
	mock := &MockEmailServicer{ctrl: ctrl}
	mock.recorder = &MockEmailServicerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEmailServicer) EXPECT() *MockEmailServicerMockRecorder {
	return m.recorder
}

// SendEmail mocks base method.
func (m *MockEmailServicer) SendEmail(ctx context.Context, dest, subject, body string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendEmail", ctx, dest, subject, body)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendEmail indicates an expected call of SendEmail.
func (mr *MockEmailServicerMockRecorder) SendEmail(ctx, dest, subject, body interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendEmail", reflect.TypeOf((*MockEmailServicer)(nil).SendEmail), ctx, dest, subject, body)
}
