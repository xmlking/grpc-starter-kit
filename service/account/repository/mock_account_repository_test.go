package repository_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"github.com/xmlking/grpc-starter-kit/ent"
	"github.com/xmlking/grpc-starter-kit/service/account/repository"
)

type mockAccountRepositorySuite struct {
	suite.Suite
	userRepository repository.UserRepository
	user           *ent.User
}

func (s *mockAccountRepositorySuite) SetupSuite() {
	s.user = &ent.User{
		FirstName: "sumo",
		LastName:  "demo",
		Email:     "sumo@demo.com",
		Username:  "sumo",
		Tenant:    "ABC Corp",
	}

	mockUserRepository := new(repository.MockUserRepository)
	// setup expectations
	mockUserRepository.On("Exist", context.Background(), s.user).Return(true, nil)

	s.userRepository = mockUserRepository
}

// User Exist test
func (s *mockAccountRepositorySuite) TestUserRepository_Exist_Integration() {
	t := s.T()

	exists, err := s.userRepository.Exist(context.Background(), s.user)
	require.Nil(t, err)
	assert.Equal(t, exists, true)
}

func TestMockAccountRepositorySuite(t *testing.T) {
	suite.Run(t, new(mockAccountRepositorySuite))
}
