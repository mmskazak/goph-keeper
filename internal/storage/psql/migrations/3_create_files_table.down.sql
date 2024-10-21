BEGIN TRANSACTION;

-- Удаление внешнего ключа
ALTER TABLE files
    DROP CONSTRAINT IF EXISTS fk_user;

-- Удаление таблицы
DROP TABLE IF EXISTS files;

COMMIT;
