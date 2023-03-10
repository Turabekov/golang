
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE investor (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR NOT NULL CHECK (length(name) < 30),
    birth_date DATE CHECK (birth_date > '1950-01-01'),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- in_stock
-- booked
-- in_use

CREATE TABLE car (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    state_number VARCHAR UNIQUE,
    price NUMERIC NOT NULL,
    investor_id UUID NOT NULL REFERENCES investor(id),
    investor_percentage INT DEFAULT 70,
    status VARCHAR NOT NULL DEFAULT 'in_stock',
    mileage NUMERIC NOT NULL,
    oil_change_status BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE customer (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);


-- new
-- client_took
-- client_returned
-- success
-- cancelled

CREATE TABLE orders (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    customer_id UUID NOT NULL REFERENCES customer(id),
    car_id UUID NOT NULL REFERENCES car(id),
    day_count INTEGER NOT NULL,
    total_price NUMERIC NOT NULL,
    paid_price NUMERIC,
    status VARCHAR,
    from_date TIMESTAMP NOT NULL CHECK (from_date::DATE >= CURRENT_DATE),
    to_date TIMESTAMP NOT NULL,
    give_km NUMERIC,
    receive_km NUMERIC,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
