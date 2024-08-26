-- +goose Up
-- +goose StatementBegin
PRAGMA foreign_keys = ON;

CREATE TABLE contributions (
    pk INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    id TEXT NOT NULL UNIQUE,
    time_spent INTEGER NOT NULL,
    habit_id INTEGER REFERENCES habits(pk) ON DELETE CASCADE,
    created_at TEXT,
    updated_at TEXT
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
PRAGMA foreign_keys = OFF;

DROP TABLE contributions;
-- +goose StatementEnd
