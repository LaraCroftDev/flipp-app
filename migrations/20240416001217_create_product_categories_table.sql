-- +goose Up
CREATE TABLE IF NOT EXISTS product_categories (
    id SERIAL PRIMARY KEY,
    category VARCHAR(50) NOT NULL
);
-- +goose Down
DROP TABLE IF EXISTS product_categories;