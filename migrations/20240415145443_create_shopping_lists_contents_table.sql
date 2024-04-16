-- +goose Up
CREATE TABLE IF NOT EXISTS shopping_lists_contents (
    id BIGSERIAL PRIMARY KEY,
    shopping_list_id BIGINT NOT NULL UNIQUE,
    FOREIGN KEY(shopping_list_id) REFERENCES shopping_lists(id) ON DELETE CASCADE,
    product_id BIGINT NOT NULL,
    FOREIGN KEY(product_id) REFERENCES products(id)
);

-- +goose Down
DROP TABLE IF EXISTS shopping_lists_contents;