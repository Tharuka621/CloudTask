package services

import (
	"errors"
	"os"
	"time"

	"cloudtask/auth-service/internal/models"
	"cloudtask/auth-service/internal/repositories"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	Register(email, password string, role models.Role) (*models.User, error)
	Login(email, password string) (string, string, error)
}

type authService struct {
	repo repositories.UserRepository
}

func NewAuthService(repo repositories.UserRepository) AuthService {
	return &authService{repo: repo}
}

func (s *authService) Register(email, password string, role models.Role) (*models.User, error) {
	count, err := s.repo.CountUsers()
	if err != nil {
		return nil, errors.New("failed to check existing users")
	}

	if count == 0 {
		role = models.AdminRole
	} else {
		role = models.MemberRole
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &models.User{
		Email:        email,
		PasswordHash: string(hashedPassword),
		Role:         role,
	}

	if err := s.repo.CreateUser(user); err != nil {
		return nil, errors.New("failed to create user, email might already exist")
	}

	return user, nil
}

func (s *authService) Login(email, password string) (string, string, error) {
	user, err := s.repo.GetUserByEmail(email)
	if err != nil {
		return "", "", errors.New("invalid credentials")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)); err != nil {
		return "", "", errors.New("invalid credentials")
	}

	// Generate Access Token
	accessToken, err := generateJWT(user.ID, string(user.Role), 15*time.Minute)
	if err != nil {
		return "", "", err
	}

	// Generate Refresh Token
	refreshToken, err := generateJWT(user.ID, string(user.Role), 7*24*time.Hour)
	if err != nil {
		return "", "", err
	}

	// Store refresh token
	rt := &models.RefreshToken{
		UserID:    user.ID,
		Token:     refreshToken,
		ExpiresAt: time.Now().Add(7 * 24 * time.Hour),
	}
	s.repo.CreateRefreshToken(rt)

	return accessToken, refreshToken, nil
}

func generateJWT(userID uint, role string, duration time.Duration) (string, error) {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		secret = "supersecret"
	}

	claims := jwt.MapClaims{
		"user_id": userID,
		"role":    role,
		"exp":     time.Now().Add(duration).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}
