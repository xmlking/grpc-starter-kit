// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('grpc');
var gkit_service_greeter_v1_greeter_service_pb = require('../../../../gkit/service/greeter/v1/greeter_service_pb.js');

function serialize_gkit_service_greeter_v1_HelloRequest(arg) {
  if (!(arg instanceof gkit_service_greeter_v1_greeter_service_pb.HelloRequest)) {
    throw new Error('Expected argument of type gkit.service.greeter.v1.HelloRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_gkit_service_greeter_v1_HelloRequest(buffer_arg) {
  return gkit_service_greeter_v1_greeter_service_pb.HelloRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_gkit_service_greeter_v1_HelloResponse(arg) {
  if (!(arg instanceof gkit_service_greeter_v1_greeter_service_pb.HelloResponse)) {
    throw new Error('Expected argument of type gkit.service.greeter.v1.HelloResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_gkit_service_greeter_v1_HelloResponse(buffer_arg) {
  return gkit_service_greeter_v1_greeter_service_pb.HelloResponse.deserializeBinary(new Uint8Array(buffer_arg));
}


var GreeterServiceService = exports.GreeterServiceService = {
  // Hello is echo method
hello: {
    path: '/gkit.service.greeter.v1.GreeterService/Hello',
    requestStream: false,
    responseStream: false,
    requestType: gkit_service_greeter_v1_greeter_service_pb.HelloRequest,
    responseType: gkit_service_greeter_v1_greeter_service_pb.HelloResponse,
    requestSerialize: serialize_gkit_service_greeter_v1_HelloRequest,
    requestDeserialize: deserialize_gkit_service_greeter_v1_HelloRequest,
    responseSerialize: serialize_gkit_service_greeter_v1_HelloResponse,
    responseDeserialize: deserialize_gkit_service_greeter_v1_HelloResponse,
  },
};

exports.GreeterServiceClient = grpc.makeGenericClientConstructor(GreeterServiceService);
