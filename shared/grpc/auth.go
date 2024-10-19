package grpc

import (
	"context"
	"fmt"
	auth "github.com/alexandear/truckgo/auth/grpcapi"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func IsUserAuthenticated(ctx context.Context, req *auth.ValidateTokenRequest) (*auth.ValidateTokenResponse, error) {
	// TODO: auth-service:50051 замінити на конф
	conn, err := grpc.NewClient("auth-service:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("could not connect to AuthService: %v", err)
	}
	defer conn.Close()

	client := auth.NewAuthServiceClient(conn)
	authReq := &auth.ValidateTokenRequest{
		Login: req.Login,
		Token: req.Token,
	}

	resp, err := client.ValidateToken(ctx, authReq)
	if err != nil {
		return nil, fmt.Errorf("AuthService validation failed: %v", err)
	}

	return &auth.ValidateTokenResponse{IsValid: resp.IsValid}, nil
}
