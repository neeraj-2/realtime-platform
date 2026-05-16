package main

import (
	"net/http"
	"realtime-platform/internal/config"
	"realtime-platform/internal/db"
	"realtime-platform/internal/handlers"
	"realtime-platform/internal/kafka"
	"realtime-platform/internal/middleware"
	"realtime-platform/internal/repository"
	"realtime-platform/internal/services"

	"github.com/gin-gonic/gin"
)

func main() {

	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		panic("Failed to load configuration: " + err.Error())
	}

	// Init logger
	config.InitLogger()

	// Connect to PostgreSQL
	dbConn, err := db.ConnectPostgres(cfg)

	// Connect to Redis
	redisClient := db.ConnectRedis(cfg)

	// Initialize Kafka producer
	kafkaProducer := kafka.NewKafkaProducer()

	if err != nil {
		config.Log.Fatal("Failed to connect to database: " + err.Error())
	}

	defer dbConn.Close()
	defer config.Log.Sync()

	// Initialize repositories, services, and handlers
	userRepo := repository.NewUserRepository(dbConn)
	userService := services.NewUserService(userRepo, redisClient, kafkaProducer)
	userHandler := handlers.NewUserHandler(userService)

	// Set up Gin router
	router := gin.Default()

	// Apply middleware
	router.Use(middleware.RequestLogger())

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
