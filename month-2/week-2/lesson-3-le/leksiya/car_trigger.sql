
CREATE OR REPLACE FUNCTION order_status() RETURNS TRIGGER LANGUAGE PLPGSQL
    AS
$$
    BEGIN

        IF new.status NOT IN ('new', 'client_took', 'client_returned', 'success', 'cancelled') THEN

            raise exception 'not allowed status';

            return null;
        END IF;

        IF new.status = 'client_took' AND old.status = 'new' THEN

            UPDATE car SET status = 'in_use' WHERE id = new.car_id;

            raise info 'OK';

            UPDATE orders SET give_km = (SELECT mileage FROM car WHERE id = new.car_id) WHERE id = new.id;

        ELSIF new.status = 'client_returned' AND old.status = 'client_took' THEN

            UPDATE car SET status = 'in_stock' WHERE id = new.car_id;

        ELSIF new.status = 'new' AND old.status IS NULL THEN

            UPDATE car SET status = 'booked' WHERE id = new.car_id;

        END IF;

        return new;
    END;
$$;

CREATE TRIGGER order_status_tg
AFTER INSERT OR UPDATE ON orders
FOR EACH ROW EXECUTE PROCEDURE order_status();


CREATE OR REPLACE FUNCTION oil_change_status() RETURNS TRIGGER LANGUAGE PLPGSQL
    AS
$$
    BEGIN

        IF (SElECT oil_change_status FROM car WHERE id = new.car_id) THEN 
            raise info 'car oil out of';
            return null;
        END IF;

        IF new.status = 'client_returned' AND old.status = 'client_took' THEN
            IF (SELECT
                SUM(receive_km - give_km)
            FROM orders
            WHERE car_id = new.car_id) > 10000 THEN
                UPDATE car SET 
                    oil_change_status = true, 
                    mileage = (
                        SELECT
                            SUM(receive_km - give_km)
                        FROM orders
                        WHERE car_id = new.car_id
                    )
                  WHERE id = new.car_id;
            ELSE 
                UPDATE car SET 
                mileage = (
                    SELECT
                        SUM(receive_km - give_km)
                    FROM orders
                    WHERE car_id = new.car_id
                  )
                WHERE id = new.car_id;
            
            END IF;
        END IF;

        return new;
    END;
$$;

CREATE TRIGGER oil_change_status_tg
BEFORE INSERT ON orders
FOR EACH ROW EXECUTE PROCEDURE oil_change_status();

drop trigger oil_change_status_tg ON orders;


