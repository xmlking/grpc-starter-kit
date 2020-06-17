package handler

import (
    "context"
    "time"

    // "github.com/golang/protobuf/ptypes"

    "github.com/jinzhu/gorm"
    "github.com/thoas/go-funk"

    ptypes1 "github.com/golang/protobuf/ptypes"
    "github.com/rs/zerolog"
    uuid "github.com/satori/go.uuid"

    account_entities "github.com/xmlking/grpc-starter-kit/mkit/service/account/entities/v1"
      "github.com/xmlking/grpc-starter-kit/mkit/service/account/profile/v1"
    profilePB "github.com/xmlking/grpc-starter-kit/service/account/proto/profile"
    "github.com/xmlking/grpc-starter-kit/service/account/repository"
    myErrors "github.com/xmlking/grpc-starter-kit/shared/errors"
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

func (ph *profileHandler) List(ctx context.Context, req *profilev1.ListRequest, ) (rsp *profilev1.ListResponse, err error) {
    ph.contextLogger.Debug().Msg("Received ProfileHandler.List request")
    preferredTheme := req.PreferredTheme.GetValue()
    model := account_entities.ProfileORM{
        // UserID:     uuid.FromStringOrNil(req.UserId.GetValue()),
        PreferredTheme: &preferredTheme,
        Gender:         account_entities.Profile_GenderType_name[int32(req.Gender)],
    }

    total, profiles, err := ph.profileRepository.List(req.Limit.GetValue(), req.Page.GetValue(), req.Sort.GetValue(), &model)
    if err != nil {
        return nil, myErrors.AppError(myErrors.DBE, err)
    }
    rsp = &profilev1.ListResponse{Total: total}
    // newProfiles := make([]*pb.Profile, len(profiles))
    // for index, profile := range profiles {
    // 	tempProfile, _ := profile.ToPB(ctx)
    // 	newProfiles[index] = &tempProfile
    // }
    newProfiles := funk.Map(profiles, func(profile *account_entities.ProfileORM) *account_entities.Profile {
        tempProfile, _ := profile.ToPB(ctx)
        return &tempProfile
    }).([]*account_entities.Profile)

    rsp.Results = newProfiles
    return
}

func (ph *profileHandler) Get(ctx context.Context, req *profilev1.GetRequest) (rsp *profilev1.GetResponse, err error) {
    ph.contextLogger.Debug().Msg("Received ProfileHandler.Get request")
    var profile *account_entities.ProfileORM

    switch id := req.Id.(type) {
    case *profilev1.GetRequest_UserId:
        println("GetRequest_UserId")
        println(req.GetId())
        profile, err = ph.profileRepository.GetByUserID(id.UserId.GetValue())
    case *profilev1.GetRequest_ProfileId:
        println("GetRequest_ProfileId")
        println(req.GetId())
        profile, err = ph.profileRepository.Get(id.ProfileId.GetValue())
    case nil:
        return nil, myErrors.ValidationError("mkit.service.account.profile.get", "validation error: Missing Id")
    default:
        return nil, myErrors.ValidationError("mkit.service.account.profile.get", "validation error: Profile.Id has unexpected type %T", id)
    }
    if err != nil {
        if err == gorm.ErrRecordNotFound {
            return &profilev1.GetResponse{Result: nil}, nil
        }
        return nil, myErrors.AppError(myErrors.DBE, err)
    }

    tempProfile, _ := profile.ToPB(ctx)
    return &profilev1.GetResponse{Result: &tempProfile}, nil
}

func (ph *profileHandler) Create(ctx context.Context, req *profilev1.CreateRequest) (rsp *profilev1.CreateResponse, err error) {
    ph.contextLogger.Debug().Msg("Received ProfileHandler.Create request")
    model := account_entities.ProfileORM{}
    userId := uuid.FromStringOrNil(req.UserId.GetValue())
    model.UserId = &userId
    model.Tz = req.Tz.GetValue()
    model.Gender = account_entities.Profile_GenderType_name[int32(req.Gender)]
    model.Avatar = req.Avatar.GetValue()
    if req.Birthday != nil {
        var t time.Time
        var err error
        if t, err = ptypes1.Timestamp(req.Birthday); err != nil {
            return nil, myErrors.ValidationError("mkit.service.account.profile.rceate", "Invalid birthday: %v", err)
        }
        model.Birthday = &t
    }
    preferredTheme := req.PreferredTheme.GetValue()
    model.PreferredTheme = &preferredTheme

    if err := ph.profileRepository.Create(&model); err != nil {
        return nil, myErrors.AppError(myErrors.DBE, err)
    }
    return
}
