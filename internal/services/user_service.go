package services


import (
	"errors"
	"realtime-platform/internal/models"
	"realtime-platform/internal/repository"
	"strings"
)


type UserService struct{
	UserRepo *repository.UserRepository
}


func NewUserService(userRepo *repository.UserRepository) *UserService {
	return &UserService{
		UserRepo: userRepo,
	}
}

func (s *UserService) CreateUser(user *models.User) error {
	
	if strings.TrimSpace(user.Name) == "" {
		return errors.New("name cannot be empty")
	}

	if strings.TrimSpace(user.Email) == "" {
		return errors.New("email cannot be empty")
	}

	return s.UserRepo.CreateUser(user)
}

func (s *UserService) GetUserByID(id int) (*models.User, error) {
	return s.UserRepo.GetUserByID(id)
}