package file_services

import (
	"context"
	"gophKeeper/internal/modules/file/file_dto/request"
)

type InfoByText struct {
	Resource string `json:"resource"`
	Login    string `json:"login"`
	Password string `json:"password"`
}

type IFileService interface {
	SaveText(ctx context.Context, dto request.SaveFileDTO) error
	DeleteText(ctx context.Context, dto request.DeleteFileDTO) error
	GetText(ctx context.Context, dto request.GetFileDTO) (string, error)
	GetAllTexts(ctx context.Context, dto request.AllFilesDTO) ([]InfoByText, error)
}
