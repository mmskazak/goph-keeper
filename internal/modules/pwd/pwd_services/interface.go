package pwd_services

import (
	"context"
	"gophKeeper/internal/modules/pwd/pwd_services/dto/request"
	"gophKeeper/internal/modules/pwd/pwd_services/dto/response"
)

type InfoByPassword struct {
	Resource string `json:"resource"`
	Login    string `json:"login"`
	Password string `json:"password"`
}

type IPwdService interface {
	SavePassword(ctx context.Context, dto request.SavePwdDTO) error
	DeletePassword(ctx context.Context, dto request.DeletePwdDTO) error
	GetAllPasswords(ctx context.Context, dto request.AllPwdDTO) ([]InfoByPassword, error)
	GetPassword(ctx context.Context, dto request.GetPwdDTO) (response.PwdDTO, error)
}
