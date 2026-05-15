package repository

import (
	"database/sql"
	"realtime-platform/internal/models"
)

type UserRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		DB: db,
	}
}

func (r *UserRepository) CreateUser(user *models.User) error {

	query := `
	INSERT INTO users(name, email)
	VALUES($1, $2)
	RETURNING id, created_at
	`

	err := r.DB.QueryRow(
		query,
		user.Name,
		user.Email,
	).Scan(&user.ID, &user.CreatedAt)

	return err
}

func (r *UserRepository) GetUserByID(id int) (*models.User, error) {

	query := `
	SELECT id, name, email, created_at
	FROM users
	WHERE id = $1
	`

	user := &models.User{}

	err := r.DB.QueryRow(query, id).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return user, nil
}