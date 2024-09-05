-- +goose Up
-- +goose StatementBegin
ALTER TABLE habits ADD COLUMN streak INTEGER DEFAULT 0;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE habits DROP COLUMN streak;
-- +goose StatementEnd
