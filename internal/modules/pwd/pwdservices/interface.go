package pwdservices

import (
	"context"
	"gophKeeper/internal/modules/pwd/pwddto"
)

type IPwdService interface {
	SavePassword(ctx context.Context, dto *pwddto.SavePwdDTO) error
	DeletePassword(ctx context.Context, dto *pwddto.DeletePwdDTO) error
	GetAllPasswords(ctx context.Context, dto *pwddto.AllPwdDTO) ([]pwddto.PwdDTO, error)
	GetPassword(ctx context.Context, dto *pwddto.GetPwdDTO) (pwddto.CredentialsDTO, error)
	UpdatePassword(ctx context.Context, dto *pwddto.UpdatePwdDTO) error
}
