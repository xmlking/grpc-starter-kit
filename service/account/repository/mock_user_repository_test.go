// Code generated by mockery v2.10.6. DO NOT EDIT.

package repository

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
	ent "github.com/xmlking/grpc-starter-kit/ent"

	uuid "github.com/google/uuid"
)

// MockUserRepository is an autogenerated mock type for the UserRepository type
type MockUserRepository struct {
	mock.Mock
}

// Count provides a mock function with given fields: ctx
func (_m *MockUserRepository) Count(ctx context.Context) (int, error) {
	ret := _m.Called(ctx)

	var r0 int
	if rf, ok := ret.Get(0).(func(context.Context) int); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Create provides a mock function with given fields: ctx, model
func (_m *MockUserRepository) Create(ctx context.Context, model *ent.User) (*ent.User, error) {
	ret := _m.Called(ctx, model)

	var r0 *ent.User
	if rf, ok := ret.Get(0).(func(context.Context, *ent.User) *ent.User); ok {
		r0 = rf(ctx, model)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ent.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *ent.User) error); ok {
		r1 = rf(ctx, model)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: ctx, id
func (_m *MockUserRepository) Delete(ctx context.Context, id uuid.UUID) (*ent.User, error) {
	ret := _m.Called(ctx, id)

	var r0 *ent.User
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) *ent.User); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ent.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uuid.UUID) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteFull provides a mock function with given fields: ctx, model
func (_m *MockUserRepository) DeleteFull(ctx context.Context, model *ent.User) (*ent.User, error) {
	ret := _m.Called(ctx, model)

	var r0 *ent.User
	if rf, ok := ret.Get(0).(func(context.Context, *ent.User) *ent.User); ok {
		r0 = rf(ctx, model)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ent.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *ent.User) error); ok {
		r1 = rf(ctx, model)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Exist provides a mock function with given fields: ctx, model
func (_m *MockUserRepository) Exist(ctx context.Context, model *ent.User) (bool, error) {
	ret := _m.Called(ctx, model)

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context, *ent.User) bool); ok {
		r0 = rf(ctx, model)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *ent.User) error); ok {
		r1 = rf(ctx, model)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Get provides a mock function with given fields: ctx, id
func (_m *MockUserRepository) Get(ctx context.Context, id uuid.UUID) (*ent.User, error) {
	ret := _m.Called(ctx, id)

	var r0 *ent.User
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) *ent.User); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ent.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uuid.UUID) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// List provides a mock function with given fields: ctx, limit, page, sort, model
func (_m *MockUserRepository) List(ctx context.Context, limit int, page int, sort string, model *ent.User) (int, []*ent.User, error) {
	ret := _m.Called(ctx, limit, page, sort, model)

	var r0 int
	if rf, ok := ret.Get(0).(func(context.Context, int, int, string, *ent.User) int); ok {
		r0 = rf(ctx, limit, page, sort, model)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 []*ent.User
	if rf, ok := ret.Get(1).(func(context.Context, int, int, string, *ent.User) []*ent.User); ok {
		r1 = rf(ctx, limit, page, sort, model)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).([]*ent.User)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, int, int, string, *ent.User) error); ok {
		r2 = rf(ctx, limit, page, sort, model)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// Update provides a mock function with given fields: ctx, model
func (_m *MockUserRepository) Update(ctx context.Context, model *ent.User) (*ent.User, error) {
	ret := _m.Called(ctx, model)

	var r0 *ent.User
	if rf, ok := ret.Get(0).(func(context.Context, *ent.User) *ent.User); ok {
		r0 = rf(ctx, model)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ent.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *ent.User) error); ok {
		r1 = rf(ctx, model)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
