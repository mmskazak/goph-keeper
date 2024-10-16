package pwd_http

import (
	"encoding/json"
	"gophKeeper/internal/modules/text/text_dto"
	"gophKeeper/internal/modules/text/text_services"
	"net/http"
)

type TextHandlers struct {
	textService text_services.ITextService
}

func NewTextHandlersHTTP(service text_services.ITextService) TextHandlers {
	return TextHandlers{
		textService: service,
	}
}

func (p TextHandlers) SaveText(w http.ResponseWriter, r *http.Request) {
	savePwdDTO, err := text_dto.SaveTextDTOFromHTTP(r)
	if err != nil {
		http.Error(w, "", http.StatusBadRequest)
		return
	}
	err = p.textService.SaveText(r.Context(), savePwdDTO)
	if err != nil {
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("OK"))
}

func (p TextHandlers) GetPassword(w http.ResponseWriter, r *http.Request) {
	getTextDTO, err := text_dto.GetTextDTOFromHTTP(r)
	if err != nil {
		http.Error(w, "", http.StatusBadRequest)
		return
	}
	password, err := p.textService.GetText(r.Context(), getTextDTO)
	if err != nil {
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(password))
}

func (p TextHandlers) DeletePassword(w http.ResponseWriter, r *http.Request) {
	deletePwdDTO, err := text_dto.DeletePwdDTOFromHTTP(r)
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

func (p TextHandlers) GetAllPasswords(w http.ResponseWriter, r *http.Request) {
	allPwdDTO, err := text_dto.AllTextsDTOFromHTTP(r)
	if err != nil {
		http.Error(w, "", http.StatusBadRequest)
		return
	}
	allPasswords, err := p.textService.GetAllTexts(r.Context(), allPwdDTO)
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
