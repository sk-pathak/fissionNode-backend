-- +goose Up
-- +goose StatementBegin

CREATE TABLE IF NOT EXISTS users (
    id BIGSERIAL PRIMARY KEY,
    username TEXT UNIQUE NOT NULL,
    email TEXT UNIQUE NOT NULL,
    password_hash TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS nodes (
    id BIGSERIAL PRIMARY KEY,
    node_name TEXT UNIQUE NOT NULL,
    ip_address TEXT NOT NULL,
    capacity JSONB NOT NULL, -- JSON to store resource details (e.g., CPU, memory).
    status TEXT NOT NULL CHECK (status IN ('online', 'offline')),
    last_heartbeat TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS services (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    node_id BIGINT NOT NULL REFERENCES nodes(id) ON DELETE CASCADE,
    image TEXT NOT NULL, -- Docker image URL
    status TEXT NOT NULL CHECK (status IN ('running', 'stopped', 'failed')),
    public_url TEXT UNIQUE, -- Public-facing URL for the service.
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS node_health_logs (
    id BIGSERIAL PRIMARY KEY,
    node_id BIGINT NOT NULL REFERENCES nodes(id) ON DELETE CASCADE,
    cpu_usage FLOAT NOT NULL, -- Percentage
    memory_usage FLOAT NOT NULL, -- Percentage
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS node_health_logs;
DROP TABLE IF EXISTS services;
DROP TABLE IF EXISTS nodes;
DROP TABLE IF EXISTS users;
-- +goose StatementEnd
