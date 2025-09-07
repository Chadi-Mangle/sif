-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS refresh_token (
    id VARCHAR(255) PRIMARY KEY,
    first_name text NOT NULL,
    last_name text NOT NULL,
    is_revoked BOOLEAN DEFAULT FALSE,
    token VARCHAR(255) NOT NULL,
    expires_at TIMESTAMP NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS refresh_token;
-- +goose StatementEnd
