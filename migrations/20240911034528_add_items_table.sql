-- +goose Up
CREATE TABLE
    items (
        id bigserial NOT NULL,
        title text NOT NULL,
        amount real NOT NULL,
        quantity integer NOT NULL,
        status text NOT NULL,
        PRIMARY KEY (id)
    );

-- +goose StatementBegin
SELECT
    'up SQL query';

-- +goose StatementEnd
-- +goose Down
DROP TABLE IF EXISTS items;

-- +goose StatementBegin
SELECT
    'down SQL query';

-- +goose StatementEnd