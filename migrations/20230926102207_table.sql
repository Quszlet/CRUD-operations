-- +goose Up
CREATE TABLE users
(
    id            serial       not null unique,
    name          varchar(255) not null,
    email         varchar(255) not null unique,
    password      varchar(255) not null
);


-- +goose Down
DROP TABLE users;
