package filegrpc

import (
	"context"
	"fmt"
	"goph-keeper/internal/helpers"
	"goph-keeper/internal/modules/file/filedto"
	"goph-keeper/internal/modules/file/fileservices"
	"goph-keeper/internal/modules/file/proto"
)

//go:generate protoc --proto_path=../proto --go_out=. --go-grpc_out=. file.proto

const ErrParsingValidateJWT = "error parsing and validating JWT token: %w"

// FileGRPCServer сервис GRPC отвечающий за работу с файлами.
type FileGRPCServer struct {
	proto.UnimplementedFileServiceServer

	fileService fileservices.IFileService
	secretKey   string
}

// NewFileGRPCServer ...
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
		return nil, fmt.Errorf(ErrParsingValidateJWT, err)
	}

	saveFileDTO := filedto.SaveFileDTO{
		UserID:      userID,
		Title:       req.GetTitle(),
		Description: req.GetDescription(),
		FileData:    req.GetFileData(),
	}

	if err := s.fileService.SaveFile(ctx, saveFileDTO); err != nil {
		return nil, fmt.Errorf("failed to save file: %w", err)
	}

	return &proto.BasicResponse{Status: "success", Message: "File saved successfully"}, nil
}

// GetFile возвращает файл пользователю.
func (s *FileGRPCServer) GetFile(ctx context.Context, req *proto.GetFileRequest) (*proto.GetFileResponse, error) {
	userID, err := helpers.ParseTokenAndExtractUserID(req.Jwt, s.secretKey)
	if err != nil {
		return nil, fmt.Errorf(ErrParsingValidateJWT, err)
	}

	getFileDTO := filedto.GetFileDTO{
		UserID: userID,
		FileID: req.GetFileId(),
	}

	// Получаем байты файла из сервиса
	fileData, err := s.fileService.GetFile(ctx, getFileDTO)
	if err != nil {
		return nil, fmt.Errorf("failed to get file: %w", err)
	}

	return &proto.GetFileResponse{
		FileId:   req.GetFileId(),
		FileData: fileData,
	}, nil
}

// DeleteFile удаляет файл с сервера.
func (s *FileGRPCServer) DeleteFile(ctx context.Context, req *proto.DeleteFileRequest) (*proto.BasicResponse, error) {
	userID, err := helpers.ParseTokenAndExtractUserID(req.GetJwt(), s.secretKey)
	if err != nil {
		return nil, fmt.Errorf(ErrParsingValidateJWT, err)
	}

	deleteFileDTO := filedto.DeleteFileDTO{
		UserID: userID,
		FileID: req.FileId,
	}

	if err := s.fileService.DeleteFile(ctx, deleteFileDTO); err != nil {
		return nil, fmt.Errorf("failed to delete file: %w", err)
	}

	return &proto.BasicResponse{Status: "success", Message: "File deleted successfully"}, nil
}

// GetAllFiles возвращает список всех файлов пользователя.
func (s *FileGRPCServer) GetAllFiles(ctx context.Context, req *proto.GetAllFilesRequest) (*proto.GetAllFilesResponse, error) {
	userID, err := helpers.ParseTokenAndExtractUserID(req.GetJwt(), s.secretKey)
	if err != nil {
		return nil, fmt.Errorf(ErrParsingValidateJWT, err)
	}

	getAllFilesDTO := filedto.AllFilesDTO{
		UserID: userID,
	}

	files, err := s.fileService.GetAllFiles(ctx, getAllFilesDTO)
	if err != nil {
		return nil, fmt.Errorf("failed to get all files: %w", err)
	}

	var fileItems []*proto.FileInfo
	for _, file := range files {
		fileItems = append(fileItems, &proto.FileInfo{
			FileId:      file.FileID,
			Title:       file.Title,
			Description: file.Description,
		})
	}

	return &proto.GetAllFilesResponse{
		Files: fileItems,
	}, nil
}
