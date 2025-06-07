package server

import (
	"net/http"
	"time"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humachi"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/redis/go-redis/v9"

	"github.com/Jesuloba-world/flowcast/internal/config"
	"github.com/Jesuloba-world/flowcast/internal/infrastructure/database"
	"github.com/Jesuloba-world/flowcast/internal/logger"
)

type Server struct {
	config *config.Config
	db     *database.DB
	redis  *redis.Client
	logger *logger.Logger
}

func New(cfg *config.Config, db *database.DB, redis *redis.Client, logger *logger.Logger) *Server {
	return &Server{
		config: cfg,
		db:     db,
		redis:  redis,
		logger: logger,
	}
}

func (s *Server) Router() http.Handler {
	r := chi.NewRouter()

	// middleware
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))

	// CORS
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   s.config.Server.CORS.AllowedOrigins,
		AllowedMethods:   s.config.Server.CORS.AllowedMethods,
		AllowedHeaders:   s.config.Server.CORS.AllowedHeaders,
		MaxAge:           300,
		AllowCredentials: true,
	}))

	r.Use(middleware.Heartbeat("/health"))

	_ = humachi.New(r, huma.DefaultConfig("Flowcast API", "1.0.0"))
	// register routes

	return r
}
