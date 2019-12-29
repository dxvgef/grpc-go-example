package main

import (
	"context"
	"errors"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	"local/pb"
)

func startGrpcClient() {
	// 连接服务端
	dialCtx, _ := context.WithTimeout(context.Background(), 3*time.Second)
	conn, err := grpc.DialContext(
		dialCtx,
		"127.0.0.1:3000",
		grpc.WithInsecure(),
		grpc.WithBlock(),
	)
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			log.Println("服务端连接超时")
		} else {
			log.Println(err.Error())
		}
		return
	}

	// 创建一个User实例
	c := pb.NewUserClient(conn)

	// 定义一个metaData，用于发送token
	ctx, _ := context.WithTimeout(context.Background(), 3*time.Second)

	// 执行服务端方法，并发送metaData
	loginResp, err := c.Login(ctx, &pb.LoginRequest{
		Username: "test",
		Password: "123456",
	})

	// 处理错误
	gErr := status.Convert(err)
	if gErr != nil {
		log.Println(gErr.Proto().GetCode(), gErr.Proto().GetMessage())
		return
	}

	log.Println("服务端响应：", loginResp.Token)

	metaData := metadata.NewOutgoingContext(ctx, metadata.Pairs("token", loginResp.Token))
	listResp, err := c.List(metaData, &pb.ListRequest{
		Columns:  []string{"id", "username"},
		Page:     1,
		PageSize: 10,
	})

	// 处理错误
	gErr = status.Convert(err)
	if gErr != nil {
		log.Println(gErr.Proto().GetCode(), gErr.Proto().GetMessage())
		return
	}

	log.Println("服务端响应：", listResp.Total)

	for k := range listResp.Rows {
		log.Println("服务端响应:", listResp.Rows[k].Id, listResp.Rows[k].Username, listResp.Rows[k].CreateTime)
	}
}
