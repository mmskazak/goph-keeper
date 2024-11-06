package pwdhttp

import (
	"context"
	"goph-keeper/internal/modules/pwd/pwddto"
	"net/http"
	"net/http/httptest"
	"testing"

	"go.uber.org/zap"

	"github.com/stretchr/testify/assert"
)

type FakePwdService struct{}

func (f FakePwdService) SavePassword(ctx context.Context, dto *pwddto.SavePwdDTO) error {
	return nil
}

func (f FakePwdService) DeletePassword(ctx context.Context, dto *pwddto.DeletePwdDTO) error {
	return nil
}
func (f FakePwdService) GetPassword(ctx context.Context, dto *pwddto.GetPwdDTO) (pwddto.ResponsePwdDTO, error) {
	return pwddto.ResponsePwdDTO{}, nil
}

func (f FakePwdService) GetAllPasswords(ctx context.Context, dto *pwddto.AllPwdDTO) ([]pwddto.ResponsePwdDTO, error) {
	return []pwddto.ResponsePwdDTO{}, nil
}
func (f FakePwdService) UpdatePassword(ctx context.Context, dto *pwddto.UpdatePwdDTO) error {
	return nil
}

func TestPwdHandlers_SavePassword(t *testing.T) {
	fake := FakePwdService{}
	authHandlers := NewPwdHandlersHTTP(fake, zap.NewNop().Sugar())
	r := httptest.NewRequest(http.MethodPost, "/pwd/save", http.NoBody)
	w := httptest.NewRecorder()
	authHandlers.SavePassword(w, r)
	assert.Equal(t, http.StatusBadRequest, w.Code)
}
