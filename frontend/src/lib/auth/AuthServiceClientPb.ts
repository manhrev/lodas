/**
 * @fileoverview gRPC-Web generated client stub for auth
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!


/* eslint-disable */
// @ts-nocheck


import * as grpcWeb from 'grpc-web';

import * as auth_pb from './auth_pb';
import * as google_protobuf_empty_pb from 'google-protobuf/google/protobuf/empty_pb';


export class AuthClient {
  client_: grpcWeb.AbstractClientBase;
  hostname_: string;
  credentials_: null | { [index: string]: string; };
  options_: null | { [index: string]: any; };

  constructor (hostname: string,
               credentials?: null | { [index: string]: string; },
               options?: null | { [index: string]: any; }) {
    if (!options) options = {};
    if (!credentials) credentials = {};
    options['format'] = 'text';

    this.client_ = new grpcWeb.GrpcWebClientBase(options);
    this.hostname_ = hostname;
    this.credentials_ = credentials;
    this.options_ = options;
  }

  methodDescriptorLogin = new grpcWeb.MethodDescriptor(
    '/auth.Auth/Login',
    grpcWeb.MethodType.UNARY,
    auth_pb.LoginRequest,
    auth_pb.LoginReply,
    (request: auth_pb.LoginRequest) => {
      return request.serializeBinary();
    },
    auth_pb.LoginReply.deserializeBinary
  );

  login(
    request: auth_pb.LoginRequest,
    metadata: grpcWeb.Metadata | null): Promise<auth_pb.LoginReply>;

  login(
    request: auth_pb.LoginRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: auth_pb.LoginReply) => void): grpcWeb.ClientReadableStream<auth_pb.LoginReply>;

  login(
    request: auth_pb.LoginRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: auth_pb.LoginReply) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/auth.Auth/Login',
        request,
        metadata || {},
        this.methodDescriptorLogin,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/auth.Auth/Login',
    request,
    metadata || {},
    this.methodDescriptorLogin);
  }

  methodDescriptorSignUp = new grpcWeb.MethodDescriptor(
    '/auth.Auth/SignUp',
    grpcWeb.MethodType.UNARY,
    auth_pb.SignUpRequest,
    auth_pb.SignUpReply,
    (request: auth_pb.SignUpRequest) => {
      return request.serializeBinary();
    },
    auth_pb.SignUpReply.deserializeBinary
  );

  signUp(
    request: auth_pb.SignUpRequest,
    metadata: grpcWeb.Metadata | null): Promise<auth_pb.SignUpReply>;

  signUp(
    request: auth_pb.SignUpRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: auth_pb.SignUpReply) => void): grpcWeb.ClientReadableStream<auth_pb.SignUpReply>;

  signUp(
    request: auth_pb.SignUpRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: auth_pb.SignUpReply) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/auth.Auth/SignUp',
        request,
        metadata || {},
        this.methodDescriptorSignUp,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/auth.Auth/SignUp',
    request,
    metadata || {},
    this.methodDescriptorSignUp);
  }

  methodDescriptorLogOut = new grpcWeb.MethodDescriptor(
    '/auth.Auth/LogOut',
    grpcWeb.MethodType.UNARY,
    google_protobuf_empty_pb.Empty,
    google_protobuf_empty_pb.Empty,
    (request: google_protobuf_empty_pb.Empty) => {
      return request.serializeBinary();
    },
    google_protobuf_empty_pb.Empty.deserializeBinary
  );

  logOut(
    request: google_protobuf_empty_pb.Empty,
    metadata: grpcWeb.Metadata | null): Promise<google_protobuf_empty_pb.Empty>;

  logOut(
    request: google_protobuf_empty_pb.Empty,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: google_protobuf_empty_pb.Empty) => void): grpcWeb.ClientReadableStream<google_protobuf_empty_pb.Empty>;

  logOut(
    request: google_protobuf_empty_pb.Empty,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: google_protobuf_empty_pb.Empty) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/auth.Auth/LogOut',
        request,
        metadata || {},
        this.methodDescriptorLogOut,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/auth.Auth/LogOut',
    request,
    metadata || {},
    this.methodDescriptorLogOut);
  }

  methodDescriptorMe = new grpcWeb.MethodDescriptor(
    '/auth.Auth/Me',
    grpcWeb.MethodType.UNARY,
    google_protobuf_empty_pb.Empty,
    auth_pb.MeReply,
    (request: google_protobuf_empty_pb.Empty) => {
      return request.serializeBinary();
    },
    auth_pb.MeReply.deserializeBinary
  );

  me(
    request: google_protobuf_empty_pb.Empty,
    metadata: grpcWeb.Metadata | null): Promise<auth_pb.MeReply>;

  me(
    request: google_protobuf_empty_pb.Empty,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: auth_pb.MeReply) => void): grpcWeb.ClientReadableStream<auth_pb.MeReply>;

  me(
    request: google_protobuf_empty_pb.Empty,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: auth_pb.MeReply) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/auth.Auth/Me',
        request,
        metadata || {},
        this.methodDescriptorMe,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/auth.Auth/Me',
    request,
    metadata || {},
    this.methodDescriptorMe);
  }

}

