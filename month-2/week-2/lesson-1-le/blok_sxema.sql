
SELECT $$ O'ktam o'qituvchi $$;

-- PL/PGSQL
-- Procedure Language Postgresql

DO 
$$
    DECLARE 
        name VARCHAR;
        actor_data RECORD;
        count INT;
    BEGIN
        SELECT 
            *
        FROM actor 
        INTO actor_data
    WHERE actor_id  = 1;
        SELECT 
            COUNT(*)
        FROM film INTO count;

    raise info 'Hello %', actor_data.actor_id;
    raise info 'Film table count %', count;
    END;
$$;


-- task 1
DO 
$$
    DECLARE 
        count INT;
    BEGIN
        SELECT 
            COUNT(*)
        FROM film 
        INTO count;

    raise info 'Film table count %', count;
    END;
$$;