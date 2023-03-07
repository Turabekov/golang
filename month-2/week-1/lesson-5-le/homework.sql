
-- USERS 
CREATE TABLE users(
    id BIGSERIAL NOT NULL PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    phone_number VARCHAR(21) NOT NULL
)
;
INSERT INTO users (name, phone_number) VALUES ('Khumoyun', '93-379-11-10');
INSERT INTO users (name, phone_number) VALUES ('Shaxzod', '97-772-07-70');
INSERT INTO users (name, phone_number) VALUES ('Baxrom', '93-802-07-70');


-- CUSTOMERS 
CREATE TABLE customers(
    id BIGSERIAL NOT NULL PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    phone VARCHAR(21) NOT NULL
)
;
INSERT INTO customers (name, phone) VALUES ('Abdulhamid', '93-111-11-11');
INSERT INTO customers (name, phone) VALUES ('Johongir', '97-122-22-22');
INSERT INTO customers (name, phone) VALUES ('Otabek', '22-333-33-33');

-- COURIERS 
CREATE TABLE couriers(
    id BIGSERIAL NOT NULL PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    phone_number VARCHAR(21) NOT NULL
)
;
INSERT INTO couriers (name, phone_number) VALUES ('Maxmud', '93-333-23-32');
INSERT INTO couriers (name, phone_number) VALUES ('Jaloliddin', '97-433-43-43');
INSERT INTO couriers (name, phone_number) VALUES ('Tahsin', '22-444-55-33');

-- PRODUCTS 
CREATE TABLE products(
    id BIGSERIAL NOT NULL PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    price NUMERIC NOT NULL
)
;
INSERT INTO products (name, price) VALUES ('Redmi Note 9 pro', 2000000);
INSERT INTO products (name, price) VALUES ('Pomidor', 10000);
INSERT INTO products (name, price) VALUES ('Olma', 12000);


-- CATEGORIES 
CREATE TABLE categories(
    id BIGSERIAL NOT NULL PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    product_id BIGSERIAL REFERENCES products(id)
)
;
INSERT INTO categories (name, product_id) VALUES ('Texnika', 1);
INSERT INTO categories (name, product_id) VALUES ('Mevalar', 3);
INSERT INTO categories (name, product_id) VALUES ('Sabzavotlar', 2);

-- ORDERS
CREATE TABLE orders(
    id BIGSERIAL NOT NULL PRIMARY KEY,
    name VARCHAR(50),
    price NUMERIC,
    phone_number VARCHAR(21),
    user_id BIGSERIAL REFERENCES users(id),
    customer_id BIGSERIAL REFERENCES customers(id),
    courier_id BIGSERIAL REFERENCES couriers(id),
    product_id BIGSERIAL REFERENCES products(id),
    quantity NUMERIC
)
;

INSERT INTO orders (name, price, phone_number, user_id, customer_id, courier_id, product_id, quantity) VALUES
('Korzinka', 20000, '93-007-07-07', 1, 2, 1, 3, 4);
INSERT INTO orders (name, price, phone_number, user_id, customer_id, courier_id, product_id, quantity) VALUES
('Makro', 32000, '93-003-03-03', 3, 1, 2, 2, 2);
INSERT INTO orders (name, price, phone_number, user_id, customer_id, courier_id, product_id, quantity) VALUES
('Havas', 22000, '91-001-01-01', 2, 3, 3, 1, 5);

-- Indexes
CREATE INDEX users_name_idx ON users(name);

CREATE INDEX orders_name_idx ON orders(name);
