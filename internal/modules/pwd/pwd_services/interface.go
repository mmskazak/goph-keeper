package pwd_services

import (
	"context"
	request2 "gophKeeper/internal/modules/pwd/pwd_dto/request"
	response2 "gophKeeper/internal/modules/pwd/pwd_dto/response"
)

type IPwdService interface {
	SavePassword(ctx context.Context, dto request2.SavePwdDTO) error
	DeletePassword(ctx context.Context, dto request2.DeletePwdDTO) error
	GetAllPasswords(ctx context.Context, dto request2.AllPwdDTO) ([]response2.PwdDTO, error)
	GetPassword(ctx context.Context, dto request2.GetPwdDTO) (response2.CredentialsDTO, error)
	UpdatePassword(ctx context.Context, dto request2.UpdatePwdDTO) error
}
