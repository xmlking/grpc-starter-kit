package repository_test

import (
	"database/sql"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	account_entities "github.com/xmlking/grpc-starter-kit/mkit/service/account/entities/v1"
	"github.com/xmlking/grpc-starter-kit/service/account/repository"
)

type accountRepositorySuite struct {
	suite.Suite
	gdb               *gorm.DB
	userRepository    repository.UserRepository
	profileRepository repository.ProfileRepository
	mock              sqlmock.Sqlmock
	user              *account_entities.UserORM
}

// SetupSuite
func (s *accountRepositorySuite) SetupSuite() {
	s.T().Log("in SetupSuite")
	var (
		db  *sql.DB
		err error
	)
	db, s.mock, err = sqlmock.New()
	require.NoError(s.T(), err)

	s.gdb, err = gorm.Open("SQLite3", db)
	require.NoError(s.T(), err)

	s.userRepository = repository.NewUserRepository(s.gdb)

	Username := "sumo"
	s.user = &account_entities.UserORM{
		Id:        uuid.NewV4(),
		FirstName: "sumo",
		LastName:  "demo",
		Email:     "sumo@demo.com",
		Username:  &Username,
	}
}

// TearDownSuite
func (s *accountRepositorySuite) TearDownSuite() {
	s.T().Log("in TearDownSuite")
	_ = s.gdb.Close()
}

// before each test
func (s *accountRepositorySuite) SetupTest() {
	s.T().Log("in SetupTest")
}

// after each test
func (s *accountRepositorySuite) TearDownTest() {
	s.T().Log("in TearDownTest")
	err := s.mock.ExpectationsWereMet() // make sure all expectations were met
	require.NoError(s.T(), err)
}

// User List test
func (s *accountRepositorySuite) TestUserRepository_List_Integration() {
	t := s.T()
	if testing.Short() {
		t.Skip("skipping long integration test")
	}

	t.Log("in TestUserRepository_Exist_Integration")

	countRows := sqlmock.NewRows([]string{"count"}).AddRow(0)

	s.mock.MatchExpectationsInOrder(false)

	s.mock.ExpectQuery("SELECT count").WillReturnRows(countRows)
	s.mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "users"`)).
		WillReturnError(sql.ErrNoRows)

	model := account_entities.UserORM{}
	total, users, err := s.userRepository.List(0, 1, "", &model)
	require.Error(s.T(), err, sql.ErrNoRows)
	s.Empty(total, 0)
	s.Empty(users, 0)
}

// User Create test
func (s *accountRepositorySuite) TestUserRepository_Create_Integration() {
	countRows := sqlmock.NewRows([]string{"count"}).AddRow(0)

	s.mock.MatchExpectationsInOrder(true)

	s.mock.ExpectQuery("SELECT count").WillReturnError(sql.ErrNoRows)
	s.mock.ExpectQuery("SELECT count").WillReturnError(sql.ErrNoRows)
	s.mock.ExpectQuery("SELECT count").WillReturnRows(countRows)
	s.mock.ExpectBegin()
	s.mock.ExpectExec(`NSERT INTO "users"`).
		WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), s.user.Email, s.user.FirstName, sqlmock.AnyArg(), s.user.LastName, sqlmock.AnyArg(), s.user.Username).
		WillReturnResult(sqlmock.NewResult(1, 1))
	s.mock.ExpectCommit()

	err := s.userRepository.Create(s.user)
	require.NoError(s.T(), err)

}

func TestExampleTestSuite(t *testing.T) {
	suite.Run(t, new(accountRepositorySuite))
}
