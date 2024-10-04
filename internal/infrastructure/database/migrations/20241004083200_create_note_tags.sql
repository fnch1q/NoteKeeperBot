-- +goose Up
-- +goose StatementBegin
CREATE TABLE note_tags (
    note_id INTEGER REFERENCES notes (id) ON DELETE CASCADE,
    tag_id INTEGER REFERENCES tags (id) ON DELETE CASCADE,
    PRIMARY KEY (note_id, tag_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS note_tags;
-- +goose StatementEnd
