package repository

import (
	"context"

	"github.com/cockroachdb/errors"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"

	"github.com/xmlking/grpc-starter-kit/ent"
	"github.com/xmlking/grpc-starter-kit/ent/predicate"
	"github.com/xmlking/grpc-starter-kit/ent/profile"
	"github.com/xmlking/grpc-starter-kit/ent/user"
)

// ProfileRepository interface
//go:generate mockery --name=ProfileRepository --case=snake --tags="mock" --inpackage --testonly
type ProfileRepository interface {
	Exist(ctx context.Context, model *ent.Profile) (bool, error)
	List(ctx context.Context, limit, page int, sort string, model *ent.Profile) (total int, profiles []*ent.Profile, err error)
	Get(ctx context.Context, id uuid.UUID) (*ent.Profile, error)
	GetByUserID(ctx context.Context, userID uuid.UUID) (*ent.Profile, error)
	Create(ctx context.Context, model *ent.Profile) (*ent.Profile, error)
}

// profileRepository struct
type profileRepository struct {
	dbClient *ent.Client
}

// NewProfileRepository returns an instance of `ProfileRepository`.
func NewProfileRepository(dbClient *ent.Client) ProfileRepository {
	return &profileRepository{
		dbClient: dbClient,
	}
}

// Exist
func (repo *profileRepository) Exist(ctx context.Context, model *ent.Profile) (bool, error) {
	log.Info().Msgf("Received profileRepository.Exist request %v", model)

	var pd []predicate.Profile
	if model.Edges.User != nil && model.Edges.User.ID != uuid.Nil { // TODO need UserId
		pd = append(pd, profile.HasUserWith(user.ID(model.Edges.User.ID)))
	}

	return repo.dbClient.Profile.
		Query().
		Where(profile.DeleteTimeIsNil()). // query only active
		Where(pd...).
		Exist(ctx)
}

// List, Cursor-based pagination
func (repo *profileRepository) List(ctx context.Context, limit, page int, sort string, model *ent.Profile) (total int, users []*ent.Profile, err error) {
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
		sort = profile.FieldCreateTime
	}

	// query only active
	var pd = []predicate.Profile{
		profile.DeleteTimeIsNil(),
	}
	if model.Edges.User != nil && model.Edges.User.ID != uuid.Nil { // TODO
		pd = append(pd, profile.HasUserWith(user.ID(model.Edges.User.ID)))

	}
	if model.PreferredTheme != "" {
		pd = append(pd, profile.PreferredTheme(model.PreferredTheme))
	}
	if model.Gender != "" {
		pd = append(pd, profile.GenderEQ(model.Gender))
	}

	total, err = repo.dbClient.Profile.
		Query().
		Where(pd...).
		Count(ctx)
	if err != nil {
		return
	}

	users, err = repo.dbClient.Profile.
		Query().
		Where(pd...).
		Limit(limit).
		Offset(offset).
		Order(ent.Desc(sort)).
		All(ctx)
	return
}

// Find by ID
func (repo *profileRepository) Get(ctx context.Context, id uuid.UUID) (*ent.Profile, error) {
	return repo.dbClient.Profile.
		Query().
		Where(profile.ID(id), profile.DeleteTimeIsNil()).
		Only(ctx)
}

// Find by UserID
func (repo *profileRepository) GetByUserID(ctx context.Context, userID uuid.UUID) (*ent.Profile, error) {
	return repo.dbClient.Profile.
		Query().
		Where(profile.HasUserWith(user.ID(userID))).
		Only(ctx)
}

// Create
func (repo *profileRepository) Create(ctx context.Context, model *ent.Profile) (*ent.Profile, error) {
	if exist, err := repo.Exist(ctx, model); err != nil {
		return nil, err
	} else if exist {
		return nil, errors.New("user already exist")
	}

	return repo.dbClient.Profile.
		Create().
		SetAge(model.Age).
		SetAvatar(model.Avatar).
		SetBirthday(model.Birthday).
		SetGender(model.Gender).
		SetTz(model.Tz).
		SetPreferredTheme(model.PreferredTheme).
		SetUserID(model.ID). // TODO
		Save(ctx)
}
