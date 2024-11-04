package filegrpc

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"goph-keeper/internal/helpers"
	"goph-keeper/internal/modules/file/filedto"
	"goph-keeper/internal/modules/file/fileservices"
	"goph-keeper/internal/modules/file/proto"
)

//go:generate protoc --proto_path=../proto --go_out=. --go-grpc_out=. file.proto

const ErrParsingValidateJWT = "error parsing and validating JWT token: %w"

// FileGRPCServer - сервис GRPC отвечающий за работу с файлами.
type FileGRPCServer struct {
	proto.UnimplementedFileServiceServer

	fileService fileservices.IFileService
	secretKey   string
}

// NewFileGRPCServer - создаёт новый FileGRPCServer.
func NewFileGRPCServer(fileService fileservices.IFileService, secretKey string) *FileGRPCServer {
	return &FileGRPCServer{
		fileService: fileService,
		secretKey:   secretKey,
	}
}

// SaveFile сохраняет файл на сервере.
func (s *FileGRPCServer) SaveFile(ctx context.Context, req *proto.SaveFileRequest) (*proto.BasicResponse, error) {
	userID, err := helpers.ParseTokenAndExtractUserID(req.GetJwt(), s.secretKey)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "Failed to validate JWT token: %v", err)
	}

	saveFileDTO := filedto.SaveFileDTO{
		UserID:   userID,
		NameFile: req.GetNameFile(),
		FileData: req.GetFileData(),
	}

	if err := s.fileService.SaveFile(ctx, saveFileDTO); err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to save file: %v", err)
	}

	// Успешный ответ
	return &proto.BasicResponse{}, nil
}

// GetFile возвращает файл пользователю.
func (s *FileGRPCServer) GetFile(req *proto.GetFileRequest, stream proto.FileService_GetFileServer) error {
	userID, err := helpers.ParseTokenAndExtractUserID(req.GetJwt(), s.secretKey)
	if err != nil {
		return status.Errorf(codes.Unauthenticated, "Failed to validate JWT token: %v", err)
	}

	getFileDTO := filedto.GetFileDTO{
		UserID: userID,
		FileID: req.GetFileId(),
	}

	// Получаем байты файла из сервиса
	fileData, err := s.fileService.GetFile(stream.Context(), getFileDTO)
	if err != nil {
		return status.Errorf(codes.Internal, "Failed to get file: %v", err)
	}

	// Отправляем данные файла в поток
	if err := stream.Send(&proto.GetFileResponse{
		FileData: fileData,
	}); err != nil {
		return status.Errorf(codes.Internal, "Failed to send file data: %v", err)
	}

	return nil
}

// DeleteFile удаляет файл с сервера.
func (s *FileGRPCServer) DeleteFile(ctx context.Context, req *proto.DeleteFileRequest) (*proto.BasicResponse, error) {
	userID, err := helpers.ParseTokenAndExtractUserID(req.GetJwt(), s.secretKey)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "Failed to validate JWT token: %v", err)
	}

	deleteFileDTO := filedto.DeleteFileDTO{
		UserID: userID,
		FileID: req.GetFileId(),
	}

	if err := s.fileService.DeleteFile(ctx, deleteFileDTO); err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to delete file: %v", err)
	}

	return &proto.BasicResponse{}, nil
}

// GetAllFiles возвращает список всех файлов пользователя.
func (s *FileGRPCServer) GetAllFiles(
	ctx context.Context,
	req *proto.GetAllFilesRequest,
) (*proto.GetAllFilesResponse, error) {
	userID, err := helpers.ParseTokenAndExtractUserID(req.GetJwt(), s.secretKey)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "Failed to validate JWT token: %v", err)
	}

	getAllFilesDTO := filedto.AllFilesDTO{
		UserID: userID,
	}

	files, err := s.fileService.GetAllFiles(ctx, getAllFilesDTO)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to get all files: %v", err)
	}

	var fileItems []*proto.FileInfo
	for _, file := range files {
		fileItems = append(fileItems, &proto.FileInfo{
			FileId:   file.FileID,
			NameFile: file.NameFile, // Используем название из DTO
		})
	}

	return &proto.GetAllFilesResponse{
		Files: fileItems,
	}, nil
}
