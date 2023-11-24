package main

import (
	"fmt"
	"github.com/ziggsdil/microservice/internal/config"
	"log/slog"
)

func main() {
	cfg := config.MustLoad()

	fmt.Println(cfg)

	// TODO: инициализировать логгер

	// TODO: инициализировать приложение (app)

	// TODO: запустить gRPC-сервер приложения
}

func setupLogger(env string) *slog.Logger {

}
