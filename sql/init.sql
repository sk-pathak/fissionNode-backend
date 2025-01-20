CREATE TABLE IF NOT EXISTS users (
    id       BIGSERIAL PRIMARY KEY,
    name     text NOT NULL,
    email    text NOT NULL UNIQUE,
    username text NOT NULL UNIQUE,
    password text NOT NULL
);
