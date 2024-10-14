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
	DeletePassword(ctx context.Context, dto pwd_dto.DeletePwdDTO) error
	GetPassword(ctx context.Context, dto pwd_dto.GetPwdDTO) (string, error)
	GetAllPasswords(ctx context.Context, dto pwd_dto.AllPwdDTO) ([]InfoByPassword, error)
}
