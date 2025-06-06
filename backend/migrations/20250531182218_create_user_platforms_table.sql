-- +goose Up
-- +goose StatementBegin
CREATE TABLE user_platforms (
    id VARCHAR(21) PRIMARY KEY,
    user_id VARCHAR(21) NOT NULL REFERENCES users (id) ON DELETE CASCADE,
    platform_id VARCHAR(21) NOT NULL REFERENCES platforms (id) ON DELETE CASCADE,
    access_token TEXT NOT NULL,
    refresh_token TEXT,
    expires_at TIMESTAMP
    WITH
        TIME ZONE,
        account_id VARCHAR(255) NOT NULL,
        account_name VARCHAR(255),
        is_active BOOLEAN DEFAULT true,
        created_at TIMESTAMP
    WITH
        TIME ZONE DEFAULT CURRENT_TIMESTAMP,
        updated_at TIMESTAMP
    WITH
        TIME ZONE DEFAULT CURRENT_TIMESTAMP,
        UNIQUE (
            user_id,
            platform_id,
            account_id
        )
);

CREATE INDEX idx_user_platforms_user_id ON user_platforms (user_id);

CREATE INDEX idx_user_platforms_platform_id ON user_platforms (platform_id);

CREATE INDEX idx_user_platforms_is_active ON user_platforms (is_active);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE user_platforms;
-- +goose StatementEnd