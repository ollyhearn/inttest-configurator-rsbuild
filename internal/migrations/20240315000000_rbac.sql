-- +goose Up
-- +goose StatementBegin

CREATE TABLE users (
    id BIGSERIAL PRIMARY KEY,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NULL,
    deleted_at TIMESTAMPTZ DEFAULT NULL,

    username VARCHAR UNIQUE NOT NULL,
    password VARCHAR NOT NULL
);

CREATE TABLE roles (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR UNIQUE NOT NULL,
    description VARCHAR
);

CREATE TABLE permissions (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR UNIQUE NOT NULL,
    description VARCHAR
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

-- insert some basic mock data
CREATE EXTENSION IF NOT EXISTS pgcrypto;

INSERT INTO users (username, password)
VALUES ('root', crypt('root', gen_salt('bf')));

INSERT INTO permissions (name)
VALUES
('perm_list_user'),
('perm_create_user'),
('perm_edit_user'),
('perm_delete_user'),
('perm_create_project'),
('perm_edit_project'),
('perm_delete_project');

INSERT INTO roles (name)
VALUES
('root');

INSERT INTO role_permissions (role_id, permission_id)
VALUES (1, 1),
(1, 2),
(1, 3),
(1, 4),
(1, 5),
(1, 6),
(1, 7);

INSERT INTO user_roles (user_id, role_id)
VALUES (1, 1);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TABLE role_permissions;
DROP TABLE user_roles;
DROP TABLE users;
DROP TABLE roles;
DROP TABLE permissions;

-- +goose StatementEnd