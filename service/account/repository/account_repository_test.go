package repository_test

import (
	"context"
	"database/sql"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/facebook/ent/dialect"
	entsql "github.com/facebook/ent/dialect/sql"
	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/xmlking/grpc-starter-kit/ent"
	"github.com/xmlking/grpc-starter-kit/service/account/repository"
	"github.com/xmlking/grpc-starter-kit/shared/config"
	_ "github.com/xmlking/grpc-starter-kit/shared/logger"
)

// https://github.com/WaranchitPk/funny_todo_list/blob/master/api/v1/tasks/repository/task_repository_test.go
type accountRepositorySuite struct {
	suite.Suite
	dbClinet       *ent.Client
	userRepository repository.UserRepository
	mock           sqlmock.Sqlmock
	user           *ent.User
}

// SetupSuite
func (s *accountRepositorySuite) SetupSuite() {
	dbConf := config.GetConfig().Database
	s.T().Logf("in SetupSuite: %v", dbConf)

	var (
		db  *sql.DB
		err error
	)
	db, s.mock, err = sqlmock.New()
	require.NoError(s.T(), err)
	drv := entsql.OpenDB(dialect.SQLite, db)

	// s.dbClinet = enttest.NewClient(s.T(), enttest.WithOptions(ent.Driver(drv), ent.Debug(), ent.Log(log.Print)))
	s.dbClinet = ent.NewClient(ent.Driver(drv), ent.Debug(), ent.Log(log.Print))

	require.NoError(s.T(), err)

	s.userRepository = repository.NewUserRepository(s.dbClinet)

	s.user = &ent.User{
		FirstName: "sumo",
		LastName:  "demo",
		Email:     "sumo@demo.com",
		Username:  "sumo",
		Tenant:    "ABC Corp",
	}
}

// TearDownSuite
func (s *accountRepositorySuite) TearDownSuite() {
	s.T().Log("in TearDownSuite")
	_ = s.dbClinet.Close()
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

	s.mock.ExpectQuery("SELECT COUNT").WillReturnRows(countRows)
	s.mock.ExpectQuery(regexp.QuoteMeta("SELECT DISTINCT `users`.`id`")).WillReturnError(sql.ErrNoRows)

	model := ent.User{}
	total, users, err := s.userRepository.List(context.Background(), 0, 1, "", &model)
	require.Error(s.T(), err, sql.ErrNoRows)
	s.Empty(total, 0)
	s.Empty(users, 0)
}

// User Create test
func (s *accountRepositorySuite) TestUserRepository_Create_Integration() {
	countRows := sqlmock.NewRows([]string{"count"}).AddRow(0)

	s.mock.MatchExpectationsInOrder(true)

	s.mock.ExpectQuery("SELECT COUNT").WillReturnRows(countRows)
	s.mock.ExpectBegin()
	s.mock.ExpectExec("INSERT INTO `users`").
		WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), s.user.Username, s.user.FirstName, s.user.LastName, s.user.Email, s.user.Tenant, sqlmock.AnyArg()).
		WillReturnResult(sqlmock.NewResult(1, 1))
	s.mock.ExpectCommit()

	usr, err := s.userRepository.Create(context.Background(), s.user)
	s.T().Log(usr)
	require.NoError(s.T(), err)

}

func TestExampleTestSuite(t *testing.T) {
	suite.Run(t, new(accountRepositorySuite))
}
