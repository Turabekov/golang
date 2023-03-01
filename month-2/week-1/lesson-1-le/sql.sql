
SQL - Structured Query Language

No SQL

SQL
    - PostgreSql
    - MySQL
    - Oracle
    - DB2

NOSQL
    - Mongodb
    - Cassandra
    - Graph
    - Redis


sudo -i -u postgres

psql

\! clear - execute terminal command
\du      - user list show
\l       - List of databases
\q       - quit postgresql
\c       - connect to database
\dt      - List of tables
\d       - table info

default port 5432

psql -U postgres -d  postgres -h localhost

CREATE ROLE abduqodir WITH LOGIN PASSWORD '12345';
DROP ROLE abduqodir;

CREATE ROLE abduqodir WITH SUPERUSER LOGIN PASSWORD '12345';

CREATE DATABASE mydb;

-- SQL Statements
DDL - Data Definition Language: CREATE, DROP
DQL - Data Query Language: SELECT
DML - Data Manipulation Language: INSERT, UPDATE, DELETE
DCL - Data Control Language: GRANT, REVOKE
TCL - Transaction Control Language: COMMIT, ROLLBACK


SELECT
INSERT
UPDATE
DELETE

SELECT current_user;

CREATE TABLE users (
    id INT PRIMARY KEY,
    first_name VARCHAR(24),
    last_name VARCHAR(24),
    age INT
);

DROP TABLE user;

INSERT INTO user(id, first_name, last_name, age) VALUES
(1, 'Shohruh', 'Safarov', 19);

INSERT INTO "user"(id, first_name, last_name, age) VALUES
(2, 'Abduqodir', 'Musayev', 17),
(3, 'Davlatbek', 'Jalolov', 20);

INSERT INTO user(id, first_name, last_name, age) VALUES
(3, 'Ali', 'Ismoilov', 20);

SELECT id, first_name, last_name, age FROM users;

SELECT * FROM users;

UPDATE users SET age = 18;

UPDATE users SET 
    first_name = 'Abdurahim', 
    last_name = 'Isanov', 
    age = 20 
WHERE id = 1;

DELETE FROM users;

DELETE FROM users WHERE id = 3;


