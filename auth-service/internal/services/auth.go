package services

import (
	"context"
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
	"regexp"
	"time"
)

var jwtSecret = []byte(viper.GetString("JWT_SECRET"))

type AuthServiceServer struct {
	*gorm.DB
	pb.UnimplementedAuthServiceServer
	*logging.Logger
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

var regexLogin = regexp.MustCompile(`^[\w-\.]+@([\w-]+\.)+[\w-]{2,4}$`)

func (s *AuthServiceServer) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	matches := regexLogin.FindAllString(req.Login, -1)
	if len(matches) == 0 {
		s.Logger.Error("invalid login format", "RegisterInvalidLogin", req.Login)
		return nil, status.Error(codes.InvalidArgument, fmt.Sprintf("invalid login format %s", req.Login))
	}
	// check if user exist
	var existingUser models.Auth
	if err := s.DB.Where("login = ?", req.Login).First(&existingUser).Error; err == nil {
		s.Logger.Info("user with login %s already exists", "RegisterLoginExist", req.Login)
		return nil, fmt.Errorf("user with login %s already exists", req.Login)
	}

	hashedPassword, err := hashPassword(req.Password)
	if err != nil {
		s.Logger.Error("failed to hash password", "RegisterHashPassword", err)
		return nil, fmt.Errorf("failed to hash password: %v", err)
	}

	newUser := models.Auth{
		Login:     req.Login,
		Password:  hashedPassword,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err = s.DB.Create(&newUser).Error; err != nil {
		s.Logger.Error("failed to create user", "RegisterDBError", err)
		return nil, fmt.Errorf("failed to create user: %v", err)
	}

	return &pb.RegisterResponse{
		Message: fmt.Sprintf("User registered successfully with login %s", req.Login),
	}, nil
}

func (s *AuthServiceServer) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	var user models.Auth
	if err := s.DB.Where("login = ?", req.Login).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			s.Logger.Error("user not found", "LoginNotFound", err)
			return nil, fmt.Errorf("user with login %s not found", req.Login)
		}
		s.Logger.Error("user not found", "LoginQueryError", err)
		return nil, fmt.Errorf("failed to query user: %v", err)
	}

	// checkPassword
	if err := checkPassword(user.Password, req.Password); err != nil {
		s.Logger.Error("invalid password", "LoginInvalidPassword", err)
		return nil, fmt.Errorf("invalid password")
	}

	// GEN JWT
	token, err := generateJWT(user.Login)
	if err != nil {
		s.Logger.Error("failed to generate token", "LoginJWTError", err)
		return nil, fmt.Errorf("failed to generate token: %v", err)
	}

	return &pb.LoginResponse{
		Token:   token,
		Message: "Login successful",
	}, nil
}

func (s *AuthServiceServer) ValidateToken(ctx context.Context, req *pb.ValidateTokenRequest) (*pb.ValidateTokenResponse, error) {
	tokenStr := req.Token
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			s.Logger.Error("unexpected signing method", "TokenError", token.Header["alg"])
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return jwtSecret, nil
	})

	if err != nil {
		s.Logger.Error("unexpected signing method", "InvalidToken", err)
		return nil, fmt.Errorf("invalid token: %v", err)
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		login := claims["login"].(string)
		s.Logger.Info("Valid token for user: %s with ID", "SuccessLogin", login)

		return &pb.ValidateTokenResponse{
			IsValid: true,
		}, nil
	} else {
		return &pb.ValidateTokenResponse{
			IsValid: false,
		}, nil
	}
}
