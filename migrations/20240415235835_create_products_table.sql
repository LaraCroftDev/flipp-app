-- +goose Up
CREATE TABLE IF NOT EXISTS products (
    id BIGSERIAL PRIMARY KEY,
    category_id INTEGER NOT NULL,
    FOREIGN KEY(category_id) REFERENCES product_categories(id),
    product_name VARCHAR(50) NOT NULL,
    product_description TEXT
);

-- +goose Down
DROP TABLE IF EXISTS products;