CREATE TABLE sale(
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    date TIMESTAMP,
    total_amount NUMERIC DEFAULT 0,
    paid_total_amount NUMERIC DEFAULT 0,
    remaining_total_amount NUMERIC DEFAULT 0,
    status VARCHAR NOT NULL DEFAULT 'new',
    products_count INT DEFAULT 0,
    cash INT DEFAULT 0,
    uzcard INT DEFAULT 0,
    humo INT DEFAULT 0,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO sale(total_amount) VALUES
(0);

UPDATE sale SET status = 'in_proccess' WHERE id = '';

CREATE TABLE sale_products( 
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    amount INT NOT NULL,
    price INT NOT NULL,
    SUM INT NOT NULL,
    sale_id UUID NOT NULL REFERENCES sale(id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO sale_products(amount, price, sale_id) VALUES 
(
    2,
    10000,
    '8735dec9-12d8-4174-af80-3daa50bff50e'
);

SELECT * FROM sale_products;

CREATE TABLE sale_payments(
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    payment_type VARCHAR NOT NULL,
    sale_id UUID NOT NULL REFERENCES sale(id),
    cash INT DEFAULT 0,
    uzcard INT DEFAULT 0,
    humo INT DEFAULT 0,
    total_payment NUMERIC NOT NULL DEFAULT 0,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO sale_payments(payment_type, sale_id, cash) VALUES
('', 
'd1451b48-82e5-4d16-82cc-71d17daadc06',
23000
);


-- status finished
-----------------------------------------------------------------------------------------------------------
CREATE TABLE client(
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR NOT NULL
);

INSERT INTO client(name) VALUES('Xumoyun');
INSERT INTO client(name) VALUES('Shaxzod');



ALTER TABLE sale
ADD COLUMN user_id UUID REFERENCES client(id),
ADD COLUMN old_client BOOLEAN DEFAULT FALSE,
ADD COLUMN old_client_amount INT DEFAULT 0;

ALTER TABLE sale
ALTER COLUMN user_id SET NOT NULL;


