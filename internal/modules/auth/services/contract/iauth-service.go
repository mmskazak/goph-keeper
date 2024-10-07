package contract

import "GophKeeper/internal/dto"

type IAuthService interface {
	Registration(dto dto.RegistrationDTO)
	Login(dto dto.LoginDTO)
	Logout(dto dto.LogoutDTO)
}
