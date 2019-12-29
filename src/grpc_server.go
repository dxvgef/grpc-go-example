package main

import (
	"context"
	"log"
	"net"

	"local/action"
	"local/pb"

	"google.golang.org/grpc"
)

func startGRpcServer() {
	// 监听一个地址
	lis, err := net.Listen("tcp", ":3000")
	if err != nil {
		log.Fatalln(err.Error())
	}

	// 创建一个grpc服务的实例
	svr := grpc.NewServer(
		grpc.UnaryInterceptor(hook), // 身份验证拦截器，一个server只支持一个拦截器
	)

	// 注册service.User服务
	pb.RegisterUserServer(svr, &action.User{})

	// 启用反射服务，允许客户端查询本实例提供的服务和方法
	// reflection.Register(svr)

	// 启动grpc服务
	log.Println("启动GRPC Server")
	if err = svr.Serve(lis); err != nil {
		log.Fatalln(err.Error())
	}
}

// 身份验证拦截器
func hook(ctx context.Context, params interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	log.Println("客户端请求方法：" + info.FullMethod)
	log.Println("客户端请求参数：", params)
	// md, ok := metadata.FromIncomingContext(ctx)
	// if !ok {
	// 	return nil, status.Error(codes.Unauthenticated, "context不正确")
	// }
	// token := md.Get("token")
	// if len(token) == 0 || token[0] == "" {
	// 	return nil, status.Error(codes.Unauthenticated, "无法获得token")
	// }
	// log.Println(md)
	// // 继续处理请求
	return handler(ctx, params)
}
