package main

import (
	"log"
	"os"
	"time"

	"local/action"
)

func main() {
	log.SetFlags(log.Lshortfile)
	log.SetOutput(os.Stdout)

	if err := action.SetDatabase(); err != nil {
		log.Println(err.Error())
		return
	}

	// 新协程中启动grpc server
	go startGRpcServer()

	//  5秒钟后在新协程中启动grpc client
	go func() {
		log.Println("5秒钟后自动启动gGPC Client...")
		time.Sleep(5 * time.Second)
		startGrpcClient()
		log.Println("示例完成，退出进程")
		os.Exit(0)
	}()

	// 启动grpc gateway server
	startGrpcGatewayServer()

}
