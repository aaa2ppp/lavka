-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
ALTER TABLE courier ADD COLUMN created_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP;
ALTER TABLE "order" ADD COLUMN created_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP;
ALTER TABLE "order" ADD COLUMN assigned_date DATE;
ALTER TABLE "order" ALTER COLUMN completed_time DROP NOT NULL;
UPDATE "order" SET completed_time=NULL WHERE completed_time='0001-01-01T00:00:00Z';

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
ALTER TABLE courier DROP COLUMN created_time;
ALTER TABLE "order" DROP COLUMN created_time;
ALTER TABLE "order" DROP COLUMN assigned_date;
UPDATE "order" SET completed_time='0001-01-01T00:00:00Z' WHERE completed_time IS NULL;
ALTER TABLE "order" ALTER COLUMN completed_time SET NOT NULL;
-- +goose StatementEnd
