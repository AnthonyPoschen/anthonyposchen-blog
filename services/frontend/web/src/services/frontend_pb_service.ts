// package: frontend
// file: frontend.proto

import * as frontend_pb from "./frontend_pb";
export class FrontendService {
  static serviceName = "frontend.FrontendService";
}
export namespace FrontendService {
  export class test {
    static readonly methodName = "test";
    static readonly service = FrontendService;
    static readonly requestStream = false;
    static readonly responseStream = false;
    static readonly requestType = frontend_pb.testMsg;
    static readonly responseType = frontend_pb.testMsg;
  }
}
