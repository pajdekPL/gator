-- +goose UP
CREATE TABLE feeds (
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    name TEXT NOT NULL,
    URL TEXT UNIQUE NOT NULL,
    user_id UUID NOT NULl REFERENCES users(id) ON DELETE CASCADE
);


-- +goose Down
DROP TABLE feeds;