package pwd_http

import (
	"encoding/json"
	"gophKeeper/internal/modules/card/card_services"
	"gophKeeper/internal/modules/text/text_dto"
	"net/http"
)

type CardHandlers struct {
	cardService card_services.ICardService
}

func NewCardHandlersHTTP(service card_services.ICardService) CardHandlers {
	return CardHandlers{
		cardService: service,
	}
}

func (p CardHandlers) SaveCard(w http.ResponseWriter, r *http.Request) {
	savePwdDTO, err := text_dto.SaveTextDTOFromHTTP(r)
	if err != nil {
		http.Error(w, "", http.StatusBadRequest)
		return
	}
	err = p.cardService.SaveCard(r.Context(), savePwdDTO)
	if err != nil {
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("OK"))
}

func (p CardHandlers) GetCard(w http.ResponseWriter, r *http.Request) {
	getTextDTO, err := text_dto.GetTextDTOFromHTTP(r)
	if err != nil {
		http.Error(w, "", http.StatusBadRequest)
		return
	}
	password, err := p.cardService.GetCard(r.Context(), getTextDTO)
	if err != nil {
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(password))
}

func (p CardHandlers) DeleteCard(w http.ResponseWriter, r *http.Request) {
	deletePwdDTO, err := text_dto.DeletePwdDTOFromHTTP(r)
	if err != nil {
		http.Error(w, "", http.StatusBadRequest)
		return
	}
	err = p.cardService.DeleteCard(r.Context(), deletePwdDTO)
	if err != nil {
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func (p CardHandlers) GetAllCards(w http.ResponseWriter, r *http.Request) {
	allPwdDTO, err := text_dto.AllTextsDTOFromHTTP(r)
	if err != nil {
		http.Error(w, "", http.StatusBadRequest)
		return
	}
	allPasswords, err := p.cardService.GetAllCards(r.Context(), allPwdDTO)
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
