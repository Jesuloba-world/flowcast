-- +goose Up
-- +goose StatementBegin
CREATE TABLE platforms (
    id VARCHAR(21) PRIMARY KEY,
    name VARCHAR(50) UNIQUE NOT NULL,
    display_name VARCHAR(100) NOT NULL,
    icon VARCHAR(255),
    is_active BOOLEAN DEFAULT true,
    created_at TIMESTAMP
    WITH
        TIME ZONE DEFAULT CURRENT_TIMESTAMP,
        updated_at TIMESTAMP
    WITH
        TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Insert default platforms with NanoIDs
INSERT INTO
    platforms (id, name, display_name)
VALUES (
        'twitter_platform_001',
        'twitter',
        'Twitter/X'
    ),
    (
        'facebook_platform_02',
        'facebook',
        'Facebook'
    ),
    (
        'instagram_platform_3',
        'instagram',
        'Instagram'
    ),
    (
        'linkedin_platform_04',
        'linkedin',
        'LinkedIn'
    ),
    (
        'youtube_platform_005',
        'youtube',
        'YouTube'
    ),
    (
        'tiktok_platform_0006',
        'tiktok',
        'TikTok'
    );

CREATE INDEX idx_platforms_name ON platforms (name);

CREATE INDEX idx_platforms_is_active ON platforms (is_active);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE platforms;
-- +goose StatementEnd