CREATE TABLE IF NOT EXISTS flyers(
    id BIGSERIAL PRIMARY KEY,
    retailer_id INTEGER NOT NULL,
    FOREIGN KEY (retailer_id) REFERENCES retailers(id) ON DELETE CASCADE,
    start_date DATE NOT NULL,
    end_date DATE NOT NULL
);