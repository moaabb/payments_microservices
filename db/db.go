package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	logging "github.com/moaabb/payments_microservices/customer/logger"
)

var logger *logging.ApplicationLogger

func ConnectToDatabase(dsn string) *pgxpool.Pool {
	logger = logging.GetLogger()

	dbpool, err := pgxpool.New(context.Background(), dsn)
	handleError(err)

	err = dbpool.Ping(context.Background())
	handleError(err)

	return dbpool
}

func handleError(err error) {
	if err != nil {
		logger.Fatalf("unable to create connection pool: %s", err.Error())
	}
}
