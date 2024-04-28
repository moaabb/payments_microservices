package customerdb

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/moaabb/payments_microservices/customer/repository"
)

type CustomerRepository struct {
	db *pgxpool.Pool
}

func NewCustomerRepository(conn *pgxpool.Pool) repository.CustomerRepository {
	return &CustomerRepository{
		db: conn,
	}
}
