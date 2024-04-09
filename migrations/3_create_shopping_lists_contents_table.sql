CREATE TABLE IF NOT EXISTS shopping_lists_contents(
    id BIGSERIAL PRIMARY KEY,
    shopping_list_id BIGSERIAL NOT NULL UNIQUE,
    FOREIGN KEY(shopping_list_id) REFERENCES shopping_lists(id) ON DELETE CASCADE,
    product_id BIGSERIAL NOT NULL,
    FOREIGN KEY(product_id) REFERENCES products(id),
    date_content_added DATE NOT NULL DEFAULT CURRENT_DATE
);