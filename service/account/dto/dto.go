package dto

import (
	ptypes1 "github.com/golang/protobuf/ptypes"
	google_protobuf1 "github.com/golang/protobuf/ptypes/wrappers"

	"github.com/xmlking/grpc-starter-kit/ent"
	"github.com/xmlking/grpc-starter-kit/ent/profile"
	account_entities "github.com/xmlking/grpc-starter-kit/mkit/service/account/entities/v1"
)

// ListMetadata struct
type ListMetadata struct {
	Count  int `json:"count"`
	Offset int `json:"offset"`
	Limit  int `json:"limit"`
	Total  int `json:"total"`
}

// ListResp struct
type ListResp struct {
	Metadata ListMetadata  `json:"metadata"`
	Results  []interface{} `json:"results"`
}

// ListRequest struct
type ListRequest struct {
	Limit  int
	Offset int
}

// UserToPB transform to PB
func UserToPB(user *ent.User) (account_entities.User, error) {
	to := account_entities.User{}
	var err error

	to.Id = &google_protobuf1.StringValue{Value: user.ID.String()}
	if user.CreateTime.IsZero() {
		if to.CreatedAt, err = ptypes1.TimestampProto(user.CreateTime); err != nil {
			return to, err
		}
	}
	if !user.UpdateTime.IsZero() {
		if to.UpdatedAt, err = ptypes1.TimestampProto(user.UpdateTime); err != nil {
			return to, err
		}
	}
	if user.DeleteTime != nil && !user.DeleteTime.IsZero() {
		if to.DeletedAt, err = ptypes1.TimestampProto(*user.DeleteTime); err != nil {
			return to, err
		}
	}
	if user.Username != "" {
		to.Username = &google_protobuf1.StringValue{Value: user.Username}
	}
	to.FirstName = user.FirstName
	to.LastName = user.LastName
	to.Email = user.Email
	if user.Edges.Profile != nil {
		tempProfile, err := ProfileToPB(user.Edges.Profile)
		if err != nil {
			return to, err
		}
		to.Profile = &tempProfile
	}

	return to, err
}

var (
	Profile_GenderType_name = map[int32]profile.Gender{
		0: profile.GenderUnspecified,
		1: profile.GenderMale,
		2: profile.GenderFemale,
	}
	Profile_GenderType_value = map[profile.Gender]int32{
		profile.GenderUnspecified: 0,
		profile.GenderMale:        1,
		profile.GenderFemale:      2,
	}
)

// ProfileToPB transform to PB
func ProfileToPB(profile *ent.Profile) (account_entities.Profile, error) {
	to := account_entities.Profile{}
	var err error

	to.Id = &google_protobuf1.StringValue{Value: profile.ID.String()}
	if !profile.CreateTime.IsZero() {
		if to.CreatedAt, err = ptypes1.TimestampProto(profile.CreateTime); err != nil {
			return to, err
		}
	}
	if !profile.UpdateTime.IsZero() {
		if to.UpdatedAt, err = ptypes1.TimestampProto(profile.UpdateTime); err != nil {
			return to, err
		}
	}
	if profile.DeleteTime != nil && !profile.DeleteTime.IsZero() {
		if to.DeletedAt, err = ptypes1.TimestampProto(*profile.DeleteTime); err != nil {
			return to, err
		}
	}
	to.Tz = profile.Tz
	to.Avatar = profile.Avatar.String()
	to.Gender = account_entities.Profile_GenderType(Profile_GenderType_value[profile.Gender])
	if !profile.Birthday.IsZero() {
		if to.Birthday, err = ptypes1.TimestampProto(profile.Birthday); err != nil {
			return to, err
		}
	}
	if profile.PreferredTheme != "" {
		to.PreferredTheme = &google_protobuf1.StringValue{Value: profile.PreferredTheme}
	}

	return to, err
}
