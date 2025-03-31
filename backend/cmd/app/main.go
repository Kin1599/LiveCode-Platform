package main

import (
	"context"
	"fmt"
	_ "livecode/docs"
	"livecode/internal/app"
	"livecode/internal/config"
	"livecode/internal/handlers"
	"livecode/internal/routes"
	"livecode/internal/services/auth"
	"livecode/internal/services/filestorage"
	"livecode/internal/services/session"
	"livecode/internal/websocket"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

// @title LiveCode API
// @version 1.0
// @description LiveCode API

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @host localhost:8080
// @BasePath /
// @securityDefinitions.oauth2.password OAuth2Password
// @tokenUrl /api/login
// @in header
// @name Authorization

type Server struct {
	config        *config.Config
	authService   *auth.Auth
	sessionClient *session.SessionService
	s3Client      *filestorage.S3Client
	router        *gin.Engine
	httpServer    *http.Server
}

func NewServer(cfg *config.Config) (*Server, error) {
	storagePath := config.BuildDBConnectionString(cfg.StoragePath)

	authService, err := app.NewAuth(storagePath)
	if err != nil {
		return nil, fmt.Errorf("failed to create auth service: %w", err)
	}
	handlers.InitAuthService(authService)

	s3Client, err := filestorage.New(cfg.StoragePath.BucketName)
	if err != nil {
		return nil, fmt.Errorf("failed to create S3 client: %w", err)
	}
	handlers.InitS3Client(s3Client)

	sessionClient, err := app.NewSessionService(storagePath)
	if err != nil {
		return nil, fmt.Errorf("failed to create session service: %w", err)
	}
	handlers.InitSessionService(sessionClient)

	router := routes.SetupRouter()

	httpServer := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	return &Server{
		config:        cfg,
		authService:   authService,
		sessionClient: sessionClient,
		s3Client:      s3Client,
		router:        router,
		httpServer:    httpServer,
	}, nil
}

func (s *Server) Run(ctx context.Context) error {
	go func() {
		if err := websocket.Run(ctx); err != nil {
			log.Println("Websocket manager failed: %w", err)
		}
	}()

	go func() {
		fmt.Printf("Starting server on %s\n", s.httpServer.Addr)
		if err := s.httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Printf("HTTP server failed: %v", err)
		}
	}()

	<-ctx.Done()

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := s.httpServer.Shutdown(shutdownCtx); err != nil {
		return fmt.Errorf("failed to shutdown HTTP server: %w", err)
	}

	log.Println("Server gracefully stopped")
	return nil
}

func main() {
	cfg := config.MustLoad()

	server, err := NewServer(cfg)
	if err != nil {
		log.Fatalf("Server initialization failed: %v", err)
	}

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	if err := server.Run(ctx); err != nil {
		log.Fatalf("Server run failed: %v", err)
	}
}
