package file_services

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"gophKeeper/internal/modules/file/file_dto/request"
)

type FileService struct {
	pool *pgxpool.Pool
}

func NewFileService(pool *pgxpool.Pool) *FileService {
	return &FileService{pool: pool}
}

func (pwd *FileService) SaveFile(ctx context.Context, dto request.SaveFileDTO) error {
	return nil
}

func (pwd *FileService) DeleteFile(ctx context.Context, dto request.DeleteFileDTO) error {
	return nil
}

func (pwd *FileService) GetPassword(ctx context.Context, dto request.GetFileDTO) (string, error) {
	return "", nil
}

func (pwd *FileService) GetAllPasswords(ctx context.Context, dto request.AllFilesDTO) ([]InfoByFiles, error) {
	return []InfoByFiles{}, nil
}
