

CREATE TABLE author (
    id SERIAL NOT NULL PRIMARY KEY,
    name VARCHAR(128) NOT NULL
);

CREATE TABLE book (
    id SERIAL NOT NULL PRIMARY KEY,
    name VARCHAR(128) NOT NULL,
    price NUMERIC NOT NULL,
    author_id INT NOT NULL REFERENCES author (id)
);

CREATE TABLE users (
    "id" INT PRIMARY KEY,
    "first_name" VARCHAR(24),
    "last_name" VARCHAR(24),
    "age" INT
);

CREATE TABLE sale (
    id SERIAL NOT NULL PRIMARY KEY,
    user_id INT NOT NULL REFERENCES users (id),
    book_id INT NOT NULL REFERENCES book (id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

TRUNCATE TABLE book RESTART IDENTITY;

INSERT INTO author (name) VALUES 
('Abdullo Qodiriy'),
('Tohir Malik'),
('G''afur G''ulom');

INSERT INTO book (name, price, author_id) VALUES
('O''tkan Kunlar', 23000, 1),
('Bolalik', 25000, 2),
('Shaytanat', 15000, 2),
('Shum bola', 16000, 3);


INSERT INTO users (id, first_name, last_name, age) VALUES
(1, 'Abduqodir', 'Musayev', 17),
(2, 'Davlatbek', 'Jalolov', 20);

SELECT
    b.name AS book_name,
    b.price AS book_price,
    a.name AS author_name
FROM
    book AS b
JOIN author AS a ON b.author_id = a.id
WHERE b.id = 1;

INSERT INTO sale (user_id, book_id) VALUES 
(1, 2),
(2, 3),
(1, 1);


SELECT
    u.first_name || ' ' || u.last_name AS full_name,
    b.name,
    b.price
FROM
    sale AS s
JOIN users AS u ON u.id = s.user_id
JOIN book AS b ON b.id = s.book_id


Task 1.

users table
    --name
    --age
    --phone_number

course table
    --name
    --duration
    --price

teacher table
    --name

       user_name       |  teacher_name | course_name | course_price | course_duration
-----------------------+---------------+-------------+--------------+----------------
 Ahadullo Habibullayev | ___________   |      ______ |  _______     | _______________
 Oybek Abduvosiqov     | ___________   |      ______ |  _______     | _______________
 Abbos Abdurashidov    | ___________   |      ______ |  _______     | _______________
 Mirsaid Akbarov       | ___________   |      ______ |  _______     | _______________
 Samandar Erkinboyev   | ___________   |      ______ |  _______     | _______________
 Samandar Erkinboyev   | ___________   |      ______ |  _______     | _______________


