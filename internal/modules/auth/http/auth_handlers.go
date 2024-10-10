package auth_http

import (
	"gophKeeper/internal/dto"
	"gophKeeper/internal/logger"
	"gophKeeper/internal/modules/auth/services/auth-service"
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
	logger.Log.Infoln("DTO: ", regDTO)
	if err != nil {
		logger.Log.Errorf("Error GetRegistrationDTOFromHTTP: %v", err)
	}
	logger.Log.Debugf("authService -> %v", s.authService)
	err = s.authService.Registration(r.Context(), regDTO)
	if err != nil {
		logger.Log.Errorf("Error authService.Registration: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("ok"))
}

func (s *AuthHandlers) Login(w http.ResponseWriter, r *http.Request) {
	inDTO := dto.GetLoginDTOFromHTTP(r)
	s.authService.Login(inDTO)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))
}

func (s *AuthHandlers) Logout(w http.ResponseWriter, r *http.Request) {
	outDTO := dto.GetLogoutDTOFromHTTP(r)
	s.authService.Logout(outDTO)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))
}
