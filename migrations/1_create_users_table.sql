CREATE TABLE IF NOT EXISTS users(
    id BIGSERIAL PRIMARY KEY,
    email varchar(256) NOT NULL UNIQUE,
    password varchar(50) NOT NULL,
    location varchar(10)
);