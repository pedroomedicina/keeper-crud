-- +goose Up
CREATE TABLE tasks (
                       id SERIAL PRIMARY KEY,
                       title VARCHAR(100) NOT NULL,
                       description TEXT,
                       created_by_user INTEGER REFERENCES users(id),
                       created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                       updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                       deleted_at TIMESTAMP DEFAULT NULL
);

-- +goose Down
DROP TABLE tasks;
