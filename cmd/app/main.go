package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/sikigasa/task-controller/cmd/config"
	"github.com/sikigasa/task-controller/internal/infra"
	postgres "github.com/sikigasa/task-controller/internal/infra/driver"
	"github.com/sikigasa/task-controller/internal/middleware"
	"github.com/sikigasa/task-controller/internal/usecase"
)

func init() {
	config.LoadEnv(".env")
}

func main() {
	if err := run(); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}

func run() error {
	// 8080番portのListenerを作成
	port := 8080
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	conn, err := postgres.NewPostgresConnection(config.Config.Postgres.User, config.Config.Postgres.Password, config.Config.Postgres.Host, config.Config.Postgres.Port, config.Config.Postgres.DBName, config.Config.Postgres.SSLMode)
	if err != nil {
		panic(err)
	}
	db, err := conn.Connection()
	if err != nil {
		panic(err)
	}
	defer conn.Close(ctx)

	// Connect RPCサーバーを作成
	taskService := usecase.NewTaskService(infra.NewTaskRepo(db), infra.NewTagRepo(db), infra.NewTaskTagRepo(db), postgres.NewPostgresTransaction(db))
	tagService := usecase.NewTagService(infra.NewTagRepo(db))

	// Routerを使用してルートを設定
	router := middleware.NewRouter(taskService, tagService)
	mux := router.SetupRoutes()

	s := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: mux,
	}
	// 作成したhttpサーバーを、8080番ポートで稼働させる
	go func() {
		log.Printf("start http server port: %v", port)
		s.Serve(listener)
	}()

	// Ctrl+Cが入力されたらGraceful shutdownされるようにする
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("stopping http server...")
	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdownCancel()
	s.Shutdown(shutdownCtx)
	log.Println("http server stopped")

	return nil
}
