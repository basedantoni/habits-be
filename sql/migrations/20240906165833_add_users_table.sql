-- +goose Up
-- +goose StatementBegin
CREATE TABLE users (
    pk INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    id TEXT NOT NULL UNIQUE,
    email TEXT NOT NULL,
    password TEXT,
    created_at TEXT NOT NULL,
    updated_at TEXT NOT NULL
);

ALTER TABLE habits ADD COLUMN user_id INTEGER REFERENCES users(pk) ON DELETE CASCADE;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users;

ALTER TABLE habits DROP COLUMN user_id;
-- +goose StatementEnd
