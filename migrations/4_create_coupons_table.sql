CREATE TABLE IF NOT EXISTS coupons(
    id BIGSERIAL PRIMARY KEY,
    flyer_id BIGSERIAL NOT NULL UNIQUE,
    FOREIGN KEY(flyer_id) REFERENCES flyers(id) ON DELETE CASCADE,
    product_id BIGSERIAL NOT NULL,
    FOREIGN KEY(product_id) REFERENCES products(id),
    measure_unit_id INTEGER,
    FOREIGN KEY(measure_unit_id) REFERENCES measurement_units(id),
    measurement_number INTEGER,
    discounted_price INTEGER NOT NULL,
    coupon_description TEXT
);