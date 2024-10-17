package response

import (
	"gophKeeper/internal/modules/pwd/pwd_services/value_obj"
)

type PwdDTO struct {
	ID          string                `json:"id"`
	Title       string                `json:"title"`
	Description string                `json:"description"`
	Credentials value_obj.Credentials `json:"credentials"`
}
