-- +goose Up
-- +goose StatementBegin
CREATE TABLE analytics (
    id VARCHAR(21) PRIMARY KEY,
    content_post_id VARCHAR(21) NOT NULL REFERENCES content_posts (id) ON DELETE CASCADE,
    likes INTEGER DEFAULT 0,
    shares INTEGER DEFAULT 0,
    comments INTEGER DEFAULT 0,
    views INTEGER DEFAULT 0,
    clicks INTEGER DEFAULT 0,
    impressions INTEGER DEFAULT 0,
    engagement DECIMAL(5, 2) DEFAULT 0,
    fetched_at TIMESTAMP
    WITH
        TIME ZONE DEFAULT CURRENT_TIMESTAMP,
        created_at TIMESTAMP
    WITH
        TIME ZONE DEFAULT CURRENT_TIMESTAMP,
        updated_at TIMESTAMP
    WITH
        TIME ZONE DEFAULT CURRENT_TIMESTAMP,
        UNIQUE (content_post_id, fetched_at)
);

CREATE INDEX idx_analytics_content_post_id ON analytics (content_post_id);

CREATE INDEX idx_analytics_fetched_at ON analytics (fetched_at);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE analytics;
-- +goose StatementEnd