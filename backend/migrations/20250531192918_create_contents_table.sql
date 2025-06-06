-- +goose Up
-- +goose StatementBegin
CREATE TYPE content_status AS ENUM(
    'draft',
    'scheduled',
    'published',
    'failed',
    'cancelled'
);

CREATE TABLE contents (
    id VARCHAR(21) PRIMARY KEY,
    user_id VARCHAR(21) NOT NULL REFERENCES users (id) ON DELETE CASCADE,
    title VARCHAR(255),
    body TEXT NOT NULL,
    media_urls TEXT[],
    status content_status DEFAULT 'draft',
    scheduled_at TIMESTAMP WITH TIME ZONE,
    published_at TIMESTAMP WITH TIME ZONE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_contents_user_id ON contents (user_id);

CREATE INDEX idx_contents_status ON contents (status);

CREATE INDEX idx_contents_scheduled_at ON contents (scheduled_at);

CREATE INDEX idx_contents_created_at ON contents (created_at);

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE contents;

DROP TYPE content_status;

-- +goose StatementEnd