CREATE TABLE car (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    state_number VARCHAR UNIQUE,
    price NUMERIC NOT NULL,
    investor_id UUID NOT NULL REFERENCES investor(id),
    investor_percentage INT DEFAULT 70,
    status VARCHAR NOT NULL DEFAULT 'in_stock',
    mileage NUMERIC NOT NULL,
    day_limit INT NOT NULL DEFAULT 200,
    limit NUMERIC NOT NULL DEFAULT 1000,
    oil_change_status BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
insert into car (state_number, price, investor_id, mileage) VALUES
('01 B 888 BB', 450000, 'e2a724cc-8f3f-426c-b7ac-b074c19fc8bb', 500),
('01 R 777 RZ', 600000, '52524086-9ac2-45cc-97af-a2e83f6c23d0', 1200),
('01 D 909 DD ', 250000, 'e2a724cc-8f3f-426c-b7ac-b074c19fc8bb', 2200);

--------------------------------------------------------------------------
CREATE TABLE client_cost(
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    customer_id UUID NOT NULL REFERENCES customer(id),
    debt NUMERIC,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

