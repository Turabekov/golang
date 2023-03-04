-- Homework

-- 1. Filmlarga qancha umumiy pul tushgan
-- SUM() foydalansiz

SELECT ROUND(SUM(amount), 4) AS total  FROM Payment; 

SELECT film.title, ROUND(SUM(payment.amount), 4) AS sum
FROM film
JOIN inventory ON inventory.film_id = film.film_id
JOIN rental ON rental.inventory_id = inventory.inventory_id
JOIN payment ON payment.rental_id = rental.rental_id
GROUP BY film.title
ORDER BY film.title
;


-- 2. Filmlarga nechta customer kirgan

SELECT COUNT(customer_id) FROM Payment;

SELECT film.title, COUNT(customer.customer_id) AS customers_count
FROM film
JOIN inventory ON inventory.film_id = film.film_id
JOIN rental ON rental.inventory_id = inventory.inventory_id
JOIN customer ON customer.customer_id = rental.customer_id
GROUP BY film.title
ORDER BY film.title
;


-- 3. Filmlarga kirgan customer lar city chiqazing

--     films  | city_names
-- -----------+-------------
--     avatar | {Chungho, Hanoi, Rio Claro, Tashkent}

SELECT film.title, ARRAY_AGG(co.country) AS films
FROM film
JOIN inventory ON inventory.film_id = film.film_id
JOIN rental ON rental.inventory_id = inventory.inventory_id
JOIN customer ON customer.customer_id = rental.customer_id
JOIN address ON address.address_id = customer.address_id
JOIN city as c ON c.city_id = address.city_id
JOIN country as co ON co.country_id = c.country_id
GROUP BY film.title
ORDER BY film.title
;


-- 4. payment_date pul tolangan kinolar nomini chiqazing va umumiy summa
-- Masalan:
-- payment_date | film_names          | total_amount
-- -------------+---------------------+--------------
--   2007-02-16 | {avatar, jon uik3 } | 220.23


SELECT cast(p.payment_date as date) AS payment_date,
ARRAY_AGG(film.title) AS film_names, 
SUM(p.amount) AS total_amount
FROM payment as p
JOIN rental ON rental.rental_id = p.rental_id
JOIN inventory ON inventory.inventory_id = rental.inventory_id
JOIN film ON film.film_id = inventory.film_id
GROUP BY cast(p.payment_date as date)
ORDER BY cast(p.payment_date as date)
;

-- 5. Eng maximum customer kirgan film nomini chiqazing.

SELECT film.title, COUNT(customer.customer_id) AS films
FROM film
JOIN inventory ON inventory.film_id = film.film_id
JOIN rental ON rental.inventory_id = inventory.inventory_id
JOIN customer ON customer.customer_id = rental.customer_id
GROUP BY film.title 
ORDER BY COUNT(customer.customer_id) DESC
LIMIT 1
;



-- 6. Biron bir country dan qatnashgan customer lar royxatini chiqazing

-- WHERE country = 'Japan'

SELECT 
country.country, 
customer.first_name AS customer_name 
FROM country 
JOIN city ON country.country_id = city.country_id
JOIN address ON address.city_id = city.city_id
JOIN customer ON customer.address_id = address.address_id
ORDER BY country.country
;
--     country | customer_name
-- ------------+-------------
--       Japan | Patricia
--       Japan | Elizabeth
--       Japan | Carol
--       Japan | Ruth
--       Japan | Sharon
--       Japan | Michelle
--       Japan | Laura
--       Japan | Sarah
--       Japan | Kimberly
--       Japan | Deborah