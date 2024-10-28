BEGIN TRANSACTION;

CREATE TABLE IF NOT EXISTS texts (
                                     id SERIAL PRIMARY KEY,
                                     user_id BIGINT,
                                     title VARCHAR(255),
                                     description VARCHAR(255),
                                     text_content TEXT
);

-- Добавляем внешний ключ отдельно
ALTER TABLE texts
    ADD CONSTRAINT fk_user_texts
        FOREIGN KEY (user_id)
            REFERENCES users(id)
            ON DELETE CASCADE;

COMMIT;
