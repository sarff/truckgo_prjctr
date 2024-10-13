package services

import (
	"fmt"
	pb "github.com/alexandear/truckgo/auth-service/generated"
	"github.com/alexandear/truckgo/auth-service/internal/models"
	"github.com/alexandear/truckgo/shared/logging"
	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
	"time"
)

var jwtSecret = []byte(viper.GetString("JWT_SECRET"))

type AuthServiceServer struct {
	*gorm.DB
	pb.UnimplementedAuthServiceServer
	*logging.Logger
	*models.Auth
}

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
		"login": login,
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
	if err := s.DB.Where("login = ?", login).First(&s.Auth).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			s.Logger.Error("user not found", logging.ErrUserNotFound, err)
			return fmt.Errorf("user with login %s not found", login)
		}
		s.Logger.Error("user not found", logging.ErrDBQueryFailed, err)
		return fmt.Errorf("failed to query user: %v", err)
	}
	return nil
}
