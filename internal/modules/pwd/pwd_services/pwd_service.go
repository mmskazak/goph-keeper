package pwd_services

import "github.com/jackc/pgx/v5/pgxpool"

type PwdService struct {
	pool *pgxpool.Pool
}

func NewPwdService(pool *pgxpool.Pool) *PwdService {
	return &PwdService{pool: pool}
}

func (pwd *PwdService) SavePassword(username string, password string) error {
	return nil
}

func (pwd *PwdService) DeletePassword(username string) error {
	return nil
}

func (pwd *PwdService) GetPassword(username string) (string, error) {
	return "secret", nil
}

func (pwd *PwdService) GetAllPasswords(username string, password string) (AllPasswords, error) {
	return AllPasswords{}, nil
}
