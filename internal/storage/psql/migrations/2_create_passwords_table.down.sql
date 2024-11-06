BEGIN TRANSACTION;

-- Удаление внешнего ключа
ALTER TABLE passwords
    DROP CONSTRAINT IF EXISTS fk_user;

-- Удаление таблицы
DROP TABLE IF EXISTS passwords;

COMMIT;
