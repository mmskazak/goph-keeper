package pwd_services

import (
	"context"
	"gophKeeper/internal/modules/pwd/pwd_services/dto/request"
	"gophKeeper/internal/modules/pwd/pwd_services/dto/response"
)

type IPwdService interface {
	SavePassword(ctx context.Context, dto request.SavePwdDTO) error
	DeletePassword(ctx context.Context, dto request.DeletePwdDTO) error
	GetAllPasswords(ctx context.Context, dto request.AllPwdDTO) ([]response.PwdDTO, error)
	GetPassword(ctx context.Context, dto request.GetPwdDTO) (response.CredentialsDTO, error)
}
