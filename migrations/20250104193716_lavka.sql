-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
ALTER TABLE courier ADD COLUMN courier_type VARCHAR(8);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
ALTER TABLE courier DROP COLUMN courier_type;
-- +goose StatementEnd
