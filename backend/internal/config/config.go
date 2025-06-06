package config

import (
	"fmt"
	"strings"

	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"

	"github.com/Jesuloba-world/flowcast/internal/logger"
)

type Config struct {
	Server    ServerConfig    `koanf:"server"`
	Database  DatabaseConfig  `koanf:"database"`
	Dragonfly DragonflyConfig `koanf:"dragonfly"`
	Auth      AuthConfig      `koanf:"auth"`
	Social    SocialConfig    `koanf:"social"`
	Logging   logger.Config   `koanf:"logging"`
}

type ServerConfig struct {
	Port        string     `koanf:"port"`
	Environment string     `koanf:"environment"`
	CORS        CORSConfig `koanf:"cors"`
}

type CORSConfig struct {
	AllowedOrigins []string `koanf:"allowed_origins"`
	AllowedMethods []string `koanf:"allowed_methods"`
	AllowedHeaders []string `koanf:"allowed_headers"`
}

type DatabaseConfig struct {
	URL             string `koanf:"url"`
	MaxOpenConns    int    `koanf:"max_open_conns"`
	MaxIdleConns    int    `koanf:"max_idle_conns"`
	ConnMaxLifetime string `koanf:"conn_max_lifetime"`
}

type DragonflyConfig struct {
	URL          string `koanf:"url"`
	Password     string `koanf:"password"`
	PoolSize     int    `koanf:"pool_size"`
	MinIdleConns int    `koanf:"min_idle_conns"`
	MaxRetries   int    `koanf:"max_retries"`
}

type AuthConfig struct {
	JWTSecret          string `koanf:"jwt_secret"`
	JWTExpirationHours int    `koanf:"jwt_expiration_hours"`
	RefreshTokenDays   int    `koanf:"refresh_token_days"`
}

type SocialConfig struct {
	Twitter   TwitterConfig   `koanf:"twitter"`
	Instagram InstagramConfig `koanf:"instagram"`
	LinkedIn  LinkedInConfig  `koanf:"linkedin"`
	Facebook  FacebookConfig  `koanf:"facebook"`
}

type TwitterConfig struct {
	APIKey      string `koanf:"api_key"`
	APISecret   string `koanf:"api_secret"`
	BearerToken string `koanf:"bearer_token"`
	RateLimit   int    `koanf:"rate_limit"`
}

type InstagramConfig struct {
	ClientID     string `koanf:"client_id"`
	ClientSecret string `koanf:"client_secret"`
	RedirectURL  string `koanf:"redirect_url"`
}

type LinkedInConfig struct {
	ClientID     string `koanf:"client_id"`
	ClientSecret string `koanf:"client_secret"`
	RedirectURL  string `koanf:"redirect_url"`
}

type FacebookConfig struct {
	AppID       string `koanf:"app_id"`
	AppSecret   string `koanf:"app_secret"`
	RedirectURL string `koanf:"redirect_url"`
}

func Load() (*Config, error) {
	k := koanf.New(".")

	setDefaults(k)

	// load from yaml: for development
	if err := k.Load(file.Provider("configs/config.yaml"), yaml.Parser()); err != nil {
		fmt.Printf("Warning: Could not load config file: %v\n", err)
	}

	// load from environment variables: for production
	if err := k.Load(env.Provider("FLOWCAST_", ".", func(s string) string {
		return strings.Replace(strings.ToLower(
			strings.TrimPrefix(s, "FLOWCAST_")), "_", ".", -1)
	}), nil); err != nil {
		fmt.Printf("Warning: Could not load config from environment variables: %v\n", err)
	}

	var cfg Config
	if err := k.Unmarshal("", &cfg); err != nil {
		return nil, fmt.Errorf("error unmarshalling config: %v", err)
	}

	return &cfg, nil
}

func setDefaults(k *koanf.Koanf) {
	// Server defaults
	k.Set("server.port", "8080")
	k.Set("server.environment", "development")
	k.Set("server.cors.allowed_origins", []string{"http://localhost:3000"})
	k.Set("server.cors.allowed_methods", []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"})
	k.Set("server.cors.allowed_headers", []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"})

	// Database defaults
	k.Set("database.url", "postgres://flowcast_user:flowcast_password@localhost:5432/flowcast?sslmode=disable")
	k.Set("database.max_open_conns", 25)
	k.Set("database.max_idle_conns", 5)
	k.Set("database.conn_max_lifetime", "5m")

	// Dragonfly defaults
	k.Set("dragonfly.url", "redis://localhost:6379")
	k.Set("dragonfly.pool_size", 20)
	k.Set("dragonfly.min_idle_conns", 5)
	k.Set("dragonfly.max_retries", 3)

	// Auth defaults
	k.Set("auth.jwt_secret", "your-super-secret-jwt-key")
	k.Set("auth.jwt_expiration_hours", 24)
	k.Set("auth.refresh_token_days", 30)
	k.Set("auth.password_min_length", 8)

	// Logging defaults
	k.Set("logging.level", "info")
	k.Set("logging.format", "json")
	k.Set("logging.output", "stdout")
}
