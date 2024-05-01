package customerdb

var GetCustomers = `
	SELECT customer_id, name, birth_date, phone, email FROM cus.customer
`
var GetCustomerById = `
	SELECT customer_id, name, birth_date, phone, email FROM cus.customer WHERE customer_id = $1
`
var CreateCustomer = `
	INSERT INTO cus.customer (name, birth_date, phone, email) VALUES ($1, $2, $3, $4) RETURNING customer_id, name, birth_date, phone, email
`
