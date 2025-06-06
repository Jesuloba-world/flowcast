package models

import (
	"context"
	"time"

	"github.com/uptrace/bun"

)

type User struct {
	bun.BaseModel `bun:"table:users,alias:u"`

	ID        string    `bun:"id,pk" json:"id"`
	Email     string    `bun:"email,unique,notnull" json:"email"`
	Username  string    `bun:"username,unique,notnull" json:"username"`
	Password  string    `bun:"password,notnull" json:"-"`
	FirstName string    `bun:"first_name" json:"first_name"`
	LastName  string    `bun:"last_name" json:"last_name"`
	Avatar    *string   `bun:"avatar" json:"avatar"`
	Timezone  string    `bun:"timezone,default:'UTC'" json:"timezone"`
	IsActive  bool      `bun:"is_active,default:true" json:"is_active"`
	CreatedAt time.Time `bun:"created_at,nullzero,notnull,default:current_timestamp" json:"created_at"`
	UpdatedAt time.Time `bun:"updated_at,nullzero,notnull,default:current_timestamp" json:"updated_at"`
}

func (u *User) BeforeAppendModel(ctx context.Context, query bun.Query) error {
	switch query.(type) {
	case *bun.InsertQuery:
		now := time.Now()
		if u.CreatedAt.IsZero() {
			u.CreatedAt = now
		}
		u.UpdatedAt = now
	case *bun.UpdateQuery:
		u.UpdatedAt = time.Now()
	}
	return nil
}

func (u *User) IsValidForLogin() bool {
	return u.IsActive && u.Email != "" && u.Password != ""
}

func (u *User) GetFullName() string {
	return u.FirstName + " " + u.LastName
}
