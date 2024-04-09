CREATE TABLE IF NOT EXISTS stores(
    id SERIAL PRIMARY KEY,
    retailer_id SERIAL NOT NULL,
    FOREIGN KEY(retailer_id) REFERENCES retailers(id) ON DELETE CASCADE,
    location VARCHAR(10) NOT NULL
);