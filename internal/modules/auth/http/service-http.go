package servive_http

import (
	"gophKeeper/internal/dto"
	"gophKeeper/internal/logger"
	"gophKeeper/internal/modules/auth/services/auth-service"
	"net/http"
)

type ServiceHTTP struct {
	authService auth_service.AuthService
}

func NewAuthServiceHTTP() *ServiceHTTP {
	return &ServiceHTTP{}
}

func (s *ServiceHTTP) Registration(w http.ResponseWriter, r *http.Request) {
	regDTO, err := dto.GetRegistrationDTOFromHTTP(r)
	if err != nil {
		logger.Log.Errorf("DTO Error: %v", err)
	}
	s.authService.Registration(regDTO)
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("ok"))
}

func (s *ServiceHTTP) Login(w http.ResponseWriter, r *http.Request) {
	inDTO := dto.GetLoginDTOFromHTTP(r)
	s.authService.Login(inDTO)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))
}

func (s *ServiceHTTP) Logout(w http.ResponseWriter, r *http.Request) {
	outDTO := dto.GetLogoutDTOFromHTTP(r)
	s.authService.Logout(outDTO)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))
}
