

ALTER SEQUENCE person_id_seq RESTART WITH 9;

SELECT * FROM pg_available_extensions;

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

SELECT uuid_generate_v4();