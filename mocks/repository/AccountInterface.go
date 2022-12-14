// Code generated by mockery v2.14.1. DO NOT EDIT.

package repository

import (
	entity "pismo/entity"

	mock "github.com/stretchr/testify/mock"
)

// AccountInterface is an autogenerated mock type for the AccountInterface type
type AccountInterface struct {
	mock.Mock
}

// Find provides a mock function with given fields: accountID
func (_m *AccountInterface) Find(accountID uint64) (entity.Account, error) {
	ret := _m.Called(accountID)

	var r0 entity.Account
	if rf, ok := ret.Get(0).(func(uint64) entity.Account); ok {
		r0 = rf(accountID)
	} else {
		r0 = ret.Get(0).(entity.Account)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint64) error); ok {
		r1 = rf(accountID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Insert provides a mock function with given fields: account
func (_m *AccountInterface) Insert(account entity.Account) (entity.Account, error) {
	ret := _m.Called(account)

	var r0 entity.Account
	if rf, ok := ret.Get(0).(func(entity.Account) entity.Account); ok {
		r0 = rf(account)
	} else {
		r0 = ret.Get(0).(entity.Account)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(entity.Account) error); ok {
		r1 = rf(account)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateBalance provides a mock function with given fields: accountID, newBalance
func (_m *AccountInterface) UpdateBalance(accountID uint64, newBalance float64) error {
	ret := _m.Called(accountID, newBalance)

	var r0 error
	if rf, ok := ret.Get(0).(func(uint64, float64) error); ok {
		r0 = rf(accountID, newBalance)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewAccountInterface interface {
	mock.TestingT
	Cleanup(func())
}

// NewAccountInterface creates a new instance of AccountInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewAccountInterface(t mockConstructorTestingTNewAccountInterface) *AccountInterface {
	mock := &AccountInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
