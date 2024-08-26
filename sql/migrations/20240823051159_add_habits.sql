-- +goose Up
-- +goose StatementBegin
CREATE TABLE habits (
    id TEXT PRIMARY KEY,
    title TEXT NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE habits;
-- +goose StatementEnd
