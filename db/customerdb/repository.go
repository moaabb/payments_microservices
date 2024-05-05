package customerdb

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/moaabb/payments_microservices/customer/models/entities"
	repository "github.com/moaabb/payments_microservices/customer/usecase"
	"go.opentelemetry.io/otel/trace"
)

const DB_OPERATION_TIMEOUT = 10 * time.Second

type CustomerRepository struct {
	db     *pgxpool.Pool
	tracer trace.Tracer
}

func NewCustomerRepository(conn *pgxpool.Pool, t trace.Tracer) repository.CustomerRepository {
	return &CustomerRepository{
		db:     conn,
		tracer: t,
	}
}

func (m *CustomerRepository) GetCustomers(ctx_t context.Context) ([]entities.Customer, error) {
	_, span := m.tracer.Start(ctx_t, GetCustomers)
	defer span.End()

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

	tx.Commit(ctx)

	return customers, nil
}
func (m *CustomerRepository) GetCustomerById(ctx_t context.Context, id int) (*entities.Customer, error) {
	_, span := m.tracer.Start(ctx_t, GetCustomerById)
	defer span.End()
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

func (m *CustomerRepository) CreateCustomer(ctx_t context.Context, payload entities.Customer) (*entities.Customer, error) {
	_, span := m.tracer.Start(ctx_t, CreateCustomer)
	defer span.End()
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
func (m *CustomerRepository) UpdateCustomer(ctx_t context.Context, customer entities.Customer) (*entities.Customer, error) {
	_, span := m.tracer.Start(ctx_t, "CustomerService")
	defer span.End()
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
