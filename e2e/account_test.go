// e2e, black-box testing
package e2e

import (
	"context"
	"testing"

	"github.com/golang/protobuf/ptypes/wrappers"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	"github.com/xmlking/grpc-starter-kit/micro/middleware/rpclog"
	"github.com/xmlking/grpc-starter-kit/mkit/service/account/profile/v1"
	"github.com/xmlking/grpc-starter-kit/mkit/service/account/user/v1"
	"github.com/xmlking/grpc-starter-kit/shared/config"
)

// Define the suite, and absorb the built-in basic suite
// functionality from testify - including a T() method which
// returns the current testing context
type AccountTestSuite struct {
	suite.Suite
	conn          *grpc.ClientConn
	userClient    userv1.UserServiceClient
	profileClient profilev1.ProfileServiceClient
}

// SetupSuite implements suite.SetupAllSuite
func (suite *AccountTestSuite) SetupSuite() {
	cfg := config.GetConfig()
	suite.T().Log("in SetupSuite")

	var err error
	suite.conn, err = grpc.Dial(
		cfg.Services.Account.Endpoint, grpc.WithInsecure(),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy":"round_robin"}`),
		grpc.WithUnaryInterceptor(grpc_middleware.ChainUnaryClient(
			rpclog.UnaryClientInterceptor(),
		)),
	)
	if err != nil {
		log.Fatal().Msgf("did not connect: %s", err)
	}

	println(suite.conn.Target())
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
	ctx := metadata.AppendToOutgoingContext(context.Background(), "X-User-Id", "john", "X-From-Id", "script")

	_, err := suite.userClient.Create(ctx, &userv1.CreateRequest{
		Username:  &wrappers.StringValue{Value: "sumo"},
		FirstName: &wrappers.StringValue{Value: "sumo"},
		LastName:  &wrappers.StringValue{Value: "demo"},
		Email:     &wrappers.StringValue{Value: "sumo@demo.com"},
	})
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

	rsp, err := suite.userClient.Exist(context.TODO(), &userv1.ExistRequest{
		Username: &wrappers.StringValue{Value: "sumo"},
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
