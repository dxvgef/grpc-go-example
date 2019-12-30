package main

import (
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
	svr := grpc.NewServer()

	// 注册service.User服务
	pb.RegisterUserServer(svr, &action.User{
		MetaData: &action.MetaData{
			TokenStr: "",
		},
	})

	// 启用反射服务，允许客户端查询本实例提供的服务和方法
	// reflection.Register(svr)

	// 启动grpc服务
	log.Println("启动GRPC Server")
	if err = svr.Serve(lis); err != nil {
		log.Fatalln(err.Error())
	}
}
