-- Homework

-- 1. Filmlarga qancha umumiy pul tushgan
-- SUM() foydalansiz



SELECT film.title, ROUND(SUM(payment.amount), 4) AS sum
FROM film
JOIN inventory ON inventory.film_id = film.film_id
JOIN rental ON rental.inventory_id = inventory.inventory_id
JOIN payment ON payment.rental_id = rental.rental_id
GROUP BY film.title
ORDER BY film.title
;



-- 2. Filmlarga nechta customer kirgan



-- 3. Filmlarga kirgan customer lar city chiqazing
--     films  | city_names
-- -----------+-------------
--     avatar | {Chungho, Hanoi, Rio Claro, Tashkent}



-- 4. payment_date pul tolangan kinolar nomini chiqazing va umumiy summa
-- Masalan:
-- payment_date | film_names          | total_amount
-- -------------+---------------------+--------------
--   2007-02-16 | {avatar, jon uik3 } | 220.23


-- 5. Eng maximum customer kirgan film nomini chiqazing.


-- 6. Biron bir country dan qatnashgan customer lar royxatini chiqazing

-- WHERE country = 'Japan'

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