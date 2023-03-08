
CREATE OR REPLACE FUNCTION getActorByID(id INT, name varchar) RETURNS RECORD LANGUAGE PLPGSQL
    AS
$$
    DECLARE
        name VARCHAR;
        actor_data RECORD;

    BEGIN

        SELECT 
            * 
        FROM actor
        INTO actor_data
        WHERE actor_id = id;
        

        return actor_data;
    END;
$$;




DO
$$
    DECLARE
        a INTEGER = 20;
        b INTEGER = 20;
    BEGIN

    IF a > b THEN
        raise info 'A soni katta';
    ELSIF a < b THEN
        raise info 'B soni katta';
    ELSE
        raise info 'Sonlar teng';
    END IF;

    END;
$$;

DO
$$
    DECLARE
        i integer = 0;
    BEGIN

    LOOP
        raise info '%', i;

        i = i + 1;

        EXIT WHEN i > 100;

        -- IF i > 100 THEN
        --     EXIT;
        -- END IF;

    END LOOP;

    END;
$$;


DO
$$
    BEGIN

    FOR i IN 1..10 BY 2 LOOP

        raise info '%', i;

    END LOOP;

    END;
$$;



-- task1
CREATE OR REPLACE FUNCTION reverse_number(number INT) RETURNS INT LANGUAGE PLPGSQL
    AS
$$
    DECLARE
        reverse integer = 0;
    BEGIN
        LOOP
            reverse = reverse * 10 + number % 10;
            number = number / 10;
            EXIT WHEN number <= 0;
        END LOOP;

        return reverse;
    END;
$$;

-- task2
CREATE OR REPLACE FUNCTION total_sum(n INT) RETURNS INT LANGUAGE PLPGSQL
    AS
$$
    DECLARE
        sum integer = 0;
    BEGIN

        FOR i IN 1..n LOOP
            sum = sum + i;

        END LOOP;
        return sum;
    END;
$$;
-----------------------------------------------------------------------------------
DO
$$
    DECLARE
        actor_data record;
    BEGIN

    FOR actor_data IN (
        SELECT
            *
        FROM actor
    ) LOOP

        raise info '%', actor_data;

    END LOOP;
    END;
$$;

-- task3
CREATE OR REPLACE FUNCTION get_film_year(from_date INT, to_date INT) RETURNS VOID LANGUAGE PLPGSQL
    AS
$$
    DECLARE
        counter integer = 0;
        film_data record;
    BEGIN   

    IF from_date > to_date THEN
        return;
    END IF;

    FOR film_data IN (
        SELECT
            *
        FROM film
    ) LOOP
        counter = counter + 1;
        IF film_data.release_year >= from_date AND film_data.release_year <= to_date THEN
            raise info '%. %', counter, film_data;
        END IF; 

    END LOOP;

    return;
    END;
$$;

-- task4

CREATE OR REPLACE FUNCTION get_film_by_language(languageName VARCHAR) RETURNS VOID LANGUAGE PLPGSQL
    AS
$$ 
    DECLARE
        counter integer = 0;
        film_data record;
    BEGIN

    FOR film_data IN (
        SELECT
            *
        FROM film
        JOIN language ON language.language_id = film.language_id
        WHERE LOWER(language.name) = LOWER(languageName)
    ) LOOP
        counter = counter + 1;
        raise info '%. %', counter, film_data;

    END LOOP;

    END;
$$;