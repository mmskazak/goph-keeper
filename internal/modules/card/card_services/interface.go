package card_services

import (
	"context"
	"gophKeeper/internal/modules/card/card_dto"
)

type ICardService interface {
	SaveCard(ctx context.Context, dto card_dto.SaveCardDTO) error
	DeleteCard(ctx context.Context, dto card_dto.DeleteCardDTO) error
	GetCard(ctx context.Context, dto card_dto.GetCardDTO) (card_dto.SaveCardDTO, error)
	GetAllCards(ctx context.Context, dto card_dto.GetAllCardsDTO) ([]card_dto.SaveCardDTO, error)
}
