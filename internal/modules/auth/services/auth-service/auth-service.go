package auth_service

import (
	"gophKeeper/internal/dto"
	"gophKeeper/internal/logger"
	"gophKeeper/internal/repositories/contract"
)

type AuthService struct {
	repo contract.IRepository
}

func NewAuthService(repo contract.IRepository) *AuthService {
	return &AuthService{repo: repo}
}

func (a *AuthService) Registration(_ *dto.RegistrationDTO) {
	logger.Log.Debugln("Регистрация пользователя.")
}

func (a *AuthService) Login(dto dto.LoginDTO) {

}

func (a *AuthService) Logout(dto dto.LogoutDTO) {

}
