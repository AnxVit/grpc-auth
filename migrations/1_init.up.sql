CREATE TABLE IF NOT EXISTS users
(
    id          INTEGER PRIMERY KEY,
    email       TEXT NOT NULL UNIQUE,
    pass_hash   BLOB NOT NULL
);
CREATE INDEX IF NOT EXISTS idx_email ON users (email);

CREATE TABLE IF NOT EXISTS apps
(
    id      INTEGER PRIMERY KEY,
    name    TEXT NOT NULL UNIQUE,
    secret  TEXT NOT NULL UNIQUE
);