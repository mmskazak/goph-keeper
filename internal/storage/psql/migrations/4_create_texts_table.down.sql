BEGIN TRANSACTION;

-- Удаление внешнего ключа
ALTER TABLE texts
    DROP CONSTRAINT IF EXISTS fk_user_texts;

-- Удаление таблицы
DROP TABLE IF EXISTS texts;

COMMIT;
