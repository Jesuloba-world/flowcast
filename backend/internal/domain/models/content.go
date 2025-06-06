package models

import (
	"context"
	"time"

	"github.com/uptrace/bun"
)

type ContentStatus string

const (
	ContentStatusDraft     ContentStatus = "draft"
	ContentStatusScheduled ContentStatus = "scheduled"
	ContentStatusPublished ContentStatus = "published"
	ContentStatusFailed    ContentStatus = "failed"
	ContentStatusCancelled ContentStatus = "cancelled"
)

type Content struct {
	bun.BaseModel `bun:"table:contents,alias:c"`

	ID          string        `bun:"id,pk" json:"id"`
	UserID      string        `bun:"user_id,notnull" json:"user_id"`
	Title       string        `bun:"title" json:"title"`
	Body        string        `bun:"body,notnull" json:"body"`
	MediaURLs   []string      `bun:"media_urls,type:text[]" json:"media_urls"`
	Status      ContentStatus `bun:"status,default:'draft'" json:"status"`
	ScheduledAt *time.Time    `bun:"scheduled_at" json:"scheduled_at"`
	PublishedAt *time.Time    `bun:"published_at" json:"published_at"`
	CreatedAt   time.Time     `bun:"created_at,nullzero,notnull,default:current_timestamp" json:"created_at"`
	UpdatedAt   time.Time     `bun:"updated_at,nullzero,notnull,default:current_timestamp" json:"updated_at"`

	// Relations
	User  *User         `bun:"rel:belongs-to,join:user_id=id" json:"user,omitempty"`
	Posts []ContentPost `bun:"rel:has-many,join:id=content_id" json:"posts,omitempty"`
}

func (c *Content) BeforeAppendModel(ctx context.Context, query bun.Query) error {
	switch query.(type) {
	case *bun.InsertQuery:
		now := time.Now()
		if c.CreatedAt.IsZero() {
			c.CreatedAt = now
		}
		c.UpdatedAt = now
	case *bun.UpdateQuery:
		c.UpdatedAt = time.Now()
	}
	return nil
}

type ContentPost struct {
	bun.BaseModel `bun:"table:content_posts,alias:cp"`

	ID             string        `bun:"id,pk" json:"id"`
	ContentID      string        `bun:"content_id,notnull" json:"content_id"`
	UserPlatformID string        `bun:"user_platform_id,notnull" json:"user_platform_id"`
	PlatformPostID *string       `bun:"platform_post_id" json:"platform_post_id"`
	CustomText     *string       `bun:"custom_text" json:"custom_text"`
	Status         ContentStatus `bun:"status,default:'draft'" json:"status"`
	ScheduledAt    *time.Time    `bun:"scheduled_at" json:"scheduled_at"`
	PublishedAt    *time.Time    `bun:"published_at" json:"published_at"`
	ErrorMessage   *string       `bun:"error_message" json:"error_message"`
	CreatedAt      time.Time     `bun:"created_at,nullzero,notnull,default:current_timestamp" json:"created_at"`
	UpdatedAt      time.Time     `bun:"updated_at,nullzero,notnull,default:current_timestamp" json:"updated_at"`

	// Relations
	Content      *Content      `bun:"rel:belongs-to,join:content_id=id" json:"content,omitempty"`
	UserPlatform *UserPlatform `bun:"rel:belongs-to,join:user_platform_id=id" json:"user_platform,omitempty"`
}

func (cp *ContentPost) BeforeAppendModel(ctx context.Context, query bun.Query) error {
	switch query.(type) {
	case *bun.InsertQuery:
		now := time.Now()
		if cp.CreatedAt.IsZero() {
			cp.CreatedAt = now
		}
		cp.UpdatedAt = now
	case *bun.UpdateQuery:
		cp.UpdatedAt = time.Now()
	}
	return nil
}
