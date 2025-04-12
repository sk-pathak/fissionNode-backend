-- +goose Up
-- +goose StatementBegin
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    email VARCHAR(255) NOT NULL UNIQUE,
    name VARCHAR(255),
    password_hash VARCHAR(255), -- Can be NULL if user only uses OAuth
    email_verified BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE user_auth_methods (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    auth_type VARCHAR(20) NOT NULL, -- 'password', 'oauth'
    identifier VARCHAR(255) NOT NULL, -- email for password auth, provider for oauth
    verified BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    UNIQUE(user_id, auth_type, identifier)
);

CREATE TABLE oauth_connections (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    provider VARCHAR(50) NOT NULL, -- 'github', 'google', etc.
    provider_user_id VARCHAR(255) NOT NULL,
    access_token TEXT,
    refresh_token TEXT,
    expires_at TIMESTAMP WITH TIME ZONE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    UNIQUE(provider, provider_user_id)
);

CREATE TABLE projects (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    subdomain VARCHAR(63) UNIQUE,
    status VARCHAR(20) NOT NULL DEFAULT 'created', -- 'created', 'running', 'stopped', 'error'
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE project_configs (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    project_id UUID NOT NULL REFERENCES projects(id) ON DELETE CASCADE,
    auth_type VARCHAR(20) NOT NULL CHECK (auth_type = 'jwt'), -- Restrict to 'jwt'
    database_type VARCHAR(20) NOT NULL CHECK (database_type = 'postgres'), -- Restrict to 'postgres'
    database_url TEXT NOT NULL, -- Database connection string
    port INTEGER CHECK (port > 0 AND port <= 65535), -- Ensure valid port range
    container_id VARCHAR(255),
    environment_variables JSONB DEFAULT '{}', -- Store environment variables as key-value pairs
    features JSONB DEFAULT '[]', -- Array of enabled features (if needed)
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE api_keys (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    project_id UUID NOT NULL REFERENCES projects(id) ON DELETE CASCADE,
    name VARCHAR(255) NOT NULL,
    key_hash VARCHAR(255) NOT NULL UNIQUE, -- Store hashed key, not plain text
    last_used_at TIMESTAMP WITH TIME ZONE,
    expires_at TIMESTAMP WITH TIME ZONE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE usage_metrics (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    project_id UUID NOT NULL REFERENCES projects(id) ON DELETE CASCADE,
    date DATE NOT NULL,
    api_calls INTEGER DEFAULT 0,
    database_size BIGINT DEFAULT 0, -- in bytes
    bandwidth_used BIGINT DEFAULT 0, -- in bytes
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    UNIQUE(project_id, date)
);

CREATE TABLE audit_logs (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID REFERENCES users(id),
    project_id UUID REFERENCES projects(id),
    action VARCHAR(255) NOT NULL,
    details JSONB,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_projects_user_id ON projects(user_id);
CREATE INDEX idx_project_configs_project_id ON project_configs(project_id);
CREATE INDEX idx_api_keys_project_id ON api_keys(project_id);
CREATE INDEX idx_usage_metrics_project_id_date ON usage_metrics(project_id, date);
CREATE INDEX idx_audit_logs_user_id ON audit_logs(user_id);
CREATE INDEX idx_audit_logs_project_id ON audit_logs(project_id);
CREATE INDEX idx_users_email ON users(email);
CREATE INDEX idx_user_auth_methods_user_id ON user_auth_methods(user_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE audit_logs;
DROP TABLE usage_metrics;
DROP TABLE api_keys;
DROP TABLE project_configs;
DROP TABLE projects;
DROP TABLE oauth_connections;
DROP TABLE user_auth_methods;
DROP TABLE users;
DROP extension IF EXISTS "uuid-ossp";
drop index if exists idx_projects_user_id;
DROP INDEX IF EXISTS idx_project_configs_project_id;
DROP INDEX IF EXISTS idx_api_keys_project_id;
DROP INDEX IF EXISTS idx_usage_metrics_project_id_date;
DROP INDEX IF EXISTS idx_audit_logs_user_id;
DROP INDEX IF EXISTS idx_audit_logs_project_id;
DROP INDEX IF EXISTS idx_users_email;
DROP INDEX IF EXISTS idx_user_auth_methods_user_id;
-- +goose StatementEnd
