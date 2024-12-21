CREATE TYPE role_type AS ENUM (
    'admin',
    'user',
    'verify_user'
);

CREATE TABLE "users_secret" (
    "id" UUID NOT NULL PRIMARY KEY,
    "username" VARCHAR(255) NOT NULL UNIQUE,
    "email" VARCHAR(255) NOT NULL UNIQUE,
    "role" role_type DEFAULT 'user' NOT NULL,
    "pass_hash" VARCHAR(255) NOT NULL,
    "token_version" INTEGER DEFAULT 0 NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP NOT NULL DEFAULT now()
);

CREATE INDEX "user_public_username_index" ON "users_secret" ("username");
CREATE INDEX "user_public_email_index" ON "users_secret" ("email");

CREATE TABLE devices (
    id UUID PRIMARY KEY NOT NULL,
    user_id UUID NOT NULL REFERENCES users_secret(id) ON DELETE CASCADE,
    factory_id TEXT NOT NULL,
    device_name TEXT,
    os_version TEXT,
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    last_used TIMESTAMP NOT NULL DEFAULT now(),
    UNIQUE (user_id, factory_id)
);

CREATE INDEX device_user_index ON devices ("user_id");
CREATE INDEX factory_id_index ON devices ("factory_id");
CREATE INDEX device_last_used_index ON devices ("last_used");

CREATE TABLE refresh_tokens (
    id UUID PRIMARY KEY NOT NULL,
    user_id UUID NOT NULL REFERENCES users_secret(id) ON DELETE CASCADE,
    token TEXT NOT NULL,
    expires_at TIMESTAMP NOT NULL,
    device_id UUID NOT NULL REFERENCES devices(id) ON DELETE CASCADE,
    ip_address TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT now()
);

CREATE INDEX refresh_token_device_index ON refresh_tokens ("device_id");
CREATE INDEX refresh_token_value_index ON refresh_tokens ("token");
CREATE INDEX refresh_token_user_index ON refresh_tokens ("user_id");
CREATE INDEX refresh_token_expiry_index ON refresh_tokens ("expires_at");

CREATE TYPE failure_reason AS ENUM (
    'invalid_password',
    'account_locked',
    'expired_token',
    'invalid_device_id',
    'invalid_refresh_token',
    'invalid_device_factory_id',
    'invalid_user_id',
    'too_many_attempts',
    'no_access',
    'success'
);

CREATE TYPE operation_type AS ENUM (
    'login',
    'refresh_token',
    'change_username',
    'change_password',
    'reset_password',
    'change_email',
    'delete_account',
    'delete_session',
    'terminate_session'
);

CREATE TABLE operation_history (
   id UUID PRIMARY KEY NOT NULL,
   user_id UUID NOT NULL REFERENCES users_secret(id) ON DELETE CASCADE,
   device_data TEXT NOT NULL,
   ip_address TEXT NOT NULL,
   operation operation_type NOT NULL,
   success BOOLEAN NOT NULL,
   failure_reason failure_reason NOT NULL,
   operation_time TIMESTAMP NOT NULL DEFAULT now()
);

CREATE INDEX operation_history_user_index ON operation_history ("user_id");
CREATE INDEX operation_history_device_data ON operation_history ("device_data");
CREATE INDEX operation_history_time_index ON operation_history ("operation_time");
CREATE INDEX operation_history_success_index ON operation_history ("success");
