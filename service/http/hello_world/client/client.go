package client

import (
	"log"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/discovery"
	etcd "github.com/kitex-contrib/registry-etcd"

	"monorepo/service/rpc/hello_word/kitex_gen/monorepo/rpc/hello_world/helloworldservice"
)

var (
	resolver         discovery.Resolver
	HelloWorldClient helloworldservice.Client
)

func Init(endpoints []string) {
	InitResolver(endpoints)
	initHelloWorldClient()
}

func InitResolver(endpoints []string) {
	var err error
	resolver, err = etcd.NewEtcdResolver(endpoints)
	if err != nil {
		log.Fatal(err)
	}
}

func initHelloWorldClient() {
	cli, err := helloworldservice.NewClient("monorepo.rpc.hello_world", client.WithResolver(resolver))
	if err != nil {
		log.Panicf("failed to new client: %s", err)
		return
	}
	HelloWorldClient = cli
}
