BEGIN TRANSACTION;

CREATE TABLE IF NOT EXISTS passwords(
                                         id SERIAL PRIMARY KEY,
                                         user_id BIGINT,
                                         resource VARCHAR(255) NOT NULL,
                                         login VARCHAR(255) NOT NULL,
                                         password VARCHAR(255) NOT NULL
);

-- Добавляем внешний ключ отдельно
ALTER TABLE passwords
    ADD CONSTRAINT fk_user
        FOREIGN KEY (user_id)
            REFERENCES users(id)
            ON DELETE CASCADE;


COMMIT;
