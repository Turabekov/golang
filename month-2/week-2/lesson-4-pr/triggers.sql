-- task1
CREATE OR REPLACE FUNCTION set_payments() RETURNS TRIGGER LANGUAGE PLPGSQL
    AS
$$
    BEGIN
        
        IF (SELECT status FROM sale WHERE id = new.sale_id) = 'in_proccess' THEN

            IF new.cash >= 0 AND new.uzcard >= 0 AND new.humo >= 0 THEN  

                new.total_payment = new.cash + new.uzcard + new.humo;

                IF (SELECT total_amount FROM sale WHERE id = new.sale_id) >= (new.total_payment + (SELECT paid_total_amount FROM sale WHERE id = new.sale_id)) THEN
                    UPDATE sale SET 
                    paid_total_amount = (
                        SELECT paid_total_amount FROM sale WHERE id = new.sale_id
                    ) + new.total_payment,
                    cash = (
                        SELECT cash FROM sale WHERE id = new.sale_id
                    ) + new.cash,
                    uzcard = (
                        SELECT uzcard FROM sale WHERE id = new.sale_id
                    ) + new.uzcard,
                    humo = (
                        SELECT humo FROM sale WHERE id = new.sale_id
                    ) + new.humo,
                    remaining_total_amount = total_amount - (paid_total_amount + new.total_payment)
                    WHERE id = new.sale_id;
                    ------------------------------------------------------------------------------------------------

                    IF (SELECT remaining_total_amount FROM sale WHERE id = new.sale_id) = 0 AND (SELECT total_amount FROM sale WHERE id = new.sale_id) > 0 AND  (SELECT paid_total_amount FROM sale WHERE id = new.sale_id) > 0 THEN 
                        UPDATE sale SET 
                            status = 'finished'
                        WHERE id = new.sale_id;
                    END IF;
                    ------------------------------------------------------------------------------------------------

                ELSE 
                    raise info 'you paid more than nessaccery';

                END IF;

            ELSE 
                raise info '400 error';
            END IF;
            
        ELSE
            raise info 'check sale status';
            return null;
        END IF;
        return new;
    END;
$$;

CREATE TRIGGER set_payments_tg
BEFORE INSERT OR UPDATE  ON sale_payments
FOR EACH ROW EXECUTE PROCEDURE set_payments();
--------------------------------------------------------------------------------------------------------------
-- Task2
CREATE OR REPLACE FUNCTION set_count_products() RETURNS TRIGGER LANGUAGE PLPGSQL
    AS
$$
    BEGIN
        
        IF (SELECT status FROM sale WHERE id = new.sale_id) IN ('in_proccess', 'new') THEN
            new.sum = new.amount * new.price;

            UPDATE sale 
            SET total_amount = (total_amount + new.sum),
            products_count = (products_count + new.amount),
            remaining_total_amount = remaining_total_amount + new.sum
            WHERE id = new.sale_id;

        ELSE
            raise info 'check sale status';
            return null;
        END IF;
        return new;
    END;
$$;

CREATE TRIGGER set_count_products_tg
BEFORE INSERT OR UPDATE  ON sale_products
FOR EACH ROW EXECUTE PROCEDURE set_count_products();
