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
    "google.golang.org/grpc/metadata"

    "github.com/xmlking/grpc-starter-kit/mkit/service/account/profile/v1"
    "github.com/xmlking/grpc-starter-kit/mkit/service/account/user/v1"
    "github.com/xmlking/grpc-starter-kit/shared/config"
    "github.com/xmlking/grpc-starter-kit/shared/constants"
    "github.com/xmlking/grpc-starter-kit/shared/util"
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
    suite.conn, err = config.GetClientConn(cfg.Services.Account)

    if err != nil {
        log.Fatal().Msgf("did not connect: %s", err)
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

    // Sending metadata - client side
    //md := metadata.Pairs("k1", "v1", "k1", "v2", "k2", "v3")
    //ctx := metadata.NewOutgoingContext(context.Background(), md)
    // create a new context with some metadata - (Optional) Just for demonstration
    ctx := metadata.AppendToOutgoingContext(context.Background(), constants.TraceIDKey, util.RandomString(8), constants.FromServiceKey, "e2e-account-test-client")
    _, err := suite.userClient.Create(ctx, &userv1.CreateRequest{
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

    ctx := metadata.AppendToOutgoingContext(context.Background(), constants.TraceIDKey, util.RandomString(8), constants.FromServiceKey, "e2e-account-test-client")
    rsp, err := suite.userClient.Exist(ctx, &userv1.ExistRequest{
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
