-- +goose Up
-- +goose StatementBegin
CREATE TABLE bungalows (
    id SERIAL PRIMARY KEY,
    capacity INTEGER NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE bungalows;
-- +goose StatementEnd
