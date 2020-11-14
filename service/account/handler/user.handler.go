package handler

import (
	"context"
	"fmt"
	"time"

	cloudevents "github.com/cloudevents/sdk-go/v2"
	"github.com/google/uuid"
	"github.com/grpc-ecosystem/go-grpc-middleware/util/metautils"
	"github.com/rs/zerolog/log"
	"github.com/thoas/go-funk"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/xmlking/toolkit/auth"

	"github.com/xmlking/grpc-starter-kit/ent"
	"github.com/xmlking/grpc-starter-kit/internal/constants"
	account_entities "github.com/xmlking/grpc-starter-kit/mkit/service/account/entities/v1"
	"github.com/xmlking/grpc-starter-kit/mkit/service/account/user/v1"
	"github.com/xmlking/grpc-starter-kit/mkit/service/emailer/v1"
	"github.com/xmlking/grpc-starter-kit/mkit/service/greeter/v1"
	"github.com/xmlking/grpc-starter-kit/service/account/dto"
	"github.com/xmlking/grpc-starter-kit/service/account/repository"
)

// UserHandler struct
type userHandler struct {
	userRepository repository.UserRepository
	emailerClient  cloudevents.Client
	greeterClient  greeterv1.GreeterServiceClient
}

// NewUserHandler returns an instance of `UserServiceHandler`.
func NewUserHandler(repo repository.UserRepository, emailerClient cloudevents.Client, greeterClient greeterv1.GreeterServiceClient) userv1.UserServiceServer {
	return &userHandler{
		userRepository: repo,
		emailerClient:  emailerClient,
		greeterClient:  greeterClient,
	}
}

func (h *userHandler) Exist(ctx context.Context, req *userv1.ExistRequest) (rsp *userv1.ExistResponse, err error) {
	log.Info().Msg("Received UserHandler.Exist request")
	model := ent.User{}
	model.ID, err = uuid.Parse(req.Id.GetValue())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "validation error: Missing Id")
	}
	model.Username = req.Username.GetValue()
	model.Email = req.Email.GetValue()

	var exists bool
	if exists, err = h.userRepository.Exist(ctx, &model); err != nil {
		return nil, status.Errorf(codes.Internal, "Database error: %v", err)
	}
	log.Info().Msgf("user exists? %t", exists)
	return &userv1.ExistResponse{Result: exists}, nil
}

func (h *userHandler) List(ctx context.Context, req *userv1.ListRequest) (rsp *userv1.ListResponse, err error) {
	log.Info().Msg("Received UserHandler.List request")
	model := ent.User{}
	model.Username = req.Username.GetValue()
	model.FirstName = req.FirstName.GetValue()
	model.LastName = req.LastName.GetValue()
	model.Email = req.Email.GetValue()

	total, users, err := h.userRepository.List(ctx, int(req.Limit.GetValue()), int(req.Page.GetValue()), req.Sort.GetValue(), &model)
	if err != nil {
		return nil, readError(err)
	}

	rsp = &userv1.ListResponse{Total: uint32(total)}

	// newUsers := make([]*accountPB.User, len(users))
	// for index, user := range users {
	// 	tmpUser, _ := user.ToPB(ctx)
	// 	newUsers[index] = &tmpUser
	// 	// *newUsers[index], _ = user.ToPB(ctx) ???
	// }
	newUsers := funk.Map(users, func(user *ent.User) *account_entities.User {
		tmpUser, _ := dto.UserToPB(user)
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
	var uid uuid.UUID
	if uid, err = uuid.Parse(id); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "validation error: Parse Id %v", err)
	}

	var model *ent.User
	if model, err = h.userRepository.Get(ctx, uid); err != nil {
		return nil, readError(err)
	}

	tempUser, _ := dto.UserToPB(model)
	return &userv1.GetResponse{Result: &tempUser}, nil
}

func (h *userHandler) Create(ctx context.Context, req *userv1.CreateRequest) (rsp *userv1.CreateResponse, err error) {
	log.Info().Msg("Received UserHandler.Create request")

	model := &ent.User{}
	model.Username = req.Username.GetValue()
	model.FirstName = req.FirstName.GetValue()
	model.LastName = req.LastName.GetValue()
	model.Email = req.Email.GetValue()

	if model, err = h.userRepository.Create(ctx, model); err != nil {
		return nil, readError(err)
	}

	// send email (TODO: async `go h.Event.Publish(...)`)
	// ctx := cecontext.WithTopic(context.Background(), topic) // for GCP PubSub
	ctxWithRetries := cloudevents.ContextWithRetriesLinearBackoff(ctx, 10*time.Millisecond, 3)

	// Create an EmailEvent.
	event := cloudevents.NewEvent()
	event.SetSource("github.com/xmlking/grpc-starter-kit/service/emailer")
	event.SetType("account.welcome.email")
	_ = event.SetData(cloudevents.ApplicationJSON, &emailerv1.Message{Subject: "Sumo", To: model.Email})
	if traceId := metautils.ExtractIncoming(ctx).Get(constants.TraceIDKey); traceId != "" {
		event.SetID(traceId)
	}

	if result := h.emailerClient.Send(ctxWithRetries, event); cloudevents.IsUndelivered(result) {
		log.Error().Err(result).Msg("EmailEvent: Failed to send. Ignoring")
		// return nil, myErrors.AppError(myErrors.PSE, err)
	} else if cloudevents.IsNACK(result) {
		log.Error().Err(result).Msg("EmailEvent: Event not accepted. Ignoring")
		// return nil, myErrors.AppError(myErrors.PSE, err)
	}

	// call greeter
	if res, err := h.greeterClient.Hello(ctx, &greeterv1.HelloRequest{Name: req.GetFirstName().GetValue()}); err != nil {
		log.Error().Err(err).Msg("Received greeterService.Hello request error")
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("broker publish error: %v", err))
	} else {
		log.Info().Msgf("Got greeterService responce %s", res.Msg)
	}

	createdUser, _ := dto.UserToPB(model)
	return &userv1.CreateResponse{Result: &createdUser}, nil
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
	var uid uuid.UUID
	if uid, err = uuid.Parse(id); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "validation error: Parse Id %v", err)
	}

	model := &ent.User{}
	model.ID = uid
	model.Username = req.Username.GetValue()
	model.FirstName = req.FirstName.GetValue()
	model.LastName = req.LastName.GetValue()
	model.Email = req.Email.GetValue()

	if model, err = h.userRepository.Update(ctx, model); err != nil {
		return nil, readError(err)
	}

	updatedUser, _ := dto.UserToPB(model)
	return &userv1.UpdateResponse{Result: &updatedUser}, nil
}

func (h *userHandler) Delete(ctx context.Context, req *userv1.DeleteRequest) (rsp *userv1.DeleteResponse, err error) {
	log.Info().Msg("Received UserHandler.Delete request")

	id := req.Id.GetValue()
	if id == "" {
		return nil, status.Errorf(codes.InvalidArgument, "validation error: Missing Id")
	}
	var uid uuid.UUID
	if uid, err = uuid.Parse(id); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "validation error: Parse Id %v", err)
	}

	var model *ent.User
	if model, err = h.userRepository.Delete(ctx, uid); err != nil {
		return nil, readError(err)
	}

	deletedUser, _ := dto.UserToPB(model)
	return &userv1.DeleteResponse{Result: &deletedUser}, nil
}

func readError(err error) error {
	if ent.IsNotFound(err) {
		return status.Errorf(codes.NotFound, "Database error: %v", err)
	} else if ent.IsConstraintError(err) {
		return status.Errorf(codes.InvalidArgument, "Database constraint error: %v", err)
	}
	return status.Errorf(codes.Internal, fmt.Sprintf("database error: %v", err))
}
