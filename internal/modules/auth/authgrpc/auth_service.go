package authgrpc

import (
	"context"
	"errors"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgconn"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"goph-keeper/internal/modules/auth/authdto"
	"goph-keeper/internal/modules/auth/authservices/authjwtservice"
	"goph-keeper/internal/modules/auth/authservices/authservice"
	pb "goph-keeper/internal/modules/auth/proto"
)

// AuthGRPCServer - сервер для Auth gRPC
type AuthGRPCServer struct {
	pb.UnimplementedAuthServiceServer
	authService *authservice.AuthService
	secretKey   string
}

// NewAuthGRPCServer - создаёт новый AuthGRPCServer
func NewAuthGRPCServer(authService *authservice.AuthService, secretKey string) *AuthGRPCServer {
	return &AuthGRPCServer{
		authService: authService,
		secretKey:   secretKey,
	}
}

// Login - авторизация пользователя
func (s *AuthGRPCServer) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	inDTO, err := authdto.LoginDTOFromLoginRequestGRPC(req)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid login data: %v", err)
	}

	userID, err := s.authService.Login(ctx, inDTO)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "Login failed: %v", err)
	}

	token, err := authjwtservice.GenerateToken(userID, s.secretKey)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to generate token: %v", err)
	}

	return &pb.LoginResponse{
		Jwt: "Bearer " + token,
	}, nil
}

// Registration - регистрация пользователя
func (s *AuthGRPCServer) Registration(ctx context.Context, req *pb.RegistrationRequest) (*pb.RegistrationResponse, error) {
	regDTO, err := authdto.GetRegistrationDTOFromRegistrationRequestGRPC(req)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid registration data: %v", err)
	}

	userID, err := s.authService.Registration(ctx, regDTO)
	if err != nil {
		if isUniqueViolationError(err) {
			return nil, status.Errorf(codes.AlreadyExists, "User already registered: %v", err)
		}
		return nil, status.Errorf(codes.Internal, "Registration failed: %v", err)
	}

	token, err := authjwtservice.GenerateToken(userID, s.secretKey)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to generate token: %v", err)
	}

	return &pb.RegistrationResponse{
		Jwt: "Bearer " + token,
	}, nil
}

// Helper function to identify unique violation errors
func isUniqueViolationError(err error) bool {
	var pgErr *pgconn.PgError
	return errors.As(err, &pgErr) && pgErr.Code == pgerrcode.UniqueViolation
}
