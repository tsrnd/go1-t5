// Code generated by MockGen. DO NOT EDIT.
// Source: user_usecase.go

// Package mock_usecase is a generated GoMock package.
package mock_usecase

import (
	gomock "github.com/golang/mock/gomock"
	user "github.com/tsrnd/goweb5/frontend/user"
	reflect "reflect"
)

// MockUserUsecase is a mock of UserUsecase interface
type MockUserUsecase struct {
	ctrl     *gomock.Controller
	recorder *MockUserUsecaseMockRecorder
}

// MockUserUsecaseMockRecorder is the mock recorder for MockUserUsecase
type MockUserUsecaseMockRecorder struct {
	mock *MockUserUsecase
}

// NewMockUserUsecase creates a new mock instance
func NewMockUserUsecase(ctrl *gomock.Controller) *MockUserUsecase {
	mock := &MockUserUsecase{ctrl: ctrl}
	mock.recorder = &MockUserUsecaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockUserUsecase) EXPECT() *MockUserUsecaseMockRecorder {
	return m.recorder
}

// GetByID mocks base method
func (m *MockUserUsecase) GetByID(id int) (*user.User, error) {
	ret := m.ctrl.Call(m, "GetByID", id)
	ret0, _ := ret[0].(*user.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID
func (mr *MockUserUsecaseMockRecorder) GetByID(id interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockUserUsecase)(nil).GetByID), id)
}

// GetByEmail mocks base method
func (m *MockUserUsecase) GetByEmail(email string) (*user.User, error) {
	ret := m.ctrl.Call(m, "GetByEmail", email)
	ret0, _ := ret[0].(*user.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByEmail indicates an expected call of GetByEmail
func (mr *MockUserUsecaseMockRecorder) GetByEmail(email interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByEmail", reflect.TypeOf((*MockUserUsecase)(nil).GetByEmail), email)
}

// GetPrivateUserDetailsByEmail mocks base method
func (m *MockUserUsecase) GetPrivateUserDetailsByEmail(email string) (*user.PrivateUserDetails, error) {
	ret := m.ctrl.Call(m, "GetPrivateUserDetailsByEmail", email)
	ret0, _ := ret[0].(*user.PrivateUserDetails)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPrivateUserDetailsByEmail indicates an expected call of GetPrivateUserDetailsByEmail
func (mr *MockUserUsecaseMockRecorder) GetPrivateUserDetailsByEmail(email interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPrivateUserDetailsByEmail", reflect.TypeOf((*MockUserUsecase)(nil).GetPrivateUserDetailsByEmail), email)
}

// Create mocks base method
func (m *MockUserUsecase) Create(email, name, password string) (int, error) {
	ret := m.ctrl.Call(m, "Create", email, name, password)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create
func (mr *MockUserUsecaseMockRecorder) Create(email, name, password interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockUserUsecase)(nil).Create), email, name, password)
}
