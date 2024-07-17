package main

import (
	"log"
	"os"

	"app/internal/post/handler"
	"app/internal/post/repository"
	"app/internal/post/service"
	"app/internal/server"
	"app/pkg/database"
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

	// Initialize post components
	postRepo := repository.NewPostgresRepository(db)
	postService := service.NewPostService(postRepo)
	postHandler := handler.NewPostHandler(postService)

	// Initialize and start server
	srv := server.NewServer(postHandler)
	srv.Run(":8080")
}
