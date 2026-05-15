package main

import (
	"net/http"
	"realtime-platform/internal/db"
	"realtime-platform/internal/handlers"
	"realtime-platform/internal/repository"
	"realtime-platform/internal/services"

	"github.com/gin-gonic/gin"
)

func main() {

	// Connect to PostgreSQL
	dbConn, err := db.ConnectPostgres()

	if err != nil {
		panic("Failed to connect to database: " + err.Error())
	}

	defer dbConn.Close()

	// Initialize repositories, services, and handlers
	userRepo := repository.NewUserRepository(dbConn)
	userService := services.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)

	// Set up Gin router
	router := gin.Default()


	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	// User routes
	router.POST("/users", userHandler.CreateUser)
	router.GET("/users/:id", userHandler.GetUserByID)

	// Start the server
	if err := router.Run(":8080"); err != nil {
		panic("Failed to start server: " + err.Error())
	}
}	