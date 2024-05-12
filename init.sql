CREATE SCHEMA IF NOT EXISTS cus;

CREATE TABLE IF NOT EXISTS cus.customers (
    customer_id BIGSERIAL PRIMARY KEY,
    name varchar(255),
    birth_date timestamp,
    email varchar(255),
    phone varchar (11)
);

INSERT INTO cus.customers(name, birth_date, email, phone) VALUES
('Luigi Cogumelo', to_date('1998-06-12', 'YYYY-MM-DD'), 'luigi@email.com', '77991507621'),
('Mario Cogumelo', to_date('1995-02-08', 'YYYY-MM-DD'), 'mario@email.com', '77991230966');