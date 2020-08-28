// Code generated by MockGen. DO NOT EDIT.
// Source: usecase.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	ent "github.com/awesomebusiness/uinvest/ent"
	model "github.com/awesomebusiness/uinvest/internal/model"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockUsecaseAuthentication is a mock of UsecaseAuthentication interface
type MockUsecaseAuthentication struct {
	ctrl     *gomock.Controller
	recorder *MockUsecaseAuthenticationMockRecorder
}

// MockUsecaseAuthenticationMockRecorder is the mock recorder for MockUsecaseAuthentication
type MockUsecaseAuthenticationMockRecorder struct {
	mock *MockUsecaseAuthentication
}

// NewMockUsecaseAuthentication creates a new mock instance
func NewMockUsecaseAuthentication(ctrl *gomock.Controller) *MockUsecaseAuthentication {
	mock := &MockUsecaseAuthentication{ctrl: ctrl}
	mock.recorder = &MockUsecaseAuthenticationMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockUsecaseAuthentication) EXPECT() *MockUsecaseAuthenticationMockRecorder {
	return m.recorder
}

// AuthenticationValidation mocks base method
func (m *MockUsecaseAuthentication) AuthenticationValidation(ctx context.Context, input model.LoginInput) (*ent.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AuthenticationValidation", ctx, input)
	ret0, _ := ret[0].(*ent.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AuthenticationValidation indicates an expected call of AuthenticationValidation
func (mr *MockUsecaseAuthenticationMockRecorder) AuthenticationValidation(ctx, input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AuthenticationValidation", reflect.TypeOf((*MockUsecaseAuthentication)(nil).AuthenticationValidation), ctx, input)
}

// RegisterValidation mocks base method
func (m *MockUsecaseAuthentication) RegisterValidation(ctx context.Context, input model.RegisterInput) (*ent.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegisterValidation", ctx, input)
	ret0, _ := ret[0].(*ent.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RegisterValidation indicates an expected call of RegisterValidation
func (mr *MockUsecaseAuthenticationMockRecorder) RegisterValidation(ctx, input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterValidation", reflect.TypeOf((*MockUsecaseAuthentication)(nil).RegisterValidation), ctx, input)
}
