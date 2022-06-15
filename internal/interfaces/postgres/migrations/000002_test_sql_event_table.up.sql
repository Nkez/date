CREATE TYPE type_request_t AS ENUM ('GET_DATE', 'GET_TIME');

ALTER TABLE users ALTER COLUMN type_request type type_request_t USING type_request::varchar::type_request_t;



