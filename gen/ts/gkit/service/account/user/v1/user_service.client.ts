// @generated by protobuf-ts 2.4.0 with parameter long_type_string,generate_dependencies,// @generated from protobuf file "gkit/service/account/user/v1/user_service.proto" (package "gkit.service.account.user.v1", syntax proto3),// tslint:disable
import type { RpcTransport } from "@protobuf-ts/runtime-rpc";
import type { ServiceInfo } from "@protobuf-ts/runtime-rpc";
import { UserService } from "./user_service";
import type { DeleteResponse } from "./user_service";
import type { DeleteRequest } from "./user_service";
import type { UpdateResponse } from "./user_service";
import type { UpdateRequest } from "./user_service";
import type { CreateResponse } from "./user_service";
import type { CreateRequest } from "./user_service";
import type { GetResponse } from "./user_service";
import type { GetRequest } from "./user_service";
import type { ListResponse } from "./user_service";
import type { ListRequest } from "./user_service";
import { stackIntercept } from "@protobuf-ts/runtime-rpc";
import type { ExistResponse } from "./user_service";
import type { ExistRequest } from "./user_service";
import type { UnaryCall } from "@protobuf-ts/runtime-rpc";
import type { RpcOptions } from "@protobuf-ts/runtime-rpc";
// Ref:
// https://github.com/seizadi/atlas-template/blob/master/resolved/pkg/pb/cmdb.proto

/**
 * User Service
 *
 * @generated from protobuf service gkit.service.account.user.v1.UserService
 */
export interface IUserServiceClient {
    /**
     * @generated from protobuf rpc: Exist(gkit.service.account.user.v1.ExistRequest) returns (gkit.service.account.user.v1.ExistResponse);
     */
    exist(input: ExistRequest, options?: RpcOptions): UnaryCall<ExistRequest, ExistResponse>;
    /**
     * @generated from protobuf rpc: List(gkit.service.account.user.v1.ListRequest) returns (gkit.service.account.user.v1.ListResponse);
     */
    list(input: ListRequest, options?: RpcOptions): UnaryCall<ListRequest, ListResponse>;
    /**
     * @generated from protobuf rpc: Get(gkit.service.account.user.v1.GetRequest) returns (gkit.service.account.user.v1.GetResponse);
     */
    get(input: GetRequest, options?: RpcOptions): UnaryCall<GetRequest, GetResponse>;
    /**
     * @generated from protobuf rpc: Create(gkit.service.account.user.v1.CreateRequest) returns (gkit.service.account.user.v1.CreateResponse);
     */
    create(input: CreateRequest, options?: RpcOptions): UnaryCall<CreateRequest, CreateResponse>;
    /**
     * @generated from protobuf rpc: Update(gkit.service.account.user.v1.UpdateRequest) returns (gkit.service.account.user.v1.UpdateResponse);
     */
    update(input: UpdateRequest, options?: RpcOptions): UnaryCall<UpdateRequest, UpdateResponse>;
    /**
     * @generated from protobuf rpc: Delete(gkit.service.account.user.v1.DeleteRequest) returns (gkit.service.account.user.v1.DeleteResponse);
     */
    delete(input: DeleteRequest, options?: RpcOptions): UnaryCall<DeleteRequest, DeleteResponse>;
}
// Ref:
// https://github.com/seizadi/atlas-template/blob/master/resolved/pkg/pb/cmdb.proto

/**
 * User Service
 *
 * @generated from protobuf service gkit.service.account.user.v1.UserService
 */
export class UserServiceClient implements IUserServiceClient, ServiceInfo {
    typeName = UserService.typeName;
    methods = UserService.methods;
    options = UserService.options;
    constructor(private readonly _transport: RpcTransport) {
    }
    /**
     * @generated from protobuf rpc: Exist(gkit.service.account.user.v1.ExistRequest) returns (gkit.service.account.user.v1.ExistResponse);
     */
    exist(input: ExistRequest, options?: RpcOptions): UnaryCall<ExistRequest, ExistResponse> {
        const method = this.methods[0], opt = this._transport.mergeOptions(options);
        return stackIntercept<ExistRequest, ExistResponse>("unary", this._transport, method, opt, input);
    }
    /**
     * @generated from protobuf rpc: List(gkit.service.account.user.v1.ListRequest) returns (gkit.service.account.user.v1.ListResponse);
     */
    list(input: ListRequest, options?: RpcOptions): UnaryCall<ListRequest, ListResponse> {
        const method = this.methods[1], opt = this._transport.mergeOptions(options);
        return stackIntercept<ListRequest, ListResponse>("unary", this._transport, method, opt, input);
    }
    /**
     * @generated from protobuf rpc: Get(gkit.service.account.user.v1.GetRequest) returns (gkit.service.account.user.v1.GetResponse);
     */
    get(input: GetRequest, options?: RpcOptions): UnaryCall<GetRequest, GetResponse> {
        const method = this.methods[2], opt = this._transport.mergeOptions(options);
        return stackIntercept<GetRequest, GetResponse>("unary", this._transport, method, opt, input);
    }
    /**
     * @generated from protobuf rpc: Create(gkit.service.account.user.v1.CreateRequest) returns (gkit.service.account.user.v1.CreateResponse);
     */
    create(input: CreateRequest, options?: RpcOptions): UnaryCall<CreateRequest, CreateResponse> {
        const method = this.methods[3], opt = this._transport.mergeOptions(options);
        return stackIntercept<CreateRequest, CreateResponse>("unary", this._transport, method, opt, input);
    }
    /**
     * @generated from protobuf rpc: Update(gkit.service.account.user.v1.UpdateRequest) returns (gkit.service.account.user.v1.UpdateResponse);
     */
    update(input: UpdateRequest, options?: RpcOptions): UnaryCall<UpdateRequest, UpdateResponse> {
        const method = this.methods[4], opt = this._transport.mergeOptions(options);
        return stackIntercept<UpdateRequest, UpdateResponse>("unary", this._transport, method, opt, input);
    }
    /**
     * @generated from protobuf rpc: Delete(gkit.service.account.user.v1.DeleteRequest) returns (gkit.service.account.user.v1.DeleteResponse);
     */
    delete(input: DeleteRequest, options?: RpcOptions): UnaryCall<DeleteRequest, DeleteResponse> {
        const method = this.methods[5], opt = this._transport.mergeOptions(options);
        return stackIntercept<DeleteRequest, DeleteResponse>("unary", this._transport, method, opt, input);
    }
}
