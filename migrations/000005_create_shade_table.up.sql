CREATE TABLE shades
(
    id         SERIAL PRIMARY KEY,
    name       VARCHAR(255) NOT NULL,
    image      VARCHAR(255) NOT NULL,
    sort_index INTEGER      NOT NULL DEFAULT 100,
    created_at TIMESTAMP    NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP    NOT NULL DEFAULT NOW()
);