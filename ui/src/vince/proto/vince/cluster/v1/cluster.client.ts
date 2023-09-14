// @generated by protobuf-ts 2.9.1 with parameter generate_dependencies
// @generated from protobuf file "vince/cluster/v1/cluster.proto" (package "v1", syntax proto3)
// tslint:disable
import type { RpcTransport } from "@protobuf-ts/runtime-rpc";
import type { ServiceInfo } from "@protobuf-ts/runtime-rpc";
import { Cluster } from "./cluster";
import type { GetClusterResponse } from "./cluster";
import type { GetClusterRequest } from "./cluster";
import { stackIntercept } from "@protobuf-ts/runtime-rpc";
import type { ApplyClusterResponse } from "./cluster";
import type { ApplyClusterRequest } from "./cluster";
import type { UnaryCall } from "@protobuf-ts/runtime-rpc";
import type { RpcOptions } from "@protobuf-ts/runtime-rpc";
/**
 * @generated from protobuf service v1.Cluster
 */
export interface IClusterClient {
    /**
     * @generated from protobuf rpc: ApplyCluster(v1.ApplyClusterRequest) returns (v1.ApplyClusterResponse);
     */
    applyCluster(input: ApplyClusterRequest, options?: RpcOptions): UnaryCall<ApplyClusterRequest, ApplyClusterResponse>;
    /**
     * @generated from protobuf rpc: GetCluster(v1.GetClusterRequest) returns (v1.GetClusterResponse);
     */
    getCluster(input: GetClusterRequest, options?: RpcOptions): UnaryCall<GetClusterRequest, GetClusterResponse>;
}
/**
 * @generated from protobuf service v1.Cluster
 */
export class ClusterClient implements IClusterClient, ServiceInfo {
    typeName = Cluster.typeName;
    methods = Cluster.methods;
    options = Cluster.options;
    constructor(private readonly _transport: RpcTransport) {
    }
    /**
     * @generated from protobuf rpc: ApplyCluster(v1.ApplyClusterRequest) returns (v1.ApplyClusterResponse);
     */
    applyCluster(input: ApplyClusterRequest, options?: RpcOptions): UnaryCall<ApplyClusterRequest, ApplyClusterResponse> {
        const method = this.methods[0], opt = this._transport.mergeOptions(options);
        return stackIntercept<ApplyClusterRequest, ApplyClusterResponse>("unary", this._transport, method, opt, input);
    }
    /**
     * @generated from protobuf rpc: GetCluster(v1.GetClusterRequest) returns (v1.GetClusterResponse);
     */
    getCluster(input: GetClusterRequest, options?: RpcOptions): UnaryCall<GetClusterRequest, GetClusterResponse> {
        const method = this.methods[1], opt = this._transport.mergeOptions(options);
        return stackIntercept<GetClusterRequest, GetClusterResponse>("unary", this._transport, method, opt, input);
    }
}