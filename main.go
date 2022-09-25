package main

import (
	"clean-architecture/constant"
	"clean-architecture/infra"
	"clean-architecture/server"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/viper"
	"google.golang.org/grpc/reflection"
)

func main() {
	// 環境変数読み込み
	loadEnv()
	// DB初期化処理
	gormHandler := infra.NewGormHandler()

	// GRPC接続
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGTERM, syscall.SIGINT)

	grpcServer := server.NewGRPCServer(gormHandler)
	go func() {
		<-sigCh
		log.Printf("shutdown signal is called")
		grpcServer.GracefulStop()
	}()

	listener, err := net.Listen("tcp", ":50051")

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("server start!")
	// サーバーリフレクションを有効にしています。
	// 有効にすることでシリアライズせずとも後述する`grpc_cli`で動作確認ができるようになります。
	reflection.Register(grpcServer)
	grpcServer.Serve(listener)
}

func loadEnv() {
	viper.AutomaticEnv()
	viper.SetDefault(constant.DBHostEnv, "127.0.0.1")
	viper.SetDefault(constant.DBPortEnv, "3306")
	viper.SetDefault(constant.DBUserEnv, "root")
	viper.SetDefault(constant.DBPasswordEnv, "")
	viper.SetDefault(constant.DBNameEnv, "dev")
}
