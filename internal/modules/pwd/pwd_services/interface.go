package pwd_services

type AllPasswords struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type IPwdService interface {
	SavePassword(username string, password string) error
	DeletePassword(username string) error
	GetPassword(username string) (string, error)
	GetAllPasswords(username string) (AllPasswords, error)
}
