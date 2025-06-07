package repositories

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"

	"github.com/Jesuloba-world/flowcast/internal/domain/repositories"
)

type sessionRepository struct {
	redis *redis.Client
}

func NewSessionRepository(redis *redis.Client) repositories.SessionRepository {
	return &sessionRepository{redis: redis}
}

func (r *sessionRepository) StoreRefreshToken(ctx context.Context, userID, token string, expiration time.Duration) error {
	key := fmt.Sprintf("refresh_token:%s", userID)
	return r.redis.Set(ctx, key, token, expiration).Err()
}

func (r *sessionRepository) GetRefreshToken(ctx context.Context, userID string) (string, error) {
	key := fmt.Sprintf("refresh_token:%s", userID)
	return r.redis.Get(ctx, key).Result()
}

func (r *sessionRepository) DeleteRefreshToken(ctx context.Context, userID string) error {
	key := fmt.Sprintf("refresh_token:%s", userID)
	return r.redis.Del(ctx, key).Err()
}

func (r *sessionRepository) ValidateRefreshToken(ctx context.Context, userID, token string) (bool, error) {
	storedToken, err := r.GetRefreshToken(ctx, userID)
	if err != nil {
		return false, err
	}
	return storedToken == token, nil
}

func (r *sessionRepository) StoreUserSession(ctx context.Context, userID, sessionID string, data map[string]interface{}, expiration time.Duration) error {
	key := fmt.Sprintf("session:%s:%s", userID, sessionID)
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}
	return r.redis.Set(ctx, key, jsonData, expiration).Err()
}

func (r *sessionRepository) GetUserSession(ctx context.Context, userID, sessionID string) (map[string]interface{}, error) {
	key := fmt.Sprintf("session:%s:%s", userID, sessionID)
	jsonData, err := r.redis.Get(ctx, key).Result()
	if err != nil {
		return nil, err
	}
	var data map[string]interface{}
	err = json.Unmarshal([]byte(jsonData), &data)
	return data, err
}

func (r *sessionRepository) DeleteUserSession(ctx context.Context, userID, sessionID string) error {
	key := fmt.Sprintf("session:%s:%s", userID, sessionID)
	return r.redis.Del(ctx, key).Err()
}

func (r *sessionRepository) DeleteAllUserSession(ctx context.Context, userID string) error {
	pattern := fmt.Sprintf("session:%s:*", userID)
	keys, err := r.redis.Keys(ctx, pattern).Result()
	if err != nil {
		return err
	}
	if len(keys) > 0 {
		return r.redis.Del(ctx, keys...).Err()
	}
	return nil
}

func (r *sessionRepository) IncrementLoginAttempts(ctx context.Context, identifier string, expiration time.Duration) (int, error) {
	key := fmt.Sprintf("login_attempts:%s", identifier)
	count, err := r.redis.Incr(ctx, key).Result()
	if err != nil {
		return 0, err
	}
	if count == 1 {
		r.redis.Expire(ctx, key, expiration)
	}
	return int(count), nil
}

func (r *sessionRepository) GetLoginAttempts(ctx context.Context, identifier string) (int, error) {
	key := fmt.Sprintf("login_attempts:%s", identifier)
	val, err := r.redis.Get(ctx, key).Result()
	if err != redis.Nil {
		return 0, nil
	}
	if err != nil {
		return 0, err
	}
	return strconv.Atoi(val)
}

func (r *sessionRepository) ResetLoginAttempts(ctx context.Context, identifiers string) error {
	key := fmt.Sprintf("login_attempts:%s", identifiers)
	return r.redis.Del(ctx, key).Err()
}
