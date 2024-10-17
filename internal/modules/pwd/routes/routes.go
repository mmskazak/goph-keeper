package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"gophKeeper/internal/modules/pwd/pwd_http"
	"gophKeeper/internal/modules/pwd/pwd_services"
	"net/http"
)

func RegistrationRoutesPwd(r chi.Router, pool *pgxpool.Pool) chi.Router {
	//Сохранить пароль
	r.Post("/pwd/save", func(w http.ResponseWriter, req *http.Request) {
		getPwdHandlers(pool).SavePassword(w, req)
	})
	//Получить все пароли
	r.Post("/pwd/all", func(w http.ResponseWriter, req *http.Request) {
		getPwdHandlers(pool).GetAllPasswords(w, req)
	})
	//Удалить пароль
	r.Delete("/pwd/delete", func(w http.ResponseWriter, req *http.Request) {
		getPwdHandlers(pool).DeletePassword(w, req)
	})
	//Получить конкретный пароль
	r.Post("/pwd/get", func(w http.ResponseWriter, req *http.Request) {
		getPwdHandlers(pool).GetPassword(w, req)
	})

	return r
}

func getPwdHandlers(pool *pgxpool.Pool) *pwd_http.PwdHandlers {
	pwdService := pwd_services.NewPwdService(pool)
	pwdHandlers := pwd_http.NewPwdHandlersHTTP(pwdService)
	return &pwdHandlers
}
