package auth_service

import (
	"GophKeeper/internal/dto"
	"GophKeeper/internal/repositories/contract"
)

type AuthService struct {
	repo contract.IRepository
}

func (a *AuthService) Registration(dto dto.RegistrationDTO) {

}

func (a *AuthService) Login(dto dto.LoginDTO) {

}

func (a *AuthService) Logout(dto dto.LogoutDTO) {

}
