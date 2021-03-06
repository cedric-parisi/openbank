// Code generated by MockGen. DO NOT EDIT.
// Source: storage/storage.go

// Package mock_storage is a generated GoMock package.
package mock_storage

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	accounts "github.com/openbank/gunk/gunk/v1/accounts"
	transactions "github.com/openbank/gunk/gunk/v1/transactions"
	reflect "reflect"
)

// MockAccountStore is a mock of AccountStore interface
type MockAccountStore struct {
	ctrl     *gomock.Controller
	recorder *MockAccountStoreMockRecorder
}

// MockAccountStoreMockRecorder is the mock recorder for MockAccountStore
type MockAccountStoreMockRecorder struct {
	mock *MockAccountStore
}

// NewMockAccountStore creates a new mock instance
func NewMockAccountStore(ctrl *gomock.Controller) *MockAccountStore {
	mock := &MockAccountStore{ctrl: ctrl}
	mock.recorder = &MockAccountStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAccountStore) EXPECT() *MockAccountStoreMockRecorder {
	return m.recorder
}

// GetAccount mocks base method
func (m *MockAccountStore) GetAccount(ctx context.Context, id string) (*accounts.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAccount", ctx, id)
	ret0, _ := ret[0].(*accounts.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAccount indicates an expected call of GetAccount
func (mr *MockAccountStoreMockRecorder) GetAccount(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccount", reflect.TypeOf((*MockAccountStore)(nil).GetAccount), ctx, id)
}

// GetAccounts mocks base method
func (m *MockAccountStore) GetAccounts(ctx context.Context, nextID string) ([]*accounts.Account, bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAccounts", ctx, nextID)
	ret0, _ := ret[0].([]*accounts.Account)
	ret1, _ := ret[1].(bool)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetAccounts indicates an expected call of GetAccounts
func (mr *MockAccountStoreMockRecorder) GetAccounts(ctx, nextID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccounts", reflect.TypeOf((*MockAccountStore)(nil).GetAccounts), ctx, nextID)
}

// CreateAccount mocks base method
func (m *MockAccountStore) CreateAccount(ctx context.Context, account *accounts.Account) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateAccount", ctx, account)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateAccount indicates an expected call of CreateAccount
func (mr *MockAccountStoreMockRecorder) CreateAccount(ctx, account interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAccount", reflect.TypeOf((*MockAccountStore)(nil).CreateAccount), ctx, account)
}

// UpdateAccount mocks base method
func (m *MockAccountStore) UpdateAccount(ctx context.Context, account *accounts.Account) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAccount", ctx, account)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateAccount indicates an expected call of UpdateAccount
func (mr *MockAccountStoreMockRecorder) UpdateAccount(ctx, account interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAccount", reflect.TypeOf((*MockAccountStore)(nil).UpdateAccount), ctx, account)
}

// DeleteAccount mocks base method
func (m *MockAccountStore) DeleteAccount(ctx context.Context, id string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAccount", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteAccount indicates an expected call of DeleteAccount
func (mr *MockAccountStoreMockRecorder) DeleteAccount(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAccount", reflect.TypeOf((*MockAccountStore)(nil).DeleteAccount), ctx, id)
}

// MockTransactionStore is a mock of TransactionStore interface
type MockTransactionStore struct {
	ctrl     *gomock.Controller
	recorder *MockTransactionStoreMockRecorder
}

// MockTransactionStoreMockRecorder is the mock recorder for MockTransactionStore
type MockTransactionStoreMockRecorder struct {
	mock *MockTransactionStore
}

// NewMockTransactionStore creates a new mock instance
func NewMockTransactionStore(ctrl *gomock.Controller) *MockTransactionStore {
	mock := &MockTransactionStore{ctrl: ctrl}
	mock.recorder = &MockTransactionStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockTransactionStore) EXPECT() *MockTransactionStoreMockRecorder {
	return m.recorder
}

// GetTransaction mocks base method
func (m *MockTransactionStore) GetTransaction(ctx context.Context, id string) (*transactions.Transaction, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTransaction", ctx, id)
	ret0, _ := ret[0].(*transactions.Transaction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTransaction indicates an expected call of GetTransaction
func (mr *MockTransactionStoreMockRecorder) GetTransaction(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTransaction", reflect.TypeOf((*MockTransactionStore)(nil).GetTransaction), ctx, id)
}

// GetTransactions mocks base method
func (m *MockTransactionStore) GetTransactions(ctx context.Context, nextID string) ([]*transactions.Transaction, bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTransactions", ctx, nextID)
	ret0, _ := ret[0].([]*transactions.Transaction)
	ret1, _ := ret[1].(bool)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetTransactions indicates an expected call of GetTransactions
func (mr *MockTransactionStoreMockRecorder) GetTransactions(ctx, nextID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTransactions", reflect.TypeOf((*MockTransactionStore)(nil).GetTransactions), ctx, nextID)
}

// CreateTransaction mocks base method
func (m *MockTransactionStore) CreateTransaction(ctx context.Context, req *transactions.Transaction) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTransaction", ctx, req)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateTransaction indicates an expected call of CreateTransaction
func (mr *MockTransactionStoreMockRecorder) CreateTransaction(ctx, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTransaction", reflect.TypeOf((*MockTransactionStore)(nil).CreateTransaction), ctx, req)
}
