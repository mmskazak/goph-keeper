package pwd_http

import (
	"context"
	"github.com/stretchr/testify/assert"
	"gophKeeper/internal/modules/pwd/pwd_dto"
	"gophKeeper/internal/modules/pwd/pwd_services"
	"net/http"
	"net/http/httptest"
	"testing"
)

type FakePwdService struct{}

func (f FakePwdService) SavePassword(ctx context.Context, dto pwd_dto.SavePwdDTO) error {
	return nil
}

func (f FakePwdService) DeletePassword(username string) error {
	return nil
}
func (f FakePwdService) GetPassword(username string) (string, error) {
	return "secret", nil
}
func (f FakePwdService) GetAllPasswords(username string) (pwd_services.AllPasswords, error) {
	return pwd_services.AllPasswords{}, nil
}

func TestPwdHandlers_SavePassword(t *testing.T) {
	fake := FakePwdService{}
	authHandlers := NewPwdHandlersHTTP(fake)
	r := httptest.NewRequest(http.MethodPost, "/save-password", http.NoBody)
	w := httptest.NewRecorder()
	authHandlers.SavePassword(w, r)
	assert.Equal(t, http.StatusBadRequest, w.Code)
}
