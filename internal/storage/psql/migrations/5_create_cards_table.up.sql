BEGIN TRANSACTION;

CREATE TABLE IF NOT EXISTS cards (
                                    id SERIAL PRIMARY KEY,
                                    user_id BIGINT,
                                    title VARCHAR(255),
                                    description VARCHAR(255),
                                    number VARCHAR(255),    -- Поле для хранения номера
                                    pincode VARCHAR(20),    -- Поле для хранения PIN-кода
                                    cvv VARCHAR(10),        -- Поле для хранения CVV
                                    expire VARCHAR(50)      -- Поле для хранения срока действия
                                    );

-- Добавляем внешний ключ отдельно
ALTER TABLE cards
    ADD CONSTRAINT fk_user_cards
        FOREIGN KEY (user_id)
            REFERENCES users(id)
            ON DELETE CASCADE;

COMMIT;
