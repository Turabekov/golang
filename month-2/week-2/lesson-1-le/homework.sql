-- Homework

-- 1. Filmlarga qancha umumiy pul tushgan
-- SUM() foydalansiz


CREATE OR REPLACE FUNCTION get_price_all_films() RETURNS TABLE(
    film_name VARCHAR,
    total_amount NUMERIC
) LANGUAGE PLPGSQL
    AS
$$
    DECLARE
        filmData record;
    BEGIN
        

    FOR filmData IN (
        SELECT film.title, ROUND(SUM(payment.amount), 4) AS sum
        FROM film
        JOIN inventory ON inventory.film_id = film.film_id
        JOIN rental ON rental.inventory_id = inventory.inventory_id
        JOIN payment ON payment.rental_id = rental.rental_id
        GROUP BY film.title
        ORDER BY film.title
    ) LOOP
        film_name := filmData.title;
        total_amount := filmData.sum;
        return next;
    END LOOP;
    END;
$$;



-- 2. Filmlarga nechta customer kirgan

CREATE OR REPLACE FUNCTION get_customer_count() RETURNS VOID LANGUAGE PLPGSQL
    AS
$$
    DECLARE
        filmTitle film.title%type;
        count integer = 0;
    BEGIN
        

        FOR filmTitle, count IN (
            SELECT film.title, COUNT(customer.customer_id) AS customers_count
            FROM film
            JOIN inventory ON inventory.film_id = film.film_id
            JOIN rental ON rental.inventory_id = inventory.inventory_id
            JOIN customer ON customer.customer_id = rental.customer_id
            GROUP BY film.title
            ORDER BY film.title
        ) LOOP
            raise info '%. %', filmTitle, count;
        END LOOP;

        return;
    END;
$$;


-- 3. Filmlarga kirgan customer lar city chiqazing
--     films  | city_names
-- -----------+-------------
--     avatar | {Chungho, Hanoi, Rio Claro, Tashkent}

CREATE OR REPLACE FUNCTION get_customer_country() RETURNS VOID LANGUAGE PLPGSQL
    AS
$$
    DECLARE
        filmTitle film.title%type;
        arr_country VARCHAR[] = ARRAY['dasd'];
    BEGIN
        FOR filmTitle, arr_country IN (
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
        ) LOOP
            raise info 'Name: %. Countries: %', filmTitle, arr_country;
        END LOOP;
        return;
    END;
$$;



-- 4. payment_date pul tolangan kinolar nomini chiqazing va umumiy summa
-- Masalan:
-- payment_date | film_names          | total_amount
-- -------------+---------------------+--------------
--   2007-02-16 | {avatar, jon uik3 } | 220.23

CREATE OR REPLACE FUNCTION get_payment_date() RETURNS VOID LANGUAGE PLPGSQL
    AS
$$
    DECLARE
        paymentDate DATE;
        arr_films VARCHAR[] = ARRAY['dasd'];
        amount NUMERIC;
    BEGIN
        FOR paymentDate, arr_films, amount IN (
            SELECT cast(p.payment_date as date) AS payment_date,
            ARRAY_AGG(film.title) AS film_names, 
            SUM(p.amount) AS total_amount
            FROM payment as p
            JOIN rental ON rental.rental_id = p.rental_id
            JOIN inventory ON inventory.inventory_id = rental.inventory_id
            JOIN film ON film.film_id = inventory.film_id
            GROUP BY cast(p.payment_date as date)
            ORDER BY cast(p.payment_date as date)
        ) LOOP
            raise info 'Payment_date: %, Film names: %, total_amount: %', paymentDate, arr_films, amount;
        END LOOP;
        return;
    END;
$$;



-- 5. Eng maximum customer kirgan film nomini chiqazing.

CREATE OR REPLACE FUNCTION get_max_film() RETURNS VOID LANGUAGE PLPGSQL
    AS
$$
    DECLARE
        filmTitle film.title%type;
        amount NUMERIC;
    BEGIN
        FOR filmTitle, amount IN (
            SELECT film.title, COUNT(customer.customer_id) AS films
            FROM film
            JOIN inventory ON inventory.film_id = film.film_id
            JOIN rental ON rental.inventory_id = inventory.inventory_id
            JOIN customer ON customer.customer_id = rental.customer_id
            GROUP BY film.title 
            ORDER BY COUNT(customer.customer_id) DESC
            LIMIT 1
        ) LOOP
            raise info 'Max enterered customers film name: %, quantity of customers: %', filmTitle, amount;
        END LOOP;
        return;
    END;
$$;



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


CREATE OR REPLACE FUNCTION get_country_customers(countryName VARCHAR) RETURNS VOID LANGUAGE PLPGSQL
    AS
$$
    DECLARE
        customer_name VARCHAR;
        idx integer = 1;
    BEGIN
        raise info 'Country: %, Customers:', countryName;
        FOR customer_name IN (
            SELECT 
            customer.first_name || ' ' || customer.last_name AS customer_name 
            FROM country 
            JOIN city ON country.country_id = city.country_id
            JOIN address ON address.city_id = city.city_id
            JOIN customer ON customer.address_id = address.address_id
            WHERE country.country = countryName
            ORDER BY country.country
        ) LOOP
            raise info '%. %', idx, customer_name;
            idx = idx + 1;
        END LOOP;
        return;
    END;
$$;

