-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
CREATE TABLE courier (
    courier_id BIGINT PRIMARY KEY,
    regions TEXT NOT NULL,
    working_hours TEXT NOT NULL
);
CREATE TABLE "order" (
    order_id BIGINT PRIMARY KEY,
    weight FLOAT NOT NULL,
    regions INT NOT NULL,
    delivery_hours TEXT NOT NULL,
    cost INT NOT NULL,
    completed_time TIMESTAMP NOT NULL
); 
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
DROP TABLE "order";
DROP TABLE courier;
-- +goose StatementEnd
