-- +goose Up
-- +goose StatementBegin
ALTER TABLE habits ADD COLUMN created_at TEXT;
ALTER TABLE habits ADD COLUMN updated_at TEXT;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE habits DROP COLUMN created_at;
ALTER TABLE habits DROP COLUMN updated_at;
-- +goose StatementEnd
