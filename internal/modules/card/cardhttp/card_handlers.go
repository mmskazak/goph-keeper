package cardhttp

import (
	"encoding/json"
	"errors"
	"gophKeeper/internal/logger"
	"gophKeeper/internal/modules/card/carddto"
	"gophKeeper/internal/modules/card/cardservices"
	"net/http"

	"github.com/jackc/pgx/v5"
)

type CardHandlers struct {
	cardService cardservices.ICardService
}

func NewCardHandlersHTTP(service cardservices.ICardService) CardHandlers {
	return CardHandlers{cardService: service}
}

func (ch CardHandlers) SaveCard(w http.ResponseWriter, r *http.Request) {
	saveCardDTO, err := carddto.SaveCardDTOFromHTTP(r)
	if err != nil {
		http.Error(w, "", http.StatusBadRequest)
		return
	}
	err = ch.cardService.SaveCard(r.Context(), saveCardDTO)
	if err != nil {
		http.Error(w, "Error saving card", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	_, err = w.Write([]byte("Card saved successfully"))
	if err != nil {
		http.Error(w, "", http.StatusInternalServerError)
		logger.Log.Errorf("failed to write resopnse for save card: %v", err)
		return
	}
}

func (ch CardHandlers) GetCard(w http.ResponseWriter, r *http.Request) {
	getCardDTO, err := carddto.GetCardDTOFromHTTP(r)
	if err != nil {
		http.Error(w, "", http.StatusBadRequest)
		return
	}
	card, err := ch.cardService.GetCard(r.Context(), getCardDTO)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			http.Error(w, "Card not found", http.StatusNotFound)
			return
		}
		http.Error(w, "Error retrieving card", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(card)
	if err != nil {
		http.Error(w, "", http.StatusInternalServerError)
		logger.Log.Errorf("failed to encode for get card: %v", err)
		return
	}
}

func (ch CardHandlers) UpdateCard(w http.ResponseWriter, r *http.Request) {
	updateCardDTO, err := carddto.UpdateCardDTOFromHTTP(r) // Предполагается, что эта функция реализована
	if err != nil {
		http.Error(w, "", http.StatusBadRequest)
		return
	}
	err = ch.cardService.UpdateCard(r.Context(), updateCardDTO)
	if err != nil {
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	_, err = w.Write([]byte("Card updated successfully"))
	if err != nil {
		http.Error(w, "", http.StatusInternalServerError)
		logger.Log.Errorf("failed to write response for update card: %v", err)
		return
	}
}

func (ch CardHandlers) DeleteCard(w http.ResponseWriter, r *http.Request) {
	deleteCardDTO, err := carddto.DeleteCardDTOFromHTTP(r) // Предполагается, что эта функция реализована
	if err != nil {
		http.Error(w, "", http.StatusBadRequest)
		return
	}
	err = ch.cardService.DeleteCard(r.Context(), deleteCardDTO)
	if err != nil {
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent) // 204 No Content
}

func (ch CardHandlers) GetAllCards(w http.ResponseWriter, r *http.Request) {
	getAllCardsDTO, err := carddto.GetAllCardsDTOFromHTTP(r) // Предполагается, что эта функция реализована
	if err != nil {
		http.Error(w, "", http.StatusBadRequest)
		return
	}
	cards, err := ch.cardService.GetAllCards(r.Context(), getAllCardsDTO)
	if err != nil {
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(cards)
	if err != nil {
		http.Error(w, "", http.StatusInternalServerError)
		logger.Log.Errorf("failed encode cards for get all: %v", err)
		return
	}
}
