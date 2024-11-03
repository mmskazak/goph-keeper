package authhttp

import (
	"encoding/json"
	"errors"
	"goph-keeper/internal/logger"
	dto "goph-keeper/internal/modules/auth/authdto"
	"goph-keeper/internal/modules/auth/authservices/authjwtservice"
	"goph-keeper/internal/modules/auth/authservices/authservice"
	"net/http"

	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgconn"
)

type AuthHandlers struct {
	authService *authservice.AuthService
	secretKey   string
}

func NewAuthHandlersHTTP(authService *authservice.AuthService, secretKey string) AuthHandlers {
	return AuthHandlers{
		authService: authService,
		secretKey:   secretKey,
	}
}

func (s *AuthHandlers) Login(w http.ResponseWriter, r *http.Request) {
	inDTO, err := dto.LoginDTOFromRequestHTTP(r)
	if err != nil {
		logger.Log.Errorf("Error parsing login request: %v", err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(map[string]string{
			"status":  "error",
			"message": "Invalid request format",
		})
		return
	}

	userID, err := s.authService.Login(r.Context(), inDTO)
	if err != nil {
		logger.Log.Errorf("Error authService.Login: %v", err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		_ = json.NewEncoder(w).Encode(map[string]string{
			"status":  "error",
			"message": "Invalid username or password",
		})
		return
	}

	token, err := authjwtservice.GenerateToken(userID, s.secretKey)
	if err != nil {
		logger.Log.Errorf("Error GenerateToken: %v", err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode(map[string]string{
			"status":  "error",
			"message": "Failed to generate token",
		})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(map[string]string{
		"status": "success",
		"token":  token,
	})
}

func (s *AuthHandlers) Registration(w http.ResponseWriter, r *http.Request) {
	regDTO, err := dto.GetRegistrationDTOFromHTTP(r)
	if err != nil {
		logger.Log.Errorf("Ошибка GetRegistrationDTOFromHTTP: %v", err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(map[string]string{
			"status":  "error",
			"message": "Invalid registration data",
		})
		return
	}

	userID, err := s.authService.Registration(r.Context(), regDTO)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == pgerrcode.UniqueViolation {
			logger.Log.Errorf("Ошибка регистрации: нарушение уникальности")
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			_ = json.NewEncoder(w).Encode(map[string]string{
				"status":  "error",
				"message": "User was already registered",
			})
			return
		}
		logger.Log.Errorf("Ошибка authService.Registration: %v", err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode(map[string]string{
			"status":  "error",
			"message": "Internal server error",
		})
		return
	}

	// Генерация и отправка токена
	token, err := authjwtservice.GenerateToken(userID, s.secretKey)
	if err != nil {
		logger.Log.Errorf("Ошибка GenerateToken: %v", err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode(map[string]string{
			"status":  "error",
			"message": "Failed to generate token",
		})
		return
	}

	// Успешный ответ при регистрации
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(map[string]string{
		"status":  "success",
		"message": "User registered successfully",
		"jwt":     "Bearer " + token,
	})
}
