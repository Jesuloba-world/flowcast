package models

import (
	"context"
	"time"

	"github.com/uptrace/bun"
)

type Analytics struct {
	bun.BaseModel `bun:"table:analytics,alias:a"`

	ID            string    `bun:"id,pk" json:"id"`
	ContentPostID string    `bun:"content_post_id,notnull" json:"content_post_id"`
	Likes         int       `bun:"likes,default:0" json:"likes"`
	Shares        int       `bun:"shares,default:0" json:"shares"`
	Comments      int       `bun:"comments,default:0" json:"comments"`
	Views         int       `bun:"views,default:0" json:"views"`
	Clicks        int       `bun:"clicks,default:0" json:"clicks"`
	Impressions   int       `bun:"impressions,default:0" json:"impressions"`
	Engagement    float64   `bun:"engagement,default:0" json:"engagement"`
	FetchedAt     time.Time `bun:"fetched_at,nullzero,notnull,default:current_timestamp" json:"fetched_at"`
	CreatedAt     time.Time `bun:"created_at,nullzero,notnull,default:current_timestamp" json:"created_at"`
	UpdatedAt     time.Time `bun:"updated_at,nullzero,notnull,default:current_timestamp" json:"updated_at"`

	// Relations
	ContentPost *ContentPost `bun:"rel:belongs-to,join:content_post_id=id" json:"content_post,omitempty"`
}

func (a *Analytics) BeforeAppendModel(ctx context.Context, query bun.Query) error {
	switch query.(type) {
	case *bun.InsertQuery:
		now := time.Now()
		if a.CreatedAt.IsZero() {
			a.CreatedAt = now
		}
		a.UpdatedAt = now
	case *bun.UpdateQuery:
		a.UpdatedAt = time.Now()
	}
	return nil
}
