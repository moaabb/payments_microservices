package entities

import (
	"encoding/json"
	"fmt"
	"time"
)

var DateFormat = time.DateOnly
var DateTimeISOFormat = "2006-01-02T15:04:05.999Z"

type Customer struct {
	CustomerId *int    `json:"customerId,omitempty"`
	Name       *string `json:"name" validate:"required,min=6,max=30"`
	BirthDate  Date    `json:"birthDate"`
	Email      *string `json:"email" validate:"required,email"`
	Phone      *string `json:"phone" validate:"required,min=11,max=11"`
}

func NewCustomer(id int, name string, birthDate Date, email string, phone string) *Customer {
	return &Customer{
		CustomerId: &id,
		Name:       &name,
		BirthDate:  birthDate,
		Email:      &email,
		Phone:      &phone,
	}
}

func (m *Customer) ToString() string {
	out, _ := json.Marshal(m)

	return string(out)
}

type Date struct {
	time.Time
}

func (t *Date) UnmarshalJSON(b []byte) (err error) {
	if string(b) == "null" {
		return nil
	}
	date, err := time.Parse(`"`+DateFormat+`"`, string(b))
	if err != nil {
		return err
	}
	t.Time = date
	return nil
}

func (t *Date) MarshalJSON() ([]byte, error) {
	formatted := fmt.Sprintf(`"%s"`, t.Time.Format(DateFormat))
	return []byte(formatted), nil
}

type DateTime struct {
	time.Time
}

func (t *DateTime) UnmarshalJSON(b []byte) (err error) {
	if string(b) == "null" {
		return nil
	}
	date, err := time.Parse(`"`+DateTimeISOFormat+`"`, string(b))
	if err != nil {
		return err
	}
	t.Time = date
	return nil
}

func (t *DateTime) MarshalJSON() ([]byte, error) {
	formatted := fmt.Sprintf(`"%s"`, t.Time.Format(DateTimeISOFormat))
	return []byte(formatted), nil
}
