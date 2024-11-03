package fileservices

import (
	"context"
	"goph-keeper/internal/modules/file/filedto"
)

type FileInfo struct {
	FileID      string `json:"file_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type IFileService interface {
	SaveFile(ctx context.Context, dto filedto.SaveFileDTO) error
	DeleteFile(ctx context.Context, dto filedto.DeleteFileDTO) error
	GetFile(context.Context, filedto.GetFileDTO) (string, error)
	GetAllFiles(ctx context.Context, dto filedto.AllFilesDTO) ([]FileInfo, error)
}
