// Code generated by mockery v1.0.0. DO NOT EDIT.

package storage

import accounts "github.com/openbank/gunk/gunk/v1/accounts"
import context "context"
import mock "github.com/stretchr/testify/mock"

// MockAccountStore is an autogenerated mock type for the AccountStore type
type MockAccountStore struct {
	mock.Mock
}

// CreateAccount provides a mock function with given fields: ctx, account
func (_m *MockAccountStore) CreateAccount(ctx context.Context, account *accounts.Account) error {
	ret := _m.Called(ctx, account)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *accounts.Account) error); ok {
		r0 = rf(ctx, account)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteAccount provides a mock function with given fields: ctx, id
func (_m *MockAccountStore) DeleteAccount(ctx context.Context, id string) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAccount provides a mock function with given fields: ctx, id
func (_m *MockAccountStore) GetAccount(ctx context.Context, id string) (*accounts.Account, error) {
	ret := _m.Called(ctx, id)

	var r0 *accounts.Account
	if rf, ok := ret.Get(0).(func(context.Context, string) *accounts.Account); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*accounts.Account)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAccounts provides a mock function with given fields: ctx, nextID
func (_m *MockAccountStore) GetAccounts(ctx context.Context, nextID string) ([]*accounts.Account, bool, error) {
	ret := _m.Called(ctx, nextID)

	var r0 []*accounts.Account
	if rf, ok := ret.Get(0).(func(context.Context, string) []*accounts.Account); ok {
		r0 = rf(ctx, nextID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*accounts.Account)
		}
	}

	var r1 bool
	if rf, ok := ret.Get(1).(func(context.Context, string) bool); ok {
		r1 = rf(ctx, nextID)
	} else {
		r1 = ret.Get(1).(bool)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, string) error); ok {
		r2 = rf(ctx, nextID)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// UpdateAccount provides a mock function with given fields: ctx, account
func (_m *MockAccountStore) UpdateAccount(ctx context.Context, account *accounts.Account) error {
	ret := _m.Called(ctx, account)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *accounts.Account) error); ok {
		r0 = rf(ctx, account)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}