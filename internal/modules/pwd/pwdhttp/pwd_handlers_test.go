package pwdhttp

import (
	"context"
	"github.com/stretchr/testify/assert"
	"goph-keeper/internal/modules/pwd/pwddto"
	"net/http"
	"net/http/httptest"
	"testing"
)

type FakePwdService struct{}

func (f FakePwdService) SavePassword(ctx context.Context, dto *pwddto.SavePwdDTO) error {
	return nil
}

func (f FakePwdService) DeletePassword(ctx context.Context, dto *pwddto.DeletePwdDTO) error {
	return nil
}
func (f FakePwdService) GetPassword(ctx context.Context, dto *pwddto.GetPwdDTO) (pwddto.CredentialsDTO, error) {
	return pwddto.CredentialsDTO{}, nil
}

func (f FakePwdService) GetAllPasswords(ctx context.Context, dto *pwddto.AllPwdDTO) ([]pwddto.PwdDTO, error) {
	return []pwddto.PwdDTO{}, nil
}
func (f FakePwdService) UpdatePassword(ctx context.Context, dto *pwddto.UpdatePwdDTO) error {
	return nil
}

func TestPwdHandlers_SavePassword(t *testing.T) {
	fake := FakePwdService{}
	authHandlers := NewPwdHandlersHTTP(fake)
	r := httptest.NewRequest(http.MethodPost, "/save-password", http.NoBody)
	w := httptest.NewRecorder()
	authHandlers.SavePassword(w, r)
	assert.Equal(t, http.StatusBadRequest, w.Code)
}
