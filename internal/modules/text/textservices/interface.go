package textservices

import (
	"context"
	"gophKeeper/internal/modules/text/textdto"
)

type ITextService interface {
	SaveText(ctx context.Context, dto textdto.SaveTextDTO) error
	DeleteText(ctx context.Context, dto textdto.DeleteTextDTO) error
	GetAllTexts(ctx context.Context, dto textdto.AllTextDTO) ([]textdto.TextDTO, error)
	GetText(ctx context.Context, dto textdto.GetTextDTO) (textdto.TextDTO, error)
	UpdateText(ctx context.Context, dto textdto.UpdateTextDTO) error
}
