package file_services

import (
	"context"
	"gophKeeper/internal/modules/file/file_dto/request"
)

type FileInfo struct {
	Name       string `json:"name"`
	PathToFile string `json:"path_to_file"`
}

type IFileService interface {
	SaveFile(ctx context.Context, dto request.SaveFileDTO) error
	DeleteFile(ctx context.Context, dto request.DeleteFileDTO) error
	GetFile(ctx context.Context, dto request.GetFileDTO) (string, error)
	GetAllFiles(ctx context.Context, dto request.AllFilesDTO) ([]FileInfo, error)
}
