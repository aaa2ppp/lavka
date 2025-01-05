-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
ALTER TABLE "order" ADD COLUMN courier_id BIGINT;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
ALTER TABLE "order" DROP COLUMN courier_id;
-- +goose StatementEnd
