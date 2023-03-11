
CREATE OR REPLACE FUNCTION perelimit() returns trigger LANGUAGE PLPGSQL
    AS
$$
    BEGIN
        raise info 'working';
        IF new.status = 'client_returned' AND old.status = 'client_took' THEN
            raise info 'Hello 1 %', (new.receive_km - old.give_km);
            raise info 'Hello 1 34 %', (SELECT day_limit FROM car WHERE id = old.car_id) * old.day_count;
            IF (new.receive_km - old.give_km) > (SELECT day_limit FROM car WHERE id = old.car_id) * old.day_count THEN
                raise info 'Hello 2';
                INSERT INTO client_cost (customer_id, debt) VALUES
                (old.customer_id, ((new.receive_km - old.give_km) - (SELECT day_limit FROM car WHERE id = old.car_id) * old.day_count) * (SELECT limit_price FROM car WHERE id = old.car_id));
            END IF; 
        END IF;
    return new;
    END;
$$;


CREATE TRIGGER customer_perelimit_tg
BEFORE UPDATE ON orders
FOR EACH ROW EXECUTE PROCEDURE perelimit();
