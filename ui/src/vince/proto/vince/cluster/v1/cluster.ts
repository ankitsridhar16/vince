// @generated by protobuf-ts 2.9.1 with parameter generate_dependencies
// @generated from protobuf file "vince/cluster/v1/cluster.proto" (package "v1", syntax proto3)
// tslint:disable
import { ServiceType } from "@protobuf-ts/runtime-rpc";
import type { BinaryWriteOptions } from "@protobuf-ts/runtime";
import type { IBinaryWriter } from "@protobuf-ts/runtime";
import { WireType } from "@protobuf-ts/runtime";
import type { BinaryReadOptions } from "@protobuf-ts/runtime";
import type { IBinaryReader } from "@protobuf-ts/runtime";
import { UnknownFieldHandler } from "@protobuf-ts/runtime";
import type { PartialMessage } from "@protobuf-ts/runtime";
import { reflectionMergePartial } from "@protobuf-ts/runtime";
import { MESSAGE_TYPE } from "@protobuf-ts/runtime";
import { MessageType } from "@protobuf-ts/runtime";
import { ClusterConfig } from "../../config/v1/config";
/**
 * @generated from protobuf message v1.ApplyClusterRequest
 */
export interface ApplyClusterRequest {
    /**
     * @generated from protobuf field: v1.ClusterConfig config = 1;
     */
    config?: ClusterConfig;
}
/**
 * @generated from protobuf message v1.ApplyClusterResponse
 */
export interface ApplyClusterResponse {
    /**
     * @generated from protobuf field: string ok = 1;
     */
    ok: string;
}
/**
 * @generated from protobuf message v1.GetClusterRequest
 */
export interface GetClusterRequest {
}
/**
 * @generated from protobuf message v1.GetClusterResponse
 */
export interface GetClusterResponse {
    /**
     * @generated from protobuf field: v1.ClusterConfig config = 1;
     */
    config?: ClusterConfig;
}
// @generated message type with reflection information, may provide speed optimized methods
class ApplyClusterRequest$Type extends MessageType<ApplyClusterRequest> {
    constructor() {
        super("v1.ApplyClusterRequest", [
            { no: 1, name: "config", kind: "message", T: () => ClusterConfig }
        ]);
    }
    create(value?: PartialMessage<ApplyClusterRequest>): ApplyClusterRequest {
        const message = {};
        globalThis.Object.defineProperty(message, MESSAGE_TYPE, { enumerable: false, value: this });
        if (value !== undefined)
            reflectionMergePartial<ApplyClusterRequest>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: ApplyClusterRequest): ApplyClusterRequest {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* v1.ClusterConfig config */ 1:
                    message.config = ClusterConfig.internalBinaryRead(reader, reader.uint32(), options, message.config);
                    break;
                default:
                    let u = options.readUnknownField;
                    if (u === "throw")
                        throw new globalThis.Error(`Unknown field ${fieldNo} (wire type ${wireType}) for ${this.typeName}`);
                    let d = reader.skip(wireType);
                    if (u !== false)
                        (u === true ? UnknownFieldHandler.onRead : u)(this.typeName, message, fieldNo, wireType, d);
            }
        }
        return message;
    }
    internalBinaryWrite(message: ApplyClusterRequest, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* v1.ClusterConfig config = 1; */
        if (message.config)
            ClusterConfig.internalBinaryWrite(message.config, writer.tag(1, WireType.LengthDelimited).fork(), options).join();
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message v1.ApplyClusterRequest
 */
export const ApplyClusterRequest = new ApplyClusterRequest$Type();
// @generated message type with reflection information, may provide speed optimized methods
class ApplyClusterResponse$Type extends MessageType<ApplyClusterResponse> {
    constructor() {
        super("v1.ApplyClusterResponse", [
            { no: 1, name: "ok", kind: "scalar", T: 9 /*ScalarType.STRING*/ }
        ]);
    }
    create(value?: PartialMessage<ApplyClusterResponse>): ApplyClusterResponse {
        const message = { ok: "" };
        globalThis.Object.defineProperty(message, MESSAGE_TYPE, { enumerable: false, value: this });
        if (value !== undefined)
            reflectionMergePartial<ApplyClusterResponse>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: ApplyClusterResponse): ApplyClusterResponse {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* string ok */ 1:
                    message.ok = reader.string();
                    break;
                default:
                    let u = options.readUnknownField;
                    if (u === "throw")
                        throw new globalThis.Error(`Unknown field ${fieldNo} (wire type ${wireType}) for ${this.typeName}`);
                    let d = reader.skip(wireType);
                    if (u !== false)
                        (u === true ? UnknownFieldHandler.onRead : u)(this.typeName, message, fieldNo, wireType, d);
            }
        }
        return message;
    }
    internalBinaryWrite(message: ApplyClusterResponse, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* string ok = 1; */
        if (message.ok !== "")
            writer.tag(1, WireType.LengthDelimited).string(message.ok);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message v1.ApplyClusterResponse
 */
export const ApplyClusterResponse = new ApplyClusterResponse$Type();
// @generated message type with reflection information, may provide speed optimized methods
class GetClusterRequest$Type extends MessageType<GetClusterRequest> {
    constructor() {
        super("v1.GetClusterRequest", []);
    }
    create(value?: PartialMessage<GetClusterRequest>): GetClusterRequest {
        const message = {};
        globalThis.Object.defineProperty(message, MESSAGE_TYPE, { enumerable: false, value: this });
        if (value !== undefined)
            reflectionMergePartial<GetClusterRequest>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: GetClusterRequest): GetClusterRequest {
        return target ?? this.create();
    }
    internalBinaryWrite(message: GetClusterRequest, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message v1.GetClusterRequest
 */
export const GetClusterRequest = new GetClusterRequest$Type();
// @generated message type with reflection information, may provide speed optimized methods
class GetClusterResponse$Type extends MessageType<GetClusterResponse> {
    constructor() {
        super("v1.GetClusterResponse", [
            { no: 1, name: "config", kind: "message", T: () => ClusterConfig }
        ]);
    }
    create(value?: PartialMessage<GetClusterResponse>): GetClusterResponse {
        const message = {};
        globalThis.Object.defineProperty(message, MESSAGE_TYPE, { enumerable: false, value: this });
        if (value !== undefined)
            reflectionMergePartial<GetClusterResponse>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: GetClusterResponse): GetClusterResponse {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* v1.ClusterConfig config */ 1:
                    message.config = ClusterConfig.internalBinaryRead(reader, reader.uint32(), options, message.config);
                    break;
                default:
                    let u = options.readUnknownField;
                    if (u === "throw")
                        throw new globalThis.Error(`Unknown field ${fieldNo} (wire type ${wireType}) for ${this.typeName}`);
                    let d = reader.skip(wireType);
                    if (u !== false)
                        (u === true ? UnknownFieldHandler.onRead : u)(this.typeName, message, fieldNo, wireType, d);
            }
        }
        return message;
    }
    internalBinaryWrite(message: GetClusterResponse, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* v1.ClusterConfig config = 1; */
        if (message.config)
            ClusterConfig.internalBinaryWrite(message.config, writer.tag(1, WireType.LengthDelimited).fork(), options).join();
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message v1.GetClusterResponse
 */
export const GetClusterResponse = new GetClusterResponse$Type();
/**
 * @generated ServiceType for protobuf service v1.Cluster
 */
export const Cluster = new ServiceType("v1.Cluster", [
    { name: "ApplyCluster", options: { "google.api.http": { post: "/v1/cluster/apply" } }, I: ApplyClusterRequest, O: ApplyClusterResponse },
    { name: "GetCluster", options: { "google.api.http": { get: "/v1/cluster" } }, I: GetClusterRequest, O: GetClusterResponse }
]);