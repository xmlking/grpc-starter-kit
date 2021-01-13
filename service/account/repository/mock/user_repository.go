package mock

import (
	"context"

	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"

	"github.com/xmlking/grpc-starter-kit/ent"
	"github.com/xmlking/grpc-starter-kit/service/account/repository"
)

// userRepository struct
type userRepository struct {
	mock.Mock
}

// NewUserRepository returns an instance of `UserRepository`.
func NewUserRepository() repository.UserRepository {
	return &userRepository{}
}

// Exist
func (repo *userRepository) Exist(ctx context.Context, model *ent.User) (bool, error) {
	//args := repo.Called(ctx, model)
	//return args.Bool(0), args.Error(1)

	ret := repo.Called(ctx, model)

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context, *ent.User) bool); ok {
		r0 = rf(ctx, model)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Bool(0)
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

// List, Cursor-based pagination
func (repo *userRepository) List(ctx context.Context, limit, page int, sort string, model *ent.User) (total int, users []*ent.User, err error) {
	ret := repo.Called(ctx, limit, page, sort, model)

	var r0 int
	if rf, ok := ret.Get(0).(func(context.Context, int, int, string, *ent.User) int); ok {
		r0 = rf(ctx, limit, page, sort, model)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Int(0)
		}
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

// Find by ID
func (repo *userRepository) Get(ctx context.Context, id uuid.UUID) (*ent.User, error) {
	ret := repo.Called(ctx, id)

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

func (repo *userRepository) Create(ctx context.Context, model *ent.User) (*ent.User, error) {
	ret := repo.Called(ctx, model)

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

// Update
func (repo *userRepository) Update(ctx context.Context, in *ent.User) (out *ent.User, err error) {
	ret := repo.Called(ctx, in)

	var r0 *ent.User
	if rf, ok := ret.Get(0).(func(context.Context, *ent.User) *ent.User); ok {
		r0 = rf(ctx, in)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ent.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *ent.User) error); ok {
		r1 = rf(ctx, in)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete, Soft Delete
func (repo *userRepository) Delete(ctx context.Context, id uuid.UUID) (*ent.User, error) {
	ret := repo.Called(ctx, id)

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

// Delete, Soft Delete
func (repo *userRepository) DeleteFull(ctx context.Context, model *ent.User) (*ent.User, error) {
	ret := repo.Called(ctx, model)

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

// Count
func (repo *userRepository) Count(ctx context.Context) (int, error) {
	ret := repo.Called(ctx)

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
