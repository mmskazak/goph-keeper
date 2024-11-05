package routespwd

import (
	"go.uber.org/zap"
	"goph-keeper/internal/modules/pwd/pwdhttp"
	"goph-keeper/internal/modules/pwd/pwdservices"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func RegistrationRoutesPwd(
	r chi.Router,
	pwdService *pwdservices.PwdService,
	zapLogger *zap.SugaredLogger,
) {
	// Сохранить пароль
	r.Post("/pwd/save", func(w http.ResponseWriter, req *http.Request) {
		getPwdHandlers(pwdService, zapLogger).SavePassword(w, req)
	})
	// Получить конкретный пароль
	r.Post("/pwd/update", func(w http.ResponseWriter, req *http.Request) {
		getPwdHandlers(pwdService, zapLogger).UpdatePassword(w, req)
	})
	// Получить все пароли
	r.Get("/pwd/all", func(w http.ResponseWriter, req *http.Request) {
		getPwdHandlers(pwdService, zapLogger).GetAllPasswords(w, req)
	})
	// Удалить пароль
	r.Get("/pwd/delete/{pwd_id}", func(w http.ResponseWriter, req *http.Request) {
		getPwdHandlers(pwdService, zapLogger).DeletePassword(w, req)
	})
	// Получить конкретный пароль
	r.Get("/pwd/get/{pwd_id}", func(w http.ResponseWriter, req *http.Request) {
		getPwdHandlers(pwdService, zapLogger).GetPassword(w, req)
	})
}

func getPwdHandlers(pwdService *pwdservices.PwdService, zapLogger *zap.SugaredLogger) *pwdhttp.PwdHandlers {
	pwdHandlers := pwdhttp.NewPwdHandlersHTTP(pwdService, zapLogger)
	return &pwdHandlers
}
