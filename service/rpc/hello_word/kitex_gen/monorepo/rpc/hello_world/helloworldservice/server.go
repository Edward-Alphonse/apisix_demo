// Code generated by Kitex v0.11.3. DO NOT EDIT.
package helloworldservice

import (
	server "github.com/cloudwego/kitex/server"
	hello_world "monorepo/service/rpc/hello_word/kitex_gen/monorepo/rpc/hello_world"
)

// NewServer creates a server.Server with the given handler and options.
func NewServer(handler hello_world.HelloWorldService, opts ...server.Option) server.Server {
	var options []server.Option

	options = append(options, opts...)
	options = append(options, server.WithCompatibleMiddlewareForUnary())

	svr := server.NewServer(options...)
	if err := svr.RegisterService(serviceInfo(), handler); err != nil {
		panic(err)
	}
	return svr
}

func RegisterService(svr server.Server, handler hello_world.HelloWorldService, opts ...server.RegisterOption) error {
	return svr.RegisterService(serviceInfo(), handler, opts...)
}
