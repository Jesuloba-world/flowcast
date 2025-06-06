package dragonfly

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"

	"github.com/Jesuloba-world/flowcast/internal/config"
)

type Client struct {
	*redis.Client
}

func New(cfg config.DragonflyConfig) (*Client, error) {
	opt, err := redis.ParseURL(cfg.URL)
	if err != nil {
		return nil, fmt.Errorf("failed to parse DragonflyDB URL: %w", err)
	}

	if cfg.Password != "" {
		opt.Password = cfg.Password
	}
	opt.PoolSize = cfg.PoolSize
	opt.MinIdleConns = cfg.MinIdleConns
	opt.MaxRetries = cfg.MaxRetries

	// optimizations
	opt.ReadTimeout = 3 * time.Second
	opt.WriteTimeout = 3 * time.Second
	opt.PoolTimeout = 4 * time.Second
	opt.ConnMaxIdleTime = 5 * time.Minute

	db := redis.NewClient(opt)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := db.Ping(ctx).Err(); err != nil {
		return nil, fmt.Errorf("failed to ping DragonflyDB: %w", err)
	}

	return &Client{db}, nil
}

func (c *Client) Close() error {
	return c.Client.Close()
}

func (c *Client) SetWithExpiration(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	return c.Client.Set(ctx, key, value, expiration).Err()
}

func (c *Client) GetJSON(ctx context.Context, key string) (string, error) {
	return c.Get(ctx, key).Result()
}

func (c *Client) SetJSON(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	return c.Set(ctx, key, value, expiration).Err()
}
