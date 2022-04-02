// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('grpc');
var gkit_service_transaction_v1_transaction_pb = require('../../../../gkit/service/transaction/v1/transaction_pb.js');
var google_protobuf_empty_pb = require('google-protobuf/google/protobuf/empty_pb.js');
var gkit_schema_transaction_v1_transaction_pb = require('../../../../gkit/schema/transaction/v1/transaction_pb.js');

function serialize_gkit_service_transaction_v1_ReadRequest(arg) {
  if (!(arg instanceof gkit_service_transaction_v1_transaction_pb.ReadRequest)) {
    throw new Error('Expected argument of type gkit.service.transaction.v1.ReadRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_gkit_service_transaction_v1_ReadRequest(buffer_arg) {
  return gkit_service_transaction_v1_transaction_pb.ReadRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_gkit_service_transaction_v1_ReadResponse(arg) {
  if (!(arg instanceof gkit_service_transaction_v1_transaction_pb.ReadResponse)) {
    throw new Error('Expected argument of type gkit.service.transaction.v1.ReadResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_gkit_service_transaction_v1_ReadResponse(buffer_arg) {
  return gkit_service_transaction_v1_transaction_pb.ReadResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_gkit_service_transaction_v1_WriteRequest(arg) {
  if (!(arg instanceof gkit_service_transaction_v1_transaction_pb.WriteRequest)) {
    throw new Error('Expected argument of type gkit.service.transaction.v1.WriteRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_gkit_service_transaction_v1_WriteRequest(buffer_arg) {
  return gkit_service_transaction_v1_transaction_pb.WriteRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_google_protobuf_Empty(arg) {
  if (!(arg instanceof google_protobuf_empty_pb.Empty)) {
    throw new Error('Expected argument of type google.protobuf.Empty');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_google_protobuf_Empty(buffer_arg) {
  return google_protobuf_empty_pb.Empty.deserializeBinary(new Uint8Array(buffer_arg));
}


var TransactionServiceService = exports.TransactionServiceService = {
  read: {
    path: '/gkit.service.transaction.v1.TransactionService/Read',
    requestStream: false,
    responseStream: false,
    requestType: gkit_service_transaction_v1_transaction_pb.ReadRequest,
    responseType: gkit_service_transaction_v1_transaction_pb.ReadResponse,
    requestSerialize: serialize_gkit_service_transaction_v1_ReadRequest,
    requestDeserialize: deserialize_gkit_service_transaction_v1_ReadRequest,
    responseSerialize: serialize_gkit_service_transaction_v1_ReadResponse,
    responseDeserialize: deserialize_gkit_service_transaction_v1_ReadResponse,
  },
  write: {
    path: '/gkit.service.transaction.v1.TransactionService/Write',
    requestStream: false,
    responseStream: false,
    requestType: gkit_service_transaction_v1_transaction_pb.WriteRequest,
    responseType: google_protobuf_empty_pb.Empty,
    requestSerialize: serialize_gkit_service_transaction_v1_WriteRequest,
    requestDeserialize: deserialize_gkit_service_transaction_v1_WriteRequest,
    responseSerialize: serialize_google_protobuf_Empty,
    responseDeserialize: deserialize_google_protobuf_Empty,
  },
};

exports.TransactionServiceClient = grpc.makeGenericClientConstructor(TransactionServiceService);
