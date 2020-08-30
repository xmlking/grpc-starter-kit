package repository

import (
	"context"
	"time"

	"github.com/cockroachdb/errors"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"

	"github.com/xmlking/grpc-starter-kit/ent"
	"github.com/xmlking/grpc-starter-kit/ent/predicate"
	"github.com/xmlking/grpc-starter-kit/ent/user"
)

// https://github.com/mazti/restless/tree/master/base
// https://github.com/WaranchitPk/funny_todo_list/tree/master/api/v1/tasks

// UserRepository interface
type UserRepository interface {
	Exist(ctx context.Context, model *ent.User) (bool, error)
	List(ctx context.Context, limit, page int, sort string, model *ent.User) (total int, users []*ent.User, err error)
	Get(ctx context.Context, id uuid.UUID) (*ent.User, error)
	Create(ctx context.Context, model *ent.User) (*ent.User, error)
	Update(ctx context.Context, model *ent.User) (*ent.User, error)
	DeleteFull(ctx context.Context, model *ent.User) (*ent.User, error)
	Delete(ctx context.Context, id uuid.UUID) (*ent.User, error)
	Count(ctx context.Context) (int, error)
}

// userRepository struct
type userRepository struct {
	dbClient *ent.Client
}

// NewUserRepository returns an instance of `UserRepository`.
func NewUserRepository(dbClient *ent.Client) UserRepository {
	return &userRepository{
		dbClient: dbClient,
	}
}

// Exist
func (repo *userRepository) Exist(ctx context.Context, model *ent.User) (bool, error) {
	log.Info().Msgf("Received userRepository.Exist request %v", model)

	var pd []predicate.User
	if model.ID != uuid.Nil {
		pd = append(pd, user.ID(model.ID))
	}
	if model.Username != "" {
		pd = append(pd, user.Username(model.Username))
	}
	if model.Email != "" {
		pd = append(pd, user.Email(model.Email))
	}

	return repo.dbClient.User.
		Query().
		Where(user.DeleteTimeIsNil()). // query only active
		Where(user.Or(pd...)).
		Exist(ctx)
}

// List, Cursor-based pagination
func (repo *userRepository) List(ctx context.Context, limit, page int, sort string, model *ent.User) (total int, users []*ent.User, err error) {
	// Set defaults
	if limit == 0 {
		limit = 10
	}
	var offset int
	if page > 1 {
		offset = (page - 1) * limit
	} else {
		offset = 0
	}
	if sort == "" {
		sort = user.FieldCreateTime
	}

	// query only active
	var pd = []predicate.User{
		user.DeleteTimeIsNil(),
	}
	if model.Username != "" {
		pd = append(pd, user.UsernameContains(model.Username))
	}
	if model.FirstName != "" {
		pd = append(pd, user.FirstNameContains(model.FirstName))
	}
	if model.LastName != "" {
		pd = append(pd, user.LastNameContains(model.LastName))
	}
	if model.Email != "" {
		pd = append(pd, user.EmailContains(model.Email))
	}

	total, err = repo.dbClient.User.
		Query().
		Where(pd...).
		Count(ctx)
	if err != nil {
		return
	}

	users, err = repo.dbClient.User.
		Query().
		Where(pd...).
		Limit(limit).
		Offset(offset).
		Order(ent.Desc(sort)).
		WithProfile().
		All(ctx)
	return
}

// Find by ID
func (repo *userRepository) Get(ctx context.Context, id uuid.UUID) (*ent.User, error) {
	return repo.dbClient.User.
		Query().
		WithProfile().
		Where(user.ID(id), user.DeleteTimeIsNil()).
		Only(ctx)
}

// Create
func (repo *userRepository) Create(ctx context.Context, model *ent.User) (*ent.User, error) {
	if exist, err := repo.Exist(ctx, model); err != nil {
		return nil, err
	} else if exist {
		return nil, errors.New("user already exist")
	}

	return repo.dbClient.User.
		Create().
		SetEmail(model.Email).
		SetUsername(model.Username).
		SetFirstName(model.FirstName).
		SetLastName(model.LastName).
		SetTenant(model.Tenant).
		Save(ctx)
}

func rollback(tx *ent.Tx, err error) error {
	if rerr := tx.Rollback(); rerr != nil {
		err = errors.Wrapf(err, "rolling back transaction: %v", rerr)
	}
	return err
}

// Update
func (repo *userRepository) Update(ctx context.Context, in *ent.User) (out *ent.User, err error) {
	var tx *ent.Tx
	tx, err = repo.dbClient.Tx(ctx)
	if err != nil {
		return nil, err
	}

	// Be specific
	//updateOp := tx.User.UpdateOneID(in.ID)
	//if in.FirstName != "" {
	//    updateOp = updateOp.SetFirstName(in.FirstName)
	//}
	//if in.LastName != "" {
	//    updateOp = updateOp.SetLastName(in.LastName)
	//}
	//if in.Email != "" {
	//    updateOp = updateOp.SetEmail(in.Email)
	//}

	updateOp := tx.User.UpdateOne(in)

	out, err = updateOp.Save(ctx)

	if err != nil {
		err = rollback(tx, err)
		return
	}

	//if err != nil {
	//   if ent.IsNotFound(err) {
	//       return nil, http.CreateNotFoundError()
	//   }
	//   if ent.IsConstraintError(err) {
	//       return nil, http.CreateBadRequestError("constraint error")
	//   }
	//   return nil, err
	//}
	//
	//if result == nil {
	//   return nil, http.CreateInternalError()
	//}

	// Commit the transaction.
	err = tx.Commit()
	return
}

// Delete, Soft Delete
func (repo *userRepository) Delete(ctx context.Context, id uuid.UUID) (*ent.User, error) {
	//return repo.dbClient.User.
	//   DeleteOneID(id).
	//   Exec(ctx)

	return repo.dbClient.User.
		UpdateOneID(id).
		SetDeleteTime(time.Now()).
		Save(ctx)

}

// Delete, Soft Delete
func (repo *userRepository) DeleteFull(ctx context.Context, model *ent.User) (*ent.User, error) {
	// here we return Update statement, that is extends delete with all the WHERE condtions, but providing a SET instruction to update single field
	return repo.dbClient.User.
		UpdateOne(model).
		SetDeleteTime(time.Now()).
		Save(ctx)
}

// Count
func (repo *userRepository) Count(ctx context.Context) (int, error) {
	return repo.dbClient.User.
		Query().
		Where(user.DeleteTimeIsNil()).
		Count(ctx)
}
