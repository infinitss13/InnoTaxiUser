// Code generated by MockGen. DO NOT EDIT.
// Source: services/service.go

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	gin "github.com/gin-gonic/gin"
	gomock "github.com/golang/mock/gomock"
	entity "github.com/infinitss13/innotaxiuser/entity"
)

// MockUserService is a mock of UserService interface.
type MockUserService struct {
	ctrl     *gomock.Controller
	recorder *MockUserServiceMockRecorder
}

// MockUserServiceMockRecorder is the mock recorder for MockUserService.
type MockUserServiceMockRecorder struct {
	mock *MockUserService
}

// NewMockUserService creates a new mock instance.
func NewMockUserService(ctrl *gomock.Controller) *MockUserService {
	mock := &MockUserService{ctrl: ctrl}
	mock.recorder = &MockUserServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserService) EXPECT() *MockUserServiceMockRecorder {
	return m.recorder
}

// Auth mocks base method.
func (m *MockUserService) Auth() gin.HandlerFunc {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Auth")
	ret0, _ := ret[0].(gin.HandlerFunc)
	return ret0
}

// Auth indicates an expected call of Auth.
func (mr *MockUserServiceMockRecorder) Auth() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Auth", reflect.TypeOf((*MockUserService)(nil).Auth))
}

// CreateUser mocks base method.
func (m *MockUserService) CreateUser(arg0 entity.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockUserServiceMockRecorder) CreateUser(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockUserService)(nil).CreateUser), arg0)
}

// DeleteProfile mocks base method.
func (m *MockUserService) DeleteProfile(string, password string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteProfile", string, password)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteProfile indicates an expected call of DeleteProfile.
func (mr *MockUserServiceMockRecorder) DeleteProfile(string, password interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteProfile", reflect.TypeOf((*MockUserService)(nil).DeleteProfile), string, password)
}

// GetRatingWithToken mocks base method.
func (m *MockUserService) GetRatingWithToken(arg0 string) (float32, string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRatingWithToken", arg0)
	ret0, _ := ret[0].(float32)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetRatingWithToken indicates an expected call of GetRatingWithToken.
func (mr *MockUserServiceMockRecorder) GetRatingWithToken(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRatingWithToken", reflect.TypeOf((*MockUserService)(nil).GetRatingWithToken), arg0)
}

// GetToken mocks base method.
func (m *MockUserService) GetToken(context *gin.Context) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetToken", context)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetToken indicates an expected call of GetToken.
func (mr *MockUserServiceMockRecorder) GetToken(context interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetToken", reflect.TypeOf((*MockUserService)(nil).GetToken), context)
}

// GetUserByToken mocks base method.
func (m *MockUserService) GetUserByToken(arg0 string) (entity.ProfileData, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByToken", arg0)
	ret0, _ := ret[0].(entity.ProfileData)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByToken indicates an expected call of GetUserByToken.
func (mr *MockUserServiceMockRecorder) GetUserByToken(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByToken", reflect.TypeOf((*MockUserService)(nil).GetUserByToken), arg0)
}

// SignInUser mocks base method.
func (m *MockUserService) SignInUser(arg0 entity.InputSignIn) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SignInUser", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SignInUser indicates an expected call of SignInUser.
func (mr *MockUserServiceMockRecorder) SignInUser(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SignInUser", reflect.TypeOf((*MockUserService)(nil).SignInUser), arg0)
}

// UpdateUserProfile mocks base method.
func (m *MockUserService) UpdateUserProfile(arg0 string, arg1 *entity.UpdateData) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUserProfile", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateUserProfile indicates an expected call of UpdateUserProfile.
func (mr *MockUserServiceMockRecorder) UpdateUserProfile(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUserProfile", reflect.TypeOf((*MockUserService)(nil).UpdateUserProfile), arg0, arg1)
}

// VerifyToken mocks base method.
func (m *MockUserService) VerifyToken(tokenSigned string) (entity.InputSignIn, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VerifyToken", tokenSigned)
	ret0, _ := ret[0].(entity.InputSignIn)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// VerifyToken indicates an expected call of VerifyToken.
func (mr *MockUserServiceMockRecorder) VerifyToken(tokenSigned interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VerifyToken", reflect.TypeOf((*MockUserService)(nil).VerifyToken), tokenSigned)
}
