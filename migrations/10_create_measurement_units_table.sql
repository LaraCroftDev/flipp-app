CREATE TABLE IF NOT EXISTS measurement_units(
    id SMALLSERIAL PRIMARY KEY;
    unit_name VARCHAR(25) NOT NULL UNIQUE
);