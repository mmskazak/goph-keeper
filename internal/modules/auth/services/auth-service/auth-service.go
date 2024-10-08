package auth_service

import (
	"gophKeeper/internal/dto"
	"gophKeeper/internal/repositories/contract"
)

type AuthService struct {
	repo contract.IRepository
}

func NewAuthService(repo contract.IRepository) *AuthService {
	return &AuthService{repo: repo}
}

func (a *AuthService) Registration(dto dto.RegistrationDTO) {

}

func (a *AuthService) Login(dto dto.LoginDTO) {

}

func (a *AuthService) Logout(dto dto.LogoutDTO) {

}
