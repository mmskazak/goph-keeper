package authgrpc

import (
	"context"
	"fmt"
	"goph-keeper/internal/modules/auth/authdto"
	"goph-keeper/internal/modules/auth/authservices/authjwtservice"
	"goph-keeper/internal/modules/auth/authservices/authservice"
	pb "goph-keeper/internal/modules/auth/proto"
)

//go:generate protoc --proto_path=../proto --go_out=. --go-grpc_out=. auth.proto

// AuthGRPCServer сервер
type AuthGRPCServer struct {
	pb.UnimplementedAuthServiceServer

	authService *authservice.AuthService
	secretKey   string
}

func NewAuthGRPCServer(authService *authservice.AuthService, secretKey string) *AuthGRPCServer {
	return &AuthGRPCServer{
		authService: authService,
		secretKey:   secretKey,
	}
}

// Login обрабатывает вход пользователя
func (s *AuthGRPCServer) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	inDTO, err := authdto.LoginDTOFromLoginRequestGRPC(req)
	if err != nil {
		return nil, fmt.Errorf("error converting login request to dto: %w", err)
	}

	userID, err := s.authService.Login(ctx, inDTO)
	if err != nil {
		return nil, fmt.Errorf("login failed: %w", err)
	}

	token, err := authjwtservice.GenerateToken(userID, s.secretKey)
	if err != nil {
		return nil, fmt.Errorf("failed to generate token: %w", err)
	}

	return &pb.LoginResponse{
		Jwt: token,
	}, nil
}

// Registration обрабатывает регистрацию пользователя
func (s *AuthGRPCServer) Registration(ctx context.Context, req *pb.RegistrationRequest) (*pb.RegistrationResponse, error) {
	// Преобразуем gRPC-запрос в DTO
	regDTO, err := authdto.GetRegistrationDTOFromRegistrationRequestGRPC(req)
	if err != nil {
		return nil, fmt.Errorf("failed to convert RegistrationRequest to RegistrationDTO: %w", err)
	}

	// Регистрация пользователя
	userID, err := s.authService.Registration(ctx, regDTO)
	if err != nil {
		return nil, fmt.Errorf("registration failed: %w", err)
	}

	// Генерация токена
	token, err := authjwtservice.GenerateToken(userID, s.secretKey)
	if err != nil {
		return nil, fmt.Errorf("failed to generate token: %w", err)
	}

	return &pb.RegistrationResponse{
		Jwt: token,
	}, nil
}
