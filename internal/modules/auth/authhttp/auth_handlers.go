package authhttp

import (
	"encoding/json"
	"errors"
	dto "goph-keeper/internal/modules/auth/authdto"
	"goph-keeper/internal/modules/auth/authservices/authjwtservice"
	"goph-keeper/internal/modules/auth/authservices/authservice"
	"net/http"

	"go.uber.org/zap"

	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgconn"
)

type AuthHandlers struct {
	zapLogger   *zap.SugaredLogger
	authService *authservice.AuthService
	secretKey   string
}

func NewAuthHandlersHTTP(authService *authservice.AuthService,
	secretKey string,
	zapLogger *zap.SugaredLogger,
) AuthHandlers {
	return AuthHandlers{
		authService: authService,
		secretKey:   secretKey,
		zapLogger:   zapLogger,
	}
}

func (s *AuthHandlers) Login(w http.ResponseWriter, r *http.Request) {
	inDTO, err := dto.LoginDTOFromRequestHTTP(r)
	if err != nil {
		s.zapLogger.Errorf("Error parsing login request: %v", err)
		http.Error(w, "Invalid request format", http.StatusBadRequest)
		return
	}

	userID, err := s.authService.Login(r.Context(), inDTO)
	if err != nil {
		s.zapLogger.Errorf("Error authService.Username: %v", err)
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	token, err := authjwtservice.GenerateToken(userID, s.secretKey)
	if err != nil {
		s.zapLogger.Errorf("Error GenerateToken: %v", err)
		http.Error(w, "Failed to generate token", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(map[string]string{
		"jwt": "Bearer " + token,
	})
}

func (s *AuthHandlers) Registration(w http.ResponseWriter, r *http.Request) {
	regDTO, err := dto.GetRegistrationDTOFromHTTP(r)
	if err != nil {
		s.zapLogger.Errorf("Ошибка GetRegistrationDTOFromHTTP: %v", err)
		http.Error(w, "Invalid registration data", http.StatusBadRequest)
		return
	}

	userID, err := s.authService.Registration(r.Context(), regDTO)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == pgerrcode.UniqueViolation {
			s.zapLogger.Errorf("Ошибка регистрации: нарушение уникальности")
			http.Error(w, "User was already registered", http.StatusBadRequest)
			return
		}
		s.zapLogger.Errorf("Ошибка authService.Registration: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	// Генерация и отправка токена
	token, err := authjwtservice.GenerateToken(userID, s.secretKey)
	if err != nil {
		s.zapLogger.Errorf("Ошибка GenerateToken: %v", err)
		http.Error(w, "Failed to generate token", http.StatusInternalServerError)
		return
	}

	// Успешный ответ при регистрации
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(map[string]string{
		"jwt": "Bearer " + token,
	})
}
