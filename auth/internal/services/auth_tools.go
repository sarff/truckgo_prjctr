package services

import (
	"fmt"
	"github.com/alexandear/truckgo/shared/logging"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
	"os"
	"time"
)

var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

func hashPassword(password string) (string, error) {
	if len(password) < 8 {
		return "", status.Error(codes.InvalidArgument, "password must be at least 8 characters")
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func checkPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func generateJWT(login string) (string, error) {
	claims := jwt.MapClaims{
		"Login": login,
		"exp":   time.Now().Add(time.Hour * 72).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (s *AuthServiceServer) checkUserByLogin(login string) error {
	if err := s.DB.Where("Login = ?", login).First(&s.Auth).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			s.Logger.Error("user not found", logging.ErrUserNotFound, err)
			return fmt.Errorf("user with Login %s not found", login)
		}
		s.Logger.Error("user not found", logging.ErrDBQueryFailed, err)
		return fmt.Errorf("failed to query user: %v", err)
	}
	return nil
}
