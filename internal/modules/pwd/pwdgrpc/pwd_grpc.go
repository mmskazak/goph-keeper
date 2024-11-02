package pwdgrpc

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v5"
	"goph-keeper/internal/logger"
	pb "goph-keeper/internal/modules/pwd/proto"
	"goph-keeper/internal/modules/pwd/pwddto"
	"goph-keeper/internal/modules/pwd/pwdservices"
	"goph-keeper/internal/modules/pwd/valueobj"
)

//go:generate protoc --proto_path=../proto --go_out=. --go-grpc_out=. pwd.proto

var ErrUpdatedRecordNotFound = errors.New("updated record not found")

// PasswordServiceServer Структура, которая реализует gRPC сервис.
type PasswordServiceServer struct {
	pb.UnimplementedPasswordServiceServer
	pwdService pwdservices.IPwdService
}

// NewPasswordServiceServer Функция для создания нового gRPC сервера.
func NewPasswordServiceServer(service pwdservices.IPwdService) *PasswordServiceServer {
	return &PasswordServiceServer{
		pwdService: service,
	}
}

// SavePassword Сохранение пароля.
func (s *PasswordServiceServer) SavePassword(ctx context.Context, req *pb.SavePwdRequest) (*pb.BasicResponse, error) {
	savePwdDTO := pwddto.SavePwdDTO{
		Title:       req.Title,
		Description: req.Description,
		Credentials: valueobj.Credentials{
			Login:    req.Credentials.Login,
			Password: req.Credentials.Password,
		},
	}
	err := s.pwdService.SavePassword(ctx, &savePwdDTO)
	if err != nil {
		logger.Log.Errorf("Error in SavePassword: %v", err)
		return nil, err
	}
	return &pb.BasicResponse{
		Status:  "success",
		Message: "Password saved successfully",
	}, nil
}

// UpdatePassword Обновление пароля.
func (s *PasswordServiceServer) UpdatePassword(ctx context.Context, req *pb.UpdatePwdRequest) (*pb.BasicResponse, error) {
	updatePwdDTO := pwddto.UpdatePwdDTO{
		PwdID:       req.PwdId,
		Title:       req.Title,
		Description: req.Description,
		Credentials: valueobj.Credentials{
			Login:    req.Credentials.Login,
			Password: req.Credentials.Password,
		},
	}
	err := s.pwdService.UpdatePassword(ctx, &updatePwdDTO)
	if err != nil {
		if errors.Is(err, ErrUpdatedRecordNotFound) {
			return &pb.BasicResponse{
				Status:  "error",
				Message: "Record not found for update",
			}, nil
		}
		logger.Log.Errorf("Error in UpdatePassword: %v", err)
		return nil, err
	}
	return &pb.BasicResponse{
		Status:  "success",
		Message: "Password updated successfully",
	}, nil
}

// DeletePassword Удаление пароля.
func (s *PasswordServiceServer) DeletePassword(ctx context.Context, req *pb.DeletePwdRequest) (*pb.BasicResponse, error) {
	deletePwdDTO := pwddto.DeletePwdDTO{PwdID: req.PwdId}
	err := s.pwdService.DeletePassword(ctx, &deletePwdDTO)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return &pb.BasicResponse{
				Status:  "error",
				Message: "Password not found",
			}, nil
		}
		logger.Log.Errorf("Error in DeletePassword: %v", err)
		return nil, err
	}
	return &pb.BasicResponse{
		Status:  "success",
		Message: "Password deleted successfully",
	}, nil
}

// GetPassword Получение одного пароля.
func (s *PasswordServiceServer) GetPassword(ctx context.Context, req *pb.GetPwdRequest) (*pb.PwdResponse, error) {
	getPwdDTO := pwddto.GetPwdDTO{PwdID: req.PwdId}
	responsePwdDTO, err := s.pwdService.GetPassword(ctx, &getPwdDTO)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			logger.Log.Infoln("Password not found")
			return nil, nil
		}
		logger.Log.Errorf("Error in GetPassword: %v", err)
		return nil, err
	}

	return &pb.PwdResponse{
		PwdId:       responsePwdDTO.PwdID,
		Title:       responsePwdDTO.Title,
		Description: responsePwdDTO.Description,
		Credentials: &pb.Credentials{
			Login:    responsePwdDTO.Credentials.Login,
			Password: responsePwdDTO.Credentials.Password,
		},
	}, nil
}

// GetAllPasswords Получение всех паролей.
func (s *PasswordServiceServer) GetAllPasswords(ctx context.Context, req *pb.AllPwdRequest) (*pb.AllPwdResponse, error) {
	allPwdDTO := pwddto.AllPwdDTO{}
	allPasswords, err := s.pwdService.GetAllPasswords(ctx, &allPwdDTO)
	if err != nil {
		logger.Log.Errorf("Error in GetAllPasswords: %v", err)
		return nil, err
	}

	pwdResponses := make([]*pb.PwdResponse, 0, len(allPasswords))
	for _, pwd := range allPasswords {
		pwdResponses = append(pwdResponses, &pb.PwdResponse{
			PwdId:       pwd.PwdID,
			Title:       pwd.Title,
			Description: pwd.Description,
			Credentials: &pb.Credentials{
				Login:    pwd.Credentials.Login,
				Password: pwd.Credentials.Password,
			},
		})
	}
	return &pb.AllPwdResponse{Passwords: pwdResponses}, nil
}
