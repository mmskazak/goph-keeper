package pwdgrpc

import (
	"context"
	"errors"
	"fmt"
	"google.golang.org/protobuf/types/known/wrapperspb"
	"goph-keeper/internal/helpers"
	pb "goph-keeper/internal/modules/pwd/proto"
	"goph-keeper/internal/modules/pwd/pwddto"
	"goph-keeper/internal/modules/pwd/pwdservices"
	"goph-keeper/internal/modules/pwd/valueobj"

	"go.uber.org/zap"

	"github.com/jackc/pgx/v5"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

//go:generate protoc --proto_path=../proto --go_out=. --go-grpc_out=. pwd.proto

var ErrUpdatedRecordNotFound = errors.New("updated record not found")

const ErrParseJWTFailed = "parse jwt failed: %w"

// PasswordGRPCServer - сервис GRPC отвечающий за работу с паролями.
type PasswordGRPCServer struct {
	pb.UnimplementedPasswordServiceServer

	zapLogger  *zap.SugaredLogger
	pwdService pwdservices.IPwdService
	secretKey  string
}

// NewPasswordGRPCServer - создаёт новый PasswordGRPCServer.
func NewPasswordGRPCServer(
	service pwdservices.IPwdService,
	secretKey string,
	zapLogger *zap.SugaredLogger,
) *PasswordGRPCServer {
	return &PasswordGRPCServer{
		pwdService: service,
		secretKey:  secretKey,
		zapLogger:  zapLogger,
	}
}

// SavePassword сохраняет пароль.
func (s *PasswordGRPCServer) SavePassword(ctx context.Context, req *pb.SavePwdRequest) (*pb.BasicResponse, error) {
	userID, err := helpers.ParseTokenAndExtractUserID(req.GetJwt().GetValue(), s.secretKey)
	if err != nil {
		return nil, fmt.Errorf(ErrParseJWTFailed, err)
	}

	savePwdDTO := pwddto.SavePwdDTO{
		UserID: userID,
		Title:  req.GetTitle().GetValue(),
		Credentials: valueobj.Credentials{
			Login:    req.GetCredentials().GetLogin().GetValue(),
			Password: req.GetCredentials().GetPassword().GetValue(),
		},
	}

	if err := s.pwdService.SavePassword(ctx, &savePwdDTO); err != nil {
		s.zapLogger.Errorf("Error in SavePassword: %v", err)
		return nil, status.Errorf(codes.Internal, "Failed to save password: %v", err)
	}

	return &pb.BasicResponse{}, nil
}

// UpdatePassword обновляет пароль.
func (s *PasswordGRPCServer) UpdatePassword(ctx context.Context, req *pb.UpdatePwdRequest) (*pb.BasicResponse, error) {
	userID, err := helpers.ParseTokenAndExtractUserID(req.GetJwt().GetValue(), s.secretKey)
	if err != nil {
		return nil, fmt.Errorf(ErrParseJWTFailed, err)
	}

	updatePwdDTO := pwddto.UpdatePwdDTO{
		UserID: userID,
		PwdID:  req.GetPwdId().GetValue(),
		Title:  req.GetTitle().GetValue(),
		Credentials: valueobj.Credentials{
			Login:    req.GetCredentials().GetLogin().GetValue(),
			Password: req.GetCredentials().GetPassword().GetValue(),
		},
	}

	if err := s.pwdService.UpdatePassword(ctx, &updatePwdDTO); err != nil {
		if errors.Is(err, ErrUpdatedRecordNotFound) {
			return nil, status.Errorf(codes.NotFound, "Record not found for update")
		}
		s.zapLogger.Errorf("Error in UpdatePassword: %v", err)
		return nil, status.Errorf(codes.Internal, "Failed to update password: %v", err)
	}

	return &pb.BasicResponse{}, nil
}

// DeletePassword удаляет пароль.
func (s *PasswordGRPCServer) DeletePassword(ctx context.Context, req *pb.DeletePwdRequest) (*pb.BasicResponse, error) {
	userID, err := helpers.ParseTokenAndExtractUserID(req.GetJwt().GetValue(), s.secretKey)
	if err != nil {
		return nil, fmt.Errorf(ErrParseJWTFailed, err)
	}

	deletePwdDTO := pwddto.DeletePwdDTO{
		PwdID:  req.GetPwdId().GetValue(),
		UserID: userID,
	}

	if err := s.pwdService.DeletePassword(ctx, &deletePwdDTO); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, status.Errorf(codes.NotFound, "Password not found")
		}
		s.zapLogger.Errorf("Error in DeletePassword: %v", err)
		return nil, status.Errorf(codes.Internal, "Failed to delete password: %v", err)
	}

	return &pb.BasicResponse{}, nil
}

// GetPassword получает один пароль.
func (s *PasswordGRPCServer) GetPassword(ctx context.Context, req *pb.GetPwdRequest) (*pb.PwdResponse, error) {
	userID, err := helpers.ParseTokenAndExtractUserID(req.GetJwt().GetValue(), s.secretKey)
	if err != nil {
		return nil, fmt.Errorf(ErrParseJWTFailed, err)
	}

	getPwdDTO := pwddto.GetPwdDTO{
		PwdID:  req.GetPwdId().GetValue(),
		UserID: userID,
	}

	responsePwdDTO, err := s.pwdService.GetPassword(ctx, &getPwdDTO)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, status.Errorf(codes.NotFound, "Password not found")
		}
		s.zapLogger.Errorf("Error in GetPassword: %v", err)
		return nil, status.Errorf(codes.Internal, "Failed to get password: %v", err)
	}

	return &pb.PwdResponse{
		PwdId: wrapperspb.String(responsePwdDTO.PwdID),
		Title: wrapperspb.String(responsePwdDTO.Title),
		Credentials: &pb.Credentials{
			Login:    wrapperspb.String(responsePwdDTO.Credentials.Login),
			Password: wrapperspb.String(responsePwdDTO.Credentials.Password),
		},
	}, nil
}

// GetAllPasswords получает все пароли.
func (s *PasswordGRPCServer) GetAllPasswords(ctx context.Context, req *pb.AllPwdRequest) (*pb.AllPwdResponse, error) {
	userID, err := helpers.ParseTokenAndExtractUserID(req.GetJwt().GetValue(), s.secretKey)
	if err != nil {
		return nil, fmt.Errorf(ErrParseJWTFailed, err)
	}

	allPwdDTO := pwddto.AllPwdDTO{
		UserID: userID,
	}

	allPasswords, err := s.pwdService.GetAllPasswords(ctx, &allPwdDTO)
	if err != nil {
		s.zapLogger.Errorf("Error in GetAllPasswords: %v", err)
		return nil, status.Errorf(codes.Internal, "Failed to get all passwords: %v", err)
	}

	pwdResponses := make([]*pb.PwdResponse, 0, len(allPasswords))
	for _, pwd := range allPasswords {
		pwdResponses = append(pwdResponses, &pb.PwdResponse{
			PwdId: wrapperspb.String(pwd.PwdID),
			Title: wrapperspb.String(pwd.Title),
			Credentials: &pb.Credentials{
				Login:    wrapperspb.String(pwd.Credentials.Login),
				Password: wrapperspb.String(pwd.Credentials.Password),
			},
		})
	}

	return &pb.AllPwdResponse{Passwords: pwdResponses}, nil
}
