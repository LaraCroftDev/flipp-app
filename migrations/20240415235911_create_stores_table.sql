-- +goose Up
CREATE TABLE IF NOT EXISTS stores (
    id SERIAL PRIMARY KEY,
    retailer_id INTEGER NOT NULL,
    FOREIGN KEY(retailer_id) REFERENCES retailers(id) ON DELETE CASCADE,
    location VARCHAR(10) NOT NULL
);

-- +goose Down
DROP TABLE IF EXISTS stores;