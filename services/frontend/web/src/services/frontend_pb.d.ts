// package: frontend
// file: frontend.proto

import * as jspb from "google-protobuf";

export class testMsg extends jspb.Message {
  getText(): string;
  setText(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): testMsg.AsObject;
  static toObject(includeInstance: boolean, msg: testMsg): testMsg.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: testMsg, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): testMsg;
  static deserializeBinaryFromReader(message: testMsg, reader: jspb.BinaryReader): testMsg;
}

export namespace testMsg {
  export type AsObject = {
    text: string,
  }
}

