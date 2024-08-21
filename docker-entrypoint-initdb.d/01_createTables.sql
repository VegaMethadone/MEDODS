CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    guid VARCHAR(256),
    refresh VARCHAR(256),
    email VARCHAR(256)
);
