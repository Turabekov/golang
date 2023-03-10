-- TRIGGER

CREATE OR FUNCTION customer_check_age() returns trigger LANGUAGE PLPGSQL
    AS
$$
    BEGIN
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

INSERT INTO customer(name, age) VALUES('Ali', 20);


-- TASK 1 

CREATE TABLE student(
    id BIGSERIAL NOT NULL PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    ball INT NOT NULL,
    task_counter NUMERIC
);

CREATE OR REPLACE FUNCTION student_calc_tasks() returns trigger LANGUAGE PLPGSQL
    AS
$$ 
    BEGIN
        IF new.ball <= 0 OR new.ball >= 100 THEN 
            raise info 'Ball out of range';
            return null;
        END IF;

        new.task_counter = ROUND(new.ball::numeric / 10, 2);

        return new;
    END;
$$;

CREATE TRIGGER student_calc_tasks_tg
BEFORE INSERT ON student
FOR EACH ROW EXECUTE PROCEDURE student_calc_tasks();

INSERT INTO student(name, ball) VALUES('Mirazam', 35);
INSERT INTO student(name, ball) VALUES('Xumoyun', 35);

