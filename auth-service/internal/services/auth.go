package services

import (
	"context"
	"fmt"
	pb "github.com/alexandear/truckgo/auth-service/generated"
)

type AuthServiceServer struct {
	pb.UnimplementedAuthServiceServer
}

func (s *AuthServiceServer) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	// Логіка реєстрації користувача
	fmt.Printf("Registering user: %s\n", req.Login)
	// Додай збереження користувача в базу даних тут

	return &pb.RegisterResponse{
		Message: "User registered successfully",
	}, nil
}

func (s *AuthServiceServer) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	// Логіка авторизації
	fmt.Printf("Logging in user: %s\n", req.Login)
	// Перевір пароль користувача та згенеруй JWT токен

	return &pb.LoginResponse{
		Token:   "generated_jwt_token", // заміни на згенерований токен
		Message: "Login successful",
	}, nil
}

func (s *AuthServiceServer) ValidateToken(ctx context.Context, req *pb.ValidateTokenRequest) (*pb.ValidateTokenResponse, error) {
	// Логіка перевірки токена
	fmt.Printf("Validating token: %s\n", req.Token)
	// Перевір валідність токена

	return &pb.ValidateTokenResponse{
		IsValid: true, // заміни на реальний результат перевірки
	}, nil
}
