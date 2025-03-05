package main

import (
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
	"log"
	"monorepo/service/rpc/hello_word/kitex_gen/monorepo/rpc/hello_world/helloworldservice"
)

func main() {
	r, err := etcd.NewEtcdRegistry([]string{"127.0.0.1:2379"})
	if err != nil {
		panic(err)
	}

	svr := helloworldservice.NewServer(new(HelloWorldServiceImpl),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: "monorepo.rpc.hello_world"}), // 指定服务名称
		server.WithRegistry(r),
	)

	if err = svr.Run(); err != nil {
		log.Println(err.Error())
	}
}
