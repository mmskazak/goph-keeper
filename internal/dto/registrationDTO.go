package dto

import "net/http"

type RegistrationDTO struct {
}

func GetRegistrationDTOFromHTTP(_ *http.Request) RegistrationDTO {
	return RegistrationDTO{}
}

func GetRegistrationDTOFromGRPC() RegistrationDTO {
	return RegistrationDTO{}
}
