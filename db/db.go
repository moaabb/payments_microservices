package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	logging "github.com/moaabb/payments_microservices/customer/logger"
)

var log = logging.GetLogger()

func ConnectToDatabase(dsn string) *pgxpool.Pool {
	dbpool, err := pgxpool.New(context.Background(), dsn)
	handleError(err)

	err = dbpool.Ping(context.Background())
	handleError(err)

	return dbpool
}

func handleError(err error) {
	if err != nil {
		log.Fatalf("unable to create connection pool: %s", err.Error())
	}
}
