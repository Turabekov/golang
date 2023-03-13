
CREATE TABLE users(
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(50) NOT NULL,
    balance INT NOT NULL
);

INSERT INTO users (name, balance) VALUES('Xumoyun', '20000');
INSERT INTO users (name, balance) VALUES('Shaxzod', '5000');

CREATE OR REPLACE procedure payment(sender_idx UUID, receive_idx UUID, price_amount INT) language PLPGSQL
    AS
$$
    DECLARE
        total int;
        komissiya_amount NUMERIC = 0.5;
        cashback int = 0;
        komissiya_sum int;
        userName VARCHAR;
    BEGIN

        IF price_amount >= 1000 THEN
            komissiya_sum = ((price_amount::numeric * komissiya_amount)/100);
            total = price_amount + komissiya_sum;

            IF (SELECT balance FROM users WHERE id = sender_idx) >= total THEN
                INSERT INTO click (sender_id, receive_id, price, komissiya) VALUES
                (sender_idx, receive_idx, price_amount, komissiya_amount);

                COMMIT;
                IF (SELECT COUNT(*) FROM click WHERE sender_id = sender_idx AND status = FALSE GROUP BY sender_id) >= 5 AND (((SELECT COUNT(*)FROM click WHERE sender_id = sender_idx AND status = FALSE GROUP BY sender_id)) % 5 = 0 ) THEN 

                cashback = ((SELECT MAX(price) FROM click WHERE status = false) * 0.5) / 100;

                total = total - cashback;
                
                UPDATE click SET status = true WHERE status = FALSE;
                
                raise info 'your cashback after 5 times is %', ROUND(cashback, 2);
                END IF;  

                UPDATE users SET balance = balance - total WHERE id = sender_idx;
                UPDATE users SET balance = balance + price_amount WHERE id = receive_idx;
                COMMIT;

            ELSE 
                raise info 'not available amount';
            END IF;
        ELSE 
            raise info 'minimum price for transaction 1000';
        END IF;
        
        raise info '    CHECK';

        SELECT name FROM users INTO userName WHERE id = sender_idx;
        
        raise info 'SENDER %', userName;
        
        SELECT name FROM users INTO userName WHERE id = receive_idx;
        raise info 'RECEIVE %', userName;
        
        raise info 'Sender Price %', price_amount;
        raise info 'Komissiya Price %', komissiya_sum;
        raise info 'CASHBACK %', cashback;
    END;
$$;

-- 2. komissiya yechib olishi kerak 0.5%

CREATE TABLE click(
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    sender_id UUID NOT NULL REFERENCES users(id),
    receive_id UUID NOT NULL REFERENCES users(id),
    price INT,
    komissiya NUMERIC DEFAULT 0.5,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP 
)

ALTER TABLE click ADD COLUMN created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP;
ALTER TABLE click ADD COLUMN status BOOLEAN DEFAULT FALSE;

-- 3. cashback har 5 martda otkazmada 0.25 cashback qaytaring



CALL payment('698067a4-023f-49a2-8141-631d8f1e598e', 'e0411345-aae5-4943-8880-03eb9380ec01', 50000);   


-- 4 status; ORDER BY DESC LIMIT 5;