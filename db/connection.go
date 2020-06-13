package db

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx"
	"github.com/jackc/pgx/v4/log/zerologadapter"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/rs/zerolog"
)

const (
	hostName     = "localhost"
	hostPort     = 5432
	userName     = "postgres"
	password     = "postgres"
	databaseName = "digester"
)

// PgConn pool struct
type PgConn struct {
	db *pgxpool.Pool
}

// New postgres connection returned
func New(ctx context.Context) *PgConn {
	pgConnString := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		hostName, hostPort, userName, password, databaseName,
	)
	poolConfig, err := pgxpool.ParseConfig(pgConnString)
	if err != nil {
		// log.Fatal().Err(err).Msg("Unable to parse database url")
		os.Exit(1)
	}
	poolConfig.ConnConfig.LogLevel = pgx.LogLevelTrace
	poolConfig.ConnConfig.Logger = zerologadapter.NewLogger(zerolog.New(zerolog.NewConsoleWriter()))
	db, err := pgxpool.ConnectConfig(context.Background(), poolConfig)
	if err != nil {
		// log.Fatal().Err(err).Msg("Unable to create connection pool")
		os.Exit(1)
	}
	return &PgConn{db}
}
