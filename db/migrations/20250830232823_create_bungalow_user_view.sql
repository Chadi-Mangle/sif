-- +goose Up
-- +goose StatementBegin
CREATE VIEW bungalows_users AS
SELECT
    b.*,
    COALESCE(
        JSON_AGG(u.*) FILTER (WHERE u.id IS NOT NULL), 
        '[]'::json
    ) AS users
FROM
    bungalows b
LEFT JOIN users u ON u.bungalow_id = b.id
GROUP BY
    b.id;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP VIEW IF EXISTS bungalows_users;
-- +goose StatementEnd
