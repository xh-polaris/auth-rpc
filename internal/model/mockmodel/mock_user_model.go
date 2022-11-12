// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/xh-polaris/account-svc/model (interfaces: UserModel)

// Package mock_model is a generated GoMock package.
package mockmodel

import (
	context "context"
	"github.com/xh-polaris/account-rpc/v3/internal/model"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockUserModel is a mock of UserModel interface.
type MockUserModel struct {
	ctrl     *gomock.Controller
	recorder *MockUserModelMockRecorder
}

// MockUserModelMockRecorder is the mock recorder for MockUserModel.
type MockUserModelMockRecorder struct {
	mock *MockUserModel
}

// NewMockUserModel creates a new mock instance.
func NewMockUserModel(ctrl *gomock.Controller) *MockUserModel {
	mock := &MockUserModel{ctrl: ctrl}
	mock.recorder = &MockUserModelMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserModel) EXPECT() *MockUserModelMockRecorder {
	return m.recorder
}

// Delete mocks base method.
func (m *MockUserModel) Delete(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockUserModelMockRecorder) Delete(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockUserModel)(nil).Delete), arg0, arg1)
}

// FindOne mocks base method.
func (m *MockUserModel) FindOne(arg0 context.Context, arg1 string) (*model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindOne", arg0, arg1)
	ret0, _ := ret[0].(*model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindOne indicates an expected call of FindOne.
func (mr *MockUserModelMockRecorder) FindOne(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindOne", reflect.TypeOf((*MockUserModel)(nil).FindOne), arg0, arg1)
}

// FindOneByAuth mocks base method.
func (m *MockUserModel) FindOneByAuth(arg0 context.Context, arg1 model.Auth) (*model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindOneByAuth", arg0, arg1)
	ret0, _ := ret[0].(*model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindOneByAuth indicates an expected call of FindOneByAuth.
func (mr *MockUserModelMockRecorder) FindOneByAuth(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindOneByAuth", reflect.TypeOf((*MockUserModel)(nil).FindOneByAuth), arg0, arg1)
}

// Insert mocks base method.
func (m *MockUserModel) Insert(arg0 context.Context, arg1 *model.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Insert", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Insert indicates an expected call of Insert.
func (mr *MockUserModelMockRecorder) Insert(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockUserModel)(nil).Insert), arg0, arg1)
}

// Update mocks base method.
func (m *MockUserModel) Update(arg0 context.Context, arg1 *model.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockUserModelMockRecorder) Update(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockUserModel)(nil).Update), arg0, arg1)
}
