package db

import (
	"context"
	"fmt"
	"ucode/ucode_go_chat_service/config"

	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Postgres struct {
	Builder squirrel.StatementBuilderType
	Pg      *pgxpool.Pool
}

func New(ctx context.Context, cfg config.Config) (*Postgres, error) {
	pg := &Postgres{}

	pgxUrl := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		cfg.PostgresUser,
		cfg.PostgresPassword,
		cfg.PostgresHost,
		cfg.PostgresPort,
		cfg.PostgresDatabase,
	)

	config, err := pgxpool.ParseConfig(pgxUrl)
	if err != nil {
		return nil, err
	}

	pool, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		return nil, err
	}

	pg.Builder = squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)
	pg.Pg = pool

	return pg, nil
}

func (p *Postgres) Close() {
	if p.Pg != nil {
		p.Pg.Close()
	}
}
