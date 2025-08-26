--migrations/001_init.sql
CREATE TABLE IF NOT EXISTS urls(
    id INTEGER PRIMARY KEY,
    url TEXT NOT NULL UNIQUE,
    alias TEXT NOT NULL
);

CREATE INDEX IF NOT EXISTS idx_url_alias ON urls(alias);

