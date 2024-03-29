// Code generated by MockGen. DO NOT EDIT.
// Source: service.go

// Package mock_service is a generated GoMock package.
package mock_service

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	models "github.com/ryodanqqe/flight-bookings/models"
	requests "github.com/ryodanqqe/flight-bookings/models/requests"
)

// MockAuthorization is a mock of Authorization interface.
type MockAuthorization struct {
	ctrl     *gomock.Controller
	recorder *MockAuthorizationMockRecorder
}

// MockAuthorizationMockRecorder is the mock recorder for MockAuthorization.
type MockAuthorizationMockRecorder struct {
	mock *MockAuthorization
}

// NewMockAuthorization creates a new mock instance.
func NewMockAuthorization(ctrl *gomock.Controller) *MockAuthorization {
	mock := &MockAuthorization{ctrl: ctrl}
	mock.recorder = &MockAuthorizationMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuthorization) EXPECT() *MockAuthorizationMockRecorder {
	return m.recorder
}

// CreateUser mocks base method.
func (m *MockAuthorization) CreateUser(user models.User) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", user)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockAuthorizationMockRecorder) CreateUser(user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockAuthorization)(nil).CreateUser), user)
}

// GenerateToken mocks base method.
func (m *MockAuthorization) GenerateToken(email, password string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateToken", email, password)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GenerateToken indicates an expected call of GenerateToken.
func (mr *MockAuthorizationMockRecorder) GenerateToken(email, password interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateToken", reflect.TypeOf((*MockAuthorization)(nil).GenerateToken), email, password)
}

// ParseToken mocks base method.
func (m *MockAuthorization) ParseToken(accessToken string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ParseToken", accessToken)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ParseToken indicates an expected call of ParseToken.
func (mr *MockAuthorizationMockRecorder) ParseToken(accessToken interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ParseToken", reflect.TypeOf((*MockAuthorization)(nil).ParseToken), accessToken)
}

// MockAdmin is a mock of Admin interface.
type MockAdmin struct {
	ctrl     *gomock.Controller
	recorder *MockAdminMockRecorder
}

// MockAdminMockRecorder is the mock recorder for MockAdmin.
type MockAdminMockRecorder struct {
	mock *MockAdmin
}

// NewMockAdmin creates a new mock instance.
func NewMockAdmin(ctrl *gomock.Controller) *MockAdmin {
	mock := &MockAdmin{ctrl: ctrl}
	mock.recorder = &MockAdminMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAdmin) EXPECT() *MockAdminMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockAdmin) Create(flight requests.CreateFlightRequest) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", flight)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockAdminMockRecorder) Create(flight interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockAdmin)(nil).Create), flight)
}

// Delete mocks base method.
func (m *MockAdmin) Delete(id string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockAdminMockRecorder) Delete(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockAdmin)(nil).Delete), id)
}

// GetAll mocks base method.
func (m *MockAdmin) GetAll() ([]models.Flight, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll")
	ret0, _ := ret[0].([]models.Flight)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll.
func (mr *MockAdminMockRecorder) GetAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockAdmin)(nil).GetAll))
}

// GetOne mocks base method.
func (m *MockAdmin) GetOne(id string) (models.Flight, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOne", id)
	ret0, _ := ret[0].(models.Flight)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOne indicates an expected call of GetOne.
func (mr *MockAdminMockRecorder) GetOne(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOne", reflect.TypeOf((*MockAdmin)(nil).GetOne), id)
}

// Update mocks base method.
func (m *MockAdmin) Update(id string, req requests.UpdateFlightRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", id, req)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockAdminMockRecorder) Update(id, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockAdmin)(nil).Update), id, req)
}

// MockUser is a mock of User interface.
type MockUser struct {
	ctrl     *gomock.Controller
	recorder *MockUserMockRecorder
}

// MockUserMockRecorder is the mock recorder for MockUser.
type MockUserMockRecorder struct {
	mock *MockUser
}

// NewMockUser creates a new mock instance.
func NewMockUser(ctrl *gomock.Controller) *MockUser {
	mock := &MockUser{ctrl: ctrl}
	mock.recorder = &MockUserMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUser) EXPECT() *MockUserMockRecorder {
	return m.recorder
}

// BookTicket mocks base method.
func (m *MockUser) BookTicket(userID string, req requests.BookTicketRequest) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BookTicket", userID, req)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BookTicket indicates an expected call of BookTicket.
func (mr *MockUserMockRecorder) BookTicket(userID, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BookTicket", reflect.TypeOf((*MockUser)(nil).BookTicket), userID, req)
}

// DeleteUser mocks base method.
func (m *MockUser) DeleteUser(id string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUser", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteUser indicates an expected call of DeleteUser.
func (mr *MockUserMockRecorder) DeleteUser(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUser", reflect.TypeOf((*MockUser)(nil).DeleteUser), id)
}

// DeleteUserBooking mocks base method.
func (m *MockUser) DeleteUserBooking(ticketID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUserBooking", ticketID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteUserBooking indicates an expected call of DeleteUserBooking.
func (mr *MockUserMockRecorder) DeleteUserBooking(ticketID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUserBooking", reflect.TypeOf((*MockUser)(nil).DeleteUserBooking), ticketID)
}

// GetOneUserBooking mocks base method.
func (m *MockUser) GetOneUserBooking(ticketID string) (models.Ticket, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOneUserBooking", ticketID)
	ret0, _ := ret[0].(models.Ticket)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOneUserBooking indicates an expected call of GetOneUserBooking.
func (mr *MockUserMockRecorder) GetOneUserBooking(ticketID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOneUserBooking", reflect.TypeOf((*MockUser)(nil).GetOneUserBooking), ticketID)
}

// GetUserBookings mocks base method.
func (m *MockUser) GetUserBookings(userID string) ([]models.Ticket, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserBookings", userID)
	ret0, _ := ret[0].([]models.Ticket)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserBookings indicates an expected call of GetUserBookings.
func (mr *MockUserMockRecorder) GetUserBookings(userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserBookings", reflect.TypeOf((*MockUser)(nil).GetUserBookings), userID)
}

// UpdateUser mocks base method.
func (m *MockUser) UpdateUser(id string, req requests.UpdateUserRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUser", id, req)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateUser indicates an expected call of UpdateUser.
func (mr *MockUserMockRecorder) UpdateUser(id, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUser", reflect.TypeOf((*MockUser)(nil).UpdateUser), id, req)
}

// UpdateUserBooking mocks base method.
func (m *MockUser) UpdateUserBooking(ticketID string, req requests.UpdateUserBookingRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUserBooking", ticketID, req)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateUserBooking indicates an expected call of UpdateUserBooking.
func (mr *MockUserMockRecorder) UpdateUserBooking(ticketID, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUserBooking", reflect.TypeOf((*MockUser)(nil).UpdateUserBooking), ticketID, req)
}
