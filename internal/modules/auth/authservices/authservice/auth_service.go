package authservice

import (
	"context"
	"errors"
	"fmt"
	dto2 "gophKeeper/internal/modules/auth/authdto"
	"gophKeeper/internal/modules/auth/authservices/authhashpwd"

	"github.com/jackc/pgx/v5/pgxpool"
)

type AuthService struct {
	pool *pgxpool.Pool
}

func NewAuthService(pool *pgxpool.Pool) *AuthService {
	return &AuthService{pool: pool}
}

func (a *AuthService) Registration(ctx context.Context, regDTO *dto2.RegistrationDTO) (int, error) {
	hashedPassword, _ := authhashpwd.HashAndStorePassword(regDTO.Password)
	sql := "INSERT INTO users (login, password) VALUES ($1, $2) RETURNING id;"
	row := a.pool.QueryRow(ctx, sql, regDTO.Login, hashedPassword)
	var id int
	err := row.Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("error while inserting user: %w", err)
	}
	return id, nil
}

func (a *AuthService) Login(ctx context.Context, logDTO *dto2.LoginDTO) (int, error) {
	sql := "SELECT id, password FROM users WHERE login = $1;"
	row := a.pool.QueryRow(ctx, sql, logDTO.Login)
	if row == nil {
		return 0, errors.New("user not found")
	}
	var id int
	var hashedPassword string
	err := row.Scan(&id, &hashedPassword)
	if err != nil {
		return 0, fmt.Errorf("error while scanning user: %w", err)
	}
	isLogin := authhashpwd.CheckHashedPassword(hashedPassword, logDTO.Password)
	if !isLogin {
		return 0, errors.New("password incorrect")
	}
	return id, nil
}
