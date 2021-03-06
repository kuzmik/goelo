CREATE TABLE IF NOT EXISTS logs (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    hash VARCHAR(64) NOT NULL,
    server VARCHAR(64) NOT NULL,
    channel VARCHAR(64) NOT NULL,
    nick VARCHAR(64) NOT NULL,
    message TEXT NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP
);

CREATE UNIQUE INDEX u_idx_hash ON logs(hash);

