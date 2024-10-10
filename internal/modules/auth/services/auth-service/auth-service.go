package auth_service

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"gophKeeper/internal/dto"
)

type AuthService struct {
	pool *pgxpool.Pool
}

func NewAuthService(pool *pgxpool.Pool) *AuthService {
	return &AuthService{pool: pool}
}

func (a *AuthService) Registration(ctx context.Context, regDTO *dto.RegistrationDTO) error {
	sql := "INSERT INTO users (login, password) VALUES ($1, $2);"
	_, err := a.pool.Exec(ctx, sql, regDTO.Login, regDTO.Password)
	if err != nil {
		return fmt.Errorf("error registration user from service: %w", err)
	}
	return nil
}

func (a *AuthService) Login(dto dto.LoginDTO) {

}

func (a *AuthService) Logout(dto dto.LogoutDTO) {

}
