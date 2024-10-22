BEGIN TRANSACTION;

CREATE TABLE IF NOT EXISTS texts (
                                     id BIGINT PRIMARY KEY,
                                     user_id BIGINT,
                                     title VARCHAR(255),
                                     description VARCHAR(255),
                                     text TEXT
);

-- Добавляем внешний ключ отдельно
ALTER TABLE texts
    ADD CONSTRAINT fk_user_texts
        FOREIGN KEY (user_id)
            REFERENCES users(id)
            ON DELETE CASCADE;

COMMIT;
