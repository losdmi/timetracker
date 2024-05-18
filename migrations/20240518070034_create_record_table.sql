-- +goose Up
-- +goose StatementBegin
CREATE TABLE record (
    id BIGINT NOT NULL GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    task TEXT NOT NULL,
    description TEXT NOT NULL,
    time_start TIMESTAMP NOT NULL,
    time_end TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE record;
-- +goose StatementEnd
