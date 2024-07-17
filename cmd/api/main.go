package main

import (
	"log"
	"os"

	"app/internal/post/handler"
	"app/internal/post/repository"
	"app/internal/post/service"
	"app/internal/server"
	"app/pkg/database"

	"go.uber.org/zap"
)

func main() {
	// Initialize database connection
	db, err := database.NewPostgresDB(database.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
	})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()
	var logger *zap.Logger
	if os.Getenv("ENV") == "development" {
		logger = zap.NewExample()
	} else {
		logger, err = zap.NewProduction()
		if err != nil {
			log.Fatalf("Failed to create logger: %v", err)
		}
		defer logger.Sync()
	}

	// Initialize post components
	postRepo := repository.NewPostgresRepository(db)
	postService := service.NewPostService(postRepo, logger)
	postHandler := handler.NewPostHandler(postService, logger)

	// Initialize and start server
	srv := server.NewServer(postHandler)
	srv.Run(":8080")
}
