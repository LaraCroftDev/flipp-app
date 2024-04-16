-- +goose Up
CREATE TABLE IF NOT EXISTS retailers (
    id SERIAL PRIMARY KEY,
    retailer_name VARCHAR(25) NOT NULL
);

-- +goose Down
DROP TABLE IF EXISTS retailers;