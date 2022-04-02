// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('grpc');
var gkit_service_account_user_v1_user_service_pb = require('../../../../../gkit/service/account/user/v1/user_service_pb.js');
var google_protobuf_wrappers_pb = require('google-protobuf/google/protobuf/wrappers_pb.js');
var gkit_service_account_entities_v1_entities_pb = require('../../../../../gkit/service/account/entities/v1/entities_pb.js');
var validate_validate_pb = require('../../../../../validate/validate_pb.js');

function serialize_gkit_service_account_user_v1_CreateRequest(arg) {
  if (!(arg instanceof gkit_service_account_user_v1_user_service_pb.CreateRequest)) {
    throw new Error('Expected argument of type gkit.service.account.user.v1.CreateRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_gkit_service_account_user_v1_CreateRequest(buffer_arg) {
  return gkit_service_account_user_v1_user_service_pb.CreateRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_gkit_service_account_user_v1_CreateResponse(arg) {
  if (!(arg instanceof gkit_service_account_user_v1_user_service_pb.CreateResponse)) {
    throw new Error('Expected argument of type gkit.service.account.user.v1.CreateResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_gkit_service_account_user_v1_CreateResponse(buffer_arg) {
  return gkit_service_account_user_v1_user_service_pb.CreateResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_gkit_service_account_user_v1_DeleteRequest(arg) {
  if (!(arg instanceof gkit_service_account_user_v1_user_service_pb.DeleteRequest)) {
    throw new Error('Expected argument of type gkit.service.account.user.v1.DeleteRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_gkit_service_account_user_v1_DeleteRequest(buffer_arg) {
  return gkit_service_account_user_v1_user_service_pb.DeleteRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_gkit_service_account_user_v1_DeleteResponse(arg) {
  if (!(arg instanceof gkit_service_account_user_v1_user_service_pb.DeleteResponse)) {
    throw new Error('Expected argument of type gkit.service.account.user.v1.DeleteResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_gkit_service_account_user_v1_DeleteResponse(buffer_arg) {
  return gkit_service_account_user_v1_user_service_pb.DeleteResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_gkit_service_account_user_v1_ExistRequest(arg) {
  if (!(arg instanceof gkit_service_account_user_v1_user_service_pb.ExistRequest)) {
    throw new Error('Expected argument of type gkit.service.account.user.v1.ExistRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_gkit_service_account_user_v1_ExistRequest(buffer_arg) {
  return gkit_service_account_user_v1_user_service_pb.ExistRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_gkit_service_account_user_v1_ExistResponse(arg) {
  if (!(arg instanceof gkit_service_account_user_v1_user_service_pb.ExistResponse)) {
    throw new Error('Expected argument of type gkit.service.account.user.v1.ExistResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_gkit_service_account_user_v1_ExistResponse(buffer_arg) {
  return gkit_service_account_user_v1_user_service_pb.ExistResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_gkit_service_account_user_v1_GetRequest(arg) {
  if (!(arg instanceof gkit_service_account_user_v1_user_service_pb.GetRequest)) {
    throw new Error('Expected argument of type gkit.service.account.user.v1.GetRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_gkit_service_account_user_v1_GetRequest(buffer_arg) {
  return gkit_service_account_user_v1_user_service_pb.GetRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_gkit_service_account_user_v1_GetResponse(arg) {
  if (!(arg instanceof gkit_service_account_user_v1_user_service_pb.GetResponse)) {
    throw new Error('Expected argument of type gkit.service.account.user.v1.GetResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_gkit_service_account_user_v1_GetResponse(buffer_arg) {
  return gkit_service_account_user_v1_user_service_pb.GetResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_gkit_service_account_user_v1_ListRequest(arg) {
  if (!(arg instanceof gkit_service_account_user_v1_user_service_pb.ListRequest)) {
    throw new Error('Expected argument of type gkit.service.account.user.v1.ListRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_gkit_service_account_user_v1_ListRequest(buffer_arg) {
  return gkit_service_account_user_v1_user_service_pb.ListRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_gkit_service_account_user_v1_ListResponse(arg) {
  if (!(arg instanceof gkit_service_account_user_v1_user_service_pb.ListResponse)) {
    throw new Error('Expected argument of type gkit.service.account.user.v1.ListResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_gkit_service_account_user_v1_ListResponse(buffer_arg) {
  return gkit_service_account_user_v1_user_service_pb.ListResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_gkit_service_account_user_v1_UpdateRequest(arg) {
  if (!(arg instanceof gkit_service_account_user_v1_user_service_pb.UpdateRequest)) {
    throw new Error('Expected argument of type gkit.service.account.user.v1.UpdateRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_gkit_service_account_user_v1_UpdateRequest(buffer_arg) {
  return gkit_service_account_user_v1_user_service_pb.UpdateRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_gkit_service_account_user_v1_UpdateResponse(arg) {
  if (!(arg instanceof gkit_service_account_user_v1_user_service_pb.UpdateResponse)) {
    throw new Error('Expected argument of type gkit.service.account.user.v1.UpdateResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_gkit_service_account_user_v1_UpdateResponse(buffer_arg) {
  return gkit_service_account_user_v1_user_service_pb.UpdateResponse.deserializeBinary(new Uint8Array(buffer_arg));
}


// Ref:
// https://github.com/seizadi/atlas-template/blob/master/resolved/pkg/pb/cmdb.proto
//
// User Service
var UserServiceService = exports.UserServiceService = {
  exist: {
    path: '/gkit.service.account.user.v1.UserService/Exist',
    requestStream: false,
    responseStream: false,
    requestType: gkit_service_account_user_v1_user_service_pb.ExistRequest,
    responseType: gkit_service_account_user_v1_user_service_pb.ExistResponse,
    requestSerialize: serialize_gkit_service_account_user_v1_ExistRequest,
    requestDeserialize: deserialize_gkit_service_account_user_v1_ExistRequest,
    responseSerialize: serialize_gkit_service_account_user_v1_ExistResponse,
    responseDeserialize: deserialize_gkit_service_account_user_v1_ExistResponse,
  },
  list: {
    path: '/gkit.service.account.user.v1.UserService/List',
    requestStream: false,
    responseStream: false,
    requestType: gkit_service_account_user_v1_user_service_pb.ListRequest,
    responseType: gkit_service_account_user_v1_user_service_pb.ListResponse,
    requestSerialize: serialize_gkit_service_account_user_v1_ListRequest,
    requestDeserialize: deserialize_gkit_service_account_user_v1_ListRequest,
    responseSerialize: serialize_gkit_service_account_user_v1_ListResponse,
    responseDeserialize: deserialize_gkit_service_account_user_v1_ListResponse,
  },
  get: {
    path: '/gkit.service.account.user.v1.UserService/Get',
    requestStream: false,
    responseStream: false,
    requestType: gkit_service_account_user_v1_user_service_pb.GetRequest,
    responseType: gkit_service_account_user_v1_user_service_pb.GetResponse,
    requestSerialize: serialize_gkit_service_account_user_v1_GetRequest,
    requestDeserialize: deserialize_gkit_service_account_user_v1_GetRequest,
    responseSerialize: serialize_gkit_service_account_user_v1_GetResponse,
    responseDeserialize: deserialize_gkit_service_account_user_v1_GetResponse,
  },
  create: {
    path: '/gkit.service.account.user.v1.UserService/Create',
    requestStream: false,
    responseStream: false,
    requestType: gkit_service_account_user_v1_user_service_pb.CreateRequest,
    responseType: gkit_service_account_user_v1_user_service_pb.CreateResponse,
    requestSerialize: serialize_gkit_service_account_user_v1_CreateRequest,
    requestDeserialize: deserialize_gkit_service_account_user_v1_CreateRequest,
    responseSerialize: serialize_gkit_service_account_user_v1_CreateResponse,
    responseDeserialize: deserialize_gkit_service_account_user_v1_CreateResponse,
  },
  update: {
    path: '/gkit.service.account.user.v1.UserService/Update',
    requestStream: false,
    responseStream: false,
    requestType: gkit_service_account_user_v1_user_service_pb.UpdateRequest,
    responseType: gkit_service_account_user_v1_user_service_pb.UpdateResponse,
    requestSerialize: serialize_gkit_service_account_user_v1_UpdateRequest,
    requestDeserialize: deserialize_gkit_service_account_user_v1_UpdateRequest,
    responseSerialize: serialize_gkit_service_account_user_v1_UpdateResponse,
    responseDeserialize: deserialize_gkit_service_account_user_v1_UpdateResponse,
  },
  delete: {
    path: '/gkit.service.account.user.v1.UserService/Delete',
    requestStream: false,
    responseStream: false,
    requestType: gkit_service_account_user_v1_user_service_pb.DeleteRequest,
    responseType: gkit_service_account_user_v1_user_service_pb.DeleteResponse,
    requestSerialize: serialize_gkit_service_account_user_v1_DeleteRequest,
    requestDeserialize: deserialize_gkit_service_account_user_v1_DeleteRequest,
    responseSerialize: serialize_gkit_service_account_user_v1_DeleteResponse,
    responseDeserialize: deserialize_gkit_service_account_user_v1_DeleteResponse,
  },
};

exports.UserServiceClient = grpc.makeGenericClientConstructor(UserServiceService);
