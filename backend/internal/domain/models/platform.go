package models

import (
	"context"
	"time"

	"github.com/uptrace/bun"
)

type Platform struct {
	bun.BaseModel `bun:"table:platforms,alias:p"`

	ID          string    `bun:"id,pk" json:"id"`
	Name        string    `bun:"name,unique,notnull" json:"name"`
	DisplayName string    `bun:"display_name,notnull" json:"display_name"`
	Icon        string    `bun:"icon" json:"icon"`
	IsActive    bool      `bun:"is_active,default:true" json:"is_active"`
	CreatedAt   time.Time `bun:"created_at,nullzero,notnull,default:current_timestamp" json:"created_at"`
	UpdatedAt   time.Time `bun:"updated_at,nullzero,notnull,default:current_timestamp" json:"updated_at"`
}

func (p *Platform) BeforeAppendModel(ctx context.Context, query bun.Query) error {
	switch query.(type) {
	case *bun.InsertQuery:
		now := time.Now()
		if p.CreatedAt.IsZero() {
			p.CreatedAt = now
		}
		p.UpdatedAt = now
	case *bun.UpdateQuery:
		p.UpdatedAt = time.Now()
	}
	return nil
}

type UserPlatform struct {
	bun.BaseModel `bun:"table:user_platforms,alias:up"`

	ID           string     `bun:"id,pk" json:"id"`
	UserID       string     `bun:"user_id,notnull" json:"user_id"`
	PlatformID   string     `bun:"platform_id,notnull" json:"platform_id"`
	AccessToken  string     `bun:"access_token,notnull" json:"-"`
	RefreshToken *string    `bun:"refresh_token" json:"-"`
	ExpiresAt    *time.Time `bun:"expires_at" json:"expires_at"`
	AccountID    string     `bun:"account_id,notnull" json:"account_id"`
	AccountName  string     `bun:"account_name" json:"account_name"`
	IsActive     bool       `bun:"is_active,default:true" json:"is_active"`
	CreatedAt    time.Time  `bun:"created_at,nullzero,notnull,default:current_timestamp" json:"created_at"`
	UpdatedAt    time.Time  `bun:"updated_at,nullzero,notnull,default:current_timestamp" json:"updated_at"`

	// Relation
	User     *User     `bun:"rel:belongs-to,join:user_id=id" json:"user,omitempty"`
	Platform *Platform `bun:"rel:belongs-to,join:platform_id=id" json:"platform,omitempty"`
}

func (up *UserPlatform) BeforeAppendModel(ctx context.Context, query bun.Query) error {
	switch query.(type) {
	case *bun.InsertQuery:
		now := time.Now()
		if up.CreatedAt.IsZero() {
			up.CreatedAt = now
		}
		up.UpdatedAt = now
	case *bun.UpdateQuery:
		up.UpdatedAt = time.Now()
	}
	return nil
}
