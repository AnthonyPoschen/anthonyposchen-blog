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

export class ServiceRequest extends jspb.Message {
  getName(): ServiceName;
  setName(value: ServiceName): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ServiceRequest.AsObject;
  static toObject(includeInstance: boolean, msg: ServiceRequest): ServiceRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: ServiceRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ServiceRequest;
  static deserializeBinaryFromReader(message: ServiceRequest, reader: jspb.BinaryReader): ServiceRequest;
}

export namespace ServiceRequest {
  export type AsObject = {
    name: ServiceName,
  }
}

export class ServiceResponse extends jspb.Message {
  getUrl(): string;
  setUrl(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ServiceResponse.AsObject;
  static toObject(includeInstance: boolean, msg: ServiceResponse): ServiceResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: ServiceResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ServiceResponse;
  static deserializeBinaryFromReader(message: ServiceResponse, reader: jspb.BinaryReader): ServiceResponse;
}

export namespace ServiceResponse {
  export type AsObject = {
    url: string,
  }
}

export enum ServiceName {
  NONE = 0,
  BLOG = 1,
  USER = 2,
}

