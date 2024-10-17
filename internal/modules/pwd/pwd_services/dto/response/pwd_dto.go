package response

import "gophKeeper/internal/modules/pwd/pwd_services/dto/common"

type PwdDTO struct {
	UserID      int                `json:"user_id"`
	Title       string             `json:"title"`
	Description string             `json:"description"`
	Credentials common.Credentials `json:"credentials"`
}
