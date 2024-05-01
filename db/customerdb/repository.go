package customerdb

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/moaabb/payments_microservices/customer/models/entities"
	"github.com/moaabb/payments_microservices/customer/repository"
)

const DB_OPERATION_TIMEOUT = 10 * time.Second

type CustomerRepository struct {
	db *pgxpool.Pool
}

func NewCustomerRepository(conn *pgxpool.Pool) repository.CustomerRepository {
	return &CustomerRepository{
		db: conn,
	}
}

func (m *CustomerRepository) GetCustomers() ([]entities.Customer, error) {
	ctx, cancel := context.WithTimeout(context.Background(), DB_OPERATION_TIMEOUT)
	defer cancel()

	tx, err := m.db.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return nil, err
	}
	defer tx.Rollback(ctx)

	rows, err := tx.Query(ctx, GetCustomers)
	if err != nil {
		return nil, err
	}

	var customers []entities.Customer
	for rows.Next() {
		customer, err := m.mapCustomer(rows)
		if err != nil {
			return nil, err
		}

		customers = append(customers, *customer)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	// tx.Commit(ctx)

	return customers, nil
}
func (m *CustomerRepository) GetCustomerById(id int) (*entities.Customer, error) {
	ctx, cancel := context.WithTimeout(context.Background(), DB_OPERATION_TIMEOUT)
	defer cancel()

	tx, err := m.db.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return nil, err
	}
	defer tx.Rollback(ctx)

	row := tx.QueryRow(ctx, GetCustomerById, id)
	customer, err := m.mapCustomer(row)
	if err != nil {
		return nil, err
	}

	tx.Commit(ctx)

	return customer, nil
}

func (m *CustomerRepository) CreateCustomer(payload entities.Customer) (*entities.Customer, error) {
	ctx, cancel := context.WithTimeout(context.Background(), DB_OPERATION_TIMEOUT)
	defer cancel()

	tx, err := m.db.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return nil, err
	}
	defer tx.Rollback(ctx)

	row := tx.QueryRow(ctx, CreateCustomer,
		payload.Name,
		payload.BirthDate.Time,
		payload.Phone,
		payload.Email,
	)
	customer, err := m.mapCustomer(row)
	if err != nil {
		return nil, err
	}

	tx.Commit(ctx)

	return customer, nil
}
func (m *CustomerRepository) UpdateCustomer(customer entities.Customer) (*entities.Customer, error) {
	return nil, nil
}

func (m *CustomerRepository) mapCustomer(rows pgx.Row) (*entities.Customer, error) {
	var customer entities.Customer
	err := rows.Scan(
		&customer.CustomerId,
		&customer.Name,
		&customer.BirthDate.Time,
		&customer.Phone,
		&customer.Email,
	)
	if err != nil {
		return nil, err
	}

	return &customer, nil
}
