CREATE TABLE IF NOT EXISTS clipped_coupons(
    id BIGSERIAL PRIMARY KEY,
    shopping_list_id BIGSERIAL NOT NULL UNIQUE,
    FOREIGN KEY(shopping_list_id) REFERENCES shopping_lists(id) on DELETE CASCADE,
    coupons_id BIGSERIAL NOT NULL UNIQUE,
    FOREIGN KEY(coupons_id) REFERENCES coupons(id) on DELETE CASCADE
);