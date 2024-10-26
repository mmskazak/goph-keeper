package texthttp

import (
	"encoding/json"
	"errors"
	"gophKeeper/internal/logger"
	"gophKeeper/internal/modules/text/textdto"
	"gophKeeper/internal/modules/text/textservices"
	"net/http"
)

var errSomeSpecificError = errors.New("error updated record not found")

type TextHandlers struct {
	textService textservices.ITextService
}

func NewTextHandlersHTTP(service textservices.ITextService) TextHandlers {
	return TextHandlers{
		textService: service,
	}
}

func (p TextHandlers) SaveText(w http.ResponseWriter, r *http.Request) {
	saveTextDTO, err := textdto.SaveTextDTOFromHTTP(r)
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
	_, err = w.Write([]byte("OK"))
	if err != nil {
		logger.Log.Error(err.Error())
	}
}

func (p TextHandlers) GetText(w http.ResponseWriter, r *http.Request) {
	getTextDTO, err := textdto.GetTextDTOFromHTTP(r)
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
	_, err = w.Write(jsonResponse)
	if err != nil {
		logger.Log.Errorf("failed to write response: %s", err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
}

func (p TextHandlers) DeleteText(w http.ResponseWriter, r *http.Request) {
	deleteTextDTO, err := textdto.DeleteTextDTOFromHTTP(r)
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
	_, err = w.Write([]byte("OK"))
	if err != nil {
		logger.Log.Errorf("failed to write response: %s", err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
}

func (p TextHandlers) GetAllTexts(w http.ResponseWriter, r *http.Request) {
	allTextDTO, err := textdto.AllTextDTOFromHTTP(r)
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
		logger.Log.Errorf("failed to marshal all texts: %s", err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(marshaledAllTexts)
	if err != nil {
		logger.Log.Errorf("error GetAllTexts writing response: %s", err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
}

func (p TextHandlers) UpdateText(w http.ResponseWriter, r *http.Request) {
	updateTextDTO, err := textdto.UpdateTextDTOFromHTTP(r)
	if err != nil {
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	err = p.textService.UpdateText(r.Context(), updateTextDTO)
	if err != nil {
		if errors.Is(err, errSomeSpecificError) {
			http.Error(w, "", http.StatusNotFound) // Запись для обновления не найдена
		} else {
			http.Error(w, "", http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	_, err = w.Write([]byte("OK"))
	if err != nil {
		logger.Log.Errorf("failed to write response: %s", err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
}
