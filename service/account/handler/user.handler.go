package handler

import (
	"context"
	"fmt"
	"time"

	cloudevents "github.com/cloudevents/sdk-go/v2"
	"github.com/jinzhu/gorm"
	"github.com/rs/zerolog/log"
	uuid "github.com/satori/go.uuid"
	"github.com/thoas/go-funk"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	account_entities "github.com/xmlking/grpc-starter-kit/mkit/service/account/entities/v1"
	"github.com/xmlking/grpc-starter-kit/mkit/service/account/user/v1"
	"github.com/xmlking/grpc-starter-kit/mkit/service/emailer/v1"
	"github.com/xmlking/grpc-starter-kit/mkit/service/greeter/v1"
	"github.com/xmlking/grpc-starter-kit/service/account/repository"
	"github.com/xmlking/grpc-starter-kit/shared/auth"
)

// UserHandler struct
type userHandler struct {
	userRepository   repository.UserRepository
	Event            cloudevents.Client
	greeterSrvClient greeterv1.GreeterServiceClient
}

// NewUserHandler returns an instance of `UserServiceHandler`.
func NewUserHandler(repo repository.UserRepository, emailerClient cloudevents.Client, greeterClient greeterv1.GreeterServiceClient) userv1.UserServiceServer {
	return &userHandler{
		userRepository:   repo,
		Event:            emailerClient,
		greeterSrvClient: greeterClient,
	}
}

func (h *userHandler) Exist(ctx context.Context, req *userv1.ExistRequest) (rsp *userv1.ExistResponse, err error) {
	log.Info().Msg("Received UserHandler.Exist request")
	model := account_entities.UserORM{}
	model.Id = uuid.FromStringOrNil(req.Id.GetValue())
	username := req.Username.GetValue()
	model.Username = &username
	model.Email = req.Email.GetValue()

	exists := h.userRepository.Exist(&model)
	log.Info().Msgf("user exists? %t", exists)
	return &userv1.ExistResponse{Result: exists}, nil
}

func (h *userHandler) List(ctx context.Context, req *userv1.ListRequest) (rsp *userv1.ListResponse, err error) {
	log.Info().Msg("Received UserHandler.List request")
	model := account_entities.UserORM{}
	username := req.Username.GetValue()
	model.Username = &username
	model.FirstName = req.FirstName.GetValue()
	model.LastName = req.LastName.GetValue()
	model.Email = req.Email.GetValue()

	total, users, err := h.userRepository.List(req.Limit.GetValue(), req.Page.GetValue(), req.Sort.GetValue(), &model)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, fmt.Sprintf("Error %v", err.Error()))
	}
	rsp = &userv1.ListResponse{Total: total}

	// newUsers := make([]*accountPB.User, len(users))
	// for index, user := range users {
	// 	tmpUser, _ := user.ToPB(ctx)
	// 	newUsers[index] = &tmpUser
	// 	// *newUsers[index], _ = user.ToPB(ctx) ???
	// }
	newUsers := funk.Map(users, func(user *account_entities.UserORM) *account_entities.User {
		tmpUser, _ := user.ToPB(ctx)
		return &tmpUser
	}).([]*account_entities.User)

	rsp.Results = newUsers
	return
}

func (h *userHandler) Get(ctx context.Context, req *userv1.GetRequest) (rsp *userv1.GetResponse, err error) {
	log.Info().Msg("Received UserHandler.Get request")

	id := req.Id.GetValue()
	if id == "" {
		return nil, status.Errorf(codes.InvalidArgument, "validation error: Missing Id")
	}
	user, err := h.userRepository.Get(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return &userv1.GetResponse{Result: nil}, nil
		}
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("database error: %v", err))
	}

	tempUser, _ := user.ToPB(ctx)
	return &userv1.GetResponse{Result: &tempUser}, nil
}

func (h *userHandler) Create(ctx context.Context, req *userv1.CreateRequest) (rsp *userv1.CreateResponse, err error) {
	log.Info().Msg("Received UserHandler.Create request")

	model := account_entities.UserORM{}
	username := req.Username.GetValue()
	model.Username = &username
	model.FirstName = req.FirstName.GetValue()
	model.LastName = req.LastName.GetValue()
	model.Email = req.Email.GetValue()

	if err := h.userRepository.Create(&model); err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("database error: %v", err))
	}

	// send email (TODO: async `go h.Event.Publish(...)`)
	// ctx := cecontext.WithTopic(context.Background(), topic) // for GCP PubSub
	ctxWithRetries := cloudevents.ContextWithRetriesLinearBackoff(ctx, 10*time.Millisecond, 3)

	if result := h.Event.Send(ctxWithRetries, createEmailEvent(model.Email)); cloudevents.IsNACK(result) { //cloudevents.IsUndelivered(result)
		log.Error().Err(result).Msg("Got Send EmailEvent error. Ignoring")
		// return nil, myErrors.AppError(myErrors.PSE, err)
	}

	// call greeter
	if res, err := h.greeterSrvClient.Hello(ctx, &greeterv1.HelloRequest{Name: req.GetFirstName().GetValue()}); err != nil {
		log.Error().Err(err).Msg("Received greeterService.Hello request error")
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("broker publish error: %v", err))
	} else {
		log.Info().Msgf("Got greeterService responce %s", res.Msg)
	}

	newUser, _ := model.ToPB(ctx)
	return &userv1.CreateResponse{Result: &newUser}, nil
}

func (h *userHandler) Update(ctx context.Context, req *userv1.UpdateRequest) (rsp *userv1.UpdateResponse, err error) {
	log.Info().Msg("Received UserHandler.Update request")
	// Identify the user
	acc, ok := auth.AccountFromContext(ctx)
	if !ok {
		return nil, status.Errorf(codes.Unauthenticated, "A valid auth token is required")
	}
	log.Info().Msgf("Caller Account: %v", acc)

	id := req.Id.GetValue()
	if id == "" {
		return nil, status.Errorf(codes.InvalidArgument, "validation error: Missing Id")
	}

	model := account_entities.UserORM{}
	username := req.Username.GetValue()
	model.Username = &username
	model.FirstName = req.FirstName.GetValue()
	model.LastName = req.LastName.GetValue()
	model.Email = req.Email.GetValue()

	if err := h.userRepository.Update(id, &model); err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("database error: %v", err))
	}

	newUser, _ := model.ToPB(ctx)
	return &userv1.UpdateResponse{Result: &newUser}, nil
}

func (h *userHandler) Delete(ctx context.Context, req *userv1.DeleteRequest) (rsp *userv1.DeleteResponse, err error) {
	log.Info().Msg("Received UserHandler.Delete request")

	id := req.Id.GetValue()
	if id == "" {
		return nil, status.Errorf(codes.InvalidArgument, "validation error: Missing Id")
	}

	model := account_entities.UserORM{}
	model.Id = uuid.FromStringOrNil(id)

	if err := h.userRepository.Delete(&model); err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("database error: %v", err))
	}

	deletedUser, _ := model.ToPB(ctx)
	return &userv1.DeleteResponse{Result: &deletedUser}, nil
}

func createEmailEvent(toEmail string) cloudevents.Event {
	// Create an Event.
	event := cloudevents.NewEvent()
	event.SetSource("github.com/xmlking/grpc-starter-kit/service/emailer")
	event.SetType("account.welcome.email")
	event.SetData(cloudevents.ApplicationJSON, &emailerv1.Message{Subject: "Sumo", To: toEmail})
	return event
}
