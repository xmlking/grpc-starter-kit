// e2e, black-box testing
package e2e

import (
	"context"
	"fmt"
	"testing"

	"github.com/golang/protobuf/ptypes/wrappers"
	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"google.golang.org/grpc"

	appendTags "github.com/xmlking/toolkit/middleware/tags/append"

	"github.com/xmlking/toolkit/util"

	"github.com/xmlking/grpc-starter-kit/mkit/service/account/profile/v1"
	"github.com/xmlking/grpc-starter-kit/mkit/service/account/user/v1"
	"github.com/xmlking/grpc-starter-kit/shared/config"
	"github.com/xmlking/grpc-starter-kit/shared/constants"
)

// Define the suite, and absorb the built-in basic suite
// functionality from testify - including a T() method which
// returns the current testing context
type AccountTestSuite struct {
	suite.Suite
	suffix        string
	conn          *grpc.ClientConn
	userClient    userv1.UserServiceClient
	profileClient profilev1.ProfileServiceClient
}

// SetupSuite implements suite.SetupAllSuite
func (suite *AccountTestSuite) SetupSuite() {
	cfg := config.GetConfig()
	suite.T().Log("in SetupSuite")

	suite.suffix = util.RandomStringLower(5)

	var err error
	var ucInterceptors = []grpc.UnaryClientInterceptor{
		appendTags.UnaryClientInterceptor(appendTags.WithTraceID(), appendTags.WithPairs(constants.FromServiceKey, constants.ACCOUNT_CLIENT)),
	}
	suite.conn, err = config.GetClientConn(cfg.Services.Account, ucInterceptors)
	if err != nil {
		log.Fatal().Err(err).Msgf("Failed connect to: %s", cfg.Services.Account.Endpoint)
	}

	suite.userClient = userv1.NewUserServiceClient(suite.conn)
	suite.profileClient = profilev1.NewProfileServiceClient(suite.conn)
}

// TearDownSuite implements suite.TearDownAllSuite
func (suite *AccountTestSuite) TearDownSuite() {
	suite.T().Log("in TearDownSuite")
	_ = suite.conn.Close()
}

// before each test
func (suite *AccountTestSuite) SetupTest() {
	t := suite.T()
	t.Log("in SetupTest - creating user")

	_, err := suite.userClient.Create(context.Background(), &userv1.CreateRequest{
		Username:  &wrappers.StringValue{Value: fmt.Sprintf("u_%s", suite.suffix)},
		FirstName: &wrappers.StringValue{Value: fmt.Sprintf("f_%s", suite.suffix)},
		LastName:  &wrappers.StringValue{Value: fmt.Sprintf("l_%s", suite.suffix)},
		Email:     &wrappers.StringValue{Value: fmt.Sprintf("e_%s@demo.com", suite.suffix)},
	})
	if err != nil {
		log.Error().Err(err).Send()
	}
	require.Nil(t, err)
}

// after each test
func (suite *AccountTestSuite) TearDownTest() {
	suite.T().Log("in TearDownTest")
}

// All methods that begin with "Test" are run as tests within a suite.
func (suite *AccountTestSuite) TestUserHandler_Exist_E2E() {
	t := suite.T()
	t.Log("in TestUserHandler_Exist_E2E, checking if user Exist")

	rsp, err := suite.userClient.Exist(context.Background(), &userv1.ExistRequest{
		Username: &wrappers.StringValue{Value: fmt.Sprintf("u_%s", suite.suffix)},
	})
	require.Nil(t, err)
	assert.Equal(suite.T(), rsp.GetResult(), true)
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestAccountTestSuite(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping e2e test")
	}
	suite.Run(t, new(AccountTestSuite))
}
