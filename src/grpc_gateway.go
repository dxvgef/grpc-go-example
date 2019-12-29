package main

import (
	"context"
	"log"
	"net/http"

	gw "github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	"local/pb"
)

type ResponseError struct {
	Error string `json:"error,omitempty"`
	// Code  int32  `json:"code,omitempty"`
}

func startGrpcGatewayServer() {
	// ------------------------- 启用grpc gateway server ---------------------------
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// 创建一个grpc gateway服务实例
	gwMux := gw.NewServeMux(
		// 将header中指定的参数append到metaDate
		gw.WithMetadata(func(ctx context.Context, req *http.Request) metadata.MD {
			md := make(map[string][]string)
			md["token"] = []string{req.Header.Get("token")}
			return md
		}),

		// 自定义grpc的错误处理
		gw.WithProtoErrorHandler(errorHandler),
	)
	opts := []grpc.DialOption{grpc.WithInsecure()}

	// 注册endpoint（上游的grpc服务的地址和端口）
	err := pb.RegisterUserHandlerFromEndpoint(ctx, gwMux, "127.0.0.1:3000", opts)
	if err != nil {
		log.Fatalln(err.Error())
		return
	}

	// 启动grpc gateway服务
	log.Println("启动GRPC Gateway Server")
	if err = http.ListenAndServe(":3002", gwMux); err != nil {
		log.Fatalln(err.Error())
		return
	}
}

// 错误响应处理器
func errorHandler(_ context.Context, _ *gw.ServeMux, marshaller gw.Marshaler, w http.ResponseWriter, _ *http.Request, err error) {
	s, _ := status.FromError(err)

	// 如果不是400错误，则不返回任何消息，可以减少json序列化的过程以及网络传输
	// 客户端可以根据状态码来决定输出什么消息
	if s.Code() != 400 {
		w.WriteHeader(gw.HTTPStatusFromCode(s.Code()))
		return
	}

	var respErr ResponseError
	// respErr.Code = s.Proto().GetCode()
	respErr.Error = s.Proto().GetMessage()

	jsonData, err := marshaller.Marshal(&respErr)
	if err != nil {
		w.WriteHeader(500)
		log.Println(err.Error())
		return
	}

	w.WriteHeader(gw.HTTPStatusFromCode(s.Code()))
	_, err = w.Write(jsonData)
	if err != nil {
		w.WriteHeader(500)
		log.Println(err.Error())
		return
	}
}
