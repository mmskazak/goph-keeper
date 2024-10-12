package pwd_http

import (
	"gophKeeper/internal/modules/pwd/pwd_dto"
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

func (p PwdHandlers) SavePassword(w http.ResponseWriter, r *http.Request) {
	savePwdDTO, err := pwd_dto.GetSavePwdDTOFromHTTP(r)
	if err != nil {
		http.Error(w, "", http.StatusBadRequest)
	}
	err = p.pwdService.SavePassword(savePwdDTO.Login, savePwdDTO.Password)
	if err != nil {
		http.Error(w, "", http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("OK"))
}

func (p PwdHandlers) GetPassword(w http.ResponseWriter, r *http.Request) {

}

func (p PwdHandlers) DeletePassword(w http.ResponseWriter, r *http.Request) {

}

func (p PwdHandlers) GetAllPasswords(w http.ResponseWriter, r *http.Request) {

}
