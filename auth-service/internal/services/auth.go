package services

import (
	"context"
	"fmt"
	pb "github.com/alexandear/truckgo/auth-service/generated"
	"github.com/alexandear/truckgo/auth-service/internal/models"
	"github.com/alexandear/truckgo/shared/logging"
	"github.com/golang-jwt/jwt/v5"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
	"regexp"
	"time"
)

var regexLogin = regexp.MustCompile(`^[\w-\.]+@([\w-]+\.)+[\w-]{2,4}$`)

// TODO: ctx
func (s *AuthServiceServer) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	matches := regexLogin.FindAllString(req.Login, -1)
	if len(matches) == 0 {
		s.Logger.Error("invalid login format", logging.ErrInvalidEmail, req.Login)
		return nil, status.Error(codes.InvalidArgument, fmt.Sprintf("invalid login format %s", req.Login))
	}
	// check if user exist
	var existingUser models.Auth
	if err := s.DB.Where("login = ?", req.Login).First(&existingUser).Error; err == nil {
		s.Logger.Info("user with login %s already exists", logging.ErrUserAlreadyExists, req.Login)
		return nil, fmt.Errorf("user with login %s already exists", req.Login)
	}

	hashedPassword, err := hashPassword(req.Password)
	if err != nil {
		s.Logger.Error("failed to hash password", logging.ErrPasswordHashingFailed, err)
		return nil, fmt.Errorf("failed to hash password: %v", err)
	}

	newUser := models.Auth{
		Login:     req.Login,
		Password:  hashedPassword,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err = s.DB.Create(&newUser).Error; err != nil {
		s.Logger.Error("failed to create user", logging.ErrDBCreateFailed, err)
		return nil, fmt.Errorf("failed to create user: %v", err)
	}

	return &pb.RegisterResponse{
		Message: fmt.Sprintf("User registered successfully with login %s", req.Login),
	}, nil
}

// TODO: ctx
func (s *AuthServiceServer) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	var user models.Auth
	if err := s.DB.Where("login = ?", req.Login).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			s.Logger.Error("user not found", logging.ErrUserNotFound, err)
			return nil, fmt.Errorf("user with login %s not found", req.Login)
		}
		s.Logger.Error("user not found", logging.ErrDBQueryFailed, err)
		return nil, fmt.Errorf("failed to query user: %v", err)
	}

	// checkPassword
	if err := checkPassword(user.Password, req.Password); err != nil {
		s.Logger.Error("invalid password", logging.ErrInvalidPassword, err)
		return nil, fmt.Errorf("invalid password")
	}

	// GEN JWT
	token, err := generateJWT(user.Login)
	if err != nil {
		s.Logger.Error("failed to generate token", logging.ErrTokenGenerationFailed, err)
		return nil, fmt.Errorf("failed to generate token: %v", err)
	}

	return &pb.LoginResponse{
		Token:   token,
		Message: "Login successful",
	}, nil
}

// TODO: ctx
func (s *AuthServiceServer) ValidateToken(ctx context.Context, req *pb.ValidateTokenRequest) (*pb.ValidateTokenResponse, error) {
	tokenStr := req.Token
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			s.Logger.Error("unexpected signing method", logging.ErrTokenValidationFailed, token.Header["alg"])
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return jwtSecret, nil
	})

	if err != nil {
		s.Logger.Error("unexpected signing method", logging.ErrInvalidEmail, err)
		return nil, fmt.Errorf("invalid token: %v", err)
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		login := claims["login"].(string)
		s.Logger.Info("Valid token for user: %s with ID", logging.MsgInfo, login)

		return &pb.ValidateTokenResponse{
			IsValid: true,
		}, nil
	} else {
		return &pb.ValidateTokenResponse{
			IsValid: false,
		}, nil
	}
}

// TODO: ctx
func (s *AuthServiceServer) ChangePassword(ctx context.Context, req *pb.ChangePasswordRequest) (*pb.ChangePasswordResponse, error) {
	var user models.Auth
	if err := s.DB.Where("login = ?", req.Login).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			s.Logger.Error("user not found", logging.ErrUserNotFound, err)
			return nil, fmt.Errorf("user not found")
		}
		s.Logger.Error("user not found", logging.ErrDBQueryFailed, err)
		return nil, fmt.Errorf("failed to query user: %v", err)
	}

	if err := checkPassword(user.Password, req.OldPassword); err != nil {
		s.Logger.Error("invalid old password", logging.ErrInvalidPassword, err)
		return nil, fmt.Errorf("old password is incorrect")
	}

	hashedNewPassword, err := hashPassword(req.NewPassword)
	if err != nil {
		s.Logger.Error("failed to hash new password", logging.ErrPasswordHashingFailed, err)
		return nil, fmt.Errorf("failed to hash new password: %v", err)
	}

	user.Password = hashedNewPassword
	if err := s.DB.Save(&user).Error; err != nil {
		s.Logger.Error("failed to update user", logging.ErrDBUpdateFailed, err)
		return nil, fmt.Errorf("failed to update password: %v", err)
	}

	return &pb.ChangePasswordResponse{
		Message: "Password is changed",
	}, nil
}
