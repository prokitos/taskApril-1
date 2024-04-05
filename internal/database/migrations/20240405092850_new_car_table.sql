-- +goose Up
-- +goose StatementBegin

ALTER TABLE car ADD COLUMN owner INTEGER REFERENCES people(id) ON DELETE CASCADE;

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP COLUMN owner;
-- +goose StatementEnd