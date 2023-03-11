
INSERT INTO investor (name, birth_date) VALUES ('Abduqodir', '2005-10-23');

insert into car (state_number, price, investor_id, mileage) VALUES
('10 A 110 AA', 450000, '979c545f-e4e6-4da3-92f6-4b578e3f2478', 15000),
('01 C 777 BA', 600000, 'bd3a2490-5ec2-45b9-8a7c-1ea0aa44dd41', 15000),
('20 B 888 BB', 250000, '979c545f-e4e6-4da3-92f6-4b578e3f2478', 15000);

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
    '66c7559f-c7a8-4153-8963-a8460608462e',
    '243452c2-2871-40ff-a41f-98c288762659',
    2,
    (
        SELECT
            price * 2
        FROM car
        WHERE id = '243452c2-2871-40ff-a41f-98c288762659'
    ),
    'new',
    CURRENT_TIMESTAMP,
    CURRENT_TIMESTAMP + INTERVAL '2 day'
);

UPDATE orders SET status = 'client_took' WHERE id = '667a868f-72eb-4613-9f24-bf16a7278f69';
UPDATE orders SET status = 'client_returned', receive_km = 35000 WHERE id = '667a868f-72eb-4613-9f24-bf16a7278f69';


new -> client_took -> client_returned -> success
new -> cancelled
client_took -> cancelled



--------------------------------------------------------------------------------------------------------
car
    mileage = 12000 KM
    day_limit = 200 KM
    limit = 1000 som
    price = 200_000 som


orders
    day_count = 2
    paid_price = 400_000

TRIGGER
    perelimit()

client_cost
    id
    customer_id
    debt
    created_at

SELECT debtor();

