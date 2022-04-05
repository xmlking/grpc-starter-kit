// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0-devel
// 	protoc        (unknown)
// source: gkit/service/account/profile/v1/profile_service.proto

package profilev1

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	v1 "github.com/xmlking/grpc-starter-kit/gen/go/gkit/service/account/entities/v1"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// FIXME: https://github.com/envoyproxy/protoc-gen-validate/issues/223
// Workaround in .override.go
type ListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Limit          *wrapperspb.UInt32Value `protobuf:"bytes,1,opt,name=limit,proto3" json:"limit,omitempty"`
	Page           *wrapperspb.UInt32Value `protobuf:"bytes,2,opt,name=page,proto3" json:"page,omitempty"`
	Sort           *wrapperspb.StringValue `protobuf:"bytes,3,opt,name=sort,proto3" json:"sort,omitempty"`
	PreferredTheme *wrapperspb.StringValue `protobuf:"bytes,4,opt,name=preferred_theme,json=preferredTheme,proto3" json:"preferred_theme,omitempty"`
	Gender         v1.Profile_GenderType   `protobuf:"varint,5,opt,name=gender,proto3,enum=gkit.service.account.entities.v1.Profile_GenderType" json:"gender,omitempty"`
}

func (x *ListRequest) Reset() {
	*x = ListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gkit_service_account_profile_v1_profile_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRequest) ProtoMessage() {}

func (x *ListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_gkit_service_account_profile_v1_profile_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRequest.ProtoReflect.Descriptor instead.
func (*ListRequest) Descriptor() ([]byte, []int) {
	return file_gkit_service_account_profile_v1_profile_service_proto_rawDescGZIP(), []int{0}
}

func (x *ListRequest) GetLimit() *wrapperspb.UInt32Value {
	if x != nil {
		return x.Limit
	}
	return nil
}

func (x *ListRequest) GetPage() *wrapperspb.UInt32Value {
	if x != nil {
		return x.Page
	}
	return nil
}

func (x *ListRequest) GetSort() *wrapperspb.StringValue {
	if x != nil {
		return x.Sort
	}
	return nil
}

func (x *ListRequest) GetPreferredTheme() *wrapperspb.StringValue {
	if x != nil {
		return x.PreferredTheme
	}
	return nil
}

func (x *ListRequest) GetGender() v1.Profile_GenderType {
	if x != nil {
		return x.Gender
	}
	return v1.Profile_GenderType(0)
}

type ListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Results []*v1.Profile `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty"`
	Total   uint32        `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
}

func (x *ListResponse) Reset() {
	*x = ListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gkit_service_account_profile_v1_profile_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListResponse) ProtoMessage() {}

func (x *ListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_gkit_service_account_profile_v1_profile_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListResponse.ProtoReflect.Descriptor instead.
func (*ListResponse) Descriptor() ([]byte, []int) {
	return file_gkit_service_account_profile_v1_profile_service_proto_rawDescGZIP(), []int{1}
}

func (x *ListResponse) GetResults() []*v1.Profile {
	if x != nil {
		return x.Results
	}
	return nil
}

func (x *ListResponse) GetTotal() uint32 {
	if x != nil {
		return x.Total
	}
	return 0
}

type GetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Id:
	//	*GetRequest_ProfileId
	//	*GetRequest_UserId
	Id isGetRequest_Id `protobuf_oneof:"id"`
}

func (x *GetRequest) Reset() {
	*x = GetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gkit_service_account_profile_v1_profile_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRequest) ProtoMessage() {}

func (x *GetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_gkit_service_account_profile_v1_profile_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRequest.ProtoReflect.Descriptor instead.
func (*GetRequest) Descriptor() ([]byte, []int) {
	return file_gkit_service_account_profile_v1_profile_service_proto_rawDescGZIP(), []int{2}
}

func (m *GetRequest) GetId() isGetRequest_Id {
	if m != nil {
		return m.Id
	}
	return nil
}

func (x *GetRequest) GetProfileId() *wrapperspb.StringValue {
	if x, ok := x.GetId().(*GetRequest_ProfileId); ok {
		return x.ProfileId
	}
	return nil
}

func (x *GetRequest) GetUserId() *wrapperspb.StringValue {
	if x, ok := x.GetId().(*GetRequest_UserId); ok {
		return x.UserId
	}
	return nil
}

type isGetRequest_Id interface {
	isGetRequest_Id()
}

type GetRequest_ProfileId struct {
	ProfileId *wrapperspb.StringValue `protobuf:"bytes,1,opt,name=profile_id,json=profileId,proto3,oneof"` // Not Working
}

type GetRequest_UserId struct {
	UserId *wrapperspb.StringValue `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3,oneof"` // Not Working
}

func (*GetRequest_ProfileId) isGetRequest_Id() {}

func (*GetRequest_UserId) isGetRequest_Id() {}

type GetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result *v1.Profile `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *GetResponse) Reset() {
	*x = GetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gkit_service_account_profile_v1_profile_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetResponse) ProtoMessage() {}

func (x *GetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_gkit_service_account_profile_v1_profile_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetResponse.ProtoReflect.Descriptor instead.
func (*GetResponse) Descriptor() ([]byte, []int) {
	return file_gkit_service_account_profile_v1_profile_service_proto_rawDescGZIP(), []int{3}
}

func (x *GetResponse) GetResult() *v1.Profile {
	if x != nil {
		return x.Result
	}
	return nil
}

type CreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId         *wrapperspb.StringValue `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"` // Not Working
	Tz             *wrapperspb.StringValue `protobuf:"bytes,3,opt,name=tz,proto3" json:"tz,omitempty"`
	Avatar         *wrapperspb.StringValue `protobuf:"bytes,4,opt,name=avatar,proto3" json:"avatar,omitempty"`
	Gender         v1.Profile_GenderType   `protobuf:"varint,5,opt,name=gender,proto3,enum=gkit.service.account.entities.v1.Profile_GenderType" json:"gender,omitempty"`
	Birthday       *timestamppb.Timestamp  `protobuf:"bytes,6,opt,name=birthday,proto3" json:"birthday,omitempty"`
	PreferredTheme *wrapperspb.StringValue `protobuf:"bytes,7,opt,name=preferred_theme,json=preferredTheme,proto3" json:"preferred_theme,omitempty"`
}

func (x *CreateRequest) Reset() {
	*x = CreateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gkit_service_account_profile_v1_profile_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRequest) ProtoMessage() {}

func (x *CreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_gkit_service_account_profile_v1_profile_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRequest.ProtoReflect.Descriptor instead.
func (*CreateRequest) Descriptor() ([]byte, []int) {
	return file_gkit_service_account_profile_v1_profile_service_proto_rawDescGZIP(), []int{4}
}

func (x *CreateRequest) GetUserId() *wrapperspb.StringValue {
	if x != nil {
		return x.UserId
	}
	return nil
}

func (x *CreateRequest) GetTz() *wrapperspb.StringValue {
	if x != nil {
		return x.Tz
	}
	return nil
}

func (x *CreateRequest) GetAvatar() *wrapperspb.StringValue {
	if x != nil {
		return x.Avatar
	}
	return nil
}

func (x *CreateRequest) GetGender() v1.Profile_GenderType {
	if x != nil {
		return x.Gender
	}
	return v1.Profile_GenderType(0)
}

func (x *CreateRequest) GetBirthday() *timestamppb.Timestamp {
	if x != nil {
		return x.Birthday
	}
	return nil
}

func (x *CreateRequest) GetPreferredTheme() *wrapperspb.StringValue {
	if x != nil {
		return x.PreferredTheme
	}
	return nil
}

type CreateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result *v1.Profile `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *CreateResponse) Reset() {
	*x = CreateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gkit_service_account_profile_v1_profile_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateResponse) ProtoMessage() {}

func (x *CreateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_gkit_service_account_profile_v1_profile_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateResponse.ProtoReflect.Descriptor instead.
func (*CreateResponse) Descriptor() ([]byte, []int) {
	return file_gkit_service_account_profile_v1_profile_service_proto_rawDescGZIP(), []int{5}
}

func (x *CreateResponse) GetResult() *v1.Profile {
	if x != nil {
		return x.Result
	}
	return nil
}

var File_gkit_service_account_profile_v1_profile_service_proto protoreflect.FileDescriptor

var file_gkit_service_account_profile_v1_profile_service_proto_rawDesc = []byte{
	0x0a, 0x35, 0x67, 0x6b, 0x69, 0x74, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2f, 0x76,
	0x31, 0x2f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1f, 0x67, 0x6b, 0x69, 0x74, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x2f, 0x67, 0x6b, 0x69, 0x74, 0x2f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2f, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x6e, 0x74, 0x69, 0x74,
	0x69, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70,
	0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x83, 0x03, 0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x3d, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x55, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x42, 0x09, 0xfa, 0x42, 0x06, 0x2a, 0x04, 0x18, 0x64, 0x28, 0x01, 0x52, 0x05, 0x6c, 0x69, 0x6d,
	0x69, 0x74, 0x12, 0x39, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x55, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x07,
	0xfa, 0x42, 0x04, 0x2a, 0x02, 0x28, 0x01, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x30, 0x0a,
	0x04, 0x73, 0x6f, 0x72, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74,
	0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x12,
	0x6c, 0x0a, 0x0f, 0x70, 0x72, 0x65, 0x66, 0x65, 0x72, 0x72, 0x65, 0x64, 0x5f, 0x74, 0x68, 0x65,
	0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e,
	0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x25, 0xfa, 0x42, 0x22, 0x72, 0x20, 0x52, 0x04, 0x64,
	0x61, 0x72, 0x6b, 0x52, 0x05, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x52, 0x06, 0x63, 0x6f, 0x73, 0x6d,
	0x69, 0x63, 0x52, 0x09, 0x63, 0x6f, 0x72, 0x70, 0x6f, 0x72, 0x61, 0x74, 0x65, 0x52, 0x0e, 0x70,
	0x72, 0x65, 0x66, 0x65, 0x72, 0x72, 0x65, 0x64, 0x54, 0x68, 0x65, 0x6d, 0x65, 0x12, 0x5a, 0x0a,
	0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x34, 0x2e,
	0x67, 0x6b, 0x69, 0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x47, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x54,
	0x79, 0x70, 0x65, 0x42, 0x0c, 0xfa, 0x42, 0x09, 0x82, 0x01, 0x06, 0x18, 0x00, 0x18, 0x01, 0x18,
	0x02, 0x52, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x22, 0x6e, 0x0a, 0x0c, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x43, 0x0a, 0x07, 0x72, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x67, 0x6b, 0x69,
	0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72,
	0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x12, 0x14,
	0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x74,
	0x6f, 0x74, 0x61, 0x6c, 0x3a, 0x03, 0xf8, 0x42, 0x01, 0x22, 0x9e, 0x01, 0x0a, 0x0a, 0x47, 0x65,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x47, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x66,
	0x69, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53,
	0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x72,
	0x03, 0xb0, 0x01, 0x01, 0x48, 0x00, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x49,
	0x64, 0x12, 0x41, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x42, 0x08, 0xfa, 0x42, 0x05, 0x72, 0x03, 0xb0, 0x01, 0x01, 0x48, 0x00, 0x52, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x42, 0x04, 0x0a, 0x02, 0x69, 0x64, 0x22, 0x55, 0x0a, 0x0b, 0x47, 0x65,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x41, 0x0a, 0x06, 0x72, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x67, 0x6b, 0x69, 0x74,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x6f,
	0x66, 0x69, 0x6c, 0x65, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x3a, 0x03, 0xf8, 0x42,
	0x01, 0x22, 0xcd, 0x03, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x3f, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x72, 0x03, 0xb0, 0x01, 0x01, 0x52, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x2c, 0x0a, 0x02, 0x74, 0x7a, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x02,
	0x74, 0x7a, 0x12, 0x3e, 0x0a, 0x06, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x42, 0x08, 0xfa, 0x42, 0x05, 0x72, 0x03, 0x90, 0x01, 0x01, 0x52, 0x06, 0x61, 0x76, 0x61, 0x74,
	0x61, 0x72, 0x12, 0x5a, 0x0a, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x34, 0x2e, 0x67, 0x6b, 0x69, 0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69,
	0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x47, 0x65,
	0x6e, 0x64, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x42, 0x0c, 0xfa, 0x42, 0x09, 0x82, 0x01, 0x06,
	0x18, 0x00, 0x18, 0x01, 0x18, 0x02, 0x52, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x40,
	0x0a, 0x08, 0x62, 0x69, 0x72, 0x74, 0x68, 0x64, 0x61, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x08, 0xfa, 0x42,
	0x05, 0xb2, 0x01, 0x02, 0x38, 0x01, 0x52, 0x08, 0x62, 0x69, 0x72, 0x74, 0x68, 0x64, 0x61, 0x79,
	0x12, 0x6f, 0x0a, 0x0f, 0x70, 0x72, 0x65, 0x66, 0x65, 0x72, 0x72, 0x65, 0x64, 0x5f, 0x74, 0x68,
	0x65, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69,
	0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x28, 0xfa, 0x42, 0x25, 0x72, 0x23, 0x52, 0x04,
	0x64, 0x61, 0x72, 0x6b, 0x52, 0x05, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x52, 0x06, 0x63, 0x6f, 0x73,
	0x6d, 0x69, 0x63, 0x52, 0x09, 0x63, 0x6f, 0x72, 0x70, 0x6f, 0x72, 0x61, 0x74, 0x65, 0xd0, 0x01,
	0x01, 0x52, 0x0e, 0x70, 0x72, 0x65, 0x66, 0x65, 0x72, 0x72, 0x65, 0x64, 0x54, 0x68, 0x65, 0x6d,
	0x65, 0x22, 0x58, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x41, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x67, 0x6b, 0x69, 0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74,
	0x69, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x06,
	0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x3a, 0x03, 0xf8, 0x42, 0x01, 0x32, 0xc2, 0x02, 0x0a, 0x0e,
	0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x63,
	0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x2c, 0x2e, 0x67, 0x6b, 0x69, 0x74, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x2d, 0x2e, 0x67, 0x6b, 0x69, 0x74, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x66,
	0x69, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x60, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x2b, 0x2e, 0x67, 0x6b, 0x69,
	0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2c, 0x2e, 0x67, 0x6b, 0x69, 0x74, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x69, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12,
	0x2e, 0x2e, 0x67, 0x6b, 0x69, 0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x2f, 0x2e, 0x67, 0x6b, 0x69, 0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x42, 0xb1, 0x02, 0x0a, 0x23, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6b, 0x69, 0x74, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x42, 0x13, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c,
	0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a,
	0x54, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x78, 0x6d, 0x6c, 0x6b,
	0x69, 0x6e, 0x67, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x73, 0x74, 0x61, 0x72, 0x74, 0x65, 0x72,
	0x2d, 0x6b, 0x69, 0x74, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x67, 0x6b, 0x69, 0x74,
	0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x2f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x70, 0x72, 0x6f, 0x66,
	0x69, 0x6c, 0x65, 0x76, 0x31, 0xa2, 0x02, 0x04, 0x47, 0x53, 0x41, 0x50, 0xaa, 0x02, 0x1f, 0x47,
	0x6b, 0x69, 0x74, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x2e, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x56, 0x31, 0xca, 0x02,
	0x1f, 0x47, 0x6b, 0x69, 0x74, 0x5c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5c, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5c, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x5c, 0x56, 0x31,
	0xe2, 0x02, 0x2b, 0x47, 0x6b, 0x69, 0x74, 0x5c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5c,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5c, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x5c,
	0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02,
	0x23, 0x47, 0x6b, 0x69, 0x74, 0x3a, 0x3a, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x3a, 0x3a,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x3a, 0x3a, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65,
	0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_gkit_service_account_profile_v1_profile_service_proto_rawDescOnce sync.Once
	file_gkit_service_account_profile_v1_profile_service_proto_rawDescData = file_gkit_service_account_profile_v1_profile_service_proto_rawDesc
)

func file_gkit_service_account_profile_v1_profile_service_proto_rawDescGZIP() []byte {
	file_gkit_service_account_profile_v1_profile_service_proto_rawDescOnce.Do(func() {
		file_gkit_service_account_profile_v1_profile_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_gkit_service_account_profile_v1_profile_service_proto_rawDescData)
	})
	return file_gkit_service_account_profile_v1_profile_service_proto_rawDescData
}

var file_gkit_service_account_profile_v1_profile_service_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_gkit_service_account_profile_v1_profile_service_proto_goTypes = []interface{}{
	(*ListRequest)(nil),            // 0: gkit.service.account.profile.v1.ListRequest
	(*ListResponse)(nil),           // 1: gkit.service.account.profile.v1.ListResponse
	(*GetRequest)(nil),             // 2: gkit.service.account.profile.v1.GetRequest
	(*GetResponse)(nil),            // 3: gkit.service.account.profile.v1.GetResponse
	(*CreateRequest)(nil),          // 4: gkit.service.account.profile.v1.CreateRequest
	(*CreateResponse)(nil),         // 5: gkit.service.account.profile.v1.CreateResponse
	(*wrapperspb.UInt32Value)(nil), // 6: google.protobuf.UInt32Value
	(*wrapperspb.StringValue)(nil), // 7: google.protobuf.StringValue
	(v1.Profile_GenderType)(0),     // 8: gkit.service.account.entities.v1.Profile.GenderType
	(*v1.Profile)(nil),             // 9: gkit.service.account.entities.v1.Profile
	(*timestamppb.Timestamp)(nil),  // 10: google.protobuf.Timestamp
}
var file_gkit_service_account_profile_v1_profile_service_proto_depIdxs = []int32{
	6,  // 0: gkit.service.account.profile.v1.ListRequest.limit:type_name -> google.protobuf.UInt32Value
	6,  // 1: gkit.service.account.profile.v1.ListRequest.page:type_name -> google.protobuf.UInt32Value
	7,  // 2: gkit.service.account.profile.v1.ListRequest.sort:type_name -> google.protobuf.StringValue
	7,  // 3: gkit.service.account.profile.v1.ListRequest.preferred_theme:type_name -> google.protobuf.StringValue
	8,  // 4: gkit.service.account.profile.v1.ListRequest.gender:type_name -> gkit.service.account.entities.v1.Profile.GenderType
	9,  // 5: gkit.service.account.profile.v1.ListResponse.results:type_name -> gkit.service.account.entities.v1.Profile
	7,  // 6: gkit.service.account.profile.v1.GetRequest.profile_id:type_name -> google.protobuf.StringValue
	7,  // 7: gkit.service.account.profile.v1.GetRequest.user_id:type_name -> google.protobuf.StringValue
	9,  // 8: gkit.service.account.profile.v1.GetResponse.result:type_name -> gkit.service.account.entities.v1.Profile
	7,  // 9: gkit.service.account.profile.v1.CreateRequest.user_id:type_name -> google.protobuf.StringValue
	7,  // 10: gkit.service.account.profile.v1.CreateRequest.tz:type_name -> google.protobuf.StringValue
	7,  // 11: gkit.service.account.profile.v1.CreateRequest.avatar:type_name -> google.protobuf.StringValue
	8,  // 12: gkit.service.account.profile.v1.CreateRequest.gender:type_name -> gkit.service.account.entities.v1.Profile.GenderType
	10, // 13: gkit.service.account.profile.v1.CreateRequest.birthday:type_name -> google.protobuf.Timestamp
	7,  // 14: gkit.service.account.profile.v1.CreateRequest.preferred_theme:type_name -> google.protobuf.StringValue
	9,  // 15: gkit.service.account.profile.v1.CreateResponse.result:type_name -> gkit.service.account.entities.v1.Profile
	0,  // 16: gkit.service.account.profile.v1.ProfileService.List:input_type -> gkit.service.account.profile.v1.ListRequest
	2,  // 17: gkit.service.account.profile.v1.ProfileService.Get:input_type -> gkit.service.account.profile.v1.GetRequest
	4,  // 18: gkit.service.account.profile.v1.ProfileService.Create:input_type -> gkit.service.account.profile.v1.CreateRequest
	1,  // 19: gkit.service.account.profile.v1.ProfileService.List:output_type -> gkit.service.account.profile.v1.ListResponse
	3,  // 20: gkit.service.account.profile.v1.ProfileService.Get:output_type -> gkit.service.account.profile.v1.GetResponse
	5,  // 21: gkit.service.account.profile.v1.ProfileService.Create:output_type -> gkit.service.account.profile.v1.CreateResponse
	19, // [19:22] is the sub-list for method output_type
	16, // [16:19] is the sub-list for method input_type
	16, // [16:16] is the sub-list for extension type_name
	16, // [16:16] is the sub-list for extension extendee
	0,  // [0:16] is the sub-list for field type_name
}

func init() { file_gkit_service_account_profile_v1_profile_service_proto_init() }
func file_gkit_service_account_profile_v1_profile_service_proto_init() {
	if File_gkit_service_account_profile_v1_profile_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_gkit_service_account_profile_v1_profile_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_gkit_service_account_profile_v1_profile_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_gkit_service_account_profile_v1_profile_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_gkit_service_account_profile_v1_profile_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_gkit_service_account_profile_v1_profile_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_gkit_service_account_profile_v1_profile_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_gkit_service_account_profile_v1_profile_service_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*GetRequest_ProfileId)(nil),
		(*GetRequest_UserId)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_gkit_service_account_profile_v1_profile_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_gkit_service_account_profile_v1_profile_service_proto_goTypes,
		DependencyIndexes: file_gkit_service_account_profile_v1_profile_service_proto_depIdxs,
		MessageInfos:      file_gkit_service_account_profile_v1_profile_service_proto_msgTypes,
	}.Build()
	File_gkit_service_account_profile_v1_profile_service_proto = out.File
	file_gkit_service_account_profile_v1_profile_service_proto_rawDesc = nil
	file_gkit_service_account_profile_v1_profile_service_proto_goTypes = nil
	file_gkit_service_account_profile_v1_profile_service_proto_depIdxs = nil
}
