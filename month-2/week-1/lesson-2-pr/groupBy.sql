SELECT country_of_birth, COUNT(*) FROM person GROUP BY country_of_birth ORDER BY country_of_birth;

ALTER TABLE person DROP CONSTRAINT person_pkey;
ALTER TABLE person ADD PRIMARY KEY (id);

-- add unique
ALTER TABLE person ADD CONSTRAINT unique_email_address UNIQUE (email);


