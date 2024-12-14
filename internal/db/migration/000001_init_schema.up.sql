CREATE TABLE "users_secret" (
    "id" UUID NOT NULL PRIMARY KEY,
    "username" VARCHAR(255) NOT NULL UNIQUE,
    "email" VARCHAR(255) NOT NULL UNIQUE,
    "email_status" BOOLEAN NOT NULL,
    "pass_hash" VARCHAR(255) NOT NULL
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
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    expires_at TIMESTAMP NOT NULL,
    device_id UUID NOT NULL REFERENCES devices(id) ON DELETE CASCADE,
    ip_address TEXT NOT NULL
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
    'success'
);

CREATE TABLE login_history (
    id UUID PRIMARY KEY NOT NULL,
    user_id UUID NOT NULL REFERENCES users_secret(id) ON DELETE CASCADE,
    device_id UUID NOT NULL REFERENCES devices(id) ON DELETE CASCADE,
    ip_address TEXT NOT NULL,
    login_time TIMESTAMP NOT NULL DEFAULT now(),
    success BOOLEAN NOT NULL,
    failure_reason failure_reason
);

CREATE INDEX login_history_user_index ON login_history ("user_id");
CREATE INDEX login_history_device_index ON login_history ("device_id");
CREATE INDEX login_history_time_index ON login_history ("login_time");
CREATE INDEX login_history_success_index ON login_history ("success");