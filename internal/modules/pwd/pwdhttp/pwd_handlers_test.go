package pwdhttp

import (
	"context"
	"github.com/stretchr/testify/assert"
	request2 "gophKeeper/internal/modules/pwd/pwddto/request"
	response2 "gophKeeper/internal/modules/pwd/pwddto/response"
	"net/http"
	"net/http/httptest"
	"testing"
)

type FakePwdService struct{}

func (f FakePwdService) SavePassword(ctx context.Context, dto request2.SavePwdDTO) error {
	return nil
}

func (f FakePwdService) DeletePassword(ctx context.Context, dto request2.DeletePwdDTO) error {
	return nil
}
func (f FakePwdService) GetPassword(ctx context.Context, dto request2.GetPwdDTO) (response2.CredentialsDTO, error) {
	return response2.CredentialsDTO{}, nil
}

func (f FakePwdService) GetAllPasswords(ctx context.Context, dto request2.AllPwdDTO) ([]response2.PwdDTO, error) {
	return []response2.PwdDTO{}, nil
}

func TestPwdHandlers_SavePassword(t *testing.T) {
	fake := FakePwdService{}
	authHandlers := NewPwdHandlersHTTP(fake)
	r := httptest.NewRequest(http.MethodPost, "/save-password", http.NoBody)
	w := httptest.NewRecorder()
	authHandlers.SavePassword(w, r)
	assert.Equal(t, http.StatusBadRequest, w.Code)
}
