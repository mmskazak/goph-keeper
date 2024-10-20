package pwd_http

import (
	"context"
	"github.com/stretchr/testify/assert"
	"gophKeeper/internal/modules/pwd/pwd_services/dto/request"
	"gophKeeper/internal/modules/pwd/pwd_services/dto/response"
	"net/http"
	"net/http/httptest"
	"testing"
)

type FakePwdService struct{}

func (f FakePwdService) SavePassword(ctx context.Context, dto request.SavePwdDTO) error {
	return nil
}

func (f FakePwdService) DeletePassword(ctx context.Context, dto request.DeletePwdDTO) error {
	return nil
}
func (f FakePwdService) GetPassword(ctx context.Context, dto request.GetPwdDTO) (response.CredentialsDTO, error) {
	return response.CredentialsDTO{}, nil
}

func (f FakePwdService) GetAllPasswords(ctx context.Context, dto request.AllPwdDTO) ([]response.PwdDTO, error) {
	return []response.PwdDTO{}, nil
}

func TestPwdHandlers_SavePassword(t *testing.T) {
	fake := FakePwdService{}
	authHandlers := NewPwdHandlersHTTP(fake)
	r := httptest.NewRequest(http.MethodPost, "/save-password", http.NoBody)
	w := httptest.NewRecorder()
	authHandlers.SavePassword(w, r)
	assert.Equal(t, http.StatusBadRequest, w.Code)
}
