package fileservices

import (
	"context"
	"goph-keeper/internal/modules/file/filedto"
)

type FileInfo struct {
	NameFile string `json:"name_file"`
	FileID   string `json:"file_id"`
}

type IFileService interface {
	SaveFile(ctx context.Context, dto filedto.SaveFileDTO) error
	DeleteFile(ctx context.Context, dto filedto.DeleteFileDTO) error
	GetFile(ctx context.Context, dto filedto.GetFileDTO) ([]byte, string, error)
	GetAllFiles(ctx context.Context, dto filedto.AllFilesDTO) ([]FileInfo, error)
}
