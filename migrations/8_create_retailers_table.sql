CREATE TABLE IF NOT EXISTS retailers(
    id SERIAL PRIMARY KEY,
    category_id INTEGER NOT NULL,
    FOREIGN KEY(category_id) REFERENCES retailer_categories(id),
    retailer_name VARCHAR(25) NOT NULL
);