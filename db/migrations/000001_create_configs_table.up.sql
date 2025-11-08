CREATE TABLE IF NOT EXISTS configs (
    id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    name VARCHAR(100) NOT NULL,
    config_values TEXT NOT NULL,
    version SMALLINT NOT NULL DEFAULT 1,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP
);
CREATE UNIQUE INDEX IF NOT EXISTS idx_config_name_version ON configs (name, version);
CREATE INDEX IF NOT EXISTS idx_config_name ON configs (name);
CREATE INDEX IF NOT EXISTS idx_config_version ON configs (version);