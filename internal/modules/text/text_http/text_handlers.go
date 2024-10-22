package text_http

import (
	"encoding/json"
	"errors"
	"gophKeeper/internal/modules/text/text_dto"
	"gophKeeper/internal/modules/text/text_services"
	"net/http"
)

var someSpecificError = errors.New("updated record not found")

type TextHandlers struct {
	textService text_services.ITextService
}

func NewTextHandlersHTTP(service text_services.ITextService) TextHandlers {
	return TextHandlers{
		textService: service,
	}
}

func (p TextHandlers) SaveText(w http.ResponseWriter, r *http.Request) {
	saveTextDTO, err := text_dto.SaveTextDTOFromHTTP(r)
	if err != nil {
		http.Error(w, "", http.StatusBadRequest)
		return
	}
	err = p.textService.SaveText(r.Context(), saveTextDTO)
	if err != nil {
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("OK"))
}

func (p TextHandlers) GetText(w http.ResponseWriter, r *http.Request) {
	getTextDTO, err := text_dto.GetTextDTOFromHTTP(r)
	if err != nil {
		http.Error(w, "", http.StatusBadRequest)
		return
	}
	text, err := p.textService.GetText(r.Context(), getTextDTO)
	if err != nil {
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	jsonResponse, err := json.Marshal(text)
	if err != nil {
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

func (p TextHandlers) DeleteText(w http.ResponseWriter, r *http.Request) {
	deleteTextDTO, err := text_dto.DeleteTextDTOFromHTTP(r)
	if err != nil {
		http.Error(w, "", http.StatusBadRequest)
		return
	}
	err = p.textService.DeleteText(r.Context(), deleteTextDTO)
	if err != nil {
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func (p TextHandlers) GetAllTexts(w http.ResponseWriter, r *http.Request) {
	allTextDTO, err := text_dto.AllTextDTOFromHTTP(r)
	if err != nil {
		http.Error(w, "", http.StatusBadRequest)
		return
	}
	allTexts, err := p.textService.GetAllTexts(r.Context(), allTextDTO)
	if err != nil {
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	marshaledAllTexts, err := json.Marshal(allTexts)
	if err != nil {
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(marshaledAllTexts)
}

func (p TextHandlers) UpdateText(w http.ResponseWriter, r *http.Request) {
	updateTextDTO, err := text_dto.UpdateTextDTOFromHTTP(r)
	if err != nil {
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	err = p.textService.UpdateText(r.Context(), updateTextDTO)
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
