CREATE
EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE TABLE users
(
    id           uuid                  DEFAULT uuid_generate_v4()
        constraint users_pk primary key,
    type_request varchar(255) NOT NULL,
    browser      varchar(255) NOT NULL,
    os           varchar(255) NOT NULL,
    device       varchar(255) NOT NULL,
    city         varchar(255) NOT NULL,
    country      varchar(255) NOT NULL,
    created_at   timestamp    NOT NULL DEFAULT (now() at time zone 'utc')
);