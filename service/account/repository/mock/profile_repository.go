package mock

import (
	"context"

	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"

	"github.com/xmlking/grpc-starter-kit/ent"
	"github.com/xmlking/grpc-starter-kit/service/account/repository"
)

// profileRepository struct
type profileRepository struct {
	mock.Mock
}

//NewProfileRepository returns an instance of `ProfileRepository`.
func NewProfileRepository() repository.ProfileRepository {
	return &profileRepository{}
}

// Exist
func (repo *profileRepository) Exist(ctx context.Context, model *ent.Profile) (bool, error) {
	//args := repo.Called(ctx, model)
	//return args.Bool(0), args.Error(1)

	ret := repo.Called(ctx, model)

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context, *ent.Profile) bool); ok {
		r0 = rf(ctx, model)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Bool(0)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *ent.Profile) error); ok {
		r1 = rf(ctx, model)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// List, Cursor-based pagination
func (repo *profileRepository) List(ctx context.Context, limit, page int, sort string, model *ent.Profile) (total int, users []*ent.Profile, err error) {
	ret := repo.Called(ctx, limit, page, sort, model)

	var r0 int
	if rf, ok := ret.Get(0).(func(context.Context, int, int, string, *ent.Profile) int); ok {
		r0 = rf(ctx, limit, page, sort, model)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Int(0)
		}
	}

	var r1 []*ent.Profile
	if rf, ok := ret.Get(1).(func(context.Context, int, int, string, *ent.Profile) []*ent.Profile); ok {
		r1 = rf(ctx, limit, page, sort, model)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).([]*ent.Profile)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, int, int, string, *ent.Profile) error); ok {
		r2 = rf(ctx, limit, page, sort, model)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// Find by ID
func (repo *profileRepository) Get(ctx context.Context, id uuid.UUID) (*ent.Profile, error) {
	ret := repo.Called(ctx, id)

	var r0 *ent.Profile
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) *ent.Profile); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ent.Profile)
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

// Find by UserID
func (repo *profileRepository) GetByUserID(ctx context.Context, userId uuid.UUID) (*ent.Profile, error) {
	ret := repo.Called(ctx, userId)

	var r0 *ent.Profile
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) *ent.Profile); ok {
		r0 = rf(ctx, userId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ent.Profile)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uuid.UUID) error); ok {
		r1 = rf(ctx, userId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (repo *profileRepository) Create(ctx context.Context, model *ent.Profile) (*ent.Profile, error) {
	ret := repo.Called(ctx, model)

	var r0 *ent.Profile
	if rf, ok := ret.Get(0).(func(context.Context, *ent.Profile) *ent.Profile); ok {
		r0 = rf(ctx, model)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ent.Profile)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *ent.Profile) error); ok {
		r1 = rf(ctx, model)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update
func (repo *profileRepository) Update(ctx context.Context, in *ent.Profile) (out *ent.Profile, err error) {
	ret := repo.Called(ctx, in)

	var r0 *ent.Profile
	if rf, ok := ret.Get(0).(func(context.Context, *ent.Profile) *ent.Profile); ok {
		r0 = rf(ctx, in)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ent.Profile)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *ent.Profile) error); ok {
		r1 = rf(ctx, in)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete, Soft Delete
func (repo *profileRepository) Delete(ctx context.Context, id uuid.UUID) (err error) {
	ret := repo.Called(ctx, id)

	var r0 *ent.Profile
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) *ent.Profile); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ent.Profile)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uuid.UUID) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	print(r0)
	return r1
}

// Delete, Soft Delete
func (repo *profileRepository) DeleteFull(ctx context.Context, model *ent.Profile) (err error) {
	ret := repo.Called(ctx, model)

	var r0 *ent.Profile
	if rf, ok := ret.Get(0).(func(context.Context, *ent.Profile) *ent.Profile); ok {
		r0 = rf(ctx, model)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ent.Profile)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *ent.Profile) error); ok {
		r1 = rf(ctx, model)
	} else {
		r1 = ret.Error(1)
	}

	print(r0)
	return r1
}

// Count
func (repo *profileRepository) Count(ctx context.Context) (int, error) {
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
