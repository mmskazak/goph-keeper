package auth_http

import (
	"gophKeeper/internal/dto"
	"gophKeeper/internal/logger"
	"gophKeeper/internal/modules/auth/services/auth_service"
	"gophKeeper/internal/modules/auth/services/jwt_service"
	"gophKeeper/internal/service_locator"
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
	if err != nil {
		logger.Log.Errorf("Error authService.Registration: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	cfg, _ := service_locator.GetConfig()
	token, err := jwt_service.GenerateToken(userID, cfg.SecretKey)
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
	cfg, _ := service_locator.GetConfig()
	token, err := jwt_service.GenerateToken(userID, cfg.SecretKey)
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
