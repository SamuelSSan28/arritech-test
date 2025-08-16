package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	httpHandler "arritech-user-management/internal/handler/http"
	"arritech-user-management/internal/repository/mysql"
	"arritech-user-management/internal/service"
	"arritech-user-management/pkg/database"
	"arritech-user-management/pkg/logger"
	"arritech-user-management/pkg/middleware"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		// Log but don't fail if .env file doesn't exist
		fmt.Println("Warning: .env file not found, using environment variables")
	}

	// Initialize logger
	log := logger.NewLogger()
	log.Info("Starting Arritech User Management API")

	// Initialize database
	dbConfig := database.GetConfigFromEnv()
	db, err := database.NewMySQLConnection(dbConfig)
	if err != nil {
		log.WithError(err).Fatal("Failed to connect to database")
	}

	// Run migrations
	if err := database.RunMigrations(db); err != nil {
		log.WithError(err).Fatal("Failed to run database migrations")
	}
	log.Info("Database migrations completed successfully")

	// Initialize validator
	validator := validator.New()

	// Initialize repository
	userRepo := mysql.NewUserRepository(db)

	// Initialize service
	userService := service.NewUserService(userRepo, log)

	// Initialize handler
	userHandler := httpHandler.NewUserHandler(userService, validator, log)

	// Initialize Gin router
	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(middleware.RequestLoggingMiddleware(log))
	router.Use(middleware.CORSMiddleware())

	// Health check endpoint
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":    "healthy",
			"timestamp": time.Now().Unix(),
		})
	})

	// API routes
	v1 := router.Group("/api/v1")
	{
		users := v1.Group("/users")
		{
			users.POST("", userHandler.CreateUser)
			users.GET("", userHandler.ListUsers)
			users.GET("/:id", userHandler.GetUser)
			users.PUT("/:id", userHandler.UpdateUser)
			users.DELETE("/:id", userHandler.DeleteUser)
		}
	}

	// Start server
	port := getEnv("SERVER_PORT", "8080")
	srv := &http.Server{
		Addr:    ":" + port,
		Handler: router,
	}

	// Graceful shutdown
	go func() {
		log.WithField("port", port).Info("Starting server")
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.WithError(err).Fatal("Failed to start server")
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Info("Shutting down server...")

	// Give outstanding requests 30 seconds to complete
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.WithError(err).Fatal("Server forced to shutdown")
	}

	log.Info("Server exited")
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
