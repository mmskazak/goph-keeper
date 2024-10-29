BEGIN TRANSACTION;

-- Удаление внешнего ключа
ALTER TABLE cards
    DROP CONSTRAINT IF EXISTS fk_user_cards;

-- Удаление таблицы
DROP TABLE IF EXISTS cards;

COMMIT;
