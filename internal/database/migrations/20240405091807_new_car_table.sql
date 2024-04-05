-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS people (
 id SERIAL PRIMARY KEY,
 name VARCHAR(255) NOT NULL,
 surname VARCHAR(255) NOT NULL,
 patronymic VARCHAR(255) default "none"
);
CREATE TABLE IF NOT EXISTS car (
 id SERIAL PRIMARY KEY,
 regNum VARCHAR(255) NOT NULL,
 mark VARCHAR(255) NOT NULL,
 model VARCHAR(255) NOT NULL,
 year VARCHAR(255) default "none"
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE people;
DROP TABLE car;
-- +goose StatementEnd