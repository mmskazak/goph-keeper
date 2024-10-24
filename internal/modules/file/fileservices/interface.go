package fileservices

import (
	"context"
	"gophKeeper/internal/modules/file/filedto/request"
)

type FileInfo struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type IFileService interface {
	SaveFile(ctx context.Context, dto request.SaveFileDTO) error
	DeleteFile(ctx context.Context, dto request.DeleteFileDTO) error
	GetFile(context.Context, request.GetFileDTO) (string, error)
	GetAllFiles(ctx context.Context, dto request.AllFilesDTO) ([]FileInfo, error)
}
