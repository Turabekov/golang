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

INSERT INTO sale(total_amount, user_id) VALUES
(0, 'fbd6dda6-9aed-4009-b568-ff550d42fbcc');

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
    7,
    10000,
    '08f86877-24e2-4707-b46c-eab43dd43a1f'
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

INSERT INTO sale_payments(payment_type, sale_id, cash, uzcard, humo) VALUES
('', 
'061d0654-0db1-4e49-9083-960b746a70eb',
20000,
5000,
15000
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


DELETE FROM sale_payments WHERE id = 'cd994076-e136-45c0-9fe5-a2b90ccaa5b0';