
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

INSERT INTO investor (name, birth_date) VALUES ('Abduqodir', '2005-10-23');

insert into car (state_number, price, investor_id, mileage) VALUES
('10 A 110 AA', 450000, 'e2a724cc-8f3f-426c-b7ac-b074c19fc8bb', 500),
('01 C 777 BA', 600000, '52524086-9ac2-45cc-97af-a2e83f6c23d0', 1200),
('20 B 888 BB', 250000, 'e2a724cc-8f3f-426c-b7ac-b074c19fc8bb', 2200);

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

INSERT INTO customer (name) VALUES
('Abdurahim'),
('Muhiddin');

INSERT INTO orders (
    customer_id,
    car_id,
    day_count,
    total_price,
    status,
    from_date,
    to_date
) VALUES (
    'e479aa14-0ff5-48d9-b863-9f42442d4b1b',
    '69340dda-6cd0-403f-b3c8-9abce3873622',
    2,
    (
        SELECT
            price * 2
        FROM car
        WHERE id = '69340dda-6cd0-403f-b3c8-9abce3873622'
    ),
    'new',
    CURRENT_TIMESTAMP,
    CURRENT_TIMESTAMP + INTERVAL '2 day'
);

UPDATE orders SET status = 'new' WHERE id = 'badae641-11d5-446a-847d-616851366592';
UPDATE orders SET status = 'client_took' WHERE id = 'badae641-11d5-446a-847d-616851366592';
UPDATE orders SET status = 'client_returned', receive_km = 2601, paid_price = 900000 WHERE id = 'badae641-11d5-446a-847d-616851366592';



new -> client_took -> client_returned -> success
new -> cancelled
client_took -> cancelled
