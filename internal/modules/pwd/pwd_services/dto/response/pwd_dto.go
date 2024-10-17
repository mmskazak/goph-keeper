package response

import (
	"gophKeeper/internal/modules/pwd/pwd_services/value_obj"
)

type PwdDTO struct {
	Title       string                `json:"title"`
	Description string                `json:"description"`
	Credentials value_obj.Credentials `json:"credentials"`
}
