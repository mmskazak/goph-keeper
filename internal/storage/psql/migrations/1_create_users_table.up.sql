BEGIN TRANSACTION;

CREATE TABLE IF NOT EXISTS users(
                                         id SERIAL PRIMARY KEY,
                                         login VARCHAR(255) NOT NULL UNIQUE,
                                         password VARCHAR(255) NOT NULL
);

COMMIT;
