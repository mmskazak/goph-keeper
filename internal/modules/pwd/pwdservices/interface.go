package pwdservices

import (
	"context"
	"goph-keeper/internal/modules/pwd/pwddto"
)

type IPwdService interface {
	SavePassword(ctx context.Context, dto *pwddto.SavePwdDTO) error
	DeletePassword(ctx context.Context, dto *pwddto.DeletePwdDTO) error
	GetAllPasswords(ctx context.Context, dto *pwddto.AllPwdDTO) ([]pwddto.ResponsePwdDTO, error)
	GetPassword(ctx context.Context, dto *pwddto.GetPwdDTO) (pwddto.ResponsePwdDTO, error)
	UpdatePassword(ctx context.Context, dto *pwddto.UpdatePwdDTO) error
}
