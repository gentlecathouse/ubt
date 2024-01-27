// @generated by protobuf-ts 2.9.0 with parameter long_type_bigint,server_generic
// @generated from protobuf file "services/block.proto" (package "ubt.services", syntax proto3)
// tslint:disable
import { Account } from "../models";
import { Block } from "../models";
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
import { FinalityStatus } from "../models";
import { ChainId } from "../models";
/**
 * @generated from protobuf message ubt.services.BlockRequest
 */
export interface BlockRequest {
    /**
     * @generated from protobuf field: ubt.ChainId chain_id = 1;
     */
    chainId?: ChainId;
    /**
     * @generated from protobuf field: bytes id = 2;
     */
    id: Uint8Array;
}
/**
 * @generated from protobuf message ubt.services.ListBlocksRequest
 */
export interface ListBlocksRequest {
    /**
     * @generated from protobuf field: ubt.ChainId chain_id = 1;
     */
    chainId?: ChainId;
    /**
     * @generated from protobuf field: uint64 start_number = 2;
     */
    startNumber: bigint;
    /**
     * @generated from protobuf field: optional uint64 count = 3;
     */
    count?: bigint;
    /**
     * @generated from protobuf field: uint32 includes = 4;
     */
    includes: number; // 1 - transactions, 2 - receipts, 4 - operations
    /**
     * @generated from protobuf field: ubt.FinalityStatus finality_status = 5;
     */
    finalityStatus: FinalityStatus; // return only blocks with that or greater finalization status
}
/**
 * @generated from protobuf enum ubt.services.ListBlocksRequest.IncludeFlags
 */
export enum ListBlocksRequest_IncludeFlags {
    /**
     * include only block header
     *
     * @generated from protobuf enum value: HEADER = 0;
     */
    HEADER = 0,
    /**
     * include transactions with header
     *
     * @generated from protobuf enum value: TRANSACTIONS = 1;
     */
    TRANSACTIONS = 1,
    /**
     * include full available data
     *
     * @generated from protobuf enum value: FULL = 4;
     */
    FULL = 4
}
/**
 * @generated from protobuf message ubt.services.GetAccountRequest
 */
export interface GetAccountRequest {
    /**
     * @generated from protobuf field: ubt.ChainId chain_id = 1;
     */
    chainId?: ChainId;
    /**
     * @generated from protobuf field: string address = 2;
     */
    address: string;
}
/**
 * @generated from protobuf message ubt.services.DeriveAccountRequest
 */
export interface DeriveAccountRequest {
    /**
     * @generated from protobuf field: ubt.ChainId chain_id = 1;
     */
    chainId?: ChainId;
    /**
     * @generated from protobuf field: bytes public_key = 2;
     */
    publicKey: Uint8Array;
}
// @generated message type with reflection information, may provide speed optimized methods
class BlockRequest$Type extends MessageType<BlockRequest> {
    constructor() {
        super("ubt.services.BlockRequest", [
            { no: 1, name: "chain_id", kind: "message", T: () => ChainId },
            { no: 2, name: "id", kind: "scalar", T: 12 /*ScalarType.BYTES*/ }
        ]);
    }
    create(value?: PartialMessage<BlockRequest>): BlockRequest {
        const message = { id: new Uint8Array(0) };
        globalThis.Object.defineProperty(message, MESSAGE_TYPE, { enumerable: false, value: this });
        if (value !== undefined)
            reflectionMergePartial<BlockRequest>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: BlockRequest): BlockRequest {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* ubt.ChainId chain_id */ 1:
                    message.chainId = ChainId.internalBinaryRead(reader, reader.uint32(), options, message.chainId);
                    break;
                case /* bytes id */ 2:
                    message.id = reader.bytes();
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
    internalBinaryWrite(message: BlockRequest, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* ubt.ChainId chain_id = 1; */
        if (message.chainId)
            ChainId.internalBinaryWrite(message.chainId, writer.tag(1, WireType.LengthDelimited).fork(), options).join();
        /* bytes id = 2; */
        if (message.id.length)
            writer.tag(2, WireType.LengthDelimited).bytes(message.id);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message ubt.services.BlockRequest
 */
export const BlockRequest = new BlockRequest$Type();
// @generated message type with reflection information, may provide speed optimized methods
class ListBlocksRequest$Type extends MessageType<ListBlocksRequest> {
    constructor() {
        super("ubt.services.ListBlocksRequest", [
            { no: 1, name: "chain_id", kind: "message", T: () => ChainId },
            { no: 2, name: "start_number", kind: "scalar", T: 4 /*ScalarType.UINT64*/, L: 0 /*LongType.BIGINT*/ },
            { no: 3, name: "count", kind: "scalar", opt: true, T: 4 /*ScalarType.UINT64*/, L: 0 /*LongType.BIGINT*/ },
            { no: 4, name: "includes", kind: "scalar", T: 13 /*ScalarType.UINT32*/ },
            { no: 5, name: "finality_status", kind: "enum", T: () => ["ubt.FinalityStatus", FinalityStatus, "FINALITY_STATUS_"] }
        ]);
    }
    create(value?: PartialMessage<ListBlocksRequest>): ListBlocksRequest {
        const message = { startNumber: 0n, includes: 0, finalityStatus: 0 };
        globalThis.Object.defineProperty(message, MESSAGE_TYPE, { enumerable: false, value: this });
        if (value !== undefined)
            reflectionMergePartial<ListBlocksRequest>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: ListBlocksRequest): ListBlocksRequest {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* ubt.ChainId chain_id */ 1:
                    message.chainId = ChainId.internalBinaryRead(reader, reader.uint32(), options, message.chainId);
                    break;
                case /* uint64 start_number */ 2:
                    message.startNumber = reader.uint64().toBigInt();
                    break;
                case /* optional uint64 count */ 3:
                    message.count = reader.uint64().toBigInt();
                    break;
                case /* uint32 includes */ 4:
                    message.includes = reader.uint32();
                    break;
                case /* ubt.FinalityStatus finality_status */ 5:
                    message.finalityStatus = reader.int32();
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
    internalBinaryWrite(message: ListBlocksRequest, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* ubt.ChainId chain_id = 1; */
        if (message.chainId)
            ChainId.internalBinaryWrite(message.chainId, writer.tag(1, WireType.LengthDelimited).fork(), options).join();
        /* uint64 start_number = 2; */
        if (message.startNumber !== 0n)
            writer.tag(2, WireType.Varint).uint64(message.startNumber);
        /* optional uint64 count = 3; */
        if (message.count !== undefined)
            writer.tag(3, WireType.Varint).uint64(message.count);
        /* uint32 includes = 4; */
        if (message.includes !== 0)
            writer.tag(4, WireType.Varint).uint32(message.includes);
        /* ubt.FinalityStatus finality_status = 5; */
        if (message.finalityStatus !== 0)
            writer.tag(5, WireType.Varint).int32(message.finalityStatus);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message ubt.services.ListBlocksRequest
 */
export const ListBlocksRequest = new ListBlocksRequest$Type();
// @generated message type with reflection information, may provide speed optimized methods
class GetAccountRequest$Type extends MessageType<GetAccountRequest> {
    constructor() {
        super("ubt.services.GetAccountRequest", [
            { no: 1, name: "chain_id", kind: "message", T: () => ChainId },
            { no: 2, name: "address", kind: "scalar", T: 9 /*ScalarType.STRING*/ }
        ]);
    }
    create(value?: PartialMessage<GetAccountRequest>): GetAccountRequest {
        const message = { address: "" };
        globalThis.Object.defineProperty(message, MESSAGE_TYPE, { enumerable: false, value: this });
        if (value !== undefined)
            reflectionMergePartial<GetAccountRequest>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: GetAccountRequest): GetAccountRequest {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* ubt.ChainId chain_id */ 1:
                    message.chainId = ChainId.internalBinaryRead(reader, reader.uint32(), options, message.chainId);
                    break;
                case /* string address */ 2:
                    message.address = reader.string();
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
    internalBinaryWrite(message: GetAccountRequest, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* ubt.ChainId chain_id = 1; */
        if (message.chainId)
            ChainId.internalBinaryWrite(message.chainId, writer.tag(1, WireType.LengthDelimited).fork(), options).join();
        /* string address = 2; */
        if (message.address !== "")
            writer.tag(2, WireType.LengthDelimited).string(message.address);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message ubt.services.GetAccountRequest
 */
export const GetAccountRequest = new GetAccountRequest$Type();
// @generated message type with reflection information, may provide speed optimized methods
class DeriveAccountRequest$Type extends MessageType<DeriveAccountRequest> {
    constructor() {
        super("ubt.services.DeriveAccountRequest", [
            { no: 1, name: "chain_id", kind: "message", T: () => ChainId },
            { no: 2, name: "public_key", kind: "scalar", T: 12 /*ScalarType.BYTES*/ }
        ]);
    }
    create(value?: PartialMessage<DeriveAccountRequest>): DeriveAccountRequest {
        const message = { publicKey: new Uint8Array(0) };
        globalThis.Object.defineProperty(message, MESSAGE_TYPE, { enumerable: false, value: this });
        if (value !== undefined)
            reflectionMergePartial<DeriveAccountRequest>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: DeriveAccountRequest): DeriveAccountRequest {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* ubt.ChainId chain_id */ 1:
                    message.chainId = ChainId.internalBinaryRead(reader, reader.uint32(), options, message.chainId);
                    break;
                case /* bytes public_key */ 2:
                    message.publicKey = reader.bytes();
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
    internalBinaryWrite(message: DeriveAccountRequest, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* ubt.ChainId chain_id = 1; */
        if (message.chainId)
            ChainId.internalBinaryWrite(message.chainId, writer.tag(1, WireType.LengthDelimited).fork(), options).join();
        /* bytes public_key = 2; */
        if (message.publicKey.length)
            writer.tag(2, WireType.LengthDelimited).bytes(message.publicKey);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message ubt.services.DeriveAccountRequest
 */
export const DeriveAccountRequest = new DeriveAccountRequest$Type();
/**
 * @generated ServiceType for protobuf service ubt.services.UbtBlockService
 */
export const UbtBlockService = new ServiceType("ubt.services.UbtBlockService", [
    { name: "getBlock", options: {}, I: BlockRequest, O: Block },
    { name: "listBlocks", serverStreaming: true, options: {}, I: ListBlocksRequest, O: Block },
    { name: "getAccount", options: {}, I: GetAccountRequest, O: Account },
    { name: "deriveAccount", options: {}, I: DeriveAccountRequest, O: Account }
]);
