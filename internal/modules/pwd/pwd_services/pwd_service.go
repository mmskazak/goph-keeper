package pwd_services

import "github.com/jackc/pgx/v5/pgxpool"

type PwdService struct {
	pool *pgxpool.Pool
}

func NewPwdService(pool *pgxpool.Pool) *PwdService {
	return &PwdService{pool: pool}
}
