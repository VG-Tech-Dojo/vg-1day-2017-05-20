-- +migrate Up
CREATE TABLE message (
    id INTEGER NOT NULL PRIMARY KEY,
    body TEXT,
    sender_name TEXT,
    created TIMESTAMP NOT NULL DEFAULT (DATETIME('now', 'localtime')),
    updated TIMESTAMP
);

-- +migrate Down
DROP TABLE message;
