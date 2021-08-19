package service

import (
	"os"
	"testing"

	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/xmlking/grpc-starter-kit/internal/config"
	"github.com/xmlking/grpc-starter-kit/internal/email"
)

type FakeEmailSender struct {
	mock.Mock
}

func (mock *FakeEmailSender) Send(subject, body string, to []string) error {
	args := mock.Called(subject, body, to)
	return args.Error(0)
}

func TestMain(m *testing.M) {
	// HINT: CertFile etc., Our schema has `file` path validation, which is relative to project root.
	if err := os.Chdir("../../.."); err != nil {
		log.Fatal().Err(err).Send()
	}
	wd, _ := os.Getwd()
	log.Debug().Msgf("Setup: changing working directory to: %s", wd)

	code := m.Run()
	os.Exit(code)
}

func TestEmailService_Welcome(t *testing.T) {
	emailer := &FakeEmailSender{}
	emailer.On("Send",
		"Welcome", "Hi Bob!", []string{"bob@smith.com"}).Return(nil)

	welcomer := NewEmailService(emailer)

	err := welcomer.Welcome("Bob", "bob@smith.com")
	assert.NoError(t, err)
	emailer.AssertExpectations(t)
}

func TestEmailService_Welcome_Integration(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping long integration test")
	}

	var (
		cfg = config.GetConfig()
	)

	emailer := email.NewSendEmail(cfg.Email)
	emailService := NewEmailService(emailer)

	err2 := emailService.Welcome("Welcome", "demo@gmail.com")
	if err2 != nil {
		t.Errorf("Send Welcome Email Failed: %v", err2)
	}
}

func TestEmailService_Welcome_E2E(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping e2e test")
	}
	t.Log("my first E2E test")
}
