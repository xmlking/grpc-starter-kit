// @generated by protobuf-ts 2.4.0 with parameter long_type_string,generate_dependencies,// @generated from protobuf file "gkit/schema/email/v1/email_message.proto" (package "gkit.schema.email.v1", syntax proto3),// tslint:disable
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
// import "third_party/proto/validate/validate.proto";

/**
 * @generated from protobuf message gkit.schema.email.v1.Message
 */
export interface Message {
    /**
     * @generated from protobuf field: string to = 1;
     */
    to: string;
    /**
     * @generated from protobuf field: string from = 2;
     */
    from: string;
    /**
     * @generated from protobuf field: string subject = 3;
     */
    subject: string;
    /**
     * @generated from protobuf field: string body = 4;
     */
    body: string;
}
// @generated message type with reflection information, may provide speed optimized methods
class Message$Type extends MessageType<Message> {
    constructor() {
        super("gkit.schema.email.v1.Message", [
            { no: 1, name: "to", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 2, name: "from", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 3, name: "subject", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 4, name: "body", kind: "scalar", T: 9 /*ScalarType.STRING*/ }
        ]);
    }
    create(value?: PartialMessage<Message>): Message {
        const message = { to: "", from: "", subject: "", body: "" };
        globalThis.Object.defineProperty(message, MESSAGE_TYPE, { enumerable: false, value: this });
        if (value !== undefined)
            reflectionMergePartial<Message>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: Message): Message {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* string to */ 1:
                    message.to = reader.string();
                    break;
                case /* string from */ 2:
                    message.from = reader.string();
                    break;
                case /* string subject */ 3:
                    message.subject = reader.string();
                    break;
                case /* string body */ 4:
                    message.body = reader.string();
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
    internalBinaryWrite(message: Message, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* string to = 1; */
        if (message.to !== "")
            writer.tag(1, WireType.LengthDelimited).string(message.to);
        /* string from = 2; */
        if (message.from !== "")
            writer.tag(2, WireType.LengthDelimited).string(message.from);
        /* string subject = 3; */
        if (message.subject !== "")
            writer.tag(3, WireType.LengthDelimited).string(message.subject);
        /* string body = 4; */
        if (message.body !== "")
            writer.tag(4, WireType.LengthDelimited).string(message.body);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message gkit.schema.email.v1.Message
 */
export const Message = new Message$Type();
