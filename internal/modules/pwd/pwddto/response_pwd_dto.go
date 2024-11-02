package pwddto

import (
	"goph-keeper/internal/modules/pwd/valueobj"
)

type ResponsePwdDTO struct {
	PwdID       string               `json:"id"`
	Title       string               `json:"title"`
	Description string               `json:"description"`
	Credentials valueobj.Credentials `json:"credentials"`
}
