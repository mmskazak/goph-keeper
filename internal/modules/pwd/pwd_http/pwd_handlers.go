package pwd_http

import (
	"encoding/json"
	"errors"
	"gophKeeper/internal/modules/pwd/pwd_services"
	"gophKeeper/internal/modules/pwd/pwd_services/dto/request"
	"net/http"
)

var someSpecificError = errors.New("updated record not found")

type PwdHandlers struct {
	pwdService pwd_services.IPwdService
}

func NewPwdHandlersHTTP(service pwd_services.IPwdService) PwdHandlers {
	return PwdHandlers{
		pwdService: service,
	}
}

func (p PwdHandlers) SavePassword(w http.ResponseWriter, r *http.Request) {
	savePwdDTO, err := request.SavePwdDTOFromHTTP(r)
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
	getPwdDTO, err := request.GetPwdDTOFromHTTP(r)
	if err != nil {
		http.Error(w, "", http.StatusBadRequest)
		return
	}
	credentials, err := p.pwdService.GetPassword(r.Context(), getPwdDTO)
	if err != nil {
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	jsonResponse, err := json.Marshal(credentials)
	if err != nil {
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	// Устанавливаем статус и отправляем ответ
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

func (p PwdHandlers) DeletePassword(w http.ResponseWriter, r *http.Request) {
	deletePwdDTO, err := request.DeletePwdDTOFromHTTP(r)
	if err != nil {
		http.Error(w, "", http.StatusBadRequest)
		return
	}
	err = p.pwdService.DeletePassword(r.Context(), deletePwdDTO)
	if err != nil {
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func (p PwdHandlers) GetAllPasswords(w http.ResponseWriter, r *http.Request) {
	allPwdDTO, err := request.AllPwdDTOFromHTTP(r)
	if err != nil {
		http.Error(w, "", http.StatusBadRequest)
		return
	}
	allPasswords, err := p.pwdService.GetAllPasswords(r.Context(), allPwdDTO)
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

func (p PwdHandlers) UpdatePassword(w http.ResponseWriter, r *http.Request) {
	updatePwdDTO, err := request.UpdatePwdDTOFromHTTP(r)
	if err != nil {
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	err = p.pwdService.UpdatePassword(r.Context(), updatePwdDTO)
	if err != nil {
		if errors.Is(err, someSpecificError) {
			http.Error(w, "", http.StatusNotFound) // Запись для обновления не найдена
		} else {
			http.Error(w, "", http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}
