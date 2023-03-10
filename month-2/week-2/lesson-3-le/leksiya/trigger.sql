

CREATE OR REPLACE FUNCTION customer_check_age() returns trigger LANGUAGE PLPGSQL
    AS
$$
    BEGIN

        -- 10::numeric
        -- cast(ball as numeric)
        -- decimal(10, 2)

        IF new.age > 18 THEN
            return new;
        END IF;

        raise info 'Yosh togri kelmaydi';
        return null;
    END;
$$;

CREATE TRIGGER customer_check_age_tg
BEFORE INSERT ON customer
FOR EACH ROW EXECUTE PROCEDURE customer_check_age();

INSERT INTO customer(name, age) VALUES
('Ali', 20);

INSERT INTO customer(name, age) VALUES
('Abduqodir', 18);
