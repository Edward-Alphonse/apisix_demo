###IDL更新
idl更新后，需要进入到 main.go所在路径下运行如下命令：
``` 
service: 服务名，上游调用创建client需要指定服务名
cd ./service/rpc/account_rpc
kitex -module monorepo -service monorepo.rpc.hello_world  idl/hello_world.thrift
```
参考：
https://www.cloudwego.io/zh/docs/kitex/getting-started/tutorial/