package pwd_services

import (
	"context"
	"gophKeeper/internal/modules/pwd/pwd_dto"
)

type InfoByPassword struct {
	Resource string `json:"resource"`
	Login    string `json:"login"`
	Password string `json:"password"`
}

type IPwdService interface {
	SavePassword(ctx context.Context, dto pwd_dto.SavePwdDTO) error
	DeletePassword(username string) error
	GetPassword(username string) (string, error)
	GetAllPasswords(username string) ([]InfoByPassword, error)
}
