package dto

import "net/http"

type LogoutDTO struct {
}

func GetLogoutDTOFromHTTP(_ *http.Request) LogoutDTO {
	return LogoutDTO{}
}

func GetLogoutDTOFromGRPC() LogoutDTO {
	return LogoutDTO{}
}
