package repositories

import (
	"context"
	"time"
)

type SessionRepository interface {
	// Token management
	StoreRefreshToken(ctx context.Context, userID, token string, expiration time.Duration) error
	GetRefreshToken(ctx context.Context, userID string) (string, error)
	DeleteRefreshToken(ctx context.Context, userID string) error
	ValidateRefreshToken(ctx context.Context, userID, token string) (bool, error)

	// Session management
	StoreUserSession(ctx context.Context, userID, sessionID string, data map[string]interface{}, expiration time.Duration) error
	GetUserSession(ctx context.Context, userID, sessionID string) (map[string]interface{}, error)
	DeleteUserSession(ctx context.Context, userID, sessionID string) error
	DeleteAllUserSession(ctx context.Context, userID string) error

	// Rate limiting
	IncrementLoginAttempts(ctx context.Context, identifier string, expiration time.Duration) (int, error)
	GetLoginAttempts(ctx context.Context, identifier string) (int, error)
	ResetLoginAttempts(ctx context.Context, identifier string) error
}
