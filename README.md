# grpc-go-example
gRPC以及gRPC Gateway的Golang示例

安装相关命令行工具：
```
go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway

go get -u github.com/golang/protobuf/protoc-gen-go

go get -u github.com/gogo/protobuf/protoc-gen-gofast
```

在`/src/pb`目录中执行以下命令生成`*.pb.go`文件：
```
protoc --proto_path=./ -I $GOPATH/src -I $GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis -I $GOPATH/src/github.com/gogo/protobuf/protobuf --gofast_out=plugins=grpc:. *.proto
```

在`/src/pb`目录中执行以下命令生成`*.pb.gw.go`文件：
```
protoc --proto_path=./ -I $GOPATH/src -I $GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --grpc-gateway_out=logtostderr=true:. *.proto
```

示例功能：
[x] gRPC Server
[x] gRPC Client
[x] 通过gofast插件生成Golang代码，实现自定义结构体的标签
[x] gRPC Gateway将JSON与gRPC互相转换
[x] 自定义gRPC Gateway的错误输出
[] 整合zerolog
[] 实现方法调用时的身份验证