package response

import (
	"gophKeeper/internal/modules/pwd/valueobj"
)

type PwdDTO struct {
	ID          string               `json:"id"`
	Title       string               `json:"title"`
	Description string               `json:"description"`
	Credentials valueobj.Credentials `json:"credentials"`
}
