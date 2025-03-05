package main

import (
	"context"
	"fmt"
	hello_world "monorepo/service/rpc/hello_word/kitex_gen/monorepo/rpc/hello_world"
)

// HelloWorldServiceImpl implements the last service interface defined in the IDL.
type HelloWorldServiceImpl struct{}

// GetHelloWorld implements the HelloWorldServiceImpl interface.
func (s *HelloWorldServiceImpl) GetHelloWorld(ctx context.Context, req *hello_world.GetHelloWorldRequest) (resp *hello_world.GetHelloWorldResponse, err error) {
	resp = hello_world.NewGetHelloWorldResponse()
	resp.Result_ = fmt.Sprintf("hello world: %v", req.Name)
	return
}
