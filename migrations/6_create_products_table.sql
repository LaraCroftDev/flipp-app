CREATE TABLE IF NOT EXISTS products(
    id BIGSERIAL PRIMARY KEY,
    product_name VARCHAR(25) NOT NULL,
    product_description TEXT,
);