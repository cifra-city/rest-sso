-- Удаляем зависимости (внешние ключи)
ALTER TABLE devices DROP CONSTRAINT devices_user_id_fkey;
ALTER TABLE refresh_tokens DROP CONSTRAINT refresh_tokens_user_id_fkey;
ALTER TABLE login_history DROP CONSTRAINT login_history_user_id_fkey;

-- Удаляем таблицы
DROP TABLE IF EXISTS login_history;
DROP TABLE IF EXISTS refresh_tokens;
DROP TABLE IF EXISTS devices;
DROP TABLE IF EXISTS users_secret;

DROP TYPE IF EXISTS failure_reason CASCADE;
DROP TYPE IF EXISTS role_type CASCADE;