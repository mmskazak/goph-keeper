package authhttp

import (
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

func (s *AuthHandlers) Registration(w http.ResponseWriter, r *http.Request) {
	regDTO, err := dto.GetRegistrationDTOFromHTTP(r)
	if err != nil {
		logger.Log.Errorf("Error GetRegistrationDTOFromHTTP: %v", err)
	}
	userID, err := s.authService.Registration(r.Context(), regDTO)
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) && pgErr.Code == pgerrcode.UniqueViolation {
		logger.Log.Errorf(`Registration failed due to unique violation`)
		http.Error(w, "Пользователь ранне уже был зарегистрирован", http.StatusBadRequest)
		return
	}
	if err != nil {
		logger.Log.Errorf("Error authService.Registration: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	token, err := authjwtservice.GenerateToken(userID, s.secretKey)
	if err != nil {
		logger.Log.Errorf("Error GenerateToken: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Authorization", "Bearer "+token)

	w.WriteHeader(http.StatusCreated)
	_, err = w.Write([]byte("ok"))
	if err != nil {
		logger.Log.Errorf("Error write: %v", err)
	}
}

func (s *AuthHandlers) Login(w http.ResponseWriter, r *http.Request) {
	inDTO, _ := dto.LoginDTOFromRequestHTTP(r)
	userID, err := s.authService.Login(r.Context(), inDTO)
	if err != nil {
		logger.Log.Errorf("Error authService.Login: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	token, err := authjwtservice.GenerateToken(userID, s.secretKey)
	if err != nil {
		logger.Log.Errorf("Error GenerateToken: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Authorization", "Bearer "+token)

	w.WriteHeader(http.StatusOK)
	_, err = w.Write([]byte("ok"))
	if err != nil {
		logger.Log.Errorf("Error write: %v", err)
	}
}
