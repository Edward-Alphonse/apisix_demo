namespace go monorepo.rpc.hello_world


struct GetHelloWorldRequest {
    1: required string name
}

struct GetHelloWorldResponse {
    1: string result
}

service HelloWorldService {
    GetHelloWorldResponse getHelloWorld(1: GetHelloWorldRequest req)
}