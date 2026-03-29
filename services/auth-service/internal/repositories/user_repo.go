package repositories

import (
	"cloudtask/auth-service/internal/models"
	DB "cloudtask/auth-service/pkg/database"
)

type UserRepository interface {
	CreateUser(user *models.User) error
	GetUserByEmail(email string) (*models.User, error)
	CreateRefreshToken(token *models.RefreshToken) error
	CountUsers() (int64, error)
}

type userRepository struct{}

func NewUserRepository() UserRepository {
	return &userRepository{}
}

func (r *userRepository) CreateUser(user *models.User) error {
	return DB.DB.Create(user).Error
}

func (r *userRepository) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	if err := DB.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) CreateRefreshToken(token *models.RefreshToken) error {
	return DB.DB.Create(token).Error
}

func (r *userRepository) CountUsers() (int64, error) {
	var count int64
	if err := DB.DB.Model(&models.User{}).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}
