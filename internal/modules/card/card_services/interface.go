package card_services

import (
	"context"
	"gophKeeper/internal/modules/text/text_dto"
)

type InfoByText struct {
	Resource string `json:"resource"`
	Login    string `json:"login"`
	Password string `json:"password"`
}

type ICardService interface {
	SaveCard(ctx context.Context, dto text_dto.SaveTextDTO) error
	DeleteCard(ctx context.Context, dto text_dto.DeleteTextDTO) error
	GetCard(ctx context.Context, dto text_dto.GetTextDTO) (string, error)
	GetAllCards(ctx context.Context, dto text_dto.AllTextDTO) ([]InfoByText, error)
}
