package auth_http

import (
	"errors"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgconn"
	"gophKeeper/internal/dig"
	"gophKeeper/internal/logger"
	dto "gophKeeper/internal/modules/auth/auth_dto"
	"gophKeeper/internal/modules/auth/auth_services/auth_jwt_service"
	"gophKeeper/internal/modules/auth/auth_services/auth_service"
	"net/http"
)

type AuthHandlers struct {
	authService *auth_service.AuthService
}

func NewAuthHandlersHTTP(authService *auth_service.AuthService) AuthHandlers {
	return AuthHandlers{
		authService: authService,
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
	cfg, _ := dig.GetConfig()
	token, err := auth_jwt_service.GenerateToken(userID, cfg.SecretKey)
	if err != nil {
		logger.Log.Errorf("Error GenerateToken: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Authorization", "Bearer "+token)

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("ok"))
}

func (s *AuthHandlers) Login(w http.ResponseWriter, r *http.Request) {
	inDTO, _ := dto.GetLoginDTOFromHTTP(r)
	userID, err := s.authService.Login(r.Context(), inDTO)
	if err != nil {
		logger.Log.Errorf("Error authService.Login: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	cfg, _ := dig.GetConfig()
	token, err := auth_jwt_service.GenerateToken(userID, cfg.SecretKey)
	if err != nil {
		logger.Log.Errorf("Error GenerateToken: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Authorization", "Bearer "+token)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))
}

func (s *AuthHandlers) Logout(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Authorization", "")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("logout"))
}
