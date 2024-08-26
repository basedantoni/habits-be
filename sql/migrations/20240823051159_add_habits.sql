-- +goose Up
-- +goose StatementBegin
CREATE TABLE habits (
    pk INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    id TEXT NOT NULL UNIQUE,
    title TEXT NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE habits;
-- +goose StatementEnd
