BEGIN TRANSACTION;

CREATE TABLE IF NOT EXISTS passwords(
                                         id SERIAL PRIMARY KEY,
                                         user_id BIGINT,
                                         title VARCHAR(255) NOT NULL,
                                         description VARCHAR(255),
                                         credentials JSONB
);

-- Добавляем внешний ключ отдельно
ALTER TABLE passwords
    ADD CONSTRAINT fk_user
        FOREIGN KEY (user_id)
            REFERENCES users(id)
            ON DELETE CASCADE;


COMMIT;

