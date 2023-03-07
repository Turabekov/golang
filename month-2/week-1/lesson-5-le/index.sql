SELECT
    *
FROM
    address
WHERE phone = '288241215394'


EXPLAIN ANALYZE SELECT
    *
FROM
    address
WHERE phone = '288241215394'

-- Execution Time: 0.258 ms

CREATE INDEX addres_phone_idx ON address(phone);

-- Execution Time: 0.113 ms


CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name VARCHAR NOT NULL,
    email VARCHAR
);

DELETE FROM users;

TRUNCATE TABLE users RESTART IDENTITY;

INSERT INTO users (name, email)
SELECT 'USER' || t.number, t.number || '@gmail.com' FROM 
    GENERATE_SERIES(1, 1000000, 1) AS t(number)

-- Execution Time: 58.757 ms

CREATE INDEX users_name_idx ON users USING HASH(name);

-- Execution Time: 0.084 ms

DROP INDEX users_name_idx;

CREATE UNIQUE INDEX users_email_idx ON users(email);

INSERT INTO users (name, email)
VALUES ('Asadbek', '1@gmail.com');


CREATE UNIQUE INDEX idx_ic_last_name
ON customer(LOWER(last_name));

-- 1 abduqodir
-- 2 ABDUQODIR error: duplicate
