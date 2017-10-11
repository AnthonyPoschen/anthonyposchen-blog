// package: helloworld
// file: helloworld.proto

import * as helloworld_pb from "./helloworld_pb";
export class HelloService {
  static serviceName = "helloworld.HelloService";
}
export namespace HelloService {
  export class sayHello {
    static readonly methodName = "sayHello";
    static readonly service = HelloService;
    static readonly requestStream = false;
    static readonly responseStream = false;
    static readonly requestType = helloworld_pb.HelloRequest;
    static readonly responseType = helloworld_pb.HelloResponse;
  }
}
