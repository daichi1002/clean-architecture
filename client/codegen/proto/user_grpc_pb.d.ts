// package: user
// file: proto/user.proto

/* tslint:disable */
/* eslint-disable */

import * as grpc from "@grpc/grpc-js";
import * as proto_user_pb from "./user_pb";
import * as google_protobuf_timestamp_pb from "google-protobuf/google/protobuf/timestamp_pb";

interface IUserServiceService
  extends grpc.ServiceDefinition<grpc.UntypedServiceImplementation> {
  listUsers: IUserServiceService_IListUsers;
  createUser: IUserServiceService_ICreateUser;
  getUser: IUserServiceService_IGetUser;
}

interface IUserServiceService_IListUsers
  extends grpc.MethodDefinition<
    proto_user_pb.ListUsersRequest,
    proto_user_pb.ListUsersResponse
  > {
  path: "/user.UserService/ListUsers";
  requestStream: false;
  responseStream: false;
  requestSerialize: grpc.serialize<proto_user_pb.ListUsersRequest>;
  requestDeserialize: grpc.deserialize<proto_user_pb.ListUsersRequest>;
  responseSerialize: grpc.serialize<proto_user_pb.ListUsersResponse>;
  responseDeserialize: grpc.deserialize<proto_user_pb.ListUsersResponse>;
}
interface IUserServiceService_ICreateUser
  extends grpc.MethodDefinition<
    proto_user_pb.CreateUserRequest,
    proto_user_pb.CreateUserResponse
  > {
  path: "/user.UserService/CreateUser";
  requestStream: false;
  responseStream: false;
  requestSerialize: grpc.serialize<proto_user_pb.CreateUserRequest>;
  requestDeserialize: grpc.deserialize<proto_user_pb.CreateUserRequest>;
  responseSerialize: grpc.serialize<proto_user_pb.CreateUserResponse>;
  responseDeserialize: grpc.deserialize<proto_user_pb.CreateUserResponse>;
}
interface IUserServiceService_IGetUser
  extends grpc.MethodDefinition<
    proto_user_pb.GetUserRequest,
    proto_user_pb.GetUserResponse
  > {
  path: "/user.UserService/GetUser";
  requestStream: false;
  responseStream: false;
  requestSerialize: grpc.serialize<proto_user_pb.GetUserRequest>;
  requestDeserialize: grpc.deserialize<proto_user_pb.GetUserRequest>;
  responseSerialize: grpc.serialize<proto_user_pb.GetUserResponse>;
  responseDeserialize: grpc.deserialize<proto_user_pb.GetUserResponse>;
}

export const UserServiceService: IUserServiceService;

export interface IUserServiceServer extends grpc.UntypedServiceImplementation {
  listUsers: grpc.handleUnaryCall<
    proto_user_pb.ListUsersRequest,
    proto_user_pb.ListUsersResponse
  >;
  createUser: grpc.handleUnaryCall<
    proto_user_pb.CreateUserRequest,
    proto_user_pb.CreateUserResponse
  >;
  getUser: grpc.handleUnaryCall<
    proto_user_pb.GetUserRequest,
    proto_user_pb.GetUserResponse
  >;
}

export interface IUserServiceClient {
  listUsers(
    request: proto_user_pb.ListUsersRequest,
    callback: (
      error: grpc.ServiceError | null,
      response: proto_user_pb.ListUsersResponse
    ) => void
  ): grpc.ClientUnaryCall;
  listUsers(
    request: proto_user_pb.ListUsersRequest,
    metadata: grpc.Metadata,
    callback: (
      error: grpc.ServiceError | null,
      response: proto_user_pb.ListUsersResponse
    ) => void
  ): grpc.ClientUnaryCall;
  listUsers(
    request: proto_user_pb.ListUsersRequest,
    metadata: grpc.Metadata,
    options: Partial<grpc.CallOptions>,
    callback: (
      error: grpc.ServiceError | null,
      response: proto_user_pb.ListUsersResponse
    ) => void
  ): grpc.ClientUnaryCall;
  createUser(
    request: proto_user_pb.CreateUserRequest,
    callback: (
      error: grpc.ServiceError | null,
      response: proto_user_pb.CreateUserResponse
    ) => void
  ): grpc.ClientUnaryCall;
  createUser(
    request: proto_user_pb.CreateUserRequest,
    metadata: grpc.Metadata,
    callback: (
      error: grpc.ServiceError | null,
      response: proto_user_pb.CreateUserResponse
    ) => void
  ): grpc.ClientUnaryCall;
  createUser(
    request: proto_user_pb.CreateUserRequest,
    metadata: grpc.Metadata,
    options: Partial<grpc.CallOptions>,
    callback: (
      error: grpc.ServiceError | null,
      response: proto_user_pb.CreateUserResponse
    ) => void
  ): grpc.ClientUnaryCall;
  getUser(
    request: proto_user_pb.GetUserRequest,
    callback: (
      error: grpc.ServiceError | null,
      response: proto_user_pb.GetUserResponse
    ) => void
  ): grpc.ClientUnaryCall;
  getUser(
    request: proto_user_pb.GetUserRequest,
    metadata: grpc.Metadata,
    callback: (
      error: grpc.ServiceError | null,
      response: proto_user_pb.GetUserResponse
    ) => void
  ): grpc.ClientUnaryCall;
  getUser(
    request: proto_user_pb.GetUserRequest,
    metadata: grpc.Metadata,
    options: Partial<grpc.CallOptions>,
    callback: (
      error: grpc.ServiceError | null,
      response: proto_user_pb.GetUserResponse
    ) => void
  ): grpc.ClientUnaryCall;
}

export class UserServiceClient
  extends grpc.Client
  implements IUserServiceClient
{
  constructor(
    address: string,
    credentials: grpc.ChannelCredentials,
    options?: Partial<grpc.ClientOptions>
  );
  public listUsers(
    request: proto_user_pb.ListUsersRequest,
    callback: (
      error: grpc.ServiceError | null,
      response: proto_user_pb.ListUsersResponse
    ) => void
  ): grpc.ClientUnaryCall;
  public listUsers(
    request: proto_user_pb.ListUsersRequest,
    metadata: grpc.Metadata,
    callback: (
      error: grpc.ServiceError | null,
      response: proto_user_pb.ListUsersResponse
    ) => void
  ): grpc.ClientUnaryCall;
  public listUsers(
    request: proto_user_pb.ListUsersRequest,
    metadata: grpc.Metadata,
    options: Partial<grpc.CallOptions>,
    callback: (
      error: grpc.ServiceError | null,
      response: proto_user_pb.ListUsersResponse
    ) => void
  ): grpc.ClientUnaryCall;
  public createUser(
    request: proto_user_pb.CreateUserRequest,
    callback: (
      error: grpc.ServiceError | null,
      response: proto_user_pb.CreateUserResponse
    ) => void
  ): grpc.ClientUnaryCall;
  public createUser(
    request: proto_user_pb.CreateUserRequest,
    metadata: grpc.Metadata,
    callback: (
      error: grpc.ServiceError | null,
      response: proto_user_pb.CreateUserResponse
    ) => void
  ): grpc.ClientUnaryCall;
  public createUser(
    request: proto_user_pb.CreateUserRequest,
    metadata: grpc.Metadata,
    options: Partial<grpc.CallOptions>,
    callback: (
      error: grpc.ServiceError | null,
      response: proto_user_pb.CreateUserResponse
    ) => void
  ): grpc.ClientUnaryCall;
  public getUser(
    request: proto_user_pb.GetUserRequest,
    callback: (
      error: grpc.ServiceError | null,
      response: proto_user_pb.GetUserResponse
    ) => void
  ): grpc.ClientUnaryCall;
  public getUser(
    request: proto_user_pb.GetUserRequest,
    metadata: grpc.Metadata,
    callback: (
      error: grpc.ServiceError | null,
      response: proto_user_pb.GetUserResponse
    ) => void
  ): grpc.ClientUnaryCall;
  public getUser(
    request: proto_user_pb.GetUserRequest,
    metadata: grpc.Metadata,
    options: Partial<grpc.CallOptions>,
    callback: (
      error: grpc.ServiceError | null,
      response: proto_user_pb.GetUserResponse
    ) => void
  ): grpc.ClientUnaryCall;
}
