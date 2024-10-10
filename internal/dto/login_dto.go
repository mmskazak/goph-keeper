package dto

import "net/http"

type LoginDTO struct {
}

func GetLoginDTOFromHTTP(_ *http.Request) LoginDTO {
	return LoginDTO{}
}

func GetLoginDTOFromGRPC() LoginDTO {
	return LoginDTO{}
}
