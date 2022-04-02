// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('grpc');
var gkit_service_account_profile_v1_profile_service_pb = require('../../../../../gkit/service/account/profile/v1/profile_service_pb.js');
var google_protobuf_timestamp_pb = require('google-protobuf/google/protobuf/timestamp_pb.js');
var google_protobuf_wrappers_pb = require('google-protobuf/google/protobuf/wrappers_pb.js');
var gkit_service_account_entities_v1_entities_pb = require('../../../../../gkit/service/account/entities/v1/entities_pb.js');
var validate_validate_pb = require('../../../../../validate/validate_pb.js');

function serialize_gkit_service_account_profile_v1_CreateRequest(arg) {
  if (!(arg instanceof gkit_service_account_profile_v1_profile_service_pb.CreateRequest)) {
    throw new Error('Expected argument of type gkit.service.account.profile.v1.CreateRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_gkit_service_account_profile_v1_CreateRequest(buffer_arg) {
  return gkit_service_account_profile_v1_profile_service_pb.CreateRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_gkit_service_account_profile_v1_CreateResponse(arg) {
  if (!(arg instanceof gkit_service_account_profile_v1_profile_service_pb.CreateResponse)) {
    throw new Error('Expected argument of type gkit.service.account.profile.v1.CreateResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_gkit_service_account_profile_v1_CreateResponse(buffer_arg) {
  return gkit_service_account_profile_v1_profile_service_pb.CreateResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_gkit_service_account_profile_v1_GetRequest(arg) {
  if (!(arg instanceof gkit_service_account_profile_v1_profile_service_pb.GetRequest)) {
    throw new Error('Expected argument of type gkit.service.account.profile.v1.GetRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_gkit_service_account_profile_v1_GetRequest(buffer_arg) {
  return gkit_service_account_profile_v1_profile_service_pb.GetRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_gkit_service_account_profile_v1_GetResponse(arg) {
  if (!(arg instanceof gkit_service_account_profile_v1_profile_service_pb.GetResponse)) {
    throw new Error('Expected argument of type gkit.service.account.profile.v1.GetResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_gkit_service_account_profile_v1_GetResponse(buffer_arg) {
  return gkit_service_account_profile_v1_profile_service_pb.GetResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_gkit_service_account_profile_v1_ListRequest(arg) {
  if (!(arg instanceof gkit_service_account_profile_v1_profile_service_pb.ListRequest)) {
    throw new Error('Expected argument of type gkit.service.account.profile.v1.ListRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_gkit_service_account_profile_v1_ListRequest(buffer_arg) {
  return gkit_service_account_profile_v1_profile_service_pb.ListRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_gkit_service_account_profile_v1_ListResponse(arg) {
  if (!(arg instanceof gkit_service_account_profile_v1_profile_service_pb.ListResponse)) {
    throw new Error('Expected argument of type gkit.service.account.profile.v1.ListResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_gkit_service_account_profile_v1_ListResponse(buffer_arg) {
  return gkit_service_account_profile_v1_profile_service_pb.ListResponse.deserializeBinary(new Uint8Array(buffer_arg));
}


var ProfileServiceService = exports.ProfileServiceService = {
  list: {
    path: '/gkit.service.account.profile.v1.ProfileService/List',
    requestStream: false,
    responseStream: false,
    requestType: gkit_service_account_profile_v1_profile_service_pb.ListRequest,
    responseType: gkit_service_account_profile_v1_profile_service_pb.ListResponse,
    requestSerialize: serialize_gkit_service_account_profile_v1_ListRequest,
    requestDeserialize: deserialize_gkit_service_account_profile_v1_ListRequest,
    responseSerialize: serialize_gkit_service_account_profile_v1_ListResponse,
    responseDeserialize: deserialize_gkit_service_account_profile_v1_ListResponse,
  },
  get: {
    path: '/gkit.service.account.profile.v1.ProfileService/Get',
    requestStream: false,
    responseStream: false,
    requestType: gkit_service_account_profile_v1_profile_service_pb.GetRequest,
    responseType: gkit_service_account_profile_v1_profile_service_pb.GetResponse,
    requestSerialize: serialize_gkit_service_account_profile_v1_GetRequest,
    requestDeserialize: deserialize_gkit_service_account_profile_v1_GetRequest,
    responseSerialize: serialize_gkit_service_account_profile_v1_GetResponse,
    responseDeserialize: deserialize_gkit_service_account_profile_v1_GetResponse,
  },
  create: {
    path: '/gkit.service.account.profile.v1.ProfileService/Create',
    requestStream: false,
    responseStream: false,
    requestType: gkit_service_account_profile_v1_profile_service_pb.CreateRequest,
    responseType: gkit_service_account_profile_v1_profile_service_pb.CreateResponse,
    requestSerialize: serialize_gkit_service_account_profile_v1_CreateRequest,
    requestDeserialize: deserialize_gkit_service_account_profile_v1_CreateRequest,
    responseSerialize: serialize_gkit_service_account_profile_v1_CreateResponse,
    responseDeserialize: deserialize_gkit_service_account_profile_v1_CreateResponse,
  },
};

exports.ProfileServiceClient = grpc.makeGenericClientConstructor(ProfileServiceService);
