-- +goose Up
-- +goose StatementBegin
CREATE TABLE content_posts (
    id VARCHAR(21) PRIMARY KEY,
    content_id VARCHAR(21) NOT NULL REFERENCES contents (id) ON DELETE CASCADE,
    user_platform_id VARCHAR(21) NOT NULL REFERENCES user_platforms (id) ON DELETE CASCADE,
    platform_post_id VARCHAR(255),
    custom_text TEXT,
    status content_status DEFAULT 'draft',
    scheduled_at TIMESTAMP
    WITH
        TIME ZONE,
        published_at TIMESTAMP
    WITH
        TIME ZONE,
        error_message TEXT,
        created_at TIMESTAMP
    WITH
        TIME ZONE DEFAULT CURRENT_TIMESTAMP,
        updated_at TIMESTAMP
    WITH
        TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_content_posts_content_id ON content_posts (content_id);

CREATE INDEX idx_content_posts_user_platform_id ON content_posts (user_platform_id);

CREATE INDEX idx_content_posts_status ON content_posts (status);

CREATE INDEX idx_content_posts_scheduled_at ON content_posts (scheduled_at);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE content_posts;
-- +goose StatementEnd