package pwdhttp

import (
	"encoding/json"
	"errors"
	"github.com/jackc/pgx/v5"
	"goph-keeper/internal/logger"
	"goph-keeper/internal/modules/pwd/pwddto"
	"goph-keeper/internal/modules/pwd/pwdservices"
	"net/http"
)

var ErrUpdatedRecordNotFound = errors.New("updated record not found")

type PwdHandlers struct {
	pwdService pwdservices.IPwdService
}

func NewPwdHandlersHTTP(service pwdservices.IPwdService) PwdHandlers {
	return PwdHandlers{
		pwdService: service,
	}
}

func (p PwdHandlers) SavePassword(w http.ResponseWriter, r *http.Request) {
	savePwdDTO, err := pwddto.SavePwdDTOFromHTTP(r)
	if err != nil {
		logger.Log.Errorf("Error make SavePwdDTOFromHTTP: %v", err)
		http.Error(w, "", http.StatusBadRequest)
		return
	}
	err = p.pwdService.SavePassword(r.Context(), &savePwdDTO)
	if err != nil {
		logger.Log.Errorf("Error work pwdService SavePwdDTO: %v", err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	_, err = w.Write([]byte("OK"))
	if err != nil {
		logger.Log.Errorf("Error writing response for handler SavePassword: %v", err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
}

func (p PwdHandlers) GetPassword(w http.ResponseWriter, r *http.Request) {
	getPwdDTO, err := pwddto.GetPwdDTOFromHTTP(r)
	if err != nil {
		logger.Log.Errorf("Error make GetPwdDTOFromHTTP: %v", err)
		http.Error(w, "", http.StatusBadRequest)
		return
	}
	credentials, err := p.pwdService.GetPassword(r.Context(), &getPwdDTO)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			logger.Log.Infoln("Get password not found")
			http.Error(w, "", http.StatusNotFound)
			return
		}
		logger.Log.Errorf("Error work pwdService GetPwdDTO: %v", err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	jsonResponse, err := json.Marshal(credentials)
	if err != nil {
		logger.Log.Errorf("Error work pwdService GetPwdDTO: %v", err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	// Устанавливаем статус и отправляем ответ
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(jsonResponse)
	if err != nil {
		logger.Log.Errorf("Error writing response for handler GetPwdDTO: %v", err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
}

func (p PwdHandlers) DeletePassword(w http.ResponseWriter, r *http.Request) {
	deletePwdDTO, err := pwddto.DeletePwdDTOFromHTTP(r)
	if err != nil {
		logger.Log.Errorf("Error make DeletePwdDTOFromHTTP: %v", err)
		http.Error(w, "", http.StatusBadRequest)
		return
	}
	err = p.pwdService.DeletePassword(r.Context(), &deletePwdDTO)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			logger.Log.Infoln("Deleted password not found")
			http.Error(w, "", http.StatusNotFound)
			return
		}
		logger.Log.Errorf("Error work pwdService DeletePwdDTO: %v", err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	_, err = w.Write([]byte("OK"))
	if err != nil {
		logger.Log.Errorf("Error writing response for handler DeletePwdDTO: %v", err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
}

func (p PwdHandlers) GetAllPasswords(w http.ResponseWriter, r *http.Request) {
	allPwdDTO, err := pwddto.AllPwdDTOFromHTTP(r)
	if err != nil {
		logger.Log.Errorf("Error make AllPwdDTOFromHTTP: %v", err)
		http.Error(w, "", http.StatusBadRequest)
		return
	}
	allPasswords, err := p.pwdService.GetAllPasswords(r.Context(), &allPwdDTO)
	if err != nil {
		logger.Log.Errorf("Error work pwdService GetAllPwdDTO: %v", err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	marshaledAllPasswords, err := json.Marshal(allPasswords)
	if err != nil {
		logger.Log.Errorf("Error work pwdService GetAllPwdDTO: %v", err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(marshaledAllPasswords)
	if err != nil {
		logger.Log.Errorf("Error writing response for handler GetAllPwdDTO: %v", err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
}

func (p PwdHandlers) UpdatePassword(w http.ResponseWriter, r *http.Request) {
	updatePwdDTO, err := pwddto.UpdatePwdDTOFromHTTP(r)
	if err != nil {
		logger.Log.Errorf("Error make UpdatePwdDTOFromHTTP: %v", err)
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	err = p.pwdService.UpdatePassword(r.Context(), &updatePwdDTO)
	if err != nil {
		if errors.Is(err, ErrUpdatedRecordNotFound) {
			logger.Log.Errorf("Error work pwdService UpdatePwdDTO: %v", err)
			http.Error(w, "", http.StatusNotFound) // Запись для обновления не найдена
		} else {
			logger.Log.Errorf("Error work pwdService UpdatePwdDTO: %v", err)
			http.Error(w, "", http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	_, err = w.Write([]byte("OK"))
	if err != nil {
		logger.Log.Errorf("Error writing response for handler UpdatePwdDTO: %v", err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
}
