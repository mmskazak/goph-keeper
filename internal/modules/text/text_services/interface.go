package text_services

import (
	"context"
	"gophKeeper/internal/modules/text/text_dto"
)

type InfoByText struct {
	Resource string `json:"resource"`
	Login    string `json:"login"`
	Password string `json:"password"`
}

type ITextService interface {
	SaveText(ctx context.Context, dto text_dto.SaveTextDTO) error
	DeleteText(ctx context.Context, dto text_dto.DeleteTextDTO) error
	GetText(ctx context.Context, dto text_dto.GetTextDTO) (string, error)
	GetAllTexts(ctx context.Context, dto text_dto.AllTextDTO) ([]InfoByText, error)
}
