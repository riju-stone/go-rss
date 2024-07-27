-- +goose Up
CREATE TABLE users(
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    fname TEXT NOT NULL,
    lname TEXT NOT NULL,
    location TEXT NOT NULL
);
-- +goose Down
DROP TABLE users;
