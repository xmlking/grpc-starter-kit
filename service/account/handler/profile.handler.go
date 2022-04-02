package handler

import (
	"context"
	"fmt"
	"net/url"
	"time"

	"github.com/google/uuid"
	"github.com/thoas/go-funk"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	ptypes1 "github.com/golang/protobuf/ptypes"
	"github.com/rs/zerolog"

	"github.com/xmlking/grpc-starter-kit/ent"
	"github.com/xmlking/grpc-starter-kit/ent/profile"
	account_entities "github.com/xmlking/grpc-starter-kit/gen/go/gkit/service/account/entities/v1"
	"github.com/xmlking/grpc-starter-kit/gen/go/gkit/service/account/profile/v1"
	"github.com/xmlking/grpc-starter-kit/service/account/dto"
	"github.com/xmlking/grpc-starter-kit/service/account/repository"
)

var (
	Profile_GenderType_name = map[int32]profile.Gender{
		0: profile.GenderUnspecified,
		1: profile.GenderMale,
		2: profile.GenderFemale,
	}
)

// ProfileHandler struct
type profileHandler struct {
	profileRepository repository.ProfileRepository
	contextLogger     zerolog.Logger
}

// NewProfileHandler returns an instance of `ProfileServiceHandler`.
func NewProfileHandler(repo repository.ProfileRepository, logger zerolog.Logger) profilev1.ProfileServiceServer {
	return &profileHandler{
		profileRepository: repo,
		contextLogger:     logger,
	}
}

func (ph *profileHandler) List(ctx context.Context, req *profilev1.ListRequest) (rsp *profilev1.ListResponse, err error) {
	ph.contextLogger.Debug().Msg("Received ProfileHandler.List request")
	preferredTheme := req.PreferredTheme.GetValue()
	model := ent.Profile{
		// UserID:     uuid.Parse(req.UserId.GetValue()),
		PreferredTheme: preferredTheme,
		Gender:         Profile_GenderType_name[int32(req.Gender)],
	}

	total, users, err := ph.profileRepository.List(ctx, int(req.Limit.GetValue()), int(req.Page.GetValue()), req.Sort.GetValue(), &model)
	if err != nil {
		return nil, readError(err)
	}

	rsp = &profilev1.ListResponse{Total: uint32(total)}

	newUsers := funk.Map(users, func(user *ent.Profile) *account_entities.Profile {
		tmpUser, _ := dto.ProfileToPB(user)
		return &tmpUser
	}).([]*account_entities.Profile)

	rsp.Results = newUsers
	return
}

func (ph *profileHandler) Get(ctx context.Context, req *profilev1.GetRequest) (rsp *profilev1.GetResponse, err error) {
	ph.contextLogger.Debug().Msg("Received ProfileHandler.Get request")
	var profile *ent.Profile

	switch id := req.Id.(type) {
	case *profilev1.GetRequest_UserId:
		println("GetRequest_UserId")
		println(req.GetId())
		userID := id.UserId.GetValue()
		if userID == "" {
			return nil, status.Errorf(codes.InvalidArgument, "validation error: Missing userId")
		}
		var userUUID uuid.UUID
		if userUUID, err = uuid.Parse(userID); err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "validation error: Parse userId %v", err)
		}
		profile, err = ph.profileRepository.GetByUserID(ctx, userUUID)
	case *profilev1.GetRequest_ProfileId:
		println("GetRequest_ProfileId")
		println(req.GetId())
		profileID := id.ProfileId.GetValue()
		if profileID == "" {
			return nil, status.Errorf(codes.InvalidArgument, "validation error: Missing profileId")
		}
		var profileUUID uuid.UUID
		if profileUUID, err = uuid.Parse(profileID); err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "validation error: Parse profileId %v", err)
		}
		profile, err = ph.profileRepository.Get(ctx, profileUUID)
	case nil:
		return nil, status.Errorf(codes.InvalidArgument, "validation error: Missing Id")
	default:
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("validation error: Profile.Id has unexpected type %T", id))
	}
	if err != nil {
		return nil, readError(err)
	}

	tempProfile, _ := dto.ProfileToPB(profile)
	return &profilev1.GetResponse{Result: &tempProfile}, nil
}

func (ph *profileHandler) Create(ctx context.Context, req *profilev1.CreateRequest) (rsp *profilev1.CreateResponse, err error) {
	ph.contextLogger.Debug().Msg("Received ProfileHandler.Create request")

	model := &ent.Profile{}
	userID := req.UserId.GetValue()
	if userID == "" {
		return nil, status.Errorf(codes.InvalidArgument, "validation error: Missing userId")
	}
	var userUUID uuid.UUID
	if userUUID, err = uuid.Parse(userID); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "validation error: Parse userId %v", err)
	}
	model.Edges.User.ID = userUUID
	model.Tz = req.Tz.GetValue()
	model.Avatar, err = url.Parse(req.Avatar.GetValue())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "validation error: Avatar Url %v", err)
	}
	model.Gender = Profile_GenderType_name[int32(req.Gender)]
	model.PreferredTheme = req.PreferredTheme.GetValue()
	if req.Birthday != nil {
		var t time.Time
		var err error
		if t, err = ptypes1.Timestamp(req.Birthday); err != nil {
			return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("Invalid birthday: %v", err))
		}
		model.Birthday = t
	}

	if model, err = ph.profileRepository.Create(ctx, model); err != nil {
		return nil, readError(err)
	}

	createdUser, _ := dto.ProfileToPB(model)
	return &profilev1.CreateResponse{Result: &createdUser}, nil
}
