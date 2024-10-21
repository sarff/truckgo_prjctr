package services

import (
	"context"
	"fmt"
	authpb "github.com/alexandear/truckgo/auth/grpcapi"
	"github.com/alexandear/truckgo/auth/internal/models"
	"github.com/alexandear/truckgo/shared/logging"
	"github.com/golang-jwt/jwt/v5"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
	"regexp"
	"time"
)

var regexLogin = regexp.MustCompile(`^[\w-\.]+@([\w-]+\.)+[\w-]{2,4}$`)

type AuthServiceServer struct {
	*gorm.DB
	authpb.UnimplementedAuthServiceServer
	*logging.Logger
	*models.Auth
}

func (s *AuthServiceServer) Register(ctx context.Context, req *authpb.RegisterRequest) (*authpb.RegisterResponse, error) {
	matches := regexLogin.FindAllString(req.Login, -1)
	if len(matches) == 0 {
		s.Logger.Error("invalid Login format", logging.ErrInvalidEmail, req.Login)
		return nil, status.Errorf(codes.InvalidArgument, "invalid Login format %s", req.Login)
	}
	if req.TypeUser != "driver" && req.TypeUser != "customer" {
		return nil, status.Errorf(codes.InvalidArgument, "invalid Login format %s", req.Login)
	}

	err := s.checkUserByLogin(req.Login)
	if err == nil {
		s.Logger.Info("user with Login %s already exists", logging.ErrUserAlreadyExists, req.Login)
		return nil, status.Errorf(codes.AlreadyExists, "user with Login %s already exists", req.Login)
	}

	hashedPassword, err := HashPassword(req.Password)
	if err != nil {
		s.Logger.Error("failed to hash password", logging.ErrPasswordHashingFailed, err)
		return nil, status.Errorf(codes.Internal, "failed to hash password: %v", err)
	}

	newUser := models.Auth{
		Login:     req.Login,
		Password:  hashedPassword,
		TypeUser:  req.TypeUser,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err = s.DB.Create(&newUser).Error; err != nil {
		s.Logger.Error("failed to create user", logging.ErrDBCreateFailed, err)
		return nil, status.Errorf(codes.Internal, "failed to create user: %v", err)
	}

	err = s.GrpcConnCreateUser(ctx, &newUser)
	if err != nil {
		s.Logger.Error("failed to create user for User service", logging.ErrDBCreateFailed, err)
		return nil, status.Errorf(codes.Internal, "failed to create user for User service: %v", err)
	}
	return &authpb.RegisterResponse{
		Message: fmt.Sprintf("User %s registered successfully ", req.Login),
	}, nil
}

func (s *AuthServiceServer) Login(_ context.Context, req *authpb.LoginRequest) (*authpb.LoginResponse, error) {
	err := s.checkUserByLogin(req.Login)
	if err != nil {
		return nil, err
	}

	// checkPassword
	if err := checkPassword(s.Auth.Password, req.Password); err != nil {
		s.Logger.Error("invalid password", logging.ErrInvalidPassword, err)
		return nil, status.Error(codes.Internal, "invalid password")
	}

	// GEN JWT
	token, err := generateJWT(s.Auth.Login)
	s.Logger.Info(s.Auth.Login)
	if err != nil {
		s.Logger.Error("failed to generate token", logging.ErrTokenGenerationFailed, err)
		return nil, status.Errorf(codes.Internal, "failed to generate token: %v", err)
	}

	return &authpb.LoginResponse{
		Token:   token,
		Message: "Login successful",
	}, nil
}

func (s *AuthServiceServer) ValidateToken(_ context.Context, req *authpb.ValidateTokenRequest) (*authpb.ValidateTokenResponse, error) {
	tokenStr := req.Token
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			s.Logger.Error("unexpected signing method", logging.ErrTokenValidationFailed, token.Header["alg"])
			return nil, status.Errorf(codes.Unknown, "unexpected signing method: %v", token.Header["alg"])
		}
		return jwtSecret, nil
	})

	if err != nil {
		s.Logger.Error("unexpected signing method", logging.ErrInvalidToken, err)
		return nil, status.Errorf(codes.InvalidArgument, "invalid token: %v", err)
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		login, exists := claims["Login"].(string)
		if !exists {
			s.Logger.Error("Login not found in token claims", logging.ErrInvalidToken, "Login not found")
			return &authpb.ValidateTokenResponse{
				IsValid: false,
			}, nil
		}
		if login != req.Login {
			s.Logger.Error("wrong Login for this token", logging.ErrInvalidEmail, login)
			return &authpb.ValidateTokenResponse{
				IsValid: false,
			}, nil
		}

		return &authpb.ValidateTokenResponse{
			IsValid: true,
		}, nil
	} else {
		return &authpb.ValidateTokenResponse{
			IsValid: false,
		}, nil
	}
}

func (s *AuthServiceServer) ChangePassword(ctx context.Context, req *authpb.ChangePasswordRequest) (*authpb.ChangePasswordResponse, error) {
	tokenValidationRes, err := s.ValidateToken(ctx, &authpb.ValidateTokenRequest{Token: req.Token})
	if err != nil || !tokenValidationRes.IsValid {
		return nil, status.Error(codes.InvalidArgument, "invalid or expired token")
	}

	err = s.checkUserByLogin(req.Login)
	if err != nil {
		return nil, err
	}
	if err := checkPassword(s.Auth.Password, req.OldPassword); err != nil {
		s.Logger.Error("invalid old password", logging.ErrInvalidPassword, err)
		return nil, status.Error(codes.InvalidArgument, "old password is incorrect")
	}

	hashedNewPassword, err := HashPassword(req.NewPassword)
	if err != nil {
		s.Logger.Error("failed to hash new password", logging.ErrPasswordHashingFailed, err)
		return nil, status.Errorf(codes.Internal, "failed to hash new password: %v", err)
	}

	s.Auth.Password = hashedNewPassword
	if err = s.DB.Save(&s.Auth).Error; err != nil {
		s.Logger.Error("failed to update user", logging.ErrDBUpdateFailed, err)
		return nil, status.Errorf(codes.Internal, "failed to update password: %v", err)
	}

	return &authpb.ChangePasswordResponse{
		Message: "Password is changed",
	}, nil
}
