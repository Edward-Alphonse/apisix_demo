### apisix 网关

#### 本地编译启动 debug
需要设置如下环境变量
```
APISIX_LISTEN_ADDRESS=unix:/tmp/runner.sock
APISIX_CONF_EXPIRE_TIME=3600
```

#### 开发环境
通过启动 docker-compose.yaml 作为本地的开发环境

#### 增加插件
参考如下文档：https://www.missquiet.com/detail?id=21