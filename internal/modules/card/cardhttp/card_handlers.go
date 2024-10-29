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
		logger.Log.Errorf("Error creating saveCardDTO: %v", err)
		http.Error(w, "", http.StatusBadRequest)
		return
	}
	err = ch.cardService.SaveCard(r.Context(), &saveCardDTO)
	if err != nil {
		logger.Log.Errorf("Error run cardService.SaveCard: %v", err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	_, err = w.Write([]byte("Card saved successfully"))
	if err != nil {
		logger.Log.Errorf("failed to write resopnse for save card: %v", err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
}

func (ch CardHandlers) GetCard(w http.ResponseWriter, r *http.Request) {
	getCardDTO, err := carddto.GetCardDTOFromHTTP(r)
	if err != nil {
		logger.Log.Errorf("Error creating getCardDTO: %v", err)
		http.Error(w, "", http.StatusBadRequest)
		return
	}
	card, err := ch.cardService.GetCard(r.Context(), &getCardDTO)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			logger.Log.Errorf("Card not found: %v", err)
			http.Error(w, "", http.StatusNotFound)
			return
		}
		logger.Log.Errorf("Error run cardService.GetCard: %v", err)
		http.Error(w, "Error retrieving card", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(card)
	if err != nil {
		logger.Log.Errorf("failed to encode for get card: %v", err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
}

func (ch CardHandlers) UpdateCard(w http.ResponseWriter, r *http.Request) {
	updateCardDTO, err := carddto.UpdateCardDTOFromHTTP(r)
	if err != nil {
		logger.Log.Errorf("Error creating updateCardDTO: %v", err)
		http.Error(w, "", http.StatusBadRequest)
		return
	}
	err = ch.cardService.UpdateCard(r.Context(), &updateCardDTO)
	if err != nil {
		logger.Log.Errorf("Error run cardService.UpdateCard: %v", err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	_, err = w.Write([]byte("Card updated successfully"))
	if err != nil {
		logger.Log.Errorf("failed to write response for update card: %v", err)
		http.Error(w, "", http.StatusInternalServerError)

		return
	}
}

func (ch CardHandlers) DeleteCard(w http.ResponseWriter, r *http.Request) {
	deleteCardDTO, err := carddto.DeleteCardDTOFromHTTP(r) // Предполагается, что эта функция реализована
	if err != nil {
		logger.Log.Errorf("Error creating deleteCardDTO: %v", err)
		http.Error(w, "", http.StatusBadRequest)
		return
	}
	err = ch.cardService.DeleteCard(r.Context(), &deleteCardDTO)
	if err != nil {
		logger.Log.Errorf("Error run cardService.DeleteCard: %v", err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (ch CardHandlers) GetAllCards(w http.ResponseWriter, r *http.Request) {
	getAllCardsDTO, err := carddto.GetAllCardsDTOFromHTTP(r)
	if err != nil {
		logger.Log.Errorf("Error creating getAllCardsDTO: %v", err)
		http.Error(w, "", http.StatusBadRequest)
		return
	}
	cards, err := ch.cardService.GetAllCards(r.Context(), &getAllCardsDTO)
	if err != nil {
		logger.Log.Errorf("Error run cardService.GetAllCards: %v", err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(cards)
	if err != nil {
		logger.Log.Errorf("failed encode cards for get all: %v", err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
}
