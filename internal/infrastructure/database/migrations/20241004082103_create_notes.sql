-- +goose Up
-- +goose StatementBegin
CREATE TABLE notes (
    id SERIAL PRIMARY KEY,
    category_id INTEGER REFERENCES categories (id) ON DELETE CASCADE,
    title VARCHAR(255) NOT NULL,
    content TEXT,
    created_date TIMESTAMP NOT NULL DEFAULT NOW()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS notes;
-- +goose StatementEnd
