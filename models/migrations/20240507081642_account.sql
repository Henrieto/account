-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

-- EXTENSIONS --
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS users (
    id UUID NOT NULL DEFAULT uuid_generate_v4(),
    username VARCHAR(225) NOT NULL,
    email VARCHAR(225) NOT NULL,
    first_name VARCHAR(225) NOT NULL,
    last_name VARCHAR(225) NOT NULL,
    gender VARCHAR(100) NOT NULL,
    password_hash TEXT NOT NULL,
    verified BOOLEAN DEFAULT FALSE,
    birthday TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    staff BOOLEAN DEFAULT FALSE,
    superuser BOOLEAN DEFAULT FALSE,
    auth_id TEXT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    group_id UUID REFERENCES groups(id),
    UNIQUE (Email),
    PRIMARY KEY (id)
);


CREATE TABLE IF NOT EXISTS groups(
    id UUID NOT NULL DEFAULT uuid_generate_v4(),
    name  VARCHAR(100) NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    UNIQUE(name),
    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS permissions (
    id UUID NOT NULL DEFAULT uuid_generate_v4(),
    model VARCHAR(100) NOT NULL,
    name VARCHAR(100) NOT NULL,
    codename VARCHAR(100) NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    UNIQUE(model , name , codename),
    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS user_permissions (
    id UUID NOT NULL DEFAULT uuid_generate_v4(),
    user_id UUID  REFERENCES users(id),
    permission_id UUID  REFERENCES permissions(id),
    CONSTRAINT user_perm_pk PRIMARY KEY (user_id, permission_id),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS group_permissions(
    id UUID NOT NULL DEFAULT uuid_generate_v4(),
    group_id UUID  REFERENCES groups(id),
    permission_id UUID  REFERENCES permissions(id),
    CONSTRAINT group_perm_pk PRIMARY KEY (group_id, permission_id),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    PRIMARY KEY (id)
);

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
