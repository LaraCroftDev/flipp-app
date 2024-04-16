-- +goose Up
CREATE TYPE measurement_units AS ENUM ('kg','lb','g','mL','L','piece','unit','each','bag','pk');

CREATE TABLE IF NOT EXISTS coupons (
    id BIGSERIAL PRIMARY KEY,
    flyer_id BIGINT NOT NULL,
    FOREIGN KEY(flyer_id) REFERENCES flyers(id) ON DELETE CASCADE,
    product_id BIGINT NOT NULL,
    FOREIGN KEY(product_id) REFERENCES products(id),
    measurement_unit MEASUREMENT_UNITS,
    measurement_number DECIMAL(10, 2),
    discounted_price DECIMAL(10, 2) NOT NULL,
    original_price DECIMAL(10, 2),
    coupon_description TEXT
);

-- +goose Down
DROP TABLE IF EXISTS coupons;