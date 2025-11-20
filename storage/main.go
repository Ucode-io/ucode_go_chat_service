package storage

import (
	"ucode/ucode_go_chat_service/config"
	"ucode/ucode_go_chat_service/pkg/db"
	"ucode/ucode_go_chat_service/pkg/logger"
	"ucode/ucode_go_chat_service/storage/postgres"
)

type StorageI interface {
	Postgres() postgres.PostgresI
}

type StoragePg struct {
	postgres postgres.PostgresI
}

func New(db *db.Postgres, log *logger.Logger, cfg config.Config) StorageI {
	return &StoragePg{
		postgres: postgres.New(db, log, cfg),
	}
}

func (s *StoragePg) Postgres() postgres.PostgresI {
	return s.postgres
}
