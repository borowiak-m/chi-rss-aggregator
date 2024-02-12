-- +goose Up
CREATE TABLE users (
    id int NOT NULL,
    title text,
    name text,
    num int,
    PRIMARY KEY(id)
);

-- +goose Down
--DROP TABLE users;