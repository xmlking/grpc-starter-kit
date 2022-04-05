// @generated by protobuf-ts 2.4.0 with parameter long_type_string,generate_dependencies,enable_angular_annotations,// @generated from protobuf file "gkit/service/transaction/v1/transaction.proto" (package "gkit.service.transaction.v1", syntax proto3),// tslint:disable
import { Inject } from "@angular/core";
import { RPC_TRANSPORT } from "@protobuf-ts/runtime-angular";
import { Injectable } from "@angular/core";
import type { RpcTransport } from "@protobuf-ts/runtime-rpc";
import type { ServiceInfo } from "@protobuf-ts/runtime-rpc";
import { TransactionService } from "./transaction";
import type { Empty } from "../../../../google/protobuf/empty";
import type { WriteRequest } from "./transaction";
import { stackIntercept } from "@protobuf-ts/runtime-rpc";
import type { ReadResponse } from "./transaction";
import type { ReadRequest } from "./transaction";
import type { UnaryCall } from "@protobuf-ts/runtime-rpc";
import type { RpcOptions } from "@protobuf-ts/runtime-rpc";
/**
 * @generated from protobuf service gkit.service.transaction.v1.TransactionService
 */
export interface ITransactionServiceClient {
    /**
     * @generated from protobuf rpc: Read(gkit.service.transaction.v1.ReadRequest) returns (gkit.service.transaction.v1.ReadResponse);
     */
    read(input: ReadRequest, options?: RpcOptions): UnaryCall<ReadRequest, ReadResponse>;
    /**
     * @generated from protobuf rpc: Write(gkit.service.transaction.v1.WriteRequest) returns (google.protobuf.Empty);
     */
    write(input: WriteRequest, options?: RpcOptions): UnaryCall<WriteRequest, Empty>;
}
/**
 * @generated from protobuf service gkit.service.transaction.v1.TransactionService
 */
@Injectable()
export class TransactionServiceClient implements ITransactionServiceClient, ServiceInfo {
    typeName = TransactionService.typeName;
    methods = TransactionService.methods;
    options = TransactionService.options;
    constructor(
    @Inject(RPC_TRANSPORT)
    private readonly _transport: RpcTransport) {
    }
    /**
     * @generated from protobuf rpc: Read(gkit.service.transaction.v1.ReadRequest) returns (gkit.service.transaction.v1.ReadResponse);
     */
    read(input: ReadRequest, options?: RpcOptions): UnaryCall<ReadRequest, ReadResponse> {
        const method = this.methods[0], opt = this._transport.mergeOptions(options);
        return stackIntercept<ReadRequest, ReadResponse>("unary", this._transport, method, opt, input);
    }
    /**
     * @generated from protobuf rpc: Write(gkit.service.transaction.v1.WriteRequest) returns (google.protobuf.Empty);
     */
    write(input: WriteRequest, options?: RpcOptions): UnaryCall<WriteRequest, Empty> {
        const method = this.methods[1], opt = this._transport.mergeOptions(options);
        return stackIntercept<WriteRequest, Empty>("unary", this._transport, method, opt, input);
    }
}
