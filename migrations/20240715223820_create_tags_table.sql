-- +goose Up
CREATE TABLE tags (
                      id SERIAL PRIMARY KEY,
                      name VARCHAR(255) NOT NULL
);

-- +goose Down
DROP TABLE tags;
