-- +goose Up
-- +goose StatementBegin

CREATE TABLE users (
    id BIGINT PRIMARY KEY,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NULL,
    deleted_at TIMESTAMPTZ DEFAULT NULL,

    username VARCHAR UNIQUE NOT NULL,
    password VARCHAR NOT NULL
);

CREATE TABLE roles (
    id BIGINT PRIMARY KEY,
    name VARCHAR UNIQUE NOT NULL
);

CREATE TABLE permissions (
    id BIGINT PRIMARY KEY,
    name VARCHAR UNIQUE NOT NULL
);

-- m2m
CREATE TABLE user_roles (
    user_id BIGINT NOT NULL,
    role_id BIGINT NOT NULL,

    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE RESTRICT,
    FOREIGN KEY (role_id) REFERENCES roles(id) ON DELETE RESTRICT,

    PRIMARY KEY (user_id, role_id)
);

-- m2m
CREATE TABLE role_permissions (
    role_id BIGINT NOT NULL,
    permission_id BIGINT NOT NULL,

    FOREIGN KEY (role_id) REFERENCES roles(id) ON DELETE RESTRICT,
    FOREIGN KEY (permission_id) REFERENCES permissions(id) ON DELETE RESTRICT,

    PRIMARY KEY (role_id, permission_id)
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TABLE role_permissions;
DROP TABLE user_roles;
DROP TABLE users;
DROP TABLE roles;
DROP TABLE permissions;

-- +goose StatementEnd