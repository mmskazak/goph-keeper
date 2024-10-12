package pwd_http

import (
	"gophKeeper/internal/modules/pwd/pwd_services"
	"net/http"
)

type PwdHandlers struct {
	pwdService *pwd_services.PwdService
}

func NewAuthHandlersHTTP(pwdService *pwd_services.PwdService) PwdHandlers {
	return PwdHandlers{
		pwdService: pwdService,
	}
}

func (h PwdHandlers) ServeHTTP(w http.ResponseWriter, r *http.Request) {}
