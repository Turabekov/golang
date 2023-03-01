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


-- =================================================================================================================
group table
    -- course_id
    -- user_id
    -- teacher_id
    -- created_at

-- Users
CREATE TABLE users(
    id SERIAL NOT NULL PRIMARY KEY,
    name VARCHAR(24) NOT NULL,
    phone_number VARCHAR(15) NOT NULL
);

INSERT INTO users(name, phone_number) VALUES
('Khumoyun', '93-379-11-10'),
('Shaxzod', '93-115-07-70'),
('Baxrom', '99-772-77-77'),
('Mirazam', '93-811-22-22'),
('Sirojiddin', '97-400-75-80');

-- Teacher
CREATE TABLE teacher(
    id SERIAL NOT NULL PRIMARY KEY,
    name VARCHAR(24) NOT NULL
);

INSERT INTO teacher(name) VALUES
('Asadbek'),
('Abdulaxad'),
('Sherzod'),
('Bexruz');


-- Course
CREATE TABLE course(
    id SERIAL NOT NULL PRIMARY KEY,
    name VARCHAR(24) NOT NULL,
    duration INT,
    price NUMERIC NOT NULL

);

INSERT INTO course(name, duration, price) VALUES 
('Go Backend', 3, 2000000),
('VueJS Frontend', 3, 1500000),
('DevOps', 5, 2000000),
('IELTS - master', 4, 1000000);



-- Group
CREATE TABLE group_class(
    course_id INT NOT NULL REFERENCES course (id),
    user_id INT NOT NULL REFERENCES users (id),
    teacher_id INT NOT NULL REFERENCES teacher (id)
)

INSERT INTO group_class(course_id, user_id, teacher_id) VALUES
(1, 1, 1),
(1, 2, 1),
(4, 4, 3),
(4, 2, 3),
(3, 3, 2);


-- Result query
SELECT
    u.name AS user_name,
    t.name AS teacher_name,
    c.name AS course_name,
    c.price AS course_price,
    c.duration AS course_duration
FROM
    group_class AS g
JOIN users AS u ON u.id = g.user_id
JOIN teacher AS t ON t.id = g.teacher_id
JOIN course AS c ON c.id = g.course_id


