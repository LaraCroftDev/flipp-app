-- +goose Up
CREATE TABLE IF NOT EXISTS users (
    id BIGSERIAL PRIMARY KEY,
    email varchar(256) NOT NULL UNIQUE,
    password varchar(255) NOT NULL,
    location varchar(10)
);

-- +goose Down
DROP TABLE IF EXISTS users;