-- +goose Up
CREATE TABLE IF NOT EXISTS favourite_flyers (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL,
    FOREIGN KEY(user_id) REFERENCES users(id) ON DELETE CASCADE,
    flyer_id BIGINT NOT NULL,
    FOREIGN KEY(flyer_id) REFERENCES flyers(id) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE IF EXISTS favourite_flyers;