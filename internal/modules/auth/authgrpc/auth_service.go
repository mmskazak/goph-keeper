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

// Login авторизация пользователя в приложении
func (s *AuthGRPCServer) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	inDTO, err := authdto.LoginDTOFromLoginRequestGRPC(req)
	if err != nil {
		return &pb.LoginResponse{
			Status:  "error",
			Message: "Invalid login data",
		}, nil
	}

	userID, err := s.authService.Login(ctx, inDTO)
	if err != nil {
		return &pb.LoginResponse{
			Status:  "error",
			Message: "Login failed",
		}, nil
	}

	token, err := authjwtservice.GenerateToken(userID, s.secretKey)
	if err != nil {
		return &pb.LoginResponse{
			Status:  "error",
			Message: "Failed to generate token",
		}, nil
	}

	return &pb.LoginResponse{
		Status: "success",
		Jwt:    "Bearer " + token,
	}, nil
}

// Registration обрабатывает регистрацию пользователя
func (s *AuthGRPCServer) Registration(ctx context.Context, req *pb.RegistrationRequest) (*pb.RegistrationResponse, error) {
	// Преобразуем gRPC-запрос в DTO
	regDTO, err := authdto.GetRegistrationDTOFromRegistrationRequestGRPC(req)
	if err != nil {
		return &pb.RegistrationResponse{
			Status:  "error",
			Message: "Invalid registration data",
		}, fmt.Errorf("failed to convert RegistrationRequest to RegistrationDTO: %w", err)
	}

	// Регистрация пользователя
	userID, err := s.authService.Registration(ctx, regDTO)
	if err != nil {
		return &pb.RegistrationResponse{
			Status:  "error",
			Message: "User was already registered",
		}, fmt.Errorf("registration failed: %w", err)
	}

	// Генерация токена
	token, err := authjwtservice.GenerateToken(userID, s.secretKey)
	if err != nil {
		return &pb.RegistrationResponse{
			Status:  "error",
			Message: "Failed to generate token",
		}, fmt.Errorf("failed to generate token: %w", err)
	}

	return &pb.RegistrationResponse{
		Status:  "success",
		Message: "User registered successfully",
		Jwt:     "Bearer " + token,
	}, nil
}
