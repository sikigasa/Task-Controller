package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"

	"google.golang.org/grpc"
)

func main() {
	// 8080番portのListenerを作成
	port := 8080
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		panic(err)
	}

	// ctx, cancel := context.WithCancel(context.Background(), 5*time.Second)
	// defer cancel()

	// gRPCサーバーを作成
	s := grpc.NewServer()

	// 作成したgRPCサーバーを、8080番ポートで稼働させる
	go func() {
		log.Printf("start gRPC server port: %v", port)
		s.Serve(listener)
	}()

	// Ctrl+Cが入力されたらGraceful shutdownされるようにする
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("stopping gRPC server...")
	s.GracefulStop()
}
