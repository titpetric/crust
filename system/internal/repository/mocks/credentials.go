// Code generated by MockGen. DO NOT EDIT.
// Source: system/internal/repository/credentials.go

// Package repository is a generated GoMock package.
package repository

import (
	context "context"
	repository "github.com/crusttech/crust/system/internal/repository"
	types "github.com/crusttech/crust/system/types"
	gomock "github.com/golang/mock/gomock"
	factory "github.com/titpetric/factory"
	reflect "reflect"
)

// MockCredentialsRepository is a mock of CredentialsRepository interface
type MockCredentialsRepository struct {
	ctrl     *gomock.Controller
	recorder *MockCredentialsRepositoryMockRecorder
}

// MockCredentialsRepositoryMockRecorder is the mock recorder for MockCredentialsRepository
type MockCredentialsRepositoryMockRecorder struct {
	mock *MockCredentialsRepository
}

// NewMockCredentialsRepository creates a new mock instance
func NewMockCredentialsRepository(ctrl *gomock.Controller) *MockCredentialsRepository {
	mock := &MockCredentialsRepository{ctrl: ctrl}
	mock.recorder = &MockCredentialsRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCredentialsRepository) EXPECT() *MockCredentialsRepositoryMockRecorder {
	return m.recorder
}

// With mocks base method
func (m *MockCredentialsRepository) With(ctx context.Context, db *factory.DB) repository.CredentialsRepository {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "With", ctx, db)
	ret0, _ := ret[0].(repository.CredentialsRepository)
	return ret0
}

// With indicates an expected call of With
func (mr *MockCredentialsRepositoryMockRecorder) With(ctx, db interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "With", reflect.TypeOf((*MockCredentialsRepository)(nil).With), ctx, db)
}

// FindByID mocks base method
func (m *MockCredentialsRepository) FindByID(ID uint64) (*types.Credentials, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByID", ID)
	ret0, _ := ret[0].(*types.Credentials)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByID indicates an expected call of FindByID
func (mr *MockCredentialsRepositoryMockRecorder) FindByID(ID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByID", reflect.TypeOf((*MockCredentialsRepository)(nil).FindByID), ID)
}

// FindByCredentials mocks base method
func (m *MockCredentialsRepository) FindByCredentials(kind, credentials string) (types.CredentialsSet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByCredentials", kind, credentials)
	ret0, _ := ret[0].(types.CredentialsSet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByCredentials indicates an expected call of FindByCredentials
func (mr *MockCredentialsRepositoryMockRecorder) FindByCredentials(kind, credentials interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByCredentials", reflect.TypeOf((*MockCredentialsRepository)(nil).FindByCredentials), kind, credentials)
}

// FindByKind mocks base method
func (m *MockCredentialsRepository) FindByKind(ownerID uint64, kind string) (types.CredentialsSet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByKind", ownerID, kind)
	ret0, _ := ret[0].(types.CredentialsSet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByKind indicates an expected call of FindByKind
func (mr *MockCredentialsRepositoryMockRecorder) FindByKind(ownerID, kind interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByKind", reflect.TypeOf((*MockCredentialsRepository)(nil).FindByKind), ownerID, kind)
}

// FindByOwnerID mocks base method
func (m *MockCredentialsRepository) FindByOwnerID(ownerID uint64) (types.CredentialsSet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByOwnerID", ownerID)
	ret0, _ := ret[0].(types.CredentialsSet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByOwnerID indicates an expected call of FindByOwnerID
func (mr *MockCredentialsRepositoryMockRecorder) FindByOwnerID(ownerID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByOwnerID", reflect.TypeOf((*MockCredentialsRepository)(nil).FindByOwnerID), ownerID)
}

// Find mocks base method
func (m *MockCredentialsRepository) Find() (types.CredentialsSet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Find")
	ret0, _ := ret[0].(types.CredentialsSet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Find indicates an expected call of Find
func (mr *MockCredentialsRepositoryMockRecorder) Find() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockCredentialsRepository)(nil).Find))
}

// Create mocks base method
func (m *MockCredentialsRepository) Create(c *types.Credentials) (*types.Credentials, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", c)
	ret0, _ := ret[0].(*types.Credentials)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create
func (mr *MockCredentialsRepositoryMockRecorder) Create(c interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockCredentialsRepository)(nil).Create), c)
}

// Update mocks base method
func (m *MockCredentialsRepository) Update(c *types.Credentials) (*types.Credentials, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", c)
	ret0, _ := ret[0].(*types.Credentials)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update
func (mr *MockCredentialsRepositoryMockRecorder) Update(c interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockCredentialsRepository)(nil).Update), c)
}

// DeleteByID mocks base method
func (m *MockCredentialsRepository) DeleteByID(id uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteByID", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteByID indicates an expected call of DeleteByID
func (mr *MockCredentialsRepositoryMockRecorder) DeleteByID(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteByID", reflect.TypeOf((*MockCredentialsRepository)(nil).DeleteByID), id)
}
