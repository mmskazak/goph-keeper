package text_services

import (
	"context"
	"gophKeeper/internal/modules/text/text_dto"
)

type ITextService interface {
	SaveText(ctx context.Context, dto text_dto.SaveTextDTO) error
	DeleteText(ctx context.Context, dto text_dto.DeleteTextDTO) error
	GetAllTexts(ctx context.Context, dto text_dto.AllTextDTO) ([]text_dto.TextDTO, error)
	GetText(ctx context.Context, dto text_dto.GetTextDTO) (text_dto.TextDTO, error)
	UpdateText(ctx context.Context, dto text_dto.UpdateTextDTO) error
}
