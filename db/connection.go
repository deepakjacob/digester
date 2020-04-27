package db

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/rs/zerolog/log"
)

const (
	hostname     = "localhost"
	host_port    = 5432
	username     = "postgres"
	password     = "postgres"
	databasename = "digester"
)

type PgConn struct {
	db *pgxpool.Pool
}

func New(ctx context.Context) *PgConn {
	pgConnString := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		hostname, host_port, username, password, databasename,
	)
	poolConfig, err := pgxpool.ParseConfig(pgConnString)
	if err != nil {
		log.Fatal().Err(err).Msg("Unable to parse database url")
		os.Exit(1)
	}

	db, err := pgxpool.ConnectConfig(context.Background(), poolConfig)
	if err != nil {
		log.Fatal().Err(err).Msg("Unable to create connection pool")
		os.Exit(1)
	}
	return &PgConn{db}
}
