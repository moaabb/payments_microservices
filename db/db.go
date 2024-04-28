package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/moaabb/payments_microservices/customer/logger"
	"go.uber.org/zap"
)

var log = logger.GetLogger()

func ConnectToDatabase(dsn string) *pgxpool.Pool {
	dbpool, err := pgxpool.New(context.Background(), dsn)
	handleError(err)

	err = dbpool.Ping(context.Background())
	handleError(err)

	return dbpool
}

func handleError(err error) {
	if err != nil {
		log.Fatal("Unable to create connection pool", zap.Error(err))
	}
}
