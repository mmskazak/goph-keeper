package cardservices

import (
	"context"
	"gophKeeper/internal/modules/card/carddto"
)

type ICardService interface {
	SaveCard(ctx context.Context, dto carddto.SaveCardDTO) error
	DeleteCard(ctx context.Context, dto carddto.DeleteCardDTO) error
	GetCard(ctx context.Context, dto carddto.GetCardDTO) (carddto.SaveCardDTO, error)
	GetAllCards(ctx context.Context, dto carddto.GetAllCardsDTO) ([]carddto.SaveCardDTO, error)
	UpdateCard(ctx context.Context, dto carddto.UpdateCardDTO) error
}
