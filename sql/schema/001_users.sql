-- +goose Up
CREATE TABLE users (
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    name TEXT NOT NULL,
    email VARCHAR NOT NULL,
    password VARCHAR NOT NULL
);

-- +goose Down
DROP TABLE users;
