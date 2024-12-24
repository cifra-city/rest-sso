-- Удаление таблиц
DROP TABLE IF EXISTS operations CASCADE;
DROP TABLE IF EXISTS sessions CASCADE;
DROP TABLE IF EXISTS account CASCADE;

-- Удаление ENUM-типов
DROP TYPE IF EXISTS operation_type;
DROP TYPE IF EXISTS failure_reason;
DROP TYPE IF EXISTS role_type;
