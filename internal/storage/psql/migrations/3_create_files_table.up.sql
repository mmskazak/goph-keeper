BEGIN TRANSACTION;

CREATE TABLE IF NOT EXISTS files (
                                     id SERIAL PRIMARY KEY,
                                     user_id BIGINT,
                                     title VARCHAR(255) NOT NULL,
                                     description VARCHAR(255),
                                     file_data BYTEA NOT NULL  -- Сохраняем данные файла в колонке BYTEA
);

-- Добавляем внешний ключ отдельно
ALTER TABLE files
    ADD CONSTRAINT fk_user
        FOREIGN KEY (user_id)
            REFERENCES users(id)
            ON DELETE CASCADE;

COMMIT;
