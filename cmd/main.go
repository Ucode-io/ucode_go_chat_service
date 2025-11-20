package main

import (
	"context"
	"ucode/ucode_go_chat_service/api"
	"ucode/ucode_go_chat_service/config"
	"ucode/ucode_go_chat_service/pkg/db"
	"ucode/ucode_go_chat_service/pkg/logger"
	"ucode/ucode_go_chat_service/storage"
)

func main() {
	cfg := config.Load()
	logger := logger.New(cfg.LogLevel)

	postgres, err := db.New(context.Background(), cfg)
	if err != nil {
		logger.Error("error while connecting to postgresql")
		return
	}

	storage := storage.New(postgres, logger, cfg)

	engine := api.New(logger, cfg, storage)

	err = engine.Run(cfg.HTTPPort)
	if err != nil {
		logger.Error("error while running rest server")
		return
	}
}
