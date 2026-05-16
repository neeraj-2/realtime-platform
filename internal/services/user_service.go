package services


import (
	"errors"
	"realtime-platform/internal/models"
	"realtime-platform/internal/repository"
	"strings"


	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"

	"realtime-platform/internal/kafka"
	"realtime-platform/internal/events"
)


type UserService struct{
	UserRepo *repository.UserRepository
	RedisClient *redis.Client
	KafkaProducer *kafka.KafkaProducer
}


func NewUserService(userRepo *repository.UserRepository, redisClient *redis.Client,
	kafkaProducer *kafka.KafkaProducer) *UserService {
	return &UserService{
		UserRepo:    userRepo,
		RedisClient: redisClient,
		KafkaProducer: kafkaProducer,
	}
}

func (s *UserService) CreateUser(user *models.User) error {
	
	if strings.TrimSpace(user.Name) == "" {
		return errors.New("name cannot be empty")
	}

	if strings.TrimSpace(user.Email) == "" {
		return errors.New("email cannot be empty")
	}

	err := s.UserRepo.CreateUser(user)
	if err != nil {
		return err
	}

	// Publish user created event to Kafka
	event := events.UserCreatedEvent{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}
	err = s.KafkaProducer.Publish("user-created", event)
	if err != nil {
		fmt.Printf("Failed to publish user created event: %v\n", err)
	}

	return nil
}

func (s *UserService) GetUserByID(id int) (*models.User, error) {
	cacheKey := fmt.Sprintf("user:%d", id)

	// Try to get user from Redis cache
	cachedUser, err := s.RedisClient.Get(context.Background(), cacheKey).Result()
	if err == nil {
		var user models.User
		err = json.Unmarshal([]byte(cachedUser), &user)
		if err == nil {
			return &user, nil
		}
	}

	// If not in cache, get from database
	user, err := s.UserRepo.GetUserByID(id)
	if err != nil {
		return nil, err
	}

	// Cache the user data in Redis with an expiration time
	userJSON, err := json.Marshal(user)
	if err == nil {
		s.RedisClient.Set(context.Background(), cacheKey, userJSON, 10*time.Minute)
	}

	return user, nil
}