package pwd_http

import (
	"encoding/json"
	"gophKeeper/internal/modules/pwd/pwd_dto"
	"gophKeeper/internal/modules/pwd/pwd_services"
	"net/http"
)

type PwdHandlers struct {
	pwdService pwd_services.IPwdService
}

func NewPwdHandlersHTTP(service pwd_services.IPwdService) PwdHandlers {
	return PwdHandlers{
		pwdService: service,
	}
}

func (p PwdHandlers) SavePassword(w http.ResponseWriter, r *http.Request) {
	savePwdDTO, err := pwd_dto.SavePwdDTOFromHTTP(r)
	if err != nil {
		http.Error(w, "", http.StatusBadRequest)
		return
	}
	err = p.pwdService.SavePassword(r.Context(), savePwdDTO)
	if err != nil {
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("OK"))
}

func (p PwdHandlers) GetPassword(w http.ResponseWriter, r *http.Request) {
	getPwdDTO, err := pwd_dto.GetPwdDTOFromHTTP(r)
	if err != nil {
		http.Error(w, "", http.StatusBadRequest)
		return
	}
	password, err := p.pwdService.GetPassword(getPwdDTO.Login)
	if err != nil {
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(password))
}

func (p PwdHandlers) DeletePassword(w http.ResponseWriter, r *http.Request) {
	getPwdDTO, err := pwd_dto.DeletePwdDTOFromHTTP(r)
	if err != nil {
		http.Error(w, "", http.StatusBadRequest)
		return
	}
	err = p.pwdService.DeletePassword(getPwdDTO.Login)
	if err != nil {
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func (p PwdHandlers) GetAllPasswords(w http.ResponseWriter, r *http.Request) {
	allPwdDTO, err := pwd_dto.AllPwdDTOFromHTTP(r)
	if err != nil {
		http.Error(w, "", http.StatusBadRequest)
		return
	}
	allPasswords, err := p.pwdService.GetAllPasswords(allPwdDTO.Login)
	if err != nil {
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	marshaledAllPasswords, err := json.Marshal(allPasswords)
	if err != nil {
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(marshaledAllPasswords)
}
