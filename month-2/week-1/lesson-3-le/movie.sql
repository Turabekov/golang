
CREATE TABLE movies (
    id SERIAL NOT NULL PRIMARY KEY,
    name VARCHAR NOT NULL,
    published_at DATE
);

CREATE TABLE comments (
    id SERIAL NOT NULL PRIMARY KEY,
    description text,
    user_id INT NOT NULL REFERENCES users(id),
    movie_id INT NOT NULL REFERENCES movies(id)
);

insert into users (id, first_name, last_name, age) VALUES
(3, 'Ali', 'Ismoilov', 19);

insert into movies (name, published_at) VALUES
('Avatar 2', '2022-12-10'),
('Qog''oz Bino', '2018-01-19'),
('Forsaj', '2000-01-23');

insert into comments (description, user_id, movie_id) VALUES
('Basharasidan malum kino qanaqaligi', 1, 2),
('Zor kino', 2, 1),
('Davomi qachon chiqadi', 1, 1),
('Yaxshi kino ekan', 3, 3)
;


SELECT
    u.first_name || ' ' || u.last_name AS full_name,
    ARRAY_AGG(m.name) AS movie_name,
    ARRAY_AGG(c.description)
FROM
    users AS u
JOIN comments AS c ON c.user_id = u.id
JOIN movies AS m ON m.id = c.movie_id
GROUP BY full_name
ORDER BY full_name

-- TASK1
SELECT 
    m.name as kino_name,
    ARRAY_AGG(c.description) as descriptions,
    ARRAY_AGG(u.name) as users
FROM comments AS c
JOIN users AS u ON c.user_id = u.id
JOIN movies AS m ON m.id = c.movie_id
GROUP BY m.name
ORDER BY m.name;

-- TASK2
SELECT 
    m.name as kino_name,
    COUNT(c.description) as descriptions
FROM comments AS c
JOIN movies AS m ON m.id = c.movie_id
GROUP BY m.name
ORDER BY m.name;

-- TASK3
SELECT 
    m.name as kino_name,
    COUNT(c.description) as descriptions
FROM comments AS c
JOIN movies AS m ON m.id = c.movie_id
GROUP BY m.name
ORDER BY COUNT(c.description)
LIMIT 1;

KINO NAME | DESCRIPTION | USER
Avatar