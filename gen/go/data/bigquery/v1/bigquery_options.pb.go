// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0-devel
// 	protoc        (unknown)
// source: data/bigquery/v1/bigquery_options.proto

package bigqueryv1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var file_data_bigquery_v1_bigquery_options_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptorpb.MessageOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         66666667,
		Name:          "data.bigquery.v1.BigQueryTableReference",
		Tag:           "bytes,66666667,opt,name=BigQueryTableReference",
		Filename:      "data/bigquery/v1/bigquery_options.proto",
	},
	{
		ExtendedType:  (*descriptorpb.MessageOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         66666668,
		Name:          "data.bigquery.v1.BigQueryTableDescription",
		Tag:           "bytes,66666668,opt,name=BigQueryTableDescription",
		Filename:      "data/bigquery/v1/bigquery_options.proto",
	},
	{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         66666667,
		Name:          "data.bigquery.v1.BigQueryFieldDescription",
		Tag:           "bytes,66666667,opt,name=BigQueryFieldDescription",
		Filename:      "data/bigquery/v1/bigquery_options.proto",
	},
	{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         66666668,
		Name:          "data.bigquery.v1.BigQueryFieldCategories",
		Tag:           "bytes,66666668,opt,name=BigQueryFieldCategories",
		Filename:      "data/bigquery/v1/bigquery_options.proto",
	},
	{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         66666669,
		Name:          "data.bigquery.v1.BigQueryFieldType",
		Tag:           "bytes,66666669,opt,name=BigQueryFieldType",
		Filename:      "data/bigquery/v1/bigquery_options.proto",
	},
	{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         66666670,
		Name:          "data.bigquery.v1.BigQueryFieldRename",
		Tag:           "bytes,66666670,opt,name=BigQueryFieldRename",
		Filename:      "data/bigquery/v1/bigquery_options.proto",
	},
	{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         66666671,
		Name:          "data.bigquery.v1.BigQueryFieldAppend",
		Tag:           "bytes,66666671,opt,name=BigQueryFieldAppend",
		Filename:      "data/bigquery/v1/bigquery_options.proto",
	},
	{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         66666672,
		Name:          "data.bigquery.v1.BigQueryFieldRegexExtract",
		Tag:           "bytes,66666672,opt,name=BigQueryFieldRegexExtract",
		Filename:      "data/bigquery/v1/bigquery_options.proto",
	},
	{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         66666673,
		Name:          "data.bigquery.v1.BigQueryFieldRegexReplace",
		Tag:           "bytes,66666673,opt,name=BigQueryFieldRegexReplace",
		Filename:      "data/bigquery/v1/bigquery_options.proto",
	},
	{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         66666674,
		Name:          "data.bigquery.v1.BigQueryFieldLocalToUtc",
		Tag:           "bytes,66666674,opt,name=BigQueryFieldLocalToUtc",
		Filename:      "data/bigquery/v1/bigquery_options.proto",
	},
	{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         66666675,
		Name:          "data.bigquery.v1.BigQueryFieldHidden",
		Tag:           "bytes,66666675,opt,name=BigQueryFieldHidden",
		Filename:      "data/bigquery/v1/bigquery_options.proto",
	},
	{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         66666676,
		Name:          "data.bigquery.v1.BigQueryFieldUseDefaultValue",
		Tag:           "bytes,66666676,opt,name=BigQueryFieldUseDefaultValue",
		Filename:      "data/bigquery/v1/bigquery_options.proto",
	},
}

// Extension fields to descriptorpb.MessageOptions.
var (
	// optional string BigQueryTableReference = 66666667;
	E_BigQueryTableReference = &file_data_bigquery_v1_bigquery_options_proto_extTypes[0]
	// optional string BigQueryTableDescription = 66666668;
	E_BigQueryTableDescription = &file_data_bigquery_v1_bigquery_options_proto_extTypes[1]
)

// Extension fields to descriptorpb.FieldOptions.
var (
	// [DescriptionString] Example: "A timestamp."
	//
	// optional string BigQueryFieldDescription = 66666667;
	E_BigQueryFieldDescription = &file_data_bigquery_v1_bigquery_options_proto_extTypes[2]
	// [PolicyTag1, PolicyTag2,...] Example:
	//
	// optional string BigQueryFieldCategories = 66666668;
	E_BigQueryFieldCategories = &file_data_bigquery_v1_bigquery_options_proto_extTypes[3]
	// [BigQueryDataType] Example: "TIMESTAMP"
	//
	// optional string BigQueryFieldType = 66666669;
	E_BigQueryFieldType = &file_data_bigquery_v1_bigquery_options_proto_extTypes[4]
	// [NewFieldName] Example: "LocalDateTime"
	//
	// optional string BigQueryFieldRename = 66666670;
	E_BigQueryFieldRename = &file_data_bigquery_v1_bigquery_options_proto_extTypes[5]
	// [AppendString] Example: "Europe/Stockholm"
	//
	// optional string BigQueryFieldAppend = 66666671;
	E_BigQueryFieldAppend = &file_data_bigquery_v1_bigquery_options_proto_extTypes[6]
	// [RegexPattern] Example "[0-9]{4}-(0[1-9]|1[0-2])-(0[1-9]|[1-2][0-9]|3[0-1])(T| )(2[0-3]|[01][0-9]):[0-5][0-9]:[0-5][0-9]"
	//
	// optional string BigQueryFieldRegexExtract = 66666672;
	E_BigQueryFieldRegexExtract = &file_data_bigquery_v1_bigquery_options_proto_extTypes[7]
	// [RegexPattern, ReplacementString] Example: "(\\+(2[0-3]|[01][0-9]):[0-5][0-9]),Europe/Stockholm"
	//
	// optional string BigQueryFieldRegexReplace = 66666673;
	E_BigQueryFieldRegexReplace = &file_data_bigquery_v1_bigquery_options_proto_extTypes[8]
	// [LocalTimezone, LocalPattern, UtcPattern ] Example: "Europe/Stockholm, yyyy-MM-dd'T'HH:mm:ss, yyyy-MM-dd'T'HH:mm:ssXXX"
	//
	// optional string BigQueryFieldLocalToUtc = 66666674;
	E_BigQueryFieldLocalToUtc = &file_data_bigquery_v1_bigquery_options_proto_extTypes[9]
	//[Hidden] Example: "true"
	//
	// optional string BigQueryFieldHidden = 66666675;
	E_BigQueryFieldHidden = &file_data_bigquery_v1_bigquery_options_proto_extTypes[10]
	//[Hidden] Example: "false"
	//
	// optional string BigQueryFieldUseDefaultValue = 66666676;
	E_BigQueryFieldUseDefaultValue = &file_data_bigquery_v1_bigquery_options_proto_extTypes[11]
)

var File_data_bigquery_v1_bigquery_options_proto protoreflect.FileDescriptor

var file_data_bigquery_v1_bigquery_options_proto_rawDesc = []byte{
	0x0a, 0x27, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x62, 0x69, 0x67, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2f,
	0x76, 0x31, 0x2f, 0x62, 0x69, 0x67, 0x71, 0x75, 0x65, 0x72, 0x79, 0x5f, 0x6f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x64, 0x61, 0x74, 0x61, 0x2e,
	0x62, 0x69, 0x67, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x1a, 0x20, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3a, 0x5a, 0x0a,
	0x16, 0x42, 0x69, 0x67, 0x51, 0x75, 0x65, 0x72, 0x79, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x65,
	0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x1f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xab, 0x81, 0xe5, 0x1f, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x16, 0x42, 0x69, 0x67, 0x51, 0x75, 0x65, 0x72, 0x79, 0x54, 0x61, 0x62, 0x6c, 0x65,
	0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x3a, 0x5e, 0x0a, 0x18, 0x42, 0x69, 0x67,
	0x51, 0x75, 0x65, 0x72, 0x79, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xac, 0x81, 0xe5, 0x1f, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x18, 0x42, 0x69, 0x67, 0x51, 0x75, 0x65, 0x72, 0x79, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x44, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x3a, 0x5c, 0x0a, 0x18, 0x42, 0x69, 0x67,
	0x51, 0x75, 0x65, 0x72, 0x79, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x18, 0xab, 0x81, 0xe5, 0x1f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x18, 0x42,
	0x69, 0x67, 0x51, 0x75, 0x65, 0x72, 0x79, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x44, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x3a, 0x5a, 0x0a, 0x17, 0x42, 0x69, 0x67, 0x51, 0x75,
	0x65, 0x72, 0x79, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69,
	0x65, 0x73, 0x12, 0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x18, 0xac, 0x81, 0xe5, 0x1f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x17, 0x42, 0x69, 0x67, 0x51,
	0x75, 0x65, 0x72, 0x79, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x69, 0x65, 0x73, 0x3a, 0x4e, 0x0a, 0x11, 0x42, 0x69, 0x67, 0x51, 0x75, 0x65, 0x72, 0x79, 0x46,
	0x69, 0x65, 0x6c, 0x64, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64,
	0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xad, 0x81, 0xe5, 0x1f, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x11, 0x42, 0x69, 0x67, 0x51, 0x75, 0x65, 0x72, 0x79, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54,
	0x79, 0x70, 0x65, 0x3a, 0x52, 0x0a, 0x13, 0x42, 0x69, 0x67, 0x51, 0x75, 0x65, 0x72, 0x79, 0x46,
	0x69, 0x65, 0x6c, 0x64, 0x52, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65,
	0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xae, 0x81, 0xe5, 0x1f, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x13, 0x42, 0x69, 0x67, 0x51, 0x75, 0x65, 0x72, 0x79, 0x46, 0x69, 0x65, 0x6c,
	0x64, 0x52, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x3a, 0x52, 0x0a, 0x13, 0x42, 0x69, 0x67, 0x51, 0x75,
	0x65, 0x72, 0x79, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x41, 0x70, 0x70, 0x65, 0x6e, 0x64, 0x12, 0x1d,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xaf, 0x81,
	0xe5, 0x1f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x13, 0x42, 0x69, 0x67, 0x51, 0x75, 0x65, 0x72, 0x79,
	0x46, 0x69, 0x65, 0x6c, 0x64, 0x41, 0x70, 0x70, 0x65, 0x6e, 0x64, 0x3a, 0x5e, 0x0a, 0x19, 0x42,
	0x69, 0x67, 0x51, 0x75, 0x65, 0x72, 0x79, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x52, 0x65, 0x67, 0x65,
	0x78, 0x45, 0x78, 0x74, 0x72, 0x61, 0x63, 0x74, 0x12, 0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64,
	0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xb0, 0x81, 0xe5, 0x1f, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x19, 0x42, 0x69, 0x67, 0x51, 0x75, 0x65, 0x72, 0x79, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x52,
	0x65, 0x67, 0x65, 0x78, 0x45, 0x78, 0x74, 0x72, 0x61, 0x63, 0x74, 0x3a, 0x5e, 0x0a, 0x19, 0x42,
	0x69, 0x67, 0x51, 0x75, 0x65, 0x72, 0x79, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x52, 0x65, 0x67, 0x65,
	0x78, 0x52, 0x65, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x12, 0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64,
	0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xb1, 0x81, 0xe5, 0x1f, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x19, 0x42, 0x69, 0x67, 0x51, 0x75, 0x65, 0x72, 0x79, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x52,
	0x65, 0x67, 0x65, 0x78, 0x52, 0x65, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x3a, 0x5a, 0x0a, 0x17, 0x42,
	0x69, 0x67, 0x51, 0x75, 0x65, 0x72, 0x79, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4c, 0x6f, 0x63, 0x61,
	0x6c, 0x54, 0x6f, 0x55, 0x74, 0x63, 0x12, 0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xb2, 0x81, 0xe5, 0x1f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x17,
	0x42, 0x69, 0x67, 0x51, 0x75, 0x65, 0x72, 0x79, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4c, 0x6f, 0x63,
	0x61, 0x6c, 0x54, 0x6f, 0x55, 0x74, 0x63, 0x3a, 0x52, 0x0a, 0x13, 0x42, 0x69, 0x67, 0x51, 0x75,
	0x65, 0x72, 0x79, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x48, 0x69, 0x64, 0x64, 0x65, 0x6e, 0x12, 0x1d,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xb3, 0x81,
	0xe5, 0x1f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x13, 0x42, 0x69, 0x67, 0x51, 0x75, 0x65, 0x72, 0x79,
	0x46, 0x69, 0x65, 0x6c, 0x64, 0x48, 0x69, 0x64, 0x64, 0x65, 0x6e, 0x3a, 0x64, 0x0a, 0x1c, 0x42,
	0x69, 0x67, 0x51, 0x75, 0x65, 0x72, 0x79, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x55, 0x73, 0x65, 0x44,
	0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x1d, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69,
	0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xb4, 0x81, 0xe5, 0x1f, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x1c, 0x42, 0x69, 0x67, 0x51, 0x75, 0x65, 0x72, 0x79, 0x46, 0x69, 0x65,
	0x6c, 0x64, 0x55, 0x73, 0x65, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x42, 0xd6, 0x01, 0x0a, 0x14, 0x63, 0x6f, 0x6d, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x62,
	0x69, 0x67, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x42, 0x14, 0x42, 0x69, 0x67, 0x71,
	0x75, 0x65, 0x72, 0x79, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x50, 0x01, 0x5a, 0x46, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x78,
	0x6d, 0x6c, 0x6b, 0x69, 0x6e, 0x67, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x73, 0x74, 0x61, 0x72,
	0x74, 0x65, 0x72, 0x2d, 0x6b, 0x69, 0x74, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x64,
	0x61, 0x74, 0x61, 0x2f, 0x62, 0x69, 0x67, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2f, 0x76, 0x31, 0x3b,
	0x62, 0x69, 0x67, 0x71, 0x75, 0x65, 0x72, 0x79, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x44, 0x42, 0x58,
	0xaa, 0x02, 0x10, 0x44, 0x61, 0x74, 0x61, 0x2e, 0x42, 0x69, 0x67, 0x71, 0x75, 0x65, 0x72, 0x79,
	0x2e, 0x56, 0x31, 0xca, 0x02, 0x10, 0x44, 0x61, 0x74, 0x61, 0x5c, 0x42, 0x69, 0x67, 0x71, 0x75,
	0x65, 0x72, 0x79, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x1c, 0x44, 0x61, 0x74, 0x61, 0x5c, 0x42, 0x69,
	0x67, 0x71, 0x75, 0x65, 0x72, 0x79, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x12, 0x44, 0x61, 0x74, 0x61, 0x3a, 0x3a, 0x42, 0x69,
	0x67, 0x71, 0x75, 0x65, 0x72, 0x79, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var file_data_bigquery_v1_bigquery_options_proto_goTypes = []interface{}{
	(*descriptorpb.MessageOptions)(nil), // 0: google.protobuf.MessageOptions
	(*descriptorpb.FieldOptions)(nil),   // 1: google.protobuf.FieldOptions
}
var file_data_bigquery_v1_bigquery_options_proto_depIdxs = []int32{
	0,  // 0: data.bigquery.v1.BigQueryTableReference:extendee -> google.protobuf.MessageOptions
	0,  // 1: data.bigquery.v1.BigQueryTableDescription:extendee -> google.protobuf.MessageOptions
	1,  // 2: data.bigquery.v1.BigQueryFieldDescription:extendee -> google.protobuf.FieldOptions
	1,  // 3: data.bigquery.v1.BigQueryFieldCategories:extendee -> google.protobuf.FieldOptions
	1,  // 4: data.bigquery.v1.BigQueryFieldType:extendee -> google.protobuf.FieldOptions
	1,  // 5: data.bigquery.v1.BigQueryFieldRename:extendee -> google.protobuf.FieldOptions
	1,  // 6: data.bigquery.v1.BigQueryFieldAppend:extendee -> google.protobuf.FieldOptions
	1,  // 7: data.bigquery.v1.BigQueryFieldRegexExtract:extendee -> google.protobuf.FieldOptions
	1,  // 8: data.bigquery.v1.BigQueryFieldRegexReplace:extendee -> google.protobuf.FieldOptions
	1,  // 9: data.bigquery.v1.BigQueryFieldLocalToUtc:extendee -> google.protobuf.FieldOptions
	1,  // 10: data.bigquery.v1.BigQueryFieldHidden:extendee -> google.protobuf.FieldOptions
	1,  // 11: data.bigquery.v1.BigQueryFieldUseDefaultValue:extendee -> google.protobuf.FieldOptions
	12, // [12:12] is the sub-list for method output_type
	12, // [12:12] is the sub-list for method input_type
	12, // [12:12] is the sub-list for extension type_name
	0,  // [0:12] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_data_bigquery_v1_bigquery_options_proto_init() }
func file_data_bigquery_v1_bigquery_options_proto_init() {
	if File_data_bigquery_v1_bigquery_options_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_data_bigquery_v1_bigquery_options_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 12,
			NumServices:   0,
		},
		GoTypes:           file_data_bigquery_v1_bigquery_options_proto_goTypes,
		DependencyIndexes: file_data_bigquery_v1_bigquery_options_proto_depIdxs,
		ExtensionInfos:    file_data_bigquery_v1_bigquery_options_proto_extTypes,
	}.Build()
	File_data_bigquery_v1_bigquery_options_proto = out.File
	file_data_bigquery_v1_bigquery_options_proto_rawDesc = nil
	file_data_bigquery_v1_bigquery_options_proto_goTypes = nil
	file_data_bigquery_v1_bigquery_options_proto_depIdxs = nil
}